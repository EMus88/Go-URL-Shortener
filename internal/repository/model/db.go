package model

type Session struct {
	ID        int    `gorm:"primary_key;auto_increment;unique"`
	SessionID string `gorm:"type:varchar(255);not null;unique"`
}

type Shorten struct {
	ID        int    `gorm:"primary_key;auto_increment;unique" json:"-"`
	URLID     string `gorm:"type:varchar(255);not null;unique" json:"-"`
	ShortURL  string `gorm:"type:varchar(255);not null" json:"short_url"`
	LongURL   string `gorm:"type:varchar(255);not null;unique" json:"original_url"`
	SessionID int    `gorm:"type:int;not null" json:"-"`
	IsDeleted bool   `gorm:"type:bool; default:false" json:"-"`
}
