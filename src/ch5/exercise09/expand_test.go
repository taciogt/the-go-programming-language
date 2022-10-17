package exercise09

import "testing"

func Test_expand(t *testing.T) {
	type args struct {
		s string
		f func(string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "function that does nothing",
		args: args{
			s: "string",
			f: func(s string) string { return s },
		},
		want: "string",
	}, {
		name: "substitute a string for itself",
		args: args{
			s: "this is a string $foo",
			f: func(s string) string {
				return s
			},
		},
		want: "this is a string foo",
	}, {
		name: "nil substitution function",
		args: args{
			s: "this is a string $foo",
			f: nil,
		},
		want: "this is a string $foo",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}
