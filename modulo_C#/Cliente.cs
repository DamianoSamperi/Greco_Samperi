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
    public partial class Cliente : Form
    {
        private string indirizzo = null;
        public Cliente()
        {
            InitializeComponent();
        }
        public Cliente(string indirizzo)
        {
            InitializeComponent();
            this.indirizzo = indirizzo;
        }

        private void btn_inserisci_Click(object sender, EventArgs e)
        {
            this.Close();
            Form newForm = new FormDatiSpedizione(indirizzo);
            newForm.ShowDialog();
            newForm = null;
        }

        private async void btn_visualizza_Click(object sender, EventArgs e)
        {
            string id_spedizione = Microsoft.VisualBasic.Interaction.InputBox("Inserisci codice tracciamento per visualizzare spedizione", "Inserimento codice tracciamento", "xxxxxxx");
            if (id_spedizione != "")
            {
                var data = new
                {
                    Id_Spedizione = id_spedizione
                };


                // Converti i dati in formato JSON
                string jsonData = JsonSerializer.Serialize(data);

                try
                {
                    using (HttpClient client = new HttpClient())
                    {
                        HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                        HttpResponseMessage response = await client.PostAsync("http://localhost:8080/tracciamento_spedizione", stringContent);
                        var contents = await response.Content.ReadAsStringAsync();
                        if (contents != "")
                        {
                            MessageBox.Show(contents);
                        }
                        else
                        {
                            MessageBox.Show("Codice non valido");
                        }
                    }

                }
                catch (Exception)
                {
                    MessageBox.Show("Connessione al server rifiutata");
                    throw;
                }
            }
            else
            {
                MessageBox.Show("Inserire un codice spedizione");
            }
        }

        private async void btn_visualizza_Click_1(object sender, EventArgs e)
        {

            // Converti i dati in formato JSON
            string jsonData = JsonSerializer.Serialize(indirizzo);

            try
            {
                using (HttpClient client = new HttpClient())
                {
                    HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                    HttpResponseMessage response = await client.PostAsync("http://localhost:8080/Visualizza_Spedizioni", stringContent);
                    var contents = await response.Content.ReadAsStringAsync();
                    if (contents != "")
                    {
                        MessageBox.Show(contents);
                    }
                    else
                    {
                        MessageBox.Show("non hai nessuna spedizione effettuata");
                    }
                }

            }
            catch (Exception)
            {
                MessageBox.Show("Connessione al server rifiutata");
                throw;
            }
        }
    }
}
