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
            labelContinua = new Label();
            checkedListBoxSiNo = new CheckedListBox();
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
            button_invia_dati_pacco.Location = new Point(62, 346);
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
            labelInserisciPeso.Size = new Size(107, 20);
            labelInserisciPeso.TabIndex = 7;
            labelInserisciPeso.Text = "Inserisci il Peso";
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
            // labelContinua
            // 
            labelContinua.AutoSize = true;
            labelContinua.Location = new Point(62, 245);
            labelContinua.Name = "labelContinua";
            labelContinua.Size = new Size(250, 20);
            labelContinua.TabIndex = 12;
            labelContinua.Text = "vuoi inserie un nuovo pacco?(si/no): ";
            // 
            // checkedListBoxSiNo
            // 
            checkedListBoxSiNo.FormattingEnabled = true;
            checkedListBoxSiNo.Items.AddRange(new object[] { "si", "no" });
            checkedListBoxSiNo.Location = new Point(62, 283);
            checkedListBoxSiNo.Name = "checkedListBoxSiNo";
            checkedListBoxSiNo.Size = new Size(150, 48);
            checkedListBoxSiNo.TabIndex = 15;
            // 
            // FormDatiPacco
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(checkedListBoxSiNo);
            Controls.Add(labelContinua);
            Controls.Add(textBoxPeso);
            Controls.Add(textBoxDimensione);
            Controls.Add(button_invia_dati_pacco);
            Controls.Add(labelInserisciPeso);
            Controls.Add(LabelInserisciDimensione);
            Name = "FormDatiPacco";
            Text = "FormDatiPacco";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private TextBox textBoxPeso;
        private TextBox textBoxDimensione;
        private Button button_invia_dati_pacco;
        private Label labelInserisciPeso;
        private Label LabelInserisciDimensione;
        private Label labelContinua;
        private CheckedListBox checkedListBoxSiNo;
    }
}