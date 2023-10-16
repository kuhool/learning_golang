package oneDayOneLibray

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//	type ServerConfig struct {
//		Database map[string]string `json:"database"`
//	}
type ServerConfig struct {
	Database struct {
		Host     string `mapstructure: "host"`
		Port     int    `mapstructure: "port"`
		Username string `mapstructure: "username"`
		Password string `mapstructure: "password"`
	} `mapstructure: "database"`
}

func ReadInit() {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	//读取配置文件的内容到viper
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//拿到配置文件的两种方式
	//1.v.Get(配置名字)
	fmt.Println(v.Get("database.host"))

	//2.把配置的值封装到结构体
	//需要反编码到结构体
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

}
func ReadInit1() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./") // 假设配置文件在当前目录的config文件夹下

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("无法读取配置文件:  %v", err)
	}

	// 读取配置项
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")

	// 打印读取到的配置项
	fmt.Printf("数据库主机:%s\n", host)
	fmt.Printf("数据库端口:  %d\n", port)
	fmt.Printf("数据库用户名:  %s\n", username)
	fmt.Printf("数据库密码:  %s\n", password)
}
