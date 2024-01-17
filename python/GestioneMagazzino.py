from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni
from flask import Flask, request, jsonify

app = Flask(__name__)
gestore_spedizioni = GestoreSpedizioni()  
magazzino = Magazzino(gestore_spedizioni)  

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
        
        return jsonify({'messaggio': 'Ordine completato. Inventario aggiornato.'})
        # Gestisci il magazzino e aggiungi il pacco del cliente




@app.route('/invia_dati_ordine', methods=['POST'])
def invia_dati_ordine():
    if request.method == 'POST':
        data = request.json
        magazzino.gestisci_magazzino(data)
        return jsonify({'messaggio': 'Dati ordine inviati correttamente.'})


        # Restituisci una risposta


@app.route('/aggiungi_pacco_cliente', methods=['POST'])
def aggiungi_pacco_cliente():
    if request.method == 'POST':
        data = request.json
        # Esegui l'elaborazione dei dati per aggiungere il pacco del cliente
        #nome = data.get('nome', '')
        #cognome = data.get('cognome', '')
        cliente = Cliente(nome="gabri", cognome="greco")
        magazzino.aggiungi_pacco_cliente(cliente, data)
        return jsonify({'messaggio': 'Pacco cliente aggiunto correttamente.'})
      

if __name__ == '__main__':
    app.run(debug=True,  host='0.0.0.0', port=8082)  # Cambia la porta se necessario


   


