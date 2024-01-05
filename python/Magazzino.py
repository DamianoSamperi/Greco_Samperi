import time
from Prodotto import Prodotto


class Magazzino:
    def __init__(self, gestore_spedizioni):
        self.codice_prodotto = None
        self.quantita = None
        #Crea un dizionario vuoto e lo assegna all'attributo di istanza 'inventario'
        self.inventario = {}

        self.limiti_dimensione = {
            "piccolo": 25,
            "medio": 15,
            "grande": 10
        }

        self.gestore_spedizioni = gestore_spedizioni



    def aggiungi_prodotto(self, prodotto):
        if prodotto.dimensione in self.limiti_dimensione:
            limite = self.limiti_dimensione[prodotto.dimensione]
            quantita_totale_dimensione = sum(
                p.quantita for p in self.inventario.values() if p.dimensione == prodotto.dimensione
            ) + prodotto.quantita

            if quantita_totale_dimensione > limite:
                print(f"Errore: Limite di pacchi {prodotto.dimensione} raggiunto ({limite}).")
                quantita_esistente = sum(
                    p.quantita for p in self.inventario.values() if p.dimensione == prodotto.dimensione
                )
                prodotto.quantita = limite - quantita_esistente
                return


        if prodotto.codice in self.inventario:
            self.inventario[prodotto.codice].quantita += prodotto.quantita
        else:
            self.inventario[prodotto.codice] = prodotto


    

    def mostra_inventario(self):
        print("Inventario del magazzino:")
        for prodotto in self.inventario.values():
            print(prodotto)




    def set_input_utente(self):
        self.ordini = []  # Lista per memorizzare gli ordini
        continua_inserimento = True
        flag = False
    
        
        while continua_inserimento:
            codice_prodotto = input("Inserisci il codice del prodotto: ")
            quantita = int(input("Inserisci la quantità desiderata da aggiungere: "))

            # Verifica se il prodotto è presente nel magazzino
            if codice_prodotto in self.inventario:
                    prodotto = self.inventario[codice_prodotto]

                # Verifica se la quantità da aggiungere supera il limite per la dimensione
                    if prodotto.dimensione in self.limiti_dimensione:
                        limite = self.limiti_dimensione[prodotto.dimensione]
                        quantita_totale_dimensione = sum(
                        ordine['quantita'] for ordine in self.ordini if ordine['codice_prodotto'] == codice_prodotto
                        ) + sum(
                        p.quantita for p in self.inventario.values() if p.codice == codice_prodotto and p.dimensione == prodotto.dimensione
                        ) + quantita
                        

                        if quantita_totale_dimensione > limite:
                            print(f"Attenzione: Limite di pacchi {prodotto.dimensione} raggiunto. non è stato aggiunto nessun pacco del nuovo ordine, riprovare con meno pacchi")
                            quantita_esistente = sum(
                            ordine['quantita'] for ordine in self.ordini if ordine['codice_prodotto'] == codice_prodotto
                            ) + sum(
                             p.quantita for p in self.inventario.values() if p.codice == codice_prodotto and p.dimensione == prodotto.dimensione
                            )

                            quantita = limite - quantita_esistente
                            flag = True
                            break
                            

                    # Aggiungi l'ordine alla lista
                    self.ordini.append({'codice_prodotto': codice_prodotto, 'quantita': quantita})

                    risposta = input("Desideri inserire un altro prodotto? (si/no): ")
                    continua_inserimento = risposta.lower() == 'si'
            else:
                print(f"Nessun prodotto trovato con il codice {codice_prodotto}.")
        return flag




    def gestisci_magazzino(self, cliente):
        flag = 0
        flag = self.set_input_utente()
        if flag == True:
            return
        
        else:

            mittente = input("Inserisci il mittente per la spedizione: ")
            destinatario = input("Inserisci il destinatario per la spedizione: ")

            for ordine in self.ordini:
                codice_prodotto = ordine['codice_prodotto']
                quantita = ordine['quantita']
            
                if codice_prodotto in self.inventario:
                    prodotto = self.inventario[codice_prodotto]
                    cliente.aggiunge_ordine(prodotto, quantita)
                    prodotto.quantita += quantita

                    print(f"Aggiunto/i {quantita} {prodotto.nome} al magazzino.")
                    time.sleep(4)
                
                    # Crea un'istanza di Spedizione utilizzando il GestoreSpedizioni
                    sped = self.gestore_spedizioni.crea_spedizione(mittente=mittente, destinatario=destinatario)
                    sped.aggiungi_evento_tracciamento(f"Prodotto/i spedito a {destinatario}")
                    sped.tracciamento()
                else:
                    print(f"Nessun prodotto trovato con il codice {codice_prodotto}.")



   
    def calcola_costo_ordine_cliente(self):
        costo_totale = 0
        
        for ordine in self.ordini:
            codice_prodotto = ordine['codice_prodotto']
            quantita = ordine['quantita']

            if codice_prodotto in self.inventario:
                prodotto = self.inventario[codice_prodotto]
                costo_totale += prodotto.calcola_prezzo() * quantita
            else:
                print(f"Nessun prodotto trovato con il codice {codice_prodotto}.")

        if costo_totale > 0:
            print(f"Costo totale dell'ordine: {costo_totale} euro.")
        


    


    def aggiungi_prodotto_cliente(self, cliente):
        while True:
            # Ottieni l'ultimo codice presente nel magazzino
            if self.inventario:
                ultimo_codice = max(self.inventario.keys())
                nuovo_codice = str(int(ultimo_codice) + 1)
            else:
                nuovo_codice = "1"

            # Richiedi all'utente di inserire le informazioni del nuovo prodotto
            nome = input("Inserisci il nome del nuovo prodotto: ")
            quantita = int(input("Inserisci la quantità del nuovo prodotto: "))
            prezzo = float(input("Inserisci il prezzo del nuovo prodotto: "))
            dimensione = input("Inserisci la dimensione del nuovo prodotto: ")

            # Crea un nuovo oggetto Prodotto con il nuovo codice
            nuovo_prodotto = Prodotto(codice=nuovo_codice, nome=nome, quantita=quantita, prezzo=prezzo, dimensione=dimensione)

            # Aggiungi il nuovo prodotto al magazzino
            self.aggiungi_prodotto(nuovo_prodotto)

            if nuovo_prodotto.dimensione in self.limiti_dimensione:
                        limite = self.limiti_dimensione[nuovo_prodotto.dimensione]
                        quantita_totale_dimensione = sum(
                        p.quantita for p in self.inventario.values() if p.dimensione == nuovo_prodotto.dimensione
                        ) + nuovo_prodotto.quantita

                        if quantita_totale_dimensione < limite:
                            print(f"Prodotto {nuovo_prodotto.nome} aggiunto al magazzino e all'inventario del cliente {cliente.nome} con codice {nuovo_prodotto.codice}.")

            continua = input("Vuoi aggiungere un altro prodotto? (si/no): ")
            if continua.lower() != 'si':
                break


    
                

