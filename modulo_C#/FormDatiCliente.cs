using System.Text;
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
            if (string.IsNullOrWhiteSpace(nome) || string.IsNullOrWhiteSpace(cognome))
            {
                MessageBox.Show("Inserire nome e cognome.", "Errore", MessageBoxButtons.OK, MessageBoxIcon.Error);
                return;
            }
            if (string.IsNullOrWhiteSpace(tb_Via_Mittente.Text) || string.IsNullOrWhiteSpace(tb_Città_Mittente.Text) || string.IsNullOrWhiteSpace(tb_Cod_Postale_Mittente.Text) || string.IsNullOrWhiteSpace(tb_Provincia_Mittente.Text))
            {
                MessageBox.Show("Compila tutti i campi del destinatario.", "Errore", MessageBoxButtons.OK, MessageBoxIcon.Error);
                return;
            }

            
            await InviaRichiestaPost(nome, cognome);

            this.Close();
            Form newForm = new Cliente(mittente);
            newForm.ShowDialog();
        }

        private async Task InviaRichiestaPost(string nome, string cognome)
        {
            string url = "http://localhost:8082/invia_dati";

            using (HttpClient client = new HttpClient())
            {               
                var data = new
                {
                    nome = nome,
                    cognome = cognome
                };

                
                string jsonData = JsonSerializer.Serialize(data);
                var content = new StringContent(jsonData, Encoding.UTF8, "application/json");
                HttpResponseMessage response = await client.PostAsync(url, content);

                
                if (response.IsSuccessStatusCode)
                {
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
