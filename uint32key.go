package rbt

// Uint32Key is the uint32 key for RbKey
type Uint32Key uint32

// ComparedTo compares the given RbKey with its self
func (ikey *Uint32Key) ComparedTo(key RbKey) KeyComparison {
    diff := int64(*ikey) - int64(*key.(*Uint32Key))
    switch {
    case diff > zeroInt64:
        return KeyIsGreater
    case diff < zeroInt64:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}