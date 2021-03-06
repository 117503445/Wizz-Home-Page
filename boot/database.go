package boot

import (
	"database/sql"
	"fmt"
	"strings"
	"wizz-home-page/app/model"
	"wizz-home-page/library"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

//InitDatabase Create database if not exists
func InitDatabase() {
	link := g.Cfg().GetString("database.link")
	// g.Log().Line().Debug(link)

	linkWithoutDbName := strings.Join(strings.Split(strings.Split(link, "/")[0], ":")[:][1:], ":") + "/"
	dbName := strings.Split(link, "/")[1]

	// g.Log().Line().Debug(linkWithoutDbName)

	sqlDB, err := sql.Open("mysql", linkWithoutDbName)
	if err != nil {
		g.Log().Line().Panic(err)
	}
	defer sqlDB.Close()

	isForceCreate := g.Cfg().GetBool("database.forceCreate")
	if isForceCreate {
		query := fmt.Sprintf("DROP DATABASE IF EXISTS %v", dbName)
		if _, err = sqlDB.Exec(query); err != nil {
			g.Log().Line().Panic(err)
		}
	}

	query := fmt.Sprintf("SELECT * FROM information_schema.SCHEMATA where SCHEMA_NAME='%v'", dbName)
	// g.Log().Debug(query)

	result, err := sqlDB.Query(query)
	if err != nil {
		g.Log().Line().Panic(err)
	}

	isDbExists := false
	for result.Next() {
		isDbExists = true // 只要查到了一条记录,就说明数据库存在
		break
	}
	// g.Log().Debug(isDbExists)
	if !isDbExists {
		g.Log().Line().Info(fmt.Sprintf("create database %v", dbName))
		if _, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName)); err != nil {
			g.Log().Line().Panic(err)
		}

		sqlMyDB, err := sql.Open("mysql", linkWithoutDbName+dbName+"?multiStatements=true")
		if err != nil {
			g.Log().Line().Panic(err)
		}

		arraySqlPath := g.Cfg().GetArray("database.sqlOnCreate")
		for i, path := range arraySqlPath {
			g.Log().Line().Debug("run ", path)
			sqlText := gfile.GetContents(path.(string))
			if _, err = sqlMyDB.Exec(sqlText); err != nil {
				g.Log().Line().Panic(err)
			}
			if i == 0 { // 在 第一条create.sql 建表以后插入数据
				// todo 实现的不好,需要重构
				adminPassword := library.RandStringRunes(12)
				if cipher, err := model.EncryptPassword(adminPassword); err != nil {
					g.Log().Line().Panic(err)
				} else {
					if err = gfile.PutContents("./tmp/password/admin.txt", adminPassword); err != nil {
						g.Log().Line().Error(err)
					}
					_, _ = g.DB().Table("role").Data(g.List{{"name": "admin"}, {"name": "user"}, {"name": "interviewer"}}).Save()
					_, _ = g.DB().Table("user").Data(g.List{{"username": "admin", "password": cipher}, {"username": "interviewer", "password": cipher}}).Save()
					_, _ = g.DB().Table("user_role").Data(g.List{{"user_id": "1", "role_id": "1"}, {"user_id": "1", "role_id": "2"}, {"user_id": "2", "role_id": "3"}}).Save()
				}
			}

		}

	}
}
