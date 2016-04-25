package rbt

// Float32Key is the float32 key for RbKey
type Float32Key float32

const zeroFloat32Key = Float32Key(0)

// ComparedTo compares the given RbKey with its self
func (fkey *Float32Key) ComparedTo(key RbKey) KeyComparison {
    diff := *fkey - *key.(*Float32Key)
    switch {
    case diff > zeroFloat32Key:
        return KeyIsGreater
    case diff < zeroFloat32Key:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}