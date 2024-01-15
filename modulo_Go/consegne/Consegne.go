package consegne

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"modulo_Go/spedizione"
	"net/http"
	"strings"
	"time"
)

type Punto_geografico struct {
	Indirizzo        string    `json:"indirizzo"`
	Latitudine       float64   `json:"latitude"`
	Longitudine      float64   `json:"longitude"`
	Consegna_Stimata time.Time `json:"consegna_Stimata"`
}
type Coordinate struct {
	Latitudine  float64 `json:"latitude"`
	Longitudine float64 `json:"longitude"`
}

//	type Direzione struct {
//		id        string
//		Distanza  float64
//		Direzione string
//	}
type Direzione struct {
	angolo_inf float64
	angolo_sup float64
}

var (
	Sud = Direzione{
		angolo_inf: 225.0,
		angolo_sup: 315.0,
	}
	Est = Direzione{
		angolo_inf: 315.0,
		angolo_sup: 45.0,
	}
	Nord = Direzione{
		angolo_inf: 45.0,
		angolo_sup: 135.0,
	}
	Ovest = Direzione{
		angolo_inf: 135.0,
		angolo_sup: 225.0,
	}
)

// distanza massima che può percorrere un corriere in una giornata in metri
const distanza_massima_percorribile = 320000.0

//	func direzione(angolo float64) string {
//		switch {
//		case angolo >= 0 && angolo < 45:
//			return "Nord"
//		case angolo >= 45 && angolo < 90:
//			return "Nord-Est"
//		case angolo >= 90 && angolo < 135:
//			return "Est"
//		case angolo >= 135 && angolo < 180:
//			return "Sud-Est"
//		case angolo >= 180 && angolo < 225:
//			return "Sud"
//		case angolo >= 225 && angolo < 270:
//			return "Sud-Ovest"
//		case angolo >= 270 && angolo < 315:
//			return "Ovest"
//		default:
//			return "Nord-Ovest"
//		}
//	}
func Todirezione(angolo float64) string {
	switch {
	case angolo >= 45 && angolo < 135:
		return "Nord"
	case (angolo >= 315 && angolo < 360) || (angolo >= 0 && angolo < 45):
		return "Est"
	case angolo >= 225 && angolo < 315:
		return "Sud"
	default:
		return "Ovest"
	}
}
func Calcola_distanza_punti(destinatario Punto_geografico, origine Punto_geografico) float64 {
	R := 6372795.477598
	latA := destinatario.Latitudine
	lonA := destinatario.Longitudine
	latB := origine.Latitudine
	lonB := origine.Longitudine
	distanza := R * math.Acos(math.Sin(latA)*math.Sin(latB)+math.Cos(latA)*math.Cos(latB)*math.Cos(lonA-lonB))
	return distanza
}
func calcola_direzione_punti(destinatario Punto_geografico, origine Punto_geografico) float64 {
	latA := origine.Latitudine
	lonA := origine.Longitudine
	latB := destinatario.Latitudine
	lonB := destinatario.Longitudine
	delta := math.Log2(math.Tan(latB/2+math.Pi/4) / math.Tan(latA/2+math.Pi/4))
	delta_lon := math.Abs(lonA - lonB)
	if delta_lon > 180 {
		delta_lon = math.Mod(delta_lon, 180.00)
	}
	angolo := math.Atan2(delta_lon, delta)
	return angolo
	// direzione := direzione(angolo)
	// return direzione
}
func calcola_punti(spedizioni []spedizione.Spedizione, sede Punto_geografico) []Punto_geografico {
	url := "https://geocoding.openapi.it/geocode"
	var punti []Punto_geografico
	punti = append(punti, sede)
	for _, spedizione := range spedizioni {
		payload := strings.NewReader("{\"address\":" + spedizione.Destinatario + "}")
		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/json")
		req.Header.Add("Authorization", "Bearer 659ad5656af8cf61ad062a3c")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		var risposta = struct {
			Success bool `json:"success"`
			Element struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"element"`
		}{}
		err = json.Unmarshal(body, &risposta)
		if err != nil {
			log.Fatal(err)
		}
		latA := risposta.Element.Latitude
		lonA := risposta.Element.Longitude
		var tupla Punto_geografico
		if spedizione.Data_consegna.IsZero() {
			tupla = Punto_geografico{Indirizzo: spedizione.ID, Latitudine: latA, Longitudine: lonA}
		} else {
			tupla = Punto_geografico{Indirizzo: spedizione.ID, Latitudine: latA, Longitudine: lonA, Consegna_Stimata: spedizione.Data_consegna}
		}
		punti = append(punti, tupla)
	}
	return punti
}

func trovaMagazzino_più_vicino(destinatario Punto_geografico, origine Punto_geografico, lista_magazzini []Punto_geografico) (float64, Punto_geografico) {
	for _, magazzino := range lista_magazzini {
		direzione := Todirezione(calcola_direzione_punti(destinatario, origine))
		direzione_magazzino := Todirezione(calcola_direzione_punti(magazzino, origine))
		if direzione == direzione_magazzino {
			distanza_magazzino := Calcola_distanza_punti(magazzino, origine)
			if distanza_magazzino <= distanza_massima_percorribile {
				return distanza_magazzino, magazzino
			}
		}
	}
	return distanza_massima_percorribile + 1, Punto_geografico{}
}
func Calcola_distanza_minima(origine Punto_geografico, Diramazioni []Punto_geografico, direzione_non_ammessa Direzione, distanza_residua_percorribile float64, lista_magazzini []Punto_geografico) (Punto_geografico, int, Direzione, float64) {
	minDistanza := math.MaxFloat64
	minDiramazione := Punto_geografico{}
	minIndice := -1
	nuovaDirezione := -1.0
	for i, p := range Diramazioni {
		direzione := calcola_direzione_punti(origine, p)
		if direzione >= direzione_non_ammessa.angolo_sup || direzione < direzione_non_ammessa.angolo_inf {
			d := Calcola_distanza_punti(origine, p)
			if d <= distanza_massima_percorribile {
				if (distanza_residua_percorribile - d) >= 0 {
					if !p.Consegna_Stimata.IsZero() && p.Consegna_Stimata == time.Now().AddDate(0, 0, 1) {
						if d < minDistanza {
							minDistanza = d
							minDiramazione = p
							minIndice = i
							nuovaDirezione = direzione
							distanza_residua_percorribile = distanza_residua_percorribile - d
						}
					}
				} else {
					d, p := trovaMagazzino_più_vicino(origine, p, lista_magazzini)
					if (distanza_residua_percorribile - d) >= 0 {
						if d < minDistanza {
							minDistanza = d
							minDiramazione = p
							minIndice = i
							nuovaDirezione = direzione
							distanza_residua_percorribile = distanza_residua_percorribile - d
						}
					}
				}
			}
		}
	}
	direzione_non_ammessa = nuovaDirezione_non_ammessa(direzione_non_ammessa, nuovaDirezione)

	return minDiramazione, minIndice, direzione_non_ammessa, distanza_residua_percorribile
}
func nuovaDirezione_non_ammessa(direzione_non_ammessa Direzione, nuovaDirezione float64) Direzione {
	if direzione_non_ammessa.angolo_inf == -1 {
		var nuovaDirezione_non_ammessa = Direzione{angolo_inf: math.Mod((nuovaDirezione + 90.0), 360), angolo_sup: math.Mod((nuovaDirezione + 270), 360)}
		return nuovaDirezione_non_ammessa
	} else {
		rotazione := nuovaDirezione - ((direzione_non_ammessa.angolo_inf+direzione_non_ammessa.angolo_sup)/2 - 180)
		var nuovaDirezione_non_ammessa = Direzione{angolo_inf: math.Mod(direzione_non_ammessa.angolo_inf+rotazione, 360), angolo_sup: math.Mod((direzione_non_ammessa.angolo_sup + rotazione), 360)}
		return nuovaDirezione_non_ammessa
	}

}

func Trova_percorso(spedizioni []spedizione.Spedizione, sede Punto_geografico, lista_magazzini []Punto_geografico) []Punto_geografico {
	punti := calcola_punti(spedizioni, sede)
	var indice int
	percorso := []Punto_geografico{}
	puntoCorrente := punti[0]               // Scelgo la sede come origine
	punti = append(punti[:0], punti[1:]...) // Rimuovo la sede dalla lista dei punti
	var direzione_non_ammessa = Direzione{angolo_inf: -1, angolo_sup: 361}
	var distanza_residua_percorribile = distanza_massima_percorribile
	for len(punti) > 0 {

		puntoCorrente, indice, direzione_non_ammessa, distanza_residua_percorribile = Calcola_distanza_minima(puntoCorrente, punti, direzione_non_ammessa, distanza_residua_percorribile, lista_magazzini)
		punti = append(punti[:indice], punti[indice+1:]...)
		if (puntoCorrente != Punto_geografico{}) {
			percorso = append(percorso, puntoCorrente)
		}
	}
	for _, punto := range percorso {
		fmt.Println(punto)
	}
	return percorso
}
