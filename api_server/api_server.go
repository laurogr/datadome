package api_server

import (
	"fmt"
	"math/rand"
)

//AnalyzeHttpRequest: REMOTE FUNCTION CALL TO AWS
//this function simulates a call to the api server, granting or not acess to the cart endpoint based on the input parameters
func AnalyzeHttpRequest(httpRequestFields string) bool {
	fmt.Println("[API_SERVER] AnalyzeHttpRequest : Calls remote server : " + httpRequestFields)
	return (rand.Intn(100)%2 == 0)
}

//Example of code we could call, like a lambda in AWS (from amazon)

/*

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
        return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
        lambda.Start(HandleRequest)
}

*/
