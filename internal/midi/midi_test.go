package midi

import (
	"fmt"
	"testing"
)

func TestNoteToNum(t *testing.T) {
	type args struct {
		name   string
		octave int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"minimum note", args{"C", -2}, 0},
		{"middle C", args{"C", 3}, 60},
		{"sharp note", args{"C#", 3}, 61},
		{"flat note", args{"Db", 3}, 61},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s%d", tt.args.name, tt.args.octave), func(t *testing.T) {
			if got := NoteToNum(tt.args.name, tt.args.octave); got != tt.want {
				t.Errorf("NoteToNum(%s, %d) = %v, want %v", tt.args.name, tt.args.octave, got, tt.want)
			}
		})
	}
}
