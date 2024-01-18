namespace Modulo_C_
{
    partial class Dati_Cliente
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            LabelInserisciNome = new Label();
            labelInserisciCognome = new Label();
            button_invia = new Button();
            textBoxCognome = new TextBox();
            textBoxNome = new TextBox();
            SuspendLayout();
            // 
            // LabelInserisciNome
            // 
            LabelInserisciNome.AutoSize = true;
            LabelInserisciNome.Location = new Point(99, 57);
            LabelInserisciNome.Name = "LabelInserisciNome";
            LabelInserisciNome.Size = new Size(141, 20);
            LabelInserisciNome.TabIndex = 0;
            LabelInserisciNome.Text = "Inserisci il tuo nome";
            // 
            // labelInserisciCognome
            // 
            labelInserisciCognome.AutoSize = true;
            labelInserisciCognome.Location = new Point(99, 157);
            labelInserisciCognome.Name = "labelInserisciCognome";
            labelInserisciCognome.Size = new Size(166, 20);
            labelInserisciCognome.TabIndex = 2;
            labelInserisciCognome.Text = "Inserisci il tuo cognome";
            // 
            // button_invia
            // 
            button_invia.Location = new Point(99, 249);
            button_invia.Name = "button_invia";
            button_invia.Size = new Size(94, 29);
            button_invia.TabIndex = 4;
            button_invia.Text = "invia";
            button_invia.UseVisualStyleBackColor = true;
            button_invia.Click += button_invia_Click;
            // 
            // textBoxCognome
            // 
            textBoxCognome.Location = new Point(99, 192);
            textBoxCognome.Name = "textBoxCognome";
            textBoxCognome.Size = new Size(184, 27);
            textBoxCognome.TabIndex = 5;
            // 
            // textBoxNome
            // 
            textBoxNome.Location = new Point(99, 104);
            textBoxNome.Name = "textBoxNome";
            textBoxNome.Size = new Size(184, 27);
            textBoxNome.TabIndex = 6;
            // 
            // Dati_Cliente
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(textBoxNome);
            Controls.Add(textBoxCognome);
            Controls.Add(button_invia);
            Controls.Add(labelInserisciCognome);
            Controls.Add(LabelInserisciNome);
            Name = "Dati_Cliente";
            Text = "Dati Cliente";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Label LabelInserisciNome;
        private Label labelInserisciCognome;
        private Button button_invia;
        private TextBox textBoxCognome;
        private TextBox textBoxNome;
    }
}