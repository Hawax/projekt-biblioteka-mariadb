-- Utworzenie tabel

-- Tabela Users
CREATE TABLE Users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100),
    Email VARCHAR(100) UNIQUE
);

-- Tabela Authors
CREATE TABLE Authors (
    AuthorID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100),
    Biography TEXT
);

-- Tabela Publishers
CREATE TABLE Publishers (
    PublisherID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100),
    Address TEXT
);

-- Tabela Categories
CREATE TABLE Categories (
    CategoryID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100),
    Description TEXT
);


CREATE TABLE Books (
    BookID INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(100),
    AuthorID INT,
    PublisherID INT,
    CategoryID INT,
    Available BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (AuthorID) REFERENCES Authors(AuthorID),
    FOREIGN KEY (PublisherID) REFERENCES Publishers(PublisherID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- Tabela Loans
CREATE TABLE Loans (
    LoanID INT AUTO_INCREMENT PRIMARY KEY,
    BookID INT,
    UserID INT,
    LoanDate DATE,
    ReturnDate DATE,
    FOREIGN KEY (BookID) REFERENCES Books(BookID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Tabela Reservations
CREATE TABLE Reservations (
    ReservationID INT AUTO_INCREMENT PRIMARY KEY,
    BookID INT,
    UserID INT,
    ReservationDate DATE,
    FOREIGN KEY (BookID) REFERENCES Books(BookID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Tabela Reviews
CREATE TABLE Reviews (
    ReviewID INT AUTO_INCREMENT PRIMARY KEY,
    BookID INT,
    UserID INT,
    Rating INT,
    Comment TEXT,
    FOREIGN KEY (BookID) REFERENCES Books(BookID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);


DELIMITER //
CREATE TRIGGER AfterBookLoan
AFTER INSERT ON Loans
FOR EACH ROW
BEGIN
   UPDATE Books SET Available = FALSE WHERE BookID = NEW.BookID;
END;

CREATE TRIGGER BeforeBookReturn
BEFORE UPDATE ON Loans
FOR EACH ROW
BEGIN
   IF NEW.ReturnDate IS NOT NULL THEN
      UPDATE Books SET Available = TRUE WHERE BookID = OLD.BookID;
   END IF;
END;

//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE LoanBook(IN book_id INT, IN user_id INT)
BEGIN
   INSERT INTO Loans (BookID, UserID, LoanDate) VALUES (book_id, user_id, CURDATE());
END;

CREATE PROCEDURE ReturnBook(IN loan_id INT)
BEGIN
   UPDATE Loans SET ReturnDate = CURDATE() WHERE LoanID = loan_id;
END;

CREATE PROCEDURE ReserveBook(IN book_id INT, IN user_id INT)
BEGIN
   INSERT INTO Reservations (BookID, UserID, ReservationDate) VALUES (book_id, user_id, CURDATE());
END;
//
DELIMITER ;

DELIMITER //
CREATE FUNCTION CalculateAverageRating(book_id INT) RETURNS DECIMAL(10,2)
BEGIN
   DECLARE avg_rating DECIMAL(10,2);
   SELECT AVG(Rating) INTO avg_rating FROM Reviews WHERE BookID = book_id;
   RETURN avg_rating;
END;

CREATE FUNCTION CountUserLoans(user_id INT) RETURNS INT
BEGIN
   DECLARE loan_count INT;
   SELECT COUNT(*) INTO loan_count FROM Loans WHERE UserID = user_id;
   RETURN loan_count;
END;

CREATE FUNCTION CheckBookAvailability(book_id INT) RETURNS BOOLEAN
BEGIN
   DECLARE is_available BOOLEAN;
   SELECT Available INTO is_available FROM Books WHERE BookID = book_id;
   RETURN is_available;
END;
//
DELIMITER ;


CREATE VIEW AvailableBooks AS
SELECT * FROM Books WHERE Available = TRUE;

CREATE VIEW UserLoanHistory AS
SELECT Users.UserID, Users.Name, Books.Title, Loans.LoanDate, Loans.ReturnDate
FROM Users
JOIN Loans ON Users.UserID = Loans.UserID
JOIN Books ON Loans.BookID = Books.BookID;

CREATE VIEW TopRatedBooks AS
SELECT Books.BookID, Books.Title, AVG(Reviews.Rating) as AverageRating
FROM Books
JOIN Reviews ON Books.BookID = Reviews.BookID
GROUP BY Books.BookID
HAVING AverageRating >= 4.0;


