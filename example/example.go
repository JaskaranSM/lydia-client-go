package main

import (
	"fmt"
	"github.com/jaskaranSM/lydia-client-go/lydia"
)

unc main(){
	var API_KEY = "<your_api_key>"
	var input string
    client := lydia.NewClient(API_KEY)
    res,err := client.CreateSession()
    if err != nil {
    	fmt.Println(err)
    }
	fmt.Scanln(&input)
	out,err := client.ThinkThought(res.Payload.SessionId,input)
    if err != nil {
    	fmt.Println(err)
    }
    fmt.Println("Lydia: "+out.Payload.Output)
}