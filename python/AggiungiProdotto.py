import GestioneMagazzino

magazzino = GestioneMagazzino.magazzino
cliente1 = GestioneMagazzino.cliente1


#aggiunge nuovo prodotto in magazzino
magazzino.aggiungi_prodotto_cliente(cliente1)

#mostra inventario aggiurnato
magazzino.mostra_inventario()
