package rbt

// Uint8Key is the uint8 key for RbKey
type Uint8Key uint8

// ComparedTo compares the given RbKey with its self
func (ikey *Uint8Key) ComparedTo(key RbKey) KeyComparison {
    diff := int16(*ikey) - int16(*key.(*Uint8Key))
    switch {
    case diff > 0:
        return KeyIsGreater
    case diff < 0:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}