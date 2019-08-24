package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/thomashedd/gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {

		//fmt.Printf("#%-5d %9.9s %.55s\n Created At Time:%v\n Updated At Time:%v\n",
		//	item.Number, item.User.Login, item.Title, item.CreatedAt, item.UpdatedAt)

		fmt.Printf("#%-5d Time:%s\n", item.Number, item.CreatedAt.String())
		fmt.Println(reflect.ValueOf(item.CreatedAt).Kind().String())
		j, err := json.MarshalIndent(item, " ", "")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(j))
		//fmt.Println(time.Now())
	}
}
