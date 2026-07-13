# Library Management System Documentation

## Overview
A console-based Library Management System built in Go that demonstrates the use of structs, interfaces, methods, slices, and maps.

## Features
- Add and remove books
- Add and manage members
- Borrow and return books
- List available books
- List borrowed books by member
- View all books and members

## Architecture

### Models
- **Book**: Represents a book with ID, Title, Author, and Status
- **Member**: Represents a member with ID, Name, and BorrowedBooks

### Services
- **LibraryManager**: Interface defining all library operations
- **Library**: Implementation of LibraryManager with in-memory storage

### Controllers
- **LibraryController**: Handles console input/output and user interaction

## Data Flow
1. User selects an option from the menu
2. Controller captures input
3. Controller calls appropriate service method
4. Service performs business logic
5. Service updates data structures
6. Controller displays results

## Error Handling
- Validation for empty inputs
- Invalid ID handling
- Book availability checks
- Member borrowing limits (max 5 books)
- Proper error messages for all scenarios

## Usage

### Running the Application
```bash
go run main.go