package main

import "fmt"

func main(){
    fmt.Println("hello world")
    //define variables
    v := 10
    d := 15.0
    s := "this is a dog"
    s_s := `this is a dog`
    c := 't'
    b := true
    fmt.Printf("%v \n", v)
    fmt.Printf("%v \n", d)
    fmt.Printf("%v \n", s)
    fmt.Printf("%v \n", s_s)
    fmt.Printf("%v \n", c)
    fmt.Printf("%v \n", b)
    //variable with zero value

    var aa int
    var bb float32
    var cc string
    var dd bool
    ff, f2, f3 := 1, 2, 3
    ss_s := "hello, how are you"
    fmt.Printf("%v \n", aa)
    fmt.Printf("%v \n", bb)
    fmt.Printf("%v \n", cc)
    fmt.Printf("%v \n", dd)
    fmt.Printf("%v \n", ss_s)
    fmt.Printf("%v %v %v\n", ff, f2, f3)

    for i := 0; i < 100; i++ {
        if i%3 == 0 {
            fmt.Print(" ", i)
        }
    }
    var i = 100
    for i < 500 {
        fmt.Print("loop\t")
        i += 100
    }
    for j := 90; j < 100 ; j++{
        fmt.Print(j, "--",string(i), "<><> ", []byte(string(i)) , "\n")
    }
    foo := "a"
    fmt.Printf("  %T \n", foo)
    for j := 10; j < 90 ; j++ {
        fmt.Printf("%v %v %v\n", j, string(j), []byte(string(j)))
    }
    
    boyB := true
    if food := "Burger"; boyB {
        fmt.Println(food)
    }
}