package rbt

// Int32Key is the float64 key for RbKey
type Int32Key int32

// ComparedTo compares the given RbKey with its self
func (ikey *Int32Key) ComparedTo(key RbKey) KeyComparison {
    diff := int32(*ikey - *key.(*Int32Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}