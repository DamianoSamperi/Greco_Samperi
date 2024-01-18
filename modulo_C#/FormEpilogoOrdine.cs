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
    public partial class FormEpilogoOrdine : Form
    {
        public FormEpilogoOrdine()
        {
            InitializeComponent();
        }

        public class DatiEpilogoOrdine
        {
            // Proprietà della classe corrispondenti ai campi del JSON
            public string tracciamento { get; set; }
            public string evento_aggiunto { get; set; }

            // Aggiungi altre proprietà se necessario
        }

        private void FormEpilogoOrdine_Load(object sender, EventArgs e)
        {
            CaricaDatiEpilogoOrdine();

        }


        private async void CaricaDatiEpilogoOrdine()
        {
            using (HttpClient client = new HttpClient())
            {
                string url = "http://localhost:8082/epilogo_ordine";

                try
                {
                    HttpResponseMessage response = await client.GetAsync(url);

                    if (response.IsSuccessStatusCode)
                    {
                        string responseData = await response.Content.ReadAsStringAsync();
                        // Parsa i dati JSON ricevuti
                        // Supponiamo che i dati siano in formato JSON
                        // Puoi usare un deserializzatore JSON come Newtonsoft.Json.JsonConvert
                        // In alternativa, puoi utilizzare il namespace System.Text.Json in C# (introdotto in .NET Core 3.0)
                        // a seconda della versione di .NET che stai utilizzando.

                        // Esempio usando Newtonsoft.Json.JsonConvert
                        DatiEpilogoOrdine datiEpilogo = JsonSerializer.Deserialize<DatiEpilogoOrdine>(responseData);

                        // Ora puoi utilizzare i datiEpilogo per popolare il tuo form in C#
                        // Ad esempio, puoi assegnare i dati a vari controlli del form.
                        label1.Text = datiEpilogo.tracciamento;
                        label2.Text = datiEpilogo.evento_aggiunto;

                    }
                    else
                    {
                        // Gestisci il caso in cui la richiesta non ha avuto successo
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
