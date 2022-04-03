package router

import (
	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/DBoyara/invest-helper/pkg/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var FUTURES fiber.Router

func SetupFuturesRoutes() {
	FUTURES.Get("", GetFutures)
	FUTURES.Get("/summary", GetFuturesSummary)
	FUTURES.Post("", CreateFutures)
	FUTURES.Put(":id", UpdateFutures)
}

func CreateFutures(c *fiber.Ctx) error {
	futures := &models.Futures{}
	if err := c.BodyParser(futures); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if futures.CommissionType != FixPRICE && futures.CommissionType != PERCENT {
		return c.Status(400).SendString(errNotValidCommissionType.Error())
	}

	futures.Amount = toFixed(futures.WarrantyProvision*float64(futures.Count), 2)
	if futures.CommissionType == FixPRICE {
		futures.CommissionAmount = futures.Commission
	} else if futures.CommissionType == PERCENT {
		commissionAmount := futures.Commission * futures.Amount / 100
		futures.CommissionAmount = toFixed(commissionAmount, 2)
	}

	res, err := repository.CreateFuture(futures)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(201).JSON(res)
}

func GetFutures(c *fiber.Ctx) error {
	dateStart := c.Query("dateStart", "")
	dateEnd := c.Query("dateEnd", "")
	showOpen := c.Query("showOpen", "true")

	logs, err := repository.GetFutures(dateStart, dateEnd, showOpen)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(logs)
}

type MarginBody struct {
	Margin float64 `json:"margin"`
	IsOpen bool    `json:"is_open"`
}

func UpdateFutures(c *fiber.Ctx) error {
	tikerId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	marginBody := &MarginBody{}
	if err := c.BodyParser(marginBody); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	futures, err := repository.GetSingleFutures(tikerId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	futures.Margin += marginBody.Margin
	futures.IsOpen = marginBody.IsOpen
	if err := repository.UpdateFutures(futures); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(futures)
}

func GetFuturesSummary(c *fiber.Ctx) error {
	dateStart := c.Query("dateStart", "")
	dateEnd := c.Query("dateEnd", "")
	showOpen := c.Query("showOpen", "false")

	futures, err := repository.GetFutures(dateStart, dateEnd, showOpen)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	income := countFuturesSummary(futures)

	return c.Status(200).JSON(income)
}
