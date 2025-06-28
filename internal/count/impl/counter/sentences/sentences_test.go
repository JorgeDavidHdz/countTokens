package sentences

import (
	"testing"
)

func TestSentences_Exec(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Sentences Success",
			args: args{
				text: "Hi, I'm Jorge. I'm 26 years old.",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sentences{}
			if got := s.Exec(tt.args.text); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
