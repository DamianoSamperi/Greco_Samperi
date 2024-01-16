namespace Modulo_C_
{
    partial class Form3
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
            label1 = new Label();
            btn_percorso = new Button();
            SuspendLayout();
            // 
            // label1
            // 
            label1.AutoSize = true;
            label1.Location = new Point(326, 20);
            label1.Name = "label1";
            label1.Size = new Size(150, 25);
            label1.TabIndex = 0;
            label1.Text = "Seleziona operazione";
            label1.UseCompatibleTextRendering = true;
            label1.Click += label1_Click;
            // 
            // btn_percorso
            // 
            btn_percorso.Location = new Point(314, 101);
            btn_percorso.Name = "btn_percorso";
            btn_percorso.Size = new Size(162, 29);
            btn_percorso.TabIndex = 1;
            btn_percorso.Text = "Ottieni spedizioni";
            btn_percorso.UseVisualStyleBackColor = true;
            btn_percorso.Click += btn_percorso_Click;
            // 
            // Form3
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(btn_percorso);
            Controls.Add(label1);
            Name = "Form3";
            Text = "Corriere";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Label label1;
        private Button btn_percorso;
    }
}