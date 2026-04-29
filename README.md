# ⚡ Pokémon Battle Strategy Engine (PBSE)

Un motor de simulación de alta performance desarrollado en **Go**, diseñado para resolver árboles de decisión complejos mediante **Backtracking** optimizado con **Pila Manual** y **Memoización Distribuida**.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Redis](https://img.shields.io/badge/Cache-Redis-D82C20?style=flat&logo=redis)](https://redis.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 🚀 Conceptos Arquitectónicos

Este proyecto demuestra patrones de nivel Senior para problemas de búsqueda exhaustiva:

* **Polimorfismo Dinámico:** Uso de interfaces en Go para inyectar diferentes comportamientos de ataque sin acoplar el motor.
* **Manual Stack Backtracking:** Evita el *Stack Overflow* inherente a la recursividad profunda mediante el uso de una pila en el Heap.
* **Memoización con Redis:** Implementación de una capa de caché persistente para estados de batalla idénticos, reduciendo la complejidad temporal de $O(b^d)$ a $O(1)$ en consultas repetidas.

---

## 📁 Estructura del Proyecto

```text
.
├── cmd/simulator/      # Punto de entrada de la aplicación
├── internal/
│   ├── pokemon/        # Lógica de dominio y polimorfismo (Interfaces)
│   └── cache/          # Wrapper para conexión y lógica de Redis
├── docker-compose.yml  # Infraestructura local (Redis)
├── go.mod              # Definición de módulos
└── README.md           # Documentación técnica

🛠️ Instalación y Configuración
1. Requisitos Previos
Go 1.21+

Docker & Docker Compose

2. Levantar Infraestructura
Inicia la instancia de Redis en segundo plano:

docker-compose up -d

3. Ejecutar el Simulador

# Descargar dependencias
go mod tidy

# Correr la simulación
go run cmd/simulator/main.go

🧪 Ejemplo de Funcionamiento
El motor evalúa el mejor camino de victoria. La primera ejecución calcula el árbol; la segunda recupera el estado instantáneamente:

Estado Inicial: Pikachu (Sparky) vs Enemigo (100 HP)

Proceso: El motor genera una firma única del estado (sim_pika_Sparky_100).

Acción: Si la firma no existe en Redis, explora todas las combinaciones de movimientos.

Almacenamiento: El resultado óptimo se persiste con un TTL definido.

☁️ Deploy (Cloud Native)
Este sistema está diseñado para ser Stateless. Para desplegar:

Redis: Crear instancia gratuita en Upstash.

App: Desplegar en Railway o Render.

Variables de Entorno:

REDIS_URL: URL de conexión de Upstash.

PORT: Puerto de escucha (opcional).

🛠️ Roadmap
[ ] Implementar Goroutines para exploración de ramas en paralelo.

[ ] Añadir soporte para tipos elementales (Fuego, Agua, Eléctrico) con multiplicadores polimórficos.

[ ] Interfaz gRPC para comunicación de microservicios.

Desarrollado con precisión para simulación de alta concurrencia.