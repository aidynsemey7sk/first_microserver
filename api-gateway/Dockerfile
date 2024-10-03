# Используем официальный образ Golang
FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN go build -o api-gateway .

# Запускаем приложение
CMD ["./api-gateway"]
