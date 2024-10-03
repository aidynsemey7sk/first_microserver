package main

import (
	"io"
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	log.Printf("Запрос: метод=%s, URL=%s", r.Method, r.URL)
}

func proxyAllUsers(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	req, err := http.NewRequest(r.Method, "http://user-service:8084/users", r.Body)
	if err != nil {
		http.Error(w, "Ошибка создания запроса к User Service", http.StatusInternalServerError)
		return
	}

	for key, value := range r.Header {
		req.Header[key] = value
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка при запросе к User Service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	for key, value := range resp.Header {
		w.Header()[key] = value
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Ошибка копирования ответа", http.StatusInternalServerError)
	}
}

func proxyUser(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	req, err := http.NewRequest(r.Method, "http://user-service:8084/user/"+r.URL.Path[len("/users/"):], r.Body)
	if err != nil {
		http.Error(w, "Ошибка создания запроса к User Service", http.StatusInternalServerError)
		return
	}

	for key, value := range r.Header {
		req.Header[key] = value
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка при запросе к User Service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	for key, value := range resp.Header {
		w.Header()[key] = value
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Ошибка копирования ответа", http.StatusInternalServerError)
	}
}

func proxyProduct(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	req, err := http.NewRequest(r.Method, "http://product-service:8086"+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "Ошибка создания запроса к Product Service", http.StatusInternalServerError)
		return
	}

	for key, value := range r.Header {
		req.Header[key] = value
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка при запросе к Product Service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	for key, value := range resp.Header {
		w.Header()[key] = value
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Ошибка копирования ответа", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/users/", proxyUser)       // Для конкретного пользователя
	http.HandleFunc("/users", proxyAllUsers)    // Для получения всех пользователей
	http.HandleFunc("/products/", proxyProduct) // Обработка запросов к продуктам
	log.Println("API Gateway is running on port 8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
