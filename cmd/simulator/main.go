package main

import (
	"context"
	"fmt"
	"log"
	"pokemon-engine/internal/pokemon"
	"os"
	"github.com/santiagourdaneta/Simulador-de-Pokemon-en-Golang-con-Memoizacion-Stack-Manual-Polimorfismo-y-Redis/internal/pokemon"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Estado struct {
	HPEnemigo int
	Pasos     int
}

func main() {
	// 1. Conexión a Redis  
	rdb := redis.NewClient(&redis.Options{
		Addr: "os.Getenv("https://proper-seasnail-108819.upstash.io")",  
	})

	// 2. Definir contendientes (Polimorfismo)
	pika := pokemon.Pikachu{Nombre: "Sparky"}
	hpEnemigoInicial := 100

	// 3. MEMOIZACIÓN: ¿Ya calculamos esto antes?
	cacheKey := fmt.Sprintf("sim_%s_%d", pika.GetID(), hpEnemigoInicial)
	val, err := rdb.Get(ctx, cacheKey).Result()
	
	if err == nil {
		fmt.Printf("🚀 [CACHE HIT] Resultado recuperado: %s\n", val)
		return
	}

	// 4. BACKTRACKING CON PILA MANUAL (Si no está en cache)
	fmt.Println("🧠 [CALCULANDO] Iniciando simulación pesada...")
	pila := []Estado{{HPEnemigo: hpEnemigoInicial, Pasos: 0}}
	mejorResultado := 999 // El mínimo de pasos

	for len(pila) > 0 {
		// Pop manual
		actual := pila[len(pila)-1]
		pila = pila[:len(pila)-1]

		if actual.HPEnemigo <= 0 {
			if actual.Pasos < mejorResultado {
				mejorResultado = actual.Pasos
			}
			continue
		}

		// Límite de profundidad para evitar ciclos infinitos
		if actual.Pasos < 10 {
			pila = append(pila, Estado{
				HPEnemigo: actual.HPEnemigo - pika.Atacar(),
				Pasos:     actual.Pasos + 1,
			})
		}
	}

	// 5. GUARDAR RESULTADO
	resStr := fmt.Sprintf("Victoria en %d turnos", mejorResultado)
	rdb.Set(ctx, cacheKey, resStr, 0)
	fmt.Printf("✅ [GUARDADO] %s\n", resStr)
}