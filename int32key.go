package rbt

// Int32Key is the int32 key for RbKey
type Int32Key int32

const zeroInt32Key = Int32Key(0)

// ComparedTo compares the given RbKey with its self
func (ikey *Int32Key) ComparedTo(key RbKey) KeyComparison {
    diff := *ikey - *key.(*Int32Key)
    switch {
    case diff > zeroInt32Key:
        return KeyIsGreater
    case diff < zeroInt32Key:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}