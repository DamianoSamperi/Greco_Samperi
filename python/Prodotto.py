
class Prodotto:
    def __init__(self, codice, nome, quantita, prezzo, dimensione):
        self.codice = codice
        self.nome = nome
        self.quantita = quantita
        self.dimensione = dimensione
        self.prezzo_base_unita = prezzo
    

    def calcola_prezzo(self):
        # Calcolo del prezzo basato sulla dimensione
        if self.dimensione == 'piccolo':
            coefficiente_dimensione = 1.0
        elif self.dimensione == 'medio':
            coefficiente_dimensione = 1.2
        elif self.dimensione == 'grande':
            coefficiente_dimensione = 1.5

        # Calcolo del prezzo finale
        prezzo_finale = self.prezzo_base_unita * coefficiente_dimensione 

        return prezzo_finale



    def __str__(self):
        return f"Prodotto: {self.nome} (Codice: {self.codice}), Dimensione: {self.dimensione}, Quantità: {self.quantita}, Prezzo: {self.calcola_prezzo()} euro"
