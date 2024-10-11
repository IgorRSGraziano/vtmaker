package main

import (
	"log"
	"os"
	"path"
	"vtmaker/file"
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

	video, err := video.NewVideo(strFilePath, mp3FilePath, gifFilePath)

	if err != nil {
		panic(err)
	}

	// defer video.Flush()

	video.CreateFromGif().AddSubtitle().AddAudio().Normalize()

	if video.Error != nil {
		panic(video.Error)
	}

	err = file.Copy(video.OutputPath, path.Join(dataDir, "final_output.mp4"))
	if err != nil {
		log.Printf("Error moving final output: %v", err)
		panic(err)
	}

	// musicDuration, err := video.GetDuration(mp3FilePath)

	// if err != nil {
	// 	panic(err)
	// }

	// createTempFile := func(format string) string {
	// 	return random.NewSHA1Hash(8) + "." + format
	// }

	// gifVideoPath := path.Join(tempFolder, createTempFile("mp4"))

	// if err := video.CreateFromGif(gifFilePath, musicDuration, gifVideoPath); err != nil {
	// 	panic(err)
	// }

	// videoSubtitlePath := path.Join(tempFolder, createTempFile("mp4"))

	// if err := video.AddSubtitle(gifVideoPath, strFilePath, videoSubtitlePath); err != nil {
	// 	panic(err)
	// }

	// finalOutputPath := path.Join(tempFolder, createTempFile("mp4"))
	// if err := video.AddAudio(videoSubtitlePath, mp3FilePath, finalOutputPath); err != nil {
	// 	panic(err)
	// }

	// finalFile := path.Join(dataDir, "final_output.mp4")
	// if err := video.Normalize(finalOutputPath, finalFile); err != nil {
	// 	panic(err)
	// }

}
