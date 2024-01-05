package magazzino

import (
	"context"
	"database/sql"
	"log"
	"modulo_Go/spedizione"
	"strconv"

	_ "go.mongodb.org/mongo-driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// _ "github.com/go-sql-driver/mysql"
)

// type Prodotto struct {
// 	ID       int    `json:"id"`
// 	Nome     string `json:"nome"`
// 	Quantita int    `json:"quantita"`
// }

type Hub struct {
	ID   int    `json:"id"`
	Sede string `json:"sede"`
	// Prodotti []Prodotto `json:"prodotti"`
	Pacchi []spedizione.Pacco
}

// la funzione modifica l'hub di un Pacco,quindi va chiamata quando
func insert_prodotto(Sede string, Pacco int) {
	db, err := sql.Open("mysql", "Greco_Samperi:apl@/Magazzino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inserisci un nuovo prodotto in un hub
	stmt, err := db.Prepare("UPDATE Prodotti SET sede = ? WHERE id = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Esegui l'istruzione SQL con i valori forniti
	_, err = stmt.Exec(Sede, Pacco)
	if err != nil {
		log.Fatal(err)
	}

}

// func ottieni_prodotti(Sede string) string {
// 	db, err := sql.Open("mysql", "Greco_Samperi:apl@/Magazzino")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Ottieni tutti i pacchi di un dato hub
// 	rows, err := db.Query("SELECT id, nome FROM Hub WHERE id = ?", 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var h Hub
// 		err := rows.Scan(&h.ID, &h.Sede)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		pacchiRows, err := db.Query("SELECT id, nome, quantita FROM Prodotti WHERE hub_id = ?", h.ID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer prodottiRows.Close()

// 		for prodottiRows.Next() {
// 			var p Prodotto
// 			err := prodottiRows.Scan(&p.ID, &p.Nome, &p.Quantita)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			h.Prodotti = append(h.Prodotti, p)
// 		}

// 		if err = prodottiRows.Err(); err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(h)
// 	}
// 	return ToString(h)
// }

func ritorna_hub_per_vicinanza(indirizzo string) string {
	//TO_DO funzione che ritorna l'hub più vicino all'indirizzo dato
	return "da implementare"
}

func OttieniPacchiPerSede(ctx context.Context, client *mongo.Client, sede string) (string, error) {
	collection := client.Database("Magazzino").Collection(sede)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return "", err
	}
	defer cursor.Close(ctx)

	var pacchi []spedizione.Pacco
	if err = cursor.All(ctx, &pacchi); err != nil {
		return "", err
	}

	return ToString(pacchi, sede), nil
}

func ToString(Pacchi []spedizione.Pacco, Sede string) string {
	String := "Hub sede: " + Sede + " prodotti:\n"
	for _, pacco := range Pacchi {
		Pacco := "Spedizione id" + strconv.Itoa(pacco.Spedizione_id) + "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		String = String + Pacco
	}
	return String
}
