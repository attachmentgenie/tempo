package backend

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	testTenantID = "fake"
)

func TestBlockMeta(t *testing.T) {
	id := uuid.New()
	b := NewBlockMeta(testTenantID, id)

	assert.Equal(t, id, b.BlockID)
	assert.Equal(t, testTenantID, b.TenantID)

	randID1 := make([]byte, 10)
	randID2 := make([]byte, 10)

	rand.Read(randID1)
	rand.Read(randID2)

	assert.Equal(t, b.StartTime, b.EndTime)

	b.ObjectAdded(randID1)
	b.ObjectAdded(randID2)
	assert.True(t, b.EndTime.After(b.StartTime))
	assert.Equal(t, 1, bytes.Compare(b.MaxID, b.MinID))

	assert.Equal(t, 2, b.TotalObjects)
}
