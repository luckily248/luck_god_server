package model

type Person struct {
	src       string
	name      string
	inventory []Item
	_class    Class
	effects   []Effect
	as        int //arm strength
	ls        int //leg strength
	rf        int //reflection
	wi        int //wisdom
}
