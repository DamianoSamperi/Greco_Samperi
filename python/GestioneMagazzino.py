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
cliente1 = Cliente(nome="Edoardo", cognome="Salamanca", eta=23, citta ="Bologna", CF="ASL00PI397HDGV2WD", indirizzo="Via Romania 16")

magazzino.aggiungi_pacco_cliente(cliente1)

magazzino.gestisci_magazzino()

    #mostra inventario aggiurnato
magazzino.mostra_inventario()



   


