package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func getInfo(filename string) (*ffprobe.ProbeData, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, filename)
	if err != nil {
		return nil, fmt.Errorf("error getting ffprobe data: %w", err)
	}
	return data, nil
}

func printInfo(info *ffprobe.ProbeData) {
	fmt.Println("BitRate:", info.Format.BitRate)
	fmt.Println("Duration:", info.Format.Duration())
	fmt.Println("Format:", info.Format.FormatName, info.Format.FormatLongName)

	fmt.Println("Size:", info.Format.Size)

	fmt.Println("NBStreams", info.Format.NBStreams)

	printVideoStreamsInfo(info.StreamType(ffprobe.StreamVideo))
	printAudioStreamsInfo(info.StreamType(ffprobe.StreamAudio))
	printSubtitleStreamsInfo(info.StreamType(ffprobe.StreamSubtitle))
}

func printVideoStreamsInfo(streams []ffprobe.Stream) {
	for idx, s := range streams {
		fmt.Printf("-> Video Stream #%d\n", idx)
		fmt.Println("---> AvgFrameRate", s.AvgFrameRate)
		fmt.Println("---> BitRate", s.BitRate)
		fmt.Println("---> SampleRate", s.SampleRate)
		fmt.Println("---> CodecName CodecLongName", s.CodecName, s.CodecLongName)
		fmt.Println("---> Width Height", s.Width, s.Height)
		fmt.Println("---> BitsPerSample", s.BitsPerSample)
	}
}

func printAudioStreamsInfo(streams []ffprobe.Stream) {
	for idx, s := range streams {
		fmt.Printf("-> Audio Stream #%d\n", idx)
		fmt.Println("---> AvgFrameRate", s.AvgFrameRate)
		fmt.Println("---> BitRate", s.BitRate)
		fmt.Println("---> SampleRate", s.SampleRate)
		fmt.Println("---> CodecName CodecLongName", s.CodecName, s.CodecLongName)
		fmt.Println("---> BitsPerSample", s.BitsPerSample)
	}
}
func printSubtitleStreamsInfo(streams []ffprobe.Stream) {
	for idx, s := range streams {
		fmt.Printf("-> Subtitle Stream #%d\n", idx)
		fmt.Println("---> AvgFrameRate", s)
		fmt.Println("---> BitRate", s.BitRate)
		fmt.Println("---> SampleRate", s.SampleRate)
		fmt.Println("---> CodecName CodecLongName", s.CodecName, s.CodecLongName)
		fmt.Println("---> BitsPerSample", s.BitsPerSample)
	}
}

func main() {

	flag.Parse()
	for _, filename := range flag.Args() {
		fmt.Printf("Probing “%s”...\n", filename)

		probeData, err := getInfo(filename)
		if err != nil {
			panic(err)
		}

		numAttachments := len(probeData.StreamType(ffprobe.StreamSubtitle))
		if numAttachments > 0 {
			fmt.Printf("Found %d attachments!\n", numAttachments)
		}

		printInfo(probeData)
	}
}
