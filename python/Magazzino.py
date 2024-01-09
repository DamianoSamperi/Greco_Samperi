import requests
import time
from Pacco import Pacco


class Magazzino:
    def __init__(self, gestore_spedizioni):
        self.codice_prodotto = None
        #Crea un dizionario vuoto e lo assegna all'attributo di istanza 'inventario'
        self.inventario = []

        self.limiti_dimensione = {
            "piccolo": 25,
            "medio": 15,
            "grande": 10
        }

        self.gestore_spedizioni = gestore_spedizioni



    def aggiungi_pacco(self, pacco):
        if pacco.dimensione in self.limiti_dimensione:
            limite = self.limiti_dimensione[pacco.dimensione]
            quantita_totale_dimensione = sum(
                1 for p in self.inventario if p.dimensione == pacco.dimensione
            )

            if quantita_totale_dimensione >= limite:
                print(f"Errore: Limite di pacchi {pacco.dimensione} raggiunto ({limite}).")
                return

        if pacco not in self.inventario:
            self.inventario.append(pacco)

    

    def mostra_inventario(self):
        print("Inventario del magazzino:")
        for pacco in self.inventario:
            print(pacco)




    def gestisci_magazzino(self):            
            mittente = input("Inserisci il mittente per la spedizione: ")
            destinatario = input("Inserisci il destinatario per la spedizione: ")

            print(f"Aggiunto al magazzino.")
            time.sleep(4)
                
            # Crea un'istanza di Spedizione utilizzando il GestoreSpedizioni
            sped = self.gestore_spedizioni.crea_spedizione(mittente=mittente, destinatario=destinatario)
            '''
            url = "http://localhost:8080/Inserisci_Spedizione"

            # Dati da inviare nella richiesta POST, se necessario
            payload = {
                    "parametro1": sped,
                    "parametro2": "sede"                    
                    
                      }   

            # Effettua la richiesta POST
            response = requests.post(url, data=payload)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)
            '''
            url2 = "http://localhost:8080/Ritorna_sede"

            # Dati da inviare nella richiesta POST, se necessario
            payload = {
                    "parametro1": mittente,
                                       
                    
                      }   

            # Effettua la richiesta POST
            response = requests.post(url2, data=payload)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)



            sped.aggiungi_evento_tracciamento(f"Pacco spedito a {destinatario}")
            sped.tracciamento()
                
            print("Ordine completato.")



    def aggiungi_pacco_cliente(self, cliente):
        self.ordini = []
        while True:
            # Ottieni l'ultimo codice presente nel magazzino
            if self.inventario:
                ultimo_codice = max(pacco.codice for pacco in self.inventario)
                nuovo_codice = ultimo_codice + str(1)
            else:
                nuovo_codice = "1"

            peso = float(input("Inserisci il peso in grammi: "))
            dimensione = input("Inserisci la dimensione del nuovo pacco: ")

            # Crea un nuovo oggetto Pacco con il nuovo codice
            nuovo_pacco = Pacco(codice=nuovo_codice, peso=peso, prezzo=100, dimensione=dimensione)

            # Aggiungi il nuovo pacco al magazzino
            self.aggiungi_pacco(nuovo_pacco)
            if nuovo_pacco.dimensione in self.limiti_dimensione:
                limite = self.limiti_dimensione[nuovo_pacco.dimensione]

                # Calcola il numero totale di pacchi della dimensione specificata presenti nell'inventario
                quantita_totale_dimensione = sum(
                    1 for ordine in self.ordini if ordine['dimensione'] == dimensione
                ) + sum(
                    1 for p in self.inventario if p.dimensione == nuovo_pacco.dimensione
                )

                if quantita_totale_dimensione > limite:
                    print(f"Attenzione: Limite di pacchi {nuovo_pacco.dimensione} raggiunto. Non è stato aggiunto nessun pacco del nuovo ordine, riprovare con meno pacchi")
                    continue
                else:
                    print(f"Pacco {nuovo_pacco.codice} aggiunto al magazzino e all'inventario con codice {nuovo_pacco.codice}.")
           
            cliente.aggiunge_ordine(nuovo_pacco)
            continua = input("Vuoi aggiungere un altro pacco? (si/no): ")
            if continua.lower() != 'si':
                break



    
                

