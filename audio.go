package tentsuyu

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

const sampleRate = 22050

var (
	audioContext   *audio.Context
	soundFilenames = []string{
		"kaboom-ah2.wav",
		"spaloosh.wav",
	}
	soundPlayers = map[string]*audio.Player{}
)

func init() {
	const sampleRate = 44100
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		panic(err)
	}
}

func (p *AudioPlayer) PlaySE(se string) error {

	sePlayer, _ := audio.NewPlayerFromBytes(p.audioContext, p.seBytes[se])
	sePlayer.SetVolume(p.seVolume[se])
	sePlayer.Play()

	return nil

}

func (p *AudioPlayer) ReturnSongPlayer(name string) *audio.Player {
	return p.songs[name]
}

func NewAudioPlayer() (*AudioPlayer, error) {
	const bytesPerSample = 4

	player := &AudioPlayer{
		audioContext: audioContext,
		//audioPlayer:  p,
		//ambience:     p1,
		//	total:     time.Second * time.Duration(s.Length()) / bytesPerSample / sampleRate,
		volume128: 128,
		seBytes:   make(map[string][]byte),
		seVolume:  make(map[string]float64),
		songs:     make(map[string]*audio.Player),
	}
	/*if player.total == 0 {
		player.total = 1
	}*/

	return player, nil
}

func (p *AudioPlayer) AddSoundEffectFromFile(name, filelocation string, volume float64) error {
	fb, err := ioutil.ReadFile(filelocation)
	if err != nil {
		return err
	}

	s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(fb))
	if err != nil {
		log.Fatal(err)

	}
	b, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatal(err)
	}

	p.seBytes[name] = b
	p.seVolume[name] = volume
	return nil
}

func (p *AudioPlayer) AddSoundEffectFromBytes(name string, fb []byte, volume float64) error {

	s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(fb))
	if err != nil {
		log.Fatal(err)

	}
	b, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatal(err)
	}

	p.seBytes[name] = b
	p.seVolume[name] = volume
	return nil
}

func (p *AudioPlayer) AddSongFromFile(name, filelocation string) error {
	b, err := ioutil.ReadFile(filelocation)
	if err != nil {
		return err
	}
	s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return err
	}
	a, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		return err
	}
	p.songs[name] = a
	return nil
}

func (p *AudioPlayer) AddSongFromBytes(name string, b []byte) error {
	s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return err
	}
	a, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		return err
	}
	p.songs[name] = a
	return nil
}

// Player represents the current audio state.
type AudioPlayer struct {
	audioContext *audio.Context
	current      time.Duration
	total        time.Duration
	seBytes      map[string][]byte
	seVolume     map[string]float64
	volume128    int
	songs        map[string]*audio.Player
}

func (p *AudioPlayer) Update() error {
	/*for _, se := range p.seSlice {
		select {
		case p.seBytes[se] = <-p.seCh[se]:
			close(p.seCh[se])
			p.seCh = nil
		default:
		}
	}*/

	/*if p.ambience.IsPlaying() == false {
		p.ambience.Rewind()
		p.ambience.Play()
	}*/

	return nil
}

func (p *AudioPlayer) UpdateVolumeIfNeeded() {
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		p.volume128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		p.volume128++
	}
	if p.volume128 < 0 {
		p.volume128 = 0
	}
	if 128 < p.volume128 {
		p.volume128 = 128
	}
	//p.audioPlayer.SetVolume(float64(p.volume128) / 128)
}
