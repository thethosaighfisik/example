package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

// Функция подключения к БД
func ConnectDB() (*sql.DB, error) {
    // Загружаем .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Ошибка загрузки .env:", err)
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    // Проверяем подключение
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    log.Println("✅ Успешное подключение к базе данных")
    return db, nil
}

