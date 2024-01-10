#!/bin/bash

# Database connection configuration
DB_HOST="localhost"
DB_USER="hawax"
DB_PASS="new_password"
DB_NAME="library"



# Funkcja do uruchamiania zapytań SQL
execute_sql() {
    local sql_query=$1
    echo "$sql_query" | mariadb -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME"
}

execute_sql_no_db() {
    local sql_query=$1
    echo "$sql_query" | mariadb -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" 
}

recreate_database() {
    echo "Dropping the existing database..."
    execute_sql "DROP DATABASE $DB_NAME;"

    echo "Creating a new database..."
    execute_sql_no_db "CREATE DATABASE $DB_NAME;"
    echo "Database recreated."
}

# Tworzenie bazy danych i tabel
echo "Tworzenie bazy danych i tabel..."
execute_sql "SOURCE database.sql"

echo "Wstawianie danych do tabel Authors, Publishers, Categories, i Users..."
execute_sql "INSERT INTO Authors (Name, Biography) VALUES ('Autor Testowy', 'Biografia testowa');"
execute_sql "INSERT INTO Publishers (Name, Address) VALUES ('Wydawnictwo Testowe', 'Adres testowy');"
execute_sql "INSERT INTO Categories (Name, Description) VALUES ('Kategoria Testowa', 'Opis testowy');"
execute_sql "INSERT INTO Users (Name, Email) VALUES ('Testowy Użytkownik', 'test@example.com');"
execute_sql "INSERT INTO Authors (Name, Biography) VALUES ('Inny Autor', 'Inna biografia');"
execute_sql "INSERT INTO Publishers (Name, Address) VALUES ('Inne Wydawnictwo', 'Inny adres');"
execute_sql "INSERT INTO Categories (Name, Description) VALUES ('Inna Kategoria', 'Inny opis');"


echo "Wstawianie danych do tabeli Books..."
execute_sql "INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID) VALUES ('Testowa Książka', 1, 1, 1);"
# Testowanie triggerów, procedur i widoków


echo "Wstawianie i testowanie recenzji..."
execute_sql "INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (1, 1, 5, 'Świetna książka');"
execute_sql "INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (1, 1, 4, 'Nawet Dobra książka');"
execute_sql "SELECT CalculateAverageRating(1);"

echo "Testowanie triggerów, procedur i widoków..."

# Testowanie triggerów
echo "Testowanie triggera AfterBookLoan..."
execute_sql "INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID) VALUES ('Testowa Książka', 1, 1, 1);"
execute_sql "INSERT INTO Users (Name, Email) VALUES ('Testowy Użytkownik', 'test@example.com');"
execute_sql "INSERT INTO Loans (BookID, UserID, LoanDate) VALUES (1, 1, CURDATE());"
execute_sql "SELECT * FROM Books WHERE BookID = 1;"

echo "Testowanie triggera BeforeBookReturn..."
execute_sql "UPDATE Loans SET ReturnDate = CURDATE() WHERE LoanID = 1;"
execute_sql "SELECT * FROM Books WHERE BookID = 1;"

# Testowanie procedur
echo "Testowanie procedury LoanBook..."
execute_sql "CALL LoanBook(1, 1);"

echo "Testowanie procedury ReturnBook..."
execute_sql "CALL ReturnBook(1);"

echo "Wstawianie dodatkowych danych do tabeli Books..."
execute_sql "INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID) VALUES ('Inna Książka', 2, 2, 2);"

echo "Testowanie procedury LoanBook z innymi książkami..."
execute_sql "CALL LoanBook(2, 1);"

echo "Testowanie procedury ReturnBook z innymi wypożyczeniami..."
execute_sql "CALL ReturnBook(2);"



echo "Testowanie historii wypożyczeń użytkownika..."
execute_sql "SELECT * FROM UserLoanHistory WHERE UserID = 1;"

# Testowanie widoków
echo "Testowanie widoku AvailableBooks..."
execute_sql "SELECT * FROM AvailableBooks;"

echo "Testowanie widoku UserLoanHistory..."
execute_sql "SELECT * FROM UserLoanHistory;"

echo "Testowanie widoku TopRatedBooks..."
execute_sql "SELECT * FROM TopRatedBooks;"

echo "Testy zakończone."



#recreate_database
