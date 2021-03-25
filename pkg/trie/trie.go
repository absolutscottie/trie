package trie

const (
	MaxKeyLen = 128
)

type TrieNode struct {
	Value    interface{}
	Children map[byte]*TrieNode //alternatively this could be an array
}

func NewTrie() *TrieNode {
	return &TrieNode{
		Children: make(map[byte]*TrieNode),
	}
}

//	validateKey checks to make sure that the key is valid (less than 128 bytes) and returns a
//	byte array representation of the key. returns ErrorKeyInvalid if the key does not conform
func validateKey(key string) ([]byte, error) {
	//validate that the key is 128 bytes or fewer. This could cause unexpected results with
	//unicode since they will be split into several bytes
	keyBytes := []byte(key)
	if len(keyBytes) > MaxKeyLen {
		return nil, ErrorInvalidKey
	}
	return keyBytes, nil
}

//	childExists checks whether the child associated with the provided byte exists or not
func (t *TrieNode) childExists(keyByte byte) bool {
	if _, ok := t.Children[keyByte]; !ok {
		return false
	}
	return true
}

//	Insert will store the provided value at the location determined by key
func (t *TrieNode) Insert(key string, value interface{}) error {
	keyBytes, err := validateKey(key)
	if err != nil {
		return err
	}

	return t.insert(keyBytes, value)
}

//	insert is the private implementation of the insertion algorithm
func (t *TrieNode) insert(key []byte, value interface{}) error {
	if len(key) == 0 {
		t.Value = value
		return nil
	}

	if !t.childExists(key[0]) {
		t.Children[key[0]] = &TrieNode{
			Children: make(map[byte]*TrieNode),
		}
	}

	return t.Children[key[0]].insert(key[1:], value)
}

//	Find will return the value at the requested location if found, otherwise it will return an non
// 	nil error
func (t *TrieNode) Find(key string) (interface{}, error) {
	keyBytes, err := validateKey(key)
	if err != nil {
		return nil, err
	}

	return t.find(keyBytes)
}

//	find is the private implementation of the find algorithm
func (t *TrieNode) find(key []byte) (interface{}, error) {
	if len(key) == 0 {
		return t.Value, nil
	}

	if !t.childExists(key[0]) {
		return nil, ErrorNotFound
	}

	return t.Children[key[0]].find(key[1:])
}
