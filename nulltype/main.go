package main

import (
	"fmt"

	nulltype "github.com/mattn/go-nulltype"
)

// User user
type User struct {
	Name nulltype.NullString `json:"name"`
}

// User2 user2
type User2 struct {
	Name string
}

// Post post
type Post struct {
	ID      int64 `db:"post_id"`
	Created int64
	Title   string              `db:"size:50"`
	Body    nulltype.NullString `db:"body,size:1024"`
}

func main() {
	var user User
	fmt.Println(user.Name)
	if user.Name.Valid() == false {
		fmt.Println("   songhq")
	}

	var user2 User2
	fmt.Println(user2.Name)
	if user2.Name == "" {
		fmt.Println("   songhq2")
	}

	user.Name.Set("Bob")
	fmt.Println(user.Name)

	fmt.Println(user.Name.StringValue() == "Bob")

	user.Name.Reset()
	fmt.Println(user.Name)
}
