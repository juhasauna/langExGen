package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"makeQuietAudio", makeQuietAudioTest},
		{"generateMp3", generateMp3Test},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func generateMp3Test(t *testing.T) {
	if err := generateMp3(); err != nil {
		t.Error(err)
	}
}

func makeQuietAudioTest(t *testing.T) {
	println(len(makeQuietAudio(11000)))
}
