package rbt

// IntKey is the integer key for RbKey
type IntKey int

// ComparedTo compares the given RbKey with its self
func (ikey *IntKey) ComparedTo(key RbKey) KeyComparison {
    diff := int(*ikey - *key.(*IntKey))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}