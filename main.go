package main

import (
	"database/sql"
	"log"

	_ "books_rent/docs"
	"books_rent/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:new_password@tcp(db)/library")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}))

	bookHandler := handlers.NewBookHandler(db)
	authorHandler := handlers.NewAuthorHandler(db)
	categoriesHandler := handlers.NewCategoryHandler(db)
	loansHandler := handlers.NewLoanHandler(db)
	reservationHandler := handlers.NewReservationHandler(db)
	reviewsHandler := handlers.NewReviewHandler(db)
	userHandler := handlers.NewUserHandler(db)

	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/available", bookHandler.GetAvailableBooks)
	r.GET("/books/top-rated", bookHandler.GetTopRatedBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	r.GET("/authors", authorHandler.GetAuthors)
	r.POST("/authors", authorHandler.CreateAuthor)
	r.GET("/authors/:id", authorHandler.GetAuthorByID)
	r.PUT("/authors/:id", authorHandler.UpdateAuthor)
	r.DELETE("/authors/:id", authorHandler.DeleteAuthor)

	r.GET("/categories", categoriesHandler.GetCategories)
	r.POST("/categories", categoriesHandler.CreateCategory)
	r.GET("/categories/:id", categoriesHandler.GetCategoryByID)
	r.PUT("/categories/:id", categoriesHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoriesHandler.DeleteCategory)

	r.GET("/loans", loansHandler.GetLoans)
	r.POST("/loans", loansHandler.CreateLoan)
	r.GET("/loans/:id", loansHandler.GetLoanByID)
	r.PUT("/loans/:id", loansHandler.UpdateLoan)
	r.DELETE("/loans/:id", loansHandler.DeleteLoan)
	r.GET("/loans/history", loansHandler.GetUserLoanHistory)

	r.GET("/reservations", reservationHandler.GetReservations)
	r.POST("/reservations", reservationHandler.CreateReservation)
	r.GET("/reservations/:id", reservationHandler.GetReservationByID)
	r.PUT("/reservations/:id", reservationHandler.UpdateReservation)
	r.DELETE("/reservations/:id", reservationHandler.DeleteReservation)

	r.GET("/reviews", reviewsHandler.GetReviews)
	r.POST("/reviews", reviewsHandler.CreateReview)
	r.GET("/reviews/:id", reviewsHandler.GetReviewByID)
	r.PUT("/reviews/:id", reviewsHandler.UpdateReview)
	r.DELETE("/reviews/:id", reviewsHandler.DeleteReview)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// Start the server
	r.Run(":8080")
}
