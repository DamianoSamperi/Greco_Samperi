package magazzino

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math"
	"modulo_Go/spedizione"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
type GestoreMagazzino struct {
	client *mongo.Client
	ctx    context.Context
}

//	type RispostaAPI struct {
//		Latitudine  float64 `json:"latitude"`
//		Longitudine float64 `json:"longitude"`
//	}
type Coordinate struct {
	Latitudine  float64 `bson:"latitude"`
	Longitudine float64 `bson:"longitude"`
}

func contiene(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
func NuovoGestoreMagazzino(ctx context.Context, uri string) (*GestoreMagazzino, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &GestoreMagazzino{client: client, ctx: ctx}, nil
}

func (g *GestoreMagazzino) Ritorna_hub_per_vicinanza(indirizzo string) string {
	//TO_DO funzione che ritorna l'hub più vicino all'indirizzo dato
	//TO_DO trasforma indirizzi in coordinate e poi calcola distanza tra due punti e moltiplica per indice curvatura terreste poi trovi il minimo delle distanze
	R := 6372795.477598
	url := "https://geocoding.openapi.it/geocode"
	collezioni, _ := g.client.Database("APL").ListCollectionNames(g.ctx, bson.M{})
	payload := strings.NewReader("{\"address\":" + indirizzo + "}")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer 659ad5656af8cf61ad062a3c")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "errore nell'indirizzo " + err.Error()
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var risposta Coordinate
	err = json.Unmarshal(body, &risposta)
	if err != nil {
		log.Fatal(err)
	}
	latA := risposta.Latitudine
	lonA := risposta.Longitudine
	min := math.MaxFloat64
	var sede string
	for _, collezione := range collezioni {
		if collezione != "spedizioni" {
			collection := g.client.Database("APL").Collection(collezione)
			var result Coordinate
			// filtro := bson.M{"latitude": bson.M{"$exists": true}}
			err := collection.FindOne(g.ctx, bson.M{}).Decode(&result)
			if err != nil {
				return "errore nella richesta al database " + err.Error()
			}
			latB := result.Latitudine
			lonB := result.Longitudine
			distanza := R * math.Acos(math.Sin(latA)*math.Sin(latB)+math.Cos(latA)*math.Cos(latB)*math.Cos(lonA-lonB))
			if min > distanza {
				min = distanza
				sede = collezione
			}
		}
	}
	return sede
}
func (g *GestoreMagazzino) Ritorna_Coordinate_hub(sede string) Coordinate {
	collection := g.client.Database("APL").Collection(sede)
	var result Coordinate
	err := collection.FindOne(g.ctx, bson.D{}).Decode(&result)
	if err != nil {
		return Coordinate{}
	}
	return result
}
func (g *GestoreMagazzino) Ottieni_Sedi(sede string) ([]Coordinate, []string) {
	collezioni, _ := g.client.Database("APL").ListCollectionNames(g.ctx, bson.M{})
	var lista_magazzini []Coordinate
	var sedi_magazzini []string
	for _, collezione := range collezioni {
		if collezione != sede && collezione != "spedizioni" {
			collection := g.client.Database("APL").Collection(collezione)
			var result Coordinate
			err := collection.FindOne(g.ctx, bson.M{}).Decode(&result)
			if err != nil {
				return nil, nil
			}
			lista_magazzini = append(lista_magazzini, result)
			sedi_magazzini = append(sedi_magazzini, collezione)
		}
	}

	return lista_magazzini, sedi_magazzini
}

func (g *GestoreMagazzino) OttieniPacchiPerSede(sede string) []spedizione.Pacco {
	collection := g.client.Database("APL").Collection(sede)
	filtro := bson.M{"spedizione_id": bson.M{"$exists": true}}
	cursor, err := collection.Find(g.ctx, filtro)
	if err != nil {
		print("error ", err)
		return []spedizione.Pacco{}
	}
	var pacchi []spedizione.Pacco
	err = cursor.All(g.ctx, &pacchi)
	if err != nil {
		print("error ", err)
		return []spedizione.Pacco{}
	}
	defer cursor.Close(g.ctx)
	return pacchi
	// return ToString(pacchi, sede)
}

func (g *GestoreMagazzino) Ottieni_Spedizioni_PerSede(sede string) []string {
	Pacchi := g.OttieniPacchiPerSede(sede)
	ids := make(map[string]bool)
	for _, pacco := range Pacchi {
		ids[pacco.Spedizione_id] = true
	}
	uniqueIDs := make([]string, 0, len(ids))
	for id := range ids {
		uniqueIDs = append(uniqueIDs, id)
	}

	return uniqueIDs
}

func (g *GestoreMagazzino) InserisciPaccoInSede(sede string, p spedizione.Pacco) string {
	collezioni, _ := g.client.Database("APL").ListCollectionNames(g.ctx, bson.M{})
	if contiene(collezioni, sede) {
		collection := g.client.Database("APL").Collection(sede)
		_, err := collection.InsertOne(g.ctx, p)
		if err != nil {
			return err.Error()
		} else {
			return "Inserimento completato"
		}
	} else {
		return "Sede non esistente"
	}

}

func (g *GestoreMagazzino) SpostaPacco(id string, vecchiaSede string, nuovaSede string) error {
	vecchiaCollection := g.client.Database("APL").Collection(vecchiaSede)
	nuovaCollection := g.client.Database("APL").Collection(nuovaSede)

	var p spedizione.Pacco
	err := vecchiaCollection.FindOneAndDelete(g.ctx, bson.M{"_id": id}).Decode(&p)
	if err != nil {
		return err
	}

	_, err = nuovaCollection.InsertOne(g.ctx, p)
	return err
}

func ToString(Pacchi []spedizione.Pacco, Sede string) string {
	String := "Hub sede: " + Sede + " prodotti:\n"
	for _, pacco := range Pacchi {
		// Pacco := "Spedizione id" + pacco.Spedizione_id + "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		Pacco := "Spedizione id" + pacco.Spedizione_id + "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Dimensione" + pacco.Dimensione + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		String = String + Pacco
	}
	return String
}
