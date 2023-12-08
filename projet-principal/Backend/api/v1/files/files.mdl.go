package files

import (
	"gorm.io/gorm"
)

type File struct {
	ID        uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	FilePath  string `gorm:"column:file_path;not null;" json:"file_path"`
	Filename  string `gorm:"column:file_name;not null;" json:"filename"`
	FileTheme string `gorm:"column:file_theme;not null;" json:"file_theme"`
	Size      int64  `gorm:"column:size;not null;" json:"size"`
	UserID    uint   `gorm:"column:user_id;not null;" json:"user_id"`
	gorm.Model
}

func GetFileById(db *gorm.DB, id uint) (file File, err error) {
	return file, db.First(&file, "id=?", id).Error
}

func DeleteFile(db *gorm.DB, file_id uint) error {
	return db.Where("id=?", file_id).Delete(&File{}).Error
}

func GetMostRecentFile(db *gorm.DB) (file File, err error) {
	return file, db.Order("created_at desc").First(&file).Error
}

func CheckFileExists(db *gorm.DB, id int64) bool {

	// init vars
	file := &File{}

	// check if row exists
	check := db.Where("id=?", id).First(&file)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get all file
func GetFiles(db *gorm.DB) (file []File, err error) {
	return file, db.Find(&file).Error
}
