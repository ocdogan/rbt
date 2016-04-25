package rbt

// Int64Key is the int64 key for RbKey
type Int64Key int64

const zeroInt64Key = Int64Key(0)

// ComparedTo compares the given RbKey with its self
func (ikey *Int64Key) ComparedTo(key RbKey) KeyComparison {
    diff := *ikey - *key.(*Int64Key)
    switch {
    case diff > zeroInt64Key:
        return KeyIsGreater
    case diff < zeroInt64Key:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}