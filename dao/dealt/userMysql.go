package dealt

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"fmt"
)

func CheckUserExist(p *models.UserSignUp) bool {
	var num int
	sqlStr := "select count(user_id) from user where username=?"
	if err := mysql.DB.Get(&num, sqlStr, p.UserName); err != nil {
		return false
	} else {
		fmt.Printf("ret=%v\n", &num)
		if num >= 1 {
			return false
		} else {
			return true
		}
	}
}

func InsertUser(p *models.UserSignUp) (err error) {
	sqlStr := "insert into user(user_id, username, password, email, gender) values (?,?,?,?,?)"
	_, err = mysql.DB.Exec(sqlStr, p.UserId, p.UserName, p.PassWord, p.Email, p.Gender)
	if err != nil {
		return err
	} else {
		return nil
	}
}
