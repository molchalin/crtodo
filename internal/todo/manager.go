package todo

import (
	"github.com/molchalin/crtodo/internal/operation"
)

type Manager struct {
	m   map[operation.DocumentID]*ToDo
	ops []operation.Operation
}

type ToDoView struct {
	Title string
	Done  bool
}

func New(ops []operation.Operation) *Manager {
	merged := operation.MergeOperations(ops)
	m := make(map[operation.DocumentID]*ToDo, len(merged))
	for dID, tdOps := range merged {
		m[dID] = Build(tdOps)
	}
	return &Manager{
		m:   m,
		ops: ops,
	}
}

func (m *Manager) Sync(m2 *Manager) {
	var ops []operation.Operation
	ops = append(ops, m.ops...)
	ops = append(ops, m2.ops...)
	ops2 := make([]operation.Operation, len(ops))
	copy(ops2, ops)

	newM := New(ops)
	m.m = newM.m
	m.ops = newM.ops

	newM2 := New(ops2)
	m2.m = newM2.m
	m2.ops = newM2.ops
}

func (m *Manager) CreateToDo(text string) operation.DocumentID {
	dID := operation.NewDocumentID()
	ops := []operation.Operation{
		operation.New(dID, "title", text, 0),
		operation.New(dID, "done", "false", 0),
	}
	m.m[dID] = Build(ops)
	m.ops = append(m.ops, ops...)
	return dID
}

func (m *Manager) UpdateText(dID operation.DocumentID, text string) {
	m.ops = append(m.ops, m.m[dID].updateText(text))
}

func (m *Manager) Done(dID operation.DocumentID) {
	m.ops = append(m.ops, m.m[dID].done())
}

func (m *Manager) List() []ToDoView {
	res := make([]ToDoView, 0, len(m.m))
	for _, v := range m.m {
		res = append(res, ToDoView{
			Title: v.Title,
			Done:  v.Done,
		})
	}
	return res
}
