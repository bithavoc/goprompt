package main

import (
    "fmt"
    "github.com/bithavoc/goprompt"
    "os"
)

func main() {
    p := &goprompt.Prompt {
        Forms: []*goprompt.Form{
            {
                Title: "Enter your Bithavoc credentials to login",
                Fields: []*goprompt.Field {
                    {
                        Name: "email",
                        Title: "Email",
                        Instructions: "Please enter your Email",
                    },
                    {
                        Name: "password",
                        Title: "Password",
                        Instructions: "Please enter your Password",
                    },
                    {
                        Name: "remember",
                        Title: "Remember credentials?",
                        DefaultValue: "true",
                        Instructions: "Do you want to persist your credentials?",
                    },
                },
            },
        },
    }
    result := p.Process(os.Args[1:])
    fmt.Printf("%+v\n", result)
}
