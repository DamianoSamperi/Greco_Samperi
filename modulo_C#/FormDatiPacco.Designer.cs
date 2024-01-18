namespace Modulo_C_
{
    partial class FormDatiPacco
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
            textBoxPeso = new TextBox();
            textBoxDimensione = new TextBox();
            button_invia_dati_pacco = new Button();
            labelInserisciPeso = new Label();
            LabelInserisciDimensione = new Label();
            buttonFineOrdine = new Button();
            SuspendLayout();
            // 
            // textBoxPeso
            // 
            textBoxPeso.Location = new Point(62, 94);
            textBoxPeso.Name = "textBoxPeso";
            textBoxPeso.Size = new Size(184, 27);
            textBoxPeso.TabIndex = 11;
            // 
            // textBoxDimensione
            // 
            textBoxDimensione.Location = new Point(62, 182);
            textBoxDimensione.Name = "textBoxDimensione";
            textBoxDimensione.Size = new Size(184, 27);
            textBoxDimensione.TabIndex = 10;
            // 
            // button_invia_dati_pacco
            // 
            button_invia_dati_pacco.Location = new Point(62, 248);
            button_invia_dati_pacco.Name = "button_invia_dati_pacco";
            button_invia_dati_pacco.Size = new Size(94, 29);
            button_invia_dati_pacco.TabIndex = 9;
            button_invia_dati_pacco.Text = "invia";
            button_invia_dati_pacco.UseVisualStyleBackColor = true;
            button_invia_dati_pacco.Click += button_invia_dati_pacco_Click_1;
            // 
            // labelInserisciPeso
            // 
            labelInserisciPeso.AutoSize = true;
            labelInserisciPeso.Location = new Point(62, 47);
            labelInserisciPeso.Name = "labelInserisciPeso";
            labelInserisciPeso.Size = new Size(189, 20);
            labelInserisciPeso.TabIndex = 7;
            labelInserisciPeso.Text = "Inserisci il Peso (in grammi)";
            // 
            // LabelInserisciDimensione
            // 
            LabelInserisciDimensione.AutoSize = true;
            LabelInserisciDimensione.Location = new Point(62, 147);
            LabelInserisciDimensione.Name = "LabelInserisciDimensione";
            LabelInserisciDimensione.Size = new Size(392, 20);
            LabelInserisciDimensione.TabIndex = 8;
            LabelInserisciDimensione.Text = "Inserisci la dimensione del pacco (piccolo/medio/grande)";
            // 
            // buttonFineOrdine
            // 
            buttonFineOrdine.Location = new Point(208, 248);
            buttonFineOrdine.Name = "buttonFineOrdine";
            buttonFineOrdine.Size = new Size(94, 29);
            buttonFineOrdine.TabIndex = 12;
            buttonFineOrdine.Text = "fine ordine";
            buttonFineOrdine.UseVisualStyleBackColor = true;
            buttonFineOrdine.Click += buttonFineOrdine_Click;
            // 
            // FormDatiPacco
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(buttonFineOrdine);
            Controls.Add(textBoxPeso);
            Controls.Add(textBoxDimensione);
            Controls.Add(button_invia_dati_pacco);
            Controls.Add(labelInserisciPeso);
            Controls.Add(LabelInserisciDimensione);
            Name = "FormDatiPacco";
            Text = "Dati Pacco";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private TextBox textBoxPeso;
        private TextBox textBoxDimensione;
        private Button button_invia_dati_pacco;
        private Label labelInserisciPeso;
        private Label LabelInserisciDimensione;
        private Button buttonFineOrdine;
    }
}