package operation

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type (
	OperationID int64
	DocumentID  string
)

type Operation struct {
	ID         OperationID
	PrevID     OperationID
	DocumentID DocumentID

	Field string
	Value string
}

func New(documentID DocumentID, field string, value string, prevID OperationID) Operation {
	return Operation{
		ID:         OperationID(time.Now().UnixNano()),
		PrevID:     prevID,
		DocumentID: documentID,
		Field:      field,
		Value:      value,
	}
}

func NewDocumentID() DocumentID {
	return DocumentID(uuid.NewString())
}

func cmp(l, r Operation) int {
	switch {
	case l.ID < r.ID:
		return 1
	case l.ID > r.ID:
		return -1
	}
	return 0
}

func eq(l, r Operation) bool {
	return l.ID == r.ID
}

func MergeOperations(ops []Operation) map[DocumentID][]Operation {
	next := make(map[OperationID][]Operation)

	slices.SortFunc(ops, cmp)
	ops = slices.CompactFunc(ops, eq)

	for _, op := range ops {
		if op.PrevID != 0 {
			next[op.PrevID] = append(next[op.PrevID], op)
		}
	}
	merged := make(map[DocumentID][]Operation)
	for _, op := range ops {
		if op.PrevID != 0 {
			continue
		}
		for {
			merged[op.DocumentID] = append(merged[op.DocumentID], op)
			n := next[op.ID]
			if len(n) == 0 {
				break
			}
			op = n[0]
		}
	}
	return merged
}
