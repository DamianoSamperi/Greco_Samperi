package spedizione

import (
	"database/sql"
	"log"
	"strconv"
	// _ "github.com/go-sql-driver/mysql"
)

type Spedizione struct {
	ID           int    `json:"id"`
	Mittente     string `json:"mittente"`
	Destinatario string `json:"destinatario"`
	Stato        string `json:"stato"`
	NumeroPacchi int    `json:"numero_pacchi"`
	Pacchi       []Pacco
}

type Pacco struct {
	ID            int `json:"id"`
	Spedizione_id int `json:"spedizione_id"`
	// Nome string  `json:"nome"`
	Peso       float64 `json:"peso"`
	Lunghezza  float64 `json:"lunghezza"`
	Altezza    float64 `json:"altezza"`
	Profondità float64 `json:"profondità"`
	Prezzo     float64 `json:"prezzo"`
}

func Visualizza_Spedizioni(Mittente string) string {
	db, err := sql.Open("mysql", "Greco_Samperi:apl@/Spedizione")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Visualizza tutte le spedizioni
	rows, err := db.Query("SELECT id, indirizzo, numero_pacchi FROM Spedizioni WHERE mittente = ?", Mittente)
	if err != nil {
		log.Fatal(err)

	}
	defer rows.Close()

	for rows.Next() {
		var s Spedizione
		err := rows.Scan(&s.ID, &s.Mittente, &s.Destinatario, &s.NumeroPacchi)
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

		// fmt.Println(s)
		return ToString(s)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return "error"
}

func Insert_Spedizione(mittente string, destinatario string, Pacchi []Pacco, sede string) {
	db, err := sql.Open("mysql", "Greco_Samperi:apl@/Spedizione")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Inserisci una nuova spedizione
	res, err := db.Exec("INSERT INTO Spedizioni (mittente, destinatario, numero_pacchi) VALUES (?, ?, ?)", mittente, destinatario, len(Pacchi))
	if err != nil {
		log.Fatal(err)
	}
	idSpedizione, _ := res.LastInsertId()
	for _, pacco := range Pacchi {
		_, err := db.Exec("INSERT INTO Pacchi (peso,lunghezza,altezza,profondità,prezzo, spedizione_id, hub) VALUES (?, ?, ?, ?, ?, ?, ?)", pacco.Peso, pacco.Lunghezza, pacco.Altezza, pacco.Profondità, pacco.Prezzo, idSpedizione, sede)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func ToString(s Spedizione) string {
	String := "Id " + strconv.Itoa(s.ID) + " Mittente" + s.Mittente + "Destinatario " + s.Destinatario + " Stato " + s.Stato + " Numero Pacchi: " + strconv.Itoa(s.NumeroPacchi)
	for _, pacco := range s.Pacchi {
		Pacco := "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		String = String + Pacco
	}
	return String
}
