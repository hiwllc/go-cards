package main

import "fmt"

func main() {
  cards := []string{"ace of diamonds", newCard()}
  cards = append(cards, "six of spades")

  for index, card := range cards {
    fmt.Println(index, card)
  }
}

func newCard() string {
  return "five of diamonds"
}
