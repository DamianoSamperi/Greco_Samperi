package main

import (
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/go-sql-driver/mysql"
)

type Spedizione struct {
	ID           int    `json:"id"`
	Nome         string `json:"nome"`
	Indirizzo    string `json:"indirizzo"`
	NumeroPacchi int    `json:"numero_pacchi"`
	Pacchi       []Pacco
}

type Pacco struct {
	ID int `json:"id"`
	// Nome string  `json:"nome"`
	Peso       float64 `json:"peso"`
	Lunghezza  float64 `json:"lunghezza"`
	Altezza    float64 `json:"altezza"`
	Profondità float64 `json:"profondità"`
}

func Visualizza_Spedizioni(db *sql.DB) {
	// Visualizza tutte le spedizioni
	rows, err := db.Query("SELECT id, nome, indirizzo, numero_pacchi FROM Spedizioni")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s Spedizione
		err := rows.Scan(&s.ID, &s.Nome, &s.Indirizzo, &s.NumeroPacchi)
		if err != nil {
			log.Fatal(err)
		}

		pacchiRows, err := db.Query("SELECT id, nome, peso FROM Pacchi WHERE spedizione_id = ?", s.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer pacchiRows.Close()

		for pacchiRows.Next() {
			var p Pacco
			err := pacchiRows.Scan(&p.ID, &p.Peso, &p.Altezza, &p.Lunghezza, &p.Profondità)
			if err != nil {
				log.Fatal(err)
			}
			s.Pacchi = append(s.Pacchi, p)
		}

		if err = pacchiRows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(s)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func Insert_Spedizione(db *sql.DB) {
	// Inserisci una nuova spedizione
	_, err := db.Exec("INSERT INTO Spedizioni (nome, indirizzo, numero_pacchi) VALUES (?, ?, ?)", "NomeSpedizione", "IndirizzoSpedizione", 10)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
