package fifoSet_test

import (
	"carted/pkg/queue/fifoSet"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonEmptyPop(t *testing.T) {
	testcases := []struct {
		toPush         []int
		expectedPopped int
	}{
		{toPush: []int{1, 2, 3}, expectedPopped: 1},
		{toPush: []int{1, 2, 3, 1}, expectedPopped: 2},
		{toPush: []int{1}, expectedPopped: 1},
	}

	for _, tc := range testcases {
		queue := fifoSet.New[int]()
		for _, el := range tc.toPush {
			queue.Push(el)
		}

		popped, err := queue.Pop()
		assert.Nil(t, err)
		assert.Equal(t, popped, tc.expectedPopped)
	}
}

func TestEmptyPop(t *testing.T) {
	queue := fifoSet.New[int]()
	_, err := queue.Pop()
	assert.ErrorIs(t, err, io.EOF)
}

func TestLength(t *testing.T) {
	testcases := []struct {
		toPush         []string
		expectedLength int
	}{
		{toPush: []string{"a", "ab", "bcd"}, expectedLength: 3},
		{toPush: []string{"aa", "aa", "aa"}, expectedLength: 1},
		{toPush: []string{"a", "ab", "bcd", "ab"}, expectedLength: 3},
	}

	for _, tc := range testcases {
		queue := fifoSet.New[string]()
		for _, el := range tc.toPush {
			queue.Push(el)
		}

		for i := 0; i < tc.expectedLength; i++ {
			_, err := queue.Pop()
			assert.Nil(t, err)
		}

		_, err := queue.Pop()
		assert.ErrorIs(t, err, io.EOF)
	}
}

func TestConcurrentConsumers(t *testing.T) {
	queue := fifoSet.New[int]()

	var wg sync.WaitGroup

	pushOnes := func(iterationCount int) {
		for i := 0; i < iterationCount; i++ {
			queue.Push(1)
			queue.Pop()
		}
		wg.Done()
	}

	wg.Add(3)
	go pushOnes(1000)
	go pushOnes(1000)
	go pushOnes(1000)

	wg.Wait()

	_, err := queue.Pop()
	assert.ErrorIs(t, err, io.EOF)
}

func BenchmarkQueue(b *testing.B) {
	queue := fifoSet.New[int]()
	for i := 0; i < b.N; i++ {
		queue.Push(1)
		queue.Pop()
	}
}
