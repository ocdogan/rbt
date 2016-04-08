package rbt

// StringKey is the string key for RbKey
type StringKey string

// ComparedTo compares the given RbKey with its self
func (skey *StringKey) ComparedTo(key RbKey) KeyComparison {
    switch {
    case *skey > *key.(*StringKey):
        return KeyIsGreater
    case *skey < *key.(*StringKey):
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}