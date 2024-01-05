class Cliente:
    def __init__(self, nome, cognome, eta, citta, CF, indirizzo):
        self.nome = nome
        self.cognome = cognome
        self.eta = eta
        self.citta = citta
        self.CF = CF
        self.indirizzo = indirizzo

    def aggiunge_ordine(self, prodotto, quantita):
        print("Caro/a {} {}, hai aggiunto un ordine di {} {}.".format(self.nome, self.cognome, quantita, prodotto.nome))



    