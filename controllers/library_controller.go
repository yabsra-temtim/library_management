package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

// LibraryController handles console interaction
type LibraryController struct {
	service services.LibraryManager
	reader  *bufio.Reader
}

// NewLibraryController creates a new controller
func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

// Run starts the console interface
func (c *LibraryController) Run() {
	for {
		c.showMainMenu()
		choice := c.getInput("Enter your choice: ")
		
		switch choice {
		case "1":
			c.addBook()
		case "2":
			c.removeBook()
		case "3":
			c.addMember()
		case "4":
			c.borrowBook()
		case "5":
			c.returnBook()
		case "6":
			c.listAvailableBooks()
		case "7":
			c.listBorrowedBooks()
		case "8":
			c.listAllBooks()
		case "9":
			c.listAllMembers()
		case "0", "exit", "quit":
			fmt.Println("\n👋 Goodbye! Thanks for using Library Management System!")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
	}
}

func (c *LibraryController) showMainMenu() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📚 LIBRARY MANAGEMENT SYSTEM")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("\n1. 📖 Add a Book")
	fmt.Println("2. 🗑️  Remove a Book")
	fmt.Println("3. 👤 Add a Member")
	fmt.Println("4. 📤 Borrow a Book")
	fmt.Println("5. 📥 Return a Book")
	fmt.Println("6. 📋 List Available Books")
	fmt.Println("7. 📋 List Borrowed Books by Member")
	fmt.Println("8. 📚 List All Books")
	fmt.Println("9. 👥 List All Members")
	fmt.Println("0. 🚪 Exit")
	fmt.Println(strings.Repeat("=", 60))
}

func (c *LibraryController) getInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (c *LibraryController) getIntInput(prompt string) (int, error) {
	input := c.getInput(prompt)
	return strconv.Atoi(input)
}

func (c *LibraryController) addBook() {
	fmt.Println("\n📖 ADD A NEW BOOK")
	fmt.Println(strings.Repeat("-", 30))
	
	id, err := c.getIntInput("Enter Book ID: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	title := c.getInput("Enter Book Title: ")
	if title == "" {
		fmt.Println("❌ Title cannot be empty.")
		return
	}
	
	author := c.getInput("Enter Book Author: ")
	if author == "" {
		fmt.Println("❌ Author cannot be empty.")
		return
	}
	
	book := models.NewBook(id, title, author)
	c.service.AddBook(book)
}

func (c *LibraryController) removeBook() {
	fmt.Println("\n🗑️  REMOVE A BOOK")
	fmt.Println(strings.Repeat("-", 30))
	
	id, err := c.getIntInput("Enter Book ID to remove: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	if err := c.service.RemoveBook(id); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}
}

func (c *LibraryController) addMember() {
	fmt.Println("\n👤 ADD A NEW MEMBER")
	fmt.Println(strings.Repeat("-", 30))
	
	id, err := c.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	name := c.getInput("Enter Member Name: ")
	if name == "" {
		fmt.Println("❌ Name cannot be empty.")
		return
	}
	
	member := models.NewMember(id, name)
	c.service.AddMember(member)
}

func (c *LibraryController) borrowBook() {
	fmt.Println("\n📤 BORROW A BOOK")
	fmt.Println(strings.Repeat("-", 30))
	
	bookID, err := c.getIntInput("Enter Book ID to borrow: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	memberID, err := c.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	if err := c.service.BorrowBook(bookID, memberID); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}
}

func (c *LibraryController) returnBook() {
	fmt.Println("\n📥 RETURN A BOOK")
	fmt.Println(strings.Repeat("-", 30))
	
	bookID, err := c.getIntInput("Enter Book ID to return: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	memberID, err := c.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	if err := c.service.ReturnBook(bookID, memberID); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}
}

func (c *LibraryController) listAvailableBooks() {
	fmt.Println("\n📋 AVAILABLE BOOKS")
	fmt.Println(strings.Repeat("-", 30))
	
	books := c.service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}
	
	for _, book := range books {
		fmt.Println(book.String())
	}
	fmt.Printf("\nTotal: %d books available\n", len(books))
}

func (c *LibraryController) listBorrowedBooks() {
	fmt.Println("\n📋 BORROWED BOOKS BY MEMBER")
	fmt.Println(strings.Repeat("-", 30))
	
	memberID, err := c.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return
	}
	
	books, err := c.service.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	
	if len(books) == 0 {
		fmt.Println("This member has no borrowed books.")
		return
	}
	
	for _, book := range books {
		fmt.Println(book.String())
	}
	fmt.Printf("\nTotal: %d books borrowed\n", len(books))
}

func (c *LibraryController) listAllBooks() {
	fmt.Println("\n📚 ALL BOOKS")
	fmt.Println(strings.Repeat("-", 30))
	
	// Type assertion to access GetAllBooks method
	if lib, ok := c.service.(*services.Library); ok {
		books := lib.GetAllBooks()
		if len(books) == 0 {
			fmt.Println("No books in the library.")
			return
		}
		
		for _, book := range books {
			fmt.Println(book.String())
		}
		fmt.Printf("\nTotal: %d books in library\n", len(books))
	} else {
		fmt.Println("❌ Error accessing library data.")
	}
}

func (c *LibraryController) listAllMembers() {
	fmt.Println("\n👥 ALL MEMBERS")
	fmt.Println(strings.Repeat("-", 30))
	
	if lib, ok := c.service.(*services.Library); ok {
		members := lib.GetAllMembers()
		if len(members) == 0 {
			fmt.Println("No members registered.")
			return
		}
		
		for _, member := range members {
			fmt.Println(member.String())
		}
		fmt.Printf("\nTotal: %d members\n", len(members))
	} else {
		fmt.Println("❌ Error accessing member data.")
	}
}