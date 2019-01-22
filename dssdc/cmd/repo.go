package cmd

import (
	"github.com/ddosakura/ds-watcher-simple-dev/repo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // use sqlite
)

func initRepo(dbname string) {
	repo.Pre = func() *gorm.DB {
		db, err := gorm.Open("sqlite3", dbname)
		if err != nil {
			er(err)
		}
		return db
	}
	repo.Init()
}
