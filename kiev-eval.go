package main

import (
    "fmt"
    "os"
)

func main() {
    
    usage := "kiev-eval card1 card2 card3 card4 card5\n"
    args := os.Args

    if len(args) < 6 {
        fmt.Printf("%s\n", usage)
        os.Exit(1)
    }

    hand := NewHand([5]Card{
            NewCard(args[1]),
            NewCard(args[2]),
            NewCard(args[3]),
            NewCard(args[4]),
            NewCard(args[5]),
    })

    fmt.Printf("Hand %s value: %d\n", hand, hand.EvalFast())
}