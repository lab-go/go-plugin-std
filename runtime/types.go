package runtime

type Context struct {
	Name string
}

type PluginFactory interface {
	Create(conf []byte) (Filter, interface{})
}

type Filter interface {
	Execute(ctx Context) error
}
