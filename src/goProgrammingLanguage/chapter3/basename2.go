package main


import (
    "fmt"
    "strings"
)


func basename(s  string) string {
    lastIndex := strings.LastIndex(s, "/")
    s = s[lastIndex+1:]
    if qot := strings.LastIndex(s, "."); qot > 0{
        s = s[:qot]
    }
    return s
}


func main(){
    s := "a.b.c.go"
    s1 := "/a/b/c/d"
    s2 := "a/b/c.d.go"
    fmt.Println(basename(s))
    fmt.Println(basename(s1))
    fmt.Println(basename(s2))

}
