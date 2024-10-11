package video

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func CreateVideoFromGif(inputPath string, duration float64, outputPath string) error {
	inputDuration, err := GetDuration(inputPath)

	if err != nil {
		return err
	}

	loopCount := uint64(duration / inputDuration)
	err = ffmpeg.Input(inputPath, ffmpeg.KwArgs{
		"stream_loop": loopCount,
	}).Output(outputPath, ffmpeg.KwArgs{
		"t": duration,
	}).Run()

	return err
}
