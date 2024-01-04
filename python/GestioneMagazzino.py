from Prodotto import Prodotto
from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni



# Creazione di un oggetto Magazzino
gestore_spedizioni = GestoreSpedizioni()
magazzino = Magazzino(gestore_spedizioni)

# Aggiunta di alcuni prodotti al magazzino
magazzino.aggiungi_prodotto(Prodotto(codice="001", nome="Computer", quantita=10, prezzo=800, dimensione='medio'))
magazzino.aggiungi_prodotto(Prodotto(codice="002", nome="Stampante", quantita=5, prezzo=150, dimensione='grande'))
magazzino.aggiungi_prodotto(Prodotto(codice="003", nome="Monitor", quantita=8, prezzo=200, dimensione=None))

# Creazione di un cliente
cliente1 = Cliente(nome="Alice", cognome="Salamanca", eta=23, citta ="Bologna", CF="ASL00PI397HDGV2WD", indirizzo="Via Romania 16")

# Visualizzazione dell'inventario aggiornato
magazzino.mostra_inventario()

# Gestione dell'ordine (rimuovendo i prodotti specificati dall'ordine)
magazzino.gestisci_ordine(cliente1)

cod_prod = magazzino.get_codice_prodotto()
quant = magazzino.get_quantita()

ordine_cliente = {cod_prod: quant}
# Calcolo del costo totale dell'ordine
costo_ordine = magazzino.calcola_costo_ordine_cliente(cliente1, ordine_cliente)

# Visualizzazione dell'inventario aggiornato
magazzino.mostra_inventario()



