# Proyecto 2: Client-Server Communication

Este proyecto demuestra comunicación cliente-servidor en sistemas distribuidos usando Go. Incluye un servidor que recibe y almacena mensajes, y un cliente que envía mensajes y consulta el estado.

## Estructura

```
proyecto2-client-server/
├── server/
│   └── main.go    # Servidor HTTP que gestiona mensajes
├── client/
│   └── main.go    # Cliente que envía mensajes al servidor
└── README.md
```

## Características del Servidor

- Recibe mensajes vía POST en formato JSON
- Almacena mensajes en memoria con timestamps
- Proporciona API para consultar todos los mensajes
- Endpoint de estado del servidor
- Thread-safe con mutex para concurrencia

## Características del Cliente

- Envía mensajes al servidor
- Consulta todos los mensajes almacenados
- Verifica estado del servidor
- Puede conectarse a servidor local o remoto (ngrok)

## Endpoints del Servidor

- `POST /api/mensaje` - Enviar un nuevo mensaje
- `GET /api/mensajes` - Obtener todos los mensajes
- `GET /api/status` - Estado del servidor

## Cómo ejecutar

### 1. Iniciar el Servidor

En una terminal:

```bash
cd proyecto2-client-server/server
go mod init server
go run main.go
```

El servidor estará escuchando en `http://localhost:8081`

### 2. Ejecutar el Cliente (Local)

En otra terminal:

```bash
cd proyecto2-client-server/client
go mod init client
go run main.go
```

El cliente se conectará a `http://localhost:8081` por defecto.

### 3. Ejecutar el Cliente (con ngrok)

Si quieres probar la comunicación distribuida real:

1. En una tercera terminal, expone el servidor con ngrok:
```bash
ngrok http 8081
```

2. Copia la URL de ngrok (ej: `https://xxxx.ngrok.io`)

3. Ejecuta el cliente con la URL de ngrok:
```bash
cd proyecto2-client-server/client
go run main.go https://xxxx.ngrok.io
```

## Probar con curl

```bash
# Enviar un mensaje
curl -X POST http://localhost:8081/api/mensaje \
  -H "Content-Type: application/json" \
  -d '{"contenido":"Hola servidor","cliente":"curl-client"}'

# Obtener todos los mensajes
curl http://localhost:8081/api/mensajes

# Ver estado del servidor
curl http://localhost:8081/api/status
```

## Conceptos de Sistemas Distribuidos

- **Cliente-Servidor**: Arquitectura fundamental en sistemas distribuidos
- **API RESTful**: Comunicación estándar entre componentes distribuidos
- **JSON**: Formato de intercambio de datos
- **Concurrencia**: Manejo seguro de múltiples clientes simultáneos
- **Comunicación remota**: Uso de ngrok para simular distribución real
- **Estado compartido**: Gestión de datos entre múltiples clientes
