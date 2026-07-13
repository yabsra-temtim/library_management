package services

import (
 	"fmt"
	"library_management/models"
	"sync"
)

// LibraryManager defines the interface for library operations
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
	GetAllBooks() []models.Book
	GetAllMembers() []models.Member
	AddMember(member models.Member)
	GetMember(memberID int) (*models.Member, error)
}

// Library implements LibraryManager interface
type Library struct {
	Books   map[int]models.Book   `json:"books"`
	Members map[int]models.Member `json:"members"`
	mu      sync.RWMutex          `json:"-"`
}

// NewLibrary creates a new Library instance
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	l.Books[book.ID] = book
	fmt.Printf("✅ Book '%s' added successfully!\n", book.Title)
}

// AddMember adds a new member to the library
func (l *Library) AddMember(member models.Member) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	l.Members[member.ID] = member
	fmt.Printf("✅ Member '%s' added successfully!\n", member.Name)
}

// GetMember retrieves a member by ID
func (l *Library) GetMember(memberID int) (*models.Member, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	member, exists := l.Members[memberID]
	if !exists {
		return nil, fmt.Errorf("member with ID %d not found", memberID)
	}
	return &member, nil
}

// RemoveBook removes a book from the library
func (l *Library) RemoveBook(bookID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	book, exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	
	if book.Status == "Borrowed" {
		return fmt.Errorf("cannot remove book '%s' as it is currently borrowed", book.Title)
	}
	
	delete(l.Books, bookID)
	fmt.Printf("✅ Book '%s' removed successfully!\n", book.Title)
	return nil
}

// BorrowBook allows a member to borrow a book
func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	// Check if book exists
	book, exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	
	// Check if book is available
	if !book.IsAvailable() {
		return fmt.Errorf("book '%s' is already borrowed", book.Title)
	}
	
	// Check if member exists
	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}
	
	// Check if member can borrow more books
	if !member.CanBorrow() {
		return fmt.Errorf("member '%s' has reached maximum borrowing limit (5 books)", member.Name)
	}
	
	// Borrow the book
	book.Borrow()
	l.Books[bookID] = book
	
	// Add to member's borrowed books
	member.AddBook(book)
	l.Members[memberID] = member
	
	fmt.Printf("✅ Book '%s' borrowed by '%s' successfully!\n", book.Title, member.Name)
	return nil
}

// ReturnBook allows a member to return a borrowed book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	// Check if book exists
	book, exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	
	// Check if book is borrowed
	if book.IsAvailable() {
		return fmt.Errorf("book '%s' is not borrowed", book.Title)
	}
	
	// Check if member exists
	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}
	
	// Check if member has borrowed this book
	if !member.HasBook(bookID) {
		return fmt.Errorf("member '%s' did not borrow '%s'", member.Name, book.Title)
	}
	
	// Return the book
	book.Return()
	l.Books[bookID] = book
	
	// Remove from member's borrowed books
	member.RemoveBook(bookID)
	l.Members[memberID] = member
	
	fmt.Printf("✅ Book '%s' returned by '%s' successfully!\n", book.Title, member.Name)
	return nil
}

// ListAvailableBooks returns all available books
func (l *Library) ListAvailableBooks() []models.Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	available := []models.Book{}
	for _, book := range l.Books {
		if book.IsAvailable() {
			available = append(available, book)
		}
	}
	return available
}

// ListBorrowedBooks returns all books borrowed by a member
func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	member, exists := l.Members[memberID]
	if !exists {
		return nil, fmt.Errorf("member with ID %d not found", memberID)
	}
	
	return member.BorrowedBooks, nil
}

// GetAllBooks returns all books in the library
func (l *Library) GetAllBooks() []models.Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	books := []models.Book{}
	for _, book := range l.Books {
		books = append(books, book)
	}
	return books
}

// GetAllMembers returns all members in the library
func (l *Library) GetAllMembers() []models.Member {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	members := []models.Member{}
	for _, member := range l.Members {
		members = append(members, member)
	}
	return members
}