package queue

type Queue[TElem comparable] interface {
	Pop() (TElem, error)
	Push(newElem TElem)
}
