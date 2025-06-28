package numbers

import "testing"

func TestNumbers_Exec(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Numbers Success",
			args: args{
				text: "H3ll0 1'm J0rg3",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Numbers{}
			if got := n.Exec(tt.args.text); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
