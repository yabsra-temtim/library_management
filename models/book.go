package models

import "fmt"

// Book represents a book in the library
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Status string `json:"status"` // "Available" or "Borrowed"
}

// NewBook creates a new book with default status "Available"
func NewBook(id int, title, author string) Book {
    return Book{
        ID:     id,
        Title:  title,
        Author: author,
        Status: "Available",
    }
}

// String returns a formatted string representation of the book
func (b Book) String() string {
    return fmt.Sprintf("ID: %d | Title: %s | Author: %s | Status: %s",
        b.ID, b.Title, b.Author, b.Status)
}

// IsAvailable checks if the book is available
func (b Book) IsAvailable() bool {
    return b.Status == "Available"
}

// Borrow marks the book as borrowed
func (b *Book) Borrow() error {
    if b.Status == "Borrowed" {
        return fmt.Errorf("book '%s' is already borrowed", b.Title)
    }
    b.Status = "Borrowed"
    return nil
}

// Return marks the book as available
func (b *Book) Return() error {
    if b.Status == "Available" {
        return fmt.Errorf("book '%s' is already available", b.Title)
    }
    b.Status = "Available"
    return nil
}