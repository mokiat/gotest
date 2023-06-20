package mock

func NewRegistry() *Registry {
	return &Registry{
		entries: make(map[string][]*Invocation),
	}
}

type Registry struct {
	entries map[string][]*Invocation
}

func (r *Registry) Match(name string, matchFunc any) *Invocation {
	invocation := &Invocation{
		matchFunc: matchFunc,
	}
	r.entries[name] = append(r.entries[name], invocation)
	return invocation
}

func (r *Registry) Invoke(name string, input []any) []any {
	for _, invocation := range r.entries[name] {
		if result, ok := invocation.Match(input); ok {
			return result
		}
	}
	return nil
}
