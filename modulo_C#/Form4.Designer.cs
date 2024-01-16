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
            textBoxMittente = new TextBox();
            textBoxDestinatario = new TextBox();
            button_invia_locazione = new Button();
            labelInserisciDestinatario = new Label();
            LabelInserisciMittente = new Label();
            SuspendLayout();
            // 
            // textBoxMittente
            // 
            textBoxMittente.Location = new Point(136, 115);
            textBoxMittente.Name = "textBoxMittente";
            textBoxMittente.Size = new Size(184, 27);
            textBoxMittente.TabIndex = 11;
            // 
            // textBoxDestinatario
            // 
            textBoxDestinatario.Location = new Point(136, 203);
            textBoxDestinatario.Name = "textBoxDestinatario";
            textBoxDestinatario.Size = new Size(184, 27);
            textBoxDestinatario.TabIndex = 10;
            // 
            // button_invia_locazione
            // 
            button_invia_locazione.Location = new Point(136, 260);
            button_invia_locazione.Name = "button_invia_locazione";
            button_invia_locazione.Size = new Size(94, 29);
            button_invia_locazione.TabIndex = 9;
            button_invia_locazione.Text = "invia";
            button_invia_locazione.UseVisualStyleBackColor = true;
            button_invia_locazione.Click += button_invia_locazione_Click;
            // 
            // labelInserisciDestinatario
            // 
            labelInserisciDestinatario.AutoSize = true;
            labelInserisciDestinatario.Location = new Point(136, 168);
            labelInserisciDestinatario.Name = "labelInserisciDestinatario";
            labelInserisciDestinatario.Size = new Size(205, 20);
            labelInserisciDestinatario.TabIndex = 8;
            labelInserisciDestinatario.Text = "Inserisci indirizzo destinatario";
            // 
            // LabelInserisciMittente
            // 
            LabelInserisciMittente.AutoSize = true;
            LabelInserisciMittente.Location = new Point(136, 68);
            LabelInserisciMittente.Name = "LabelInserisciMittente";
            LabelInserisciMittente.Size = new Size(182, 20);
            LabelInserisciMittente.TabIndex = 7;
            LabelInserisciMittente.Text = "Inserisci indirizzo mittente";
            // 
            // Form4
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(textBoxMittente);
            Controls.Add(textBoxDestinatario);
            Controls.Add(button_invia_locazione);
            Controls.Add(labelInserisciDestinatario);
            Controls.Add(LabelInserisciMittente);
            Name = "Form4";
            Text = "Form4";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private TextBox textBoxMittente;
        private TextBox textBoxDestinatario;
        private Button button_invia_locazione;
        private Label labelInserisciDestinatario;
        private Label LabelInserisciMittente;
    }
}