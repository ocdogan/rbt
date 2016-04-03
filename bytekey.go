package rbt

// ByteKey is the float64 key for RbKey
type ByteKey byte

// ComparedTo compares the given RbKey with its self
func (bkey *ByteKey) ComparedTo(key RbKey) KeyComparison {
    switch {
    case *bkey > *key.(*ByteKey):
        return KeyIsGreater
    case *bkey < *key.(*ByteKey):
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}