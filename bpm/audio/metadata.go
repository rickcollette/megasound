package audio

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rickcollette/megasound/flac"
	"github.com/rickcollette/megasound/mp3"
	"github.com/rickcollette/megasound/vorbis"
	"github.com/rickcollette/megasound/wav"
)

// AudioMetadata contains audio properties extracted from a file.
type AudioMetadata struct {
	Rate     int // Sample rate in Hz
	Channels int // Number of audio channels
}

// GetMetadata extracts metadata (sample rate, channels) based on file type using the megasound package.
func GetMetadata(filePath string) (*AudioMetadata, error) {
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".wav":
		return getWavMetadata(filePath)
	case ".mp3":
		return getMp3Metadata(filePath)
	case ".flac":
		return getFlacMetadata(filePath)
	case ".ogg":
		return getOggMetadata(filePath)
	default:
		return nil, errors.New("unsupported file format")
	}
}

// getWavMetadata extracts metadata for WAV files using megasound/wav.
func getWavMetadata(filePath string) (*AudioMetadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open WAV file: %v", err)
	}
	defer file.Close()

	_, format, err := wav.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode WAV file: %v", err)
	}

	return &AudioMetadata{
		Rate:     int(format.SampleRate), // Explicit conversion to int
		Channels: format.NumChannels,
	}, nil
}

// getMp3Metadata extracts metadata for MP3 files using megasound/mp3.
func getMp3Metadata(filePath string) (*AudioMetadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open MP3 file: %v", err)
	}
	defer file.Close()

	_, format, err := mp3.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MP3 file: %v", err)
	}

	return &AudioMetadata{
		Rate:     int(format.SampleRate), // Explicit conversion to int
		Channels: format.NumChannels,
	}, nil
}

// getFlacMetadata extracts metadata for FLAC files using megasound/flac.
func getFlacMetadata(filePath string) (*AudioMetadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open FLAC file: %v", err)
	}
	defer file.Close()

	_, format, err := flac.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode FLAC file: %v", err)
	}

	return &AudioMetadata{
		Rate:     int(format.SampleRate), // Explicit conversion to int
		Channels: format.NumChannels,
	}, nil
}

// getOggMetadata extracts metadata for OGG-Vorbis files using megasound/vorbis.
func getOggMetadata(filePath string) (*AudioMetadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open OGG file: %v", err)
	}
	defer file.Close()

	_, format, err := vorbis.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OGG file: %v", err)
	}

	return &AudioMetadata{
		Rate:     int(format.SampleRate), // Explicit conversion to int
		Channels: format.NumChannels,
	}, nil
}

// CalculateInterval dynamically calculates INTERVAL based on the sample rate.
func CalculateInterval(rate int, targetFrameSize int) int {
	// INTERVAL is computed based on the target frame size relative to the sample rate.
	return targetFrameSize
}
