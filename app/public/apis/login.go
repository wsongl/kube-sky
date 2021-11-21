package apis

import (
	"kube-sky/app/system/models"
	"kube-sky/common/server/loginlog"
	"kube-sky/pkg/conn"
	"kube-sky/pkg/jwtauth"
	"kube-sky/pkg/tools/response"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// Login 登陆
func Login(c *gin.Context) {
	var (
		err       error
		token     string
		user      models.UserRequest
		loginUser struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	)
	err = c.ShouldBind(&loginUser)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	// 查询用户信息
	err = conn.Orm.Model(&models.User{}).Where("username = ?", loginUser.Username).Find(&user).Error
	if err != nil {
		response.Error(c, err, response.QueryUserError)
		return
	}

	if user.Username == "" {
		response.Error(c, nil, response.UserNotExistError)
		return
	}

	if !user.Status {
		response.Error(c, nil, response.UserUnavailableError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		go loginlog.Create(c, user.Username, "密码不正确")
		response.Error(c, err, response.IncorrectPasswordError)
		return
	}

	token, err = jwtauth.GenToken(&user)
	if err != nil {
		response.Error(c, err, response.GenerateTokenError)
		return
	}

	go loginlog.Create(c, user.Username, "登陆成功")

	response.OK(c, token, "")
}
