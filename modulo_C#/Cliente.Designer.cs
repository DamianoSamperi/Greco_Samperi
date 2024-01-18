namespace Modulo_C_
{
    partial class Cliente
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
            btn_visualizza = new Button();
            btn_inserisci = new Button();
            label1 = new Label();
            SuspendLayout();
            // 
            // btn_visualizza
            // 
            btn_visualizza.Location = new Point(290, 177);
            btn_visualizza.Name = "btn_visualizza";
            btn_visualizza.Size = new Size(162, 29);
            btn_visualizza.TabIndex = 5;
            btn_visualizza.Text = "Visualizza Spedizione";
            btn_visualizza.UseVisualStyleBackColor = true;
            btn_visualizza.Click += btn_visualizza_Click;
            // 
            // btn_inserisci
            // 
            btn_inserisci.Location = new Point(290, 89);
            btn_inserisci.Name = "btn_inserisci";
            btn_inserisci.Size = new Size(162, 29);
            btn_inserisci.TabIndex = 4;
            btn_inserisci.Text = "Crea spedizione";
            btn_inserisci.UseVisualStyleBackColor = true;
            btn_inserisci.Click += btn_inserisci_Click;
            // 
            // label1
            // 
            label1.AutoSize = true;
            label1.Location = new Point(302, 9);
            label1.Name = "label1";
            label1.Size = new Size(150, 25);
            label1.TabIndex = 3;
            label1.Text = "Seleziona operazione";
            label1.UseCompatibleTextRendering = true;
            // 
            // Cliente
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(btn_visualizza);
            Controls.Add(btn_inserisci);
            Controls.Add(label1);
            Name = "Cliente";
            Text = "Cliente";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Button btn_visualizza;
        private Button btn_inserisci;
        private Label label1;
    }
}