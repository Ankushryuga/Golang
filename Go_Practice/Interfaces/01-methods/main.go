package main

func main() {
	william := book{
		title: "Willam Sakeshpere",
		price: 100,
	}

	pokemonGO := game{
		title: "Pokemon GO",
		price: 200,
	}

	shadowFight := game{
		title: "Shadow Fight 5",
		price: 400,
	}

	// #3
	william.print()     // sends `william` value to `book.print`.
	pokemonGO.print()   // sends `pokemonGo` value to `game.print`.
	shadowFight.print() // sends `shadowFight` value to `game.print`.

	// #2
	// william.printBook()
	// pokemonGo.printGame()

	// #1
	// printBook(william)
	// printGame(pokemonGO)
}
