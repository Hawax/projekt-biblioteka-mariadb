package handlers

import (
	"books_rent/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	DB *sql.DB
}

func NewReviewHandler(db *sql.DB) *ReviewHandler {
	return &ReviewHandler{DB: db}
}

// GetReviews godoc
// @Summary Get a list of reviews
// @Description Get a list of all reviews
// @Tags reviews
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Review
// @Router /reviews [get]
func (h *ReviewHandler) GetReviews(c *gin.Context) {
	var reviews []models.Review
	rows, err := h.DB.Query("SELECT * FROM Reviews")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ReviewID, &review.BookID, &review.UserID, &review.Rating, &review.Comment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		reviews = append(reviews, review)
	}
	c.JSON(http.StatusOK, reviews)
}

// CreateReview godoc
// @Summary Create a new review
// @Description Add a new review to the database
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param review body models.Review true "Create Review"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(review.BookID, review.UserID, review.Rating, review.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Review created"})
}

// GetReviewByID godoc
// @Summary Get details of a specific review
// @Description Get details of a review given its ID
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Review ID"
// @Success 200 {object} models.Review
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [get]
func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var review models.Review
	err := h.DB.QueryRow("SELECT * FROM Reviews WHERE ReviewID = ?", id).Scan(&review.ReviewID, &review.BookID, &review.UserID, &review.Rating, &review.Comment)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, review)
}

// UpdateReview godoc
// @Summary Update a review
// @Description Update details of a review given its ID
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Review ID"
// @Param review body models.Review true "Update Review"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [put]
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.Exec("UPDATE Reviews SET BookID = ?, UserID = ?, Rating = ?, Comment = ? WHERE ReviewID = ?", review.BookID, review.UserID, review.Rating, review.Comment, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review updated"})
}

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review given its ID
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Review ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.DB.Exec("DELETE FROM Reviews WHERE ReviewID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}
