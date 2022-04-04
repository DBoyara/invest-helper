package router

import (
	"github.com/DBoyara/invest-helper/pkg/models"
	"math"
)

const (
	SELL     string = "sell"
	BUY             = "buy"
	FixPRICE        = "fix_price"
	PERCENT         = "percent"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func countEquitySummary(logs []*models.TradingLog) *models.Summary {
	summary := &models.Summary{}

	for _, log := range logs {
		if log.Type == BUY {
			summary.Buy += log.Amount
		} else if log.Type == SELL {
			summary.Sell += log.Amount
		}
		summary.Commission += log.CommissionAmount
	}

	summary.Income = (summary.Sell - summary.Buy - summary.Commission) / summary.Buy * 100
	summary.Income = toFixed(summary.Income, 2)
	summary.Commission = toFixed(summary.Commission, 2)
	summary.Buy = toFixed(summary.Buy, 2)
	summary.Sell = toFixed(summary.Sell, 2)
	return summary
}

func countFuturesSummary(logs []*models.Futures) *models.Summary {
	summary := &models.Summary{}

	for _, log := range logs {
		summary.TurnoverMargin += log.Margin
		summary.TurnoverWP += log.Amount
		summary.Commission += log.Commission
	}

	summary.Income = (summary.TurnoverMargin - summary.Commission) / summary.TurnoverWP * 100
	summary.Income = toFixed(summary.Income, 2)
	summary.Commission = toFixed(summary.Commission, 2)
	return summary
}
