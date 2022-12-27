package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Book struct {
	id     string
	title  string
	author string
	price  float32
}

type Books struct {
	items []Book
}

//add book
func (b *Books) Add(book Book) {
	b.items = append(b.items, book)
}


func main() {
	myBookStore := Books{}
	newBook := Book{}
	newBook.id = uuid.New().String()
	newBook.title = "Harry Potter"
	newBook.author = "J K Rowling"
	newBook.price = 9.80

	myBookStore.Add(newBook)

	newBook.id = uuid.New().String()
	newBook.title = "Harry Potter 2"
	newBook.author = "J K Rowling 435"
	newBook.price = 345

	myBookStore.Add(newBook)

	fmt.Println(myBookStore)
}
