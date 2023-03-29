package conf

import (
	"IM/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	LoadServer()
	LoadMysqlData()
	//MySQL
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}

func LoadMysqlData() {
	Db = viper.GetString("datasource.driverName")
	DbHost = viper.GetString("datasource.host")
	DbPort = viper.GetString("datasource.port")
	DbName = viper.GetString("datasource.database")
	DbUser = viper.GetString("datasource.username")
	DbPassWord = viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
		charset)
	db, err := gorm.Open(Db, args)
	db.LogMode(true) //GORM的打印
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true) //默认不加复数s
}

func LoadServer() {
	HttpPort = ":" + viper.GetString("server.port")
}
