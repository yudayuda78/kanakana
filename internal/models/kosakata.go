package models

type Kosakata struct {
	ID       int    `json:"id"`
	Kanji    string `json:"kanji"`
	Reading  string `json:"reading"` // Hiragana atau Katakana
	Romaji   string `json:"romaji"`
	Arti     string `json:"arti"`
	LevelJLPT int    `json:"level_jlpt"` // Level JLPT (N5 - N1)
}