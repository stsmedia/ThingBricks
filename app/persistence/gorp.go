package persistence

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"log"
	"os"
)

var (
	Dbm *gorp.DbMap
)

func InitDb() {
	dbopen, err := sql.Open("postgres", "host=localhost user=thingbricks dbname=thingbricks sslmode=disable")
	checkErr(err, "sql.Open failed")

	Dbm = &gorp.DbMap{Db: dbopen, Dialect: gorp.PostgresDialect{}}

	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	Dbm.AddTableWithName(models.Account{}, "accounts").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Login{}, "logins").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.KeyPair{}, "api_keys").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.AccountGroup{}, "account_groups").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Project{}, "projects").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.DataStreamGroup{}, "data_stream_groups").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.DataStream{}, "data_streams").SetKeys(true, "Id")

	err = Dbm.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	revel.INFO.Println("initializing db")
	log.Println("setting db " + Dbm.Dialect.TruncateClause())

	//	return &DB{Gorp: Dbm}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
