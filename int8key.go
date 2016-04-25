package rbt

// Int8Key is the int8 key for RbKey
type Int8Key int8

const zeroInt8Key = Int8Key(0)

// ComparedTo compares the given RbKey with its self
func (ikey *Int8Key) ComparedTo(key RbKey) KeyComparison {
    diff := *ikey - *key.(*Int8Key)
    switch {
    case diff > zeroInt8Key:
        return KeyIsGreater
    case diff < zeroInt8Key:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}