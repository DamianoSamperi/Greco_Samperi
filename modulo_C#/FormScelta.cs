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
            //CustomControl customControl = new CustomControl();
            Form newForm = new Dati_Cliente();
            //newForm.Controls.Add(customControl);
            newForm.ShowDialog();
            newForm = null;
           
            this.Show();
        }

        private void btn_corriere_Click(object sender, EventArgs e)
        {
            this.Hide();
            //CustomControl customControl = new CustomControl();
            Form newForm = new FormCorriere();
            //newForm.Controls.Add(customControl);
            newForm.ShowDialog();
            newForm = null;
            this.Show();
        }
    }
}
