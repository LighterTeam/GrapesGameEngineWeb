using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;
using System.Text;
using System.Diagnostics;
using System.IO;
using System.Runtime.Serialization.Json;

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
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(new Form1());
        }

        public static void SendData(MessagePacket mp)
        {
            MemoryStream stream = new MemoryStream();
            DataContractJsonSerializer ser = new DataContractJsonSerializer(typeof(MessagePacket));
            ser.WriteObject(stream, mp);
            var packet = Encoding.UTF8.GetString(stream.GetBuffer(), 0, (int)stream.Length);
            client.SendData(packet);
        }

        public static void SendDataBroad(string data)
        {
            MessagePacket mp = new MessagePacket();
            mp.D = data;
            mp.T = (int)SendMode.SendBroad;
            MemoryStream stream = new MemoryStream();
            DataContractJsonSerializer ser = new DataContractJsonSerializer(typeof(MessagePacket));
            ser.WriteObject(stream, mp);
            var packet = Encoding.UTF8.GetString(stream.GetBuffer(), 0, (int)stream.Length);
            client.SendData(packet);
        }

        public static void SendDataBroadNotSelf(string data)
        {
            MessagePacket mp = new MessagePacket();
            mp.D = data;
            mp.T = (int)SendMode.SendBroadNotSelf;
            MemoryStream stream = new MemoryStream();
            DataContractJsonSerializer ser = new DataContractJsonSerializer(typeof(MessagePacket));
            ser.WriteObject(stream, mp);
            var packet = Encoding.UTF8.GetString(stream.GetBuffer(), 0, (int)stream.Length);
            client.SendData(packet);
        }

        public static void SendDataSingle(string data, Int64 uuid)
        {
            MessagePacket mp = new MessagePacket();
            mp.D = data;
            mp.T = (int)SendMode.SendSingle;
            mp.S = uuid;
            MemoryStream stream = new MemoryStream();
            DataContractJsonSerializer ser = new DataContractJsonSerializer(typeof(MessagePacket));
            ser.WriteObject(stream, mp);
            var packet = Encoding.UTF8.GetString(stream.GetBuffer(), 0, (int)stream.Length);
            client.SendData(packet);
        }
    }
}
