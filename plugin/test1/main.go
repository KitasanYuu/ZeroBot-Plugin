package main

import (
	"fmt"
	"test1/ID_Search"
	"test1/handle"
)

var USERID = "KitasanYuu"

func main() {
	Users := "https://robertsspaceindustries.com/citizens/" + USERID

	h := handle.UserHandle{}
	Search := ID_Search.NewSearch()
	request, err := ID_Search.NewRequest("GET", Users, ID_Search.UserAgent, &h, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	Search.Request = request
	Search.Visit()
}
