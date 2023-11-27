package operation

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingle(t *testing.T) {
	dID := DocumentID(uuid.NewString())
	opID1 := OperationID(1)
	opID21 := OperationID(3)
	opID22 := OperationID(2)
	opID3 := OperationID(4)

	ops := []Operation{
		{
			ID:         opID1,
			DocumentID: dID,
			PrevID:     0,
		},
		{
			ID:         opID21,
			DocumentID: dID,
			PrevID:     opID1,
		},
		{
			ID:         opID22,
			DocumentID: dID,
			PrevID:     opID1,
		},
		{
			ID:         opID3,
			DocumentID: dID,
			PrevID:     opID21,
		},
	}

	merged := MergeOperations(ops)
	require.Len(t, merged, 1)

	td, ok := merged[dID]
	require.True(t, ok)

	require.Len(t, td, 3)

	assert.Equal(t, opID1, td[0].ID)
	assert.Equal(t, opID21, td[1].ID)
	assert.Equal(t, opID3, td[2].ID)
}
