package ds

type List interface {
	Size() int
	IsEmpty() bool
	Contains(elem interface{}) bool
	ContainsAll(elemList ...interface{}) bool
	ToArray() []interface{}
	Add(elem interface{}) bool
	AddAll(elemList ...interface{}) bool
	Remove(elem interface{}) bool
	RemoveAll(elem ...interface{}) bool
	RemoveAt(idx int) (interface{}, bool)
	SubList(from, to int) (List, error)
	Get(idx int) (interface{}, error)
	Set(idx int, elem interface{}) error
	IndexOf(elem interface{}) int
	LastIndexOf(elem interface{}) int
	Clear()
}

/* Comparator compares priorities of 2 elements.
 * if result > 0, priority of 'a' is higher,
 * if result == 0, priorities of 'a' && 'b' are the same,
 * if result < 0, priority of 'a' is lower.
 */
type Comparator func(a, b interface{}) int

/* TypeChecker checks if the type of the current elem is valid */
type TypeChecker func(elem interface{}) bool

func IsLinkedList(l List) bool {
	_, ok := l.(*LinkedList)
	return ok
}