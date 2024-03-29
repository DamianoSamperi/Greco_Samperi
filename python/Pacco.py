
class Pacco:
    def __init__(self, codice_sped, peso, prezzo, dimensione):
        self.codice_sped = codice_sped
        self.peso = peso
        self.dimensione = dimensione
        self.prezzo_base_unita = prezzo
    
    

    def calcola_prezzo(self):
        if self.dimensione == 'piccolo':
            coefficiente_dimensione = 1.0
        elif self.dimensione == 'medio':
            coefficiente_dimensione = 1.2
        elif self.dimensione == 'grande':
            coefficiente_dimensione = 1.5

        if self.peso > 0 and self.peso < 5000:
            coefficiente_peso = 1.8
        elif self.peso > 5000 and self.peso < 20000:
            coefficiente_peso = 2.5
        elif self.peso > 20000:
            coefficiente_peso = 3.0

        prezzo_finale = self.prezzo_base_unita * coefficiente_dimensione * coefficiente_peso
        return prezzo_finale



    def __str__(self):
        return f"(Cod_sped: {self.codice_sped} Peso: {self.peso}, Dimensione: {self.dimensione}, Prezzo: {self.calcola_prezzo()} euro"
    

    def to_dict(self):
        return {
            'spedizione_id': self.codice_sped,
            'peso': self.peso,
            'prezzo': self.calcola_prezzo(),
            'dimensione': self.dimensione
        }