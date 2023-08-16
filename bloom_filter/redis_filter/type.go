package redis_filter

//https://redis.uptrace.dev/zh/guide/lua-scripting.html#lua-%E5%92%8C-go-%E7%B1%BB%E5%9E%8B
//https://hur.st/bloomfilter/?n=10000000&p=&m=512M&k=6

type RedisFilter interface {
	// Set  设置对应的key
	Set(key string) error
	// Get  获取对应的key 是不是在 bitmap 中对应的位置都存在
	Get(key string) (bool, error)
	// Delete 删除使用的键
	Delete()
	//Expire 刷新过期时间
	Expire()
}

type Encryptor interface {
	Encrypt(string) uint32
}
