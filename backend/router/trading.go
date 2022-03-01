package router

import (
	"context"
	"errors"
	"github.com/DBoyara/invest-helper/common"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"strings"
	"time"

	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/DBoyara/invest-helper/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

var TRADING fiber.Router

var (
	errNotValidTradeType      = errors.New("not valid trade type")
	errNotValidCommissionType = errors.New("not valid commission type")
)

func SetupTradingRoutes() {
	TRADING.Get("", GetTradingLogs)
	TRADING.Get("/commissions", GetCommissions)
	TRADING.Get("/tiker", GetTikerInfo)
	TRADING.Get("/summary/:type", GetSummary)
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

	log.Amount = toFixed(log.Price*float64(log.Count)*float64(log.Lot), 2)
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
	dateStart := c.Query("dateStart", "")
	dateEnd := c.Query("dateEnd", "")
	showOpen := c.Query("showOpen", "true")
	currency := c.Query("showOpen", "rub")
	tikerType := c.Query("tikerType", "equity")

	logs, err := repository.GetTradeLogs(dateStart, dateEnd, showOpen, tikerType, currency)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(logs)
}

func GetTikerInfo(c *fiber.Ctx) error {
	tiker := c.Query("tiker")
	t := strings.ToUpper(tiker)
	settings := common.GetSettings()

	client := sdk.NewRestClient(settings.TinkoffToken)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	instruments, err := client.InstrumentByTicker(ctx, t)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(instruments)
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

func GetSummary(c *fiber.Ctx) error {
	dateStart := c.Query("dateStart", "")
	dateEnd := c.Query("dateEnd", "")
	showOpen := c.Query("showOpen", "false")
	currency := c.Query("showOpen", "rub")
	tikerType := c.Params("tikerType", "equity")

	logs, err := repository.GetTradeLogs(dateStart, dateEnd, showOpen, tikerType, currency)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	income := countEquitySummary(logs)

	return c.Status(200).JSON(income)
}
