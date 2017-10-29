package main

func main() {
	// cards := newDeckFromFile("teste.txt")
	// cards.print()
	cards := newDeck()
	cards.saveToFile("teste.txt")

}
