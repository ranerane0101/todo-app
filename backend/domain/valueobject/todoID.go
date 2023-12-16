package valueobject

type TodoID struct {
	id string
}

func NewTodoID(id string) TodoID {
	return TodoID{id: id}
}

func (id TodoID) String() string {
	return id.id
}

func (id TodoID) Value() string {
	return id.id
}
