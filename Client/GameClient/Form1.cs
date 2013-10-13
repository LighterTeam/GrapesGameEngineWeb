using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.Runtime.Serialization.Json;
using System.Runtime.Serialization;
using System.IO;
using LitJson;

namespace GameClient
{
    public partial class Form1 : Form
    {
        public static TSGame game;
        public static Form1 GF;

        public Form1()
        {
            InitializeComponent();
        }

        public delegate void OutDelegate(string csData);
        public void MessageProc(string csData)
        {
            game.MessageProc(csData);
        }

        static void ClientBinaryInputHandler(byte[] data)
        {
            var csData = Encoding.UTF8.GetString(data, 0, data.Length);
            Console.WriteLine("Recv:" + csData);
            Form1.OutDelegate outdelegate = new Form1.OutDelegate(GF.MessageProc);
            GF.BeginInvoke(outdelegate, csData);
        }

        static void ClientMessageInputHandler(string message)
        {
            Console.WriteLine("与服务器端断开连接");
            Program.client.Close();
        }

        protected override bool ProcessCmdKey(ref System.Windows.Forms.Message msg, System.Windows.Forms.Keys keyData)  
        {
            if (game.KeyDown(keyData))
            {
                return true;
            }
            return base.ProcessCmdKey(ref msg, keyData);  
        }

        private void button1_Click(object sender, EventArgs e)
        {
            JsonData pjd = new JsonData();
            pjd["OPCODE"] = (Int32)GameOpcode.CreateSprite;
            pjd["UUID"] = game.SelfUUID;
            Program.SendDataSingle(pjd.ToJson(), game.SelfUUID);
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            GF = this;
            game = new TSGame(this);
            Program.client = new SocketClient();
            Program.client.BinaryInput = new ClientBinaryInputHandler(ClientBinaryInputHandler); //设置数据包处理回调方法
            Program.client.MessageInput = new ClientMessageInputHandler(ClientMessageInputHandler);//断开处理
            if (Program.client.Connect("10.211.55.4", 9901)) //连接到服务器
            {
                Program.client.StartRead(); //开始监听读取
            }
        }
    }
}
