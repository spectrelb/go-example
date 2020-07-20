package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var Db *sqlx.DB

func init()  {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func main() {
	insert()
	selectSql()
	update()
	delete()
	transaction()
}

func insert()  {
	r, err := Db.Exec("insert into person(username, sex, email) values (?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func selectSql()  {
	var person []Person
	err := Db.Select(&person, "select * from  person where user_id = ?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}

func update()  {
	res, err := Db.Exec("update person set username=? where user_id= ?", "stu0003", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ",err)
	}
	fmt.Println("update succ:",row)
}


func delete()  {

	res, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row,err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ",err)
	}

	fmt.Println("delete succ: ",row)
}

func transaction()  {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}