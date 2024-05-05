package main
import (
	"fmt"
//	"encoding/csv"
//"os"
//"log"
"strconv"
_"github.com/go-sql-driver/mysql"
"database/sql"
)
type Security struct{
	
	secid string
	seccode string	

}

type ShoppingRecord struct {
	ID      string
	Seccode string
}


func main(){


	var Sec [] Security
	//var EGX Egx= make("",[] OrderBook,0,300)
	records := readCsvFile("./Securities.csv")
	shoppingList := createShoppingList(records)
	var n, m string
	for i := 0; i < len(shoppingList); i++ {

		fmt.Sscan(shoppingList[i].ID, &n)
		fmt.Sscan(shoppingList[i].Seccode, &m)
		//fmt.Println(n, m)
	
		Sec1 := Security{secid: n, seccode: m}
		
		Sec = append(Sec, Sec1)
		//fmt.Println(Sec1)
		
	}

fmt.Println(Sec)

//we then wrtite this to secuirity table in the DB
db, err := sql.Open("mysql", "root@tcp(192.168.169.50:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var s string
	fmt.Println("Connecting to MySQL Success!")
	
	s="INSERT INTO Security (secid,seccode) values "
	for i:=0;i<len(Sec)-1;i++{
		
		//fmt.Println(Sec[i].seccode)
		 s=s+"("+Sec[i].secid+",\""+Sec[i].seccode+"\""+"),"
	}
	_,error:=db.Exec("Truncate table Security;")
	if error != nil {
		panic(err.Error())
	}
	//fmt.Println(result)

	s=s+"("+Sec[len(Sec)-1].secid+",\""+Sec[len(Sec)-1].seccode+"\""+")"
	
	insert, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}
	//defer insert.Close()
	
	

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Trauncated Security Table")
	fmt.Println("Added "+ strconv.Itoa(len(Sec))+" Securities to Security Table")
	}


