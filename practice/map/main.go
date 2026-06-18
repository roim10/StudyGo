package main

import (
	"fmt"
	"maps"
)

func main() {
	users := make(map[string]int)
	users["admin"] = 5
	users["vise-admin"] = 3
	users["users"] = 2
	users["bots"] = 5
	users["guest"] = 1
	if val, ok := users["admin"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println(nil)
	}
	users["bots"] = 4
	delete(users, "guest")
	fmt.Println(users["guest"])
	for _, v := range users {
		fmt.Println(v)
	}
	backupUsers := map[string]int{
		"admin":      5,
		"vise-sdmin": 3,
		"users":      2,
		"bots":       5,
		"guest":      1,
	}
	fmt.Println(maps.Equal(backupUsers, users))
}
