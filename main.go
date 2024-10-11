package main

import (
	"fmt"
	"log"
	"os"
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

	alreadyExists := file.Exists(args.OutputPath)

	if alreadyExists {
		log.Printf("File %s already exists", args.OutputPath)
		log.Printf("Do you want to overwrite it? [y/n]")
		var answer string
		fmt.Scanln(&answer)

		if answer != "y" {
			log.Println("Exiting...")
			return
		}

		err = os.Remove(args.OutputPath)

		if err != nil {
			log.Printf("Error removing file: %v", err)
			panic(err)
		}
	}

	err = file.Copy(video.OutputPath, args.OutputPath)

	if err != nil {
		log.Printf("Error moving final output: %v", err)
		panic(err)
	}

}
