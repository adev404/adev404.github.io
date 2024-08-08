package main

type indexData struct {
	Title string
	H1    string
	P     string
}

type videoData struct {
	Title     string
	H1        string
	StartName string
	StartSrc  string
	Sources   map[string]string
}

type audioData struct {
	Title   string
	H1      string
	Sources map[string]string
}

type imageData struct {
	Title    string
	H1       string
	StartSrc string
	Sources  map[string]string
}

type textData struct {
	Title   string
	H1      string
	Sources map[string]string
}
