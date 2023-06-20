package mock

type Invocation struct {
	matchFunc any
}

func (i *Invocation) Match(input []any) ([]any, bool) {
	panic("TODO")
}

func (i *Invocation) SetStub(stub any) {

}

func (i *Invocation) SetReturnValues(values []any) {

}

func (i *Invocation) SetCallCount(count int) {

}
