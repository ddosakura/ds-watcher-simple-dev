package repo

import "github.com/jinzhu/gorm"

var Pre func() *gorm.DB

func Init() {
	db := Pre()
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Notes{})
}

func Note(note *Notes) {
	db := Pre()
	defer db.Close()
	db.Create(note)
}

func Detail(name string) *[]Notes {
	db := Pre()
	defer db.Close()
	var notes []Notes
	db.Find(&notes, "developer = ?", name)
	return &notes
}

func Developers() []string {
	db := Pre()
	defer db.Close()
	var notes []Notes
	db.Select("developer").Group("developer").Find(&notes)
	result := []string{}
	for _, v := range notes {
		result = append(result, v.Developer)
	}
	return result
}
