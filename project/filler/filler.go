package main

import (
	"flag"
	"fmt"
	"path"
	"syscall"
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
	"github.com/shopspring/decimal"

	//"reflect"

	fix44er "github.com/quickfixgo/fix44/executionreport"
	fix44nos "github.com/quickfixgo/fix44/newordersingle"

	///"github.com/quickfixgo/fix44/ordercancelreplacerequest"
	OCRR "github.com/quickfixgo/fix44/ordercancelreplacerequest"
	OCR "github.com/quickfixgo/fix44/ordercancelrequest"

	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/Shopify/sarama"
	//cluster "github.com/bsm/sarama-cluster"
	//"timestamppb"
	// "time"
)

type Order struct {
	ID          int32   `json:"id"`
	Side        Side    `json:"side"`
	Quantity    int32   `json:"Quantity"`
	Price       float64 `json:"Price"`
	Timestamp   int64   `json:"timestamp"`
	Seccode     string  `json:"Seccode"`
	Custodian   string  `json:"Custodian"`
	HQty        int32   `json:"HQty"`
	User        string  `json:"User"`
	TrdAcc      string  `json:"TrdAcc"`
	MsgType     string  `json:"MsgType"`
	TimeInForce string  `json:"TimeInForce"`
	OrdType     string  `json:"OrdType"`
}

type executor struct {
	orderID int
	execID  int
	*quickfix.MessageRouter
}

func (order *Order) toJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}

func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

func newExecutor() *executor {
	e := &executor{MessageRouter: quickfix.NewMessageRouter()}

	e.AddRoute(fix44nos.Route(e.OnFIX44NewOrderSingle))
	e.AddRoute(OCR.Route(e.OnOrderCancelRequest))
	e.AddRoute(OCRR.Route(e.OnOrderCancelReplaceRequest))

	return e
}

func (e *executor) genOrderID() field.OrderIDField {
	e.orderID++
	return field.NewOrderID(strconv.Itoa(e.orderID))
}

func (e *executor) genExecID() field.ExecIDField {
	e.execID++
	return field.NewExecID(strconv.Itoa(e.execID))
}

// quickfix.Application interface
func (e executor) OnCreate(sessionID quickfix.SessionID)                           {}
func (e executor) OnLogon(sessionID quickfix.SessionID)                            {}
func (e executor) OnLogout(sessionID quickfix.SessionID)                           {}
func (e executor) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID)     {}
func (e executor) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error { return nil }
func (e executor) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

// Use Message Cracker on Incoming Application Messages
func (e *executor) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	return e.Route(msg, sessionID)
}

func (e *executor) OnOrderCancelReplaceRequest(msg OCRR.OrderCancelReplaceRequest, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	if err != nil {
		return err
	}
	fmt.Println("Request to Amend order recived")
	fmt.Println(msg.ToMessage().String())
	return
	//e.Route(msg.Message, sessionID)
}

func (e *executor) OnOrderCancelRequest(msg OCR.OrderCancelRequest, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	if err != nil {
		return err
	}
	fmt.Println("Request to cancel order recived")
	fmt.Println(msg.ToMessage().String())
	side, err := msg.GetSide()
	if err != nil {
		return
	}
	orderQty := decimal.New(0,10)

	if err != nil {
		return
	}
	var price decimal.Decimal = decimal.New(0, 10)

	execReport := fix44er.New(
		e.genOrderID(),
		e.genExecID(),
		field.NewExecType(enum.ExecType_CANCELED),
		field.NewOrdStatus(enum.OrdStatus_CANCELED),
		field.NewSide(side),
		field.NewLeavesQty(decimal.Zero, 2),
		field.NewCumQty(orderQty, 2),
		field.NewAvgPx(price, 2),
	)
	clOrdID, err := msg.GetClOrdID()
	if err != nil {
		return
	}
	symbol, err := msg.GetSymbol()
	if err != nil {
		return
	}

	execReport.SetClOrdID(clOrdID)
	execReport.SetSymbol(symbol)
	execReport.SetOrderQty(orderQty, 2)
	execReport.SetLastQty(orderQty, 2)
	execReport.SetLastPx(price, 2)


	sendErr := quickfix.SendToTarget(execReport, sessionID)
	if sendErr != nil {
		fmt.Println(sendErr)
	}
	fmt.Println(msg.ToMessage().String())

	return
	//e.Route(msg.Message, sessionID)
}

func (e *executor) OnFIX44NewOrderSingle(msg fix44nos.NewOrderSingle, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	ordType, err := msg.GetOrdType()
	if err != nil {
		return err
	}

	if ordType != enum.OrdType_LIMIT {
		return quickfix.ValueIsIncorrect(tag.OrdType)
	}

	symbol, err := msg.GetSymbol()
	if err != nil {
		return
	}

	side, err := msg.GetSide()
	if err != nil {
		return
	}

	orderQty, err := msg.GetOrderQty()
	if err != nil {
		return
	}

	price, err := msg.GetPrice()
	if err != nil {
		return
	}

	clOrdID, err := msg.GetClOrdID()
	if err != nil {
		return
	}

	firmid, err := msg.GetSenderCompID()
	if err != nil {
		return err
	}

	var sside Side
	if side == "2" {
		sside = 2
	}
	if side == "1" {
		sside = 1
	}
	//	t := time.Date(2020, 11, 14, 11, 30, 32, 0, time.UTC)

	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	// Timestamp := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix()
	t := (date.Unix())
	if err != nil {
		panic(err)
	}
	// t := timestamppb.Now() //(time.Unix(timestamp, 0)).String()
	fmt.Println(t)
	var ord Order
	// var ordID1 int64
	// ordID1,err=strconv(clOrdID, 10, 32)
	//
	//	if err != nil {
	//	  panic(err)/
	//	}
	var err1 error
	var x int64
	x, err1 = strconv.ParseInt(clOrdID, 10, 64)
	if err1 != nil {
		panic(err)
	}
	ord.ID = int32(x)
	//(clOrdID,

	ord.Side = sside

	x, err1 = strconv.ParseInt(orderQty.String(), 10, 64)
	if err1 != nil {
		panic(err)
	}
	ord.Quantity = int32(x)
	var tx float64
	tx, err1 = strconv.ParseFloat(price.String(), 64)
	if err1 != nil {
		panic(err)
	}
	ord.Price = tx

	// price.String()
	ord.User = firmid
	ord.Seccode = symbol
	//f:=ord.toJSON()

	var custodian string

	custodian, err = msg.GetString(tag.PartyID)
	if err != nil {
		return
	}
	ord.Custodian = custodian

	ord.TimeInForce, err = msg.GetString(tag.TimeInForce)
	if err != nil {
		return
	}

	ord.OrdType, err = msg.GetString(tag.OrdType)
	if err != nil {
		return
	}

	ord.TrdAcc, err = msg.GetAccount()
	if err != nil {
		return
	}

	// f:="{\"id\":" + clOrdID + ",\"side\":\"" + sside + "\",\"price\":" + price.String() + ",\"quantity\":" + orderQty.String() + ",\"firmid\":\"" + firmid + "\",\"timestamp\":" + strconv.FormatInt(t,10) + ",\"Seccode\":\"" + symbol + "\"}"
	f := ord.toJSON()
	fmt.Println(string(f[:]))

	producer := createProducer()
	producer.Input() <- &sarama.ProducerMessage{
		Topic: "orders",
		Value: sarama.ByteEncoder(f),
		Key:   sarama.ByteEncoder(symbol),
	}

	execReport := fix44er.New(
		e.genOrderID(),
		e.genExecID(),
		field.NewExecType(enum.ExecType_PENDING_NEW),
		field.NewOrdStatus(enum.OrdStatus_PENDING_NEW),
		field.NewSide(side),
		field.NewLeavesQty(decimal.Zero, 2),
		field.NewCumQty(orderQty, 2),
		field.NewAvgPx(price, 2),
	)

	execReport.SetClOrdID(clOrdID)
	execReport.SetSymbol(symbol)
	execReport.SetOrderQty(orderQty, 2)
	execReport.SetLastQty(orderQty, 2)
	execReport.SetLastPx(price, 2)

	if msg.HasAccount() {
		acct, err := msg.GetAccount()
		if err != nil {
			return err
		}
		execReport.SetAccount(acct)
	}

	sendErr := quickfix.SendToTarget(execReport, sessionID)
	if sendErr != nil {
		fmt.Println(sendErr)
	}

	return
}

func main() {

	flag.Parse()

	cfgFileName := path.Join("config", "executor.cfg")
	if flag.NArg() > 0 {
		cfgFileName = flag.Arg(0)
	}

	cfg, err := os.Open(cfgFileName)
	if err != nil {
		fmt.Printf("Error opening %v, %v\n", cfgFileName, err)
		return
	}

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Error reading cfg,", err)
		return
	}

	logFactory, _ := quickfix.NewFileLogFactory(appSettings)
	app := newExecutor()

	acceptor, err := quickfix.NewAcceptor(app, quickfix.NewMemoryStoreFactory(), appSettings, logFactory)
	if err != nil {
		fmt.Printf("Unable to create Acceptor: %s\n", err)
		return
	}

	err = acceptor.Start()
	if err != nil {
		fmt.Printf("Unable to start Acceptor: %s\n", err)
		return
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	acceptor.Stop()
}
func createProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false         // fire and forget
	config.Producer.Return.Errors = true             // notify on failed
	config.Producer.RequiredAcks = sarama.WaitForAll // waits for all insync replicas to commit

	producer, err := sarama.NewAsyncProducer([]string{"kafka.default.svc.cluster.local:9092"}, config)
	//producer, err := sarama.NewAsyncProducer([]string{"192.168.1.71:9094"}, config)
	if err != nil {
		log.Fatal("Unable to connect producer to kafka server")
	}

	return producer

}
