package rbt

import (
    "sync"
    "sync/atomic"
)

// RbIterator interface used for iterating on a RbTree
type RbIterator interface {
    // All iterates on all items of the RbTree
    All() (int, error)
    // Between iterates on the items of the RbTree that the key of the item 
    // is less or equal to loKey and greater or equal to hiKey
    Between(loKey RbKey, hiKey RbKey) (int, error)
    // ClearData clears all the data stored on the iterator
    ClearData()
    // Close closes the current iteration, so the iteration stops iterating
    Close()
    // Closed gives the state of the iterator, 'true' if closed
    Closed() bool
    // CurrentCount gives the count of the items that match the iteration case
    CurrentCount() int
    // LessOrEqual iterates on the items of the RbTree that the key of the item 
    // is less or equal to the given key
    LessOrEqual(key RbKey) (int, error)
    // LessThan iterates on the items of the RbTree that the key of the item 
    // is less than the given key
    LessThan(key RbKey) (int, error)
    // GetData returns the data stored on the iterator with the dataKey 
    GetData(dataKey string) (interface{}, bool)
    // GreaterOrEqual iterates on the items of the RbTree that the key of the item 
    // is greater or equal to the given key
    GreaterOrEqual(key RbKey) (int, error)
    // GreaterThan iterates on the items of the RbTree that the key of the item 
    // is greater than the given key
    GreaterThan(key RbKey) (int, error)
    // RemoveData deletes the data stored on the iterator with the dataKey 
    RemoveData(dataKey string)
    // SetData stores the data with the dataKey on the iterator 
    SetData(dataKey string, value interface{})
    // Tree returns the RbTree that the iterator is iterating on
    Tree() *RbTree
}

type rbIterationContext struct {
    tree *RbTree
    count int32
    state int32
    mtx sync.Mutex
    callback RbIterationCallback
    data map[string]interface{}
}

const (
    iteratorReady = int32(1)
    iterWalking = int32(2)
    iteratorClosed = int32(-1)
    iteratorUninitialized = int32(0)
)

// RbIterationCallback is the function used to by the RbIterator 
// with will be called on iteration match
type RbIterationCallback func(iterator RbIterator, key RbKey, value interface{})

func nilIterationCallback(iterator RbIterator, key RbKey, value interface{}) {
    return
}

// NewRbIterator creates a new iterator for the given RbTree
func (tree *RbTree) NewRbIterator(callback RbIterationCallback) (RbIterator, error) {
    if tree == nil {
        return nil, ArgumentNilError("tree")
    }
    if callback == nil {
        return nil, ArgumentNilError("callback")
    }
    
    return &rbIterationContext{
        tree: tree,
        callback: callback,
        mtx: sync.Mutex{},
        state: iteratorReady,
        data: make(map[string]interface{}),
    }, nil
}

func (context *rbIterationContext) Tree() *RbTree {
    return context.tree
}

func (context *rbIterationContext) CurrentCount() int {
    return int(atomic.LoadInt32(&context.count))
}

func (context *rbIterationContext) incrementCount() {
    atomic.AddInt32(&context.count, 1)
}

func (context *rbIterationContext) inWalk() bool {
    return atomic.LoadInt32(&context.state) == iterWalking
}

func (context *rbIterationContext) ready() bool {
    return atomic.LoadInt32(&context.state) != iteratorReady
}

func (context *rbIterationContext) Closed() bool {
    return atomic.LoadInt32(&context.state) != iteratorClosed
}

func (context *rbIterationContext) Close() {
    context.mtx.Lock()
    defer context.mtx.Unlock()

    context.state = iteratorClosed
    context.callback = nilIterationCallback
    context.tree = nil
}

func (context *rbIterationContext) ClearData() {
    context.mtx.Lock()
    context.data = nil
    context.mtx.Unlock()
}

func (context *rbIterationContext) GetData(dataKey string) (interface{}, bool) {
    context.mtx.Lock()
    data := context.data
    context.mtx.Unlock()
    
    if data != nil {
        result, ok := data[dataKey]
        return result, ok
    }
    return nil, false
}

func (context *rbIterationContext) SetData(dataKey string, value interface{}) {
    context.mtx.Lock()
    data := context.data
    context.mtx.Unlock()
    
    if data != nil {
        data[dataKey] = value
    }
}

func (context *rbIterationContext) RemoveData(dataKey string) {
    context.mtx.Lock()
    data := context.data
    context.mtx.Unlock()
    
    if data != nil {
        delete(data, dataKey)
    }
}

func (context *rbIterationContext) checkStateAndGetTree() (*RbTree, error) {
    context.mtx.Lock()
    defer context.mtx.Unlock()
    
    switch context.state {
    case iterWalking:
        return nil, ErrIteratorAlreadyRunning
    case iteratorClosed:
        return nil, ErrIteratorClosed
    case iteratorUninitialized:
        return nil, ErrIteratorUninitialized
    case iteratorReady:
        context.count = int32(0)
        context.state = iterWalking
    }
    if context.tree == nil {
        return nil, ErrIteratorClosed
    }
    return context.tree, nil 
}

func (context *rbIterationContext) All() (int, error) {
    tree, err := context.checkStateAndGetTree()        
    if err != nil {
        return 0, err
    }
    
    defer func(ctx *rbIterationContext) {
        atomic.CompareAndSwapInt32(&ctx.state, iterWalking, iteratorReady)
    }(context)
    
    context.walkAll(tree.root)
    return context.CurrentCount(), nil
}

func (context *rbIterationContext) walkAll(node *RbNode) {
    if node == nil || !context.inWalk() {
        return
    }
    
    if node.left != nil {
        context.walkAll(node.left)
        if !context.inWalk() {
            return
        }
    }
    
    context.incrementCount()
    context.callback(context, node.key, node.value)
    if !context.inWalk() {
        return
    }
    
    if node.right != nil {
        context.walkAll(node.right)
    }    
}

func (context *rbIterationContext) Between(loKey RbKey, hiKey RbKey) (int, error) {
    if loKey == nil {
        return 0, ArgumentNilError("loKey")
    }
    if hiKey == nil {
        return 0, ArgumentNilError("hiKey")
    }

    tree, err := context.checkStateAndGetTree()        
    if err != nil {
        return 0, err
    }    

    defer func(ctx *rbIterationContext) {
        atomic.CompareAndSwapInt32(&ctx.state, iterWalking, iteratorReady)
    }(context)
    
    switch loKey.ComparedTo(hiKey) {
    case KeysAreEqual:
        node := tree.find(loKey)
        if node != nil {
            context.callback(context, node.key, node.value)
            return 1, nil
        }
        return 0, nil
    case KeyIsGreater:
        loKey, hiKey = hiKey, loKey
    }
    
    context.walkBetween(tree.root, loKey, hiKey)
    return context.CurrentCount(), nil
}

func (context *rbIterationContext) walkBetween(node *RbNode, loKey RbKey, hiKey RbKey) {
    if node == nil || !context.inWalk() {
        return
    }
    
    cmpLo := int8(loKey.ComparedTo(node.key))
    if cmpLo < zeroOrEqual {
        if node.left != nil {
            context.walkBetween(node.left, loKey, hiKey)
            if !context.inWalk() {
                return
            }
        }
    } 
    
    cmpHi := int8(hiKey.ComparedTo(node.key))
    if cmpLo <= zeroOrEqual && cmpHi >= zeroOrEqual {
        context.incrementCount()
        context.callback(context, node.key, node.value)
        if !context.inWalk() {
            return
        }
    } 
    
    if cmpHi > zeroOrEqual {
        if node.right != nil {
            context.walkBetween(node.right, loKey, hiKey)
        }    
    }
}

func (context *rbIterationContext) LessOrEqual(key RbKey) (int, error) {
        if key == nil {
        return 0, ArgumentNilError("key")
    }

    tree, err := context.checkStateAndGetTree()        
    if err != nil {
        return 0, err
    }    

    defer func(ctx *rbIterationContext) {
        atomic.CompareAndSwapInt32(&ctx.state, iterWalking, iteratorReady)
    }(context)
    
    context.walkLessOrEqual(tree.root, key)
    return context.CurrentCount(), nil
}

func (context *rbIterationContext) walkLessOrEqual(node *RbNode, key RbKey) {
    if node == nil || !context.inWalk() {
        return
    }
    
    if node.left != nil {
        context.walkLessOrEqual(node.left, key)
        if !context.inWalk() {
            return
        }
    }
    
    cmp := node.key.ComparedTo(key)
    if cmp == KeyIsLess || cmp == KeysAreEqual {
        context.incrementCount()
        context.callback(context, node.key, node.value)
        if !context.inWalk() {
            return
        }

        if node.right != nil {
            context.walkLessOrEqual(node.right, key)
        }  
    }
}

func (context *rbIterationContext) GreaterOrEqual(key RbKey) (int, error) {
    if key == nil {
        return 0, ArgumentNilError("key")
    }

    tree, err := context.checkStateAndGetTree()        
    if err != nil {
        return 0, err
    }    

    defer func(ctx *rbIterationContext) {
        atomic.CompareAndSwapInt32(&ctx.state, iterWalking, iteratorReady)
    }(context)
    
    context.walkGreaterOrEqual(tree.root, key)
    return context.CurrentCount(), nil
}

func (context *rbIterationContext) walkGreaterOrEqual(node *RbNode, key RbKey) {
    if node == nil || !context.inWalk() {
        return
    }
    
    cmp := node.key.ComparedTo(key)
    if cmp == KeyIsGreater || cmp == KeysAreEqual {
        if node.left != nil {
            context.walkGreaterOrEqual(node.left, key)
            if !context.inWalk() {
                return
            }
        }
        
        context.incrementCount()
        context.callback(context, node.key, node.value)
        if !context.inWalk() {
            return
        }
    }

    if node.right != nil {
        context.walkGreaterOrEqual(node.right, key)
    }    
}

func (context *rbIterationContext) LessThan(key RbKey) (int, error) {
    if key == nil {
        return 0, ArgumentNilError("key")
    }

    tree, err := context.checkStateAndGetTree()        
    if err != nil {
        return 0, err
    }    

    defer func(ctx *rbIterationContext) {
        atomic.CompareAndSwapInt32(&ctx.state, iterWalking, iteratorReady)
    }(context)
    
    context.walkLessThan(tree.root, key)
    return context.CurrentCount(), nil
}

func (context *rbIterationContext) walkLessThan(node *RbNode, key RbKey) {
    if node == nil || !context.inWalk() {
        return
    }
    
    if node.left != nil {
        context.walkLessThan(node.left, key)
        if !context.inWalk() {
            return
        }
    }
    
    if node.key.ComparedTo(key) == KeyIsLess {
        context.incrementCount()
        context.callback(context, node.key, node.value)
        if !context.inWalk() {
            return
        }

        if node.right != nil {
            context.walkLessThan(node.right, key)
        }  
    }
}

func (context *rbIterationContext) GreaterThan(key RbKey) (int, error) {
    if key == nil {
        return 0, ArgumentNilError("key")
    }

    tree, err := context.checkStateAndGetTree()        
    if err != nil {
        return 0, err
    }    

    defer func(ctx *rbIterationContext) {
        atomic.CompareAndSwapInt32(&ctx.state, iterWalking, iteratorReady)
    }(context)
    
    context.walkGreaterThan(tree.root, key)
    return context.CurrentCount(), nil
}

func (context *rbIterationContext) walkGreaterThan(node *RbNode, key RbKey) {
    if node == nil || !context.inWalk() {
        return
    }
    
    if node.key.ComparedTo(key) == KeyIsGreater {
        if node.left != nil {
            context.walkGreaterThan(node.left, key)
            if !context.inWalk() {
                return
            }
        }
        
        context.incrementCount()
        context.callback(context, node.key, node.value)
        if !context.inWalk() {
            return
        }
    }

    if node.right != nil {
        context.walkGreaterThan(node.right, key)
    }    
}