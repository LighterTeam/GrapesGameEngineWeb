using System;
using System.Collections.Generic;
using System.Text;
using System.Net.Sockets;
using System.Net;

namespace MessageClient
{
    //数据包输入代理
    public delegate void ClientBinaryInputHandler(byte[] data);

    // 异常错误通常是用户断开的代理
    public delegate void ClientMessageInputHandler(string message);

    public delegate void ConnectionHandler(bool conn);

    public class SocketClient
    {
        private Socket sock;
        
        // Socket对象
        public Socket Sock { get { return sock; } }

        // 数据包长度
        public int BuffLength { get; set; }

        // 数据输入处理
        public ClientBinaryInputHandler BinaryInput { get; set; }

        // 异常错误通常是用户断开处理
        public ClientMessageInputHandler MessageInput { get; set; }

        // 有连接
        public ConnectionHandler ConnInput { get; set; }

        private SocketError socketError;

        public SocketClient()
        {
            BuffLength = 4096;
            sock = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
        }

        //连接到目标主机
        public bool Connect(string host, int prot)
        {
            try
            {
                IPEndPoint myEnd = new IPEndPoint(IPAddress.Parse(host), prot);
                sock.Connect(myEnd);
                if (sock.Connected)
                {
                    return true;
                }
                else
                {
                    return false;
                }

            }
            catch (SocketException)
            {
                return false;
            }
            catch
            {
                throw;
            }
        }
        
        //连接到目标主机
        public void BeginConnect(string host, int prot)
        {
            try
            {
                IPEndPoint myEnd = new IPEndPoint(IPAddress.Parse(host), prot);
                sock.BeginConnect(myEnd, new AsyncCallback(ConnAsyncCallBack), sock);             
            }           
            catch
            {
                throw;
            }
        }

        void ConnAsyncCallBack(IAsyncResult result)
        {
            try
            {
                sock.EndConnect(result);
                if (sock.Connected)
                {
                    if(ConnInput!=null)
                        ConnInput(true);
                }
                else
                    if (ConnInput != null)
                        ConnInput(false);
            }
            catch(Exception)
            {
                if (ConnInput != null)
                     ConnInput(false);
            }
        }

        // 开始读取数据
        public void StartRead()
        {
            BeginReceive();
        }

        void BeginReceive()
        {
            byte[] data = new byte[BuffLength];
            sock.BeginReceive(data, 0, data.Length, SocketFlags.None, out socketError, args_Completed, data);
        }

        public static byte[] buf;
        public static byte[] bu;
        public static int iLen = 0;

        public int GetDataLen(byte[] data, int index)
        {
            var len = BitConverter.ToInt32(data, index);
            var bys = BitConverter.GetBytes(len);
            bys = Buffers.ReverseBytes(bys);
            len = BitConverter.ToInt32(bys, 0); //获取到包头多少个字节
            return len;
        }

        void DataMM(byte[] data)
        {
            var index = 0;

            if (iLen > 0)
            {
                if (iLen <= data.Length)
                {
                    var tmp = new byte[iLen];
                    Array.Copy(data, 0, tmp, 0, tmp.Length);
                    buf = Buffers.MergeBytes(buf, tmp);
                    this.BinaryInput(buf);
                    buf = null;
                    index = iLen;
                }
                else
                {
                    buf = Buffers.MergeBytes(buf, data);
                    iLen -= data.Length;
                    return;
                }

                if (bu != null)
                {
                    data = Buffers.MergeBytes(bu, data);
                    bu = null;
                }
            }

            while (true)
            {
                if (data.Length - index < 4)
                {
                    bu = new byte[data.Length - index];
                    Array.Copy(data, index, bu, 0, bu.Length);
                    iLen = 0;
                    break;
                }

                iLen = GetDataLen(data, index);
                if (index + iLen + 4 > data.Length)
                {
                    //包不完整
                    buf = new byte[data.Length - index - 4];
                    Array.Copy(data, index + 4, buf, 0, buf.Length);
                    iLen -= buf.Length;
                    break;
                }
                else if (index + iLen + 4 == data.Length)
                {
                    buf = new byte[iLen];
                    Array.Copy(data, index + 4, buf, 0, buf.Length);
                    //回调函数
                    this.BinaryInput(buf);
                    buf = null;
                    iLen = 0;
                    break;
                }
                else
                {
                    buf = new byte[iLen];
                    Array.Copy(data, index + 4, buf, 0, buf.Length);
                    //回调函数
                    this.BinaryInput(buf);
                    buf = null;
                    index += iLen + 4;
                }
            }
        }

        void args_Completed(IAsyncResult reault)
        {
            int cout = 0;
            try
            {
                cout = sock.EndReceive(reault);
            }
            catch (SocketException e)
            {
                socketError = e.SocketErrorCode;
            }
            catch
            {
                socketError = SocketError.HostDown;
            }


            if (socketError == SocketError.Success && cout > 0)
            {
                byte[] buffer = reault.AsyncState as byte[];
                byte[] data = new byte[cout];
                Array.Copy(buffer, 0, data, 0, data.Length);
                if (this.BinaryInput != null)
                    DataMM(data);
                BeginReceive();
            }
            else
            {
                sock.Close();
                if (MessageInput != null)
                    MessageInput("与服务器连接断开");
            }
        }

        // 发送数据包
        public void SendData(byte[] data)
        {
            sock.BeginSend(data, 0, data.Length, SocketFlags.None, AsynCallBack, sock);
        }

        void AsynCallBack(IAsyncResult result)
        {
            try
            {
                Socket sock = result.AsyncState as Socket;

                if (sock != null)
                {
                    sock.EndSend(result);
                }
            }
            catch
            {

            }
        }
        public void Close()
        {
            sock.Close();
        }
    }
}
