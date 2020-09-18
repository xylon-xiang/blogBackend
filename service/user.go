package service

import (
	"blogBackend/model"
	"blogBackend/util"
)

func GetAllUser() (*[]model.UserInfoReturnModule, error) {

	results, err := model.FindAll("user", "", "")
	if err != nil {
		return nil, err
	}

	allUser := results.([]model.User)

	userInfo := make([]model.UserInfoReturnModule, len(allUser))
	for key, value := range allUser{
		userInfo[key].UserId = value.UserId
		userInfo[key].Email = value.Email
		userInfo[key].Phone = value.Phone
		userInfo[key].UserName = value.UserName
		userInfo[key].RegisterTime = value.RegisterTime
		userInfo[key].UserMotto = value.UserMotto
		userInfo[key].UserImgAddress = value.UserImgAddress
	}

	return &userInfo, nil
}

func GetSpecificUser(userId string) (*model.UserInfoReturnModule, error) {
	result, err := model.FindById("user", userId)

	if err != nil{
		return nil, err
	}

	user := result.(model.User)

	var userInfo model.UserInfoReturnModule
	userInfo.UserImgAddress = user.UserImgAddress
	userInfo.UserMotto = user.UserMotto
	userInfo.RegisterTime = user.RegisterTime
	userInfo.UserName = user.UserName
	userInfo.Phone = user.Phone
	userInfo.UserId	= user.UserId
	userInfo.Email = user.Email

	return &userInfo, nil
}

// the password will be encryption in service layer
func LogJudge(userId string, userPassword string) (bool, error) {

	userPassword = util.EncryptPassword(userPassword)

	result, err := model.FindById("user", userId)
	if err != nil{
		return false, err
	}

	return userPassword == result.(model.User).UserPassword, nil

}
