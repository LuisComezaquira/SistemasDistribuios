package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Mensaje representa la estructura de datos que se envía
type Mensaje struct {
	Contenido string `json:"contenido"`
	Cliente   string `json:"cliente"`
}

// Respuesta representa la respuesta del servidor
type Respuesta struct {
	Exito   bool   `json:"exito"`
	Mensaje string `json:"mensaje"`
	ID      int    `json:"id,omitempty"`
}

// MensajeCompleto representa un mensaje almacenado en el servidor
type MensajeCompleto struct {
	ID        int       `json:"id"`
	Contenido string    `json:"contenido"`
	Timestamp time.Time `json:"timestamp"`
	Cliente   string    `json:"cliente"`
}

func main() {
	// URL del servidor (puede ser localhost o ngrok)
	serverURL := "http://localhost:8081"
	if len(os.Args) > 1 {
		serverURL = os.Args[1]
	}

	log.Printf("Cliente conectándose a: %s", serverURL)
	
	// Probar conexión
	if err := probarConexion(serverURL); err != nil {
		log.Printf("Advertencia: No se pudo conectar al servidor: %v", err)
		log.Println("Asegúrate de que el servidor esté ejecutándose")
	}

	// Enviar algunos mensajes de ejemplo
	clienteID := fmt.Sprintf("Cliente-%d", time.Now().Unix())
	
	mensajes := []string{
		"Hola desde el cliente!",
		"Este es un sistema distribuido en Go",
		"Mensaje de prueba número 3",
	}

	for i, contenido := range mensajes {
		log.Printf("\n--- Enviando mensaje %d ---", i+1)
		if err := enviarMensaje(serverURL, clienteID, contenido); err != nil {
			log.Printf("Error al enviar mensaje: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	// Obtener todos los mensajes
	log.Println("\n--- Obteniendo todos los mensajes ---")
	if err := obtenerMensajes(serverURL); err != nil {
		log.Printf("Error al obtener mensajes: %v", err)
	}
}

func probarConexion(serverURL string) error {
	resp, err := http.Get(serverURL + "/api/status")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var status map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return err
	}

	log.Printf("Estado del servidor: %v", status)
	return nil
}

func enviarMensaje(serverURL, cliente, contenido string) error {
	msg := Mensaje{
		Contenido: contenido,
		Cliente:   cliente,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error al codificar JSON: %v", err)
	}

	resp, err := http.Post(
		serverURL+"/api/mensaje",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("error al enviar solicitud: %v", err)
	}
	defer resp.Body.Close()

	var respuesta Respuesta
	if err := json.NewDecoder(resp.Body).Decode(&respuesta); err != nil {
		return fmt.Errorf("error al decodificar respuesta: %v", err)
	}

	if respuesta.Exito {
		log.Printf("✓ Mensaje enviado exitosamente (ID: %d)", respuesta.ID)
	} else {
		log.Printf("✗ Error del servidor: %s", respuesta.Mensaje)
	}

	return nil
}

func obtenerMensajes(serverURL string) error {
	resp, err := http.Get(serverURL + "/api/mensajes")
	if err != nil {
		return fmt.Errorf("error al obtener mensajes: %v", err)
	}
	defer resp.Body.Close()

	var mensajes []MensajeCompleto
	if err := json.NewDecoder(resp.Body).Decode(&mensajes); err != nil {
		return fmt.Errorf("error al decodificar mensajes: %v", err)
	}

	log.Printf("Total de mensajes en el servidor: %d\n", len(mensajes))
	for _, msg := range mensajes {
		log.Printf("  [%d] %s - %s (enviado: %s)",
			msg.ID,
			msg.Cliente,
			msg.Contenido,
			msg.Timestamp.Format("15:04:05"),
		)
	}

	return nil
}
