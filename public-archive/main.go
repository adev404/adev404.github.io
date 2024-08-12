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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/public-archive", http.StatusFound)
	})

	http.HandleFunc("/public-archive", archive.IndexHandler)
	http.HandleFunc("/public-archive/video", archive.VideoHandler)
	http.HandleFunc("/public-archive/audio", archive.AudioHandler)
	http.HandleFunc("/public-archive/image", archive.ImageHandler)
	http.HandleFunc("/public-archive/text", archive.TextHandler)

	log.Printf("Serving on %s ...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
