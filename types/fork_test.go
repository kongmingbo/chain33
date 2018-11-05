package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForks(t *testing.T) {
	setLocalFork()
	assert.Equal(t, systemFork.IsFork("abc", 1, "ForkV1"), false)
	assert.Equal(t, systemFork.IsFork("abc", 1, "ForkV12"), false)
	assert.Equal(t, systemFork.IsFork("bityuan", 1, "ForkTransferExec"), false)
	assert.Equal(t, systemFork.IsFork("local", 0, "ForkBlockHash"), false)
	assert.Equal(t, systemFork.IsFork("local", 1, "ForkBlockHash"), true)
	assert.Equal(t, systemFork.IsFork("local", 1, "ForkTransferExec"), true)
}