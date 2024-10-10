package main

import (
	"os"
	"path"
	"vtmaker/random"
	"vtmaker/video"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	dataDir := path.Join(cwd, "data")

	var strFilePath, mp3FilePath, gifFilePath string
	//each and find srt file, mp3 file, and gif

	files, err := os.ReadDir(dataDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		switch path.Ext(file.Name()) {
		case ".srt":
			strFilePath = path.Join(dataDir, file.Name())
		case ".mp3":
			mp3FilePath = path.Join(dataDir, file.Name())
		case ".gif":
			gifFilePath = path.Join(dataDir, file.Name())
		}
	}

	if strFilePath == "" {
		panic("subtitle file not found")
	}
	if mp3FilePath == "" {
		panic("mp3 file not found")
	}
	if gifFilePath == "" {
		panic("gif file not found")
	}

	tempFolder := path.Join(dataDir, "temp")

	defer os.RemoveAll(tempFolder)

	if err := os.Mkdir(tempFolder, 0755); err != nil {
		panic(err)
	}

	musicDuration, err := video.GetDuration(mp3FilePath)

	if err != nil {
		panic(err)
	}

	createTempFile := func(format string) string {
		return random.NewSHA1Hash(8) + "." + format
	}

	gifVideoPath := path.Join(tempFolder, createTempFile("mp4"))

	if err := video.CreateVideoFromGif(gifFilePath, musicDuration, gifVideoPath); err != nil {
		panic(err)
	}

	videoSubtitlePath := path.Join(tempFolder, createTempFile("mp4"))

	if err := video.AddSubtitleToVideo(gifVideoPath, strFilePath, videoSubtitlePath); err != nil {
		panic(err)
	}

	finalOutputPath := path.Join(tempFolder, createTempFile("mp4"))
	if err := video.AddAudioToVideo(videoSubtitlePath, mp3FilePath, finalOutputPath); err != nil {
		panic(err)
	}

	finalFile := path.Join(dataDir, "final_output.mp4")
	if err := video.NormalizeVideo(finalOutputPath, finalFile); err != nil {
		panic(err)
	}

}
