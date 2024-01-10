package consegne

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"modulo_Go/spedizione"
	"net/http"
	"strings"
)

type RispostaAPI struct {
	Latitudine  float64 `json:"latitude"`
	Longitudine float64 `json:"longitude"`
}

// deve ricevere tutte le spedizioni con i pacchi in una sede, la sede dove si trova
func Calcola_consegne(spedizioni []spedizione.Spedizione, sede RispostaAPI) []spedizione.Pacco {
	//mi calcolo la distanza tra tutti i destinatari delle spedizioni dalla sede corrispondente e la direzione
	R := 6372795.477598
	url := "https://geocoding.openapi.it/geocode"

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
		var risposta RispostaAPI
		err = json.Unmarshal(body, &risposta)
		if err != nil {
			log.Fatal(err)
		}
		latA := risposta.Latitudine
		lonA := risposta.Longitudine
		latB := sede.Latitudine
		lonB := sede.Longitudine
		distanza := R * math.Acos(math.Sin(latA)*math.Sin(latB)+math.Cos(latA)*math.Cos(latB)*math.Cos(lonA-lonB))
		delta := math.Log2(math.Tan(latB/2+math.Pi/4) / math.Tan(latA/2+math.Pi/4))
		delta_lon := math.Abs(lonA - lonB)
		if delta_lon > 180 {
			delta_lon = math.Mod(delta_lon, 180.00)
		}
		direzione := math.Atan2(delta_lon, delta)
	}
}
