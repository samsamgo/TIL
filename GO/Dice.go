package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // 주사위를 굴립니다.
    diceRoll := rand.Intn(6) + 1

    // 주사위 눈을 출력합니다.
    fmt.Println("주사위 눈:", diceRoll)

    // 주사위 눈이 3이면 "박수!"를 출력합니다.
    if diceRoll == 3 {
        fmt.Println("박수!")
    }
}
