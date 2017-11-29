package sound_effect

import (
	"os"

	"errors"
	"github.com/cryptix/wav"
	"pkg.deepin.io/lib/asound"
	paSimple "pkg.deepin.io/lib/pulse/simple"
)

type WavDecoder struct {
	reader     *wav.Reader
	f          *os.File
	sampleSpec *SampleSpec
	bufSize    int
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

	bufSize := int(wavFile.SampleRate/8) * int(wavFile.Channels) * int(wavFile.SignificantBits/8)
	return &WavDecoder{
		f:          f,
		reader:     wavReader,
		sampleSpec: sampleSpec,
		bufSize:    bufSize,
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
	n, err := wavRead(d.reader, buf)
	var data []byte
	if n > 0 {
		data = buf[:n]
	}
	return data, err
}

// return num of bytes
func wavRead(wavReader *wav.Reader, buf []byte) (int, error) {
	var i int
	for i < len(buf) {
		data, err := wavReader.ReadRawSample()
		if err != nil {
			return i, err
		}
		i += copy(buf[i:], data)
	}
	return i, nil
}

func (d *WavDecoder) Close() error {
	return d.f.Close()
}
