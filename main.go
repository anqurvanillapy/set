package set

// USet is an unordered set based on the official map implementation.
type USet map[interface{}]struct{}

// NewUSet creats a preallocated unordered set.
func NewUSet() USet {
	return make(USet)
}

// Add a new key to the set.  Return value for whether it already exists.
func (us USet) Add(k interface{}) bool {
	_, ok := us[k]
	us[k] = struct{}{}
	return ok
}

// Has a specific key in the set?
func (us USet) Has(k interface{}) bool {
	_, ok := us[k]
	return ok
}

// Delete a key.  Note that `delete' is quicker than assigning its zero value.
func (us USet) Delete(k interface{}) {
	delete(us, k)
}
