from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni
from flask import Flask, request, jsonify

app = Flask(__name__)
gestore_spedizioni = GestoreSpedizioni()  # Assicurati di importare la classe 'GestoreSpedizioni'
magazzino = Magazzino(gestore_spedizioni)  # Assicurati di importare la classe 'Magazzino'

@app.route('/invia_dati', methods=['POST'])
def invia_dati():
    if request.method == 'POST':
        # Ottieni i dati dal corpo della richiesta
        data = request.json

        # Esegui l'elaborazione dei dati
        nome = data.get('nome', '')
        cognome = data.get('cognome', '')

        # Crea un oggetto Cliente con i dati ricevuti
        cliente = Cliente(nome=nome, cognome=cognome)

        # Gestisci il magazzino e aggiungi il pacco del cliente
        magazzino.gestisci_magazzino()
        magazzino.aggiungi_pacco_cliente(cliente)

        # Restituisci una risposta
        return jsonify({'messaggio': 'Ordine completato. Inventario aggiornato.'})

if __name__ == '__main__':
    app.run(debug=True,  host='0.0.0.0', port=8082)  # Cambia la porta se necessario


   


