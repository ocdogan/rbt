package rbt

// Int64Key is the int64 key for RbKey
type Int64Key int64

// ComparedTo compares the given RbKey with its self
func (ikey *Int64Key) ComparedTo(key RbKey) KeyComparison {
    diff := int64(*ikey - *key.(*Int64Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}