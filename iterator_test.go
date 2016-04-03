package rbt

import (
    "fmt"
    "runtime"
	"testing"
    "time"
)

func TestIterate(t *testing.T) {
    mem1 := new(runtime.MemStats)
    runtime.ReadMemStats(mem1)
    
    t1 := time.Now()

    tree := NewRbTree(nil, nil)
    for i := 1; i <= 1000000; i++ {
        key := IntKey(i)
        tree.Insert(&key, 10 + i)
    }

    fmt.Printf("Insert time: %.5f sec\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()))
    
    count := 0
    iterator, err := tree.NewRbIterator(func(iterator RbIterator, key RbKey, value interface{}){
        count++
    })
    
    if err != nil {
        return
    }
    
    t1 = time.Now()

    count = 0
    iterator.All()
    fmt.Printf("All completed in: %.5f sec with count %d\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()), count)
    
    t1 = time.Now()

    count = 0
    loKey, hiKey := IntKey(0), IntKey(2000000)
    iterator.Between(&loKey, &hiKey)
    fmt.Printf("Between completed in: %.5f sec with count %d\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()), count)
    
    t1 = time.Now()

    count = 0
    key := IntKey(900001)
    iterator.LessThan(&key)
    fmt.Printf("LessThan completed in: %.5f sec with count %d\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()), count)
    
    t1 = time.Now()

    count = 0
    key = IntKey(100000)
    iterator.GreaterThan(&key)
    fmt.Printf("GreaterThan completed in: %.5f sec with count %d\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()), count)
    
    t1 = time.Now()

    count = 0
    key = IntKey(1000000)
    iterator.LessOrEqual(&key)
    fmt.Printf("LessOrEqual completed in: %.5f sec with count %d\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()), count)
    
    t1 = time.Now()

    count = 0
    key = IntKey(0)
    iterator.GreaterOrEqual(&key)
    fmt.Printf("GreaterOrEqual completed in: %.5f sec with count %d\n", float64(time.Now().Sub(t1).Nanoseconds())/float64(time.Second.Nanoseconds()), count)
    
    mem2 := new(runtime.MemStats)
    runtime.ReadMemStats(mem2)
    fmt.Printf("Mem allocated: %9.3f MB\n", float64(mem2.Alloc - mem1.Alloc)/(1024*1024))
}