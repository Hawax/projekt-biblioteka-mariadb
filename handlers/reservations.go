package handlers

import (
	"books_rent/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ReservationHandler struct {
	DB *sql.DB
}

func NewReservationHandler(db *sql.DB) *ReservationHandler {
	return &ReservationHandler{DB: db}
}

// GetReservations godoc
// @Summary Get a list of reservations
// @Description Get a list of all reservations
// @Tags reservations
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Reservation
// @Router /reservations [get]
func (h *ReservationHandler) GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	rows, err := h.DB.Query("SELECT * FROM Reservations")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var reservation models.Reservation
		var reservationDate string // handling date as a string for simplicity
		if err := rows.Scan(&reservation.ReservationID, &reservation.BookID, &reservation.UserID, &reservationDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		reservation.ReservationDate, _ = time.Parse("2006-01-02", reservationDate)
		reservations = append(reservations, reservation)
	}
	c.JSON(http.StatusOK, reservations)
}

// CreateReservation godoc
// @Summary Create a new reservation
// @Description Add a new reservation to the database
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param reservation body models.Reservation true "Create Reservation"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservations [post]
func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.BindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("INSERT INTO Reservations (BookID, UserID, ReservationDate) VALUES (?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(reservation.BookID, reservation.UserID, reservation.ReservationDate.Format("2006-01-02"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created"})
}

// GetReservationByID godoc
// @Summary Get details of a specific reservation
// @Description Get details of a reservation given its ID
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "Reservation ID"
// @Success 200 {object} models.Reservation
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservations/{id} [get]
func (h *ReservationHandler) GetReservationByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var reservation models.Reservation
	var reservationDate string
	err := h.DB.QueryRow("SELECT * FROM Reservations WHERE ReservationID = ?", id).Scan(&reservation.ReservationID, &reservation.BookID, &reservation.UserID, &reservationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Reservation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	reservation.ReservationDate, _ = time.Parse("2006-01-02", reservationDate)
	c.JSON(http.StatusOK, reservation)
}

// UpdateReservation godoc
// @Summary Update a reservation
// @Description Update details of a reservation given its ID
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "Reservation ID"
// @Param reservation body models.Reservation true "Update Reservation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservations/{id} [put]
func (h *ReservationHandler) UpdateReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var reservation models.Reservation
	if err := c.BindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.Exec("UPDATE Reservations SET BookID = ?, UserID = ?, ReservationDate = ? WHERE ReservationID = ?", reservation.BookID, reservation.UserID, reservation.ReservationDate.Format("2006-01-02"), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation updated"})
}

// DeleteReservation godoc
// @Summary Delete a reservation
// @Description Delete a reservation given its ID
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "Reservation ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /reservations/{id} [delete]
func (h *ReservationHandler) DeleteReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.DB.Exec("DELETE FROM Reservations WHERE ReservationID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
}
