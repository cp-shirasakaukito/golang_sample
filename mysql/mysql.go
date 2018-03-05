package mysql

import (
	"os"
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

//環境変数を取得fallbackはローカル検証用
var (
	endpoint	string	= getEnv("DB_ENDPOINT", "127.0.0.1")
	user 		string	= getEnv("DB_USER", "root")
	password	string	= getEnv("DB_PASSWORD", "root")
	database	string	= getEnv("DB_DATABASE", "dezura")
	port		string	= getEnv("DB_PORT", "33306")
)

//環境変数が存在していれば返却、なければfallbackで指定した文字列を返却
func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallback
}

func Exec() {
	// initiate
	dbmap := InitDb()
	defer dbmap.Db.Close()

	// make insert datum
	InsertStrt := newCompany("10001","test", "hoge@example.com","06-6012-3456")
	err := dbmap.Insert(&InsertStrt)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("insert succeeded!")

}

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", user + ":" + password + "@tcp(" + endpoint + ":" + port + ")/" + database)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// attach the table 'companies' to Company strut
	dbmap.AddTableWithName(Company{}, "companies").SetKeys(true, "Id")


	return dbmap
}

func newCompany(code, name, mail, tel string) Company {
	return Company{
		Code:code,
		Name:name,
		Mail:mail,
		Tel:tel,
		DelFlg:0,
	}
}

type Company struct {
	Id			int64	`db:"id"`
	Code		string	`db:"code"`
	Name		string	`db:"name"`
	Mail		string	`db:"mail"`
	Tel			string	`db:"tel"`
	DelFlg		int8	`db:"delete_flg"`
}



