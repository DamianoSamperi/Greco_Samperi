import random
import string
from datetime import datetime

class Spedizione:
    def __init__(self, codice_spedizione, mittente, destinatario):
        self.codice_spedizione = codice_spedizione
        self.mittente = mittente
        self.destinatario = destinatario
        self.stato_attuale = "In transito"
        self.data_spedizione = datetime.now()
        #self.data_consegna_stimata = self.data_spedizione + timedelta(days=5)
        self.eventi_tracciamento = []

    #aggiorna stato e aggiungi tracciamento sono insieme, bisogna chiamare la funzione che farà Damiano
    def aggiorna_stato(self, nuovo_stato):
        self.stato_attuale = nuovo_stato
        self.eventi_tracciamento.append((datetime.now(), nuovo_stato))


    def aggiungi_evento_tracciamento(self, evento):
        self.eventi_tracciamento.append((datetime.now(), evento))

    def tracciamento(self):
        print(f"Codice Spedizione: {self.codice_spedizione}")
        print(f"Mittente: {self.mittente}")
        print(f"Destinatario: {self.destinatario}")
        print(f"Stato attuale: {self.stato_attuale}")
        print(f"Data Spedizione: {self.data_spedizione.strftime('%Y-%m-%d %H:%M:%S')}")
        #print(f"Data Consegna Stimata: {self.data_consegna_stimata.strftime('%Y-%m-%d %H:%M:%S')}")
        print("Eventi di Tracciamento:")
        for data, evento in self.eventi_tracciamento:
            print(f"{data.strftime('%Y-%m-%d %H:%M:%S')}: {evento}")


class GestoreSpedizioni:
    def __init__(self):
        self.spedizioni_attive = {}

    def genera_codice_spedizione(self):
       lunghezza_codice = 8
       caratteri_validi = string.ascii_uppercase + string.digits
       return ''.join(random.choice(caratteri_validi) for _ in range(lunghezza_codice))    
       

    def crea_spedizione(self, mittente, destinatario):
        codice_spedizione = self.genera_codice_spedizione()
        #data_spedizione = datetime.now() + timedelta(hours=2)
        spedizione = Spedizione(codice_spedizione, mittente, destinatario)
        self.spedizioni_attive[codice_spedizione] = spedizione
        return spedizione
       

    

