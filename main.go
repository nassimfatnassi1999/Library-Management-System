package main

import (
	"fmt"
)

type LibraryInterface interface {
	addBook(b *Book)
	addUser(u *User)
	removeBook(b *Book)
	removeUser(u *User)
	searchBook(id int) *Book
	searchUser(id int) *User
	issueBook(u *User, b *Book)
	returnBook(u *User, b *Book)
	getAllBooks() []Book
	getAllUsers() []User
	getBooksIssuedToUser(u *User) []Book
	getUsersWhoIssuedBook(b *Book) []User
	getUserDetails() *User
	getBookDetails() *Book
}

type Date struct {
	Day   int
	Month int
	Year  int
}
type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	PublishedAt Date
}
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Books     []Book
}
type Library struct {
	Books []Book
	Users []User
}

// function to add book to library
func (l *Library) addBook(b *Book) {
	l.Books = append(l.Books, *b)
}

// function to add user to library
func (l *Library) addUser(u *User) {
	l.Users = append(l.Users, *u)
}

// function to remove book from library
func (l *Library) removeBook(b *Book) {
	for i, book := range l.Books {
		if book.ID == b.ID {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			break
		}
	}
}

// function to remove user from library
func (l *Library) removeUser(u *User) {
	for i, user := range l.Users {
		if user.ID == u.ID {
			l.Users = append(l.Users[:i], l.Users[i+1:]...)
			break
		}
	}
}

// function to search book
func (l *Library) searchBook(id int) *Book {
	for i, book := range l.Books {
		if book.ID == id {
			return &l.Books[i]
		}
	}
	return nil
}

// function to search user
func (l *Library) searchUser(id int) *User {
	for i, user := range l.Users {
		if user.ID == id {
			return &l.Users[i]
		}
	}
	return nil
}

// function to associate book with user
func issueBook(u *User, b *Book) {
	u.Books = append(u.Books, *b)
}

// function to remove book from user
func returnBook(u *User, b *Book) {
	for i, book := range u.Books {
		if book.ID == b.ID {
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			break
		}
	}
}

// function to get all books
func (l *Library) getAllBooks() []Book {
	return l.Books
}

// function to get all users
func (l *Library) getAllUsers() []User {
	return l.Users
}

// function to get all books issued to user
func getBooksIssuedToUser(u *User) []Book {
	return u.Books
}

// function to get all users who have issued book
func (l *Library) getUsersWhoIssuedBook(b *Book) []User {
	var users []User
	for _, user := range l.Users {
		for _, book := range user.Books {
			if book.ID == b.ID {
				users = append(users, user)
			}
		}
	}
	return users
}

// function to get user details from cli input
func getUserDetails() *User {
	var id int
	var firstName, lastName, email string
	fmt.Println("Enter User ID:")
	fmt.Scan(&id)
	fmt.Println("Enter First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter Last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter Email:")
	fmt.Scan(&email)
	return &User{id, firstName, lastName, email, []Book{}}
}

// function to get book details from cli input
func getBookDetails() *Book {
	var id int
	var title, author, isbn string
	var day, month, year int
	fmt.Println("Enter Book ID:")
	fmt.Scan(&id)
	fmt.Println("Enter Title:")
	fmt.Scan(&title)
	fmt.Println("Enter Author:")
	fmt.Scan(&author)
	fmt.Println("Enter ISBN:")
	fmt.Scan(&isbn)
	fmt.Println("Enter Published Day:")
	fmt.Scan(&day)
	fmt.Println("Enter Published Month:")
	fmt.Scan(&month)
	fmt.Println("Enter Published Year:")
	fmt.Scan(&year)
	return &Book{id, title, author, isbn, Date{day, month, year}}
}

func main() {
	// Initialize the Library
	library := &Library{
		Users: []User{},
		Books: []Book{},
	}

	for {
		fmt.Println("\n Welcome To Library Management System")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Add a User")
		fmt.Println("3. Remove a Book")
		fmt.Println("4. Remove a User")
		fmt.Println("5. Search for a Book")
		fmt.Println("6. Search for a User")
		fmt.Println("7. Issue a Book")
		fmt.Println("8. Return a Book")
		fmt.Println("9. Get All Books")
		fmt.Println("10. Get All Users")
		fmt.Println("11. Get Books Issued to a User")
		fmt.Println("12. Get Users Who Issued a Book")
		fmt.Println("13. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			book := getBookDetails()
			library.addBook(book)

		case 2:
			user := getUserDetails()
			library.addUser(user)

		case 3:
			var bookID int
			fmt.Print("Enter Book ID to remove: ")
			_, err := fmt.Scanf("%d", &bookID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			book := library.searchBook(bookID)
			if book != nil {
				library.removeBook(book)
				fmt.Println("Book removed successfully.")
			} else {
				fmt.Println("Book not found.")
			}

		case 4:
			var userID int
			fmt.Print("Enter User ID to remove: ")
			_, err := fmt.Scan(&userID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			user := library.searchUser(userID)
			if user != nil {
				library.removeUser(user)
			} else {
				fmt.Println("User not found.")
			}

		case 5:
			var bookID int
			fmt.Print("Enter Book ID to search: ")
			_, err := fmt.Scan(&bookID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			book := library.searchBook(bookID)
			if book != nil {
				fmt.Printf("Book Found: ID=%d, Title=%s, Author=%s\n", book.ID, book.Title, book.Author)
			} else {
				fmt.Println("Book not found.")
			}

		case 6:
			var userID int
			fmt.Print("Enter User ID to search: ")
			_, err := fmt.Scan(&userID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			user := library.searchUser(userID)
			if user != nil {
				fmt.Printf("User Found: ID=%d, Name=%s %s , Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
			} else {
				fmt.Println("User not found.")
			}

		case 7:
			var userID int
			fmt.Print("Enter User ID to issue book: ")
			_, err := fmt.Scan(&userID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			user := library.searchUser(userID)
			if user != nil {
				var bookID int
				fmt.Print("Enter Book ID to issue: ")
				_, err := fmt.Scan(&bookID)
				if err != nil {
					fmt.Println("Invalid input, please enter a valid integer ID.")
					continue
				}

				book := library.searchBook(bookID)
				if book != nil {
					issueBook(user, book)
				} else {
					fmt.Println("Book not found.")
				}
			} else {
				fmt.Println("User not found.")
			}

		case 8:
			var userID int
			fmt.Print("Enter User ID to return book: ")
			_, err := fmt.Scan(&userID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			user := library.searchUser(userID)
			if user != nil {
				var bookID int
				fmt.Print("Enter Book ID to return: ")
				_, err := fmt.Scan(&bookID)
				if err != nil {
					fmt.Println("Invalid input, please enter a valid integer ID.")
					continue
				}

				book := library.searchBook(bookID)
				if book != nil {
					returnBook(user, book)
				} else {
					fmt.Println("Book not found.")
				}
			} else {
				fmt.Println("User not found.")
			}

		case 9:
			books := library.getAllBooks()
			if len(books) == 0 {
				fmt.Println("No books available.")
			} else {
				fmt.Println("Books in the Library:")
				for _, book := range books {
					fmt.Printf("ID=%d, Title=%s, Author=%s\n", book.ID, book.Title, book.Author)
				}
			}

		case 10:
			users := library.getAllUsers()
			if len(users) == 0 {
				fmt.Println("No users registered.")
			} else {
				fmt.Println("Users in the Library:")
				for _, user := range users {
					fmt.Printf("User Found: ID=%d, Name=%s %s , Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
				}
			}

		case 11:
			var userID int
			fmt.Print("Enter User ID to view issued books: ")
			_, err := fmt.Scan(&userID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			user := library.searchUser(userID)
			if user != nil {
				books := getBooksIssuedToUser(user)
				if len(books) == 0 {
					fmt.Println("No books issued to this user.")
				} else {
					fmt.Println("Books issued to the user:")
					for _, book := range books {
						fmt.Printf("ID=%d, Title=%s, Author=%s\n", book.ID, book.Title, book.Author)
					}
				}
			} else {
				fmt.Println("User not found.")
			}

		case 12:
			var bookID int
			fmt.Print("Enter Book ID to view users who issued it: ")
			_, err := fmt.Scan(&bookID)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid integer ID.")
				continue
			}

			book := library.searchBook(bookID)
			if book != nil {
				users := library.getUsersWhoIssuedBook(book)
				if len(users) == 0 {
					fmt.Println("No users have issued this book.")
				} else {
					fmt.Println("Users who issued this book:")
					for _, user := range users {
						fmt.Printf("User Found: ID=%d, Name=%s %s , Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
					}
				}
			} else {
				fmt.Println("Book not found.")
			}

		case 13:
			fmt.Println("Thank you for using the Library Management System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
