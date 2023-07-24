package oneDayOneLibray

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"
)

var Engine *xorm.Engine

func init() {
	engine, err := xorm.NewEngine("mysql", "lvweb1:lavion#2013@tcp(polar-dev.rwlb.rds.aliyuncs.com:3306)/db_product_core?charset=utf8")
	if err != nil {
		fmt.Println("error =>>>>>>>", err)
	}
	Engine = engine
	Engine.ShowSQL(true)                     //则会在控制台打印出生成的SQL语句；
	Engine.Logger().SetLevel(core.LOG_DEBUG) //则会在控制台打印调试及以上的信息；
}

type TbLog struct {
	Id         int       `json:"id" xorm:"'id' not null pk autoincr comment('自增id') INT(10)"`
	LogItemId  int       `json:"log_item_id" xorm:"'log_item_id' not null default 0 comment('日志项ID') index(idx_log_item_id) INT(10)"`
	Source     int       `json:"source" xorm:"'source' not null default 0 comment('源 1-spu 2-cspu 3-product 4-sku 5-sku_item') TINYINT(3)"`
	OpType     int       `json:"op_type" xorm:"'op_type' not null default 0 comment('操作类型') index(idx_log_item_id) TINYINT(1)"`
	Notice     string    `json:"notice" xorm:"'notice' comment('操作原因') MEDIUMTEXT"`
	Content    string    `json:"content" xorm:"'content' comment('详情') TEXT"`
	OpUid      int       `json:"op_uid" xorm:"'op_uid' not null default 0 comment('操作人uid') INT(10)"`
	OpUidType  int       `json:"op_uid_type" xorm:"'op_uid_type' not null default 0 comment('操作人账号类型1-hive 2-用户表') INT(1)"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' not null default CURRENT_TIMESTAMP created comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' not null default CURRENT_TIMESTAMP updated comment('更新时间') DATETIME"`
}

func XormGetTest() {
	TbLog := TbLog{}
	//hasFound, err = t.BaseSession(ctx).Where(whereStr, args...).Desc(t.GetPk()).Get(bean)
	//return impl.GetByWhere(ctx, "biz_id = ? and biz_type = ? ", []interface{}{bizId, bizType})
	args := []interface{}{101, 255}
	//Engine.Alias("o").Where("o.log_item_id = ? and o.source = ? ", args...).Get(&TbLog)
	hasFound, err := Engine.Where("log_item_id = ? and source = ? ", args...).Desc("id").Get(&TbLog)
	fmt.Println(TbLog)
	fmt.Println(hasFound, err)
}

func XormInsertTest() {

	TbLog := []TbLog{
		{LogItemId: 100, Source: 10000},
		{LogItemId: 101, Source: 10001},
	}

	beansInterface := []interface{}{}
	for _, bean := range TbLog {
		beansInterface = append(beansInterface, bean)
	}

	insertId, err := Engine.Insert(beansInterface...)

	fmt.Println(insertId)
	fmt.Println(err)
	fmt.Println(TbLog)

}
