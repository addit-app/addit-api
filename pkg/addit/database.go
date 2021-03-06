package addit

import (
	"fmt"
	"time"
	"log/syslog"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	engine := connectDB()
	syncDB(engine)
}

func connectDB() *xorm.Engine {
	var err error
	var engine *xorm.Engine

	var (
		DBMS_ID = GetEnv("DBMS_ID", "root")
		DBMS_PW = GetEnv("DBMS_PW", "op0023")
		CONNECT = GetEnv("CONNECT", "127.0.0.1:3306")
		TABLE   = GetEnv("TABLE",   "addit")
		DbURI  = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?%s",
			DBMS_ID,
			DBMS_PW,
			CONNECT,
			TABLE,
			"charset=utf8&parseTime=True&loc=Local",
		)

	)

	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-api-log")
	if err != nil {
		panic(err)
	}

	xormLog := xorm.NewSimpleLogger(logWriter)
	xormLog.ShowSQL(true)
	xormLog.SetLevel(core.LOG_DEBUG)
	xormLog.IsShowSQL()

	engine, err = xorm.NewEngine("mysql", DbURI)
	engine.SetLogger(xormLog)
	engine.TZLocation, _ = time.LoadLocation("Asia/Seoul")

	return engine
}

func syncDB( engine *xorm.Engine ) {
	if err := engine.Sync2(new(UrlIndex)); err != nil {
		panic(err)
	}

	if err := engine.Sync2(new(ChainIndex)); err != nil {
		panic(err)
	}
}

func InsertContents(url, hash string) (error) {
	contents := new(UrlIndex)

	contents.Url = url
	contents.Hash = hash
	contents.Count = 1

	engine := connectDB()
	defer engine.Close()

	_, err := engine.InsertOne(contents)
	if err != nil {
		return err
	}
	return nil

}

func SelectContents(hash string) (UrlIndex, bool, error) {
	var contents = UrlIndex{Hash:hash}
	engine := connectDB()
	defer engine.Close()

	has, err := engine.Get(&contents)
	if err != nil {
		return contents, false, err
	}

	return contents, has, nil
}

func UpdateContents(hash string, count int) (UrlIndex, error) {
	var contents UrlIndex

	contents.Count = count

	engine := connectDB()
	defer engine.Close()

	_, err := engine.In("hash", hash).Update(&contents)
	if err != nil {
		return contents, err
	}

	index, _, err := SelectContents(hash)
	return index, err
}

func InsertChainIndex(hash string, index int) (error) {
	contents := new(ChainIndex)

	contents.Hash = hash
	contents.Chainid = index

	engine := connectDB()
	defer engine.Close()

	_, err := engine.InsertOne(contents)
	if err != nil {
		return err
	}

	return nil
}

func SelectChainIndex(hash string) ([]ChainIndex, error) {
	var indexes []ChainIndex

	engine := connectDB()
	defer engine.Close()

	err := engine.Find(&indexes, &ChainIndex{Hash:hash})
	if err != nil {
		return indexes, err
	}

	return indexes, nil
}