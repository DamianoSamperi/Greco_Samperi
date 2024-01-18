namespace Modulo_C_
{
    partial class Form4
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
            tb_Via_Mittente = new TextBox();
            button_invia_locazione = new Button();
            label_dati_mittente = new Label();
            label_via_mittente = new Label();
            label_città_mittente = new Label();
            label_CP_mittente = new Label();
            tb_Città_Mittente = new TextBox();
            tb_Cod_Postale_Mittente = new TextBox();
            label_provincia_mittente = new Label();
            tb_Provincia_Mittente = new TextBox();
            tb_Provincia_Destinatario = new TextBox();
            label1 = new Label();
            tb_Cod_Postale_Destinatario = new TextBox();
            tb_Città_Destinatario = new TextBox();
            label2 = new Label();
            label3 = new Label();
            label4 = new Label();
            label_dati_destinatario = new Label();
            tb_Via_Destinatario = new TextBox();
            SuspendLayout();
            // 
            // tb_Via_Mittente
            // 
            tb_Via_Mittente.Location = new Point(89, 115);
            tb_Via_Mittente.Name = "tb_Via_Mittente";
            tb_Via_Mittente.PlaceholderText = "Inserisci indirizzo";
            tb_Via_Mittente.Size = new Size(184, 27);
            tb_Via_Mittente.TabIndex = 11;
            tb_Via_Mittente.TextChanged += textBoxMittente_TextChanged;
            // 
            // button_invia_locazione
            // 
            button_invia_locazione.Location = new Point(89, 318);
            button_invia_locazione.Name = "button_invia_locazione";
            button_invia_locazione.Size = new Size(94, 29);
            button_invia_locazione.TabIndex = 9;
            button_invia_locazione.Text = "invia";
            button_invia_locazione.UseVisualStyleBackColor = true;
            button_invia_locazione.Click += button_invia_locazione_Click;
            // 
            // label_dati_mittente
            // 
            label_dati_mittente.AutoSize = true;
            label_dati_mittente.Location = new Point(12, 89);
            label_dati_mittente.Name = "label_dati_mittente";
            label_dati_mittente.Size = new Size(184, 20);
            label_dati_mittente.TabIndex = 12;
            label_dati_mittente.Text = "Inserisci i dati del mittente";
            // 
            // label_via_mittente
            // 
            label_via_mittente.AutoSize = true;
            label_via_mittente.Location = new Point(44, 118);
            label_via_mittente.Name = "label_via_mittente";
            label_via_mittente.Size = new Size(30, 20);
            label_via_mittente.TabIndex = 13;
            label_via_mittente.Text = "Via";
            label_via_mittente.Click += label2_Click;
            // 
            // label_città_mittente
            // 
            label_città_mittente.AutoSize = true;
            label_città_mittente.Location = new Point(303, 118);
            label_città_mittente.Name = "label_città_mittente";
            label_città_mittente.Size = new Size(38, 20);
            label_città_mittente.TabIndex = 14;
            label_città_mittente.Text = "città";
            // 
            // label_CP_mittente
            // 
            label_CP_mittente.AutoSize = true;
            label_CP_mittente.Location = new Point(471, 118);
            label_CP_mittente.Name = "label_CP_mittente";
            label_CP_mittente.Size = new Size(26, 20);
            label_CP_mittente.TabIndex = 15;
            label_CP_mittente.Text = "CP";
            // 
            // tb_Città_Mittente
            // 
            tb_Città_Mittente.Location = new Point(347, 115);
            tb_Città_Mittente.Name = "tb_Città_Mittente";
            tb_Città_Mittente.Size = new Size(97, 27);
            tb_Città_Mittente.TabIndex = 16;
            tb_Città_Mittente.TextChanged += textBox1_TextChanged;
            // 
            // tb_Cod_Postale_Mittente
            // 
            tb_Cod_Postale_Mittente.Location = new Point(503, 115);
            tb_Cod_Postale_Mittente.MaxLength = 5;
            tb_Cod_Postale_Mittente.Name = "tb_Cod_Postale_Mittente";
            tb_Cod_Postale_Mittente.Size = new Size(57, 27);
            tb_Cod_Postale_Mittente.TabIndex = 17;
            // 
            // label_provincia_mittente
            // 
            label_provincia_mittente.AutoSize = true;
            label_provincia_mittente.Location = new Point(579, 122);
            label_provincia_mittente.Name = "label_provincia_mittente";
            label_provincia_mittente.Size = new Size(69, 20);
            label_provincia_mittente.TabIndex = 18;
            label_provincia_mittente.Text = "Provincia";
            // 
            // tb_Provincia_Mittente
            // 
            tb_Provincia_Mittente.Location = new Point(654, 118);
            tb_Provincia_Mittente.MaxLength = 2;
            tb_Provincia_Mittente.Name = "tb_Provincia_Mittente";
            tb_Provincia_Mittente.Size = new Size(42, 27);
            tb_Provincia_Mittente.TabIndex = 19;
            // 
            // tb_Provincia_Destinatario
            // 
            tb_Provincia_Destinatario.Location = new Point(654, 225);
            tb_Provincia_Destinatario.MaxLength = 2;
            tb_Provincia_Destinatario.Name = "tb_Provincia_Destinatario";
            tb_Provincia_Destinatario.Size = new Size(42, 27);
            tb_Provincia_Destinatario.TabIndex = 28;
            // 
            // label1
            // 
            label1.AutoSize = true;
            label1.Location = new Point(579, 229);
            label1.Name = "label1";
            label1.Size = new Size(69, 20);
            label1.TabIndex = 27;
            label1.Text = "Provincia";
            // 
            // tb_Cod_Postale_Destinatario
            // 
            tb_Cod_Postale_Destinatario.Location = new Point(503, 222);
            tb_Cod_Postale_Destinatario.MaxLength = 5;
            tb_Cod_Postale_Destinatario.Name = "tb_Cod_Postale_Destinatario";
            tb_Cod_Postale_Destinatario.Size = new Size(57, 27);
            tb_Cod_Postale_Destinatario.TabIndex = 26;
            // 
            // tb_Città_Destinatario
            // 
            tb_Città_Destinatario.Location = new Point(347, 222);
            tb_Città_Destinatario.Name = "tb_Città_Destinatario";
            tb_Città_Destinatario.Size = new Size(97, 27);
            tb_Città_Destinatario.TabIndex = 25;
            // 
            // label2
            // 
            label2.AutoSize = true;
            label2.Location = new Point(471, 225);
            label2.Name = "label2";
            label2.Size = new Size(26, 20);
            label2.TabIndex = 24;
            label2.Text = "CP";
            // 
            // label3
            // 
            label3.AutoSize = true;
            label3.Location = new Point(303, 225);
            label3.Name = "label3";
            label3.Size = new Size(38, 20);
            label3.TabIndex = 23;
            label3.Text = "città";
            // 
            // label4
            // 
            label4.AutoSize = true;
            label4.Location = new Point(44, 225);
            label4.Name = "label4";
            label4.Size = new Size(30, 20);
            label4.TabIndex = 22;
            label4.Text = "Via";
            // 
            // label_dati_destinatario
            // 
            label_dati_destinatario.AutoSize = true;
            label_dati_destinatario.Location = new Point(12, 196);
            label_dati_destinatario.Name = "label_dati_destinatario";
            label_dati_destinatario.Size = new Size(207, 20);
            label_dati_destinatario.TabIndex = 21;
            label_dati_destinatario.Text = "Inserisci i dati del destinatario";
            // 
            // tb_Via_Destinatario
            // 
            tb_Via_Destinatario.Location = new Point(89, 222);
            tb_Via_Destinatario.Name = "tb_Via_Destinatario";
            tb_Via_Destinatario.PlaceholderText = "Inserisci Indirizzo";
            tb_Via_Destinatario.Size = new Size(184, 27);
            tb_Via_Destinatario.TabIndex = 20;
            // 
            // Form4
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(tb_Provincia_Destinatario);
            Controls.Add(label1);
            Controls.Add(tb_Cod_Postale_Destinatario);
            Controls.Add(tb_Città_Destinatario);
            Controls.Add(label2);
            Controls.Add(label3);
            Controls.Add(label4);
            Controls.Add(label_dati_destinatario);
            Controls.Add(tb_Via_Destinatario);
            Controls.Add(tb_Provincia_Mittente);
            Controls.Add(label_provincia_mittente);
            Controls.Add(tb_Cod_Postale_Mittente);
            Controls.Add(tb_Città_Mittente);
            Controls.Add(label_CP_mittente);
            Controls.Add(label_città_mittente);
            Controls.Add(label_via_mittente);
            Controls.Add(label_dati_mittente);
            Controls.Add(tb_Via_Mittente);
            Controls.Add(button_invia_locazione);
            Name = "Form4";
            Text = "Form4";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private TextBox tb_Via_Mittente;
        private Button button_invia_locazione;
        private Label label_dati_mittente;
        private Label label_via_mittente;
        private Label label_città_mittente;
        private Label label_CP_mittente;
        private TextBox tb_Città_Mittente;
        private TextBox tb_Cod_Postale_Mittente;
        private Label label_provincia_mittente;
        private TextBox tb_Provincia_Mittente;
        private TextBox tb_Provincia_Destinatario;
        private Label label1;
        private TextBox tb_Cod_Postale_Destinatario;
        private TextBox tb_Città_Destinatario;
        private Label label2;
        private Label label3;
        private Label label4;
        private Label label_dati_destinatario;
        private TextBox tb_Via_Destinatario;
    }
}