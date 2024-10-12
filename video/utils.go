package video

import (
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetDuration(inputPath string) (float64, error) {
	out, err := ffmpeg.Probe(inputPath, []ffmpeg.KwArgs{
		{"show_entries": "format=duration"},
		{"v": "quiet"},
	}...)

	if err != nil {
		return 0, err
	}

	return gjson.Get(out, "format.duration").Float(), nil
}

func Normalize(inputPath string, outputPath string) error {
	err := ffmpeg.Input(inputPath).Output(outputPath, ffmpeg.KwArgs{
		"vf":        "pad=ceil(iw/2)*2:ceil(ih/2)*2",
		"c:v":       "libx264",
		"profile:v": "baseline",
		"level":     "3.0",
		"pix_fmt":   "yuv420p",
	}).ErrorToStdOut().Run()

	return err
}
