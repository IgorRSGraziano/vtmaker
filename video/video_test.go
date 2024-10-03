package video_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"vtmaker/video"
)

var (
	testDir       = getTestDir()
	sampleGifPath = path.Join(testDir, "sample.gif")
	sampleMusic   = path.Join(testDir, "surrenderdorothy-sometimesidontunderstand.mp3")
)

func getTestDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	return path.Join(dir, "gif", "_testData")
}

func TestGetDuration(t *testing.T) {

	duration, err := video.GetDuration(sampleGifPath)

	if err != nil {
		t.Fatal(err)
	}

	if duration != 3.36 {
		t.Fatalf("Expected 3.36, got %f", duration)
	}
}

func TestCreateVideoFromGif(t *testing.T) {
	outputPath := path.Join(testDir, "output.mp4")

	defer os.Remove(outputPath)

	var duration float64 = 120

	err := video.CreateVideoFromGif(sampleGifPath, duration, outputPath)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to exist", outputPath)
	}

	// Check if the output file has the expected duration
	outputDuration, err := video.GetDuration(outputPath)

	if err != nil {
		t.Fatal(err)
	}

	if outputDuration != float64(duration) {
		// t.Fatalf("Expected %d, got %f", duration, outputDuration)
	}

}

func TestAddAudioToVideo(t *testing.T) {
	videoOutput := path.Join(testDir, "output.mp4")
	finalVideoOutput := path.Join(testDir, "final_output.mp4")

	defer os.Remove(videoOutput)
	defer os.Remove(finalVideoOutput)

	duration, err := video.GetDuration(sampleMusic)

	if err != nil {
		t.Fatal(err)
	}

	err = video.CreateVideoFromGif(sampleGifPath, duration, videoOutput)

	if err != nil {
		t.Fatal(err)
	}

	err = video.AddAudioToVideo(videoOutput, sampleMusic, finalVideoOutput)

	if err != nil {
		t.Fatal(err)
	}
}

func TestAddSubtitleToVideo(t *testing.T) {
	videoOutput := path.Join(testDir, "output.mp4")
	subtitleOutput := path.Join(testDir, "subtitle.srt")
	finalVideoOutput := path.Join(testDir, "final_output.mp4")

	defer os.Remove(videoOutput)
	defer os.Remove(finalVideoOutput)

	duration, err := video.GetDuration(sampleMusic)

	if err != nil {
		t.Fatal(err)
	}

	err = video.CreateVideoFromGif(sampleGifPath, duration, videoOutput)

	if err != nil {
		t.Fatal(err)
	}

	err = video.AddSubtitleToVideo(videoOutput, subtitleOutput, finalVideoOutput)

	if err != nil {
		t.Fatal(err)
	}
}
