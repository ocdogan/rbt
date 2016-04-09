package rbt

import (
    "fmt"
    "runtime"
	"testing"
    "time"
)

func TestInsertDeleteAndGet(t *testing.T) {
    fmt.Println("\nTestInsertDeleteAndGet\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

    mem1 := new(runtime.MemStats)
    runtime.ReadMemStats(mem1)
    
    t1 := time.Now()

    tree := NewRbTree(nil, nil)
    for i := 0; i < 1000000; i++ {
        key := IntKey(i)
        tree.Insert(&key, 10 + i)
    }

    t2 := time.Now()
    fmt.Printf("Insert time: %.5f sec\n", float64(t2.Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()))
    
    count := 0
    for i := 0; i < 1500000; i++ {
        key := IntKey(i)
        if _, ok := tree.Get(&key); ok {
            count++
        }
    }

    t3 := time.Now()
    fmt.Printf("Search time: %.5f sec with count %d\n", float64(t3.Sub(t2).Nanoseconds())/float64(time.Second.Nanoseconds()), count)

    for i := 1; i < 1000000; i++ {
        key := IntKey(i)
        tree.Delete(&key)
    }
    
    t4 := time.Now()
    fmt.Printf("Delete time: %.5f sec\n", float64(t4.Sub(t3).Nanoseconds())/float64(time.Second.Nanoseconds()))
    
    mem2 := new(runtime.MemStats)
    runtime.ReadMemStats(mem2)
    if mem2.Alloc <= mem1.Alloc {
        fmt.Printf("Mem allocated: 0 MB\n")
    } else {
        fmt.Printf("Mem allocated: %3.3f MB\n", float64(mem2.Alloc - mem1.Alloc)/(1024*1024))
    }
}

type mapIntKey int

func (ikey mapIntKey) ComparedTo(key RbKey) KeyComparison {
    diff := int(ikey - key.(mapIntKey))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}

func TestInsertDeleteAndGetMap(t *testing.T) {
    fmt.Println("\nTestInsertDeleteAndGetMap\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

    mem1 := new(runtime.MemStats)
    runtime.ReadMemStats(mem1)
    
    t1 := time.Now()

    tree := make(map[RbKey]interface{})
    for i := 0; i < 1000000; i++ {
        key := mapIntKey(i)
        tree[key] = 10 + i
    }

    t2 := time.Now()
    fmt.Printf("Insert map time: %.5f sec\n", float64(t2.Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()))
    
    count := 0
    for i := 0; i < 1500000; i++ {
        key := mapIntKey(i)
        _, ok := tree[key]
        if ok {
            count++
        }
    }

    t3 := time.Now()
    fmt.Printf("Search map time: %.5f sec with count %d\n", float64(t3.Sub(t2).Nanoseconds())/float64(time.Second.Nanoseconds()), count)

    for i := 1; i < 1000000; i++ {
        key := mapIntKey(i)
        delete(tree, key)
    }
    
    t4 := time.Now()
    fmt.Printf("Delete map time: %.5f sec\n", float64(t4.Sub(t3).Nanoseconds())/float64(time.Second.Nanoseconds()))
    
    mem2 := new(runtime.MemStats)
    runtime.ReadMemStats(mem2)
    if mem2.Alloc <= mem1.Alloc {
        fmt.Printf("Mem map allocated: 0 MB\n")
    } else {
        fmt.Printf("Mem map allocated: %3.3f MB\n", float64(mem2.Alloc - mem1.Alloc)/(1024*1024))
    }
}