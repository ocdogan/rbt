package rbt

// Float64Key is the float64 key for RbKey
type Float64Key float64

// ComparedTo compares the given RbKey with its self
func (fkey *Float64Key) ComparedTo(key RbKey) KeyComparison {
    diff := float64(*fkey - *key.(*Float64Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}