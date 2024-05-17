package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func evenNames(slice []string) []string {
    //Tulis kode disini
    var evenName []string
    for _, name := range slice {
        if len(name) % 2 == 0 {
            evenName = append(evenName, name)
        }
    }
    return evenName
}
 

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    x := scanner.Text()
    slice := strings.Split(x, " ")
    names := evenNames(slice)
    for _, pname := range names {
        fmt.Printf("%s ", pname)
    }
    fmt.Println()
    fmt.Println(names)
}

