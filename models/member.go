package models

import "fmt"

// Member represents a library member
type Member struct {
    ID            int    `json:"id"`
    Name          string `json:"name"`
    BorrowedBooks []Book `json:"borrowed_books"`
}

// NewMember creates a new member with empty borrowed books list
func NewMember(id int, name string) Member {
    return Member{
        ID:            id,
        Name:          name,
        BorrowedBooks: []Book{},
    }
}

// String returns a formatted string representation of the member
func (m Member) String() string {
    return fmt.Sprintf("ID: %d | Name: %s | Books Borrowed: %d",
        m.ID, m.Name, len(m.BorrowedBooks))
}

// CanBorrow checks if the member can borrow more books
func (m Member) CanBorrow() bool {
    return len(m.BorrowedBooks) < 5 // Max 5 books per member
}

// HasBook checks if the member has borrowed a specific book
func (m Member) HasBook(bookID int) bool {
    for _, book := range m.BorrowedBooks {
        if book.ID == bookID {
            return true
        }
    }
    return false
}

// AddBook adds a book to member's borrowed list
func (m *Member) AddBook(book Book) {
    m.BorrowedBooks = append(m.BorrowedBooks, book)
}

// RemoveBook removes a book from member's borrowed list
func (m *Member) RemoveBook(bookID int) error {
    for i, book := range m.BorrowedBooks {
        if book.ID == bookID {
            m.BorrowedBooks = append(m.BorrowedBooks[:i], m.BorrowedBooks[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("book with ID %d not found in member's borrowed list", bookID)
}