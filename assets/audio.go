package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const (
	defaultSFXVolume   = 0.5
	defaultMusicVolume = 1.0
)

var (
	ctx = audio.NewContext(44100)

	//go:embed audio/gamemusic.mp3
	gameMusicBytes  []byte
	gameMusicPlayer *audio.Player
)

func init() {
	var err error

	mp3Reader, err := mp3.Decode(ctx, bytes.NewReader(gameMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	introLength := mp3Reader.Length() * 3 / 4
	infiniteReader := audio.NewInfiniteLoopWithIntro(mp3Reader, introLength, mp3Reader.Length()-introLength)
	gameMusicPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	gameMusicPlayer.SetVolume(defaultMusicVolume)
}

func ReplayGameMusic() {
	gameMusicPlayer.Rewind()
	gameMusicPlayer.Play()
}

func StopGameMusic() {
	if gameMusicPlayer.IsPlaying() {
		gameMusicPlayer.Pause()
	}
}

func ResumeGameMusic() {
	if !gameMusicPlayer.IsPlaying() {
		gameMusicPlayer.Play()
	}
}
