package main

import (
    "fmt"
    "os"
    "os/user"
    "monkey/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hey %s! This is the Monkey programming language!", user.Username)
    fmt.Printf("Ready to read your commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
