package  main

import "encoding/json"

type Trade struct {
	TakerOrderID int32 `json:"taker_order_id"`
	MakerOrderID int32 `json:"maker_order_id"`
	Quantity     int32
	Price        float64 `json:"price"`
	Timestamp    int64   `json:"timestamp"`
}

// struct to json
func (trade *Trade) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, trade)
}

func (trade *Trade) ToJSON() []byte {
	str, _ := json.Marshal(trade)
	return str
}
