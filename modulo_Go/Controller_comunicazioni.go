package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modulo_Go/spedizione"
	"net/http"
)

func Inserimento_spedizione(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
			return
		}
		var dati spedizione.Spedizione
		_ = json.Unmarshal(body, &dati)
		spedizione.Insert_Spedizione(dati.Mittente, dati.Destinatario, dati.Indirizzo, dati.Pacchi)
	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}

func Visualizza_spedizioni(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati string
	_ = json.Unmarshal(body, &dati)
	fmt.Fprint(w, spedizione.Visualizza_Spedizioni(dati))
}
func Inserimento_prodotto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hai richiesto: %s!", r.URL.Path[1:])
}

func Ottieni_prodotti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Questo è un altro handler!")
}

func main() {
	http.HandleFunc("/Inserisci_Spedizione", Inserimento_spedizione)
	http.HandleFunc("/Visualizza_Spedizioni", Visualizza_spedizioni)
	http.HandleFunc("/Ottieni_Prodotti_Hub", Ottieni_prodotti)
	http.HandleFunc("/Inserisci_Prodotto_Hub", Inserimento_prodotto)
	// TO_DO inserire error handler nel listen
	http.ListenAndServe(":8080", nil)
}
