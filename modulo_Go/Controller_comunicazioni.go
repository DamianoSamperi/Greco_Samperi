package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"modulo_Go/magazzino"
	"modulo_Go/spedizione"
	"net/http"
)

type richiesta struct {
	Stringa string           `json:"stringa"`
	Pacco   spedizione.Pacco `json:"pacco"`
}

func Inserimento_spedizione(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
			return
		}
		var dati spedizione.Spedizione
		_ = json.Unmarshal(body, &dati)
		Sede := g.Ritorna_hub_per_vicinanza(dati.Mittente)
		spedizione.Insert_Spedizione(dati.Mittente, dati.Destinatario, dati.Pacchi, Sede)
	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}

func Visualizza_spedizioni(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati string
	_ = json.Unmarshal(body, &dati)
	fmt.Fprint(w, spedizione.Visualizza_Spedizioni(dati))
}
func Inserimento_prodotto(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati richiesta
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	g.InserisciPaccoInSede(dati.Stringa, dati.Pacco)
}

func Ottieni_prodotti(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati string
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, g.OttieniPacchiPerSede(dati))
}

// func Ritorna_sede(w http.ResponseWriter, r *http.Request) string{
// 	ctx := context.TODO()
// 	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb://localhost:27017")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
// 		return
// 	}
// 	var dati string
// 	err = json.Unmarshal(body, &dati)
// 	if err != nil {
// 		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
// 		return
// 	}
// 	return g.Ritorna_hub_per_vicinanza(dati)
// }

func main() {
	http.HandleFunc("/Inserisci_Spedizione", Inserimento_spedizione)
	http.HandleFunc("/Visualizza_Spedizioni", Visualizza_spedizioni)
	//passare la sede
	http.HandleFunc("/Ottieni_Prodotti_Hub", Ottieni_prodotti)
	http.HandleFunc("/Inserisci_Prodotto_Hub", Inserimento_prodotto)
	//TO_DO funzione che passa tutti gli id delle spedizioni
	//TO_DO funzione che modifica lo stato della spedizione
	// http.HandleFunc("/Ritorna_Sede", Ritorna_sede)
	// TO_DO inserire error handler nel listen
	http.ListenAndServe(":8080", nil)

}
