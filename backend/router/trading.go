package router

import (
	"errors"
	"time"

	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/DBoyara/invest-helper/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

var TRADING fiber.Router

const layout = "2006-01-02"
const (
	SELL     string = "sell"
	BUY             = "buy"
	FixPRICE        = "fix_price"
	PERCENT         = "percent"
)

var (
	errNotValidTradeType      = errors.New("not valid trade type")
	errNotValidCommissionType = errors.New("not valid commission type")
)

func SetupTradingRoutes() {
	TRADING.Get("", GetTradingLogs)
	TRADING.Get("/commissions", GetCommissions)
	TRADING.Post("", CreateTradingLog)
	TRADING.Put("/close", CloseDeals)
}

func CreateTradingLog(c *fiber.Ctx) error {
	log := &models.TradingLog{}
	if err := c.BodyParser(log); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if log.Type != SELL && log.Type != BUY {
		return c.Status(400).SendString(errNotValidTradeType.Error())
	}
	if log.CommissionType != FixPRICE && log.CommissionType != PERCENT {
		return c.Status(400).SendString(errNotValidCommissionType.Error())
	}

	log.Amount = log.Price * float64(log.Count) * float64(log.Lot)
	if log.CommissionType == FixPRICE {
		log.CommissionAmount = log.Commission
	} else if log.CommissionType == PERCENT {
		commissionAmount := log.Commission * log.Amount / 100
		log.CommissionAmount = toFixed(commissionAmount, 2)
	}

	res, err := repository.CreateTradeLog(log)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(201).JSON(res)
}

func GetTradingLogs(c *fiber.Ctx) error {
	dateStart := c.Query("dateStart")
	dateEnd := c.Query("dateEnd")

	if dateStart == "" && dateEnd == "" {
		logs, err := repository.GetTradeLogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(logs)
	}

	DateStart, err := time.Parse(layout, dateStart)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	DateEnd, err := time.Parse(layout, dateEnd)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	logs, err := repository.GetTradeLogsByDatetime(DateStart, DateEnd)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(logs)
}

type IdList struct {
	Ids    []string `json:"ids"`
	IsOpen bool     `json:"is_open"`
}

func CloseDeals(c *fiber.Ctx) error {
	idList := &IdList{}
	if err := c.BodyParser(idList); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := repository.UpdateLogsStatusByID(idList.Ids, idList.IsOpen); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}

func GetCommissions(c *fiber.Ctx) error {
	com, err := repository.GetCommissions()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(com)
}
