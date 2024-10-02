package video

import (
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

func CreateVideoFromGif(inputPath string, duration uint, outputPath string) error {
	return nil
}
