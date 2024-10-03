# Используем официальный образ Golang
FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN go build -o product-service .

# Запускаем приложение
CMD ["./product-service"]
