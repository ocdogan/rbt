package rbt

// Int8Key is the float64 key for RbKey
type Int8Key int8

// ComparedTo compares the given RbKey with its self
func (ikey *Int8Key) ComparedTo(key RbKey) KeyComparison {
    diff := int8(*ikey - *key.(*Int8Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}