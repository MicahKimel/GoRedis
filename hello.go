package main
import (
    "fmt"
    "context"
	"net/http"
	"io/ioutil"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		d, _ := ioutil.ReadAll(r.Body)
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		err := rdb.Set(ctx, "key", d, 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)
		fmt.Fprintf(rw, "Hello %s\n", d)

		val2, err := rdb.Get(ctx, "key2").Result()
		if err == redis.Nil {
			fmt.Println("key2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key2", val2)
    }})
	http.ListenAndServe(":9090",nil)
}

// import (
// _ "github.com/alexbrainman/odbc"
// "database/sql"
// "log"
// "fmt"
// )
 
// func main() {
// db, err := sql.Open("odbc",
// "DSN=CData DB2 Source")
// if err != nil {
// log.Fatal(err)
// }
 
// var (
// ordername string
// freight string
// )
 
// rows, err := db.Query("SELECT OrderName, Freight FROM Orders WHERE ShipCity = ?", "New York")
// if err != nil {
// log.Fatal(err)
// }
// defer rows.Close()
// for rows.Next() {
// err := rows.Scan(&ordername, &freight)
// if err != nil {
// log.Fatal(err)
// }
// fmt.Println(ordername, freight)
// }
// err = rows.Err()
// if err != nil {
// log.Fatal(err)
// }
 
// defer db.Close()
// }



