package transcription

type Transcription struct {
	Name   string  `yaml:"name"`
	Lyrics []Lyric `yaml:"lyrics"`
}

type Lyric struct {
	StartSecond    float64 `yaml:"start_second"`
	EndSecond      float64 `yaml:"end_second"`
	OriginalText   string  `yaml:"original_text"`
	TranslatedText string  `yaml:"translated_text"`
}
