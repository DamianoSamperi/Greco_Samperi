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
type modifica_data struct {
	Id_spedizione string `json:"id"`
	Data          string `json:"data"`
}
type richiesta_spedizione struct {
	ID           string `json:"id"`
	Mittente     string `json:"mittente"`
	Destinatario string `json:"destinatario"`
	Sede         string `json:"sede"`
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
		g.Insert_Spedizione(dati.ID, dati.Mittente, dati.Destinatario, dati.Sede)
	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}
func Inserimento_pacco(w http.ResponseWriter, r *http.Request) {
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
		var dati = struct {
			Spedizione_id string  `json:"id_spedizione"`
			Peso          float64 `json:"peso"`
			Dimensione    string  `json:"dimensione"`
			Prezzo        float64 `json:"prezzo"`
		}{}
		err = json.Unmarshal(body, &dati)
		if err != nil {
			http.Error(w, "Errore formato inviato", http.StatusBadRequest)
		}
		err = g.Insert_Pacco_spedizione(dati.Spedizione_id, dati.Peso, dati.Dimensione, dati.Prezzo)
		if err != nil {
			http.Error(w, "Errore inserimento "+err.Error(), http.StatusBadRequest)
			//TO_DO andrebbe fatto il rollback
		} else {
			fmt.Fprint(w, "Pacco inserito con successo")
		}
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
		Mittente string `json:"mittente"`
	}{}
	_ = json.Unmarshal(body, &dati)
	fmt.Fprint(w, g.Visualizza_Spedizioni(dati.Mittente))
}
func Traccia_spedizione(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal("Errore connessione spedizioni ", err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati = struct {
		Id_Spedizione string `json:"id_spedizione"`
	}{}
	_ = json.Unmarshal(body, &dati)
	fmt.Fprint(w, g.Traccia_Spedizione(dati.Id_Spedizione))
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
	result := g.InserisciPaccoInSede(dati.Sede, dati.Pacco)
	if result != "Inserimento completato" {
		http.Error(w, result, http.StatusBadRequest)
	} else {
		fmt.Fprint(w, result)
	}

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
		Sede string `json:"sede"`
	}{}
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, g.OttieniPacchiPerSede(dati.Sede))
}
func Consegna_hub(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	s, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati = struct {
		Nuovo_Hub     string `json:"nuovo_hub"`
		Vecchio_Hub   string `json:"vecchio_hub"`
		Id_Spedizione string `json:"id_Spedizione"`
	}{}
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	// spedizione := s.Trova_spedizioni_per_ID(dati.Id_Spedizione)
	// pacchi := spedizione.Pacchi
	err = g.SpostaPacco(dati.Id_Spedizione, dati.Vecchio_Hub, dati.Nuovo_Hub)
	if err != nil {
		if err.Error() == "vecchia sede inesistente" || err.Error() == "nuova sede inesistente" {
			fmt.Fprint(w, err.Error())
			return
		}
		http.Error(w, "Errore spostamento pacco", http.StatusBadRequest)
		return
	}
	// s.Modifica_Stato_Spedizione(dati.Id_Spedizione, "Consegnato all'Hub")
	fmt.Fprint(w, s.Modifica_Stato_Spedizione(dati.Id_Spedizione, "Consegnato all'Hub"))
	// for _,_ = range pacchi{
	// 	g.SpostaPacco(dati.Id_Spedizione,dati.vecchio_Hub,dati.nuovo_Hub)
	// }
}
func Ritorna_sede(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		http.Error(w, "Errore connessione al database", http.StatusBadRequest)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della richiesta", http.StatusBadRequest)
		return
	}
	var dati = struct {
		Indirizzo string `json:"indirizzo"`
	}{}
	err = json.Unmarshal(body, &dati)
	if err != nil {
		http.Error(w, "Formato json non corretto", http.StatusBadRequest)
		return
	}
	print("ind ", dati.Indirizzo)
	hub := g.Ritorna_hub_per_vicinanza(dati.Indirizzo)
	if hub != "" {
		fmt.Fprint(w, hub)
	} else {
		http.Error(w, "Errore nella ricerca della sede", http.StatusBadRequest)
	}
}
func Ritorna_id(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, g.RitornaID())
}
func Ottieni_data(w http.ResponseWriter, r *http.Request) {
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
	var id_spedizione string
	_ = json.Unmarshal(body, &id_spedizione)
	fmt.Fprint(w, g.Ritorna_Data_Spedizione(id_spedizione))
}
func Modifica_stato(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	m, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
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
		err = json.Unmarshal(body, &dati)
		if err != nil {
			http.Error(w, "Formato json non corretto", http.StatusBadRequest)
			return
		}
		result := g.Modifica_Stato_Spedizione(dati.Id_spedizione, dati.Stato)
		if result != "codice spedizione non valido" && result != "Il pacco è già Consegnato" && dati.Stato == "Consegnato" {
			err := m.Delete_pacchi(dati.Id_spedizione)
			if err != nil {
				http.Error(w, "Errore eliminazione Pacco", http.StatusBadRequest)
				return
			}
		}
		fmt.Fprint(w, result)
	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}
func Inserimento_data_consegna(w http.ResponseWriter, r *http.Request) {
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
		var dati modifica_data
		err = json.Unmarshal(body, &dati)
		if err != nil {
			http.Error(w, "Formato json non corretto", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, g.Modifica_Data_Consegna_Spedizione(dati.Id_spedizione, dati.Data))

	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}
func Ottieni_percorso(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	g, err := magazzino.NuovoGestoreMagazzino(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal("Errore connessione magazzino ", err)
	}
	s, err := spedizione.NuovoGestoreSpedizioni(ctx, "mongodb+srv://root:yWP2DlLumOz07vNv@apl.yignw97.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal("Errore connessione spedizioni ", err)
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
			if spedizione.ID != "nulla" {
				spedizioni = append(spedizioni, spedizione)
			}
		}
		if len(spedizioni) > 0 {
			Coordinata_Sede := g.Ritorna_Coordinate_hub(dati.Sede)
			if Coordinata_Sede == (magazzino.Coordinate{}) {
				http.Error(w, "Sede non valida", http.StatusMethodNotAllowed)
			}
			var Sede = consegne.Punto_percorso{Indirizzo: dati.Sede, Latitudine: Coordinata_Sede.Latitudine, Longitudine: Coordinata_Sede.Longitudine}
			Lista_coordinate_magazzini, sedi_magazzini := g.Ottieni_Sedi(Sede.Indirizzo)
			var lista_magazzini []consegne.Punto_percorso
			for i, coordinate := range Lista_coordinate_magazzini {
				magazzino := consegne.Punto_percorso{Indirizzo: sedi_magazzini[i], Latitudine: coordinate.Latitudine, Longitudine: coordinate.Longitudine}
				lista_magazzini = append(lista_magazzini, magazzino)
			}
			percorso := consegne.Trova_percorso(spedizioni, Sede, lista_magazzini)
			var indirizzi []string
			for _, p := range percorso {
				if p.Indirizzo == "" {
					destinatario := s.Ritorna_Destinatario_Spedizione(p.Id)
					indirizzi = append(indirizzi, p.Id+" "+destinatario+"\n")
					s.Modifica_Stato_Spedizione(p.Id, "InTransito")
				} else {
					indirizzi = append(indirizzi, p.Id+" da consegnare all'hub di "+p.Indirizzo+"\n")
					s.Modifica_Stato_Spedizione(p.Id, "InTransito")
				}
			}
			if len(indirizzi) > 0 {
				fmt.Fprint(w, indirizzi)
			} else {
				fmt.Fprint(w, "Nessuna spedizione da consegnare")
			}
		} else {
			fmt.Fprint(w, "Nessuna spedizione da consegnare")
		}

	} else {
		http.Error(w, "Metodo non valido", http.StatusMethodNotAllowed)
	}
}

func main() {
	//passi una spedizione e la sede che puoi farti tornare da ritorna sede(o faccio io vedi tu) e la inserisce nel database
	http.HandleFunc("/Inserisci_Spedizione", Inserimento_spedizione)
	//funzione dove passi i pacchi e li inserisce in una spedizione creata in caso di errore non inserire più pacchi ma riprova ad inserire la spedizione
	http.HandleFunc("/Inserisci_Pacco_spedizione", Inserimento_pacco)
	//passi un mittente e ti torna tutte le sue spedizioni
	http.HandleFunc("/Visualizza_Spedizioni", Visualizza_spedizioni)
	//passi la sede e ti ritorna tutti i pacchi in quella sede
	http.HandleFunc("/Ottieni_Prodotti_Hub", Ottieni_prodotti)
	//Passi la sede e un pacco e lo inserisce nel Hub corrispondente
	http.HandleFunc("/Inserisci_Prodotto_Hub", Inserimento_prodotto)
	//TO_DO funzione che passa tutti gli id delle spedizioni
	http.HandleFunc("/RitornaId_Spedizionie", Ritorna_id)
	//TO_DO funzione che modifica lo stato della spedizione, devi passare il nuovo stato e l'id
	http.HandleFunc("/Modifica_Stato_Spedizione", Modifica_stato)
	//passi l'indirizzo e ti torna la sede più vicina, il formato indirizzo deve essere Via, Città Provincia ex. "Via Cristoforo Colombo, Roma RM"
	http.HandleFunc("/Ritorna_Sede", Ritorna_sede)
	//funziona che torna il percorso o i pacchi vedo per il corriere, bisogna passargli la sede
	http.HandleFunc("/Ottieni_Percorso", Ottieni_percorso)
	//funzione che mi ritorna data spedizione, in modo da fare il controllo per l'aggiunta nuova data su c#
	http.HandleFunc("/Ottieni_data_spedizione", Ottieni_data)
	//funzione che permette all'utente di scegliere una data di consegna, di default è impostata come il prima possibile
	http.HandleFunc("/Scegli_data_consegna", Inserimento_data_consegna)
	//TO_DO mi serve una funzione dove passi id_spedizione e cambia hub a tutti i pacchi presenti, viene passato il nuovo hub,vecchio hub, id spedizione
	http.HandleFunc("/Consegna_hub", Consegna_hub)
	//Visualizza TRacciamento dato id spedizione
	http.HandleFunc("/tracciamento_spedizione", Traccia_spedizione)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Print("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
