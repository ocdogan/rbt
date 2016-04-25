package rbt

// Uint16Key is the uint16 key for RbKey
type Uint16Key uint16

const zeroInt32 = int32(0)

// ComparedTo compares the given RbKey with its self
func (ikey *Uint16Key) ComparedTo(key RbKey) KeyComparison {
    diff := int32(*ikey) - int32(*key.(*Uint16Key))
    switch {
    case diff > zeroInt32:
        return KeyIsGreater
    case diff < zeroInt32:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}