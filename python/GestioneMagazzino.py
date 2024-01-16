from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni
from fastapi import FastAPI
import sys

#app = FastAPI()

#@app.get("/")
#def requestMagazzino():
def main():
    gestore_spedizioni = GestoreSpedizioni()
    # Creazione di un oggetto Magazzino
    magazzino = Magazzino(gestore_spedizioni)
    # Creazione di un cliente1

    if len(sys.argv) >= 3:
        # Il primo argomento (sys.argv[0]) è il nome dello script, quindi il nome è il secondo argomento
        nome = sys.argv[1]
        cognome = sys.argv[2]


    print(f"Nome ricevuto dal programma: {nome}")
    print(f"Cognome ricevuto dal programma: {cognome}")
    cliente1 = Cliente(nome=nome, cognome=cognome)

    magazzino.gestisci_magazzino()

    magazzino.aggiungi_pacco_cliente(cliente1)

    print("Ordine completato.")
    #mostra inventario aggiurnato
    magazzino.mostra_inventario()

if __name__ == "__main__":
    main()


   


