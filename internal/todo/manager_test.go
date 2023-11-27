package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	laptop := New(nil)

	td1 := laptop.CreateToDo("read about local-first")

	phone := New(nil)

	phone.Sync(laptop)

	phone.UpdateText(td1, "read about local-first and crdt")

	longMsg := "read about local-first, crdt and automerge"
	laptop.UpdateText(td1, longMsg)

	phone.Sync(laptop)

	l1 := phone.List()
	require.Len(t, l1, 1)
	l2 := laptop.List()
	require.Len(t, l2, 1)

	assert.Equal(t, longMsg, l1[0].Title)
	assert.Equal(t, longMsg, l2[0].Title)
}
