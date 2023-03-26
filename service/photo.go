package service

import (
	"my-admin/global"
	"my-admin/model"
)

func CreatePhoto(photo model.Photo) {

}

func DeletePhoto(photo model.Photo) {

}

func insertPhoto(phote model.Photo) error {
	return global.DBClient.Create(&phote).Error
}

func deletePhoto(photo model.Photo) error {
	return global.DBClient.Delete(photo).Error
}

func isPhotoOutOfCapacity() bool {
	return false
}
