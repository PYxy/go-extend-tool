# go-extend-tool
go 扩展工具

slice 文件夹

```Bash
#index.go
Index         查询命中的第一个索引
IndexFunc
IndexAll       查询命中的所有索引
IndexAllFunc

#add.go
AddIndex       指定位置添加元素

#delete.go
DeleteByIndex      删除指定位置的多个元素
DeleteByValue      删除指定值的多个元素
DeleteByValueFunc  删除指定值的多个元素的补充


#aggregate.go
Max                  获取切片中的最大值以及对应索引
Min                  获取切片中的最小值以及对应索引
Sum                  切片数据求和
Intersection         可比较类型(交集)
IntersectionFunc     交集
Union                可比较类型(并集)
UnionFunc            并集
DifferenceSet        可比较类型(差集)
DifferenceSetFunc    差集
#unique.go
Unique          切片去重(保留出现的第一个)
UniqueFuncFirst 切片去重(保留出现的第一个)
UniqueFuncLast  切片去重(保留出现的最后一个)

#reverse.go
Reverse        创建一个新的切片,进行翻转
ReverseOnSelf  在 src 上进行翻转

#toMap.go
ToMap          切片转map
ToMapFunc      
```

list 文件夹

```Bash
#linked_list go
NewLinkedList      创建空链表
NewLinkedBySlice   传入切片,创建链表
Get                返回对应下标的元素
Append             在末尾追加元素
Add                在特定下标处增加一个新元素
Set                重置 index 位置的值
Delete             删除目标元素的位置,并且返回该位置的值
Len                返回长度
Cap                返回容量
Range              遍历 List 的所有元素
AsSliceAsc         将 List 转化为一个切片(顺序)
AsSliceRev         将 List 转化为一个切片(逆序)
```


bloom_filter/redis_filter 文件夹
```Bash
#redis.go
NewRedisBloomFilter  创建一个基于redis的布隆过滤器
Set                  根据key 设置状态
Get                  根据key 获取状态
Delete               删除布隆过滤器 生成的键
Expire               给布隆过滤器 生成的键 设置过期时间
```