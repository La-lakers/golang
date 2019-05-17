package main


import (
    "fmt"
)

// a.c -> a; /a/b/c.go -> c ; a.b.c -> a.b
func  basename( str string ) string {
    //  去掉最后一个/
    for i := len(str) -1; i >= 0; i--{
        if str[i] == '/'{
	    str = str[i+1:]
	    break
	}
    }
    // 去掉最后一个.
    for i := len(str) -1; i >= 0; i--{
        if str[i] == '.'{
	   str = str[:i]
	   break
	}
    }
    return str

}


func main(){
    s := "a.b.c"
    fmt.Println(basename(s))
    s1 := "/a/b/c.go"
    fmt.Println(basename(s1))

}
