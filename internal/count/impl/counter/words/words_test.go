package words

import "testing"

func TestWords_Exec(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Words Success",
			args: args{
				text: "Hello I'm Jorge",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Words{}
			if got := w.Exec(tt.args.text); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
