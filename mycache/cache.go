package mycache

import "time"

// 全局map
var (
	obj1 map[string]interface{} // 缓存对象,里面存放的是缓存的字符串
	obj2 map[string]int64       // 缓存时间Map,里面存放的是缓存对象的缓存开始时间
	obj3 map[string]int64       // 缓存过期时间,里面存放的是缓存对象的过期时间
)

// Set 设置缓存, timeout = 0 或 < 0 都表示立即过期
func Set(key string, value interface{}, timeout int64) {
	initObj()                     // 初始化
	obj1[key] = value             // 设置缓存对象
	obj2[key] = time.Now().Unix() // 设置缓存开始时间
	obj3[key] = timeout           // 设置缓存超时时间
}

// Get 获取缓存
func Get(key string) interface{} {
	initObj()
	if _, ok := obj1[key]; !ok {
		return nil
	}
	if (time.Now().Unix() - obj2[key]) >= obj3[key] {
		delete(obj1, key)
		delete(obj2, key)
		delete(obj3, key)
		return nil
	}
	return obj1[key]
}

// initObj 初始化对象
func initObj() {
	if obj1 == nil {
		obj1 = make(map[string]interface{})
	}
	if obj2 == nil {
		obj2 = make(map[string]int64)
	}
	if obj3 == nil {
		obj3 = make(map[string]int64)
	}
}
