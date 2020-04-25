package runtime

type Engine struct {
	Loader *Loader
}

func (e *Engine) Load(name, path string, conf []byte) {
	e.Loader.Load(name, path, conf)
}

func (e *Engine) Run(ctx Context) {
	for _, filter := range e.Loader.Filters {
		filter.Execute(ctx)
	}
}
