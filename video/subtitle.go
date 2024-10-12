package video

import (
	"errors"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AddSubtitle(inputVideoPath string, inputSubtitlePath string, outputPath string) error {
	if inputSubtitlePath == "" {
		return errors.New("subtitle path is empty")
	}

	if inputVideoPath == "" {
		return errors.New("video path is empty")
	}

	if outputPath == "" {
		return errors.New("output path is empty")
	}

	if outputPath == inputVideoPath || outputPath == inputSubtitlePath {
		return errors.New("output path is the same as input path")
	}

	err := ffmpeg.Input(inputVideoPath).Output(outputPath, ffmpeg.KwArgs{
		"vf": "subtitles=" + inputSubtitlePath,
	}).ErrorToStdOut().Run()

	return err
}
