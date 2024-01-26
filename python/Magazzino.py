import requests
from Pacco import Pacco
import json



class Magazzino:
    def __init__(self, gestore_spedizioni):
        self.cod_sped = None
        self.sede = None
        #lista vuota assegnata all'attributo di istanza 'inventario'
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

    

    
    def gestisci_magazzino(self, data):
            #elaborazione dei dati passati tramite richiesta Post     
            mittente = data.get('mittente', '')   
            destinatario = data.get('destinatario', '')   

            
            # time.sleep(4)
                
            # Creazione istanza Spedizione utilizzando il GestoreSpedizioni
            self.sped, self.cod_sped = self.gestore_spedizioni.crea_spedizione(mittente=mittente, destinatario=destinatario)
            

            print("Ritorna Sede")
            url2 = "http://go:8080/Ritorna_Sede"
            payload = {"indirizzo": mittente }
            headers = {"Content-Type": "application/json"}
            response = requests.post(url2, json=payload, headers=headers)
            print("ci sono arrivato tanto")
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)

            self.sede = response.text
            

    

            print("Inserisci Spedizione")
            url = "http://go:8080/Inserisci_Spedizione"
            payload = {
                    "id": self.cod_sped,
                    "mittente": mittente,
                    "destinatario": destinatario, 
                    "sede": self.sede                    
                      }   
            print("Id ",self.cod_sped)
            headers = {"Content-Type": "application/json"}
            response = requests.post(url, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)




            print("Ottieni Prodotti")
            url3 = "http://go:8080/Ottieni_Prodotti_Hub"
            payload = {
                    "sede": self.sede                   
                      }   
            headers = {"Content-Type": "application/json"}
            # Effettua la richiesta POST
            response = requests.post(url3, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)



            print("Visualizza Spedizioni")
            url5 = "http://go:8080/Visualizza_Spedizioni"
            payload = {"Mittente": mittente }
            payload_json = json.dumps(payload)  
            headers = {"Content-Type": "application/json"}
            response = requests.post(url5, json=payload_json, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)

            


    
   
   



    def aggiungi_pacco_cliente(self, data):
            self.ordini = []
            while True:
                try:
                    peso = float(data.get('peso', ''))
                    break
                except ValueError:
                    print("Errore: Inserisci un valore numerico.")

            
            while True:
                dimensione = data.get('dimensione', '')    
                if dimensione in ["piccolo", "medio", "grande"]:
                    break 
                else:
                    print("Errore: Inserisci una dimensione valida (piccolo, medio o grande).")

            #nuovo oggetto Pacco con il nuovo codice
            nuovo_pacco = Pacco(codice_sped=self.cod_sped, peso=peso, prezzo=100, dimensione=dimensione)

            # Aggiunta nuovo pacco al magazzino
            self.aggiungi_pacco(self.nuovo_pacco)
            

           
            print("Inserisci Prodotto Hub")
            url4 = "http://go:8080/Inserisci_Prodotto_Hub"
            payload = {
                       "sede": self.sede,
                       "pacco": self.nuovo_pacco.to_dict()              
                      }   
            
            headers = {"Content-Type": "application/json"}
            response = requests.post(url4, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)
           



            print("Inserisci Pacco spedizione")
            url7 = "http://go:8080/Inserisci_Pacco_spedizione"
            payload = {
                       "id_spedizione": nuovo_pacco.codice_sped,
                       "peso": nuovo_pacco.peso,
                       "dimensione": nuovo_pacco.dimensione,
                       "prezzo": nuovo_pacco.calcola_prezzo()          
                      }   
            
            headers = {"Content-Type": "application/json"}
            response = requests.post(url7, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)

            self.preventivo =  sum(nuovo_pacco.calcola_prezzo() for nuovo_pacco in self.inventario)

    def riepilogo_ordine(self):            
            tracciamento_corrente = self.sped.tracciamento()
            evento_aggiunto = self.sped.aggiungi_evento_tracciamento(f"Pacco in preparazione consegna")
            return {
                "tracciamento": tracciamento_corrente,
                "evento_aggiunto": evento_aggiunto,
                "prezzo": f"prezzo: {str(self.preventivo)}"
                }
                    

            

           


    
                

