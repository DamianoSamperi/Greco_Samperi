package magazzino

import (
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/go-sql-driver/mysql"
)

type Prodotto struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Quantita int    `json:"quantita"`
}

type Hub struct {
	ID       int        `json:"id"`
	Sede     string     `json:"sede"`
	Prodotti []Prodotto `json:"prodotti"`
}

func insert_prodotto() {
	db, err := sql.Open("mysql", "Greco_Samperi:apl@/Magazzino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inserisci un nuovo prodotto in un hub
	_, err = db.Exec("INSERT INTO Prodotti (nome, quantita, hub_id) VALUES (?, ?, ?)", "NomeProdotto", 10, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func ottieni_prodotti() {
	db, err := sql.Open("mysql", "Greco_Samperi:apl@/Magazzino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ottieni tutti i prodotti di un dato hub
	rows, err := db.Query("SELECT id, nome FROM Hub WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var h Hub
		err := rows.Scan(&h.ID, &h.Sede)
		if err != nil {
			log.Fatal(err)
		}

		prodottiRows, err := db.Query("SELECT id, nome, quantita FROM Prodotti WHERE hub_id = ?", h.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer prodottiRows.Close()

		for prodottiRows.Next() {
			var p Prodotto
			err := prodottiRows.Scan(&p.ID, &p.Nome, &p.Quantita)
			if err != nil {
				log.Fatal(err)
			}
			h.Prodotti = append(h.Prodotti, p)
		}

		if err = prodottiRows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(h)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
