package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Mensaje representa la estructura de datos que se intercambia
type Mensaje struct {
	ID        int       `json:"id"`
	Contenido string    `json:"contenido"`
	Timestamp time.Time `json:"timestamp"`
	Cliente   string    `json:"cliente"`
}

// Respuesta representa la respuesta del servidor
type Respuesta struct {
	Exito   bool   `json:"exito"`
	Mensaje string `json:"mensaje"`
	ID      int    `json:"id,omitempty"`
}

var (
	mensajes []Mensaje
	mu       sync.Mutex
	contador int
)

func main() {
	mensajes = make([]Mensaje, 0)
	
	// Endpoint para recibir mensajes
	http.HandleFunc("/api/mensaje", manejarMensaje)
	
	// Endpoint para obtener todos los mensajes
	http.HandleFunc("/api/mensajes", obtenerMensajes)
	
	// Endpoint de estado
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":           "online",
			"mensajes_totales": len(mensajes),
			"timestamp":        time.Now(),
		})
	})

	port := ":8081"
	log.Printf("Servidor iniciado en http://localhost%s", port)
	log.Println("Endpoints disponibles:")
	log.Println("  - POST http://localhost:8081/api/mensaje")
	log.Println("  - GET  http://localhost:8081/api/mensajes")
	log.Println("  - GET  http://localhost:8081/api/status")
	log.Println("\nPara exponer con ngrok:")
	log.Println("  ngrok http 8081")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func manejarMensaje(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Respuesta{
			Exito:   false,
			Mensaje: "Método no permitido. Use POST",
		})
		return
	}

	var msg Mensaje
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Respuesta{
			Exito:   false,
			Mensaje: "Error al decodificar JSON: " + err.Error(),
		})
		return
	}

	mu.Lock()
	contador++
	msg.ID = contador
	msg.Timestamp = time.Now()
	mensajes = append(mensajes, msg)
	mu.Unlock()

	log.Printf("Mensaje recibido de %s: %s (ID: %d)", msg.Cliente, msg.Contenido, msg.ID)

	json.NewEncoder(w).Encode(Respuesta{
		Exito:   true,
		Mensaje: "Mensaje recibido correctamente",
		ID:      msg.ID,
	})
}

func obtenerMensajes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Respuesta{
			Exito:   false,
			Mensaje: "Método no permitido. Use GET",
		})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	log.Printf("Enviando lista de %d mensajes", len(mensajes))
	json.NewEncoder(w).Encode(mensajes)
}
