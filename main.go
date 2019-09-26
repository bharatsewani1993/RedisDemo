package main

import(
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

//CreateConnection test connection with redis server
func CreateConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, err)
	return client
}

//SetValue set key
func SetValue(client *redis.Client, key,value string){
	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key Set Done!")
}

//GetValue get key
func GetValue(client *redis.Client, key string) string {
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("value ==>", val)
	fmt.Println("Get value done.")
	return val
}

func main(){
 fmt.Println("Declare key and value")
 var key,value string = "a","apple" 

 fmt.Println("Start Connection Test")
 connection := CreateConnection()

 fmt.Println("Set key and value")
 SetValue(connection,key,value)

 fmt.Println("Get value")
 GetValue(connection,key)
 
}