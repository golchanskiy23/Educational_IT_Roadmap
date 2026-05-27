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

	router := &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
	router.Handle("GET", "/users", http.HandlerFunc(listUsers))

	if err := http.ListenAndServe(":8080", router);  err != nil{
		log.Fatalf("error during server starts: %v", err)
	}
}