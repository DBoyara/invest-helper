package models

import "time"

type TradingLog struct {
	Id               int       `json:"id,omitempty"`
	Datetime         time.Time `json:"datetime,omitempty"`
	Tiker            string    `json:"tiker"`
	Type             string    `json:"type"`
	IsOpen           bool      `json:"is_open,omitempty"`
	Price            float64   `json:"price"`
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
