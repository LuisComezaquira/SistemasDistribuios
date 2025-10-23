# Proyecto 1: Simple HTTP Server

Este es un servidor HTTP básico escrito en Go que demuestra conceptos fundamentales de sistemas distribuidos.

## Características

- Servidor HTTP que responde a múltiples endpoints
- Endpoints RESTful con respuestas JSON
- Registro de solicitudes con timestamps
- Fácil exposición a internet mediante ngrok

## Endpoints

- `GET /` - Página principal con mensaje de bienvenida
- `GET /api/status` - Estado del servidor en formato JSON
- `GET /api/info` - Información del proyecto en formato JSON

## Cómo ejecutar

1. Navega al directorio del proyecto:
```bash
cd proyecto1-simple-server
```

2. Inicializa el módulo de Go (primera vez):
```bash
go mod init proyecto1-simple-server
```

3. Ejecuta el servidor:
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## Exponer con ngrok

En otra terminal, ejecuta:
```bash
ngrok http 8080
```

ngrok te proporcionará una URL pública (ej: `https://xxxx.ngrok.io`) que puedes compartir para acceder a tu servidor desde cualquier lugar.

## Probar el servidor

```bash
# Probar endpoint raíz
curl http://localhost:8080/

# Probar endpoint de status
curl http://localhost:8080/api/status

# Probar endpoint de info
curl http://localhost:8080/api/info
```

## Conceptos de Sistemas Distribuidos

- **Comunicación HTTP**: Protocolo estándar para sistemas distribuidos
- **API RESTful**: Arquitectura común en servicios distribuidos
- **Exposición pública**: ngrok simula un entorno de producción distribuido
