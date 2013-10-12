using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using LitJson;
using System.Drawing;

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
            var spr = GetSprite(SelfUUID);
            var sprMove = spr.Location;
            if (keyData == (Keys.W))
            {
                sprMove.Y--;
            }

            if (keyData == (Keys.S))
            {
                sprMove.Y++;
            }

            if (keyData == (Keys.A))
            {
                sprMove.X--;
            }

            if (keyData == (Keys.D))
            {
                sprMove.X++;
            }

            spr.Location = sprMove;

            JsonData jd = new JsonData();
            jd["OPCODE"] = (Int32)GameOpcode.MoveSprite;
            jd["UUID"] = SelfUUID;
            jd["X"] = spr.Location.X;
            jd["Y"] = spr.Location.Y;
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
                        Point p = new Point(50,50);
                        CreateSprite(uuid, p);
                        fm.Text = "" + uuid;
                    }
                    break;
                case (Int32)GameOpcode.CreateSprite:
                    {
                        var uuid = (Int64)jd["UUID"];
                        Point p = new Point(50, 50);
                        CreateSprite(uuid, p);
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
                case (Int32)GameOpcode.RemoveSprite:
                    {
                        var uuid = (Int64)jd["UUID"];
                        RemoveSprite(uuid);
                    }
                    break;
            }
        }

        public bool RemoveSprite(Int64 uuid)
        {
            var sp = SpriteList[uuid];
            fm.Controls.Remove(sp);
            fm.ResumeLayout(false);
            SpriteList.Remove(uuid);
            return true;
        }
            
        public Sprite CreateSprite(Int64 uuid, Point p )
        {
            Sprite sp = new Sprite();
            sp.Location = p;
            sp.Name = ""+uuid;
            sp.Size = new System.Drawing.Size(75, 23);
            sp.TabIndex = 1;
            sp.Text = ""+uuid;
            sp.UseVisualStyleBackColor = true;
            SpriteList[uuid] = sp;

            fm.Controls.Add(sp);
            fm.ResumeLayout(false);
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
