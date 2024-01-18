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
    public partial class FormDatiSpedizione : Form
    {
        public FormDatiSpedizione()
        {
            InitializeComponent();
        }

        private async void button_invia_locazione_Click(object sender, EventArgs e)
        {
            string destinatario = tb_Via_Destinatario.Text + "," + tb_Città_Destinatario.Text + " " + tb_Cod_Postale_Destinatario.Text + " " + tb_Provincia_Destinatario.Text;

            await InviaRichiestaPost(mittente, destinatario);

            this.Close();
            Form newForm = new FormDatiPacco();
            newForm.ShowDialog();
            newForm = null;
        }

        private async Task InviaRichiestaPost(string mittente, string destinatario)
        {
            string url = "http://localhost:8082/invia_dati_ordine";

            using (HttpClient client = new HttpClient())
            {

                var data = new
                {
                    mittente = mittente,
                    destinatario = destinatario
                };


                string jsonData = JsonSerializer.Serialize(data);

                var content = new StringContent(jsonData, Encoding.UTF8, "application/json");

                HttpResponseMessage response = await client.PostAsync(url, content);


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
