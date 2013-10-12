using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using LitJson;

namespace GameClient
{
    public class Sprite : Button
    {
        public void SetLoc(int x, int y)
        {
            this.Location = new System.Drawing.Point(x, y);
        }
    }

    public class TSGame
    {
        private Form fm;
        public Int64 SelfUUID;

        public TSGame(Form fm)
        {
            this.fm = fm;
        }

        public Dictionary<Int64, Sprite> SpriteList = new Dictionary<Int64, Sprite>();

        public void Init() {

        }

        public bool KeyDown(Keys keyData)
        {
            if (keyData == (Keys.W))
            {
            }

            if (keyData == (Keys.S))
            {
            }

            if (keyData == (Keys.A))
            {
            }

            if (keyData == (Keys.D))
            {
            }

            JsonData jd = new JsonData();
            jd["OPCODE"] = (Int32)GameOpcode.MoveSprite;
            jd["UUID"] = SelfUUID;
            jd["X"] = GetSprite(SelfUUID).Location.X;
            jd["Y"] = GetSprite(SelfUUID).Location.Y;
            Program.SendDataBroad(jd.ToJson());
            return true;
        }

        public void MessageProc(string msg)
        {
            JsonData jd = JsonMapper.ToObject(msg);
            if (jd == null)
            {
                Console.WriteLine("JsonData Error! msg:" + msg);
                return;
            }
            var opCode = (Int32)jd["OPCODE"];
            switch (opCode)
            {
                case (Int32)GameOpcode.RegistServer:
                    {
                        var uuid = (Int64)jd["UUID"];
                        SelfUUID = uuid;
                    }
                    break;
                case (Int32)GameOpcode.CreateSprite:
                    {
                        var uuid = (Int64)jd["UUID"];
                        CreateSprite(uuid);
                    }
                    break;
                case (Int32)GameOpcode.MoveSprite:
                    {
                        var uuid = (Int64)jd["UUID"];
                        var x = (Int32)jd["X"];
                        var y = (Int32)jd["Y"];
                        GetSprite(uuid).SetLoc(x, y);
                    }
                    break;
            }
        }
            
        public Sprite CreateSprite(Int64 uuid)
        {
            Sprite sp = new Sprite();
            sp.Location = new System.Drawing.Point(208, 400);
            sp.Name = ""+uuid;
            sp.Size = new System.Drawing.Size(75, 23);
            sp.TabIndex = 1;
            sp.Text = ""+uuid;
            sp.UseVisualStyleBackColor = true;
            SpriteList[uuid] = sp;

            return sp;
        }

        public Sprite GetSprite(Int64 uuid)
        {
            if (SpriteList.ContainsKey(uuid))
            {
                return SpriteList[uuid];
            }
            return null;
        }
    }
}
