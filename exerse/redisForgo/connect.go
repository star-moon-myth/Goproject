package main

import (
	"context"
	"fmt"
	_"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx=context.Background()
func initRedis()error{
	rdb=redis.NewClient(&redis.Options{
		Addr: "redis-18612.c90.us-east-1-3.ec2.redns.redis-cloud.com:18612",
		Username: "default",
		Password: "d7B8UXU2wR8inNk0yPeFsaouPP2Wo1q5",
		DB: 0,
		PoolSize:     10,                    // 连接池最大连接数
		MinIdleConns: 5,                     // 最小空闲连接数
		PoolTimeout:  30 * time.Second,      // 从池中获取连接的超时时间
		DialTimeout:  5 * time.Second,       // 建立新连接的超时时间
		ReadTimeout:  3 * time.Second,       // 读取操作的超时时间
		WriteTimeout: 3 * time.Second,       // 写入操作的超时时间
	})
	pong,err:=rdb.Ping(ctx).Result()
	if err!=nil{
		fmt.Println("failed to connect redis :",err)
		return err
	}
	fmt.Println("Successfully connected to Redis,Ping Response:",pong)
	return nil 
}
func main(){
	err:=initRedis()
	if err!=nil{
		fmt.Println("init error:",err)
		return
	}
	// BasicOperation()
	// hashOperation()
	// listOpraton()
	// SetOperation()
	// Zsetoperation()
	// ExpirationOpetation()
	// examplePipeline()
	// transactionOperation()
	// transactionOperation()
	pubsub()
}

// func BasicOperation(){
// 	fmt.Println("=========String类型操作==============")
// 	key:="user:1:name"
// 	value:="wangming"
// 	expiration:=time.Minute*5
// 	err:=rdb.Set(ctx,key,value,expiration).Err()
// 	if err!=nil{
// 		fmt.Println("failed to set key:",err)
// 		return 
// 	}
// 	fmt.Printf("Successfully set key:%s,value:%s,expiration:%v\n",key,value,expiration)

// 	getname,err:=rdb.Get(ctx,key).Result()
// 	if err!=nil{
// 		if err==redis.Nil{
// 			fmt.Println("key does not exist")
// 		}else{
// 			fmt.Println("failed to get key:",err)
// 			return
// 		}
// 	}
// 	fmt.Println("the key is :",getname)

// 	deleteuser,err:=rdb.Del(ctx,key,"persistent_key").Result()
// 	if err!=nil{
// 		fmt.Println("failed to delete:",err)
// 	}else{
// 		fmt.Printf("Successfully deleted %d key:",deleteuser)
// 	}

// 	_,err=rdb.Get(ctx,key).Result()
// 	if err==redis.Nil{
// 		fmt.Println("key already deleted")
// 	}

// 	countekeys:="page_views"
// 	currentviews,err:=rdb.Incr(ctx,countekeys).Result()
// 	if err!=nil{
// 		fmt.Println("failed to increment:",err)
// 		return 
// 	}else{
// 		fmt.Println("Successfully incremented key:",countekeys,currentviews)
// 	}
// 	currentviews,_=rdb.Incr(ctx,countekeys).Result()
// 	fmt.Println("Successfully incremented key:",countekeys,currentviews)

// 	currentviews,_=rdb.IncrBy(ctx,countekeys,10).Result()
// 	fmt.Println("Successfully incremented key:",countekeys,currentviews)

// 	currentviews,_=rdb.Decr(ctx,countekeys).Result()
// 	fmt.Println("Successfully incremented key:",countekeys,currentviews)

// 	rdb.Del(ctx,countekeys)
// }

// func hashOperation(){
// 	fmt.Println("========hash==========")
// 	hashkey:="user:2:name"
// 	err:=rdb.HSet(ctx,hashkey,"name","Alice","age",34,"email","2198853596@qq.com").Err()
// 	if err!=nil{
// 		fmt.Println("Hset error:",err)		
// 		return
// 	}else{
// 		fmt.Println("🎉had hset :",hashkey)
// 	}

// 	name,err:=rdb.HGet(ctx,hashkey,"name").Result()
// 	if err==redis.Nil{
// 		fmt.Println("The field not exists:",hashkey)
// 	}else if err!=nil{
// 		fmt.Println("Hget error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully hget:",name)
// 	}

// 	allfields,err:=rdb.HGetAll(ctx,hashkey).Result()
// 	if err!=nil{
// 		fmt.Println("Hgetall error:",err)
// 	}else{
// 		fmt.Println("📖All fields:------")
// 		for field,value:=range allfields{
// 			fmt.Println(field,value)
// 		}
// 		fmt.Println("🔚----==----🔚")
// 	}
	
// 	deletecount,err:=rdb.HDel(ctx,hashkey,"email").Result()
// 	if err!=nil{
// 		fmt.Println("Hdel error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully：",deletecount)
// 	}

// 	_,err=rdb.HGet(ctx,hashkey,"email").Result()
// 	if err==redis.Nil{
// 		fmt.Println("Its OK to delete email")
// 	}
// 	rdb.Del(ctx,hashkey)
// }

// func listOpraton(){
// 	listkey:="tasks:pending"
// 	err:=rdb.LPush(ctx,listkey,"task4","task3","task2","task1").Err()
// 	if err!=nil{
// 		fmt.Println("Lpush error:",err)
// 	}else{
// 		fmt.Println("🎉success Lpush!")
// 	}
// 	err=rdb.RPush(ctx,listkey,"task5","task6").Err()
// 	if err!=nil{
// 		fmt.Println("Rpush error:",err)
// 	}else{
// 		fmt.Println("🎉success Rpush!")
// 	}
// 	len,err:=rdb.LLen(ctx,listkey).Result()
// 	if err!=nil{
// 		fmt.Println("LLen error:",err)
// 	}else{
// 		fmt.Println("The length of listlen:",len)
// 	}
// 	lfield,err:=rdb.LRange(ctx,listkey,0,-1).Result()
// 	if err!=nil{
// 		fmt.Println("Lrange error:",err)
// 	}else{
// 		fmt.Printf("🎉success Lrange,%s:%v\n",listkey,lfield)
// 	}

// 	task,err:=rdb.LPop(ctx,listkey).Result()
// 	if err!=nil{
// 		fmt.Println("LPop error:",err)
// 	}else{
// 		fmt.Println("🎉success LPop:",task)
// 	}
// 	task,err=rdb.RPop(ctx,listkey).Result()
// 	if err!=nil{
// 		fmt.Println("RPop error:",err)
// 	}else{
// 		fmt.Println("🎉success RPop:",task)
// 	}
// 	rdb.Del(ctx,listkey)
// }
// func SetOperation(){
// 	setkey1:="user:1:tag"
// 	setkey2:="user:2:tag"
// 	addcount,err:=rdb.SAdd(ctx,setkey1,"linux","apple","huawei").Result()
// 	if err!=nil{
// 		fmt.Println("Sadd error:",err)
// 	}else{
// 		fmt.Println("🎉The first time successfully Sadd:",addcount)
// 	}

// 	addcount,err=rdb.SAdd(ctx,setkey1,"vivo","apple","huawei").Result()
// 	if err!=nil{
// 		fmt.Println("Sadd error:",err)
// 	}else{
// 		fmt.Println("🎉The second time successfully Sadd:",addcount)
// 	}
// 	rdb.SAdd(ctx,setkey2,"python","java","golang")
// 	member,err:=rdb.SMembers(ctx,setkey2).Result()
// 	if err!=nil{
// 		fmt.Println("Smembers error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully Smembers:",member)
// 	}
// 	ismember,err:=rdb.SIsMember(ctx,setkey2,"golang").Result()
// 	if err!=nil{
// 		fmt.Println("Sismember error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully Sismember:",ismember)
// 	}
// 	ismember,_=rdb.SIsMember(ctx,setkey2,"haha").Result()
// 	fmt.Println("🎉Successfully Sismember:",ismember)

// 	scount,err:=rdb.SCard(ctx,setkey2).Result()
// 	if err!=nil{
// 		fmt.Println("Scard error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully Scard:",scount)
// 	}

// 	remove,err:=rdb.SRem(ctx,setkey2,"java","golang","heeh").Result()
// 	if err!=nil{
// 		fmt.Println("Srem error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully Srem:",remove)
// 	}

// 	members,_:=rdb.SMembers(ctx,setkey2).Result()
// 	fmt.Println("🎉Successfully Srem:",members)

// 	difftags,_:=rdb.SDiff(ctx,setkey1,setkey2).Result()
// 	fmt.Println("🎉Successfully Sdiff:",difftags)

// 	rdb.Del(ctx,setkey1,setkey2)
// }

// func Zsetoperation(){
// 	fmt.Println("=========zset==========")
// 	zsetkey:="game:leaderboard"
// 	memberstoadd:=[]redis.Z{
// 		{Score: 98.5,Member: "player1"},
// 		{Score: 100.0,Member: "player2"},
// 		{Score: 75.2,Member: "player3"},
// 	}
// 	addedcount,err:=rdb.ZAdd(ctx,zsetkey,memberstoadd ...).Result()
// 	if err!=nil{
// 		fmt.Println("Zadd error:",err)
// 	}else{
// 		fmt.Printf("🎉Successfully Zadd %d members\n",addedcount)
// 	}
// 	memberstoupdate:=[]redis.Z{
// 		{Score: 99,Member: "player1"},
// 		{Score: 88.8,Member: "player4"},
		
// 	}
// 	addedcount,_=rdb.ZAdd(ctx,zsetkey,memberstoupdate...).Result()
// 	fmt.Printf("🎉Successfully Zadd %d members\n",addedcount)

// 	playerAsc,err:=rdb.ZRange(ctx,zsetkey,0,-1).Result()
// 	if err!=nil{
// 		fmt.Println("Zrange error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully zrange:",playerAsc)
// 	}
// 	playerdesc,err:=rdb.ZRevRange(ctx,zsetkey,0,-1).Result()
// 	if err!=nil{
// 		fmt.Println("Zrevrange error:",err)
// 	}
// 	fmt.Println("🎉Successfully zrevrange:",playerdesc)

// 	playerAscwithScore,err:=rdb.ZRangeWithScores(ctx,zsetkey,0,-1).Result()
// 	if err!=nil{
// 		fmt.Println("ZrangeWithScores error:",err)
// 	}else{
// 		fmt.Println("🎉Successfully zrangeWithScores===========")
// 		for _,pp:=range playerAscwithScore{
// 			fmt.Printf("%s:%.0f\n",pp.Member,pp.Score)
// 		}
// 	}

// 	score,err:=rdb.ZScore(ctx,zsetkey,"player2").Result()
// 	if err==redis.Nil{
// 		fmt.Println("The member player2 not exists",)
// 	}else if err!=nil{
// 		fmt.Println("Zscore error:",err)
// 	}else{
// 		fmt.Printf("🎉Successfully Zscore:player2 score is %.0f\n",score)
// 	}
// 	zcount,_:=rdb.ZCard(ctx,zsetkey).Result()
// 	fmt.Printf("🎉Successfully Zcard:the zset has %d members\n",zcount)

// 	remove,_:=rdb.ZRem(ctx,zsetkey,"player3").Result()
// 	fmt.Printf("🎉Successfully Zrem:remove %d members\n",remove)

// 	rdb.Del(ctx,zsetkey)
// }

// func ExpirationOpetation(){
// 	fmt.Println("==========ExpirationOpetation=========")
// 	tempkey:="session:user123"
// 	tempval:="some session data"
// 	rdb.Set(ctx,tempkey,tempval,time.Second*10).Err()
// 	fmt.Println("set expiration session successfully:",tempkey)

// 	anotherkey:="session:new"
// 	rdb.Set(ctx,anotherkey,"new data",0).Err()
// 	result,err:=rdb.Expire(ctx,anotherkey,10*time.Second).Result()
// 	if err!=nil{
// 		fmt.Println("Expire error:",err)
// 	}else if result{
// 		fmt.Printf("🎉Successfully Expire,%s:10*time.Second\n",anotherkey,)
// 	}else{
// 		fmt.Println("maybe the key is not existed!❌")
// 	}

// 	ttl,err:=rdb.TTL(ctx,tempkey).Result()
// 	if err!=nil{
// 		fmt.Println("TTL error:",err)
// 	}else{
// 		fmt.Printf("🎉Successfully TTL,%s:%.0f\n",tempkey,ttl.Seconds())
// 	}
// 	ttl,_=rdb.TTL(ctx,anotherkey).Result()
// 	fmt.Printf("🎉Successfully TTL,%s:%.0f\n",tempkey,ttl.Seconds())

// 	persisted,err:=rdb.Persist(ctx,anotherkey).Result()
// 	if err!=nil{
// 		fmt.Println("Persist error:",err)
// 	}else if !persisted{
// 		fmt.Println("Persist error:",err)
// 	}else{
// 		fmt.Printf("🎉Successfully Persist a key,%s\n",anotherkey)
// 	}
// 	rdb.Del(ctx,anotherkey,tempkey)
// }

// func examplePipeline()  {
// 	fmt.Println("==========examplePipeline=========")
// 	pipe:=rdb.Pipeline()
// 	pipe.Incr(ctx,"pipelinekey")
// 	pipe.Set(ctx,"pipelinekey2","key2hello",time.Minute)
// 	pipe.Expire(ctx,"pipelinekey2",time.Second*30)
// 	pipe.Get(ctx,"pipelinekey2")

// 	cmds,err:=pipe.Exec(ctx)
// 	if err!=nil&&err==redis.Nil{
// 		fmt.Println("Exec error:",err)
// 	}else{
// 		fmt.Println("🎉Success for pipeline!")
// 	}
// 	if len(cmds)>1{
// 		if err:=cmds[0].Err();err!=nil{
// 			fmt.Printf("The first cmd have error %v!",err)
// 		}else{
// 			if intcmd,ok:=cmds[0].(*redis.IntCmd);ok{
// 				fmt.Printf("The first cmd result is %d!",intcmd.Val())
// 			}
// 		}
// 	}
// 	rdb.Del(ctx,"pipelinekey","pipelinekey2")
// }

// func transactionOperation(){
// 	fmt.Println("==========transactionOperation=========")
// 	key1:="tx_key1"
// 	key2:="tx_key2"
// 	rdb.Set(ctx,key1,"10",0)
// 	rdb.Set(ctx,key2,"hello",0)
// 	pipe:=rdb.TxPipeline()
// 	pipe.Incr(ctx,key1)
// 	pipe.Set(ctx,key2,"hehe",0)
// 	_,err:=pipe.Exec(ctx)

// 	if err!=nil&&err!=redis.Nil{
// 		fmt.Println("Exec error:",err)
// 		get1,_:=rdb.Get(ctx,key1).Result()
// 		get2,_:=rdb.Get(ctx,key2).Result()
// 		fmt.Println("get1:",get1,"get2:",get2)
// 	}else{
// 		fmt.Println("🎉Success for transaction!")
// 		get1,_:=rdb.Get(ctx,key1).Result()
// 		get2,_:=rdb.Get(ctx,key2).Result()
// 		fmt.Println("get1:",get1,"get2:",get2)
// 	}

// 	fmt.Println("=========乐观锁watch=========")
// 	accountKey := "user:3:balance"
// 	rdb.Set(ctx, accountKey, "100", 0)


// 	deductBalance := func(amount int) error {
//         err := rdb.Watch(ctx, func(tx *redis.Tx) error {
//             balanceStr, err := tx.Get(ctx, accountKey).Result()
//             if err != nil && err != redis.Nil {
//                 return err
//             }
//             balance, _ := strconv.Atoi(balanceStr)
//             if balance < amount {
//                 return fmt.Errorf("余额不足") 
//             }
//             _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
//                 pipe.DecrBy(ctx, accountKey, int64(amount))
//                 return nil // Pipelined 内部通常返回 nil
//             })
//             return err // 如果 DecrBy 出错，返回错误
//         }, accountKey) // 指定要 Watch 的 key(s)

//         if err != nil {
//             if err == redis.TxFailedErr {
//                 fmt.Println("❌ 乐观锁失败：余额在操作期间被修改，请重试！")
//             } else {
//                 fmt.Printf("❌ 事务处理失败: %v\n", err)
//             }
//             return err
//         }
//         fmt.Printf("✅ 成功扣除 %d 余额\n", amount)
//         return nil
//     }

// 	go func() { 
// 	time.Sleep(58 * time.Millisecond)
// 	rdb.Set(ctx, accountKey, "50", 0)
// 	fmt.Println("后台修改了余额!") }()

//     time.Sleep(50 * time.Millisecond) // 确保 Watch 先执行

// 	deductBalance(20) // 第一次通常成功 (如果没有并发修改)
//     newBalance, _ := rdb.Get(ctx, accountKey).Result()
//     fmt.Printf("  当前余额: %s\n", newBalance) // 应为 80 (如果成功)
// 	rdb.Del(ctx, key1, key2, accountKey)
// }

func pubsub(){
	fmt.Println("=======pubsub==========")
	channel:="my_chat_channel"
	pubsub:=rdb.Subscribe(ctx,channel)
	ch:=pubsub.Channel()
	defer pubsub.Close()
	go func ()  {
		fmt.Println("订阅者监听中..........")
		for msg:=range ch{
			fmt.Printf("监听到消息来自%s的频道:%s\n",msg.Channel,msg.Payload)
			if msg.Payload=="quit"{
				fmt.Println("Have reived the massage 'quit' ")
				pubsub.Unsubscribe(ctx,channel)
				break
			}
		}
		fmt.Println("监听结束!")
	}()
	time.Sleep(1*time.Second)
	err:=rdb.Publish(ctx,channel,"hello redis,🎉").Err()
	if err!=nil{
		fmt.Println("Publish error:",err)
	}else{
		fmt.Println("🎉Success for publish!")
	}
	err=rdb.Publish(ctx,channel,"Second message hello redis,🎉").Err()
	if err!=nil{
		fmt.Println("Publish error:",err)
	}

	time.Sleep(1*time.Second)
	rdb.Publish(ctx,channel,"quit")
	fmt.Println("发布了quit消息！")

	time.Sleep(2*time.Second)
}