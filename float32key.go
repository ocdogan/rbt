package rbt

// Float32Key is the float64 key for RbKey
type Float32Key float32

// ComparedTo compares the given RbKey with its self
func (fkey *Float32Key) ComparedTo(key RbKey) KeyComparison {
    diff := float32(*fkey - *key.(*Float32Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}