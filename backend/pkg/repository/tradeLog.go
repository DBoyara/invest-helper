package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/jackc/pgx/v4"
)

const (
	selectedFields string = "id, datetime, tiker, type, is_open, price, count, lot, amount, commission, commission_amount, tiker_type, currency"
	layout                = "2006-01-02"
	Equity                = "equity"
	RUB                   = "rub"
)

func scanLog(rows pgx.Row, model *models.TradingLog) error {
	return rows.Scan(
		&model.Id,
		&model.Datetime,
		&model.Tiker,
		&model.Type,
		&model.IsOpen,
		&model.Price,
		&model.Count,
		&model.Lot,
		&model.Amount,
		&model.Commission,
		&model.CommissionAmount,
		&model.TikerType,
		&model.Currency,
	)
}

func constructQuery(dateStart, dateEnd, showOpen, tikerType, currency string) (string, error) {
	selectQuery := fmt.Sprintf("select %s from trading_logs", selectedFields)
	filterQuery := fmt.Sprintf("where is_open = %s and tiker_type = '%s' and currency = '%s'", showOpen, tikerType, currency)
	orderQuery := "order by datetime desc"
	dateFilter := ""

	if dateStart != "" && dateEnd != "" {
		ds, err := time.Parse(layout, dateStart)
		if err != nil {
			return "", err
		}
		de, err := time.Parse(layout, dateEnd)
		if err != nil {
			return "", err
		}
		dateFilter = fmt.Sprintf(" and datetime between '%v' and '%v'", ds, de)
	}
	query := fmt.Sprintf("%s %s%s %s", selectQuery, filterQuery, dateFilter, orderQuery)
	return query, nil
}

func CreateTradeLog(log *models.TradingLog) (*models.TradingLog, error) {
	if log.TikerType == "" {
		log.TikerType = Equity
	}
	if log.Currency == "" {
		log.Currency = RUB
	}

	err := connPool.BeginFunc(context.Background(), func(t pgx.Tx) error {
		return t.QueryRow(
			context.Background(),
			`insert into trading_logs 
				(tiker, type, price, count, lot, amount, commission, commission_amount, tiker_type, currency) 
			values 
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
			returning id`,
			log.Tiker,
			log.Type,
			log.Price,
			log.Count,
			log.Lot,
			log.Amount,
			log.Commission,
			log.CommissionAmount,
			log.TikerType,
			log.Currency,
		).Scan(&log.Id)
	})

	return log, err
}

func GetTradeLogs(dateStart, dateEnd, showOpen, tikerType, currency string) ([]*models.TradingLog, error) {
	var res []*models.TradingLog

	query, err := constructQuery(dateStart, dateEnd, showOpen, tikerType, currency)
	if err != nil {
		return nil, err
	}

	rows, err := connPool.Query(
		context.Background(),
		query,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	for rows.Next() {
		tmp := &models.TradingLog{}

		if err = scanLog(rows, tmp); err != nil {
			return res, err
		}

		res = append(res, tmp)
	}

	return res, nil
}

func UpdateLogsStatusByID(ids []string, isOpen bool) error {
	idsToStr := strings.Join(ids, ",")
	str := fmt.Sprintf("update trading_logs set is_open = %v where id in (%v)", isOpen, idsToStr)

	_, err := connPool.Exec(
		context.Background(),
		str,
	)

	if err == pgx.ErrNoRows {
		return nil
	}

	return err
}
