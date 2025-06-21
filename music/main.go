package main

import (
	"encoding/binary"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"io"
	"log"
	"os/exec"
	"time"
)

// YouTubeStreamer implements beep.Streamer for raw PCM audio
type YouTubeStreamer struct {
	cmd     *exec.Cmd
	reader  io.Reader
	err     error
	quality string // audio quality preference
}

// Stream reads PCM data and converts to audio samples
func (s *YouTubeStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	buf := make([]byte, len(samples)*4) // 16-bit stereo = 4 bytes/sample
	bytesRead, err := io.ReadFull(s.reader, buf)

	if err == io.EOF || err == io.ErrUnexpectedEOF {
		s.err = err
		return bytesRead / 4, false
	}
	if err != nil {
		s.err = err
		return 0, false
	}

	// Convert bytes to audio samples
	for i := 0; i < bytesRead/4; i++ {
		left := int16(binary.LittleEndian.Uint16(buf[i*4 : i*4+2]))
		right := int16(binary.LittleEndian.Uint16(buf[i*4+2 : i*4+4]))
		samples[i] = [2]float64{
			float64(left) / 32768.0,
			float64(right) / 32768.0,
		}
	}

	return bytesRead / 4, true
}

// Err returns any playback error
func (s *YouTubeStreamer) Err() error {
	return s.err
}

// Close stops the streaming process
func (s *YouTubeStreamer) Close() {
	if s.cmd != nil && s.cmd.Process != nil {
		s.cmd.Process.Kill()
	}
}

// NewYouTubeStreamer creates a new YouTube audio stream
func NewYouTubeStreamer(url, quality string) (*YouTubeStreamer, error) {
	// Get audio URL using yt-dlp
	ytCmd := exec.Command("yt-dlp",
		"-f", "bestaudio[ext=webm]",
		"-g",
		url,
	)

	audioURL, err := ytCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get audio URL: %w", err)
	}

	// Start ffmpeg to decode audio stream
	ffmpegCmd := exec.Command("ffmpeg",
		"-i", string(audioURL[:len(audioURL)-1]), // Remove trailing newline
		"-f", "s16le",          // 16-bit little-endian PCM
		"-acodec", "pcm_s16le", // Raw PCM codec
		"-ar", "44100",         // Sample rate
		"-ac", "2",             // Stereo audio
		"-loglevel", "quiet",   // Suppress ffmpeg logs
		"pipe:1",               // Output to stdout
	)

	stdout, err := ffmpegCmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := ffmpegCmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	return &YouTubeStreamer{
		cmd:     ffmpegCmd,
		reader:  stdout,
		quality: quality,
	}, nil
}

func main() {
	// Replace with your YouTube URL
	youtubeURL := "https://www.youtube.com/watch?v=Am4wYTiHHx8"

	// Create streamer
	streamer, err := NewYouTubeStreamer(youtubeURL, "best")
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize speaker
	sampleRate := beep.SampleRate(44100)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))

	// Play stream
	speaker.Play(streamer)

	// Wait for playback to finish
	for streamer.Err() == nil {
		time.Sleep(500 * time.Millisecond)
	}

	if streamer.Err() != io.EOF {
		log.Println("Playback error:", streamer.Err())
	}
}

/*
package main

import (
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

func main() {
	
	f, err := os.Open("heartbeat.mp3") // Replace with your path
	if err != nil {
		panic(err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	// speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	select {} // keep playing
}
*/