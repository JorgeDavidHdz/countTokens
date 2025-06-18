package symbols

import "testing"

func TestSymbols_Exec(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Symbols Success",
			args: args{
				text: "Hello I'm Jorg#",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Symbols{}
			if got := s.Exec(tt.args.text); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
