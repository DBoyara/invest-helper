package repository

import (
	"context"
	"time"

	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/jackc/pgx/v4"
)

func CreateTradeLog(log *models.TradingLog) (int, error) {
	err := connPool.BeginFunc(context.Background(), func(t pgx.Tx) error {
		return t.QueryRow(
			context.Background(),
			`insert into trading_logs 
				(tiker, type, price, count, lot, amount, commission, commission_amount) 
			values 
				($1, $2, $3, $4, $5, $6, $7, $8) 
			returning id`,
			log.Tiker,
			log.Type,
			log.Price,
			log.Count,
			log.Lot,
			log.Amount,
			log.Commission,
			log.CommissionAmount,
		).Scan(&log.Id)
	})

	return log.Id, err
}

func GetTradeLogs() ([]*models.TradingLog, error) {
	var res []*models.TradingLog

	rows, err := connPool.Query(
		context.Background(),
		"select id, datetime, tiker, type, price, count, lot, amount, commission, commission_amount from trading_logs order by datetime",
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

func GetTradeLogsByDatetime(startDate, endDate time.Time) ([]*models.TradingLog, error) {
	var res []*models.TradingLog

	rows, err := connPool.Query(
		context.Background(),
		"select id, datetime, tiker, type, price, count, lot, amount, commission, commission_amount from trading_logs where datetime between $1 and $2 order by datetime",
		startDate,
		endDate,
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

func scanLog(rows pgx.Row, model *models.TradingLog) error {
	return rows.Scan(
		&model.Id,
		&model.Datetime,
		&model.Tiker,
		&model.Type,
		&model.Price,
		&model.Count,
		&model.Lot,
		&model.Amount,
		&model.Commission,
		&model.CommissionAmount)
}
