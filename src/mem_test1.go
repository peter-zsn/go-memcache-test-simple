package main


import(
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	server = "192.168.7.250:11211"
)

func main()  {
	fmt.Println("starting...... ")
	mc := memcache.New(server)
	if mc == nil{
		fmt.Println("memcache init errot")
	}
	
	// set
	err := mc.Set(&memcache.Item{Key: "foo", Value:[]byte("zhangshuainan")})
	if err != nil{
		fmt.Println(err)
	}
	// get
	value, err :=mc.Get("foo")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(value.Value), value.Key)
	
	// 可以直接对某个key进行更新 set:  add and replace
	mc.Set(&memcache.Item{Key: "foo", Value:[]byte("zhangshuainan111")})
	value1, err :=mc.Get("foo")
	fmt.Println(string(value1.Value))
	
	// 也可以用replace进行更新  存在直接进行replace， 不存在返回item not stored
	err = mc.Replace(&memcache.Item{Key: "foo", Value:[]byte("big man")})
	if err != nil{
		value, err :=mc.Get("foo")
		if err != nil{
			fmt.Println(err)
		}else{
			fmt.Println(string(value.Value), value.Key)
		}
	}
	err = mc.Replace(&memcache.Item{Key: "foo111", Value:[]byte("big man 111")})
	if err != nil{
		fmt.Println(err, 2222)
	}else {
		value, err :=mc.Get("foo111")
		if err != nil{
			fmt.Println(err, "get foo111")
		}else {
			fmt.Println(string(value.Value), value.Key)
		}
	}
	
	// delete, 不存在的key删除返回错误信息， 存在的key进行删除
	err = mc.Delete("noexists")
	fmt.Println(err, "delete noexists")
	err = mc.Delete("foo")
	if err == nil{
		value, err :=mc.Get("foo")
		if err != nil{
			fmt.Println(err, "get foo")
		}else {
			fmt.Println(string(value.Value), value.Key)
		}
	}
	
	// incrby 对指定的key的value， 增加指定量,  不存在的key直接显示错误, 存在的直接返回增加后的value值
	InValue, err := mc.Increment("noexists", 20)
	if err != nil{
		fmt.Println(err, "Increment no key")
	}else{
		fmt.Println(InValue, "InValue")
		value, err :=mc.Get("noexists")
		if err != nil{
			fmt.Println(err, "get noexists")
		}else {
			fmt.Println(string(value.Value), value.Key)
		}
	}
	mc.Set(&memcache.Item{Key: "infoo", Value:[]byte("1")})
	InValue, err = mc.Increment("infoo", 20)
	if err != nil{
		fmt.Println(err, "Increment no key")
	}else{
		fmt.Println(InValue, "InValue")
		value, err :=mc.Get("infoo")
		if err != nil{
			fmt.Println(err, "get infoo")
		}else {
			fmt.Println(string(value.Value), value.Key)
		}
	}
	
	//decrby 同incrby  减少制定的量
	InValue, err = mc.Decrement("infoo", 10)
	fmt.Println(InValue)
}
