package magazzino

import (
	"context"
	"modulo_Go/spedizione"
	"strconv"

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

func NuovoGestoreMagazzino(ctx context.Context, uri string) (*GestoreMagazzino, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &GestoreMagazzino{client: client, ctx: ctx}, nil
}

func Ritorna_hub_per_vicinanza(indirizzo string) string {
	//TO_DO funzione che ritorna l'hub più vicino all'indirizzo dato
	return "da implementare"
}

func (g *GestoreMagazzino) OttieniPacchiPerSede(sede string) string {
	collection := g.client.Database("Magazzino").Collection(sede)
	cursor, err := collection.Find(g.ctx, bson.M{})
	if err != nil {
		return err.Error()
	}
	defer cursor.Close(g.ctx)

	var pacchi []spedizione.Pacco
	if err = cursor.All(g.ctx, &pacchi); err != nil {
		return err.Error()
	}

	return ToString(pacchi, sede)
}

func (g *GestoreMagazzino) InserisciPaccoInSede(sede string, p spedizione.Pacco) error {
	collection := g.client.Database("nomeDelTuoDatabase").Collection(sede)
	_, err := collection.InsertOne(g.ctx, p)
	return err
}

func (g *GestoreMagazzino) SpostaPacco(id string, vecchiaSede string, nuovaSede string) error {
	vecchiaCollection := g.client.Database("nomeDelTuoDatabase").Collection(vecchiaSede)
	nuovaCollection := g.client.Database("nomeDelTuoDatabase").Collection(nuovaSede)

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
		Pacco := "Spedizione id" + strconv.Itoa(pacco.Spedizione_id) + "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		String = String + Pacco
	}
	return String
}
