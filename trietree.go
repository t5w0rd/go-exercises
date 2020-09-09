package exercises

type trieTreeNode struct {
	children [26]*trieTreeNode
	exists   interface{}
}

type TrieTree struct {
	root     *trieTreeNode
	size     int
	pool     []trieTreeNode
	poolUsed int
}

func NewTrieTree() *TrieTree {
	pool := make([]trieTreeNode, 100)
	return &TrieTree{
		root:     &pool[0],
		pool:     pool,
		poolUsed: 1,
	}
}

func (t *TrieTree) Set(s string) (exist bool) {
	where := &t.root
	for _, c := range s {
		index := c - 'a'
		where = &(*where).children[index]
		if *where == nil {
			*where = &trieTreeNode{}
		}
	}

	if (*where).exists == struct{}{} {
		return true
	}

	(*where).exists = struct{}{}
	t.size++
	return false
}

func (t *TrieTree) Set2(s string) (exist bool) {
	where := &t.root
	for _, c := range s {
		index := c - 'a'
		where = &(*where).children[index]
		if *where == nil {
			*where = &t.pool[t.poolUsed]
			t.poolUsed++
		}
	}

	if (*where).exists == struct{}{} {
		return true
	}

	(*where).exists = struct{}{}
	return false
}

func (t *TrieTree) Get(s string) (exist bool) {
	where := t.root
	for _, c := range s {
		index := c - 'a'
		where = where.children[index]
		if where == nil {
			return false
		}
	}
	return where.exists == struct{}{}
}
