package media

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
	_ "golang.org/x/image/webp" // register WebP decoder
)

// Image sanitization constants.
const (
	// ImageMaxSide is the maximum pixels per side before resize.
	ImageMaxSide = 1200
	// ImageSanitizeMaxBytes is the max file size after compression (5MB, Anthropic API limit).
	ImageSanitizeMaxBytes = 5 * 1024 * 1024
)

// jpegQualities is the grid of quality levels to try during sanitization.
var jpegQualities = []int{85, 75, 65, 55, 45, 35}

func init() {
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "\x89PNG", png.Decode, png.DecodeConfig)
}

// DetectImageFormat reads the first 12 bytes of a file to identify the image format
// via magic bytes. Returns "png", "jpeg", "webp", "gif", or "jpeg" as default fallback.
func DetectImageFormat(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open for format detection: %w", err)
	}
	defer f.Close()

	buf := make([]byte, 12)
	if _, err := f.Read(buf); err != nil {
		return "", fmt.Errorf("read header bytes: %w", err)
	}

	switch {
	case bytes.HasPrefix(buf, []byte("\x89PNG")):
		return "png", nil
	case bytes.HasPrefix(buf, []byte("\xff\xd8\xff")):
		return "jpeg", nil
	case bytes.HasPrefix(buf, []byte("RIFF")) && len(buf) >= 12 && bytes.Equal(buf[8:12], []byte("WEBP")):
		return "webp", nil
	case bytes.HasPrefix(buf, []byte("GIF8")):
		return "gif", nil
	default:
		return "jpeg", nil
	}
}

// PreserveFormat re-encodes an image to preserve its original format (PNG alpha, etc).
// Does NOT resize or compress — only normalizes format for storage.
// PNG → PNG (preserves alpha), JPEG → JPEG (quality 95), WebP/GIF → JPEG (fallback).
// Returns: output path, output MIME type, error.
func PreserveFormat(inputPath string) (string, string, error) {
	srcFormat, err := DetectImageFormat(inputPath)
	if err != nil {
		srcFormat = "jpeg"
	}

	img, err := imaging.Open(inputPath, imaging.AutoOrientation(true))
	if err != nil {
		return "", "", fmt.Errorf("open image: %w", err)
	}

	var buf bytes.Buffer
	var ext, mime string

	switch srcFormat {
	case "png":
		if err := png.Encode(&buf, img); err != nil {
			return "", "", fmt.Errorf("encode png: %w", err)
		}
		ext, mime = ".png", "image/png"
	default:
		if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 95}); err != nil {
			return "", "", fmt.Errorf("encode jpeg: %w", err)
		}
		ext, mime = ".jpg", "image/jpeg"
	}

	return writeTempImage(buf.Bytes(), ext, mime)
}

// SanitizeForVision resizes and compresses an image for LLM API limits.
// Pipeline: decode → auto-orient EXIF → resize if >1200px → compress until <5MB.
// PNG inputs preserve alpha channel when output <5MB, otherwise fallback to JPEG.
// Returns: compressed image data, MIME type, error.
func SanitizeForVision(inputPath string, inputMime string) ([]byte, string, error) {
	img, err := imaging.Open(inputPath, imaging.AutoOrientation(true))
	if err != nil {
		return nil, "", fmt.Errorf("open image: %w", err)
	}

	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	if w > ImageMaxSide || h > ImageMaxSide {
		img = imaging.Fit(img, ImageMaxSide, ImageMaxSide, imaging.Lanczos)
	}

	isPNG := inputMime == "image/png"

	if isPNG {
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err == nil && buf.Len() <= ImageSanitizeMaxBytes {
			return buf.Bytes(), "image/png", nil
		}
	}

	for _, quality := range jpegQualities {
		var buf bytes.Buffer
		if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality}); err != nil {
			return nil, "", fmt.Errorf("encode jpeg (q=%d): %w", quality, err)
		}
		if buf.Len() <= ImageSanitizeMaxBytes {
			return buf.Bytes(), "image/jpeg", nil
		}
	}

	return nil, "", fmt.Errorf("image too large even at lowest quality (dimensions: %dx%d)", w, h)
}

// writeTempImage writes data to a temp file with the given extension and returns
// the file path and MIME type. Cleans up on write failure.
func writeTempImage(data []byte, ext, mime string) (string, string, error) {
	f, err := os.CreateTemp("", "goclaw_preserved_*"+ext)
	if err != nil {
		return "", "", fmt.Errorf("create temp file: %w", err)
	}
	outPath := f.Name()
	if _, err := f.Write(data); err != nil {
		f.Close()
		os.Remove(outPath)
		return "", "", fmt.Errorf("write image: %w", err)
	}
	f.Close()
	return outPath, mime, nil
}
