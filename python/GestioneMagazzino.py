from Prodotto import Prodotto
from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni

# Creazione di un oggetto Magazzino
gestore_spedizioni = GestoreSpedizioni()
magazzino = Magazzino(gestore_spedizioni)

# Aggiunta di alcuni prodotti al magazzino
magazzino.aggiungi_prodotto(Prodotto(codice="1", nome="Computer", quantita=8, prezzo=800, dimensione='medio'))
magazzino.aggiungi_prodotto(Prodotto(codice="2", nome="Stampante", quantita=5, prezzo=150, dimensione='grande'))
magazzino.aggiungi_prodotto(Prodotto(codice="3", nome="Monitor", quantita=3, prezzo=200, dimensione='medio'))

# Creazione di un cliente1
cliente1 = Cliente(nome="Edoardo", cognome="Salamanca", eta=23, citta ="Bologna", CF="ASL00PI397HDGV2WD", indirizzo="Via Romania 16")

magazzino.mostra_inventario()



if __name__ == "__main__":
    #Gestione dell'ordine (rimuovendo i prodotti specificati dall'ordine)
    magazzino.gestisci_magazzino(cliente1)

    #calcola il costo dell'ordine
    costo_ordine = magazzino.calcola_costo_ordine_cliente()

    #Visualizzazione dell'inventario aggiornato
    magazzino.mostra_inventario()



