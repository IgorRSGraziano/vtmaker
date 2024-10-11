package video

import (
	"log"
	"os"
	"path"
)

type Video struct {
	SrtFilePath string
	Mp3FilePath string
	GifFilePath string
	OutputPath  string
	tmpFolder   string
	Error       error
}

func NewVideo(srtFilePath, mp3FilePath, gifFilePath string) (*Video, error) {
	if !fileExists(srtFilePath) {
		log.Printf("Subtitle file not found: %s", srtFilePath)
		return nil, os.ErrNotExist
	}

	if !fileExists(mp3FilePath) {
		log.Printf("MP3 file not found: %s", mp3FilePath)
		return nil, os.ErrNotExist
	}

	if !fileExists(gifFilePath) {
		log.Printf("GIF file not found: %s", gifFilePath)
		return nil, os.ErrNotExist
	}

	tmpFolder := path.Join(os.TempDir(), "vtmaker")

	if err := os.MkdirAll(tmpFolder, 0755); err != nil {
		log.Printf("Error creating temp folder: %v", err)
		return nil, err
	}

	return &Video{
		SrtFilePath: srtFilePath,
		Mp3FilePath: mp3FilePath,
		GifFilePath: gifFilePath,
		tmpFolder:   tmpFolder,
	}, nil
}

func (v *Video) Flush() {
	os.RemoveAll(v.tmpFolder)
}

func (v *Video) CreateFromGif() *Video {
	if v.Error != nil {
		return v
	}

	outputPath := path.Join(v.tmpFolder, "giftovid.mp4")

	duration, err := GetDuration(v.Mp3FilePath)

	if err != nil {
		log.Printf("Error getting duration of gif: %v", err)
		v.Error = err
		return v
	}

	err = CreateFromGif(v.GifFilePath, duration, outputPath)

	if err != nil {
		log.Printf("Error creating video from gif: %v", err)
		v.Error = err
		return v
	}

	v.OutputPath = outputPath

	return v
}

func (v *Video) AddSubtitle() *Video {
	if v.Error != nil {
		return v
	}

	outputPath := path.Join(v.tmpFolder, "subtitled.mp4")

	err := AddSubtitle(v.OutputPath, v.SrtFilePath, outputPath)

	if err != nil {
		log.Printf("Error adding subtitle to video: %v", err)
		v.Error = err
		return v
	}

	v.OutputPath = outputPath

	return v
}

func (v *Video) Normalize() *Video {
	if v.Error != nil {
		return v
	}

	outputPath := path.Join(v.tmpFolder, "normalized.mp4")

	err := Normalize(v.OutputPath, outputPath)

	if err != nil {
		log.Printf("Error normalizing video: %v", err)
		v.Error = err
		return v
	}

	v.OutputPath = outputPath

	return v
}

func (v *Video) AddAudio() *Video {
	if v.Error != nil {
		return v
	}

	outputPath := path.Join(v.tmpFolder, "audioadded.mp4")

	err := AddAudio(v.OutputPath, v.Mp3FilePath, outputPath)

	if err != nil {
		log.Printf("Error adding audio to video: %v", err)
		v.Error = err
		return v
	}

	v.OutputPath = outputPath

	return v
}
