namespace Modulo_C_
{
    partial class Consegna
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
            btn_hub = new Button();
            btn_consegna = new Button();
            label_id = new Label();
            textBox_id = new TextBox();
            SuspendLayout();
            // 
            // btn_hub
            // 
            btn_hub.Location = new Point(218, 173);
            btn_hub.Name = "btn_hub";
            btn_hub.Size = new Size(177, 29);
            btn_hub.TabIndex = 0;
            btn_hub.Text = "Consegna all'hub";
            btn_hub.UseVisualStyleBackColor = true;
            btn_hub.Click += btn_hub_ClickAsync;
            // 
            // btn_consegna
            // 
            btn_consegna.Location = new Point(63, 173);
            btn_consegna.Name = "btn_consegna";
            btn_consegna.Size = new Size(94, 29);
            btn_consegna.TabIndex = 1;
            btn_consegna.Text = "Consegna";
            btn_consegna.UseVisualStyleBackColor = true;
            btn_consegna.Click += btn_consegna_Click;
            // 
            // label_id
            // 
            label_id.AutoSize = true;
            label_id.Location = new Point(63, 99);
            label_id.Name = "label_id";
            label_id.Size = new Size(98, 20);
            label_id.TabIndex = 2;
            label_id.Text = "Id spedizione";
            // 
            // textBox_id
            // 
            textBox_id.Location = new Point(218, 101);
            textBox_id.Name = "textBox_id";
            textBox_id.PlaceholderText = "Inserire codice spedizione";
            textBox_id.Size = new Size(177, 27);
            textBox_id.TabIndex = 3;
            // 
            // Consegna
            // 
            AutoScaleDimensions = new SizeF(8F, 20F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(textBox_id);
            Controls.Add(label_id);
            Controls.Add(btn_consegna);
            Controls.Add(btn_hub);
            Name = "Consegna";
            Text = "Consegna";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Button btn_hub;
        private Button btn_consegna;
        private Label label_id;
        private TextBox textBox_id;
    }
}