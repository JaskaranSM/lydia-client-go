# Lydia API Wrapper [unofficial]

A Lydia Chatbot API wrapper written in Go.

# Installation

`
go get -u github.com/jaskaranSM/lydia-client-go
`

# Usage

lydia.NewClient Method takes an API_KEY and returns an instance of lydiaAI. this client have CreateSession and ThinkThought Methods which can be used to consume session creation and thought endpoints of the API.

# Example 

`
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
`

