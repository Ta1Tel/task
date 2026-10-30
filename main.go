package main

import (
	"fmt"
	task "task/crud"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	var users []*task.User

	task.Add(&users, 1, "Иванв", 25)
	task.Add(&users, 2, "Мария", 30)
	task.Add(&users, 3, "Пётр", 22)

	task.Update(users, 1, "Иван Петров", 26)
	task.Delete(&users, 2)

	for _, u := range users {
		fmt.Printf("ID: %d | Имя: %s | Возраст: %d\n", u.ID, u.Name, u.Age)
	}
}
