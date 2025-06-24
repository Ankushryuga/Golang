package main
import "fmt"

type Books struct{
	title string
	author string
	bookdId int
}

func printBooks(book Books){
	fmt.Println("Book title is: ", book.title)
	fmt.Println("Book author is: ", book.author)
	fmt.Println("Book Id is: ", book.bookdId)
}


func main(){
	var book1 Books
	var book2 Books
	//first book infroamtions:
	book1.title="Jupiter Notebook"
	book1.author="Ankush"
	book1.bookdId=300304

	//second book informationsbook
	book2.title="Vscode"
	book2.author="microsoft"
	book2.bookdId=39399494

	printBooks(book1)
	fmt.Println("***********************##################")
	printBooks(book2)
}