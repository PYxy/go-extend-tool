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