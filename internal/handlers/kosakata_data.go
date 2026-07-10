package handlers

import "kanakana/internal/models"

// daftarKosakata adalah data dummy kosakata bahasa Jepang
var daftarKosakata = []models.Kosakata{
	{
		ID:        1,
		Kanji:     "日本語",
		Reading:   "にほんご",
		Romaji:    "Nihongo",
		Arti:      "Bahasa Jepang",
		LevelJLPT: 5,
	},
	{
		ID:        2,
		Kanji:     "食べる",
		Reading:   "たべる",
		Romaji:    "Taberu",
		Arti:      "Makan",
		LevelJLPT: 5,
	},
	{
		ID:        3,
		Kanji:     "先生",
		Reading:   "せんせい",
		Romaji:    "Sensei",
		Arti:      "Guru",
		LevelJLPT: 5,
	},
}
