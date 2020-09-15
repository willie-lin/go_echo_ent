package datasource

import (
	"context"
	"fmt"
	"github.com/facebook/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_echo_ent/ent"
	"go_echo_ent/ent/migrate"
	"log"
	"time"
)

//func Client() *ent.Client {
//	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
//	if err != nil {
//		log.Fatalf("failed opening connection to sqlite: %v", err)
//	}
//
//	defer client.Close()
//
//	if err := client.Schema.Create(context.Background()); err != nil {
//		log.Fatalf("failed creating schema resources: %v", err)
//	}
//	return client
//}

const (
	driverName      = "mysql"
	dataSourceName  = "root:root1234@tcp(127.0.0.1:3306)/ent?charset=utf8&parseTime=true"
	maxIdleConns    = 6
	maxOpenConns    = 100
	connMaxLifetime = time.Hour * 2
)

func Clients() (*ent.Client, error) {
	drv, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/ent?charset=utf8&parseTime=true")
	//drv, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/ent")
	if err != nil {
		return nil, err
	}
	db := drv.DB()
	//db.SetMaxIdleConns(maxIdleConns)
	//db.SetConnMaxLifetime(connMaxLifetime)
	//db.SetMaxOpenConns(maxOpenConns)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	fmt.Println("ccc")
	return ent.NewClient(ent.Driver(drv)), nil

}

//func Clients() (*ent.Client, error) {
//	db, err := sql.Open(driverName, dataSourceName)
//	if err != nil {
//		return nil, err
//	}
//	//db := drv.DB()
//	db.SetMaxIdleConns(maxIdleConns)
//	db.SetConnMaxLifetime(connMaxLifetime)
//	db.SetMaxOpenConns(maxOpenConns)
//	drv := entsql.OpenDB(driverName, db)
//	return ent.NewClient(ent.Driver(drv)), nil
//}

func AutoMigration(client *ent.Client, ctx context.Context) {
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func DebugMode(err error, client *ent.Client, ctx context.Context) {
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
