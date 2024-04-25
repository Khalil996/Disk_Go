package models

import "time"

type Share struct {
	Id          string    `xorm:"pk varchar(255) notnull default '' 'id' comment('分享id')" json:"id"`
	Pwd         string    `xorm:"varchar(8) notnull default '' 'pwd' comment('分享密码')" json:"pwd"`
	Name        string    `xorm:"varchar(64) notnull default '' 'name'" json:"name"`
	UserId      int64     `xorm:"bigint notnull default 0 'user_id'" json:"userId"`
	Url         string    `xorm:"varchar(255) notnull default '' 'url'" json:"url"`
	Created     time.Time `xorm:"created" json:"created"`
	Expired     int64     `xorm:"bigint notnull default 0 'expired' comment('到期时间')" json:"expired"`
	DownloadNum int64     `xorm:"bigint notnull default 0 'download_num'" json:"downloadNum"`
	ClickNum    int64     `xorm:"bigint notnull default 0 'click_num'" json:"clickNum"`
	Status      int8      `xorm:"tinyint notnull default 0 'status'" json:"status"`
	Type        int8      `xorm:"tinyint notnull default 0 'type'" json:"type"`
	Reason      string    `xorm:"varchar(1023) notnull default '' 'reason'" json:"reason"`
}

func (table *Share) Share() string {
	return "share"
}
