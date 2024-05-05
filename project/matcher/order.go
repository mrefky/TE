package main

import (
	"encoding/json"
	//"github.com/shopspring/decimal"
)

// Order Type
type Order struct {
	ID        int32 `json:"id"`
	Side      Side  `json:"side"`
	Quantity  int32  `json:"Quantity"`
	Price     float64  `json:"Price"`
	Timestamp int64 `json:"timestamp"`
	Seccode   string `json:"Seccode"`
	Custodian string `json:"Custodian"`
	HQty int32 `json:"HQty"`
	User string `json:"User"`
	TrdAcc string `json:"TrdAcc"`
	MsgType string `json:"MsgType"`
	TimeInForce string `json:"TimeInForce"`
	OrdType string `json:"OrdType"`

}

// Convert order to struct from json
func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

// Convert order to json from order struct
func (order *Order) toJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}
