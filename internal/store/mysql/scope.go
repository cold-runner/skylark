package mysql

import "gorm.io/gorm"

func LarkByStuNum(stuNum string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stu_num = ?", stuNum)
	}
}

func LarkByPhone(phone string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("phone = ?", phone)
	}
}

func LarkByQqUnionId(qqUnionId string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("qq_union_id = ?", qqUnionId)
	}
}

func LarkQqUnionIdIsNull(db *gorm.DB) *gorm.DB {
	return db.Where("qq_union_id is NULL")
}
