package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	CategoryID int    `json:"category_id"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var books []Book
var authors []Author
var categories []Category

var bookID, authorID, categoryID int = 1, 1, 1

func main() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookHandler)
	http.HandleFunc("/authors", authorsHandler)
	http.HandleFunc("/categories", categoriesHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Books Handlers
func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(books)
	case http.MethodPost:
		var b Book
		json.NewDecoder(r.Body).Decode(&b)
		b.ID = bookID
		bookID++
		books = append(books, b)
		json.NewEncoder(w).Encode(b)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		for i, b := range books {
			if b.ID == id {
				json.NewDecoder(r.Body).Decode(&books[i])
				books[i].ID = id
				json.NewEncoder(w).Encode(books[i])
				return
			}
		}
		http.NotFound(w, r)
	case http.MethodDelete:
		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.NotFound(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Authors Handlers
func authorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(authors)
	case http.MethodPost:
		var a Author
		json.NewDecoder(r.Body).Decode(&a)
		a.ID = authorID
		authorID++
		authors = append(authors, a)
		json.NewEncoder(w).Encode(a)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Categories Handlers
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(categories)
	case http.MethodPost:
		var c Category
		json.NewDecoder(r.Body).Decode(&c)
		c.ID = categoryID
		categoryID++
		categories = append(categories, c)
		json.NewEncoder(w).Encode(c)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
