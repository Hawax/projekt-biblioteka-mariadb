package handlers

import (
	"books_rent/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	DB *sql.DB
}

func NewBookHandler(db *sql.DB) *BookHandler {
	return &BookHandler{DB: db}
}

// GetBooks godoc
// @Summary Get a list of books
// @Description Get a list of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book
	rows, err := h.DB.Query("SELECT * FROM Books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookID, &book.Title, &book.AuthorID, &book.PublisherID, &book.CategoryID, &book.Available); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Create Book"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID, Available) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.AuthorID, book.PublisherID, book.CategoryID, book.Available)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

// GetBookByID godoc
// @Summary Get details of a specific book
// @Description Get details of a book given its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	err := h.DB.QueryRow("SELECT * FROM Books WHERE BookID = ?", id).Scan(&book.BookID, &book.Title, &book.AuthorID, &book.PublisherID, &book.CategoryID, &book.Available)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update details of a book given its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Update Book"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.Exec("UPDATE Books SET Title = ?, AuthorID = ?, PublisherID = ?, CategoryID = ?, Available = ? WHERE BookID = ?", book.Title, book.AuthorID, book.PublisherID, book.CategoryID, book.Available, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book given its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.DB.Exec("DELETE FROM Books WHERE BookID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// GetAvailableBooks godoc
// @Summary Get available books
// @Description Get a list of all books that are available
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books/available [get]
func (h *BookHandler) GetAvailableBooks(c *gin.Context) {
	var books []models.Book
	rows, err := h.DB.Query("SELECT * FROM AvailableBooks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookID, &book.Title, &book.AuthorID, &book.PublisherID, &book.CategoryID, &book.Available); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

// GetTopRatedBooks godoc
// @Summary Get top-rated books
// @Description Get a list of books with an average rating of 4 or higher
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books/top-rated [get]
func (h *BookHandler) GetTopRatedBooks(c *gin.Context) {
	var books []models.Book
	rows, err := h.DB.Query("SELECT * FROM TopRatedBooks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		var averageRating float64
		if err := rows.Scan(&book.BookID, &book.Title, &averageRating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		book.AverageRating = averageRating
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}
