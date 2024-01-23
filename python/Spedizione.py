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
        self.data_consegna_stimata = datetime.min
        self.eventi_tracciamento = []

    
    def aggiungi_evento_tracciamento(self, nuovo_stato):
        self.stato_attuale = nuovo_stato
        self.eventi_tracciamento.append((datetime.now(), nuovo_stato))

    def tracciamento(self):
        tracciamento_str = (
            f"Codice Spedizione: {self.codice_spedizione}\n"
            f"Mittente: {self.mittente}\n"
            f"Destinatario: {self.destinatario}\n"
            f"Stato attuale: {self.stato_attuale}\n"
            f"Data Spedizione: {self.data_spedizione.strftime('%Y-%m-%d %H:%M:%S')}\n"
            f"Data Consegna Stimata: {self.data_consegna_stimata.strftime('%Y-%m-%d %H:%M:%S')}\n"
            f"Eventi di Tracciamento: {self.eventi_tracciamento}\n"
        )

        for data, evento in self.eventi_tracciamento:
            tracciamento_str += f"{data.strftime('%Y-%m-%d %H:%M:%S')}: {evento}\n"

        return tracciamento_str


class GestoreSpedizioni:
    def __init__(self):
        self.spedizioni_attive = {}

    def genera_codice_spedizione(self):
       lunghezza_codice = 8
       caratteri_validi = string.ascii_uppercase + string.digits
       return ''.join(random.choice(caratteri_validi) for _ in range(lunghezza_codice))    
       

    def crea_spedizione(self, mittente, destinatario):
        codice_spedizione = self.genera_codice_spedizione()
        spedizione = Spedizione(codice_spedizione, mittente, destinatario)
        self.spedizioni_attive[codice_spedizione] = spedizione
        return spedizione, codice_spedizione
            
       

    

