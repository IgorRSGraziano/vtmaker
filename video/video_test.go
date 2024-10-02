package video_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"vtmaker/video"
)

func TestGetDuration(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	file := path.Join(dir, "gif", "_testData", "sample.gif")

	duration, err := video.GetDuration(file)

	if err != nil {
		t.Fatal(err)
	}

	if duration != 3.36 {
		t.Fatalf("Expected 2.0, got %f", duration)
	}
}
