    
class Cliente:
    def __init__(self, nome, cognome, eta, citta, CF, indirizzo):
        self.nome = nome
        self.cognome = cognome
        self.eta = eta
        self.citta = citta
        self.CF = CF
        self.indirizzo = indirizzo

    def effettua_ordine(self, prodotto, quantita):
        print("Caro/a {} {}, hai effettuato un ordine di {} {}.".format(self.nome, self.cognome, quantita, prodotto.nome))