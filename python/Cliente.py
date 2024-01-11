class Cliente:
    def __init__(self, nome, cognome, indirizzo):
        self.nome = nome
        self.cognome = cognome        
        self.indirizzo = indirizzo
        self.ordini = []

    def aggiunge_ordine(self, pacco):
        self.ordini.append(pacco)
        print("Caro/a {} {}, il tuo ordine è stato preso in carico.".format(self.nome, self.cognome))



    