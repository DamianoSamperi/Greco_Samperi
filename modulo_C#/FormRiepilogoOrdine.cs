using System.Text.Json;

namespace Modulo_C_
{
    public partial class FormRiepilogoOrdine : Form
    {
        public FormRiepilogoOrdine()
        {
            InitializeComponent();
        }

        public class DatiEpilogoOrdine
        {
            // Proprietà della classe corrispondenti ai campi del JSON
            public string tracciamento { get; set; }
            public string evento_aggiunto { get; set; }
            public string prezzo { get; set; }


        }

        private void FormRiepilogoOrdine_Load(object sender, EventArgs e)
        {
            CaricaDatiEpilogoOrdine();

        }


        private async void CaricaDatiEpilogoOrdine()
        {
            using (HttpClient client = new HttpClient())
            {
                string url = "http://localhost:8082/riepilogo_ordine";

                try
                {
                    HttpResponseMessage response = await client.GetAsync(url);

                    if (response.IsSuccessStatusCode)
                    {
                        string responseData = await response.Content.ReadAsStringAsync();
                        DatiEpilogoOrdine datiEpilogo = JsonSerializer.Deserialize<DatiEpilogoOrdine>(responseData);

                        label1.Text = datiEpilogo.tracciamento;
                        label2.Text = datiEpilogo.evento_aggiunto;
                        label3.Text = datiEpilogo.prezzo;

                    }
                    else
                    {
                        Console.WriteLine($"Errore: {response.StatusCode}");
                    }
                }
                catch (Exception ex)
                {
                    // Gestisci eventuali eccezioni durante la richiesta
                    Console.WriteLine($"Errore: {ex.Message}");
                }
            }
        }


    }
}
