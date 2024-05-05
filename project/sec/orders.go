package main

import (
	"encoding/csv"
	"fmt"
	"log"

	//"time"
	"encoding/json"
	//	"github.com/shopspring/decimal"
	//"golang.org/x/text/date"
	//	"google.golang.org/genproto/googleapis/type/datetime"

	//"strconv"
	//"datetime"
	"os"
	"sort"
)

type Order struct {
	Id        int32   `json:"Id"`
	Seccode   string  `json:"Seccode"`
	Myside    string  `json:"side"`
	Quantity  int32   `json:"Quantity"`
	Price     float64 `json:"Price"`
	Timestamp int64   `json:"Timestamp"`
  Custodian string `json:"Custodian"`
        HQty int32 `json:"HQty"`
        User string `json:"User"`
        TrdAcc string `json:"TrdAcc"`
        MsgType string `json:"MsgType"`
		TimeInForce string `json:"TimeInForce"`
		OrdType string `json:"OrdType"`

}
type Orderbook struct {
	//secid string
	seccode string
	Buys    []Order `json:"Buys"`
	Sells   []Order `json:"Sells"`
}

type rawOrders struct {
	Id         string `json:"Id"`
	Seccode    string `json:"Seccode"`
	Price      string `json:"Price"`
	Ord_qty    string `json:"Quantity"`
	Myside     string `json:"side"`
	Timestamp string `json:"Timestamp"`
}

type SecurityRecord struct {
	id      string
	Seccode string
}

func CreateOrdersList(data [][]string) []rawOrders {
	var OrdersList []rawOrders
	for i, line := range data {
		if i > 0 { // omit header line
			var rec rawOrders
			for j, field := range line {
				if j == 0 {
					rec.Id = field
				} else if j == 1 {
					rec.Seccode = field
				} else if j == 2 {
					rec.Price = field
				} else if j == 3 {
					rec.Ord_qty = field
				} else if j == 4 {
					rec.Myside = field
				} else if j == 5 {
					rec.Timestamp = field
				}
			}

			OrdersList = append(OrdersList, rec)
		}
	}
	return OrdersList
}

func CreateSecurityList(data [][]string) []SecurityRecord {
	var shoppingList []SecurityRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec SecurityRecord
			for j, field := range line {
				if j == 0 {
					rec.id = field
				} else if j == 1 {
					rec.Seccode = field
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getorderbook(ex []Orderbook, seccode string) int {
	var u int
	u = 0
	for i := 0; i < len(ex); i++ {
		var y Orderbook
		y = ex[i]

		if y.seccode == seccode {
			u = i
			return i
			//fmt.Println(y)

		}

		//break

	}
	return u
}

func AddBuyorder(ex []Orderbook, o Order, k int) []Orderbook {
	var iegx []Orderbook

	iegx = Egx
	var n int
	Egx[k].Buys = append(Egx[k].Buys, o)

	//if it is the only order then add at the end of the buy book
	n = len(ex[k].Buys)
	if n != 0 {
		sort.Slice(Egx[k].Buys, func(i, j int) bool {
			return Egx[k].Buys[i].Price > (Egx[k].Buys[j].Price)
		})

	}

	//  Printall(Egx.books[k].buys)
	return iegx
}

// Convert order to json from order struct
func (order *Order) toJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}

func Printall(d []Order) {
	//var s string
	for i := 0; i < len(d); i++ {
		fmt.Print(d[i].Price, " ")
		//fmt.Print(s," ")
	}
	fmt.Println()
}
