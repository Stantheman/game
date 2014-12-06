package sound

import (
	"errors"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type SoundManager struct{}

func Init() (*SoundManager, error) {
	if ret := sdl.Init(sdl.INIT_AUDIO); ret != 0 {
		return nil, errors.New("Couldn't initialize audio")
	}
	return &SoundManager{}, nil
}

func (s *SoundManager) Beep() error {
	audioSpec := new(sdl.AudioSpec)

	audioSpec.Freq = 28000
	audioSpec.Format = sdl.AUDIO_S16SYS
	audioSpec.Channels = 1
	audioSpec.Samples = 2048
	audioSpec.UserData = unsafe.Pointer(s)

	f := callback
	fp := unsafe.Pointer(&f)

	audioSpec.Callback = sdl.AudioCallback(fp)

	resultSpec := new(sdl.AudioSpec)
	if ret := sdl.OpenAudioDevice("", 0, audioSpec, resultSpec, sdl.AUDIO_ALLOW_ANY_CHANGE); ret != 0 {
		return fmt.Errorf("Couldn't initialize audio device: returned %v", ret)
	}

	sdl.PauseAudio(0)
	return nil
}

func callback(ptr unsafe.Pointer, stream *uint, length int) {
	fmt.Printf("stream is %v, len is %v\n", stream, length)
}

// func callback(ptr unsafe.Pointer, stream *uint, length int) {
// 	fmt.Printf("stream is %v, len is %v\n", stream, length)
// }

// const int AMPLITUDE = 28000;
// const int FREQUENCY = 44100;

// struct BeepObject
// {
//     double freq;
//     int samplesLeft;
// };

// class Beeper
// {
// private:
//     double v;
//     std::queue<BeepObject> beeps;
// public:
//     Beeper();
//     ~Beeper();
//     void beep(double freq, int duration);
//     void generateSamples(Sint16 *stream, int length);
//     void wait();
// };

// void audio_callback(void*, Uint8*, int);

// Beeper::Beeper()
// {
//     SDL_AudioSpec desiredSpec;

//     desiredSpec.freq = FREQUENCY;
//     desiredSpec.format = AUDIO_S16SYS;
//     desiredSpec.channels = 1;
//     desiredSpec.samples = 2048;
//     desiredSpec.callback = audio_callback;
//     desiredSpec.userdata = this;

//     SDL_AudioSpec obtainedSpec;

//     // you might want to look for errors here
//     SDL_OpenAudio(&desiredSpec, &obtainedSpec);

//     // start play audio
//     SDL_PauseAudio(0);
// }

// Beeper::~Beeper()
// {
//     SDL_CloseAudio();
// }

// void Beeper::generateSamples(Sint16 *stream, int length)
// {
//     int i = 0;
//     while (i < length) {

//         if (beeps.empty()) {
//             while (i < length) {
//                 stream[i] = 0;
//                 i++;
//             }
//             return;
//         }
//         BeepObject& bo = beeps.front();

//         int samplesToDo = std::min(i + bo.samplesLeft, length);
//         bo.samplesLeft -= samplesToDo - i;

//         while (i < samplesToDo) {
//             stream[i] = AMPLITUDE * std::sin(v * 2 * M_PI / FREQUENCY);
//             i++;
//             v += bo.freq;
//         }

//         if (bo.samplesLeft == 0) {
//             beeps.pop();
//         }
//     }
// }

// void Beeper::beep(double freq, int duration)
// {
//     BeepObject bo;
//     bo.freq = freq;
//     bo.samplesLeft = duration * FREQUENCY / 1000;

//     SDL_LockAudio();
//     beeps.push(bo);
//     SDL_UnlockAudio();
// }

// void Beeper::wait()
// {
//     int size;
//     do {
//         SDL_Delay(20);
//         SDL_LockAudio();
//         size = beeps.size();
//         SDL_UnlockAudio();
//     } while (size > 0);

// }

// void audio_callback(void *_beeper, Uint8 *_stream, int _length)
// {
//     Sint16 *stream = (Sint16*) _stream;
//     int length = _length / 2;
//     Beeper* beeper = (Beeper*) _beeper;

//     beeper->generateSamples(stream, length);
// }

// int main(int argc, char* argv[])
// {
//     SDL_Init(SDL_INIT_AUDIO);

//     int duration = 1000;
//     double Hz = 440;

//     Beeper b;
//     b.beep(Hz, duration);
//     b.wait();

//     return 0;
// }
