package fifoSet

import (
	"carted/pkg/queue"
	"container/list"
	"io"
	"sync"
)

type fifoSet[TElem comparable] struct {
	mutex         sync.Mutex
	elemContainer *list.List
	elemMap       map[TElem]*list.Element
}

func New[TElem comparable]() *fifoSet[TElem] {
	return &fifoSet[TElem]{
		elemContainer: list.New(),
		elemMap:       make(map[TElem]*list.Element),
	}
}

// queue.Queue implementation
var _ queue.Queue[int] = (*fifoSet[int])(nil)

func (set *fifoSet[T]) Pop() (T, error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	if set.elemContainer.Len() == 0 {
		return *new(T), io.EOF
	}

	poppedElem := set.elemContainer.Remove(set.elemContainer.Front()).(T)
	delete(set.elemMap, poppedElem)
	return poppedElem, nil
}

func (set *fifoSet[T]) Push(newElem T) {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	if existingElem, ok := set.elemMap[newElem]; ok {
		set.elemContainer.MoveToBack(existingElem)
	} else {
		containerEntry := set.elemContainer.PushBack(newElem)
		set.elemMap[newElem] = containerEntry
	}
}
