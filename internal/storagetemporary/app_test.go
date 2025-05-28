package storagetemporary_test

import (
	"testing"
	"time"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/storagetemporary"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	id, name string
}

func TestStorageTemporary(t *testing.T) {
	st, err := storagetemporary.New(
		t.Context(),
		storagetemporary.WithTimeTick[TestData](3),
		storagetemporary.WithTimeToLive[TestData](3))
	assert.NoError(t, err)

	st.SetData("hash_one", TestData{
		id:   "4av2h4b3",
		name: "One element description",
	})
	st.SetData("hash_two", TestData{
		id:   "72gd8824",
		name: "Two element description",
	})
	assert.Equal(t, st.DataSize(), 2)

	elem, ok := st.GetData("hash_one")
	assert.True(t, ok)
	assert.Equal(t, elem.id, "4av2h4b3")

	st.DelData("hash_one")
	elem, ok = st.GetData("hash_one")
	assert.False(t, ok)
	assert.Equal(t, st.DataSize(), 1)

	st.Cancel()
	assert.Equal(t, st.DataSize(), 0)

	st.SetData("hash_three", TestData{
		id:   "fe838582",
		name: "Three element description",
	})
	assert.Equal(t, st.DataSize(), 1)

	time.Sleep(4 * time.Second)
	assert.Equal(t, st.DataSize(), 0)
}
