package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/", http.RedirectHandler("/ready", http.StatusSeeOther))
	http.HandleFunc("/alive", liveness)
	http.HandleFunc("/ready", readiness)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), getEnvOrDefault("PORT", "8080"))
	log.Printf("Starting server on: '%s'\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readiness(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"name":     os.Getenv("SERVICE_NAME"),
		"secret":   maskSecret(os.Getenv("SERVICE_SECRET")),
		"database": os.Getenv("DATABASE_NAME"),
	}

	jsonBytes, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func maskSecret(secret string) string {
	if len(secret) > 2 {
		secret = secret[:1] + strings.Repeat("*", len(secret)-2) + secret[len(secret)-1:]
	}
	return secret
}
