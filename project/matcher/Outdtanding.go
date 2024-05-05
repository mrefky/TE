package main
import( 
"fmt"
"log"
"database/sql"
"strconv"
)
func Outstanding( gx []Egx)  []Egx{

	fmt.Println("Synchronizing .............................. ")


var Dbcon string =MYSQL
	//Dbcon="root@tcp(192.168.1.50:3306)/test"
	db, err := sql.Open("mysql", Dbcon)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select orderid,Quantity,Remaining,ordtype,secid,timestamp,price from test.traded where (Remaining !=0)or (Remaining is null) ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var sid string
	var mtype string
	var n,m,timest int
	var p float64
	var c sql.NullString
	for rows.Next() {
	err := rows.Scan(&n,&m,&c,&mtype,&sid,&timest,&p)
		if err != nil {
			log.Fatal(err)
		}
		var Remain int
		if c.Valid {
			xsv,_:= strconv.Atoi(c.String)
			Remain=xsv
		}else {
			Remain=m
		}
		
		var g1 Order
		g1.ID=int32(n)
		g1.Quantity=int32(Remain)
		g1.Price=p 
		g1.Seccode=sid
		if mtype=="sell"{
			g1.Side=Sell
			//fmt.Println(g1.Side)
		}else {g1.Side=Buy
			//fmt.Println(g1.Side)
		}
		
		g1.Timestamp=int64(timest)
		//fmt.Println(g1)
		var xyz int = getorderbook(gx, g1.Seccode)
		if g1.Side==Buy{
		gx[xyz].orderbook.addBuyOrder(g1)

		}else{gx[xyz].orderbook.addSellOrder(g1)}

	//	fmt.Println(gx[xyz])
		
	}

	
		//fmt.Println(n,":",m,":",Remain," ",mtype," ",timest," ",sid)
	
	


	
fmt.Println("Ready ..........................")
return gx

}
