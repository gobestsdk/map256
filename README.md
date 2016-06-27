# map256
map256是一种利用256叉树实现的树索引

## 用法
+ func GET(key)
+ func PUT(key)
+ func DET(key)
> same as PUT(nil)
+ func EXIST(key)
> same as (GET(key)!=nil)
```