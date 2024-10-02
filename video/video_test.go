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
		t.Fatalf("Expected 2.0, got %f", duration)
	}
}

func TestCreateVideoFromGif(t *testing.T) {
	outputPath := path.Join(testDir, "output.mp4")

	defer os.Remove(outputPath)

	var duration uint = 120

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
		t.Fatalf("Expected %d, got %f", duration, outputDuration)
	}
}
