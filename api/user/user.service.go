package user

import (
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"gorm.io/gorm"
)

func CheckPhoneService(phone string) bool {
	client := global.DB.GetClient()
	if client.Model(&model.User{}).Where("phone = ?", phone).First(&model.User{}).RowsAffected > 0 {
		return true
	}
	return false
}

func CheckNameService(name string) bool {
	client := global.DB.GetClient()
	if client.Model(&model.User{}).Where("name = ?", name).First(&model.User{}).RowsAffected > 0 {
		return true
	}
	return false
}

func CreateUserService(user *model.User) (err error) {
	client := global.DB.GetClient()
	tx := client.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	// If there are no issues, submit the transaction; if there are issues, roll back to ensure that the data is error free
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func GetUserServiceByName(name string) (user model.User, err error) {
	client := global.DB.GetClient()
	err = client.Model(&model.User{}).Where("name = ?", name).First(&user).Error
	if user.Uid == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func GetUserServiceByUid(uid uint) (user model.User, err error) {
	client := global.DB.GetClient()
	err = client.Model(&model.User{}).Where("uid = ?", uid).First(&user).Error
	if user.Uid == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func UpdateUserService(user *model.User) (err error) {
	client := global.DB.GetClient()
	if err = client.Model(&model.User{}).Where("uid = ?", user.Uid).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserServiceList(pageSize, pageNum int) (users []model.User, err error) {
	client := global.DB.GetClient()
	err = client.Model(&model.User{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	return
}
