package models

import "time"

type LogBasic struct {
	Model         `xorm:"extends"`
	Ip            string        `json:"ip,omitempty"`
	StatusCode    int           `json:"status_Code,omitempty"`
	TimeConsuming time.Duration `json:"time_Consuming,omitempty"`
	UserId        int64         `json:"userIdentity,omitempty"`
	Methods       string        `json:"methods,omitempty"`
	Path          string        `json:"path,omitempty"`
	RequestTime   time.Time     `json:"requestTime"`
	Role          string        `json:"role,omitempty"`
}
