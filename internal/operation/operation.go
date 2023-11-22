package operation

import "sort"

type (
	OperationID string
	DocumentID  string
)

type Operation struct {
	ID         OperationID
	PrevID     *OperationID
	DocumentID DocumentID

	Field string
	Value string
}

func MergeOperations(ops []Operation) map[DocumentID][]Operation {
	next := make(map[OperationID][]Operation)

	for _, op := range ops {
		if op.PrevID != nil {
			next[*op.PrevID] = append(next[*op.PrevID], op)
		}
	}
	merged := make(map[DocumentID][]Operation)
	for _, op := range ops {
		if op.PrevID != nil {
			continue
		}
		for {
			merged[op.DocumentID] = append(merged[op.DocumentID], op)
			n := next[op.ID]
			if len(n) == 0 {
				break
			}
			sort.Slice(n, func(i, j int) bool {
				return n[i].ID > n[j].ID
			})
			op = n[0]
		}
	}
	return merged
}
