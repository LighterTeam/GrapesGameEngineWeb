using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Runtime.Serialization;

namespace GameClient
{
    [DataContract]
    public class MessagePacket
    {
        [DataMember]
        internal Int32 T;

        [DataMember]
        internal Int64 S;

        [DataMember]
        internal String D;
    }

    public enum SendMode
    {
    	SendBroad,          //广播
	    SendSingle,			//单发
	    SendBroadNotSelf,	//广播没有自己
    }

    public enum GameOpcode
    {
        RegistServer,
        CreateSprite,
        MoveSprite,
    }
}
