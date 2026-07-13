package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/services"
	"library_management/models"
)

func main() {
	// Initialize library service
	library := services.NewLibrary()
	
	// Seed some initial data
	seedData(library)
	
	// Initialize controller
	controller := controllers.NewLibraryController(library)
	
	// Run the application
	controller.Run()
}

func seedData(library *services.Library) {
	// Add some sample books
	library.AddBook(models.NewBook(1, "The Great Gatsby", "F. Scott Fitzgerald"))
	library.AddBook(models.NewBook(2, "To Kill a Mockingbird", "Harper Lee"))
	library.AddBook(models.NewBook(3, "1984", "George Orwell"))
	library.AddBook(models.NewBook(4, "Pride and Prejudice", "Jane Austen"))
	library.AddBook(models.NewBook(5, "The Catcher in the Rye", "J.D. Salinger"))
	
	// Add some sample members
	library.AddMember(models.NewMember(1, "Alice Johnson"))
	library.AddMember(models.NewMember(2, "Bob Smith"))
	library.AddMember(models.NewMember(3, "Charlie Brown"))
	
	fmt.Println("📚 Sample data loaded successfully!")
}