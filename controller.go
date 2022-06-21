package main
import (
	"crypto/sha256"
	"fmt"
    "database/sql"
	"context"
    _ "github.com/go-sql-driver/mysql"
	"encoding/base64"
	 "net/http"
	"io/ioutil"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// struct user() {
// 	username string
// 	password string
// 	phone string
// 	email string
// }

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
    	}
	})
	http.HandleFunc("/createuser", func(rw http.ResponseWriter, r*http.Request){
		//next cast input body to type user and post
		//d, _ := ioutil.ReadAll(r.Body)
		s := "Foo"
		hsha256 := sha256.Sum256([]byte(s))

		fmt.Printf("SHA256: %x\n", hsha256)

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db")

		// if there is an error opening the connection, handle it
		if err != nil {
			panic(err.Error())
		}

		// defer the close till after the main function has finished
		// executing
		defer db.Close()

		myhash := base64.StdEncoding.EncodeToString(hsha256[:])

		mystring:= string("call db.create_user('Test1', '" + myhash + "', 'TEST', 'TETS' )")

		insert, err := db.Query(mystring)

		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
	})
	http.ListenAndServe(":9090",nil)
}


