package rbt

// IntKey is the integer key for RbKey
type IntKey int

const zeroIntKey = IntKey(0)

// ComparedTo compares the given RbKey with its self
func (ikey *IntKey) ComparedTo(key RbKey) KeyComparison {
    diff := *ikey - *key.(*IntKey)
    switch {
    case diff > zeroIntKey:
        return KeyIsGreater
    case diff < zeroIntKey:
        return KeyIsLess
    default:
        return KeysAreEqual
    }
}