package video

import (
	"errors"

	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// Aumenta a duracao do gif, 4 é a quantidade de vezes necessária (precisa ser calculada), 10 é o tempo em segundos
// ffmpeg -stream_loop 4 -i sample.gif -t 10 out.gif;

//Transforma gif em mp4
// ffmpeg -i sample.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" out.mp4;

//Processo:
// 1. Aumentar a duracao do gif
// 2. Transformar gif em mp4
// 3. Adicionar audio ao mp4

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

func AddAudioToVideo(inputVideoPath string, inputAudioPath string, outputPath string) error {

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
	}).Run()

	return err
}

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

func AddSubtitleToVideo(inputVideoPath string, inputSubtitlePath string, outputPath string) error {
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

	// ffmpeg -vf subtitles=/home/prplexx/projects/vtmaker/video/gif/_testData/subtitle.srt -i /home/prplexx/projects/vtmaker/video/gif/_testData/output.mp4 /home/prplexx/projects/vtmaker/video/gif/_testData/final_output.mp4

	err := ffmpeg.Input(inputVideoPath).Output(outputPath, ffmpeg.KwArgs{
		"vf": "subtitles=" + inputSubtitlePath,
	}).Run()

	return err
}

func NormalizeVideo(inputPath string, outputPath string) error {
	// ffmpeg -i <input> -c:v libx264 -profile:v baseline -level 3.0 -pix_fmt yuv420p <output>
	err := ffmpeg.Input(inputPath).Output(outputPath, ffmpeg.KwArgs{
		"c:v":       "libx264",
		"profile:v": "baseline",
		"level":     "3.0",
		"pix_fmt":   "yuv420p",
	}).Run()

	return err
}
