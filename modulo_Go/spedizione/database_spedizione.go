package spedizione

import (
	//TO_DO da passare come non relazionale, per aggiungere lista stati per il tracciamento
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// _ "github.com/go-sql-driver/mysql"
)

type Stato int

const (
	InPreparazione Stato = iota
	InTransito
	Hub
	Consegnato
	Errore
)

type Spedizione struct {
	ID           string  `bson:"id"`
	Mittente     string  `bson:"mittente"`
	Destinatario string  `bson:"destinatario"`
	Stato        []Stato `bson:"stato"`
	NumeroPacchi int     `bson:"numero_pacchi"`
	Pacchi       []Pacco `bson:"pacchi"`
}

type GestoreSpedizioni struct {
	client *mongo.Client
	ctx    context.Context
}

type Pacco struct {
	ID            int
	Spedizione_id int
	// Nome string  `json:"nome"`
	Peso       float64
	Lunghezza  float64
	Altezza    float64
	Profondità float64
	Prezzo     float64
}

func (s Stato) String() string {
	switch s {
	case InPreparazione:
		return "InPreparazione"
	case InTransito:
		return "InTransito"
	case Hub:
		return "Consegnato all'Hub"
	case Consegnato:
		return "Consegnato"
	}
	return "unknown"
}
func ToStato(s string) Stato {
	switch s {
	case "InPreparazione":
		return InPreparazione
	case "InTransito":
		return InTransito
	case "Consegnato all'Hub":
		return Hub
	case "Consegnato":
		return Consegnato
	}
	return Errore
}
func NuovoGestoreSpedizioni(ctx context.Context, uri string) (*GestoreSpedizioni, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &GestoreSpedizioni{client: client, ctx: ctx}, nil
}

func (g *GestoreSpedizioni) Visualizza_Spedizioni(Mittente string) string {
	collection := g.client.Database("Magazzino").Collection("spedizioni")
	var spedizioni []Spedizione
	filter := bson.D{{Key: "mittente", Value: Mittente}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result Spedizione
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		spedizioni = append(spedizioni, result)
	}

	return ToString(spedizioni)
}

func (g *GestoreSpedizioni) Insert_Spedizione(ID string, mittente string, destinatario string, Pacchi []Pacco, sede string) {
	collection := g.client.Database("Magazzino").Collection("spedizioni")
	var Stati []Stato
	Stati = append(Stati, InPreparazione)
	spedizione := Spedizione{ID, mittente, destinatario, Stati, len(Pacchi), Pacchi}
	insertResult, err := collection.InsertOne(context.TODO(), spedizione)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserita una nuova spedizione con ID:", insertResult.InsertedID)
	//TO_DO potrebbbe essere necessario inserire il pacco nel magazzino
}

func (g *GestoreSpedizioni) RitornaID() []string {
	collection := g.client.Database("Magazzino").Collection("spedizioni")
	opts := options.Find().SetProjection(bson.D{{Key: "ID", Value: 1}}) //TO_DO il key-value viene aggiunto da Golang perchè richiede i key name della struct, andrebbe controllato se funziona anche cosi l'option
	cur, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var IDs []string
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result string
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		IDs = append(IDs, result)
	}
	return IDs
}
func (g *GestoreSpedizioni) Modifica_Stato_Spedizione(id string, stato string) {
	//TO_DO funzione modifica, però prima va cambiato il database in non relazionale
	collection := g.client.Database("Magazzino").Collection("spedizioni")
	filter := bson.D{{Key: "idspedizione", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "stato", Value: ToStato(stato)}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Modificati %v documenti\n", updateResult.ModifiedCount)
}

func ToString(spedizioni []Spedizione) string {
	var String string
	for _, s := range spedizioni {
		SpedizioneString := "Id " + s.ID + " Mittente" + s.Mittente + "Destinatario " + s.Destinatario + " Stato " + s.Stato[len(s.Stato)-1].String() + " Numero Pacchi: " + strconv.Itoa(s.NumeroPacchi)
		for _, pacco := range s.Pacchi {
			Pacco := "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
			SpedizioneString = SpedizioneString + Pacco
		}
		String = String + SpedizioneString
	}
	return String
}
