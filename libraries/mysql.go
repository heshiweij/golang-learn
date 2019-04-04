package libraries

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// sqlx.db 已经自带了 "连接池"、"线程安全"，可以直接使用
	DB *sqlx.DB
)

func init() {
	// 初始化 mysql
	database, err := sqlx.Open("mysql", "homestead:secret@tcp(192.168.11.11)/sakila")
	if err != nil {
		panic(err)
	}

	DB = database
}

func TestMySQL() {
	// 查询
	//testSelect()

	// 修改
	//testUpdate()

	// 新增
	//testInsert()

	// 删除
	testDelete()
}

type Actor struct {
	ActorId    int    `db:"actor_id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	LastUpdate string `db:"last_update"`
}

func testSelect() {
	var actors []Actor
	err := DB.Select(&actors, "SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id < ?", 5)

	if err != nil {
		panic(err)
	}

	//fmt.Println(actors)
	for i, data := range actors {
		fmt.Println("i => ", i, "data => ", data.ActorId, "|", data.FirstName, "|", data.LastName, "|", data.LastUpdate)
	}
}

func testUpdate() {
	result, err := DB.Exec("update actor set first_name = ? where actor_id = ?", "hsw3", 1)
	if err != nil {
		panic(err)
	}

	//fmt.Println(result)
	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId()) // 只有 insert 语句才会有值
}

func testInsert() {

	result, err := DB.Exec("insert into actor (first_name, last_name) values (?,?)", "何世威", "2018-11-11")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())
}

func testDelete() {
	result, err := DB.Exec("delete from actor where actor_id = ?", 201)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())
}
