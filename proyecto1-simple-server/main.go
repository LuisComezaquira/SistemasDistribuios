package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Handler para la ruta raíz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Solicitud recibida: %s %s desde %s", r.Method, r.URL.Path, r.RemoteAddr)
		fmt.Fprintf(w, "¡Hola desde el servidor de Sistemas Distribuidos!\n")
		fmt.Fprintf(w, "Hora del servidor: %s\n", time.Now().Format(time.RFC3339))
	})

	// Handler para /api/status
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"online","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	// Handler para /api/info
	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"proyecto":"Sistemas Distribuidos","tecnologias":["Go","ngrok"],"version":"1.0"}`)
	})

	port := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s", port)
	log.Println("Presiona Ctrl+C para detener el servidor")
	log.Println("\nEndpoints disponibles:")
	log.Println("  - http://localhost:8080/")
	log.Println("  - http://localhost:8080/api/status")
	log.Println("  - http://localhost:8080/api/info")
	log.Println("\nPara exponer con ngrok, ejecuta en otra terminal:")
	log.Println("  ngrok http 8080")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
