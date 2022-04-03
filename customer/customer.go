package customer

import (
	"fmt"
	"net/http"
)

//customer_cart:
//this is the workflow for connections that are allowed
//this will call the backend to continue the workflow
func Cart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[CUSTOMER] ACESS TO CART GRANTED, CONTINUE WORKFLOW WITH ORIGINAL HTTP REQUEST")
	fmt.Println("[CUSTOMER] ACESS TO CART GRANTED, CONTINUE WORKFLOW WITH ORIGINAL HTTP REQUEST")
}
