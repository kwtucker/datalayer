package datalayer

import "sync"

// Registry is a container for objects.
type Registry struct {
	sync.Mutex
	dataStores map[string]*Object
}

// Object represents an object and its metadata in a registry
type Object struct {
	Value   DataStorer
	Name    string
	Type    string
	Enabled bool
}

func (r *Registry) initDataStores() {
	if r.dataStores == nil {
		r.dataStores = make(map[string]*Object)
	}
}

// Register adds objects to the registry.
func (r *Registry) Register(ds ...*Object) {
	r.Lock()
	defer r.Unlock()
	r.initDataStores()

	for _, obj := range ds {
		name := obj.Name
		r.dataStores[name] = obj
		r.dataStores[name].Enabled = true
	}
}

// Lookup returns the requested object or nil if it is not registerd.
func (r *Registry) Lookup(name string) *Object {
	r.Lock()
	defer r.Unlock()

	if obj, ok := r.dataStores[name]; ok {
		if obj.Enabled {
			return obj
		}
	}
	return nil
}

// RegisteredModulesState returns a map with module name and if it is enabled.
func (r *Registry) RegisteredModulesState() map[string]bool {
	r.Lock()
	defer r.Unlock()
	state := map[string]bool{}

	for key, val := range r.dataStores {
		state[key] = val.Enabled
	}
	return state
}

// Enable enables the module names passed in if they exist.
func (r *Registry) Enable(names ...string) {
	r.Lock()
	defer r.Unlock()

	for _, name := range names {
		if obj, ok := r.dataStores[name]; ok {
			obj.Enabled = true
		}
	}
}

// Disable disables the module names passed in if they exist.
func (r *Registry) Disable(names ...string) {
	r.Lock()
	defer r.Unlock()

	for _, name := range names {
		if obj, ok := r.dataStores[name]; ok {
			obj.Enabled = false
		}
	}
}
