// Package Classification of API
//
// Documentation for API
// Schemes: http
// BasePath: /
// Version 1.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MicahKimel/GoRedis/data"
	"github.com/MicahKimel/GoRedis/jwt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) RedisTest(rw http.ResponseWriter, r *http.Request) {
	setupCORS(&rw, r)
	fmt.Print("REDIS TEST FUNC \n")
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
}

func (u *Users) AddUser(rw http.ResponseWriter, r *http.Request) {
	setupCORS(&rw, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	fmt.Print("CREATE USER CALLED\n")
	user := &data.User{}
	err := user.FromJSON(r.Body)
	if err != nil {
		fmt.Print("Unable to unmarshal json\n")
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	hsha256 := sha256.Sum256([]byte(string(user.Password)))

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

	mystring := string("call db.create_user('" + string(user.Username) + `', 
	'` + myhash + "', '" + string(user.Phone) + "', '" + string(user.Email) + "' )")

	fmt.Print(mystring)

	insert, err := db.Query(mystring)

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func (u *Users) Authenticate(rw http.ResponseWriter, r *http.Request) {
	setupCORS(&rw, r)
	fmt.Print("GET USER CALLED\n")
	user := r.URL.Query().Get("user")
	password := r.URL.Query().Get("password")
	fmt.Print("PASSWORD: " + password + "\n")

	hsha256 := sha256.Sum256([]byte(string(password)))

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

	mystring := "call db.get_user('" + string(user) + `', '` + string(myhash) + "' );"

	fmt.Print(mystring)

	var name string

	row := db.QueryRow(mystring)
	err = row.Scan(&name)
	// be careful deferring Queries if you are using transactions
	defer db.Close()

	if err != nil {
		fmt.Print("ERROR: Bad Username or Password\n")
		http.Error(rw, "ERROR: Bad Username or Password", http.StatusBadRequest)
	} else {
		fmt.Print(name)
		fmt.Print(err)
	}
	if name == user {
		myjwt := jwt.NewJwt(name)
		token, err := myjwt.GetToken(name)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			rw.Header().Set("Authorization", "Bearer "+token)
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Token: " + token))
		}
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Name and password do not match"))
		return
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
