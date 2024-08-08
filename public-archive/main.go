package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
)

var (
	port = flag.String("p", "80", "Port to serve on")

	static           = flag.Bool("static", false, "Generate static website")
	generateDefaults = flag.Bool("defaults", false, "Generate default (video, audio, image, and text) source files")
)

type Archive struct {
	T *template.Template
}

func main() {
	flag.Parse()

	archive := &Archive{
		T: template.Must(template.ParseGlob("templates/*")),
	}

	if *generateDefaults {
		log.Println("Generating default source files ...")
		archive.SaveSources("video", _defaultVideoSources)
		archive.SaveSources("audio", _defaultAudioSources)
		archive.SaveSources("image", _defaultImageSources)
		archive.SaveSources("text", _defaultTextSources)
	}

	if *static {
		archive.GenerateStaticSite()
		log.Println("Successfully generated static website ...")
		os.Exit(0)
	}

	http.HandleFunc("/", archive.IndexHandler)
	http.HandleFunc("/video.html", archive.VideoHandler)
	http.HandleFunc("/audio.html", archive.AudioHandler)
	http.HandleFunc("/image.html", archive.ImageHandler)
	http.HandleFunc("/text.html", archive.TextHandler)

	log.Printf("Serving on %s ...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
