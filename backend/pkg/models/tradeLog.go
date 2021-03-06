package models

import "time"

type TradingLog struct {
	Id               int       `json:"id,omitempty"`
	Datetime         time.Time `json:"datetime,omitempty"`
	Tiker            string    `json:"tiker"`
	TikerType        string    `json:"tiker_type"`
	Type             string    `json:"type"`
	IsOpen           bool      `json:"is_open,omitempty"`
	Price            float64   `json:"price"`
	Currency         string    `json:"currency"`
	Count            int       `json:"count"`
	Lot              int       `json:"lot"`
	Amount           float64   `json:"amount,omitempty"`
	Commission       float64   `json:"commission"`
	CommissionAmount float64   `json:"commission_amount"`
	CommissionType   string    `json:"commission_type,omitempty"`
}

type Commission struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type Summary struct {
	Buy            float64 `json:"buy,omitempty"`
	Sell           float64 `json:"sell,omitempty"`
	TurnoverMargin float64 `json:"turnover_margin,omitempty"`
	TurnoverWP     float64 `json:"turnover_wp,omitempty"`
	Commission     float64 `json:"commission"`
	Income         float64 `json:"income"`
}
