package set

// USet is an unordered set based on the official map implementation.
type USet map[interface{}]bool

// NewUSet creats a preallocated unordered set.
func NewUSet() USet {
	return make(USet)
}

// Add a new key to the set.  Return value for whether it already exists.
func (us USet) Add(k interface{}) bool {
	exists := us[k]
	us[k] = true
	return exists
}

// Has a specific key in the set?
func (us USet) Has(k interface{}) bool {
	return us[k]
}

// Delete a key.  Note that `delete' is quicker than assigning its zero value.
func (us USet) Delete(k interface{}) {
	delete(us, k)
}
