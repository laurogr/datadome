package main

import (
	"datadome/customer"
	"datadome/module"
	"fmt"
	"log"
	"net/http"
)

//cart
//this function simulates the customer endpoint
//it calls first the local datadome module to analyze the http headers
//if it is not a bot we continue the workflow
func cart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[MAIN] /cart")

	var isBot = module.AnalyzeHttpRequest(r)
	//I am only considering here 2 states : Is/Isnt bot.
	//however, we could other states such as captcha. (for simplicity, I am just exploring those cases)
	if isBot {
		http.Error(w, "ERROR : FORBIDDEN ACCESS", http.StatusForbidden)
	} else {
		customer.Cart(w, r)
	}
}

//main
//loop dealing with the calls to the /cart endpoint
func main() {
	http.HandleFunc("/cart", cart)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
