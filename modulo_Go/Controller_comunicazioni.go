package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"modulo_Go/consegne"
	"modulo_Go/magazzino"
	"modulo_Go/spedizione"
	"net/http"
	"os"
)

type richiesta struct {
	Sede  string           `json:"sede"`
	Pacco spedizione.Pacco `json:"pacco"`
}
type modifica_stato struct {
	Stato         string `json:"stato"`
	Id_spedizione string `json:"id"`
}
type richiesta_spedizione struct {
	Spedizione spedizione.Spedizione `json:"spedizione"`
	Sede       string                `json:"sede"`
}

func Inserimento_spedizione(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
			return
		}
		var dati richiesta_spedizione
		_ = json.Unmarshal(body, &dati)
		// Sede := g.Ritorna_hub_per_vicinanza(dati.Mittente)
		g.Insert_Spedizione(dati.Spedizione.ID, dati.Spedizione.Mittente, dati.Spedizione.Destinatario, dati.Spedizione.Pacchi, dati.Sede)
	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}

func Visualizza_spedizioni(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati = struct {
		Mittente string
	}{}
	_ = json.Unmarshal(body, &dati)
	fmt.Fprint(w, g.Visualizza_Spedizioni(dati.Mittente))
}
func Inserimento_prodotto(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
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
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati = struct {
		Sede string
	}{}
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, g.OttieniPacchiPerSede(dati.Sede))
}

func Ritorna_sede(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati = struct {
		Indirizzo string
	}{}
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, g.Ritorna_hub_per_vicinanza(dati.Indirizzo))
}
func Ritorna_id(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, g.RitornaID())
}
func Modifica_stato(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
			return
		}
		var dati modifica_stato
		_ = json.Unmarshal(body, &dati)
		g.Modifica_Stato_Spedizione(dati.Id_spedizione, dati.Stato)

	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}
func Ottieni_percorso(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	s, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
			return
		}
		var dati = struct {
			Sede string
		}{}
		_ = json.Unmarshal(body, &dati)
		//TO_DO funzione che torna gli id delle spedizioni di un magazzino passata la sede
		ids := g.Ottieni_Spedizioni_PerSede(dati.Sede)
		//TO_DO funzione in spedizione e trovare un modo per passargli la sede come punto geografico o lo calcolo all'interno mi sa meglio
		var spedizioni []spedizione.Spedizione
		for _, id := range ids {
			//TO_DO andrebbero tornate solo quelle che sono in preparazione
			spedizione := s.Trova_spedizioni_per_ID(id)
			spedizioni = append(spedizioni, spedizione)
		}
		Coordinata_Sede := g.Ritorna_Coordinate_hub(dati.Sede)
		if Coordinata_Sede == (magazzino.Coordinate{}) {
			http.Error(w, "Sede non valida", http.StatusMethodNotAllowed)
		}
		var Sede = consegne.Punto_geografico{Indirizzo: dati.Sede, Latitudine: Coordinata_Sede.Latitudine, Longitudine: Coordinata_Sede.Longitudine}
		percorso := consegne.Trova_percorso(spedizioni, Sede)
		var indirizzi []string
		for _, p := range percorso {
			indirizzi = append(indirizzi, p.Indirizzo)
			if p.Indirizzo != dati.Sede {
				s.Modifica_Stato_Spedizione(p.Indirizzo, "InTransito")
			}
		}
		fmt.Fprint(w, indirizzi)
	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}
func main() {
	//passi una spedizione e la sede che puoi farti tornare da ritorna sede(o faccio io vedi tu) e la inserisce nel database
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
	//funziona che torna il percorso o i pacchi vedo per il corriere, bisogna passargli la sede
	http.HandleFunc("/Ottieni_Percorso", Ottieni_percorso)
	// TO_DO inserire error handler nel listen
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Print("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
