package rbt

// Float64Key is the float64 key for RbKey
type Float64Key float64

const zeroFloat64Key = Float64Key(0)

// ComparedTo compares the given RbKey with its self
func (fkey *Float64Key) ComparedTo(key RbKey) KeyComparison {
    diff := *fkey - *key.(*Float64Key)
    switch {
    case diff > zeroFloat64Key:
        return KeyIsGreater
    case diff < zeroFloat64Key:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}