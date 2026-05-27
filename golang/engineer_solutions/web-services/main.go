package main

import(
	"net/http"
	"log"
)

func main(){
	/*mux := http.NewServeMux()
	mux.Handle("GET /users", chain(
		http.HandlerFunc(listUsers),
		Auth, Logger, Recovery,
	))*/

	router := &Router{}
	
	router.Use(Logger)
	router.Use(Recovery)
	router.Use(Auth)

	router.Handle("GET", "/users", listUsers)
	router.Handle("GET", "/users/:id", getUser)

	if err := http.ListenAndServe(":8080", router);  err != nil{
		log.Fatalf("error during server starts: %v", err)
	}
}