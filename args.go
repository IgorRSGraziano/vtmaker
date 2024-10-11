package main

import (
	"flag"
	"log"
	"os"
	"path"
)

type Args struct {
	// --audio || -a
	AudioPath string
	// --gif || -g
	GifPath string
	// --output || -o
	OutputPath string
	// --subtitle || -s
	SubtitlePath string

	// --help || -h
	Help bool

	// --folder || -f
	FolderPath string

	//--output || -o
	// OutputPath string
}

func setDefaultArgs(args *Args) {
	cwd, _ := os.Getwd()
	args.OutputPath = path.Join(cwd, "output.mp4")

	if args.FolderPath != "" {
		files, err := os.ReadDir(args.FolderPath)
		if err != nil {
			log.Printf("Error reading folder %s: %s", args.FolderPath, err)
			panic(err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			switch path.Ext(file.Name()) {
			case ".srt":
				args.SubtitlePath = path.Join(args.FolderPath, file.Name())
			case ".mp3":
				args.AudioPath = path.Join(args.FolderPath, file.Name())
			case ".gif":
				args.GifPath = path.Join(args.FolderPath, file.Name())
			}
		}
	}
}

func GetArgs() *Args {

	// Inicializa a estrutura Args
	args := &Args{}

	// Definindo as flags
	flag.StringVar(&args.AudioPath, "audio", "", "Path to the audio file")
	flag.StringVar(&args.AudioPath, "a", "", "Path to the audio file (shorthand)")

	flag.StringVar(&args.GifPath, "gif", "", "Path to the gif file")
	flag.StringVar(&args.GifPath, "g", "", "Path to the gif file (shorthand)")

	flag.StringVar(&args.OutputPath, "output", "", "Path to the output file")
	flag.StringVar(&args.OutputPath, "o", "", "Path to the output file (shorthand)")

	flag.StringVar(&args.SubtitlePath, "subtitle", "", "Path to the subtitle file")
	flag.StringVar(&args.SubtitlePath, "s", "", "Path to the subtitle file (shorthand)")

	flag.BoolVar(&args.Help, "help", false, "Show usage information")
	flag.BoolVar(&args.Help, "h", false, "Show usage information (shorthand)")

	flag.StringVar(&args.FolderPath, "folder", "", "Path to the folder containing files")
	flag.StringVar(&args.FolderPath, "f", "", "Path to the folder containing files (shorthand)")

	flag.Parse()

	if args.Help {
		flag.Usage()
		os.Exit(0)
	}

	setDefaultArgs(args)

	return args
}
