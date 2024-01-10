package models

import "time"

type Book struct {
	BookID        int     `json:"book_id"`
	Title         string  `json:"title"`
	AuthorID      int     `json:"author_id"`
	PublisherID   int     `json:"publisher_id"`
	CategoryID    int     `json:"category_id"`
	Available     bool    `json:"available"`
	AverageRating float64 `json:"average_rating,omitempty"`
}

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type Author struct {
	AuthorID  int    `json:"author_id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type Publisher struct {
	PublisherID int    `json:"publisher_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

type Category struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Loan struct {
	LoanID     int        `json:"loan_id"`
	BookID     int        `json:"book_id"`
	UserID     int        `json:"user_id"`
	LoanDate   *time.Time `json:"loan_date"`
	ReturnDate *time.Time `json:"return_date"`
}

type Reservation struct {
	ReservationID   int       `json:"reservation_id"`
	BookID          int       `json:"book_id"`
	UserID          int       `json:"user_id"`
	ReservationDate time.Time `json:"reservation_date"`
}

type Review struct {
	ReviewID int    `json:"review_id"`
	BookID   int    `json:"book_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}

type UserLoanHistory struct {
	UserID     int        `json:"user_id"`
	UserName   string     `json:"user_name"`
	BookTitle  string     `json:"book_title"`
	LoanDate   *time.Time `json:"loan_date"`
	ReturnDate *time.Time `json:"return_date"`
}
