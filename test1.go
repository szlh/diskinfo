package main
import "fmt"

func main() {
    var m map[string]string
    m["name"] = "zzy"
}

func main() {
    var m map[string]string
    if m == nil {
        fmt.Println("this map is a nil map")
    }
}

import (
    "errors"
    "fmt"
)

func main() {
    i := 2
    if i > 1 {
        i, err := doDivision(i, 2)
        if err != nil {
            panic(err)
        }
        fmt.Println(i)
    }
    fmt.Println(i)
}

func doDivision(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("input is invalid")
    }
    return x / y, nil
}
type person struct {
    name   string
    age    byte
    isDead bool
}

func main() {
    p1 := person{name: "zzy", age: 100}
    p2 := person{name: "dj", age: 99}
    p3 := person{name: "px", age: 20}
    people := []person{p1, p2, p3}
    whoIsDead(people)
    for _, p := range people {
        if p.isDead {
            fmt.Println("who is dead?", p.name)
        }
    }
}

func whoIsDead(people []person) {
    for _, p := range people {
        if p.age < 50 {
            p.isDead = true
        }
    }
}
