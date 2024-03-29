﻿using System.Text;
using System.Text.Json;

namespace Modulo_C_
{
    public partial class FormCorriere : Form
    {
        public FormCorriere()
        {
            InitializeComponent();
        }

        private void label1_Click(object sender, EventArgs e)
        {

        }

        private async void btn_percorso_Click(object sender, EventArgs e)
        {
            string sede = Microsoft.VisualBasic.Interaction.InputBox("Inserisci sede in cui ti trovi", "Scelta hub", "Catania");
            if (sede != "")
            {
                try
                {
                    var data = new
                    {
                        sede = sede
                    };
                    HttpClient httpClient = new HttpClient();
                    string jsonData = JsonSerializer.Serialize(data);
                    HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                    HttpResponseMessage response = await httpClient.PostAsync("http://localhost:8080/Ottieni_Percorso", stringContent);
                    var contents = await response.Content.ReadAsStringAsync();
                    MessageBox.Show(contents);
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Connessione al server rifiutata: {ex.Message}");
                }

            }

        }

        private void btn_consegna_Click(object sender, EventArgs e)
        {
            this.Hide();
            //CustomControl customControl = new CustomControl();
            Form newForm = new Consegna();
            //newForm.Controls.Add(customControl);
            newForm.ShowDialog();
            this.Show();
        }
    }
}
