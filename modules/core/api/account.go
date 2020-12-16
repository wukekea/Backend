package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type HobbyName = string
type VipLevel = int8
type VipName = string

var VipMap map[VipLevel]VipName = map[VipLevel]VipName{
	Soldier:      "Soldier",       // 士兵
	Sergeant:     "Sergeant",      // 中士
	Captain:      "Captain",       // 上尉
	Major:        "Major",         // 少校
	Colonel:      "Colonel",       // 上校
	General:      "General",       // 上将
	FieldMarshal: "Field Marshal", // 元帅
}

const (
	PlayingBasketball HobbyName = "playing basketball"
	PlayingFootBall   HobbyName = "playing football"
	PlayingPingPang   HobbyName = "playing ping pang"
	PlayingTennis     HobbyName = "playing tennis"
	Walking           HobbyName = "walking"
	SingingSongs      HobbyName = "singing songs"
	Dancing           HobbyName = "dancing"
	PlayingGame       HobbyName = "playing games"
	WatchingMovie     HobbyName = "watching movies"
	ReadingBooks      HobbyName = "reading books"
	Studying          HobbyName = "studying"
)
const (
	Soldier VipLevel = iota
	Sergeant
	Captain
	Major
	Colonel
	General
	FieldMarshal
)

const (
	SexUnknown int8 = iota
	SexBoy
	SexGirl
)

// AccountInfo为account表储存的数据
type AccountInfo struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	MID      string             `json:"id" bson:"id"`             // 账户id
	Name     string             `json:"name" bson:"name"`         // 用户名
	Password string             `json:"password" bson:"password"` // 用户密码
	Age      int8               `json:"age" bson:"age"`           // 用户年龄
	Sex      int8               `json:"sex" bson:"sex"`           // 用户性别；0：未知；1：男；2：女
	Hobbies  []Hobby            `json:"hobbies" bson:"hobbies"`   // 用户爱好
	Address  string             `json:"address" bson:"address"`   // 用户地址
	Vip      Vip                `json:"vip" bson:"vip"`           // vip信息
}
type Hobby struct {
	Name      string `json:"name" bson:"name"`             // 爱好名称
	StartTime int64  `json:"start_time" bson:"start_time"` // 爱好开始时间 毫秒表示
	Weight    int8   `json:"weight" bson:"weight"`         // 爱好权重（权重越高，用户越喜欢该爱好）
}
type Vip struct {
	VipName      VipName  `json:"vip_name" bson:"vip_name"`           // vip名字
	VipLevel     VipLevel `json:"vip_level" bson:"vip_level"`         // vip等级
	StartTime    int64    `json:"start_time" bson:"start_time"`       // vip获取的开始时间 毫秒表示
	DurationTime int64    `json:"duration_time" bson:"duration_time"` // vip可以持续的时间
}
