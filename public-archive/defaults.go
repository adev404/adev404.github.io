package main

import "log"

var (
	_defaultVideoSources = map[string]string{
		"Code Rush 2000":        "https://ia801809.us.archive.org/6/items/code-rush-2000/Code%20Rush%20%282000%29-Decombed.ia.mp4",
		"RIP a Remix Manifesto": "https://ia600403.us.archive.org/27/items/rip-a-remix-manifesto-8040182/RIP%20%EF%BC%9A%20A%20Remix%20Manifesto%20%5B8040182%5D.mp4",
		"Good Copy Bad Copy":    "https://ia802809.us.archive.org/10/items/Good_Copy_Bad_Copy-film/GCBC_SCR.XviD_512kb.mp4",
		"TPB AFK 2013":          "https://ia802904.us.archive.org/0/items/TpbAfkThePirateBayAwayFromKeyboard/TPB.AFK.2013.480p.vp8-SimonKlose.mp4",
		"Cosmos Laundromat":     "https://upload.wikimedia.org/wikipedia/commons/transcoded/3/36/Cosmos_Laundromat_-_First_Cycle_-_Official_Blender_Foundation_release.webm/Cosmos_Laundromat_-_First_Cycle_-_Official_Blender_Foundation_release.webm.720p.vp9.webm",
		"Spring Blender 2019":   "https://upload.wikimedia.org/wikipedia/commons/transcoded/a/a5/Spring_-_Blender_Open_Movie.webm/Spring_-_Blender_Open_Movie.webm.720p.vp9.webm",
	}

	_defaultAudioSources = map[string]string{
		"Debussy - Clair de lune":                    "https://upload.wikimedia.org/wikipedia/commons/b/be/Clair_de_lune_%28Claude_Debussy%29_Suite_bergamasque.ogg",
		"Vivaldi - Winter movement 1":                "https://upload.wikimedia.org/wikipedia/commons/0/04/Vivaldi_Winter_mvt_1_Allegro_non_molto_-_The_USAF_Concert.ogg",
		"Grieg - Morning Mood":                       "https://upload.wikimedia.org/wikipedia/commons/1/1a/Musopen_-_Morning.ogg",
		"Beethoven - Symphony No. 5 movement 1":      "https://upload.wikimedia.org/wikipedia/commons/5/5b/Ludwig_van_Beethoven_-_Symphonie_5_c-moll_-_1._Allegro_con_brio.ogg",
		"Beethoven - Fur Elise":                      "https://upload.wikimedia.org/wikipedia/commons/7/7b/FurElise.ogg",
		"Beethoven - Moonlight Sonata movement 1":    "https://upload.wikimedia.org/wikipedia/commons/e/eb/Beethoven_Moonlight_1st_movement.ogg",
		"Beethoven - Moonlight Sonata movement 3":    "https://upload.wikimedia.org/wikipedia/commons/d/d4/Beethoven_Moonlight_3rd_movement.ogg",
		"Vivaldi - Summer movement 3":                "https://upload.wikimedia.org/wikipedia/commons/f/fa/Vivaldi_-_Four_Seasons_2_Summer_mvt_3_Presto_-_John_Harrison_violin.oga",
		"Wagner - Rise of the Valkyries":             "https://upload.wikimedia.org/wikipedia/commons/9/91/Ride_of_the_Valkyries.ogg",
		"Chopin - Nocturnes Op. 9":                   "https://upload.wikimedia.org/wikipedia/commons/5/5c/Frederic_Chopin_-_Nocturne_Eb_major_Opus_9%2C_number_2.ogg",
		"Pachelbel - Pachelbel's canon":              "https://upload.wikimedia.org/wikipedia/commons/5/59/Kevin_MacLeod_-_Canon_in_D_Major.ogg",
		"Orff - O Fortuna":                           "https://upload.wikimedia.org/wikipedia/en/4/4e/Carl_Orff-Carmina_Burana-O_Fortuna.ogg",
		"Strauss - The Blue Danube":                  "https://upload.wikimedia.org/wikipedia/commons/9/91/Strauss%2C_An_der_sch%C3%B6nen_blauen_Donau.ogg",
		"Satie - Gymnopedies":                        "https://upload.wikimedia.org/wikipedia/commons/9/90/Erik_Satie_-_gymnopedies_-_la_1_ere._lent_et_douloureux.ogg",
		"Mozart - Introitus":                         "https://upload.wikimedia.org/wikipedia/commons/6/66/W._A._Mozart_-_Requiem%2C_K._626_%28Bruno_Walter%2C_Wiener_Philharmoniker%2C_Wiener_Staatsopernchor%2C_1956%29_-_01._Requiem_aeternam.ogg",
		"Mozart - Dies irae":                         "https://upload.wikimedia.org/wikipedia/commons/1/10/W._A._Mozart_-_Requiem%2C_K._626_%28Bruno_Walter%2C_Wiener_Philharmoniker%2C_Wiener_Staatsopernchor%2C_1956%29_-_03._Dies_irae.ogg",
		"Mozart - Lacrimosa":                         "https://upload.wikimedia.org/wikipedia/commons/5/5b/W._A._Mozart_-_Requiem%2C_K._626_%28Bruno_Walter%2C_Wiener_Philharmoniker%2C_Wiener_Staatsopernchor%2C_1956%29_-_08._Lacrimosa.ogg",
		"Mozart - Eine kleine Nachtmusik movement 1": "https://upload.wikimedia.org/wikipedia/commons/2/24/Mozart_-_Eine_kleine_Nachtmusik_-_1._Allegro.ogg",
		"Mozart - Symphony No. 40 movement 1":        "https://upload.wikimedia.org/wikipedia/commons/9/99/Wolfgang_Amadeus_Mozart_-_Symphony_40_g-moll_-_1._Molto_allegro.ogg",
		"Mozart - Piano Sonata No. 11 movement 3":    "https://upload.wikimedia.org/wikipedia/commons/b/bf/Mozart_-_Piano_Sonata_No._11_in_A_major_-_III._Allegro_%28Turkish_March%29.ogg",
		"Bach - Ave Maria":                           "https://upload.wikimedia.org/wikipedia/commons/d/d4/JOHN_MICHEL_CELLO-BACH_AVE_MARIA.ogg",
	}

	_defaultImageSources = map[string]string{
		"Schlacht bei Issus":    "https://upload.wikimedia.org/wikipedia/commons/4/4a/Albrecht_Altdorfer_-_Schlacht_bei_Issus_%28Alte_Pinakothek%2C_M%C3%BCnchen%29_-_Google_Art_Project.jpg",
		"Cafe Terrace at Night": "https://upload.wikimedia.org/wikipedia/commons/b/b0/Vincent_van_Gogh_%281853-1890%29_Caf%C3%A9terras_bij_nacht_%28place_du_Forum%29_Kr%C3%B6ller-M%C3%BCller_Museum_Otterlo_23-8-2016_13-35-40.JPG",
		"El Coloso":             "https://upload.wikimedia.org/wikipedia/commons/b/be/El_coloso.jpg",
	}

	_defaultTextSources = map[string]string{
		"Nineteen Eighty-Four": "http://gutenberg.net.au/ebooks01/0100021.txt",
		"Principa Mathematica": "https://en.wikisource.org/wiki/Russell_%26_Whitehead%27s_Principia_Mathematica",
		"Don Quixote":          "https://www.gutenberg.org/cache/epub/996/pg996.txt",
		"The Republic":         "https://www.gutenberg.org/cache/epub/1497/pg1497.txt",
	}

	_defaultIndexData = indexData{
		Title: "Index",
		H1:    "Public Archive",
		P: `Public Archive is a collection of technical implementations for learning,
		and permissive licensed works for sharing.`,
	}
)

func (a *Archive) LoadVideoData() *videoData {
	sources, err := a.LoadSources("video")
	if err != nil {
		log.Printf("a.LoadSources(video): %s", err.Error())
	}

	startName := "Code Rush 2000"
	startSrc := sources[startName]

	vd := &videoData{
		Title:     "Video",
		H1:        "Public Archive: Video",
		StartName: startName,
		StartSrc:  startSrc,
		Sources:   sources,
	}

	return vd
}

func (a *Archive) LoadAudioData() *audioData {
	sources, err := a.LoadSources("audio")
	if err != nil {
		log.Printf("a.LoadSources(audio): %s", err.Error())
	}

	ad := &audioData{
		Title:   "Audio",
		H1:      "Public Archive: Audio",
		Sources: sources,
	}

	return ad
}

func (a *Archive) LoadImageData() *imageData {
	sources, err := a.LoadSources("image")
	if err != nil {
		log.Printf("a.LoadSources(image): %s", err.Error())
	}

	id := &imageData{
		Title:   "Image",
		H1:      "Public Archive: Image",
		Sources: sources,
	}

	return id
}

func (a *Archive) LoadTextData() *textData {
	sources, err := a.LoadSources("text")
	if err != nil {
		log.Printf("a.LoadSources(text): %s", err.Error())
	}

	td := &textData{
		Title:   "Text",
		H1:      "Public Archive: Text",
		Sources: sources,
	}

	return td
}
