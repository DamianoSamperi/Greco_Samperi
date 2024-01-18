using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Diagnostics;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Net.Http;
using System.Text.Json;

namespace Modulo_C_
{
    public partial class Dati_Cliente : Form
    {

        public Dati_Cliente()
        {
            InitializeComponent();
        }




        public async void button_invia_Click(object sender, EventArgs e)
        {
            string nome = textBoxNome.Text;
            string cognome = textBoxCognome.Text;
            string mittente = tb_Via_Mittente.Text + "," + tb_Città_Mittente.Text + " " + tb_Cod_Postale_Mittente.Text + " " + tb_Provincia_Mittente.Text;
            // Esegui la richiesta POST
            await InviaRichiestaPost(nome, cognome);

            this.Close();
            Form newForm = new Cliente(mittente);
            newForm.ShowDialog();
            newForm = null;
            

        }

        private async Task InviaRichiestaPost(string nome, string cognome)
        {
            string url = "http://localhost:8082/invia_dati";

            using (HttpClient client = new HttpClient())
            {
                // Dati da inviare
                var data = new
                {
                    nome = nome,
                    cognome = cognome
                };

                // Converti i dati in formato JSON
                string jsonData = JsonSerializer.Serialize(data);

                // Crea il contenuto della richiesta
                var content = new StringContent(jsonData, Encoding.UTF8, "application/json");

                // Effettua la richiesta POST
                HttpResponseMessage response = await client.PostAsync(url, content);

                // Verifica se la richiesta ha avuto successo
                if (response.IsSuccessStatusCode)
                {
                    // Leggi la risposta
                    string responseContent = await response.Content.ReadAsStringAsync();
                    MessageBox.Show(responseContent, "Risposta dal server", MessageBoxButtons.OK, MessageBoxIcon.Information);
                }
                else
                {
                    MessageBox.Show($"Errore: {response.StatusCode}", "Errore", MessageBoxButtons.OK, MessageBoxIcon.Error);
                }
            }

        }


    }
}
