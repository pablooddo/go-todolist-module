package todolist

import "fmt"

type TodoList struct {
	lista []*TodoItem
}

func (tl *TodoList) Add(t *TodoItem) {
	tl.lista = append(tl.lista, t)
}

func (tl *TodoList) Remove(index int) {
	tl.lista = append(tl.lista[:index], tl.lista[index+1:]...)
}

type TodoItem struct {
	name        string
	description string
	done        bool
}

func (t *TodoItem) markAsDone() {
	t.done = true
}

func Show(tl TodoList) {
	fmt.Println("")
	fmt.Println("---------------")
	for index, t := range tl.lista {
		fmt.Println(index, t.name, t.done)
	}
}

/*
func main() {
	a := todoItem{
		name:        "item 1",
		description: "description of item 1",
		done:        false,
	}

	b := &todoItem{
		name:        "item 2",
		description: "description of item 2",
		done:        false,
	}

	tl := todoList{}

	show(tl)

	a.markAsDone()
	tl.add(&a)
	show(tl)

	tl.add(b)
	show(tl)

	tl.remove(0)
	show(tl)


	todoMap := make(map[string]*todoList)
	todoMap["PABLO"] = &tl
}
*/
