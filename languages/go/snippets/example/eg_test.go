package example

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		to string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "World",
			args: args{to: "World"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hello(tt.args.to)
		})
	}
}
