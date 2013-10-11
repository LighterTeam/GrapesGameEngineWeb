using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;
using System.Text;
using System.Diagnostics;

namespace GameClient
{
    static class Program
    {
        public static SocketClient client;

        /// <summary>
        /// The main entry point for the application.
        /// </summary>
        [STAThread]
        static void Main()
        {
            client = new SocketClient();
            client.BinaryInput = new ClientBinaryInputHandler(ClientBinaryInputHandler); //设置数据包处理回调方法
            client.MessageInput = new ClientMessageInputHandler(ClientMessageInputHandler);//断开处理
            if (client.Connect("127.0.0.1", 9901)) //连接到服务器
            {
                client.StartRead(); //开始监听读取

                Application.EnableVisualStyles();
                Application.SetCompatibleTextRenderingDefault(false);
                Application.Run(new Form1());
            }
        }

        static void ClientBinaryInputHandler(byte[] data)
        {
            var csData = Encoding.UTF8.GetString(data, 0, data.Length);
            Console.WriteLine("Recv:" + csData);
            
        }

        static void ClientMessageInputHandler(string message)
        {
            Console.WriteLine("与服务器端断开连接");
            client.Close();
        }
    }
}
