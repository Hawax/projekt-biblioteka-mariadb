-- Insert dummy data into Users
INSERT INTO Users (Name, Email) VALUES ('Jan Kowalski', 'jan.kowalski@example.com');
INSERT INTO Users (Name, Email) VALUES ('Anna Nowak', 'anna.nowak@example.com');
INSERT INTO Users (Name, Email) VALUES ('Piotr Wiśniewski', 'piotr.wisniewski@example.com');
INSERT INTO Users (Name, Email) VALUES ('Katarzyna Zielińska', 'katarzyna.zielinska@example.com');

-- Insert dummy data into Authors
INSERT INTO Authors (Name, Biography) VALUES ('Henryk Sienkiewicz', 'Polish novelist and Nobel Prize laureate');
INSERT INTO Authors (Name, Biography) VALUES ('Stanisław Lem', 'A prominent Polish science fiction writer');
INSERT INTO Authors (Name, Biography) VALUES ('Adam Mickiewicz', 'A principal figure in Polish Romanticism');
INSERT INTO Authors (Name, Biography) VALUES ('Wislawa Szymborska', 'Nobel Prize-winning Polish poet');

-- Insert dummy data into Publishers
INSERT INTO Publishers (Name, Address) VALUES ('Wydawnictwo Literackie', 'Kraków, Poland');
INSERT INTO Publishers (Name, Address) VALUES ('Znak', 'Kraków, Poland');
INSERT INTO Publishers (Name, Address) VALUES ('WAB', 'Warszawa, Poland');
INSERT INTO Publishers (Name, Address) VALUES ('Czytelnik', 'Warszawa, Poland');

-- Insert dummy data into Categories
INSERT INTO Categories (Name, Description) VALUES ('Fiction', 'Fiction books');
INSERT INTO Categories (Name, Description) VALUES ('Science Fiction', 'Science fiction and fantasy');
INSERT INTO Categories (Name, Description) VALUES ('Poetry', 'Collections of poetry');
INSERT INTO Categories (Name, Description) VALUES ('History', 'Historical books and biographies');

-- Insert dummy data into Books
INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID, Available) VALUES ('Quo Vadis', 1, 1, 4, TRUE);
INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID, Available) VALUES ('Solaris', 2, 2, 2, FALSE);
INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID, Available) VALUES ('Pan Tadeusz', 3, 3, 1, TRUE);
INSERT INTO Books (Title, AuthorID, PublisherID, CategoryID, Available) VALUES ('Miracle Fair', 4, 4, 3, TRUE);

-- Insert dummy data into Loans
INSERT INTO Loans (BookID, UserID, LoanDate, ReturnDate) VALUES (1, 1, '2024-01-01', NULL);
INSERT INTO Loans (BookID, UserID, LoanDate, ReturnDate) VALUES (3, 2, '2024-01-05', '2024-02-05');

-- Insert dummy data into Reservations
INSERT INTO Reservations (BookID, UserID, ReservationDate) VALUES (2, 3, '2024-01-10');
INSERT INTO Reservations (BookID, UserID, ReservationDate) VALUES (4, 4, '2024-01-15');

-- Insert dummy data into Reviews
INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (1, 1, 5, 'Klasyczna powieść historyczna, polecam!');
INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (2, 2, 4, 'Interesująca książka z gatunku science fiction.');
INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (3, 3, 5, 'Najlepsza polska epopeja narodowa.');
INSERT INTO Reviews (BookID, UserID, Rating, Comment) VALUES (4, 4, 4, 'Piękne wiersze, bardzo liryczne.');
