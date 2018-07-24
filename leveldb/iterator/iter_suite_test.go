package iterator_test

import (
	"testing"

	"github.com/red010b37/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
