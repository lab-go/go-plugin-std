package runtime

import (
	"encoding/json"
	"log"
	"plugin"
)

type Loader struct {
	Filters map[string]Filter
	Config  map[string][]byte
}

func (l *Loader) Load(name, path string, conf []byte) {
	pdll, err := plugin.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	// 在插件中寻找相关的对象，将其方法加载
	v, err := pdll.Lookup("Builder")

	if err != nil {
		log.Fatalln(err)
	}

	if vp, ok := v.(func() PluginFactory); ok {
		factory := vp()

		f, config := factory.Create(conf)

		l.Filters[name] = f
		marshal, _ := json.Marshal(config)
		l.Config[name] = marshal
	}
}
