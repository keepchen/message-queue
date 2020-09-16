package queue

import (
	"fmt"
	"log"
	"sync"
)

type queue struct {
	size uint
	list *doubleLinkedList
}

//DBSet 数据库设置
type DBSet struct {
	mu sync.Mutex
	db *map[string]*queue
}

//Instance 实例
type Instance struct {
	dbName string
	debug  bool
}

var pool *DBSet

//GetDBInstance 获取实例
func GetDBInstance(dbName string) *Instance {
	if len(dbName) == 0 {
		dbName = "db0"
	}
	pool.mu.Lock()
	defer pool.mu.Unlock()

	_, ok := (*pool.db)[dbName]
	if !ok {
		//当前库不存在，先创建
		(*pool.db)[dbName] = &queue{}
	}
	return &Instance{dbName: dbName, debug: true}
}

//SetDebugMode 设置调试模式
func (instance *Instance) SetDebugMode(isDebug bool) *Instance {
	instance.debug = isDebug
	return instance
}

//LPush 向队列头部添加节点
func (instance *Instance) LPush(val interface{}) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	list, err := (*pool.db)[instance.dbName].list.lPush(val)
	if err != nil {
		if instance.debug {
			log.Printf("db {%s} [LPush] error: %#v\n", instance.dbName, err)
		}
		return err
	}

	//当前队列长度加1
	(*pool.db)[instance.dbName].size++
	(*pool.db)[instance.dbName].list = list
	return nil
}

//LPop 从队列头部返回节点值
func (instance *Instance) LPop() (interface{}, error) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	list, val, err := (*pool.db)[instance.dbName].list.lPop()
	if err != nil {
		if instance.debug {
			log.Printf("db {%s} [LPop] error: %#v\n", instance.dbName, err)
		}
		return nil, err
	}

	//当前队列长度减1
	(*pool.db)[instance.dbName].size--
	(*pool.db)[instance.dbName].list = list
	return val, nil
}

//RPush 向队列尾部追加节点
func (instance *Instance) RPush(val interface{}) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	list, err := (*pool.db)[instance.dbName].list.rPush(val)
	if err != nil {
		if instance.debug {
			log.Printf("db {%s} [RPush] error: %#v\n", instance.dbName, err)
		}
		return err
	}

	//当前队列长度加1
	(*pool.db)[instance.dbName].size++
	(*pool.db)[instance.dbName].list = list
	return nil
}

//RPop 从队列尾部弹出节点值
func (instance *Instance) RPop() (interface{}, error) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	list, val, err := (*pool.db)[instance.dbName].list.rPop()
	if err != nil {
		if instance.debug {
			log.Printf("db {%s} [RPop] error: %#v\n", instance.dbName, err)
		}
		return nil, err
	}

	//当前队列长度减1
	(*pool.db)[instance.dbName].size--
	(*pool.db)[instance.dbName].list = list
	return val, nil
}

//GetDBList 获取当前数据库
func (instance *Instance) GetDBList() []string {
	var result []string
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if len(*pool.db) == 0 {
		return result
	}
	for dbName := range *pool.db {
		result = append(result, dbName)
	}
	return result
}

//GetSize 获取当前队列长度
func (instance *Instance) GetSize() uint {
	var length uint
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if list, ok := (*pool.db)[instance.dbName]; ok {
		length = list.size
	}
	return length
}

//DisplayQueue 打印队列信息
func (instance *Instance) DisplayQueue() {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	(*pool.db)[instance.dbName].list.displayQueue()

	return
}

//FlushDB 删除数据库
//
//警告：此操作后，当前实例将被销毁，若继续调用，将会panic
//
//如：
//
//instance := queue.GetDBInstance("db1")
//
//instance.FlushDB()
//
//instance.LPush(1)  # <- 此处将会panic
func (instance *Instance) FlushDB() error {
	pool.mu.Lock()
	defer func() {
		pool.mu.Unlock()
		//销毁实例
		instance = (*Instance)(nil)
	}()

	if _, ok := (*pool.db)[instance.dbName]; ok {
		if instance.debug {
			log.Printf("db {%s} will be delete\n", instance.dbName)
		}
		delete(*pool.db, instance.dbName)
		return nil
	}
	return fmt.Errorf("db {%s} not exist", instance.dbName)
}
