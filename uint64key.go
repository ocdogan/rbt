package rbt

// Uint64Key is the uint64 key for RbKey
type Uint64Key uint64

// ComparedTo compares the given RbKey with its self
func (ikey *Uint64Key) ComparedTo(key RbKey) KeyComparison {
    key1 := uint64(*ikey)
    key2 := uint64(*key.(*Uint64Key))
    switch {
    case key1 > key2:
        return KeyIsGreater
    case key1 < key2:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}