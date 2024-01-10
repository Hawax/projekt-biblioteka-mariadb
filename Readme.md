# README - Wypożyczalnia Książek

## Opis Projektu
Projekt "Wypożyczalnia Książek" to aplikacja backendowa stworzona w języku Go, używająca frameworka Gin do obsługi REST API oraz MariaDB jako bazy danych. Aplikacja umożliwia zarządzanie księgozbiorem, wypożyczeniami, rezerwacjami oraz recenzjami książek. Dzięki wykorzystaniu Swaggera, zapewnia także interaktywny interfejs dokumentacji API.

## Funkcjonalności
- CRUD (Create, Read, Update, Delete) dla książek, użytkowników, autorów, wydawców, kategorii.
- Zarządzanie wypożyczeniami i rezerwacjami książek.
- Dodawanie recenzji do książek.
- Wyświetlanie dostępnych książek i książek o wysokiej ocenie.
- Przeglądanie historii wypożyczeń użytkowników.

## Uruchomienie Projektu
Projekt wykorzystuje Docker i Docker Compose do łatwego uruchomienia aplikacji wraz z bazą danych.

### Wymagania
- Zainstalowany Docker
- Zainstalowany Docker Compose

### Kroki Uruchomienia
1. Sklonuj repozytorium projektu na swoją maszynę.
2. W głównym katalogu projektu znajdują się pliki `Dockerfile` oraz `docker-compose.yml`. Upewnij się, że są one obecne.
3. Otwórz terminal w katalogu projektu i uruchom następujące polecenie:

   ```
   docker-compose up --build
   ```

   To polecenie zbuduje obrazy Docker i uruchomi kontenery dla aplikacji oraz bazy danych.

4. Po uruchomieniu, aplikacja będzie dostępna pod adresem `http://localhost:8080`.
5. Dokumentacja API w formacie Swagger jest dostępna pod adresem `http://localhost:8080/swagger/index.html`.

### Struktura Projektu
- `/handlers` - Zawiera handlery obsługujące różne endpointy API.
- `/models` - Definicje modeli danych używanych w aplikacji.
- `main.go` - Główny plik aplikacji, konfiguruje i uruchamia serwer.
- `Dockerfile` - Instrukcje do stworzenia obrazu Docker dla aplikacji.
- `docker-compose.yml` - Konfiguracja Docker Compose do uruchomienia aplikacji wraz z bazą danych.
- `database.sql` - Skrypt SQL do stworzenia schematu bazy danych.
- `dummy_data.sql` - Skrypt SQL do wypełnienia bazy danych przykładowymi danymi.