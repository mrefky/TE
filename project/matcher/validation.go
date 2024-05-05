package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	//"strconv"
)

func validate(Ex []Egx){

fmt.Println(len(Ex))
var producer1 =createProducer()
var consumer1=createConsumer1()


go func(){

//fmt.Print(".")	
for msg := range consumer1.Messages() {
//	fmt.Println(msg.Value)	
//	raw :=String(msg[:].ToJSON()
if Valid(Ex,msg.Value){
producer1.Input() <- &sarama.ProducerMessage{
	Topic: "orders",
	Value: sarama.ByteEncoder(string(msg.Value)),
}
consumer1.MarkOffset(msg, "")
}
}

}()
}

func Valid(E1 []Egx,m []byte)bool{
return true	
	//fmt.Println(m)
	var order Order
	order.FromJSON(m)
	
if order.Side.String()=="buy" {
	return true
}
var x int =getorderbook(E1,order.Seccode)
if x ==-1{ 
fmt.Println("Unkown security")
	return false

}
if x !=-1{
var nn []Holding=E1[x].orderbook.Holdings
for i:=0;i<len(nn);i++{
	if nn[i].Quantity>=order.Quantity{
		return true
	}
}
fmt.Println("No Holdings; order rejected")
	return false
	
	}
	fmt.Println("No Holdings; order rejected")
return false
}
