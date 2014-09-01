package main

import "github.com/bithavoc/goprompt"

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
    p.Process()
}
