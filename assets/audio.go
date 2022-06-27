package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume   = 1.0
	defaultMusicVolume = 1.0
)

var (
	ctx = audio.NewContext(44100)

	//go:embed audio/gamemusic.wav
	gameMusicBytes  []byte
	gameMusicPlayer *audio.Player
	//go:embed audio/portal.wav
	portalSoundBytes  []byte
	portalSoundPlayer *audio.Player
	//go:embed audio/death.wav
	deathSoundBytes  []byte
	deathSoundPlayer *audio.Player
	//go:embed audio/dash.wav
	dashSoundBytes  []byte
	dashSoundPlayer *audio.Player
)

func init() {
	wavReader, err := wav.Decode(ctx, bytes.NewReader(gameMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	introLength := wavReader.Length() * 3 / 4
	infiniteReader := audio.NewInfiniteLoopWithIntro(wavReader, introLength, wavReader.Length()-introLength)
	gameMusicPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	gameMusicPlayer.SetVolume(defaultMusicVolume)

	wavReader, err = wav.Decode(ctx, bytes.NewReader(portalSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	portalSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}
	portalSoundPlayer.SetVolume(defaultSFXVolume)

	wavReader, err = wav.Decode(ctx, bytes.NewReader(deathSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	deathSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}
	deathSoundPlayer.SetVolume(defaultSFXVolume)

	wavReader, err = wav.Decode(ctx, bytes.NewReader(dashSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	dashSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}
	dashSoundPlayer.SetVolume(defaultSFXVolume)
}

// Musics

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

// Sfx

func PlayPortalSound() {
	portalSoundPlayer.Rewind()
	portalSoundPlayer.Play()
}

func PlayDeathSound() {
	deathSoundPlayer.Rewind()
	deathSoundPlayer.Play()
}

func PlayDashSound() {
	dashSoundPlayer.Rewind()
	dashSoundPlayer.Play()
}

func StopDashSound() {
	if dashSoundPlayer.IsPlaying() {
		dashSoundPlayer.Pause()
	}
}
