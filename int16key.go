package rbt

// Int16Key is the int16 key for RbKey
type Int16Key int16

// ComparedTo compares the given RbKey with its self
func (ikey *Int16Key) ComparedTo(key RbKey) KeyComparison {
    diff := int16(*ikey - *key.(*Int16Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}