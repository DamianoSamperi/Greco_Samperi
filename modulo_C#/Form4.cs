using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Modulo_C_
{
    public partial class Form4 : Form
    {
        public Form4()
        {
            InitializeComponent();
        }

        private async void button_invia_locazione_Click(object sender, EventArgs e)
        {
            string mittente = textBoxMittente.Text;
            string destinatario = textBoxDestinatario.Text;

            await InviaRichiestaPost(mittente, destinatario);

            this.Hide();
            //CustomControl customControl = new CustomControl();
            Form newForm = new FormDatiPacco();
            //newForm.Controls.Add(customControl);
            newForm.ShowDialog();
            newForm = null;

            this.Show();
        }

        private async Task InviaRichiestaPost(string mittente, string destinatario)
        {
            string url = "http://localhost:8082/invia_dati_ordine"; // Cambia la porta e il percorso a seconda delle tue esigenze

            using (HttpClient client = new HttpClient())
            {
                // Dati da inviare
                var data = new
                {
                    mittente = mittente,
                    destinatario = destinatario
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
