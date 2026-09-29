package main

import (
	"fmt"
	"kanakana/internal/config"
	"kanakana/internal/models"
	"kanakana/internal/pkg/database"
	"log"
)

func main() {
	fmt.Println("Memulai Seeder Kosakata N4...")

	cfg := config.Load()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	if err := db.AutoMigrate(&models.KosakataLevel{}, &models.Kosakata{}); err != nil {
		log.Fatalf("Gagal auto migrate: %v", err)
	}

	n4Level := models.KosakataLevel{Name: "N4"}
	if err := db.FirstOrCreate(&n4Level, &models.KosakataLevel{Name: "N4"}).Error; err != nil {
		log.Fatalf("Gagal membuat level N4: %v", err)
	}
	fmt.Println("Level N4 siap")

	kosakataList := []models.Kosakata{
		// ===== Kata Kerja (動詞) =====
		{Kanji: "会う", Reading: "あう", Romaji: "au", Arti: "bertemu"},
		{Kanji: "洗う", Reading: "あらう", Romaji: "arau", Arti: "mencuci"},
		{Kanji: "言う", Reading: "いう", Romaji: "iu", Arti: "mengatakan"},
		{Kanji: "動く", Reading: "うごく", Romaji: "ugoku", Arti: "bergerak"},
		{Kanji: "歌う", Reading: "うたう", Romaji: "utau", Arti: "bernyanyi"},
		{Kanji: "打つ", Reading: "うつ", Romaji: "utsu", Arti: "memukul"},
		{Kanji: "送る", Reading: "おくる", Romaji: "okuru", Arti: "mengirim"},
		{Kanji: "押す", Reading: "おす", Romaji: "osu", Arti: "mendorong/menekan"},
		{Kanji: "考える", Reading: "かんがえる", Romaji: "kangaeru", Arti: "berpikir"},
		{Kanji: "消す", Reading: "けす", Romaji: "kesu", Arti: "mematikan/menghapus"},
		{Kanji: "答える", Reading: "こたえる", Romaji: "kotaeru", Arti: "menjawab"},
		{Kanji: "閉める", Reading: "しめる", Romaji: "shimeru", Arti: "menutup"},
		{Kanji: "知る", Reading: "しる", Romaji: "shiru", Arti: "mengetahui"},
		{Kanji: "吸う", Reading: "すう", Romaji: "suu", Arti: "menghisap/merokok"},
		{Kanji: "住む", Reading: "すむ", Romaji: "sumu", Arti: "tinggal"},
		{Kanji: "座る", Reading: "すわる", Romaji: "suwaru", Arti: "duduk"},
		{Kanji: "立つ", Reading: "たつ", Romaji: "tatsu", Arti: "berdiri"},
		{Kanji: "着く", Reading: "つく", Romaji: "tsuku", Arti: "tiba"},
		{Kanji: "作る", Reading: "つくる", Romaji: "tsukuru", Arti: "membuat"},
		{Kanji: "勤める", Reading: "つとめる", Romaji: "tsutomeru", Arti: "bekerja di"},
		{Kanji: "泊まる", Reading: "とまる", Romaji: "tomaru", Arti: "menginap"},
		{Kanji: "乗る", Reading: "のる", Romaji: "noru", Arti: "naik (kendaraan)"},
		{Kanji: "始まる", Reading: "はじまる", Romaji: "hajimaru", Arti: "mulai (intransitif)"},
		{Kanji: "働く", Reading: "はたらく", Romaji: "hataraku", Arti: "bekerja"},
		{Kanji: "引く", Reading: "ひく", Romaji: "hiku", Arti: "menarik"},
		{Kanji: "待つ", Reading: "まつ", Romaji: "matsu", Arti: "menunggu"},
		{Kanji: "休む", Reading: "やすむ", Romaji: "yasumu", Arti: "beristirahat"},
		{Kanji: "渡る", Reading: "わたる", Romaji: "wataru", Arti: "menyeberang"},
		{Kanji: "遊ぶ", Reading: "あそぶ", Romaji: "asobu", Arti: "bermain"},
		{Kanji: "泳ぐ", Reading: "およぐ", Romaji: "oyogu", Arti: "berenang"},
		{Kanji: "返す", Reading: "かえす", Romaji: "kaesu", Arti: "mengembalikan"},
		{Kanji: "手伝う", Reading: "てつだう", Romaji: "tetsudau", Arti: "membantu"},
		{Kanji: "持つ", Reading: "もつ", Romaji: "motsu", Arti: "membawa/memegang"},
		{Kanji: "直す", Reading: "なおす", Romaji: "naosu", Arti: "memperbaiki"},
		{Kanji: "習う", Reading: "ならう", Romaji: "narau", Arti: "belajar (meniru)"},
		{Kanji: "登る", Reading: "のぼる", Romaji: "noboru", Arti: "mendaki"},
		{Kanji: "運ぶ", Reading: "はこぶ", Romaji: "hakobu", Arti: "membawa (barang)"},
		{Kanji: "曲がる", Reading: "まがる", Romaji: "magaru", Arti: "berbelok"},
		{Kanji: "渡す", Reading: "わたす", Romaji: "watasu", Arti: "menyerahkan"},
		{Kanji: "怒る", Reading: "おこる", Romaji: "okoru", Arti: "marah"},
		{Kanji: "下りる", Reading: "おりる", Romaji: "oriru", Arti: "turun"},
		{Kanji: "足りる", Reading: "たりる", Romaji: "tariru", Arti: "cukup"},
		{Kanji: "着る", Reading: "きる", Romaji: "kiru", Arti: "memakai (baju)"},
		{Kanji: "教える", Reading: "おしえる", Romaji: "oshieru", Arti: "mengajar"},
		{Kanji: "借りる", Reading: "かりる", Romaji: "kariru", Arti: "meminjam"},
		{Kanji: "食べる", Reading: "たべる", Romaji: "taberu", Arti: "makan"},
		{Kanji: "寝る", Reading: "ねる", Romaji: "neru", Arti: "tidur"},
		{Kanji: "起きる", Reading: "おきる", Romaji: "okiru", Arti: "bangun"},
		{Kanji: "見る", Reading: "みる", Romaji: "miru", Arti: "melihat"},
		{Kanji: "出る", Reading: "でる", Romaji: "deru", Arti: "keluar"},
		{Kanji: "入れる", Reading: "いれる", Romaji: "ireru", Arti: "memasukkan"},
		{Kanji: "開ける", Reading: "あける", Romaji: "akeru", Arti: "membuka"},
		{Kanji: "止める", Reading: "とめる", Romaji: "tomeru", Arti: "menghentikan"},
		{Kanji: "覚える", Reading: "おぼえる", Romaji: "oboeru", Arti: "mengingat"},
		{Kanji: "始める", Reading: "はじめる", Romaji: "hajimeru", Arti: "memulai"},
		{Kanji: "見せる", Reading: "みせる", Romaji: "miseru", Arti: "memperlihatkan"},
		{Kanji: "調べる", Reading: "しらべる", Romaji: "shiraberu", Arti: "memeriksa"},
		{Kanji: "忘れる", Reading: "わすれる", Romaji: "wasureru", Arti: "lupa"},
		{Kanji: "生まれる", Reading: "うまれる", Romaji: "umareru", Arti: "dilahirkan"},
		{Kanji: "決める", Reading: "きめる", Romaji: "kimeru", Arti: "memutuskan"},
		{Kanji: "集める", Reading: "あつめる", Romaji: "atsumeru", Arti: "mengumpulkan"},
		{Kanji: "運動する", Reading: "うんどうする", Romaji: "undou suru", Arti: "berolahraga"},
		{Kanji: "買い物する", Reading: "かいものする", Romaji: "kaimono suru", Arti: "berbelanja"},
		{Kanji: "散歩する", Reading: "さんぽする", Romaji: "sanpo suru", Arti: "berjalan-jalan"},
		{Kanji: "掃除する", Reading: "そうじする", Romaji: "souji suru", Arti: "membersihkan"},
		{Kanji: "洗濯する", Reading: "せんたくする", Romaji: "sentaku suru", Arti: "mencuci pakaian"},
		{Kanji: "卒業する", Reading: "そつぎょうする", Romaji: "sotsugyou suru", Arti: "lulus"},
		{Kanji: "電話する", Reading: "でんわする", Romaji: "denwa suru", Arti: "menelepon"},
		{Kanji: "旅行する", Reading: "りょこうする", Romaji: "ryokou suru", Arti: "bepergian"},
		{Kanji: "予約する", Reading: "よやくする", Romaji: "yoyaku suru", Arti: "memesan/reservasi"},
		{Kanji: "結婚する", Reading: "けっこんする", Romaji: "kekkon suru", Arti: "menikah"},
		{Kanji: "練習する", Reading: "れんしゅうする", Romaji: "renshuu suru", Arti: "berlatih"},
		{Kanji: "紹介する", Reading: "しょうかいする", Romaji: "shoukai suru", Arti: "memperkenalkan"},
		{Kanji: "案内する", Reading: "あんないする", Romaji: "annai suru", Arti: "memandu/mengantar"},
		{Kanji: "心配する", Reading: "しんぱいする", Romaji: "shinpai suru", Arti: "khawatir"},
		{Kanji: "説明する", Reading: "せつめいする", Romaji: "setsumei suru", Arti: "menjelaskan"},
		{Kanji: "注文する", Reading: "ちゅうもんする", Romaji: "chuumon suru", Arti: "memesan (makanan)"},
		{Kanji: "準備する", Reading: "じゅんびする", Romaji: "junbi suru", Arti: "mempersiapkan"},
		{Kanji: "勉強する", Reading: "べんきょうする", Romaji: "benkyou suru", Arti: "belajar"},
		{Kanji: "料理する", Reading: "りょうりする", Romaji: "ryouri suru", Arti: "memasak"},
		{Kanji: "運転する", Reading: "うんてんする", Romaji: "unten suru", Arti: "mengemudi"},
		{Kanji: "仕事する", Reading: "しごとする", Romaji: "shigoto suru", Arti: "bekerja"},

		// ===== Kata Sifat-i (い形容詞) =====
		{Kanji: "明るい", Reading: "あかるい", Romaji: "akarui", Arti: "terang/ceria"},
		{Kanji: "痛い", Reading: "いたい", Romaji: "itai", Arti: "sakit"},
		{Kanji: "嬉しい", Reading: "うれしい", Romaji: "ureshii", Arti: "senang"},
		{Kanji: "美味しい", Reading: "おいしい", Romaji: "oishii", Arti: "enak"},
		{Kanji: "大きい", Reading: "おおきい", Romaji: "ookii", Arti: "besar"},
		{Kanji: "怖い", Reading: "こわい", Romaji: "kowai", Arti: "menakutkan"},
		{Kanji: "寂しい", Reading: "さびしい", Romaji: "sabishii", Arti: "sepi/sedih"},
		{Kanji: "楽しい", Reading: "たのしい", Romaji: "tanoshii", Arti: "menyenangkan"},
		{Kanji: "正しい", Reading: "ただしい", Romaji: "tadashii", Arti: "benar"},
		{Kanji: "近い", Reading: "ちかい", Romaji: "chikai", Arti: "dekat"},
		{Kanji: "強い", Reading: "つよい", Romaji: "tsuyoi", Arti: "kuat"},
		{Kanji: "長い", Reading: "ながい", Romaji: "nagai", Arti: "panjang"},
		{Kanji: "難しい", Reading: "むずかしい", Romaji: "muzukashii", Arti: "sulit"},
		{Kanji: "優しい", Reading: "やさしい", Romaji: "yasashii", Arti: "lembut/baik hati"},
		{Kanji: "安い", Reading: "やすい", Romaji: "yasui", Arti: "murah"},
		{Kanji: "悪い", Reading: "わるい", Romaji: "warui", Arti: "buruk"},
		{Kanji: "若い", Reading: "わかい", Romaji: "wakai", Arti: "muda"},
		{Kanji: "広い", Reading: "ひろい", Romaji: "hiroi", Arti: "luas"},
		{Kanji: "狭い", Reading: "せまい", Romaji: "semai", Arti: "sempit"},
		{Kanji: "遠い", Reading: "とおい", Romaji: "tooi", Arti: "jauh"},
		{Kanji: "低い", Reading: "ひくい", Romaji: "hikui", Arti: "rendah/pendek"},
		{Kanji: "暗い", Reading: "くらい", Romaji: "kurai", Arti: "gelap"},
		{Kanji: "忙しい", Reading: "いそがしい", Romaji: "isogashii", Arti: "sibuk"},
		{Kanji: "眠い", Reading: "ねむい", Romaji: "nemui", Arti: "mengantuk"},
		{Kanji: "細い", Reading: "ほそい", Romaji: "hosoi", Arti: "tipis/langsing"},
		{Kanji: "太い", Reading: "ふとい", Romaji: "futoi", Arti: "tebal/gemuk"},
		{Kanji: "軽い", Reading: "かるい", Romaji: "karui", Arti: "ringan"},
		{Kanji: "重い", Reading: "おもい", Romaji: "omoi", Arti: "berat"},
		{Kanji: "丸い", Reading: "まるい", Romaji: "marui", Arti: "bulat"},
		{Kanji: "白い", Reading: "しろい", Romaji: "shiroi", Arti: "putih"},
		{Kanji: "黒い", Reading: "くろい", Romaji: "kuroi", Arti: "hitam"},
		{Kanji: "赤い", Reading: "あかい", Romaji: "akai", Arti: "merah"},
		{Kanji: "青い", Reading: "あおい", Romaji: "aoi", Arti: "biru"},
		{Kanji: "黄色い", Reading: "きいろい", Romaji: "kiiroi", Arti: "kuning"},
		{Kanji: "暑い", Reading: "あつい", Romaji: "atsui", Arti: "panas (cuaca)"},
		{Kanji: "熱い", Reading: "あつい", Romaji: "atsui", Arti: "panas (benda)"},
		{Kanji: "寒い", Reading: "さむい", Romaji: "samui", Arti: "dingin (cuaca)"},
		{Kanji: "冷たい", Reading: "つめたい", Romaji: "tsumetai", Arti: "dingin (benda)"},
		{Kanji: "温かい", Reading: "あたたかい", Romaji: "atatakai", Arti: "hangat"},
		{Kanji: "早い", Reading: "はやい", Romaji: "hayai", Arti: "cepat/pagi"},
		{Kanji: "遅い", Reading: "おそい", Romaji: "osoi", Arti: "lambat/terlambat"},
		{Kanji: "良い", Reading: "よい", Romaji: "yoi", Arti: "baik"},
		{Kanji: "少ない", Reading: "すくない", Romaji: "sukunai", Arti: "sedikit"},
		{Kanji: "多い", Reading: "おおい", Romaji: "ooi", Arti: "banyak"},
		{Kanji: "古い", Reading: "ふるい", Romaji: "furui", Arti: "tua/lama"},
		{Kanji: "新しい", Reading: "あたらしい", Romaji: "atarashii", Arti: "baru"},
		{Kanji: "高い", Reading: "たかい", Romaji: "takai", Arti: "tinggi/mahal"},

		// ===== Kata Sifat-na (な形容詞) =====
		{Kanji: "元気", Reading: "げんき", Romaji: "genki", Arti: "sehat/energik"},
		{Kanji: "便利", Reading: "べんり", Romaji: "benri", Arti: "praktis"},
		{Kanji: "不便", Reading: "ふべん", Romaji: "fuben", Arti: "tidak praktis"},
		{Kanji: "有名", Reading: "ゆうめい", Romaji: "yuumei", Arti: "terkenal"},
		{Kanji: "必要", Reading: "ひつよう", Romaji: "hitsuyou", Arti: "perlu"},
		{Kanji: "静か", Reading: "しずか", Romaji: "shizuka", Arti: "tenang"},
		{Kanji: "賑やか", Reading: "にぎやか", Romaji: "nigiyaka", Arti: "ramai"},
		{Kanji: "暇", Reading: "ひま", Romaji: "hima", Arti: "senggang"},
		{Kanji: "簡単", Reading: "かんたん", Romaji: "kantan", Arti: "mudah/sederhana"},
		{Kanji: "大切", Reading: "たいせつ", Romaji: "taisetsu", Arti: "penting"},
		{Kanji: "大丈夫", Reading: "だいじょうぶ", Romaji: "daijoubu", Arti: "baik-baik saja"},
		{Kanji: "立派", Reading: "りっぱ", Romaji: "rippa", Arti: "hebat/megah"},
		{Kanji: "安全", Reading: "あんぜん", Romaji: "anzen", Arti: "aman"},
		{Kanji: "危険", Reading: "きけん", Romaji: "kiken", Arti: "berbahaya"},
		{Kanji: "複雑", Reading: "ふくざつ", Romaji: "fukuzatsu", Arti: "kompleks"},
		{Kanji: "真面目", Reading: "まじめ", Romaji: "majime", Arti: "serius/rajin"},
		{Kanji: "自由", Reading: "じゆう", Romaji: "jiyuu", Arti: "bebas"},
		{Kanji: "幸せ", Reading: "しあわせ", Romaji: "shiawase", Arti: "bahagia"},
		{Kanji: "得意", Reading: "とくい", Romaji: "tokui", Arti: "pandai/bangga"},
		{Kanji: "苦手", Reading: "にがて", Romaji: "nigate", Arti: "tidak pandai"},
		{Kanji: "好き", Reading: "すき", Romaji: "suki", Arti: "suka"},
		{Kanji: "嫌い", Reading: "きらい", Romaji: "kirai", Arti: "tidak suka"},
		{Kanji: "上手", Reading: "じょうず", Romaji: "jouzu", Arti: "pandai"},
		{Kanji: "下手", Reading: "へた", Romaji: "heta", Arti: "tidak pandai"},

		// ===== Kata Benda: Orang/Keluarga (名詞: 人・家族) =====
		{Kanji: "家族", Reading: "かぞく", Romaji: "kazoku", Arti: "keluarga"},
		{Kanji: "両親", Reading: "りょうしん", Romaji: "ryoushin", Arti: "orang tua"},
		{Kanji: "親", Reading: "おや", Romaji: "oya", Arti: "orang tua"},
		{Kanji: "息子", Reading: "むすこ", Romaji: "musuko", Arti: "anak laki-laki"},
		{Kanji: "娘", Reading: "むすめ", Romaji: "musume", Arti: "anak perempuan"},
		{Kanji: "夫", Reading: "おっと", Romaji: "otto", Arti: "suami"},
		{Kanji: "妻", Reading: "つま", Romaji: "tsuma", Arti: "istri"},
		{Kanji: "兄", Reading: "あに", Romaji: "ani", Arti: "kakak laki-laki (saya)"},
		{Kanji: "姉", Reading: "あね", Romaji: "ane", Arti: "kakak perempuan (saya)"},
		{Kanji: "弟", Reading: "おとうと", Romaji: "otouto", Arti: "adik laki-laki"},
		{Kanji: "妹", Reading: "いもうと", Romaji: "imouto", Arti: "adik perempuan"},
		{Kanji: "友達", Reading: "ともだち", Romaji: "tomodachi", Arti: "teman"},
		{Kanji: "隣", Reading: "となり", Romaji: "tonari", Arti: "tetangga"},
		{Kanji: "先生", Reading: "せんせい", Romaji: "sensei", Arti: "guru"},
		{Kanji: "医者", Reading: "いしゃ", Romaji: "isha", Arti: "dokter"},
		{Kanji: "店員", Reading: "てんいん", Romaji: "tenin", Arti: "pegawai toko"},
		{Kanji: "警察", Reading: "けいさつ", Romaji: "keisatsu", Arti: "polisi"},
		{Kanji: "弁護士", Reading: "べんごし", Romaji: "bengoshi", Arti: "pengacara"},
		{Kanji: "会社員", Reading: "かいしゃいん", Romaji: "kaishain", Arti: "karyawan"},

		// ===== Kata Benda: Makanan (名詞: 食べ物) =====
		{Kanji: "朝ご飯", Reading: "あさごはん", Romaji: "asagohan", Arti: "sarapan"},
		{Kanji: "昼ご飯", Reading: "ひるごはん", Romaji: "hirugohan", Arti: "makan siang"},
		{Kanji: "晩ご飯", Reading: "ばんごはん", Romaji: "bangohan", Arti: "makan malam"},
		{Kanji: "肉", Reading: "にく", Romaji: "niku", Arti: "daging"},
		{Kanji: "魚", Reading: "さかな", Romaji: "sakana", Arti: "ikan"},
		{Kanji: "野菜", Reading: "やさい", Romaji: "yasai", Arti: "sayuran"},
		{Kanji: "果物", Reading: "くだもの", Romaji: "kudamono", Arti: "buah"},
		{Kanji: "牛乳", Reading: "ぎゅうにゅう", Romaji: "gyuunyuu", Arti: "susu sapi"},
		{Kanji: "卵", Reading: "たまご", Romaji: "tamago", Arti: "telur"},
		{Kanji: "醤油", Reading: "しょうゆ", Romaji: "shouyu", Arti: "kecap asin"},
		{Kanji: "味噌汁", Reading: "みそしる", Romaji: "misoshiru", Arti: "sup miso"},
		{Kanji: "お弁当", Reading: "おべんとう", Romaji: "obentou", Arti: "bekal"},
		{Kanji: "お菓子", Reading: "おかし", Romaji: "okashi", Arti: "camilan/manisan"},
		{Kanji: "料理", Reading: "りょうり", Romaji: "ryouri", Arti: "masakan"},
		{Kanji: "材料", Reading: "ざいりょう", Romaji: "zairyou", Arti: "bahan"},
		{Kanji: "お茶", Reading: "おちゃ", Romaji: "ocha", Arti: "teh"},
		{Kanji: "水", Reading: "みず", Romaji: "mizu", Arti: "air"},
		{Kanji: "酒", Reading: "さけ", Romaji: "sake", Arti: "minuman keras/mirin"},

		// ===== Kata Benda: Tempat/Bangunan (名詞: 場所・建物) =====
		{Kanji: "病院", Reading: "びょういん", Romaji: "byouin", Arti: "rumah sakit"},
		{Kanji: "銀行", Reading: "ぎんこう", Romaji: "ginkou", Arti: "bank"},
		{Kanji: "郵便局", Reading: "ゆうびんきょく", Romaji: "yuubinkyoku", Arti: "kantor pos"},
		{Kanji: "図書館", Reading: "としょかん", Romaji: "toshokan", Arti: "perpustakaan"},
		{Kanji: "公園", Reading: "こうえん", Romaji: "kouen", Arti: "taman"},
		{Kanji: "駅", Reading: "えき", Romaji: "eki", Arti: "stasiun"},
		{Kanji: "空港", Reading: "くうこう", Romaji: "kuukou", Arti: "bandara"},
		{Kanji: "ホテル", Reading: "ほてる", Romaji: "hoteru", Arti: "hotel"},
		{Kanji: "神社", Reading: "じんじゃ", Romaji: "jinja", Arti: "kuil Shinto"},
		{Kanji: "お寺", Reading: "おてら", Romaji: "otera", Arti: "kuil Buddha"},
		{Kanji: "交差点", Reading: "こうさてん", Romaji: "kousaten", Arti: "persimpangan"},
		{Kanji: "信号", Reading: "しんごう", Romaji: "shingou", Arti: "lampu lalu lintas"},
		{Kanji: "角", Reading: "かど", Romaji: "kado", Arti: "sudut/pojok"},
		{Kanji: "店", Reading: "みせ", Romaji: "mise", Arti: "toko"},
		{Kanji: "会社", Reading: "かいしゃ", Romaji: "kaisha", Arti: "perusahaan"},
		{Kanji: "大学", Reading: "だいがく", Romaji: "daigaku", Arti: "universitas"},
		{Kanji: "旅館", Reading: "りょかん", Romaji: "ryokan", Arti: "penginapan tradisional"},

		// ===== Kata Benda: Waktu/Frekuensi (名詞: 時間・頻度) =====
		{Kanji: "今週", Reading: "こんしゅう", Romaji: "konshuu", Arti: "minggu ini"},
		{Kanji: "来週", Reading: "らいしゅう", Romaji: "raishuu", Arti: "minggu depan"},
		{Kanji: "先週", Reading: "せんしゅう", Romaji: "senshuu", Arti: "minggu lalu"},
		{Kanji: "今月", Reading: "こんげつ", Romaji: "kongetsu", Arti: "bulan ini"},
		{Kanji: "来月", Reading: "らいげつ", Romaji: "raigetsu", Arti: "bulan depan"},
		{Kanji: "先月", Reading: "せんげつ", Romaji: "sengetsu", Arti: "bulan lalu"},
		{Kanji: "今年", Reading: "ことし", Romaji: "kotoshi", Arti: "tahun ini"},
		{Kanji: "来年", Reading: "らいねん", Romaji: "rainen", Arti: "tahun depan"},
		{Kanji: "去年", Reading: "きょねん", Romaji: "kyonen", Arti: "tahun lalu"},
		{Kanji: "毎朝", Reading: "まいあさ", Romaji: "maiasa", Arti: "setiap pagi"},
		{Kanji: "毎晩", Reading: "まいばん", Romaji: "maiban", Arti: "setiap malam"},
		{Kanji: "毎週", Reading: "まいしゅう", Romaji: "maishuu", Arti: "setiap minggu"},
		{Kanji: "毎月", Reading: "まいつき", Romaji: "maitsuki", Arti: "setiap bulan"},
		{Kanji: "毎年", Reading: "まいとし", Romaji: "maitoshi", Arti: "setiap tahun"},
		{Kanji: "朝", Reading: "あさ", Romaji: "asa", Arti: "pagi"},
		{Kanji: "昼", Reading: "ひる", Romaji: "hiru", Arti: "siang"},
		{Kanji: "晩", Reading: "ばん", Romaji: "ban", Arti: "malam"},
		{Kanji: "夜", Reading: "よる", Romaji: "yoru", Arti: "malam"},
		{Kanji: "夕方", Reading: "ゆうがた", Romaji: "yuugata", Arti: "sore"},
		{Kanji: "終わり", Reading: "おわり", Romaji: "owari", Arti: "akhir"},
		{Kanji: "始まり", Reading: "はじまり", Romaji: "hajimari", Arti: "awal"},
		{Kanji: "真ん中", Reading: "まんなか", Romaji: "mannaka", Arti: "tengah"},

		// ===== Kata Benda: Alam/Cuaca (名詞: 自然・天気) =====
		{Kanji: "天気", Reading: "てんき", Romaji: "tenki", Arti: "cuaca"},
		{Kanji: "天気予報", Reading: "てんきよほう", Romaji: "tenki yohou", Arti: "ramalan cuaca"},
		{Kanji: "気候", Reading: "きこう", Romaji: "kikou", Arti: "iklim"},
		{Kanji: "季節", Reading: "きせつ", Romaji: "kisetsu", Arti: "musim"},
		{Kanji: "春", Reading: "はる", Romaji: "haru", Arti: "musim semi"},
		{Kanji: "夏", Reading: "なつ", Romaji: "natsu", Arti: "musim panas"},
		{Kanji: "秋", Reading: "あき", Romaji: "aki", Arti: "musim gugur"},
		{Kanji: "冬", Reading: "ふゆ", Romaji: "fuyu", Arti: "musim dingin"},
		{Kanji: "台風", Reading: "たいふう", Romaji: "taifuu", Arti: "topan"},
		{Kanji: "地震", Reading: "じしん", Romaji: "jishin", Arti: "gempa bumi"},
		{Kanji: "天気図", Reading: "てんきず", Romaji: "tenkizu", Arti: "peta cuaca"},

		// ===== Kata Benda: Lainnya (名詞: その他) =====
		{Kanji: "海外", Reading: "かいがい", Romaji: "kaigai", Arti: "luar negeri"},
		{Kanji: "旅行", Reading: "りょこう", Romaji: "ryokou", Arti: "perjalanan"},
		{Kanji: "仕事", Reading: "しごと", Romaji: "shigoto", Arti: "pekerjaan"},
		{Kanji: "買い物", Reading: "かいもの", Romaji: "kaimono", Arti: "belanja"},
		{Kanji: "住所", Reading: "じゅうしょ", Romaji: "juusho", Arti: "alamat"},
		{Kanji: "名前", Reading: "なまえ", Romaji: "namae", Arti: "nama"},
		{Kanji: "電話番号", Reading: "でんわばんごう", Romaji: "denwa bangou", Arti: "nomor telepon"},

		// ===== Adverbia (副詞) =====
		{Kanji: "いつも", Reading: "いつも", Romaji: "itsumo", Arti: "selalu"},
		{Kanji: "よく", Reading: "よく", Romaji: "yoku", Arti: "sering/baik"},
		{Kanji: "時々", Reading: "ときどき", Romaji: "tokidoki", Arti: "kadang-kadang"},
		{Kanji: "たまに", Reading: "たまに", Romaji: "tamani", Arti: "kadang-kadang (jarang)"},
		{Kanji: "ほとんど", Reading: "ほとんど", Romaji: "hotondo", Arti: "hampir semua"},
		{Kanji: "もう", Reading: "もう", Romaji: "mou", Arti: "sudah"},
		{Kanji: "まだ", Reading: "まだ", Romaji: "mada", Arti: "masih"},
		{Kanji: "すぐ", Reading: "すぐ", Romaji: "sugu", Arti: "segera"},
		{Kanji: "とても", Reading: "とても", Romaji: "totemo", Arti: "sangat"},
		{Kanji: "大変", Reading: "たいへん", Romaji: "taihen", Arti: "sangat/sukar"},
		{Kanji: "ちょっと", Reading: "ちょっと", Romaji: "chotto", Arti: "sedikit"},
		{Kanji: "あまり", Reading: "あまり", Romaji: "amari", Arti: "tidak terlalu (dgn negatif)"},
		{Kanji: "ぜひ", Reading: "ぜひ", Romaji: "zehi", Arti: "pasti (harapan)"},
		{Kanji: "きっと", Reading: "きっと", Romaji: "kitto", Arti: "pasti (keyakinan)"},
		{Kanji: "たぶん", Reading: "たぶん", Romaji: "tabun", Arti: "mungkin"},
		{Kanji: "もちろん", Reading: "もちろん", Romaji: "mochiron", Arti: "tentu saja"},
		{Kanji: "確かに", Reading: "たしかに", Romaji: "tashikani", Arti: "memang"},
		{Kanji: "一緒に", Reading: "いっしょに", Romaji: "isshoni", Arti: "bersama-sama"},
		{Kanji: "急に", Reading: "きゅうに", Romaji: "kyuuni", Arti: "tiba-tiba"},
		{Kanji: "全然", Reading: "ぜんぜん", Romaji: "zenzen", Arti: "sama sekali (dgn negatif)"},
		{Kanji: "ちゃんと", Reading: "ちゃんと", Romaji: "chanto", Arti: "dengan benar"},
		{Kanji: "ゆっくり", Reading: "ゆっくり", Romaji: "yukkuri", Arti: "perlahan"},
		{Kanji: "はっきり", Reading: "はっきり", Romaji: "hakkiri", Arti: "dengan jelas"},
		{Kanji: "たくさん", Reading: "たくさん", Romaji: "takusan", Arti: "banyak"},
		{Kanji: "全部", Reading: "ぜんぶ", Romaji: "zenbu", Arti: "semua"},

		// ===== Kata Sambung/Konektor (接続詞) =====
		{Kanji: "それで", Reading: "それで", Romaji: "sorede", Arti: "karena itu"},
		{Kanji: "それに", Reading: "それに", Romaji: "soreni", Arti: "lagi pula"},
		{Kanji: "しかし", Reading: "しかし", Romaji: "shikashi", Arti: "namun"},
		{Kanji: "ところで", Reading: "ところで", Romaji: "tokorode", Arti: "omong-omong"},
		{Kanji: "だから", Reading: "だから", Romaji: "dakara", Arti: "jadi"},
		{Kanji: "けれども", Reading: "けれども", Romaji: "keredomo", Arti: "tetapi"},
		{Kanji: "そして", Reading: "そして", Romaji: "soshite", Arti: "dan kemudian"},
		{Kanji: "それから", Reading: "それから", Romaji: "sorekara", Arti: "setelah itu"},
		{Kanji: "または", Reading: "または", Romaji: "matawa", Arti: "atau"},
		{Kanji: "でも", Reading: "でも", Romaji: "demo", Arti: "tetapi"},

		// ===== Kata Tanya (疑問詞) =====
		{Kanji: "いつ", Reading: "いつ", Romaji: "itsu", Arti: "kapan"},
		{Kanji: "どこ", Reading: "どこ", Romaji: "doko", Arti: "di mana"},
		{Kanji: "なぜ", Reading: "なぜ", Romaji: "naze", Arti: "mengapa"},
		{Kanji: "どうして", Reading: "どうして", Romaji: "doushite", Arti: "mengapa (kasual)"},
		{Kanji: "どの", Reading: "どの", Romaji: "dono", Arti: "yang mana"},
		{Kanji: "どんな", Reading: "どんな", Romaji: "donna", Arti: "jenis apa"},

		// ===== Partikel (助詞) =====
		{Kanji: "～ごろ", Reading: "～ごろ", Romaji: "goro", Arti: "sekitar (waktu)"},
		{Kanji: "～まで", Reading: "～まで", Romaji: "made", Arti: "sampai"},
		{Kanji: "～から", Reading: "～から", Romaji: "kara", Arti: "dari"},
		{Kanji: "～しか", Reading: "～しか", Romaji: "shika", Arti: "hanya (dgn negatif)"},

		// ===== Counter (助数詞) =====
		{Kanji: "一つ", Reading: "ひとつ", Romaji: "hitotsu", Arti: "satu (benda umum)"},
		{Kanji: "二つ", Reading: "ふたつ", Romaji: "futatsu", Arti: "dua (benda umum)"},
		{Kanji: "三つ", Reading: "みっつ", Romaji: "mittsu", Arti: "tiga (benda umum)"},
		{Kanji: "四つ", Reading: "よっつ", Romaji: "yottsu", Arti: "empat (benda umum)"},
		{Kanji: "五つ", Reading: "いつつ", Romaji: "itsutsu", Arti: "lima (benda umum)"},
		{Kanji: "一人", Reading: "ひとり", Romaji: "hitori", Arti: "satu orang"},
		{Kanji: "二人", Reading: "ふたり", Romaji: "futari", Arti: "dua orang"},

		// ===== Kata Depan/Posisi (位置) =====
		{Kanji: "上", Reading: "うえ", Romaji: "ue", Arti: "atas"},
		{Kanji: "下", Reading: "した", Romaji: "shita", Arti: "bawah"},
		{Kanji: "前", Reading: "まえ", Romaji: "mae", Arti: "depan"},
		{Kanji: "後ろ", Reading: "うしろ", Romaji: "ushiro", Arti: "belakang"},
		{Kanji: "右", Reading: "みぎ", Romaji: "migi", Arti: "kanan"},
		{Kanji: "左", Reading: "ひだり", Romaji: "hidari", Arti: "kiri"},
		{Kanji: "中", Reading: "なか", Romaji: "naka", Arti: "dalam"},
		{Kanji: "外", Reading: "そと", Romaji: "soto", Arti: "luar"},
		{Kanji: "間", Reading: "あいだ", Romaji: "aida", Arti: "antara"},
		{Kanji: "横", Reading: "よこ", Romaji: "yoko", Arti: "samping"},
	}

	for i := range kosakataList {
		kosakataList[i].Levels = []models.KosakataLevel{n4Level}
	}

	batchSize := 50
	total := len(kosakataList)
	inserted := 0

	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}
		batch := kosakataList[i:end]

		if err := db.CreateInBatches(&batch, len(batch)).Error; err != nil {
			log.Fatalf("Gagal insert batch %d-%d: %v", i+1, end, err)
		}
		inserted += len(batch)
		fmt.Printf("Insert %d/%d kosakata N4...\n", inserted, total)
	}

	fmt.Println("==================================================")
	fmt.Printf("Sukses! %d kosakata N4 berhasil ditambahkan.\n", inserted)
	fmt.Println("==================================================")
}