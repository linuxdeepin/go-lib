package sound_effect

import (
	"errors"
	"os"
	"time"

	"github.com/cryptix/wav"
	"pkg.deepin.io/lib/asound"
	paSimple "pkg.deepin.io/lib/pulse/simple"
)

type WavDecoder struct {
	reader        *wav.Reader
	f             *os.File
	sampleSpec    *SampleSpec
	bufSize       int
	bytesPerFrame int
	duration      time.Duration
}

func newWavDecoder(filename string, fileInfo os.FileInfo) (Decoder, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	wavReader, err := wav.NewReader(f, fileInfo.Size())
	if err != nil {
		return nil, err
	}
	wavFile := wavReader.GetFile()
	paFormat, pcmFormat, err := getWavFormat(wavFile)
	if err != nil {
		return nil, err
	}

	sampleSpec := &SampleSpec{
		paFormat:  paFormat,
		pcmFormat: pcmFormat,
		rate:      int(wavFile.SampleRate),
		channels:  int(wavFile.Channels),
	}

	bytesPerFrame := int(wavFile.Channels) * int(wavFile.SignificantBits/8)
	bufSize := int(wavFile.SampleRate/8) * bytesPerFrame
	// NOTE: do not use wavFile.Duration
	duration := time.Duration(
		float64(wavReader.GetSampleCount()/uint32(wavFile.Channels)) / float64(wavFile.SampleRate) *
			float64(time.Second),
	)

	return &WavDecoder{
		f:             f,
		reader:        wavReader,
		sampleSpec:    sampleSpec,
		bufSize:       bufSize,
		bytesPerFrame: bytesPerFrame,
		duration:      duration,
	}, nil
}

func getWavFormat(wavFile wav.File) (paSimple.SampleFormat, asound.PCMFormat, error) {
	switch wavFile.SignificantBits {
	case 8:
		return paSimple.SampleFormatU8, asound.PCMFormatU8, nil
	case 16:
		return paSimple.SampleFormatS16LE, asound.PCMFormatS16LE, nil
	case 32:
		return paSimple.SampleFormatS32LE, asound.PCMFormatS32LE, nil
	default:
		return 0, 0, errors.New("unsupported format")
	}
}

func (d *WavDecoder) GetSampleSpec() *SampleSpec {
	return d.sampleSpec
}

func (d *WavDecoder) Decode() ([]byte, error) {
	buf := make([]byte, d.bufSize)
	n, err := d.read(buf)
	return buf[:n], err
}

func (d *WavDecoder) GetDuration() time.Duration {
	return d.duration
}

func (d *WavDecoder) read(buf []byte) (int, error) {
	var n int
	for n < len(buf) {
		sample, err := d.reader.ReadRawSample()
		if err != nil {
			return n - (n % d.bytesPerFrame), err
		}
		n += copy(buf[n:], sample)
	}
	return n, nil
}

func (d *WavDecoder) Close() error {
	return d.f.Close()
}
