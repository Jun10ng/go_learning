package github

import (
	"fmt"
	"log"
	"time"
)

// items 是所有issue
func searchByDate(items []*Issue)  {
	now := time.Now()
	beforeOneDay := now.AddDate(0,0,-1)
	beforeOneMonth := now.AddDate(0,-1,0)
	beforeOneYear := now.AddDate(-1,0,0)
	// 一天前发的
	postDay := make([]*Issue,0)
	// 一月前
	postMonth := make([]*Issue,0)
	// 一年前发
	postYear := make([]*Issue,0)
	for _,item :=range items{
		if item.CreatedAt.Before(beforeOneDay){
			postDay=append(postDay,item)
		}else
		if item.CreatedAt.Before(beforeOneMonth){
			postMonth=append(postMonth,item)
		}else
		if item.CreatedAt.Before(beforeOneYear){
			postYear=append(postYear,item)
		}
	}
	log.Printf("一天前有%d\n",len(postDay))
	for _,e:= range postDay {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			e.Number, e.User.Login, e.Title)
	}
	log.Printf("一月前有%d\n",len(postMonth))
	for _,e:= range postMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			e.Number, e.User.Login,e.Title)
	}
	log.Printf("一年前有%d\n",len(postYear))
	for _,e:= range postYear{
		fmt.Printf("#%-5d %9.9s %.55s\n",
			e.Number, e.User.Login, e.Title)
	}
}
