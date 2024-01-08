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
	Sede  string           `json:"sede"`
	Pacco spedizione.Pacco `json:"pacco"`
}
type modifica_stato struct {
	Stato         string `json:"stato"`
	Id_spedizione string `json:"id"`
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
	g.InserisciPaccoInSede(dati.Sede, dati.Pacco)
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

func Ritorna_sede(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprint(w, g.Ritorna_hub_per_vicinanza(dati))
}
func Ritorna_id(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, spedizione.RitornaID())
}
func Modifica_stato(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
			return
		}
		var dati modifica_stato
		_ = json.Unmarshal(body, &dati)
		spedizione.Modifica_Stato_Spedizione(dati.Id_spedizione, dati.Stato)

	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}
func main() {
	//passi una spedizione e la inserisce nel database
	http.HandleFunc("/Inserisci_Spedizione", Inserimento_spedizione)
	//passi un mittente e ti torna tutte le sue spedizioni
	http.HandleFunc("/Visualizza_Spedizioni", Visualizza_spedizioni)
	//passi la sede e ti ritorna tutti i pacchi in quella sede
	http.HandleFunc("/Ottieni_Prodotti_Hub", Ottieni_prodotti)
	//Passi la sede e un pacco e lo inserisce nel Hub corrispondente
	http.HandleFunc("/Inserisci_Prodotto_Hub", Inserimento_prodotto)
	//TO_DO funzione che passa tutti gli id delle spedizioni
	http.HandleFunc("/RitornaId_Spedizionie", Ritorna_id)
	//TO_DO funzione che modifica lo stato della spedizione
	http.HandleFunc("/Modifica_Stato_Spedizione", Modifica_stato)
	//passi l'indirizzo e ti torna la sede più vicina
	http.HandleFunc("/Ritorna_Sede", Ritorna_sede)
	// TO_DO inserire error handler nel listen
	http.ListenAndServe(":8080", nil)

}
