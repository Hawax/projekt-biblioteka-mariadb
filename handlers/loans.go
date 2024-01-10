package handlers

import (
	"books_rent/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type LoanHandler struct {
	DB *sql.DB
}

func NewLoanHandler(db *sql.DB) *LoanHandler {
	return &LoanHandler{DB: db}
}

// GetLoans godoc
// @Summary Get a list of loans
// @Description Get a list of all loans
// @Tags loans
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Loan
// @Router /loans [get]

func (h *LoanHandler) GetLoans(c *gin.Context) {
	var loans []models.Loan
	rows, err := h.DB.Query("SELECT LoanID, BookID, UserID, LoanDate, ReturnDate FROM Loans")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var loan models.Loan
		var loanDate sql.NullString
		var returnDate sql.NullString
		if err := rows.Scan(&loan.LoanID, &loan.BookID, &loan.UserID, &loanDate, &returnDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if loanDate.Valid {
			parsedDate, _ := time.Parse("2006-01-02", loanDate.String)
			loan.LoanDate = &parsedDate
		}

		if returnDate.Valid {
			parsedDate, _ := time.Parse("2006-01-02", returnDate.String)
			loan.ReturnDate = &parsedDate
		}

		loans = append(loans, loan)
	}
	c.JSON(http.StatusOK, loans)
}

// CreateLoan godoc
// @Summary Create a new loan
// @Description Add a new loan to the database
// @Tags loans
// @Accept  json
// @Produce  json
// @Param loan body models.Loan true "Create Loan"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /loans [post]
func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var loan models.Loan
	if err := c.BindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("INSERT INTO Loans (BookID, UserID, LoanDate, ReturnDate) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(loan.BookID, loan.UserID, loan.LoanDate.Format("2006-01-02"), loan.ReturnDate.Format("2006-01-02"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Loan created"})
}

// GetLoanByID godoc
// @Summary Get details of a specific loan
// @Description Get details of a loan given its ID
// @Tags loans
// @Accept  json
// @Produce  json
// @Param id path int true "Loan ID"
// @Success 200 {object} models.Loan
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /loans/{id} [get]
func (h *LoanHandler) GetLoanByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var loan models.Loan
	var loanDate sql.NullString
	var returnDate sql.NullString

	err := h.DB.QueryRow("SELECT LoanID, BookID, UserID, LoanDate, ReturnDate FROM Loans WHERE LoanID = ?", id).Scan(&loan.LoanID, &loan.BookID, &loan.UserID, &loanDate, &returnDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Loan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if loanDate.Valid {
		parsedDate, _ := time.Parse("2006-01-02", loanDate.String)
		loan.LoanDate = &parsedDate
	}

	if returnDate.Valid {
		parsedDate, _ := time.Parse("2006-01-02", returnDate.String)
		loan.ReturnDate = &parsedDate
	}

	c.JSON(http.StatusOK, loan)
}

// UpdateLoan godoc
// @Summary Update a loan
// @Description Update details of a loan given its ID
// @Tags loans
// @Accept  json
// @Produce  json
// @Param id path int true "Loan ID"
// @Param loan body models.Loan true "Update Loan"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /loans/{id} [put]
func (h *LoanHandler) UpdateLoan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var loan models.Loan
	if err := c.BindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.Exec("UPDATE Loans SET BookID = ?, UserID = ?, LoanDate = ?, ReturnDate = ? WHERE LoanID = ?", loan.BookID, loan.UserID, loan.LoanDate.Format("2006-01-02"), loan.ReturnDate.Format("2006-01-02"), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Loan updated"})
}

// DeleteLoan godoc
// @Summary Delete a loan
// @Description Delete a loan given its ID
// @Tags loans
// @Accept  json
// @Produce  json
// @Param id path int true "Loan ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /loans/{id} [delete]
func (h *LoanHandler) DeleteLoan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.DB.Exec("DELETE FROM Loans WHERE LoanID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted"})
}

// GetUserLoanHistory godoc
// @Summary Get user loan history
// @Description Get the loan history of all users
// @Tags loans
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UserLoanHistory
// @Router /loans/history [get]
func (h *LoanHandler) GetUserLoanHistory(c *gin.Context) {
	var histories []models.UserLoanHistory
	rows, err := h.DB.Query("SELECT * FROM UserLoanHistory")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var history models.UserLoanHistory
		var loanDate sql.NullString
		var returnDate sql.NullString

		if err := rows.Scan(&history.UserID, &history.UserName, &history.BookTitle, &loanDate, &returnDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if loanDate.Valid {
			parsedDate, _ := time.Parse("2006-01-02", loanDate.String)
			history.LoanDate = &parsedDate
		}

		if returnDate.Valid {
			parsedDate, _ := time.Parse("2006-01-02", returnDate.String)
			history.ReturnDate = &parsedDate
		}

		histories = append(histories, history)
	}
	c.JSON(http.StatusOK, histories)
}
