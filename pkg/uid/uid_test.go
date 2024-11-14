package uid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	id1 := Get(100)
	id2 := Get(100)

	assert.NotEqual(t, id1, id2)
}

func ExampleGet() {
	id1 := Get(0)
	id2 := Get(0)
	fmt.Printf("%v %v\n", id1, id2)
}