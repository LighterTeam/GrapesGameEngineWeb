using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.IO;

namespace MessageClient
{
    class Program
    {
        static SocketClient client;

        static void Main(string[] args)
        {
            client = new SocketClient();
            client.BinaryInput = new ClientBinaryInputHandler(ClientBinaryInputHandler); //设置数据包处理回调方法
            client.MessageInput = new ClientMessageInputHandler(ClientMessageInputHandler);//断开处理
            if (client.Connect(Config.Default.IP, Config.Default.Port)) //连接到服务器
            {             
                client.StartRead(); //开始监听读取
                while (true)
                {
                    string mess = Console.ReadLine();
                    if (mess.Length == 0)
                        continue;
                    byte[] messdata = Buffers.GetSocketBytes(mess);
                    messdata = Buffers.MergeBytes(Buffers.GetSocketBytes(messdata.Length), messdata);
                    client.SendData(messdata);
                }
            }
            else
            {
                Console.WriteLine("无法连接服务器");
            }
        }

        static void ClientBinaryInputHandler(byte[] data)
        {
            Console.WriteLine(Encoding.UTF8.GetString(data, 0, data.Length));
        }

        static void ClientMessageInputHandler(string message)
        {
            Console.WriteLine("与服务器端断开连接");
            client.Close();
        }
    }
}
