import requests
import time
from Pacco import Pacco
import json
from flask import Flask, request, jsonify


class Magazzino:
    def __init__(self, gestore_spedizioni):
        self.cod_sped = None
        #self.codice_prodotto = None
        self.sede = None
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


    
    def gestisci_magazzino(self, data):


        # Esegui l'elaborazione dei dati      
            mittente = data.get('mittente', '')   
            destinatario = data.get('destinatario', '')   

            
            time.sleep(4)
                
            # Crea un'istanza di Spedizione utilizzando il GestoreSpedizioni
            sped, self.cod_sped = self.gestore_spedizioni.crea_spedizione(mittente=mittente, destinatario=destinatario)
            

            print("Ritorna Sede")
            url2 = "http://localhost:8080/Ritorna_Sede"

            payload = {"indirizzo": mittente }
            payload_json = json.dumps(payload)
            headers = {"Content-Type": "application/json"}
            response = requests.post(url2, data=payload_json, headers=headers)

            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)

            
            self.sede = response.text
            

    

            print("Inserisci Spedizione")
            url = "http://localhost:8080/Inserisci_Spedizione"
            payload = {
                    "id": self.cod_sped,
                    "mittente": mittente,
                    "destinatario": destinatario, 
                    "sede": self.sede                    
                      }   
            print("Id ",self.cod_sped)
            # Effettua la richiesta POST
            headers = {"Content-Type": "application/json"}
            response = requests.post(url, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)

            print("Ottieni Prodotti")
            url3 = "http://localhost:8080/Ottieni_Prodotti_Hub"
            payload = {
                    "sede": self.sede                   
                      }   

            #payload_json = json.dumps(payload)
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
            url5 = "http://localhost:8080/Visualizza_Spedizioni"

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

            



            sped.aggiungi_evento_tracciamento(f"Pacco in preparazione consegna")
            sped.tracciamento()
                
   
   



    def aggiungi_pacco_cliente(self, cliente, data):
        self.ordini = []
        while True:
            # Ottieni l'ultimo codice presente nel magazzino
            while True:
                try:
                    peso = float(data.get('peso', ''))
                    break
                except ValueError:
                    print("Errore: Inserisci un valore numerico.")

            
            while True:
                dimensione = data.get('dimensione', '')
    
                if dimensione in ["piccolo", "medio", "grande"]:
                    break  # Esci dal ciclo se l'input è valido
                else:
                    print("Errore: Inserisci una dimensione valida (piccolo, medio o grande).")

            # Crea un nuovo oggetto Pacco con il nuovo codice
            nuovo_pacco = Pacco(codice_sped=self.cod_sped, peso=peso, prezzo=100, dimensione=dimensione)

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
                    print(f"Pacco  aggiunto al magazzino e all'inventario con codice {nuovo_pacco.codice_sped}.")
            
            #pacco_json = json.dumps(nuovo_pacco.to_dict())
            print("Inserisci Prodotto Hub")
            url4 = "http://localhost:8080/Inserisci_Prodotto_Hub"
            payload = {
                       "sede": self.sede,
                       "pacco": nuovo_pacco.to_dict()              
                      }   
            
            #payload_json = json.dumps(payload)
            headers = {"Content-Type": "application/json"}
            # Effettua la richiesta POST
            response = requests.post(url4, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)
           

            print("Inserisci Pacco spedizione")
            url7 = "http://localhost:8080/Inserisci_Pacco_spedizione"
            payload = {
                       "id_spedizione": nuovo_pacco.codice_sped,
                       "peso": nuovo_pacco.peso,
                       "dimensione": nuovo_pacco.dimensione,
                       "prezzo": nuovo_pacco.calcola_prezzo()          
                      }   
            
            #payload_json = json.dumps(payload)
            headers = {"Content-Type": "application/json"}
            # Effettua la richiesta POST
            response = requests.post(url7, json=payload, headers=headers)
            if response.status_code == 200:
                print("Richiesta POST eseguita con successo!")
                print(response.text)
            else:
                print(f"Errore nella richiesta POST. Codice di stato: {response.status_code}")
                print(response.text)
           



            cliente.aggiunge_ordine(nuovo_pacco)
            continua = data.get('risposta', '')
            if continua.lower() != 'si':
                break

           


    
                

