using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Modulo_C_
{
    public partial class Consegna : Form
    {
        public Consegna()
        {
            InitializeComponent();
        }

        private async void btn_hub_ClickAsync(object sender, EventArgs e)
        {
            string id_spedizione = textBox_id.Text;
            if (id_spedizione!="")
            {
                string hub = Microsoft.VisualBasic.Interaction.InputBox("Inserisci sede in cui consegnare il pacco", "Scelta hub", "Catania");
                if (hub != "")
                {
                    // Dati da inviare
                    var data = new
                    {
                        id = id_spedizione,
                        hub = hub
                    };

                    // Converti i dati in formato JSON
                    string jsonData = JsonSerializer.Serialize(data);

                    try
                    {
                        using (HttpClient client = new HttpClient())
                        {
                            HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                            HttpResponseMessage response = await client.PostAsync("http://localhost:8080//Modifica_Stato_Spedizione", stringContent);
                            var contents = await response.Content.ReadAsStringAsync();
                            MessageBox.Show(contents);
                        }

                    }
                    catch (Exception)
                    {
                        MessageBox.Show("Connessione al server rifiutata");
                        throw;
                    }

                }
            }
            else
            {
                MessageBox.Show("Inserire un codice spedizione");
            }

        }
    }
}
