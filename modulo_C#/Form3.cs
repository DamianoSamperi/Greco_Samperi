using Microsoft.VisualBasic;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Modulo_C_
{
    public partial class Form3 : Form
    {
        public Form3()
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
                    HttpClient httpClient = new HttpClient();

                    HttpContent stringContent = new StringContent(sede, Encoding.UTF8, "application/json");
                    HttpResponseMessage response = await httpClient.PostAsync("http://localhost:8080/Ottieni_Percorso", stringContent);
                    var contents = await response.Content.ReadAsStringAsync();
                    MessageBox.Show(contents);
                }
                catch (Exception)
                {
                    MessageBox.Show("Connessione al server rifiutata");
                    throw;
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
            newForm = null;
            this.Show();
        }
    }
}
