package rbt

// Uint16Key is the uint16 key for RbKey
type Uint16Key uint16

// ComparedTo compares the given RbKey with its self
func (ikey *Uint16Key) ComparedTo(key RbKey) KeyComparison {
    diff := int32(*ikey) - int32(*key.(*Uint16Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}