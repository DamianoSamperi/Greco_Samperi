package spedizione

import (
	//TO_DO da passare come non relazionale, per aggiungere lista stati per il tracciamento
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

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
	ID              string    `bson:"id"`
	Mittente        string    `bson:"mittente"`
	Destinatario    string    `bson:"destinatario"`
	Stato           []Stato   `bson:"stato"`
	Data_spedizione time.Time `bson:"data_spedizione"`
	Data_consegna   time.Time `bson:"data_consegna"`
	NumeroPacchi    int       `bson:"numero_pacchi"`
	Pacchi          []Pacco   `bson:"pacchi"`
}

type GestoreSpedizioni struct {
	client *mongo.Client
	ctx    context.Context
}

type Pacco struct {
	Spedizione_id string
	// Nome string  `json:"nome"`
	Peso float64
	// Lunghezza  float64
	// Altezza    float64
	// Profondità float64
	Dimensione string
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
	collection := g.client.Database("APL").Collection("spedizioni")
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
func (g *GestoreSpedizioni) Traccia_Spedizione(ID string) string {
	collection := g.client.Database("APL").Collection("spedizioni")
	filter := bson.D{{Key: "id", Value: ID}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())
	var result Spedizione
	err = cur.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return Tracciamento(result)

}
func (g *GestoreSpedizioni) Trova_spedizioni_per_ID(ID string) Spedizione {
	collection := g.client.Database("APL").Collection("spedizioni")
	filter := bson.D{{Key: "id", Value: ID}}
	var result Spedizione
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal("Errore ricerca nella collezione ", err)
	}
	return result
}

// func (g *GestoreSpedizioni) Ritorna_pacchi_spedizione(ID string) []Pacco{
// 	collection := g.client.Database("APL").Collection("spedizioni")
// 	filter := bson.D{{Key: "id", Value: ID}}
// 	cur, err := collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(context.TODO())
// 	var result Spedizione
// 	err = cur.Decode(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return result.Pacchi

// }

func (g *GestoreSpedizioni) Insert_Spedizione(ID string, mittente string, destinatario string, sede string) {
	collection := g.client.Database("APL").Collection("spedizioni")
	var Stati []Stato
	Stati = append(Stati, InPreparazione)
	spedizione := Spedizione{ID, mittente, destinatario, Stati, time.Now(), time.Time{}, 0, []Pacco{}}
	insertResult, err := collection.InsertOne(context.TODO(), spedizione)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserita una nuova spedizione con ID:", insertResult.InsertedID)
	//TO_DO potrebbbe essere necessario inserire il pacco nel magazzino
}
func (g *GestoreSpedizioni) Insert_Pacco_spedizione(ID string, Peso float64, Dimensione string, Prezzo float64) error {
	collection := g.client.Database("APL").Collection("spedizioni")
	Pacco := Pacco{Spedizione_id: ID, Peso: Peso, Dimensione: Dimensione, Prezzo: Prezzo}
	// filter := bson.M{{"id": ID}}
	var spedizione Spedizione
	err := collection.FindOne(g.ctx, bson.M{"id": ID}).Decode(&spedizione)
	if err != nil {
		return errors.New("Spedizione inesistente")
	}
	// err = cur.Decode(&spedizione)
	// if err != nil {
	// 	print("Errore decode")
	// 	return err
	// }
	Pacchi := spedizione.Pacchi
	Pacchi = append(Pacchi, Pacco)
	// defer cur.Close(context.TODO())
	numero_pacchi := len(Pacchi)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "numero_pacchi", Value: numero_pacchi}, {Key: "pacchi", Value: Pacchi}}}}
	// update := bson.D{{Key: "$set", Value: bson.D{{Key: "numero_pacchi", Value: numero_pacchi}}}, {Key: "$push", Value: bson.D{{Key: "pacchi", Value: Pacchi}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), bson.M{"id": ID}, update)
	if err != nil {
		print("Errore update")
		return err
	}
	fmt.Printf("Modificati %v documenti\n", updateResult.ModifiedCount)
	//TO_DO potrebbbe essere necessario inserire il pacco nel magazzino
	return nil
}

func (g *GestoreSpedizioni) RitornaID() []string {
	collection := g.client.Database("APL").Collection("spedizioni")
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
func (g *GestoreSpedizioni) Modifica_Data_Consegna_Spedizione(id string, data string) {
	date, err := time.Parse("2006/01/02", data)
	if err != nil {
		log.Fatal(err)
	}
	collection := g.client.Database("APL").Collection("spedizioni")
	filter := bson.D{{Key: "idspedizione", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "data_consegna", Value: date}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Modificati %v documenti\n", updateResult.ModifiedCount)
}
func (g *GestoreSpedizioni) Ritorna_Data_Spedizione(id string) time.Time {
	//TO_DO funzione modifica, però prima va cambiato il database in non relazionale
	collection := g.client.Database("APL").Collection("spedizioni")
	filter := bson.D{{Key: "idspedizione", Value: id}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())
	var date time.Time
	err = cur.Decode(&date)
	if err != nil {
		log.Fatal(err)
	}
	return date
}
func (g *GestoreSpedizioni) Modifica_Stato_Spedizione(id string, stato string) {
	//TO_DO funzione modifica, però prima va cambiato il database in non relazionale
	collection := g.client.Database("APL").Collection("spedizioni")
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "stato", Value: ToStato(stato)}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Modificati %v documenti\n", updateResult.ModifiedCount)
}

func Tracciamento(spedizione Spedizione) string {
	// var String string
	SpedizioneString := "Id " + spedizione.ID + " Mittente" + spedizione.Mittente + "Destinatario " + spedizione.Destinatario + " Numero Pacchi: " + strconv.Itoa(spedizione.NumeroPacchi) + "Pacchi: "
	for _, pacco := range spedizione.Pacchi {
		// Pacco := "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		Pacco := "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Dimensione" + pacco.Dimensione + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
		SpedizioneString = SpedizioneString + Pacco
	}
	SpedizioneString = SpedizioneString + "Tracciamento eventi: "
	for _, stato := range spedizione.Stato {
		SpedizioneString = SpedizioneString + stato.String()
	}
	// String = SpedizioneString
	return SpedizioneString
}
func ToString(spedizioni []Spedizione) string {
	var SpedizioneString string
	for _, s := range spedizioni {
		SpedizioneString := "Id " + s.ID + " Mittente" + s.Mittente + "Destinatario " + s.Destinatario + "Stato: " + s.Stato[len(s.Stato)-1].String() + " Numero Pacchi: " + strconv.Itoa(s.NumeroPacchi) + "Pacchi: "
		for _, pacco := range s.Pacchi {
			// Pacco := "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Lunghezza" + strconv.FormatFloat(pacco.Lunghezza, 'f', -1, 64) + "Altezza" + strconv.FormatFloat(pacco.Altezza, 'f', -1, 64) + "Profondità" + strconv.FormatFloat(pacco.Profondità, 'f', -1, 64) + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
			Pacco := "Peso" + strconv.FormatFloat(pacco.Peso, 'f', -1, 64) + "Dimensione" + pacco.Dimensione + "Prezzo" + strconv.FormatFloat(pacco.Prezzo, 'f', -1, 64)
			SpedizioneString = SpedizioneString + Pacco
		}
	}

	// String = SpedizioneString
	return SpedizioneString
}
