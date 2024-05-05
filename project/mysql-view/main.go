package main

import (
    "fmt"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "os"
    "html/template"
)

type Order struct {
    Orderid string ;
    Price string;
    Remaining string;
    Quantity   string;
    Secid   string;
    Ordtype string
}

func helloWorld(w http.ResponseWriter, r *http.Request){
    name, err := os.Hostname()
    checkErr(err)
    fmt.Fprintf(w, "HOSTNAME : %s\n", name)
}

func dbConnect() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    //dbPass := ""
    dbHost := "192.168.169.50"
    dbPort := "3306"
    dbName := "test"
    db, err := sql.Open(dbDriver, dbUser +"@tcp("+ dbHost +":"+ dbPort +")/"+ dbName +"?charset=utf8")
    checkErr(err)
    return db
}

func dbSelect() []Order{
    db := dbConnect()
    rows, err := db.Query("select Orderid,Price,CONVERT(COALESCE(Remaining,Quantity),SIGNED INTEGER) as Remaining ,Quantity,Secid,Ordtype from traded where Remaining>0 or remaining is null  order by secid,`Remaining` desc,Orderid ")

    checkErr(err)

    order := Order{}
    orders := []Order{}

    for rows.Next() {
        var Orderid,Price,Remaining,Quantity,Secid,Ordtype string
        err = rows.Scan(&Orderid, &Price,&Remaining,&Quantity,&Secid,&Ordtype)
        checkErr(err)
        order.Orderid = Orderid
        order.Secid=Secid
        order.Price = Price
        order.Remaining=Remaining
        order.Quantity=Quantity
        order.Ordtype=Ordtype
        orders= append(orders, order)

    }
    defer db.Close()
    //fmt.Println(employees)
    return orders
}

var tmpl = template.Must(template.ParseFiles("layout.html"))
//var tmpl = template.Must(template.ParseGlob("layout.html"))
func dbTableHtml(w http.ResponseWriter, r *http.Request){
    table := dbSelect()
    err := tmpl.ExecuteTemplate(w, "Index", table)
    if err != nil {
                // Do something with the error
                fmt.Println(err)

        }
    //tmpl.ExecuteTemplate(w, "Index", table)
}

func dbTable(w http.ResponseWriter, r *http.Request){
    table := dbSelect()
    for i := range(table) {
        emp := table[i]
        fmt.Fprintf(w,"YESS|%12s|%12s|%12s|%20s|\n" ,emp.Orderid ,emp.Price,emp.Quantity,emp.Secid,emp.Ordtype)
    }
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/view", dbTableHtml)
    http.HandleFunc("/raw", dbTable)
    http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
