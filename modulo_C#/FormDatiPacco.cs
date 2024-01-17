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
    public partial class FormDatiPacco : Form
    {
        public FormDatiPacco()
        {
            InitializeComponent();
        }


        private async void button_invia_dati_pacco_Click_1(object sender, EventArgs e)
        {
            string peso = textBoxPeso.Text;
            string dimensione = textBoxDimensione.Text;
            string risposta = "";

            for (int i = 0; i < checkedListBoxSiNo.Items.Count; i++)
            {
                if (checkedListBoxSiNo.GetItemChecked(i))
                {
                    risposta = checkedListBoxSiNo.GetItemText(checkedListBoxSiNo.Items[i]);
                    break;  // Esci dal ciclo una volta trovata la checkbox selezionata
                }
            }



            await InviaRichiestaPost(peso, dimensione, risposta);

            if (risposta == "si")
            {
                await InviaRichiestaPost(peso, dimensione, risposta);
            }
        }

        private async Task InviaRichiestaPost(string peso, string dimensione, string risposta)
        {
            string url = "http://localhost:8082/aggiungi_pacco_cliente"; // Cambia la porta e il percorso a seconda delle tue esigenze

            using (HttpClient client = new HttpClient())
            {
                // Dati da inviare
                var data = new
                {
                    peso = peso,
                    dimensione = dimensione,
                    risposta = risposta
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
