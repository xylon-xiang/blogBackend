package model

type User struct {
	UserId         string `json:"user_id"`
	UserName       string `json:"user_name"`
	UserPassword   string `json:"user_password"`
	UserImgAddress string `json:"user_img_address"`
	UserMotto      string `json:"user_motto"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	RegisterTime   int64  `json:"register_time"`
}

type UserLogPost struct {
	UserId       string `json:"user_id"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type UserLogReturnModule struct {
	JwtToken string `json:"jwt_token"`
}

type UserInfoReturnModule struct {
	UserId         string `json:"user_id"`
	UserName       string `json:"user_name"`
	UserImgAddress string `json:"user_img_address"`
	UserMotto      string `json:"user_motto"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	RegisterTime   int64  `json:"register_time"`
}
