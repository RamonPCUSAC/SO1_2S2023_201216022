package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MisDatos struct {
	Carnet string `json:"id"`
	Nombre string `json:"name"`
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	student := MisDatos{
		Carnet: "201216022",
		Nombre: "Ramon Osvaldo Patzan Caballeros",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func main() {
	http.HandleFunc("/data", getDataHandler)
	port := "8080"
	fmt.Printf("Servidor en puerto %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
