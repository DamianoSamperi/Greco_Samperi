class Magazzino:
    def __init__(self, gestore_spedizioni):
        self.codice_prodotto = None
        self.quantita = None
        #Crea un dizionario vuoto e lo assegna all'attributo di istanza 'inventario'
        self.inventario = {}
        self.gestore_spedizioni = gestore_spedizioni



    def aggiungi_prodotto(self, prodotto):
        if prodotto.codice in self.inventario:
            self.inventario[prodotto.codice].quantita += prodotto.quantita
        else:
            self.inventario[prodotto.codice] = prodotto

    

    def mostra_inventario(self):
        print("Inventario del magazzino:")
        for prodotto in self.inventario.values():
            print(prodotto)

    def set_input_utente(self):
        self.codice_prodotto = input("Inserisci il codice del prodotto: ")
        self.quantita = int(input("Inserisci la quantità desiderata: "))

    
    
    def gestisci_ordine(self, cliente):
        self.set_input_utente()
        if self.codice_prodotto in self.inventario:
            prodotto = self.inventario[self.codice_prodotto]
                 # Effettua l'ordine
            if prodotto.quantita >= self.quantita:
                    cliente.effettua_ordine(prodotto, self.quantita)
                    prodotto.quantita -= self.quantita

                    # Crea un'istanza di Spedizione utilizzando il GestoreSpedizioni
                    sped = self.gestore_spedizioni.crea_spedizione(mittente="Samperi_Greco_SPA", destinatario=cliente.nome + " " + cliente.cognome)

                    # Aggiungi un evento di tracciamento personalizzato
                    sped.aggiungi_evento_tracciamento(f"Prodotto/i spedito a {cliente.nome} {cliente.cognome}")

                    # Stampa le informazioni di tracciamento della spedizione
                    sped.tracciamento()
            else:
                print(f"Quantità insufficiente di {prodotto.nome} in magazzino. quantita massima disponibile: {prodotto.quantita}")
        else:
            print(f"Nessun prodotto trovato con il codice {self.codice_prodotto}.")



    def get_codice_prodotto(self):
        return self.codice_prodotto


    def get_quantita(self):
        return self.quantita


    def calcola_costo_ordine_cliente(self, cliente, ordine):
        costo_totale = 0
        for cod_prod, quant in ordine.items():
            if cod_prod in self.inventario:
                prodotto = self.inventario[cod_prod]
                costo_totale += prodotto.calcola_prezzo() * quant
            else:
                print(f"Nessun prodotto trovato con il codice {self.codice_prodotto}.")

        print(f"Costo totale dell'ordine per {cliente.nome}: {costo_totale} euro.")
