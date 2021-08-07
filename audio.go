package tentsuyu

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/h2non/filetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	audioContext *audio.Context
)

// AudioPlayer represents the current audio state.
type AudioPlayer struct {
	audioContext      *audio.Context
	current           time.Duration
	total             time.Duration
	seBytes           map[string][]byte
	seVolume          map[string]float64
	volume128         int
	songs             map[string]*audio.Player
	muteSE, muteMusic bool
}

//NewAudioPlayer returns a new AudioPlayer
func NewAudioPlayer() (*AudioPlayer, error) {
	sampleRate := 44100
	if audioContext == nil {
		audioContext = audio.NewContext(sampleRate)
	}
	//const bytesPerSample = 4

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

//PlaySE playes the sound effect with the given name
func (p *AudioPlayer) PlaySE(se string) error {
	if p.muteSE {
		return nil
	}
	sePlayer := audio.NewPlayerFromBytes(p.audioContext, p.seBytes[se])
	sePlayer.SetVolume(p.seVolume[se])
	sePlayer.Play()

	return nil

}

//MuteAll sets the mute state of both SoundEffects and Music
func (p *AudioPlayer) MuteAll(m bool) {
	p.muteSE = m
	p.muteMusic = m
}

//MuteSE sets the mute state of SoundEffects
func (p *AudioPlayer) MuteSE(m bool) {
	p.muteSE = m
}

//MuteMusic sets the mute state of Music
func (p *AudioPlayer) MuteMusic(m bool) {
	p.muteMusic = m
	for _, player := range p.songs {
		player.Pause()
	}
}

func (p *AudioPlayer) PauseAllMusic() {
	for _, player := range p.songs {
		player.Pause()
	}
}

//IsSEMuted returns true if the sound effects are muted for the AudioPlayer
func (p *AudioPlayer) IsSEMuted() bool {
	return p.muteSE
}

//IsMusicMuted returns true if the music is muted for the AudioPlayer
func (p *AudioPlayer) IsMusicMuted() bool {
	return p.muteMusic
}

//ReturnSongPlayer returns the player for the song audio
func (p *AudioPlayer) ReturnSongPlayer(name string) *audio.Player {
	return p.songs[name]
}

//AddSoundEffectFromFile adds a SE of the given name and volume at the file location.
func (p *AudioPlayer) AddSoundEffectFromFile(name, filelocation string, volume float64) error {
	fb, err := ioutil.ReadFile(filelocation)
	if err != nil {
		return err
	}

	return p.AddSoundEffectFromBytes(name, fb, volume)
}

//AddSoundEffectFromBytes adds a new sound effect file from a byte slice
func (p *AudioPlayer) AddSoundEffectFromBytes(name string, fb []byte, volume float64) error {
	var s io.Reader
	var err error
	if filetype.IsExtension(fb, "wav") {
		s, err = wav.Decode(audioContext, bytes.NewReader(fb))
		if err != nil {
			log.Fatal(err)

		}
	} else if filetype.IsExtension(fb, "mp3") {
		s, err = mp3.Decode(audioContext, bytes.NewReader(fb))
		if err != nil {
			log.Fatal(err)

		}
	} else if filetype.IsExtension(fb, "ogg") {
		s, err = vorbis.Decode(audioContext, bytes.NewReader(fb))
		if err != nil {
			log.Fatal(err)

		}
	}
	b, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatal(err)
	}

	p.seBytes[name] = b
	p.seVolume[name] = volume
	return nil
}

//AddSongFromFile to the AudioPlayer
func (p *AudioPlayer) AddSongFromFile(name, filelocation string) error {

	fb, err := ioutil.ReadFile(filelocation)
	if err != nil {
		return err
	}
	return p.AddSongFromBytes(name, fb)
}

//AddSongFromBytes takes the byte slice of the song file
func (p *AudioPlayer) AddSongFromBytes(name string, fb []byte) error {
	var s io.Reader
	var err error
	if filetype.IsExtension(fb, "wav") {
		s, err = wav.Decode(audioContext, bytes.NewReader(fb))
		if err != nil {
			log.Fatal(err)

		}
	} else if filetype.IsExtension(fb, "mp3") {
		s, err = mp3.Decode(audioContext, bytes.NewReader(fb))
		if err != nil {
			log.Fatal(err)

		}
	} else if filetype.IsExtension(fb, "ogg") {
		s, err = vorbis.Decode(audioContext, bytes.NewReader(fb))
		if err != nil {
			log.Fatal(err)

		}
	}
	a, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		return err
	}
	p.songs[name] = a
	return nil
}

//Update isn't currently used for the AudioPlayer
//TODO: Implement this ... if necessary
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

//UpdateVolumeIfNeeded should be used to listen to changing the volume level
//TODO: Implement this
func (p *AudioPlayer) UpdateVolumeIfNeeded() {
	if ebiten.IsKeyPressed(ebiten.KeyMinus) {
		p.volume128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyEqual) {
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

//JukeBox for playing songs

//JukeBox is used to control BG Music at least.
//Possible SE support
type JukeBox struct {
	AudioPlayer  *AudioPlayer
	SongList     []string
	songRepeat   int
	CurrSong     string
	NextSong     string
	PrevSong     string
	BGDefaultVol float64
	CurrVol      float64
	currIndx     int
}

//CreateJukeBox from passed audioplayer, music and sound effect lists
func CreateJukeBox(audioPlayer *AudioPlayer, musicList []string) *JukeBox {
	j := &JukeBox{
		AudioPlayer:  audioPlayer,
		SongList:     musicList,
		BGDefaultVol: 0.3,
		CurrVol:      0.3,
		currIndx:     -1,
	}
	return j
}

//ContinuousPlay of the song based on passed song name
func (j *JukeBox) ContinuousPlay(songName string) {
	/*for i, s := range j.SongList {
		if s == songName {
			j.currIndx = i
		}
		break
	}*/
	if !j.AudioPlayer.ReturnSongPlayer(songName).IsPlaying() && !j.AudioPlayer.IsMusicMuted() {
		j.AudioPlayer.ReturnSongPlayer(songName).Rewind()
		j.AudioPlayer.ReturnSongPlayer(songName).Play()
		j.AudioPlayer.ReturnSongPlayer(songName).SetVolume(j.CurrVol)
	}
}

//PlaySong with the given name at the passed volume
func (j *JukeBox) PlaySong(songName string, vol float64) {
	j.AudioPlayer.ReturnSongPlayer(songName).Rewind()
	j.AudioPlayer.ReturnSongPlayer(songName).Play()
	j.AudioPlayer.ReturnSongPlayer(songName).SetVolume(vol)
}

//PlayBG plays background music in the given order
func (j *JukeBox) PlayBG() {
	if j.currIndx == -1 {
		j.currIndx = 0
		j.CurrSong = j.SongList[j.currIndx]
		j.PlaySong(j.CurrSong, j.CurrVol)
	}
	if !j.AudioPlayer.ReturnSongPlayer(j.CurrSong).IsPlaying() {
		j.currIndx++
		if j.currIndx < len(j.SongList) {
			j.CurrSong = j.SongList[j.currIndx]
			j.PlaySong(j.CurrSong, j.CurrVol)
		} else {
			j.currIndx = -1
		}
	}
}

//CurrentSongName returns the name of the song that is currently being played
func (j *JukeBox) CurrentSongName() string {
	if j.currIndx != -1 {
		return j.SongList[j.currIndx]
	}
	return "No current song"
}
