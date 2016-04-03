package rbt

// KeyComparison structure used as result of comparing two keys 
type KeyComparison int8

const (
    // KeyIsLess is returned as result of key comparison if the first key is less than the second key
    KeyIsLess KeyComparison = iota - 1 
    // KeysAreEqual is returned as result of key comparison if the first key is equal to the second key
    KeysAreEqual
    // KeyIsGreater is returned as result of key comparison if the first key is greater than the second key
    KeyIsGreater
)

const (
    red byte = byte(0)
    black byte = byte(1)
    zeroOrEqual = int8(0)
)

func (tree KeyComparison) String() string {
    switch tree {
    case KeyIsLess:
        return "lessThan"
    case KeyIsGreater:
        return "greaterThan"
    default:
        return "equalTo"
    }
}

// RbKey interface
type RbKey interface {
    ComparedTo(key RbKey) KeyComparison
}

// RbNode structure used for storing key and value pairs
type RbNode struct {
    key RbKey
    value interface{}
    color byte
    left, right, parent *RbNode
}

// RbTree structure
type RbTree struct {
    root *RbNode
    count int
    onInsert KeyValueEvent
    onDelete KeyValueEvent
}

// KeyValueEvent function used on Insert or Delete operations
type KeyValueEvent func(key RbKey, currValue interface{}) (newValue interface{})

// NewRbTree creates a new RbTree and returns its address
func NewRbTree(onInsert, onDelete KeyValueEvent) *RbTree {
    return &RbTree{
        onInsert: onInsert,
        onDelete: onDelete,
    }
}

// newRbNode creates a new RbNode and returns its address
func newRbNode(key RbKey, value interface{}) *RbNode {
    result := &RbNode{
        key: key,
        value: value,
        color: red,
    }
    return result
}

// isRed checks if node exists and its color is red
func isRed(node *RbNode) bool {
    return node != nil && node.color == red
}

// isBlack checks if node exists and its color is black
func isBlack(node *RbNode) bool { 
    return node != nil && node.color == black 
}

// min finds the smallest node key including the given node
func min(node *RbNode) *RbNode {
    if node != nil {
        for node.left != nil {
            node = node.left
        }
    }
    return node
}

// max finds the greatest node key including the given node
func max(node *RbNode) *RbNode {
    if node != nil {
        for node.right != nil {
            node = node.right
        }
    }
    return node
}

// floor returns the largest key node in the subtree rooted at x less than or equal to the given key
func floor(node *RbNode, key RbKey) *RbNode {
    if node == nil {
        return nil
    }
    
    switch key.ComparedTo(node.key) {
    case KeysAreEqual:
        return node
    case KeyIsLess:
        return floor(node.left, key)
    default:
        fn := floor(node.right, key)
        if fn != nil {
            return fn
        }
        return node
    }
}

// ceilig returns the smallest key node in the subtree rooted at x greater than or equal to the given key
func ceiling(node *RbNode, key RbKey) *RbNode {  
    if node == nil {
        return nil
    }
    
    switch key.ComparedTo(node.key) {
    case KeysAreEqual:
        return node
    case KeyIsGreater:
        return ceiling(node.right, key)
    default:
        cn := ceiling(node.left, key)
        if cn != nil {
            return cn
        }
        return node
    }
}

// flipColor switchs the color of the node from red to black or black to red
func flipColor(node *RbNode) {
    if node.color == black {
        node.color = red
    } else {
        node.color = black
    }
}

// colorFlip switchs the color of the node and its children from red to black or black to red
func colorFlip(node *RbNode) {
    flipColor(node)
    flipColor(node.left)
    flipColor(node.right)
}

// rotateLeft makes a right-leaning link lean to the left
func rotateLeft(node *RbNode) *RbNode {
    child := node.right
    node.right = child.left
    child.left = node
    child.color = node.color
    node.color = red

    node.parent = child
    if node.right != nil {
        node.right.parent = node
    }

    return child
}

// rotateRight makes a left-leaning link lean to the right
func rotateRight(node *RbNode) *RbNode {
    child := node.left
    node.left = child.right
    child.right = node
    child.color = node.color
    node.color = red

    node.parent = child
    if node.left != nil {
        node.left.parent = node
    }

    return child
}

// moveRedLeft makes node.left or one of its children red,
// assuming that node is red and both children are black.
func moveRedLeft(node *RbNode) *RbNode {
    colorFlip(node)
    if isRed(node.right.left) {
        node.right = rotateRight(node.right)
        if node.right != nil {
            node.right.parent = node
        }
        node = rotateLeft(node)
        colorFlip(node)
    }
    return node
}

// moveRedRight makes node.right or one of its children red,
// assuming that node is red and both children are black.
func moveRedRight(node *RbNode) *RbNode {
    colorFlip(node)
    if isRed(node.left.left) {
        node = rotateRight(node)
        colorFlip(node)
    }
    return node
}

// balance restores red-black tree invariant
func balance(node *RbNode) *RbNode {
    if isRed(node.right) {
        node = rotateLeft(node)
    }
    if isRed(node.left) && isRed(node.left.left) {
        node = rotateRight(node)
    }
    if isRed(node.left) && isRed(node.right) {
        colorFlip(node)
    }
    return node
}

// deleteMin removes the smallest key and associated value from the tree
func deleteMin(node *RbNode) *RbNode {
    if node.left == nil {
        return nil
    }    
    if isBlack(node.left) && !isRed(node.left.left) {
        node = moveRedLeft(node)
    }
    node.left = deleteMin(node.left)
    if node.left != nil {
        node.left.parent = node
    }
    return balance(node)
}

// Count returns if count of the nodes stored.
func (tree *RbTree) Count() int {
    return tree.count
}

// IsEmpty returns if the tree has any node.
func (tree *RbTree) IsEmpty() bool {
    return tree.root == nil
}

// Min returns the smallest key in the tree.
func (tree *RbTree) Min() (RbKey, interface{}) {
    if tree.root != nil {
        result := min(tree.root)
        return result.key, result.value
    }
    return nil, nil
} 

// Max returns the largest key in the tree.
func (tree *RbTree) Max() (RbKey, interface{}) {
    if tree.root != nil {
        result := max(tree.root)
        return result.key, result.value
    }
    return nil, nil
} 

// Floor returns the largest key in the tree less than or equal to key
func (tree *RbTree) Floor(key RbKey) (RbKey, interface{}) {
    if key != nil && tree.root != nil {
        node := floor(tree.root, key)
        if node == nil {
            return nil, nil
        }
        return node.key, node.value
    }
    return nil, nil
}    

// Ceiling returns the smallest key in the tree greater than or equal to key
func (tree *RbTree) Ceiling(key RbKey) (RbKey, interface{}) {
    if key != nil && tree.root != nil {
        node := ceiling(tree.root, key)
        if node == nil {
            return nil, nil
        }
        return node.key, node.value
    }
    return nil, nil
}

// Get returns the stored value if key found and 'true', 
// otherwise returns 'false' with second return param if key not found 
func (tree *RbTree) Get(key RbKey) (interface{}, bool) {
    if key != nil && tree.root != nil {
        node := tree.find(key)
        if node != nil {
            return node.value, true
        }
    }
    return nil, false
}

// find returns the node if key found, otherwise returns nil 
func (tree *RbTree) find(key RbKey) *RbNode {
    for node := tree.root; node != nil; { 
        switch key.ComparedTo(node.key) {
        case KeyIsLess:
            node = node.left
        case KeyIsGreater:
            node = node.right
        default:
            return node
        }    
    }
    return nil
}

// Insert inserts the given key and value into the tree
func (tree *RbTree) Insert(key RbKey, value interface{}) {
    if key != nil {
        tree.root = tree.insertNode(tree.root, key, value);
        tree.root.color = black
        tree.root.parent = nil
    }
}

// insertNode adds the given key and value into the node
func (tree *RbTree) insertNode(node *RbNode, key RbKey, value interface{}) *RbNode {
    if node == nil {
        tree.count++
        return newRbNode(key, value)
    }

    switch key.ComparedTo(node.key) {
    case KeyIsLess:
        node.left  = tree.insertNode(node.left,  key, value)
        node.left.parent = node
    case KeyIsGreater:
        node.right = tree.insertNode(node.right, key, value)
        node.right.parent = node
    default:
        if tree.onInsert == nil {
            node.value = value
        } else {
            node.value = tree.onInsert(key, value)
        }
    }
    return balance(node)
}

// Delete deletes the given key from the tree
func (tree *RbTree) Delete(key RbKey) {
    tree.root = tree.deleteNode(tree.root, key)
    if tree.root != nil {
        tree.root.color = black
        tree.root.parent = nil
    }
}

// deleteNode deletes the given key from the node
func (tree *RbTree) deleteNode(node *RbNode, key RbKey) *RbNode {
    if node == nil {
        return nil
    }
    
    cmp := key.ComparedTo(node.key)
    if cmp == KeyIsLess {
        if isBlack(node.left) && !isRed(node.left.left) {
            node = moveRedLeft(node)
        }
        node.left = tree.deleteNode(node.left, key)
        if node.left != nil {
            node.left.parent = node
        }
    } else {
        if cmp == KeysAreEqual && tree.onDelete != nil {
            value := tree.onInsert(key, node.value)        
            if value != nil {
                node.value = value
                return node
            }
        }
        
        if isRed(node.left) {
            node = rotateRight(node)
        }
        
        if isBlack(node.right) && !isRed(node.right.left) {
            node = moveRedRight(node)
        }
        
        if key.ComparedTo(node.key) != KeysAreEqual {
            node.right = tree.deleteNode(node.right, key)
            if node.right != nil {
                node.right.parent = node
            }
        } else {
            if node.right == nil {
                return nil
            }

            rm := min(node.right)
            node.key   = rm.key
            node.value = rm.value
            node.right = deleteMin(node.right)

            if node.right != nil {
                node.right.parent = node
            }
            
            rm.left = nil
            rm.right = nil
            rm.parent = nil
            
            tree.count--
        }
    }
    return balance(node)
}