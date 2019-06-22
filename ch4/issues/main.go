package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/thomashedd/gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Now())
	//month
	mL, err := time.ParseDuration("720h0m0s") // don't understaand why I need this?
	if err != nil {
		log.Fatal(err)
	}
	//year
	yL, err := time.ParseDuration("8760h0m0s")
	if err != nil {
		log.Fatal(err)
	}

	// alternative method is just compare UTC time of time.Now() with time of issue creation?

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		category := "Less than a year old"
		if time.Now().Sub(item.CreatedAt) < mL {
			category = "Less than a month old"
		}
		if time.Now().Sub(item.CreatedAt) > yL {
			category = "more than a year old"
		}
		fmt.Printf("#%-5d %9.9s %.55s\n %s\n Time:%v\n",
			item.Number, item.User.Login, item.Title, category, item.CreatedAt)
	}
}
