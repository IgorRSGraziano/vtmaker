package main

import (
	"fmt"
	"log"
	"vtmaker/file"
	"vtmaker/video"
)

func main() {
	args := GetArgs()
	fmt.Println(args)

	video, err := video.NewVideo(args.SubtitlePath, args.AudioPath, args.GifPath)

	if err != nil {
		panic(err)
	}

	defer video.Flush()

	video.CreateFromGif().AddSubtitle().AddAudio().Normalize()

	if video.Error != nil {
		panic(video.Error)
	}

	err = file.Copy(video.OutputPath, args.OutputPath)

	if err != nil {
		log.Printf("Error moving final output: %v", err)
		panic(err)
	}

}
