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
	TRADING.Post("", CreateTradingLog)
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
		log.CommissionAmount = log.Commission * log.Amount / 100
	}

	_, err := repository.CreateTradeLog(log)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.SendStatus(201)
}

type DateInterval struct {
	DateStart time.Time `query:"dateStart"`
	DateEnd   time.Time `query:"dateEnd"`
}

func GetTradingLogs(c *fiber.Ctx) error {
	logs, err := repository.GetTradeLogs()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(logs)
}
