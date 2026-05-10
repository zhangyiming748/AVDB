package sqlite

import (
	"gorm.io/gorm"
)

type AVDB struct {
	gorm.Model        // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	NO         string `gorm:"uniqueIndex;size:64;not null"` // 作品番号，唯一索引
	Title      string `gorm:"size:2048;not null"`           // 原始标题
	ZhCnTitle  string `gorm:"size:2048"`                    // 中文标题
	Pretty     string `gorm:"size:255"`                     // 美化后的显示名称
}

// CreateAVDB 创建记录
func CreateAVDB(db *gorm.DB, avdb *AVDB) error {
	return db.Create(avdb).Error
}

// GetAVDBByID 根据ID查询
func GetAVDBByID(db *gorm.DB, id uint) (*AVDB, error) {
	var avdb AVDB
	err := db.First(&avdb, id).Error
	if err != nil {
		return nil, err
	}
	return &avdb, nil
}

// GetAVDBByNO 根据番号查询
func GetAVDBByNO(db *gorm.DB, no string) (*AVDB, error) {
	var avdb AVDB
	err := db.Where("no = ?", no).First(&avdb).Error
	if err != nil {
		return nil, err
	}
	return &avdb, nil
}

// SearchAVDBByTitle 根据标题模糊查询（支持原始标题和中文标题）
func SearchAVDBByTitle(db *gorm.DB, keyword string, limit int) ([]AVDB, error) {
	var avdbs []AVDB
	searchPattern := "%" + keyword + "%"
	err := db.Where("title LIKE ? OR zh_cn_title LIKE ?", searchPattern, searchPattern).
		Limit(limit).
		Find(&avdbs).Error
	return avdbs, err
}

// SearchAVDBByNO 根据番号模糊查询
func SearchAVDBByNO(db *gorm.DB, keyword string, limit int) ([]AVDB, error) {
	var avdbs []AVDB
	searchPattern := "%" + keyword + "%"
	err := db.Where("no LIKE ?", searchPattern).
		Limit(limit).
		Find(&avdbs).Error
	return avdbs, err
}

// UpdateAVDB 更新记录
func UpdateAVDB(db *gorm.DB, avdb *AVDB) error {
	return db.Save(avdb).Error
}

// UpdateAVDBFields 更新指定字段
func UpdateAVDBFields(db *gorm.DB, id uint, updates map[string]interface{}) error {
	return db.Model(&AVDB{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteAVDB 软删除记录
func DeleteAVDB(db *gorm.DB, id uint) error {
	return db.Delete(&AVDB{}, id).Error
}

// DeleteAVDBByNO 根据番号软删除记录
func DeleteAVDBByNO(db *gorm.DB, no string) error {
	return db.Where("no = ?", no).Delete(&AVDB{}).Error
}

// ListAVDBs 分页查询所有记录
func ListAVDBs(db *gorm.DB, page, pageSize int) ([]AVDB, int64, error) {
	var avdbs []AVDB
	var total int64

	// 查询总数
	if err := db.Model(&AVDB{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Find(&avdbs).Error
	return avdbs, total, err
}
