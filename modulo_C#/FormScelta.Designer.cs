namespace Modulo_C_
{
    partial class FormScelta
    {
        /// <summary>
        ///  Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        ///  Clean up any resources being used.
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
        ///  Required method for Designer support - do not modify
        ///  the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            btn_cliente = new Button();
            btn_corriere = new Button();
            label1 = new Label();
            SuspendLayout();
            // 
            // btn_cliente
            // 
            btn_cliente.Location = new Point(187, 180);
            btn_cliente.Name = "btn_cliente";
            btn_cliente.Size = new Size(128, 59);
            btn_cliente.TabIndex = 0;
            btn_cliente.Text = "Cliente";
            btn_cliente.UseVisualStyleBackColor = true;
            btn_cliente.Click += btn_cliente_Click;
            // 
            // btn_corriere
            // 
            btn_corriere.Location = new Point(407, 180);
            btn_corriere.Name = "btn_corriere";
            btn_corriere.Size = new Size(128, 59);
            btn_corriere.TabIndex = 1;
            btn_corriere.Text = "Corriere";
            btn_corriere.UseVisualStyleBackColor = true;
            btn_corriere.Click += btn_corriere_Click;
            // 
            // label1
            // 
            label1.AutoSize = true;
            label1.Location = new Point(318, 60);
            label1.Name = "label1";
            label1.Size = new Size(80, 20);
            label1.TabIndex = 2;
            label1.Text = "Identificati";
            label1.Click += label1_Click;
            // 
            // Form1
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(label1);
            Controls.Add(btn_corriere);
            Controls.Add(btn_cliente);
            Name = "Form1";
            Text = "Form1";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Button btn_cliente;
        private Button btn_corriere;
        private Label label1;
    }
}
