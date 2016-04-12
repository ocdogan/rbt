package rbt

import (
    "encoding/binary"
)

const (
    c1 uint32 = 0xcc9e2d51
    c2 uint32 = 0x1b873593
)

func MurmurHash3(data []byte, length uint32, seed uint32) uint32 {
    h1 := seed
    nblocks := length >> 2

    i := 0
    for j := nblocks; j > 0; j-- {
        k1l := binary.LittleEndian.Uint32(data[i:])

        k1l *= c1
        k1l = mmhRotateLeft(k1l, 15)
        k1l *= c2

        h1 ^= k1l
        h1 = mmhRotateLeft(h1, 13)
        h1 = h1 * 5 + 0xe6546b64

        i += 4
    }

    nblocks <<= 2
    k1 := uint32(0)
    tailLength := length & 3

    if tailLength == 3 {
        k1 ^= uint32(data[2 + nblocks]) << 16
    }
    if tailLength >= 2 {
        k1 ^= uint32(data[1 + nblocks]) << 8
    }
    if tailLength >= 1 {
        k1 ^= uint32(data[nblocks])
        k1 *= c1 
        k1 = mmhRotateLeft(k1, 15) 
        k1 *= c2
        h1 ^= k1
    }

    h1 ^= length
    return mmhFMix(h1)
}

func mmhFMix(h uint32) uint32 {
    h ^= h >> 16
    h *= 0x85ebca6b
    h ^= h >> 13
    h *= 0xc2b2ae35
    h ^= h >> 16

    return h
}

func mmhRotateLeft(x uint32, r byte) uint32 {
    return uint32((x << r) | (x >> (32 - r)))
}
