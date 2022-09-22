package main

import (
	"fmt"
	"net/http"

	"github.com/nicholasjackson/env"
)

var name = env.String("NAME",false,"Silas10","Name of the person to say Goodmorning")
func main(){
	env.Parse()

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Hello %s", *name)
	})

	http.ListenAndServe(":9090",nil)
}