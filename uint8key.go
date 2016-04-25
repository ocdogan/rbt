package rbt

// Uint8Key is the uint8 key for RbKey
type Uint8Key uint8

const zeroInt16 = int16(0)

// ComparedTo compares the given RbKey with its self
func (ikey *Uint8Key) ComparedTo(key RbKey) KeyComparison {
    diff := int16(*ikey) - int16(*key.(*Uint8Key))
    switch {
    case diff > zeroInt16:
        return KeyIsGreater
    case diff < zeroInt16:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}