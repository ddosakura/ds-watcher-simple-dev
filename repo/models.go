package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Notes struct {
	gorm.Model
	Developer  string
	Project    string
	File       string
	Ext        string
	ChangeTime time.Time
}
