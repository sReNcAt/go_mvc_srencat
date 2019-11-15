package main

import (
	"srencat/mvc_route"
	"net/http"
	"log"
	"fmt"
)

func main() {
	PORT := "8080"
	println("Start Server PORT '"+PORT+"' to Listening")
	start_server(PORT)
}

func start_server(PORT string){
	defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
	
	route := mvc_route.InitRouter()
	log.Fatal(http.ListenAndServe(":"+PORT, route))
}