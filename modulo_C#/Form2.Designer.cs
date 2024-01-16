namespace Modulo_C_
{
    partial class Form2
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
            textBoxNome = new TextBox();
            labelInserisciCognome = new Label();
            textBoxCognome = new TextBox();
            button_invia = new Button();
            SuspendLayout();
            // 
            // LabelInserisciNome
            // 
            LabelInserisciNome.AutoSize = true;
            LabelInserisciNome.Location = new Point(93, 64);
            LabelInserisciNome.Name = "LabelInserisciNome";
            LabelInserisciNome.Size = new Size(141, 20);
            LabelInserisciNome.TabIndex = 0;
            LabelInserisciNome.Text = "Inserisci il tuo nome";
            LabelInserisciNome.Click += label_Nome_Click;
            // 
            // textBoxNome
            // 
            textBoxNome.Location = new Point(93, 97);
            textBoxNome.Name = "textBoxNome";
            textBoxNome.Size = new Size(190, 27);
            textBoxNome.TabIndex = 1;
            textBoxNome.TextChanged += textBoxNome_TextChanged;
            // 
            // labelInserisciCognome
            // 
            labelInserisciCognome.AutoSize = true;
            labelInserisciCognome.Location = new Point(93, 147);
            labelInserisciCognome.Name = "labelInserisciCognome";
            labelInserisciCognome.Size = new Size(166, 20);
            labelInserisciCognome.TabIndex = 2;
            labelInserisciCognome.Text = "Inserisci il tuo cognome";
            labelInserisciCognome.Click += label_cognome;
            // 
            // textBoxCognome
            // 
            textBoxCognome.Location = new Point(93, 185);
            textBoxCognome.Name = "textBoxCognome";
            textBoxCognome.Size = new Size(190, 27);
            textBoxCognome.TabIndex = 3;
            textBoxCognome.TextChanged += textBoxCognome_TextChanged;
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
            // Form2
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(button_invia);
            Controls.Add(textBoxCognome);
            Controls.Add(labelInserisciCognome);
            Controls.Add(textBoxNome);
            Controls.Add(LabelInserisciNome);
            Name = "Form2";
            Text = "Form2";
            Load += Form2_Load;
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Label LabelInserisciNome;
        private TextBox textBoxNome;
        private Label labelInserisciCognome;
        private TextBox textBoxCognome;
        private Button button_invia;
    }
}