package todo

import (
	"github.com/molchalin/crtodo/internal/operation"
)

type ToDo struct {
	id operation.DocumentID

	titleLastOp operation.OperationID
	Title       string
	doneLastOp  operation.OperationID
	Done        bool
}

func Build(ops []operation.Operation) *ToDo {
	td := &ToDo{}
	if len(ops) == 0 {
		panic("elements expected")
	}
	for _, op := range ops {
		td.id = op.DocumentID
		switch op.Field {
		case "title":
			td.Title = op.Value
			td.titleLastOp = op.ID
		case "done":
			td.Done = op.Value != ""
			td.doneLastOp = op.ID
		}
	}
	return td
}

func (t *ToDo) updateText(text string) operation.Operation {
	op := operation.New(t.id, "title", text, t.titleLastOp)
	t.Title = text
	t.titleLastOp = op.ID
	return op
}

func (t *ToDo) done() operation.Operation {
	op := operation.New(t.id, "done", "true", t.doneLastOp)
	t.Done = true
	t.doneLastOp = op.ID
	return op
}
