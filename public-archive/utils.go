package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	_dataDir      = "data"
	_videoSrcFile = "video-sources.json"
	_audioSrcFile = "audio-sources.json"
	_imageSrcFile = "image-sources.json"
	_textSrcFile  = "text-sources.json"
)

var (
	_sources = map[string]string{
		"video": _videoSrcFile,
		"audio": _audioSrcFile,
		"image": _imageSrcFile,
		"text":  _textSrcFile,
	}

	_pages = map[string]string{
		"index.html": "index",
		"video.html": "video",
		"audio.html": "audio",
		"image.html": "image",
		"text.html":  "text",
	}
)

func (a *Archive) LoadSources(src string) (map[string]string, error) {
	if _, e := _sources[src]; !e {
		return nil, fmt.Errorf("%s not a valid source", src)
	}
	if _, err := os.Stat(_dataDir); err != nil {
		return nil, err
	}

	b, err := os.ReadFile(fmt.Sprintf("%s/%s", _dataDir, _sources[src]))
	if err != nil {
		return nil, err
	}

	srcs := make(map[string]string)
	err = json.Unmarshal(b, &srcs)
	if err != nil {
		return nil, err
	}

	return srcs, nil
}

func (a *Archive) SaveSources(src string, srcs map[string]string) error {
	if _, e := _sources[src]; !e {
		return fmt.Errorf("%s not a valid source", src)
	}

	if _, err := os.Stat(_dataDir); err != nil {
		if err := os.Mkdir(_dataDir, 0700); err != nil {
			return err
		}
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", _dataDir, _sources[src]))
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.MarshalIndent(srcs, "", "\t")
	if err != nil {
		return err
	}

	f.Write(b)

	return nil
}

func (a *Archive) GenerateStaticSite() error {
	for fname, fdata := range _pages {
		f, err := os.Create(fname)
		if err != nil {
			return err
		}
		defer f.Close()

		// TODO: Better
		var fn any
		switch fdata {
		case "index":
			fn = _defaultIndexData
		case "video":
			fn = a.LoadVideoData()
		case "audio":
			fn = a.LoadAudioData()
		case "image":
			fn = a.LoadImageData()
		case "text":
			fn = a.LoadTextData()
		}

		a.T.ExecuteTemplate(f, fname, fn)
	}

	return nil
}
