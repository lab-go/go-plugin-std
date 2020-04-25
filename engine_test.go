package go_plugin_std

import (
	"github.com/lab-go/go-plugin-std/runtime"
	"testing"
)

func TestEngine_Run(t *testing.T) {
	type fields struct {
		loader *runtime.Loader
	}
	type args struct {
		ctx  runtime.Context
		name string
		path string
		conf []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "1",
			fields: fields{
				loader: &runtime.Loader{
					Filters: map[string]runtime.Filter{},
					Config:  map[string][]byte{},
				},
			},
			args: args{
				ctx: runtime.Context{
					Name: "test",
				},
				name: "p1",
				path: "./demo/demo.so",
				conf: []byte(`{"name": "xx"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &runtime.Engine{
				Loader: tt.fields.loader,
			}

			e.Load(tt.args.name, tt.args.path, tt.args.conf)
			e.Run(tt.args.ctx)
		})
	}
}
