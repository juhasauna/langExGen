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

func generateMp3() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:/Users/FIJUSAU/go/src/firebase/serviceAccount.json")

	client, err := texttospeech.NewClient(ctx, opt)
	if err != nil {
		return err
	}

	sentence := getSentences().småferetagIEnLitenStad
	sentence = strings.ReplaceAll(sentence, "\t", "")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	seps := regexp.MustCompile(`[,|.|:|?|!]`)
	splitIt := seps.Split(sentence, -1)
	data := []inputItem{}
	for i, v := range splitIt {
		if v == "" {
			continue
		}
		audio := inputItem{sentence: v, id: fmt.Sprintf("%d_1", i), voice: "sv-SE-Wavenet-A"}
		data = append(data, audio)
	}
	var audioContent []byte
	for _, v := range data {
		count := 3
		rmWhiteSpace := strings.TrimSpace(v.sentence)
		if strings.Index(rmWhiteSpace, "_") == 0 {
			count = strings.Count(v.sentence, "_")
			v.sentence = strings.ReplaceAll(v.sentence, "_", "")
		}

		reqAudio := audioRequest(v)

		respAudio, err := client.SynthesizeSpeech(ctx, &reqAudio)
		if err != nil {
			return err
		}
		quietAudio := makeQuietAudio(len(respAudio.AudioContent))
		for i := 0; i < count; i++ {
			audioContent = append(audioContent, respAudio.AudioContent...)
			audioContent = append(audioContent, quietAudio...)
		}
		fmt.Printf("Wrote audio to []byte: %v\n", v.id)
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
	length = int(float64(length) * 1.6)
	bytes, err := os.ReadFile("C:/Users/FIJUSAU/go/src/langExGen/audio/space.mp3")
	if err != nil {
		log.Fatal(err)
	}
	spaceLen := len(bytes)
	quietLen := length/spaceLen + 1
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
