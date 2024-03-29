﻿using System.Text;
using System.Text.Json;

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
                catch (Exception ex)
                {
                    MessageBox.Show($"Connessione al server rifiutata: {ex.Message}");
                }
            }
            else
            {
                MessageBox.Show("Inserire un codice spedizione");
            }
        }

        private async void btn_visualizza_Click_1(object sender, EventArgs e)
        {
            var data = new
            {
                mittente = indirizzo
            };
            
            string jsonData = JsonSerializer.Serialize(data);

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
            catch (Exception ex)
            {
                MessageBox.Show($"Connessione al server rifiutata: {ex.Message}");
            }
        }

        private async void btn_data_Click(object sender, EventArgs e)
        {
            string id_spedizione = Microsoft.VisualBasic.Interaction.InputBox("Inserisci codice tracciamento spedizione", "Inserimento codice tracciamento", "xxxxxxx");
            if (id_spedizione != "")
            {
                string nuova_data = Microsoft.VisualBasic.Interaction.InputBox("Inserisci nuova data consegna (superiore a 4 giorni da oggi)", "Inserimento data consegna", DateTime.Today.AddDays(4).ToString("yyyy/MM/dd"));
                if (DateTime.Parse(nuova_data)>= DateTime.Today.AddDays(4)){

                    var data = new
                    {
                        id = id_spedizione,
                        data = nuova_data

                    };
                    string jsonData = JsonSerializer.Serialize(data);

                    try
                    {
                        using (HttpClient client = new HttpClient())
                        {
                            HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                            HttpResponseMessage response = await client.PostAsync("http://localhost:8080/Scegli_data_consegna", stringContent);
                            var contents = await response.Content.ReadAsStringAsync();
                            MessageBox.Show(contents);
                          
                        }

                    }
                    catch (Exception ex)
                    {
                        MessageBox.Show($"Connessione al server rifiutata: {ex.Message}");
                    }
                }
                else
                {
                    MessageBox.Show("Inserire data consegna valida");
                }

            }else{
                MessageBox.Show("Inserire un codice spedizione");
            }
            
        }
    }
}
