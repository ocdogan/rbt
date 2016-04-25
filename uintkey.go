package rbt

// UintKey is the uint key for RbKey
type UintKey uint

const zeroInt64 = int64(0)

// ComparedTo compares the given RbKey with its self
func (ikey *UintKey) ComparedTo(key RbKey) KeyComparison {
    diff := int64(*ikey) - int64(*key.(*UintKey))
    switch {
    case diff > zeroInt64:
        return KeyIsGreater
    case diff < zeroInt64:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}