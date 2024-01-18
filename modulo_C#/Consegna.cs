﻿using System;
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
            if (id_spedizione != "")
            {
                string nuovo_hub = Microsoft.VisualBasic.Interaction.InputBox("Inserisci sede in cui consegnare il pacco", "Scelta hub", "Catania");
                if (nuovo_hub != "")
                {
                    string vecchio_hub = Microsoft.VisualBasic.Interaction.InputBox("Inserisci sede in cui hai ricevuto il pacco", "Scelta hub", "Catania");

                    // Dati da inviare
                    var data = new
                    {
                        Nuovo_Hub = nuovo_hub,
		                Vecchio_Hub   = vecchio_hub,
		                Id_Spedizione =id_spedizione
                    };
                    

                    // Converti i dati in formato JSON
                    string jsonData = JsonSerializer.Serialize(data);

                    try
                    {
                        using (HttpClient client = new HttpClient())
                        {
                            HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                            HttpResponseMessage response = await client.PostAsync("http://localhost:8080/Consegna_hub", stringContent);
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

        private async void btn_consegna_Click(object sender, EventArgs e)
        {
            string id_spedizione = textBox_id.Text;
            if (id_spedizione != "")
            {

                var data = new
                {
                    stato= "Consegnato",
                    id = id_spedizione
                };
                // Converti i dati in formato JSON
                string jsonData = JsonSerializer.Serialize(id_spedizione);

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
            else
            {
                MessageBox.Show("Inserire un codice spedizione");
            }
        }
    }
}
