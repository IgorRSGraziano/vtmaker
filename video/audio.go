package video

import (
	"errors"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AddAudio(inputVideoPath string, inputAudioPath string, outputPath string) error {

	if inputAudioPath == "" {
		return errors.New("audio path is empty")
	}

	if inputVideoPath == "" {
		return errors.New("video path is empty")
	}

	if outputPath == "" {
		return errors.New("output path is empty")
	}

	if outputPath == inputVideoPath || outputPath == inputAudioPath {
		return errors.New("output path is the same as input path")
	}

	err := ffmpeg.Input(inputAudioPath, ffmpeg.KwArgs{
		"i": inputVideoPath,
	}).Output(outputPath, ffmpeg.KwArgs{
		"map":      []string{"0:v", "1:a"},
		"c":        "copy",
		"shortest": "",
	}).ErrorToStdOut().Run()

	return err
}
