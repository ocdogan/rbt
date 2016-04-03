[RbTree](<https://godoc.org/github.com/ocdogan/rbt>)
====================================================

An iterable basic Red Black tree implementation in Golang.

Installation
------------

`go get `[github.com/ocdogan/rbt](<https://github.com/ocdogan/rbt>)

Doc
---

### type RbIterator

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
type RbIterator interface {
    // All iterates on all items of the RbTree
    All() (int, error)
    // Between iterates on the items of the RbTree that the key of the item
    // is less or equal to loKey and greater or equal to hiKey
    Between(loKey RbKey, hiKey RbKey) (int, error)
    // ClearData clears all the data stored on the iterator
    ClearData()
    // Close closes the current iteration, so the iteration stops iterating
    Close()
    // Closed gives the state of the iterator, 'true' if closed
    Closed() bool
    // CurrentCount gives the count of the items that match the iteration case
    CurrentCount() int
    // LessOrEqual iterates on the items of the RbTree that the key of the item
    // is less or equal to the given key
    LessOrEqual(key RbKey) (int, error)
    // LessThan iterates on the items of the RbTree that the key of the item
    // is less than the given key
    LessThan(key RbKey) (int, error)
    // GetData returns the data stored on the iterator with the dataKey
    GetData(dataKey string) (interface{}, bool)
    // GreaterOrEqual iterates on the items of the RbTree that the key of the item
    // is greater or equal to the given key
    GreaterOrEqual(key RbKey) (int, error)
    // GreaterThan iterates on the items of the RbTree that the key of the item
    // is greater than the given key
    GreaterThan(key RbKey) (int, error)
    // RemoveData deletes the data stored on the iterator with the dataKey
    RemoveData(dataKey string)
    // SetData stores the data with the dataKey on the iterator
    SetData(dataKey string, value interface{})
    // Tree returns the RbTree that the iterator is iterating on
    Tree() *RbTree
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

RbIterator interface used for iterating on a RbTree

### type RbTree

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
type RbTree struct {
    // contains filtered or unexported fields
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

RbTree structure

#### func NewRbTree

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func NewRbTree(onInsert, onDelete KeyValueEvent) *RbTree
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

NewRbTree creates a new RbTree and returns its address

#### func (\*RbTree) Ceiling

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Ceiling(key RbKey) (RbKey, interface{})
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Ceiling returns the smallest key in the tree greater than or equal to key

#### func (\*RbTree) Count

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Count() int
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Count returns if count of the nodes stored.

#### func (\*RbTree) Delete

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Delete(key RbKey)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Delete deletes the given key from the tree

#### func (\*RbTree) Floor

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Floor(key RbKey) (RbKey, interface{})
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Floor returns the largest key in the tree less than or equal to key

#### func (\*RbTree) Get

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Get(key RbKey) (interface{}, bool)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Get returns the stored value if key found and 'true', otherwise returns 'false'
with second return param if key not found

#### func (\*RbTree) Insert

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Insert(key RbKey, value interface{})
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Insert inserts the given key and value into the tree

#### func (\*RbTree) IsEmpty

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) IsEmpty() bool
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

IsEmpty returns if the tree has any node.

#### func (\*RbTree) Max

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Max() (RbKey, interface{})
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Max returns the largest key in the tree.

#### func (\*RbTree) Min

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) Min() (RbKey, interface{})
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Min returns the smallest key in the tree.

#### func (\*RbTree) NewRbIterator

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (tree *RbTree) NewRbIterator(callback RbIterationCallback) (RbIterator, error)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

NewRbIterator creates a new iterator for the given RbTree

 

Examples
--------

Add, delete, get operations.

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
package main

import (
    "fmt"
    "runtime"
    "testing"
    "time"

    "github.com/ocdogan/rbt"
)

func main() {
    mem1 := new(runtime.MemStats)
    runtime.ReadMemStats(mem1)

    tree := NewRbTree(nil, nil)
    t1 := time.Now()

    for i := 0; i < 1000000; i++ {
        key := IntKey(i)
        tree.Insert(&key, 10 + i)
    }

    t2 := time.Now()
    fmt.Printf("Insert time: %.5f sec\n", 
        float64(t2.Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()))

    for i := 0; i < 1500000; i++ {
        key := IntKey(i)
        tree.Get(&key)
    }

    t3 := time.Now()
    fmt.Printf("Search time: %.5f sec\n",
        float64(t3.Sub(t2).Nanoseconds())/float64(time.Second.Nanoseconds()))

    for i := 1; i < 1000000; i++ {
        key := IntKey(i)
        tree.Delete(&key)
    }

    t4 := time.Now()
    fmt.Printf("Delete time: %.5f sec\n",
        float64(t4.Sub(t3).Nanoseconds())/float64(time.Second.Nanoseconds()))

    mem2 := new(runtime.MemStats)
    runtime.ReadMemStats(mem2)
    fmt.Printf("Mem allocated: %9.3f MB\n",
        float64(mem2.Alloc - mem1.Alloc)/(1024*1024))
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

 

Iterating on the tree.

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
package main

import (
    "fmt"
    "runtime"
    "testing"
    "time"

    "github.com/ocdogan/rbt"
)

func main() {
    mem1 := new(runtime.MemStats)
    runtime.ReadMemStats(mem1)

    tree := NewRbTree(nil, nil)
    t1 := time.Now()

    // Insert    
    for i := 1; i <= 1000000; i++ {
        key := IntKey(i)
        tree.Insert(&key, 10 + i)
    }

    fmt.Printf("Insert time: %.5f sec\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()))

    // Initialize iterator    
    count := 0
    iterator, err := tree.NewRbIterator(func(iterator RbIterator, 
        key RbKey, value interface{}){
        count++
    })

    if err != nil {
        return
    }

    // All    
    t1 = time.Now()
    count = 0
    iterator.All()
    fmt.Printf("All completed in: %.5f sec with count %d\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()), count)

    // Between    
    t1 = time.Now()
    count = 0
    loKey, hiKey := IntKey(0), IntKey(2000000)
    iterator.Between(&loKey, &hiKey)
    fmt.Printf("Between completed in: %.5f sec with count %d\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()), count)

    // LessThan    
    t1 = time.Now()
    count = 0
    key := IntKey(900001)
    iterator.LessThan(&key)
    fmt.Printf("LessThan completed in: %.5f sec with count %d\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()), count)

    // GreaterThan    
    t1 = time.Now()
    count = 0
    key = IntKey(100000)
    iterator.GreaterThan(&key)
    fmt.Printf("GreaterThan completed in: %.5f sec with count %d\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()), count)

    // LessOrEqual    
    t1 = time.Now()
    count = 0
    key = IntKey(1000000)
    iterator.LessOrEqual(&key)
    fmt.Printf("LessOrEqual completed in: %.5f sec with count %d\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()), count)

    // GreaterOrEqual    
    t1 = time.Now()
    count = 0
    key = IntKey(0)
    iterator.GreaterOrEqual(&key)
    fmt.Printf("GreaterOrEqual completed in: %.5f sec with count %d\n",
        float64(time.Now().Sub(t1).Nanoseconds())
            /float64(time.Second.Nanoseconds()), count)

    mem2 := new(runtime.MemStats)
    runtime.ReadMemStats(mem2)
    fmt.Printf("Mem allocated: %9.3f MB\n", 
        float64(mem2.Alloc -    mem1.Alloc)/(1024*1024))
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
