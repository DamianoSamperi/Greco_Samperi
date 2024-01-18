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
            tb_Provincia_Mittente = new TextBox();
            label_provincia_mittente = new Label();
            tb_Cod_Postale_Mittente = new TextBox();
            tb_Città_Mittente = new TextBox();
            label_CP_mittente = new Label();
            label_città_mittente = new Label();
            label_via_mittente = new Label();
            label_dati_mittente = new Label();
            tb_Via_Mittente = new TextBox();
            SuspendLayout();
            // 
            // LabelInserisciNome
            // 
            LabelInserisciNome.AutoSize = true;
            LabelInserisciNome.Location = new Point(56, 60);
            LabelInserisciNome.Name = "LabelInserisciNome";
            LabelInserisciNome.Size = new Size(141, 20);
            LabelInserisciNome.TabIndex = 0;
            LabelInserisciNome.Text = "Inserisci il tuo nome";
            // 
            // labelInserisciCognome
            // 
            labelInserisciCognome.AutoSize = true;
            labelInserisciCognome.Location = new Point(56, 160);
            labelInserisciCognome.Name = "labelInserisciCognome";
            labelInserisciCognome.Size = new Size(166, 20);
            labelInserisciCognome.TabIndex = 2;
            labelInserisciCognome.Text = "Inserisci il tuo cognome";
            // 
            // button_invia
            // 
            button_invia.Location = new Point(99, 409);
            button_invia.Name = "button_invia";
            button_invia.Size = new Size(94, 29);
            button_invia.TabIndex = 4;
            button_invia.Text = "invia";
            button_invia.UseVisualStyleBackColor = true;
            button_invia.Click += button_invia_Click;
            // 
            // textBoxCognome
            // 
            textBoxCognome.Location = new Point(56, 195);
            textBoxCognome.Name = "textBoxCognome";
            textBoxCognome.Size = new Size(184, 27);
            textBoxCognome.TabIndex = 5;
            // 
            // textBoxNome
            // 
            textBoxNome.Location = new Point(56, 107);
            textBoxNome.Name = "textBoxNome";
            textBoxNome.Size = new Size(184, 27);
            textBoxNome.TabIndex = 6;
            // 
            // tb_Provincia_Mittente
            // 
            tb_Provincia_Mittente.Location = new Point(666, 298);
            tb_Provincia_Mittente.MaxLength = 2;
            tb_Provincia_Mittente.Name = "tb_Provincia_Mittente";
            tb_Provincia_Mittente.Size = new Size(42, 27);
            tb_Provincia_Mittente.TabIndex = 28;
            // 
            // label_provincia_mittente
            // 
            label_provincia_mittente.AutoSize = true;
            label_provincia_mittente.Location = new Point(591, 302);
            label_provincia_mittente.Name = "label_provincia_mittente";
            label_provincia_mittente.Size = new Size(69, 20);
            label_provincia_mittente.TabIndex = 27;
            label_provincia_mittente.Text = "Provincia";
            // 
            // tb_Cod_Postale_Mittente
            // 
            tb_Cod_Postale_Mittente.Location = new Point(515, 295);
            tb_Cod_Postale_Mittente.MaxLength = 5;
            tb_Cod_Postale_Mittente.Name = "tb_Cod_Postale_Mittente";
            tb_Cod_Postale_Mittente.Size = new Size(57, 27);
            tb_Cod_Postale_Mittente.TabIndex = 26;
            // 
            // tb_Città_Mittente
            // 
            tb_Città_Mittente.Location = new Point(359, 295);
            tb_Città_Mittente.Name = "tb_Città_Mittente";
            tb_Città_Mittente.Size = new Size(97, 27);
            tb_Città_Mittente.TabIndex = 25;
            // 
            // label_CP_mittente
            // 
            label_CP_mittente.AutoSize = true;
            label_CP_mittente.Location = new Point(483, 298);
            label_CP_mittente.Name = "label_CP_mittente";
            label_CP_mittente.Size = new Size(26, 20);
            label_CP_mittente.TabIndex = 24;
            label_CP_mittente.Text = "CP";
            // 
            // label_città_mittente
            // 
            label_città_mittente.AutoSize = true;
            label_città_mittente.Location = new Point(315, 298);
            label_città_mittente.Name = "label_città_mittente";
            label_città_mittente.Size = new Size(38, 20);
            label_città_mittente.TabIndex = 23;
            label_città_mittente.Text = "città";
            // 
            // label_via_mittente
            // 
            label_via_mittente.AutoSize = true;
            label_via_mittente.Location = new Point(56, 298);
            label_via_mittente.Name = "label_via_mittente";
            label_via_mittente.Size = new Size(30, 20);
            label_via_mittente.TabIndex = 22;
            label_via_mittente.Text = "Via";
            // 
            // label_dati_mittente
            // 
            label_dati_mittente.AutoSize = true;
            label_dati_mittente.Location = new Point(56, 261);
            label_dati_mittente.Name = "label_dati_mittente";
            label_dati_mittente.Size = new Size(184, 20);
            label_dati_mittente.TabIndex = 21;
            label_dati_mittente.Text = "Inserisci i dati del mittente";
            // 
            // tb_Via_Mittente
            // 
            tb_Via_Mittente.Location = new Point(101, 295);
            tb_Via_Mittente.Name = "tb_Via_Mittente";
            tb_Via_Mittente.PlaceholderText = "Inserisci indirizzo";
            tb_Via_Mittente.Size = new Size(184, 27);
            tb_Via_Mittente.TabIndex = 20;
            // 
            // Dati_Cliente
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(tb_Provincia_Mittente);
            Controls.Add(label_provincia_mittente);
            Controls.Add(tb_Cod_Postale_Mittente);
            Controls.Add(tb_Città_Mittente);
            Controls.Add(label_CP_mittente);
            Controls.Add(label_città_mittente);
            Controls.Add(label_via_mittente);
            Controls.Add(label_dati_mittente);
            Controls.Add(tb_Via_Mittente);
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
        private TextBox tb_Provincia_Mittente;
        private Label label_provincia_mittente;
        private TextBox tb_Cod_Postale_Mittente;
        private TextBox tb_Città_Mittente;
        private Label label_CP_mittente;
        private Label label_città_mittente;
        private Label label_via_mittente;
        private Label label_dati_mittente;
        private TextBox tb_Via_Mittente;
    }
}