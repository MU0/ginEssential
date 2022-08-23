package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(20);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}

func main() {

	db := InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证

		if len(telephone) != 11 {
			ctx.JSON(422, gin.H{"code": 422, "msg": "手机号必须要11位"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		//如果名称没有传，给一个10为的随着字符串
		if len(name) == 0 {
			name = RandomString(10)
			return
		}

		log.Println(name, telephone, password)
		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(422, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}

		db.Create(&newUser)

		//返回结果
		ctx.JSON(200, gin.H{"msg": "注册成功"})
	})
	panic(r.Run())
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8mb4"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	return db
}
