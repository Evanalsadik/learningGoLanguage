package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter text: ")
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Println("You entered:", text)
        fmt.Println("Enter more text (or press Ctrl+D to exit):")
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

