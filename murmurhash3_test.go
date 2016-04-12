package rbt

import (
    "encoding/binary"
    "fmt"
    "runtime"
    "time"
	"testing"
)

func TestMurmurHash3(t *testing.T) {
    fmt.Println("\nTestMurmurHash3\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

    mem1 := new(runtime.MemStats)
    runtime.ReadMemStats(mem1)
    
    key := make([]byte, 256)
    hashes := make([]byte, 1024)

    ellapsed := int64(0)
    for i := 0; i < 256; i++ {
        t1 := time.Now()
        key[i] = byte(i)
        result := MurmurHash3(key, uint32(i), uint32(256 - i))
        ellapsed += time.Now().Sub(t1).Nanoseconds()

        binary.LittleEndian.PutUint32(hashes[4*i:4*i + 4], result)
    }
    fmt.Printf("MurmurHash3 avg time: %.5f nanosec\n", float64(ellapsed)/float64(256))
    
    mem2 := new(runtime.MemStats)
    runtime.ReadMemStats(mem2)
    if mem2.Alloc <= mem1.Alloc {
        fmt.Printf("Mem allocated: 0 MB\n")
    } else {
        fmt.Printf("Mem allocated: %3.3f MB\n", float64(mem2.Alloc - mem1.Alloc)/(1024*1024))
    }

    finalr := MurmurHash3(hashes, 1024, 0)
    var verification uint32 = 0xB0F57EE3

    fmt.Printf("MurmurHash3 finalr: %d, verification: %d\n", finalr, verification)
    if verification != finalr {
        t.Fail()
    } else {
        fmt.Println("passed")
    }
}