package rbt

// NilKey is the nil value key for RbKey
type NilKey struct {}

// ComparedTo compares the given RbKey with its self
func (nkey *NilKey) ComparedTo(key RbKey) KeyComparison {
    if key == nil {
        return KeysAreEqual
    } 
    if _, ok := key.(*NilKey); ok {
        return KeysAreEqual
    }
    return KeyIsLess
}