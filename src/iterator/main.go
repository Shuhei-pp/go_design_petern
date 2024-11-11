package iterator

import "fmt"

func Main() {
	bookShelf := &BookShelf{
		Books: []Book{
			{Name: "Book1"},
			{Name: "Book2"},
			{Name: "Book3"},
		},
	}

	for bookShelf.Iterator().HasNext() {
		bookShelfIterator := bookShelf.Iterator()
		book := bookShelfIterator.Next()
		fmt.Println(book.Name)
	}
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Book struct {
	Name string
}

type BookShelf struct {
	Books []Book
}

func (b *BookShelf) Iterator() BookShelfIterator {
	return BookShelfIterator{b}
}

type BookShelfIterator struct {
	BookShelf *BookShelf
}

func (bs BookShelfIterator) HasNext() bool {
	return len(bs.BookShelf.Books) > 0
}

func (bs BookShelfIterator) Next() Book {
	book := bs.BookShelf.Books[0]
	bs.BookShelf.Books = bs.BookShelf.Books[1:]
	return book
}
