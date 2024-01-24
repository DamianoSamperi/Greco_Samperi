using System.Text;
using System.Text.Json;

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

            //Nel contesto di decimal.TryParse, il metodo accetta un parametro out per memorizzare
            //il risultato della conversione:
            if (!decimal.TryParse(peso, out decimal pesoDecimal) || pesoDecimal <= 0)
            {
                MessageBox.Show("Inserire un peso valido.", "Errore", MessageBoxButtons.OK, MessageBoxIcon.Error);
                return;
            }

            // Verifica che la dimensione sia tra le opzioni valide
            List<string> opzioniDimensione = new List<string> { "piccolo", "medio", "grande" };
            if (string.IsNullOrWhiteSpace(dimensione) || !opzioniDimensione.Contains(dimensione.ToLower()))
            {
                MessageBox.Show("Inserire una dimensione valida (piccolo, medio o grande).", "Errore", MessageBoxButtons.OK, MessageBoxIcon.Error);
                return;
            }


            await InviaRichiestaPost(peso, dimensione);

            textBoxPeso.Text = string.Empty;
            textBoxDimensione.Text = string.Empty;
        }

        private async Task InviaRichiestaPost(string peso, string dimensione)
        {
            string url = "http://localhost:8082/aggiungi_pacco_cliente";
            try
            {
                using (HttpClient client = new HttpClient())
                {
                    
                    var data = new
                    {
                        peso = peso,
                        dimensione = dimensione
                    };

                    
                    string jsonData = JsonSerializer.Serialize(data);                   
                    var content = new StringContent(jsonData, Encoding.UTF8, "application/json");                    
                    HttpResponseMessage response = await client.PostAsync(url, content);

                   
                    if (response.IsSuccessStatusCode)
                    {                     
                        string responseContent = await response.Content.ReadAsStringAsync();
                        MessageBox.Show(responseContent, "Risposta dal server", MessageBoxButtons.OK, MessageBoxIcon.Information);
                    }

                }
            }
            catch (Exception ex)
            {
                MessageBox.Show($"Errore durante la richiesta POST: {ex.Message}", "Errore", MessageBoxButtons.OK, MessageBoxIcon.Error);
            }
        }

        private void buttonFineOrdine_Click(object sender, EventArgs e)
        {
            this.Close();

            Form newForm = new FormRiepilogoOrdine();
            newForm.ShowDialog();
            newForm = null;
        }
    }
}
