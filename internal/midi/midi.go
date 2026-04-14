package midi

var notesToNum = map[string]int{
	"C":  0,
	"C#": 1,
	"Db": 1,
	"D":  2,
	"D#": 3,
	"Eb": 3,
	"E":  4,
	"F":  5,
	"F#": 6,
	"Gb": 6,
	"G":  7,
	"G#": 8,
	"Ab": 8,
	"A":  9,
	"A#": 10,
	"Bb": 10,
	"B":  11,
}

// NoteToNum maps a note name and octave to its assigned numeric value. Per the
// MIDI specification middle C has a reference value of 60. Switchboard follows
// the convention used by most manufacturers where middle C is defined as C3,
// meaning the lowest possible MIDI note (note #0) is called C–2.
func NoteToNum(name string, octave int) int {
	key, ok := notesToNum[name]
	if !ok {
		panic("unknown note: " + name)
	}
	return key + (octave+2)*12
}
