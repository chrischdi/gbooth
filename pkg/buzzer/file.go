package buzzer

import "os"

type FileBuzzer struct {
	TrueBuzzer
	File string
}

func (b *FileBuzzer) Pressed() bool {
	_, err := os.Stat(b.File)
	return err != nil
}
