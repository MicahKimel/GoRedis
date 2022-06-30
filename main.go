package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/MicahKimel/GoRedis/handlers"
	myjwt "github.com/MicahKimel/GoRedis/jwt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	//env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handlers
	uh := handlers.NewUsers(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	getUnauth := sm.Methods(http.MethodGet, "OPTIONS").Subrouter()
	getUnauth.HandleFunc("/authenticate", uh.Authenticate)
	//getUnauth.Use(addCors)

	postUnauth := sm.Methods(http.MethodPost, "OPTIONS").Subrouter()
	postUnauth.HandleFunc("/createaccount", uh.AddUser)
	//postUnauth.Use(addCors)

	getR := sm.Methods(http.MethodGet, "OPTIONS").Subrouter()
	getR.HandleFunc("/id/{id}", uh.RedisTest)
	//getR.Use(addCors)
	getR.Use(authMiddleware)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getUnauth.Handle("/docs", sh)
	getUnauth.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := http.Server{
		Addr:         "localhost:9090",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Auth Called \n")
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		tmpjwt := myjwt.NewJwt(mux.Vars(r)["id"])
		claims, err := tmpjwt.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}
