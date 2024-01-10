package handlers

import (
	"books_rent/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	DB *sql.DB
}

func NewCategoryHandler(db *sql.DB) *CategoryHandler {
	return &CategoryHandler{DB: db}
}

// GetCategories godoc
// @Summary Get a list of categories
// @Description Get a list of all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	var categories []models.Category
	rows, err := h.DB.Query("SELECT * FROM Categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.CategoryID, &category.Name, &category.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, category)
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Add a new category to the database
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body models.Category true "Create Category"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("INSERT INTO Categories (Name, Description) VALUES (?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Name, category.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

// GetCategoryByID godoc
// @Summary Get details of a specific category
// @Description Get details of a category given its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	err := h.DB.QueryRow("SELECT * FROM Categories WHERE CategoryID = ?", id).Scan(&category.CategoryID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update details of a category given its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Update Category"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.Exec("UPDATE Categories SET Name = ?, Description = ? WHERE CategoryID = ?", category.Name, category.Description, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category given its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.DB.Exec("DELETE FROM Categories WHERE CategoryID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
