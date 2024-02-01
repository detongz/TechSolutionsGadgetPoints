package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

// 用户登录注册注销，用户信息

var jwtKey = []byte("your_secret_key")

// Users struct represents the Users table
type Users struct {
	UserID   int       `json:"userID"`   // Unique identifier for the user
	Username string    `json:"username"` // User's chosen username
	Password string    `json:"password"` // User's hashed password
	Email    string    `json:"email"`    // User's email address
	JoinDate time.Time `json:"joinDate"` // Timestamp of when the user joined
}

// Claims contains the claims of the token
type Claims struct {
	UserID   int
	Username string
	jwt.StandardClaims
}

// Login handles user login and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	// 这里应该有检查用户名和密码的逻辑，我们假设它已经完成
	// 假设user.Username和user.Password是正确和合法的用户凭证
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		user.UserID,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}
	// 假设我们存储了用户信息，我们在这里可以更新用户的登录时间等
	// 返回token
	w.Write([]byte(tokenString))
}

// 注册函数
func Register(id, username, password string) bool {

}

// 查询用户信息函数
func QueryUserInfo(username string) (User, bool) {
	userMutex.Lock()
	defer userMutex.Unlock()
	// 检查用户是否存在
	user, exists := users[username]
	if !exists {
		return User{}, false
	}
	return user, true
}

// 注销用户函数
func Deregister(username string) bool {
	userMutex.Lock()
	defer userMutex.Unlock()
	// 检查用户是否存在
	if _, exists := users[username]; !exists {
		return false
	}
	// 删除用户
	delete(users, username)
	return true
}

func Logout(w http.ResponseWriter, r *http.Request) {}

// 用户查询积分进度
func Progress() {}
