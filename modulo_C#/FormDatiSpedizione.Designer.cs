namespace Modulo_C_
{
    partial class FormDatiSpedizione
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
            button_invia_locazione = new Button();
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
            // button_invia_locazione
            // 
            button_invia_locazione.Location = new Point(89, 171);
            button_invia_locazione.Name = "button_invia_locazione";
            button_invia_locazione.Size = new Size(94, 29);
            button_invia_locazione.TabIndex = 9;
            button_invia_locazione.Text = "invia";
            button_invia_locazione.UseVisualStyleBackColor = true;
            button_invia_locazione.Click += button_invia_locazione_Click;
            // 
            // tb_Provincia_Destinatario
            // 
            tb_Provincia_Destinatario.Location = new Point(654, 78);
            tb_Provincia_Destinatario.MaxLength = 2;
            tb_Provincia_Destinatario.Name = "tb_Provincia_Destinatario";
            tb_Provincia_Destinatario.Size = new Size(42, 27);
            tb_Provincia_Destinatario.TabIndex = 28;
            // 
            // label1
            // 
            label1.AutoSize = true;
            label1.Location = new Point(579, 82);
            label1.Name = "label1";
            label1.Size = new Size(69, 20);
            label1.TabIndex = 27;
            label1.Text = "Provincia";
            // 
            // tb_Cod_Postale_Destinatario
            // 
            tb_Cod_Postale_Destinatario.Location = new Point(503, 75);
            tb_Cod_Postale_Destinatario.MaxLength = 5;
            tb_Cod_Postale_Destinatario.Name = "tb_Cod_Postale_Destinatario";
            tb_Cod_Postale_Destinatario.Size = new Size(57, 27);
            tb_Cod_Postale_Destinatario.TabIndex = 26;
            // 
            // tb_Città_Destinatario
            // 
            tb_Città_Destinatario.Location = new Point(347, 75);
            tb_Città_Destinatario.Name = "tb_Città_Destinatario";
            tb_Città_Destinatario.Size = new Size(97, 27);
            tb_Città_Destinatario.TabIndex = 25;
            // 
            // label2
            // 
            label2.AutoSize = true;
            label2.Location = new Point(471, 78);
            label2.Name = "label2";
            label2.Size = new Size(26, 20);
            label2.TabIndex = 24;
            label2.Text = "CP";
            // 
            // label3
            // 
            label3.AutoSize = true;
            label3.Location = new Point(303, 78);
            label3.Name = "label3";
            label3.Size = new Size(38, 20);
            label3.TabIndex = 23;
            label3.Text = "città";
            // 
            // label4
            // 
            label4.AutoSize = true;
            label4.Location = new Point(44, 78);
            label4.Name = "label4";
            label4.Size = new Size(30, 20);
            label4.TabIndex = 22;
            label4.Text = "Via";
            // 
            // label_dati_destinatario
            // 
            label_dati_destinatario.AutoSize = true;
            label_dati_destinatario.Location = new Point(12, 49);
            label_dati_destinatario.Name = "label_dati_destinatario";
            label_dati_destinatario.Size = new Size(207, 20);
            label_dati_destinatario.TabIndex = 21;
            label_dati_destinatario.Text = "Inserisci i dati del destinatario";
            // 
            // tb_Via_Destinatario
            // 
            tb_Via_Destinatario.Location = new Point(89, 75);
            tb_Via_Destinatario.Name = "tb_Via_Destinatario";
            tb_Via_Destinatario.PlaceholderText = "Inserisci Indirizzo";
            tb_Via_Destinatario.Size = new Size(184, 27);
            tb_Via_Destinatario.TabIndex = 20;
            // 
            // FormDatiSpedizione
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
            Controls.Add(button_invia_locazione);
            Name = "FormDatiSpedizione";
            Text = "Dati Spedizione";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion
        private Button button_invia_locazione;
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