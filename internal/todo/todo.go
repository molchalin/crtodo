package todo

import (
	"encoding"
	"fmt"
	"strconv"

	"github.com/molchalin/crtodo/internal/operation"
)

type ToDo struct {
	id   operation.DocumentID
	Name Field[*str]
	Done Field[*bol]
}

func Build(id operation.DocumentID, ops []operation.Operation) (ToDo, error) {
	td := ToDo{
		id: id,
	}
	for _, op := range ops {
		var err error
		switch op.Field {
		case "name":
			err = td.Name.apply(op)
		case "done":
			err = td.Done.apply(op)
		}
		if err != nil {
			return ToDo{}, fmt.Errorf("apply operation: %v", err)
		}
	}
	return td, nil
}

type marshaler interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type Field[T marshaler] struct {
	Value         T
	LastOperation operation.OperationID
}

func (f *Field[T]) apply(op operation.Operation) error {
	err := f.Value.UnmarshalText([]byte(op.Value))
	f.LastOperation = op.ID
	return err
}

type str string

func (s *str) MarshalText() ([]byte, error) {
	return []byte(*s), nil
}

func (s *str) UnmarshalText(b []byte) error {
	*s = str(b)
	return nil
}

type bol bool

func (b *bol) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatBool(bool(*b))), nil
}

func (s *bol) UnmarshalText(b []byte) error {
	v, err := strconv.ParseBool(string(b))
	*s = bol(v)
	return err
}
