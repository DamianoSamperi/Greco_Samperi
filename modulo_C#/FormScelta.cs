using System.Text.Json;
using System.Text;
using System.Net.Http.Json;

namespace Modulo_C_
{
    public partial class FormScelta : Form
    {
        public FormScelta()
        {
            InitializeComponent();
        }

        private void label1_Click(object sender, EventArgs e)
        {

        }

        private void btn_cliente_Click(object sender, EventArgs e)
        {
            this.Hide();
            Form newForm = new Dati_Cliente();
            newForm.ShowDialog();
           
            this.Show();
        }

        private async void btn_corriere_Click(object sender, EventArgs e)
        {
            string identificativo = Microsoft.VisualBasic.Interaction.InputBox("Inserisci identificativo corriere", "Identificazione", "000000");
            if (identificativo != "")
            {
                try
                {
                    //var data = new
                    //{
                    //    identificativo = identificativo
                    //};
                    var data = identificativo;
                    HttpClient httpClient = new HttpClient();
                    string jsonData = JsonSerializer.Serialize(data);
                    HttpContent stringContent = new StringContent(jsonData, Encoding.UTF8, "application/json");
                    HttpResponseMessage response = await httpClient.PostAsync("http://localhost:8080/identifica_corriere", stringContent);
                    //var contents = await response.Content.ReadAsStringAsync();
                    try {
                        bool contents = await response.Content.ReadFromJsonAsync<bool>();
                        if (contents)
                        {
                            this.Hide();
                            Form newForm = new FormCorriere();
                            newForm.ShowDialog();
                            this.Show();
                        }
                        else
                        {
                            MessageBox.Show("identificativo non valido");
                        }
                    }catch (JsonException)
                    {
                        MessageBox.Show("Errata risposta dal server");
                    }

                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Connessione al server rifiutata: {ex.Message}");
                }

            }
            else
            {
                MessageBox.Show("Inserisci un identificativo");
            }

        }
    }
}
