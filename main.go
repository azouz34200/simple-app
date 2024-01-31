package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.elastic.co/apm/module/apmhttp"
)

func main() {
	// Enveloppe le gestionnaire HTTP avec l'instrumentation APM
	http.Handle("/", apmhttp.Wrap(http.HandlerFunc(helloHandler)))

	log.Println("Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// helloHandler gère les requêtes vers l'endpoint /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Simule une charge de travail
	generateLoad()

	duration := time.Since(start)
	log.Printf("Durée du traitement: %s\n", duration)
	fmt.Fprint(w, "Hello World")
}

// generateLoad simule une charge de travail avec une attente et des calculs
func generateLoad() {
	// Attend de manière aléatoire jusqu'à 1 seconde pour simuler un travail I/O
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10000)))
	// Effectue des calculs pour simuler une charge CPU
	for i := 0; i < 10000000; i++ {
		_ = rand.Float64() * rand.Float64()
	}
}

func init() {
	// Initialise le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())
}
