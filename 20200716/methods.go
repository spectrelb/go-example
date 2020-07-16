package main

import "fmt"

type Role struct {
	Name string
	Ability string
	Level int
	Kill float64
}

func (r Role) Kungfu() {
	fmt.Printf("我是:%s，我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.Kill)
}

func (r *Role) Kungfu1() {
	fmt.Printf("我是:%s，我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.Kill)
}

func main() {
	role := Role{"1", "2", 3, 4.0}
	role.Kungfu()

	role1 := &Role{"1", "2", 3, 4.0}
	role1.Kungfu1()
}

