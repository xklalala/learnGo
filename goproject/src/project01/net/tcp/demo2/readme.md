经典案例--海量用户即时通讯

redis：

redis共有0-15号数据库
1.添加 set k v
2.获得 get k
3.切换数据库 select
4.查看数据库有多少个字段 dbsize
5.清空数据库，flush  flushall


redis五大数据类型
    string(字符串)  hash(哈希)    list(列表)  set(集合) zset(有序集合)

字符串：string是redis最基本的类型，一个key对应一个value
string类型是二进制安全的，处普通的字符串外，也可以放图片等数据
redis字符串value最大值为512M
set adress 北京
get adress

set指令：如果存在就是修改，不存在就是添加
del指令：删除
del adress

setex message 10 hello  //设置10秒超时销毁

mset可以同时设置多个k->v
mset k1 v1 k2 v2
mget k1 k2


Hash(哈希)
键值对的集合
hash适合存储对象
hset student name "tom"
hset student age "25"
hgetall student
hdel student

hmset dog color "black" age "4"
hmget dog color age
hlen

hexits dog color


List(列表)
list的本质是一个链表
元素的值可以重复
可以头插或者尾插

lpush city beijing nanjing nanchang tianjin
lrange city 0 -1 //

lpush  左边插入
rpush  右边插入
lrange 【key start stop】， 0代表第一个元素，-1代表列表的最后一个元素
lpop    从列表左边取出一个数据
rpop    从列表右边取出一个数据
del

set(集合)
底层是hashtable数据结构
字符串是无序的，元素值不能重复
sadd emails abc@qq.com abc@sina.com uu@aa.com
srem emails abc@qq.com

smembers 取出所有值
sismember 判断值是否是成员
srem    删除指定值
