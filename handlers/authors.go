package handlers

import (
	"books_rent/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	DB *sql.DB
}

func NewAuthorHandler(db *sql.DB) *AuthorHandler {
	return &AuthorHandler{DB: db}
}

// GetAuthors godoc
// @Summary Get a list of authors
// @Description Get a list of all authors
// @Tags authors
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Author
// @Router /authors [get]
func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	var authors []models.Author
	rows, err := h.DB.Query("SELECT * FROM Authors")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.AuthorID, &author.Name, &author.Biography); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		authors = append(authors, author)
	}
	c.JSON(http.StatusOK, authors)
}

// CreateAuthor godoc
// @Summary Create a new author
// @Description Add a new author to the database
// @Tags authors
// @Accept  json
// @Produce  json
// @Param author body models.Author true "Create Author"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors [post]
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("INSERT INTO Authors (Name, Biography) VALUES (?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(author.Name, author.Biography)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Author created"})
}

// GetAuthorByID godoc
// @Summary Get details of a specific author
// @Description Get details of an author given their ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors/{id} [get]
func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author
	err := h.DB.QueryRow("SELECT * FROM Authors WHERE AuthorID = ?", id).Scan(&author.AuthorID, &author.Name, &author.Biography)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Author not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, author)
}

// UpdateAuthor godoc
// @Summary Update an author
// @Description Update details of an author given their ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Param author body models.Author true "Update Author"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author
	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.Exec("UPDATE Authors SET Name = ?, Biography = ? WHERE AuthorID = ?", author.Name, author.Biography, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author updated"})
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Description Delete an author given their ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.DB.Exec("DELETE FROM Authors WHERE AuthorID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
