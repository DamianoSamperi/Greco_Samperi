from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni
from flask import Flask, request, jsonify, json

app = Flask(__name__)
gestore_spedizioni = GestoreSpedizioni()  
magazzino = Magazzino(gestore_spedizioni)  

cliente_data = {}

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

        cliente_data['cliente'] = cliente
        
        return jsonify({'messaggio':'dati cliente inviati correttamnete'})
        




@app.route('/invia_dati_ordine', methods=['POST'])
def invia_dati_ordine():
    if request.method == 'POST':
        data = request.json
        magazzino.gestisci_magazzino(data)
        return jsonify({'messaggio':'Dati ordine inviati correttamente.'})



@app.route('/aggiungi_pacco_cliente', methods=['POST'])
def aggiungi_pacco_cliente():
    if request.method == 'POST':
        data = request.json
        # Esegui l'elaborazione dei dati per aggiungere il pacco del cliente
        
        cliente = cliente_data.get('cliente')
        nome_cliente = cliente.nome
        cognome_cliente = cliente.cognome

        # Costruisci manualmente la risposta JSON senza ordinare i campi
        response_data = {
            'messaggio': 'Pacco aggiunto correttamente a nome del cliente',
            'nome': nome_cliente,
            'cognome': cognome_cliente
        }

        # Converti il dizionario in una stringa JSON mantenendo l'ordine dei campi
        response_json = json.dumps(response_data, sort_keys=False)

        magazzino.aggiungi_pacco_cliente(data)        
        return response_json, 200, {'Content-Type': 'application/json'}
    
@app.route('/epilogo_ordine', methods=['GET'])
def epilogo_ordine():
    dati_epilogo = magazzino.epilogo_ordine()
    # Utilizza jsonify per convertire i dati in formato JSON e inviarli come risposta
    return jsonify(dati_epilogo)


if __name__ == '__main__':
    app.run(debug=True,  host='0.0.0.0', port=8082)  

'''
from flask import Flask, request, jsonify
from concurrent.futures import ThreadPoolExecutor
import json
import threading
from Cliente import Cliente
from Magazzino import Magazzino
from Spedizione import GestoreSpedizioni


app = Flask(__name__)
gestore_spedizioni = GestoreSpedizioni()  
magazzino = Magazzino(gestore_spedizioni)  
cliente_data = {}

executor = ThreadPoolExecutor(max_workers=2)  # Imposta il numero desiderato di worker

@app.route('/invia_dati', methods=['POST'])
def invia_dati():
    if request.method == 'POST':
        data = request.json
        nome = data.get('nome', '')
        cognome = data.get('cognome', '')
        cliente = Cliente(nome=nome, cognome=cognome)
        cliente_data['cliente'] = cliente
        return jsonify({'messaggio': 'dati cliente inviati correttamente'})

@app.route('/invia_dati_ordine', methods=['POST'])
def invia_dati_ordine():
    if request.method == 'POST':
        data = request.json
        # Esegui la gestione del magazzino in un thread separato
        executor.submit(magazzino.gestisci_magazzino, data)
        return jsonify({'messaggio': 'Dati ordine inviati correttamente.'})

@app.route('/aggiungi_pacco_cliente', methods=['POST'])
def aggiungi_pacco_cliente():
    if request.method == 'POST':
        data = request.json
        cliente = cliente_data.get('cliente')
        nome_cliente = cliente.nome
        cognome_cliente = cliente.cognome
        response_data = {
            'messaggio': 'Pacco aggiunto correttamente a nome del cliente',
            'nome': nome_cliente,
            'cognome': cognome_cliente
        }
        response_json = json.dumps(response_data, sort_keys=False)
        # Esegui l'aggiunta del pacco del cliente in un thread separato
        executor.submit(magazzino.aggiungi_pacco_cliente, data)
        return response_json, 200, {'Content-Type': 'application/json'}

@app.route('/epilogo_ordine', methods=['GET'])
def epilogo_ordine():
    # Esegui la funzione epilogo_ordine in un thread separato
    future = executor.submit(magazzino.epilogo_ordine)
    dati_epilogo = future.result()
    return jsonify(dati_epilogo)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8082)

'''
   


