package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	//dsn := "root:123456@tcp(115.159.69.196:3306)/mxshop_user?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // 禁用彩色打印
	//	},
	//)
	//
	//// 全局模式
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		// 生成的表为创建实体的结构体名
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}

	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(newPassword)
	//
	//for i := 0; i < 10; i++ {
	//	user := model.User{
	//		NickName: fmt.Sprintf("bobby%d", i),
	//		Mobile:   fmt.Sprintf("1878222222%d", i),
	//		Password: newPassword,
	//	}
	//	db.Save(&user)
	//}

	//设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	//sql才是最重要的，本着这个原则我尽量的给大家看到每个api背后的sql语句是什么

	//定义一个表结构， 将表结构直接生成对应的表 - migrations
	// 迁移 schema
	//_ = db.AutoMigrate(&model.User{}) //此处应该有sql语句

	// fmt.Println(genMd5("123456"))
	//将用户的密码变一下 随机字符串+用户密码
	//暴力破解 123456 111111 000000 彩虹表 盐值
	//e10adc3949ba59abbe56e057f20f883e
	//e10adc3949ba59abbe56e057f20f883e

	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(len(newPassword)) // 96
	fmt.Println(newPassword)      //$pbkdf2-sha512$uk2H3xakGvJipKdf$5b4703ff649d39e453caba1a164b0b3e00631cdbc642a5cd0ff242b6a17fb215

	passwordInfo := strings.Split(newPassword, "$")
	fmt.Println(passwordInfo) // [ pbkdf2-sha512 uk2H3xakGvJipKdf 5b4703ff649d39e453caba1a164b0b3e00631cdbc642a5cd0ff242b6a17fb215]
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
}