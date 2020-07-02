package orm

import "sync"

var instant *MongoOperation
var once sync.Once

// 不要用结构体直接生成实例
//请使用GetInstant()方法
type MongoOperation struct{}

func (this *MongoOperation) GetInstant() *MongoOperation {
	once.Do(func() {
		instant = &MongoOperation{}
	})
	return instant
}

func (this *MongoOperation) findOne() interface{} {

}
