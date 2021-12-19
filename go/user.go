package main

import "fmt"

type User struct {
	ID   int
	Name string
}

// User 構造体が持つ関数
func (u *User) GetIDName() string {
	return fmt.Sprintf("%d: %s", u.ID, u.Name)
}

// User 構造体が持つ関数
func (u *User) private() string {
	return fmt.Sprintf("This is private method")
}

func main() {
	user := User{ID: 1, Name: "Test User"}
	fmt.Println(user.GetIDName()) // "1: Test User"

	fmt.Println(user.private())
}
