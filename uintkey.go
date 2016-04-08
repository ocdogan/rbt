package rbt

// UintKey is the uint key for RbKey
type UintKey uint

// ComparedTo compares the given RbKey with its self
func (ikey *UintKey) ComparedTo(key RbKey) KeyComparison {
    diff := int64(*ikey) - int64(*key.(*UintKey))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}