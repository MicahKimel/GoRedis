package main

import (
	"net/http"
	"github.com/MicahKimel/GoRedis/handlers"
)


func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		handlers.redisTest(rw, r)
	})
	http.HandleFunc("/createuser", func(rw http.ResponseWriter, r*http.Request){
		handlers.addUser(rw, r)
	})
	http.ListenAndServe(":9090",nil)
}



