package rbt

// BoolKey is the float64 key for RbKey
type BoolKey bool

// ComparedTo compares the given RbKey with its self
func (bkey *BoolKey) ComparedTo(key RbKey) KeyComparison {
    var key1 = bool(*bkey)
    var key2 = bool(*key.(*BoolKey))
    
    switch {
    case key1 && !key2:
        return KeyIsGreater
    case !key1 && key2:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}