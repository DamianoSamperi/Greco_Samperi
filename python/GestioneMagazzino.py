from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni
from fastapi import FastAPI

#app = FastAPI()

#@app.get("/")
#def requestMagazzino():
gestore_spedizioni = GestoreSpedizioni()
# Creazione di un oggetto Magazzino
magazzino = Magazzino(gestore_spedizioni)
# Creazione di un cliente1

nome = input("Inserisci il tuo nome: ")
cognome = input("Inserisci il tuo cognome: ")
indirizzo = input("Inserisci il tuo indirizzo: ")
cliente1 = Cliente(nome=nome, cognome=cognome, indirizzo=indirizzo)

magazzino.gestisci_magazzino()

magazzino.aggiungi_pacco_cliente(cliente1)

print("Ordine completato.")
#mostra inventario aggiurnato
magazzino.mostra_inventario()



   


