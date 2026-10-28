package task

type User struct {
	ID   int
	Name string
	Age  int
}

func Add(users *[]*User, id int, name string, age int) {
	*users = append(*users, &User{ID: id, Name: name, Age: age})
}

func Update(users []*User, id int, name string, age int) {
	for _, u := range users {
		if u.ID == id {
			if name != "" {
				u.Name = name
			}
			if age > 0 {
				u.Age = age
			}
			return
		}
	}
}

func Delete(users *[]*User, id int) {
	for i, u := range *users {
		if u.ID == id {
			*users = append((*users)[:i], (*users)[i+1:]...)
			return
		}
	}
}
