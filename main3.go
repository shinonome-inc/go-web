package main

import (
        "encoding/json"
        "fmt"
        "net/http"
)

type ErrorResponse struct {
        Status string   `json:"status:`
        Errors []string `json:"errors"`}
func jsonResponse(rw http.ResponseWriter, req *http.Request) {
        resp := ErrorResponse{"failure", []string{"user not found", "invalid password", "invalid username"}}
        response, _ := json.Marshal(resp)

        defer func() {
                rw.Header().Set("Content-Type", "application/json")
                fmt.Fprint(rw, string(response))
        }()
}
func main() {
        http.HandleFunc("/json", jsonResponse)
        http.ListenAndServe(":8080", nil)
}
