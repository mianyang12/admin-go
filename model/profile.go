package model

type Profile struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	UserId   int    `gorm:"foreignKey;not null" json:"-"`
	NickName string `gorm:"-" json:"nickName"`
}

func (Profile) TableName() string {
	return "profile"
}
