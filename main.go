package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/teris-io/shortid"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

var urlMap = make(map[string]string)
var mu sync.Mutex

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 1. Verificar o método HTTP (deve ser POST)
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	// 2. Decodificar o JSON da requisição
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Requisição inválida. Certifique-se de que o JSON está correto: {\"long_url\": \"sua_url\"}", http.StatusBadRequest)
		return
	}

	// 3. Validar a URL longa
	if req.LongURL == "" {
		http.Error(w, "O campo 'long_url' é obrigatório.", http.StatusBadRequest)
		return
	}

	// 4. Gerar um código curto único usando a função de nível de pacote
	shortCode, err := shortid.Generate()
	if err != nil {
		http.Error(w, "Erro interno ao gerar o código curto.", http.StatusInternalServerError)
		log.Printf("Erro ao gerar shortID: %v", err)
		return
	}

	// 5. Armazenar o mapeamento (com proteção de concorrência)
	mu.Lock()
	urlMap[shortCode] = req.LongURL
	mu.Unlock()

	// 6. Construir a URL curta completa
	shortURL := fmt.Sprintf("http://%s/%s", r.Host, shortCode)

	// 7. Enviar a resposta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 1. Verificar o método HTTP (deve ser GET)
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	// 2. Extrair o código curto da URL
	shortCode := r.URL.Path[1:]

	// 3. Buscar a URL longa correspondente (com proteção de concorrência)
	mu.Lock()
	longURL, ok := urlMap[shortCode]
	mu.Unlock()

	// 4. Verificar se o código curto foi encontrado
	if !ok {
		http.Error(w, "URL curta não encontrada.", http.StatusNotFound)
		return
	}

	// 5. Redirecionar para a URL longa
	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shortenURL)
	http.HandleFunc("/", redirectURL)

	port := ":8080"
	log.Printf("Servidor GoURL iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
