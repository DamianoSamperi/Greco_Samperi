using System.Windows.Forms;

namespace Modulo_C_
{
    public partial class Form1 : Form
    {
        public Form1()
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
            Form newForm = new Form2();
            //newForm.Controls.Add(customControl);
            newForm.ShowDialog();
            newForm = null;
            this.Show();
        }

        private void btn_corriere_Click(object sender, EventArgs e)
        {
            this.Hide();
            //CustomControl customControl = new CustomControl();
            Form newForm = new Form3();
            //newForm.Controls.Add(customControl);
            newForm.ShowDialog();
            newForm = null;
            this.Show();
        }
    }
}
