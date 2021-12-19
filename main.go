package main

func main() {
	cards := deck{"Ace of Diamonds", "King of Diamonds"}
	cards = append(cards, "Queen of Diamonds")

	cards.print()
}
