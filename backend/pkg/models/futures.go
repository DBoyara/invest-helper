package models

import "time"

type Futures struct {
	Id                int64     `json:"id,omitempty"`
	Datetime          time.Time `json:"datetime,omitempty"`
	Tiker             string    `json:"tiker"`
	IsOpen            bool      `json:"is_open,omitempty"`
	WarrantyProvision float64   `json:"warranty_provision"`
	Count             int       `json:"count"`
	Amount            float64   `json:"amount,omitempty"`
	Margin            float64   `json:"margin,omitempty"`
	Commission        float64   `json:"commission"`
	CommissionAmount  float64   `json:"commission_amount,omitempty"`
	CommissionType    string    `json:"commission_type,omitempty"`
}
