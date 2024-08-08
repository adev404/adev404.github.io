package main

import (
	"net/http"
)

func (a *Archive) IndexHandler(w http.ResponseWriter, r *http.Request) {
	a.T.ExecuteTemplate(w, "index.html", _defaultIndexData)
}

func (a *Archive) VideoHandler(w http.ResponseWriter, r *http.Request) {
	a.T.ExecuteTemplate(w, "video.html", a.LoadVideoData())
}

func (a *Archive) AudioHandler(w http.ResponseWriter, r *http.Request) {
	a.T.ExecuteTemplate(w, "audio.html", a.LoadAudioData())
}

func (a *Archive) ImageHandler(w http.ResponseWriter, r *http.Request) {
	a.T.ExecuteTemplate(w, "image.html", a.LoadImageData())
}

func (a *Archive) TextHandler(w http.ResponseWriter, r *http.Request) {
	a.T.ExecuteTemplate(w, "text.html", a.LoadTextData())
}
