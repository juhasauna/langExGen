package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

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

type voices struct {
	se []string
	en string
	fi string
}

func newVoices() voices {
	return voices{se: []string{"sv-SE-Wavenet-B", "sv-SE-Wavenet-C", "sv-SE-Wavenet-D", "sv-SE-Wavenet-E", "sv-SE-Wavenet-A"},
		en: "en-US-Wavenet-G",
		fi: "fi-FI-Wavenet-A",
	}
}

func generateMp3() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:/Users/FIJUSAU/go/src/firebase/serviceAccount.json")

	client, err := texttospeech.NewClient(ctx, opt)
	if err != nil {
		return err
	}
	sentence := bliExamenstuderande

	audioInputs := []inputItem{}
	for i := 0; i < len(sentence); i = i + 5000 {
		to := 0
		if len(sentence) > i+5000 {
			to = i + 5000
		} else {
			to = len(sentence)
		}
		audioInputs = append(audioInputs, inputItem{sentence: sentence[i:to], voice: "en-US-Neural2-J"})
		// audioInputs = append(audioInputs, inputItem{sentence: sentence[i:to], voice: "sv-SE-Wavenet-B"})
	}

	completeAudio := []byte{}
	for i, v := range audioInputs {
		reqAudio := audioRequest(v, "en-US")
		respAudio, err := client.SynthesizeSpeech(ctx, &reqAudio)
		if err != nil {
			return err
		}

		completeAudio = append(completeAudio, respAudio.AudioContent...)
		sentLen := len(v.sentence)
		printLen := 1
		if sentLen > 20 {
			printLen = 20
		}
		fmt.Println("successfully added part", i, v.sentence[len(v.sentence)-printLen:len(v.sentence)-1])
	}

	path := "C:/Users/FIJUSAU/go/src/langExGen/audio/"
	filePath := path + "test.mp3"
	err = os.WriteFile(filePath, completeAudio, 0644)
	if err != nil {
		return err
	}

	return nil
}
func generateLanguageLearningMp3() error {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Int()
	ctx := context.Background()
	opt := option.WithCredentialsFile("C:/Users/FIJUSAU/go/src/firebase/serviceAccount.json")

	client, err := texttospeech.NewClient(ctx, opt)
	if err != nil {
		return err
	}
	voices := newVoices()
	sentence := synkronmaskiner2
	sentence = strings.ReplaceAll(sentence, "\t", "")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	seps := regexp.MustCompile(`[,|.|:|?|!]`)
	splitIt := seps.Split(sentence, -1)
	data := []inputItem{}
	seVoicesLen := len(voices.se)
	voiceIndex := randInt % seVoicesLen
	voice := voices.se[voiceIndex]
	for i, v := range splitIt {
		if v == "" {
			continue
		}
		if strings.Contains(v, "_lang_en_") {
			voice = voices.en
			v = strings.ReplaceAll(v, "_lang_en_", "")
		} else if strings.Contains(v, "_lang_se_") {
			voice = voices.se[voiceIndex]
			v = strings.ReplaceAll(v, "_lang_se_", "")
		} else if strings.Contains(v, "_next_voice_") {
			voiceIndex++
			voiceIndex = voiceIndex % seVoicesLen
			voice = voices.se[voiceIndex]
			v = strings.ReplaceAll(v, "_next_voice_", "")
		}
		audio := inputItem{sentence: v, id: fmt.Sprintf("%d_1", i), voice: voice}
		data = append(data, audio)
	}
	var audioContent []byte
	for _, v := range data {

		count := 3
		rmWhiteSpace := strings.TrimSpace(v.sentence)
		skip := false
		if strings.Index(rmWhiteSpace, "$") == 0 {
			v.sentence = strings.ReplaceAll(v.sentence, "$", "")
			skip = true
		} else if strings.Index(rmWhiteSpace, "_") == 0 {
			count = strings.Count(v.sentence, "_")
			v.sentence = strings.ReplaceAll(v.sentence, "_", "")
		}

		reqAudio := audioRequest(v, "sv-SE")

		respAudio, err := client.SynthesizeSpeech(ctx, &reqAudio)
		if err != nil {
			return err
		}
		quietAudio := makeQuietAudio(len(respAudio.AudioContent))
		for i := 0; i < count; i++ {
			audioContent = append(audioContent, respAudio.AudioContent...)
			if skip {
				break
			}
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

func audioRequest(v inputItem, langCode string) texttospeechpb.SynthesizeSpeechRequest {
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: v.sentence},
		},

		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: langCode,
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
