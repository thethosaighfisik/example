package handlers

import (
    "encoding/json"
    "net/http"
    "auth_service/db"
    "auth_service/models"
    "log"
)

// Обработчик регистрации пользователя
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Некорректные данные", http.StatusBadRequest)
        return
    }

    // Проверяем, чтобы email и пароль были не пустыми
    if user.Email == "" || user.Password == "" {
        http.Error(w, "Email и пароль обязательны", http.StatusBadRequest)
        return
    }

    // Подключаемся к БД
    database, err := db.ConnectDB()
    if err != nil {
        http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
        return
    }
    defer database.Close()

    // Вызываем функцию регистрации пользователя
    err = models.CreateUser(database, user.Email, user.Password)
    if err != nil {
        http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
        log.Println("Ошибка регистрации:", err)
        return
    }

    // Успешный ответ
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно зарегистрирован"})
}

