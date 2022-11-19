package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type inputItem struct {
	id           string
	sentence     string
	voice        string
	volume       float64
	speakingRate float64
	audioContent []byte
}

type audioAndPause struct {
	audio   inputItem
	pause   inputItem
	repeats int
}

func generateMp3() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:/Users/FIJUSAU/go/src/firebase/serviceAccount.json")

	client, err := texttospeech.NewClient(ctx, opt)
	if err != nil {
		return err
	}

	sentence := getSentences().test
	sentence = strings.ReplaceAll(sentence, "\t", "")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	seps := regexp.MustCompile(`[,|.|:|?|!]`)
	splitIt := seps.Split(sentence, -1)
	data := []audioAndPause{}
	for i, v := range splitIt {
		if v == "" {
			continue
		}
		audio := inputItem{sentence: v, id: fmt.Sprintf("%d_1", i), voice: "sv-SE-Wavenet-A"}
		data = append(data, audioAndPause{audio: audio})
	}
	var audioContent []byte
	for _, v := range data {
		count := 3
		rmWhiteSpace := strings.TrimSpace(v.audio.sentence)
		if strings.Index(rmWhiteSpace, "_") == 0 {
			count = strings.Count(v.audio.sentence, "_")
			v.audio.sentence = strings.ReplaceAll(v.audio.sentence, "_", "")
		}

		reqAudio := audioRequest(v.audio)

		respAudio, err := client.SynthesizeSpeech(ctx, &reqAudio)
		if err != nil {
			return err
		}
		fmt.Println(len(respAudio.AudioContent))
		quietAudio := makeQuietAudio(len(respAudio.AudioContent) * 3 / 2)
		for i := 0; i < count; i++ {
			audioContent = append(audioContent, respAudio.AudioContent...)
			audioContent = append(audioContent, quietAudio...)
		}
		fmt.Printf("Wrote audio to []byte: %v\n", v.audio.id)
	}
	path := "C:/Users/FIJUSAU/go/src/langExGen/audio/"
	filePath := path + "test.mp3"
	err = os.WriteFile(filePath, audioContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func makeQuietAudio(length int) []byte {
	bytes, err := os.ReadFile("C:/Users/FIJUSAU/go/src/langExGen/audio/space.mp3")
	if err != nil {
		log.Fatal(err)
	}
	spaceLen := len(bytes)
	quietLen := length/spaceLen + 1
	fmt.Println(spaceLen, "\t", quietLen)
	var quietBytes []byte
	for i := 0; i < quietLen; i++ {
		quietBytes = append(quietBytes, bytes...)
	}
	return quietBytes
}

func audioRequest(v inputItem) texttospeechpb.SynthesizeSpeechRequest {
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: v.sentence},
		},

		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "sv-SE",
			Name:         v.voice,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			VolumeGainDb:  v.volume,
			SpeakingRate:  v.speakingRate,
		},
	}
	return req
}

func generateQuiet() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:/Users/FIJUSAU/go/src/firebase/serviceAccount.json")

	client, err := texttospeech.NewClient(ctx, opt)
	if err != nil {
		return err
	}

	type volRate struct {
		volume       float64
		speakingRate float64
	}

	data := []audioAndPause{}

	audio := inputItem{sentence: " ", voice: "sv-SE-Wavenet-A"}
	data = append(data, audioAndPause{audio: audio})

	var audioContent []byte
	for _, v := range data {

		reqAudio := audioRequest(v.audio)

		respAudio, err := client.SynthesizeSpeech(ctx, &reqAudio)
		if err != nil {
			return err
		}

		audioContent = append(audioContent, respAudio.AudioContent...)
		fmt.Printf("Wrote audio to []byte: %v\n", v.audio.id)
	}
	path := "C:/Users/FIJUSAU/go/src/langExGen/audio/"
	filePath := path + "test.mp3"
	err = os.WriteFile(filePath, audioContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GenerateCustomMp3_old() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:/Users/FIJUSAU/go/src/firebase/serviceAccount.json")

	client, err := texttospeech.NewClient(ctx, opt)
	if err != nil {
		return err
	}

	type volRate struct {
		volume       float64
		speakingRate float64
	}
	// normal := volRate{0, 1}
	pause := volRate{-96, .6}

	sentence := getSentences().test
	sentence = strings.ReplaceAll(sentence, "\t", "")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	seps := regexp.MustCompile(`[,|.|:|?|!]`)
	splitIt := seps.Split(sentence, -1)
	data := []audioAndPause{}
	for i, v := range splitIt {
		if v == "" {
			continue
		}
		audio := inputItem{sentence: v, id: fmt.Sprintf("%d_1", i), voice: "sv-SE-Wavenet-A"}
		pause := inputItem{sentence: v, id: fmt.Sprintf("%d_1pause", i), voice: "sv-SE-Wavenet-A", volume: pause.volume, speakingRate: pause.speakingRate}
		data = append(data, audioAndPause{audio: audio, pause: pause})
	}
	var audioContent []byte
	for _, v := range data {
		count := 3
		rmWhiteSpace := strings.TrimSpace(v.audio.sentence)
		if strings.Index(rmWhiteSpace, "_") == 0 {
			count = strings.Count(v.audio.sentence, "_")
			v.audio.sentence = strings.ReplaceAll(v.audio.sentence, "_", "")
			v.pause.sentence = strings.ReplaceAll(v.pause.sentence, "_", "")
		}

		reqAudio := audioRequest(v.audio)
		reqPause := audioRequest(v.pause)

		respAudio, err := client.SynthesizeSpeech(ctx, &reqAudio)
		if err != nil {
			return err
		}
		respPause, err := client.SynthesizeSpeech(ctx, &reqPause)
		if err != nil {
			return err
		}
		fmt.Println(len(respAudio.AudioContent))
		fmt.Println(len(respPause.AudioContent))
		for i := 0; i < count; i++ {
			audioContent = append(audioContent, respAudio.AudioContent...)
			audioContent = append(audioContent, respPause.AudioContent...)
		}
		fmt.Printf("Wrote audio to []byte: %v\n", v.audio.id)
	}
	path := "C:/Users/FIJUSAU/go/src/langExGen/audio/"
	filePath := path + "test.mp3"
	err = os.WriteFile(filePath, audioContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
