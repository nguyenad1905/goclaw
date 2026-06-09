package media

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

// createTestPNG creates a small PNG image file and returns its path.
func createTestPNG(t *testing.T, width, height int) string {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatalf("png.Encode: %v", err)
	}
	path := filepath.Join(t.TempDir(), "test.png")
	if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	return path
}

// createTestJPEG creates a small JPEG image file and returns its path.
func createTestJPEG(t *testing.T, width, height int) string {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90}); err != nil {
		t.Fatalf("jpeg.Encode: %v", err)
	}
	path := filepath.Join(t.TempDir(), "test.jpg")
	if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	return path
}

func TestPreserveFormat_PNG(t *testing.T) {
	srcPath := createTestPNG(t, 100, 100)

	outPath, outMime, err := PreserveFormat(srcPath)
	if err != nil {
		t.Fatalf("PreserveFormat: %v", err)
	}
	defer os.Remove(outPath)

	if outMime != "image/png" {
		t.Errorf("MIME = %q, want %q", outMime, "image/png")
	}
	if ext := filepath.Ext(outPath); ext != ".png" {
		t.Errorf("extension = %q, want .png", ext)
	}

	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !bytes.HasPrefix(data, []byte("\x89PNG")) {
		t.Errorf("output is not valid PNG, first 4 bytes: %x", data[:4])
	}
}

func TestPreserveFormat_JPEG(t *testing.T) {
	srcPath := createTestJPEG(t, 100, 100)

	outPath, outMime, err := PreserveFormat(srcPath)
	if err != nil {
		t.Fatalf("PreserveFormat: %v", err)
	}
	defer os.Remove(outPath)

	if outMime != "image/jpeg" {
		t.Errorf("MIME = %q, want %q", outMime, "image/jpeg")
	}
	if ext := filepath.Ext(outPath); ext != ".jpg" {
		t.Errorf("extension = %q, want .jpg", ext)
	}

	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !bytes.HasPrefix(data, []byte("\xff\xd8\xff")) {
		t.Errorf("output is not valid JPEG, first 3 bytes: %x", data[:3])
	}
}

func TestSanitizeForVision_PNG(t *testing.T) {
	srcPath := createTestPNG(t, 100, 100)

	data, outMime, err := SanitizeForVision(srcPath, "image/png")
	if err != nil {
		t.Fatalf("SanitizeForVision: %v", err)
	}

	if outMime != "image/png" {
		t.Errorf("MIME = %q, want %q", outMime, "image/png")
	}
	if !bytes.HasPrefix(data, []byte("\x89PNG")) {
		t.Errorf("output is not valid PNG, first 4 bytes: %x", data[:4])
	}
}

func TestSanitizeForVision_JPEG(t *testing.T) {
	srcPath := createTestJPEG(t, 100, 100)

	data, outMime, err := SanitizeForVision(srcPath, "image/jpeg")
	if err != nil {
		t.Fatalf("SanitizeForVision: %v", err)
	}

	if outMime != "image/jpeg" {
		t.Errorf("MIME = %q, want %q", outMime, "image/jpeg")
	}
	if !bytes.HasPrefix(data, []byte("\xff\xd8\xff")) {
		t.Errorf("output is not valid JPEG, first 3 bytes: %x", data[:3])
	}
}

func TestSanitizeForVision_LargeImageResize(t *testing.T) {
	srcPath := createTestJPEG(t, 2000, 2000)

	data, _, err := SanitizeForVision(srcPath, "image/jpeg")
	if err != nil {
		t.Fatalf("SanitizeForVision: %v", err)
	}

	if len(data) > ImageSanitizeMaxBytes {
		t.Errorf("output size %d exceeds max %d", len(data), ImageSanitizeMaxBytes)
	}
}

func TestSanitizeForVision_PNGAlphaPreserved(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 10, 10))
	img.SetNRGBA(5, 5, color.NRGBA{R: 255, G: 0, B: 0, A: 128})
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatalf("png.Encode: %v", err)
	}
	srcPath := filepath.Join(t.TempDir(), "alpha.png")
	if err := os.WriteFile(srcPath, buf.Bytes(), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	data, outMime, err := SanitizeForVision(srcPath, "image/png")
	if err != nil {
		t.Fatalf("SanitizeForVision: %v", err)
	}

	if outMime != "image/png" {
		t.Errorf("MIME = %q, want %q (alpha should be preserved)", outMime, "image/png")
	}
	if !bytes.HasPrefix(data, []byte("\x89PNG")) {
		t.Errorf("output is not valid PNG")
	}
}

func TestDetectImageFormat(t *testing.T) {
	tests := []struct {
		name   string
		create func(t *testing.T) string
		want   string
	}{
		{"png", func(t *testing.T) string { return createTestPNG(t, 10, 10) }, "png"},
		{"jpeg", func(t *testing.T) string { return createTestJPEG(t, 10, 10) }, "jpeg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.create(t)
			got, err := DetectImageFormat(path)
			if err != nil {
				t.Fatalf("DetectImageFormat: %v", err)
			}
			if got != tt.want {
				t.Errorf("DetectImageFormat = %q, want %q", got, tt.want)
			}
		})
	}
}
