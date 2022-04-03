package module

import (
	"datadome/api_server"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpFields struct {
	Method     string `json:"method"`
	Host       string `json:"host"`
	RequestURI string `json:"requestURI"`
	Ip         string `json:"ip"`
	Referer    string `json:"referer"`
	UserAgent  string `json:"useragent"`
}

//datadome_local_parser_IP
//this function parses the IP of the requests
func ipParser(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

//datadome_local_parser
//this function creates the datastructure used to query the apiserver
func httpParser(r *http.Request) HttpFields {
	var httpReqFields HttpFields
	httpReqFields.Method = r.Method
	httpReqFields.Host = r.Host
	httpReqFields.RequestURI = r.RequestURI
	httpReqFields.Ip = ipParser(r)
	httpReqFields.Referer = r.Referer()
	httpReqFields.UserAgent = r.UserAgent()

	return httpReqFields
}

//AnalyzeHttpRequest:
//this function calls the parser to create the httpRequest data structure
//creates the jason format
//calls the api_server to analyze the request
func AnalyzeHttpRequest(r *http.Request) bool {
	fmt.Println("[MODULE] AnalyzeHttpRequest")

	var httpReqFields = httpParser(r)
	httpReqFieldsJson, err := json.Marshal(httpReqFields)
	if err != nil {
		fmt.Println(err)
	}

	return api_server.AnalyzeHttpRequest(string(httpReqFieldsJson))
}
