/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-5
 * Time: 上午11:11
 * To change this template use File | Settings | File Templates.
 */
package commDefine

import (
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	Op_GetUserIDList = iota // 获取用户ID列表
	Op_GetUser              // 获取用户信息
	Op_GetUserByID          // 根据用户ID获取用户信息
	Op_CreateUser           // 创建一个新用户
	Op_UpdateUser           // 更新用户信息

	Op_GetHeroList // 获取名将列表
	Op_GetHero     // 获取名将信息
	Op_AddHero     // 增加名将
	Op_RemoveHero  // 删除名将
	Op_UpdateHero  // 更新名将信息

	Op_GetHeroSpiritList // 获取武灵列表
	Op_AddHeroSpirit     // 增加武灵
	Op_RemoveHeroSpirit  // 删除武灵

	Op_GetSpiritChipList //获取所有的武灵碎片
	Op_SpiritChipAdd     //增加武灵碎片
	Op_SpiritChipRemove  //删除武灵碎片
	Op_SpiritChipUpdate  //更新武灵碎片

	Op_GetHeroSoulList // 获取武魄列表
	Op_AddHeroSoul     // 增加武魄
	Op_RemoveHeroSoul  // 删除武魄

	Op_DeleteAlly    // 删除好友
	Op_AddAlly       // 增加好友
	Op_DeleteAllyReq // 删除好友请求
	Op_AddAllyReq    // 增加好友请求

	Op_UpdateMagicTrain // 更新诡术修炼的值

	Op_Sign      // 签到操作
	Op_SignClear // 清除签到
	Op_SignGift  // 签到奖励记录

	Op_MissionAdd    //插入一个关卡
	Op_MissionRemove //删除一个关卡
	Op_MissionCover  //覆盖整个关卡数据 修改里面的数组
	Op_GetAllMission //获取所有的关卡

	Op_MsgSystemAdd    //插入一个消息
	Op_MsgSystemRemove //删除一个消息
	Op_MsgSystemUpdate //更新一个消息
	Op_MsgSystemGetUserAll //获取单个角色全部的消息

)

const (
	MAX_FORMATION_POS uint8 = 9 // 阵型位置数
)

// 玩家数据
type DBUser struct {
	Id       bson.ObjectId "_id"
	UserID   uint64        // 账户唯一ID
	Name     string        // 名字
	NickName string        // 昵称
	Password string        // 密码
	Sex      uint8         // 性别 1:女 2:男
	IconID   uint32        // 头像ID

	Level  uint16 // 等级
	Exp    uint32 // 经验
	Energy uint16 // 精力

	AssistHeroId uint64 // 援助名将ID

	HeroGhost uint32 // 武魂数
	Gold      uint32 // 元宝

	AssistPoint uint32 // 支援奖励点数

	AllyList DBAllyList      // 好友列表
	AllyReq  map[string]bool // 好友请求列表

	Formation DBFormation // 阵型

	OnlineTime  int64 // 上线时间
	OfflineTime int64 // 下线时间
	CreateTime  int64 // 创建时间

	// 诡术修炼 加成值
	HP     int32 // 生命
	Attack int32 // 攻击
	Defend int32 // 防御
	Hit    int32 // 命中
	Dodge  int32 // 闪避

	// 签到
	SignList  map[string]bool // 好友请求列表
	SignGift  [5]int          // 签到奖励
	SignMonth int             // 签到月份
}

// 阵型
type DBFormation struct {
	ID           uint8                     // 阵型ID
	PositionList [MAX_FORMATION_POS]uint64 // 阵型各位置名将ID
}

// 名将
type DBHero struct {
	ID               uint64         // 唯一ID
	ConfigId         uint32         // 配置ID
	IsActive         bool           // 是否激活
	IsLocked         bool           // 是否锁定
	StarClass        uint8          // 星级
	Level            uint16         // 等级
	Potential        uint16         // 潜力值
	BreakTimes       uint8          // 名将突破次数
	PotentialAddList map[string]int // 潜力突破表
}

// 名将列表
type DBHeroList struct {
	UserID uint64
	DBHero map[string]*DBHero
}

// 武灵
type DBHeroSpirit struct {
	ID       uint64 // 武灵ID
	ConfigId uint32 // 配置ID
	IsLocked bool   // 是否锁定
}

// 武灵列表
type DBHeroSpiritList struct {
	UserID     uint64
	HeroSpirit map[string]*DBHeroSpirit
}

// 武灵碎片
type DBHeroSpiritChip struct {
	ID            uint64 // 武灵碎片ID
	ConfigId      uint32 // 配置ID
	Count         uint16 // 堆叠数
	SealRuneCount uint16 // 武灵碎片上的封印符数量
}

// 武灵碎片列表
type DBHeroSpiritChipList struct {
	UserID         uint64
	HeroSpiritChip map[string]*DBHeroSpiritChip
}

// 武魄
type DBHeroSoul struct {
	ID       uint64 // 唯一ID
	ConfigId uint32 // 配置ID
	Count    uint16 // 堆叠数
}

// 武魄列表
type DBHeroSoulList struct {
	UserID   uint64
	HeroSoul map[string]*DBHeroSoul
}

// 盟友
type DBAlly struct {
	Id   uint64
	Name string
}

// 盟友列表
type DBAllyList map[string]*DBAlly

// 关卡状态
type DBMissionState struct {
	ConfigId uint32   // 配置ID
	Reward   []uint16 // 已获取的奖励
}

// 关卡列表
type DBMissionStateList struct {
	UserID           uint64
	MissionStateList map[string]*DBMissionState
}

// 消息列表
type DBMessage struct {
	Type int32
	CreateTime int64
	Info string
}

type DBMessageList struct {
	UserID uint64
	MList  map[string]*DBMessage
}

// 活跃玩家列表
type DBActiveUser struct {
	UserID uint64
	OffTime time.Time
}

