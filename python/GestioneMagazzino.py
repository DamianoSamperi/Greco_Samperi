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
        
        
        cliente = cliente_data.get('cliente')
        nome_cliente = cliente.nome
        cognome_cliente = cliente.cognome

        #risposta JSON senza ordinare i campi
        response_data = {
            'messaggio': 'Pacco aggiunto correttamente a nome del cliente',
            'nome': nome_cliente,
            'cognome': cognome_cliente
        }

        # Conversione dizionario in una stringa JSON mantenendo l'ordine dei campi
        response_json = json.dumps(response_data, sort_keys=False)
        magazzino.aggiungi_pacco_cliente(data)        
        return response_json, 200, {'Content-Type': 'application/json'}
    


    
@app.route('/epilogo_ordine', methods=['GET'])
def epilogo_ordine():
    dati_epilogo = magazzino.epilogo_ordine()
    #jsonify per convertire i dati in formato JSON e inviarli come risposta
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

class ThreadManager:
    def __init__(self):
        self.executor = ThreadPoolExecutor(max_workers=4) #numero max di threads in un pool(core)
        self.cliente_data = {}
        self.cliente_data_mutex = threading.Lock()
        self.gestore_spedizioni = GestoreSpedizioni()
        self.magazzino = Magazzino(self.gestore_spedizioni)

    def invia_dati(self, data):
        nome = data.get('nome', '')
        cognome = data.get('cognome', '')
        cliente = Cliente(nome=nome, cognome=cognome)

        with self.cliente_data_mutex:
            self.cliente_data['cliente'] = cliente

        return {'messaggio': 'dati cliente inviati correttamente'}

    def invia_dati_ordine(self, data):
        self.executor.submit(self.magazzino.gestisci_magazzino, data)
        return {'messaggio': 'Dati ordine inviati correttamente.'}

    def aggiungi_pacco_cliente(self, data):
        with self.cliente_data_mutex:
            cliente = self.cliente_data.get('cliente')
            nome_cliente = cliente.nome
            cognome_cliente = cliente.cognome

        response_data = {
            'messaggio': 'Pacco aggiunto correttamente a nome del cliente',
            'nome': nome_cliente,
            'cognome': cognome_cliente
        }

        self.executor.submit(self.magazzino.aggiungi_pacco_cliente, data)
        return response_data

    def epilogo_ordine(self):
        future = self.executor.submit(self.magazzino.epilogo_ordine)
        dati_epilogo = future.result()
        return dati_epilogo

thread_manager = ThreadManager()

@app.route('/invia_dati', methods=['POST'])
def invia_dati_route():
    if request.method == 'POST':
        data = request.json
        for _ in range(1):  # Esegui la funzione invia_dati 10 volte
            result = thread_manager.invia_dati(data)
        return jsonify(result)

@app.route('/invia_dati_ordine', methods=['POST'])
def invia_dati_ordine_route():
    if request.method == 'POST':
        data = request.json
        for _ in range(1):  # Esegui la funzione invia_dati_ordine 10 volte
            result = thread_manager.invia_dati_ordine(data)
        return jsonify(result)

@app.route('/aggiungi_pacco_cliente', methods=['POST'])
def aggiungi_pacco_cliente_route():
    if request.method == 'POST':
        data = request.json
        for _ in range(1):  # Esegui la funzione aggiungi_pacco_cliente 10 volte
            result = thread_manager.aggiungi_pacco_cliente(data)
        return jsonify(result)

@app.route('/epilogo_ordine', methods=['GET'])
def epilogo_ordine_route():
    for _ in range(1):  # Esegui la funzione epilogo_ordine 10 volte
        result = thread_manager.epilogo_ordine()
    return jsonify(result)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8082)
'''

   


