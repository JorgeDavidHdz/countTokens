package letters

import "testing"

func TestLetters_Exec(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Letters Success",
			args: args{
				text: "Hello I'm Jorge",
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Letters{}
			if got := l.Exec(tt.args.text); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
