class Cliente:
    def __init__(self, nome, cognome, eta, citta, CF, indirizzo):
        self.nome = nome
        self.cognome = cognome
        self.eta = eta
        self.citta = citta
        self.CF = CF
        self.indirizzo = indirizzo
        self.ordini = []

    def aggiunge_ordine(self, pacco):
        self.ordini.append(pacco)
        print("Caro/a {} {}, il tuo ordine è stato preso in carico da ...".format(self.nome, self.cognome))



    