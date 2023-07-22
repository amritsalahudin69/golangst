package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	ArticleID    uint          `json:"articleId"`
	Author       *Author       `json:"author"`
	Game         *Game         `json:"game,omitempty"`
	Genre        *Genre        `json:"genre,omitempty"`
	Platform     *Platform     `json:"platform,omitempty"`
	ImageArticle *ImageArticle `json:"imageArticle"`
	Translation  *Translation  `json:"translation"`
}

type Author struct {
	Fullname       string `json:"fullname"`
	ProfilePicture string `json:"profilePicture"`
	UserName       string `json:"userName"`
}

type Game struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Genre    *Genre    `json:"genre"`
	Platform *Platform `json:"platform"`
}

type Genre struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	IsHighLight int    `json:"isHighLight"`
}

type Platform struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	IsHighLight int    `json:"isHighLight"`
	Logo        string `json:"logo"`
	Image       string `json:"image"`
}

type ImageArticle struct {
	Image          string `json:"image"`
	ImageDesktop   string `json:"imageDesktop"`
	ImageHighLight string `json:"imageHighLight"`
}

type Translation struct {
	Title         string  `json:"title"`
	Slug          string  `json:"slug"`
	Excerpt       string  `json:"excerpt"`
	Locale        string  `json:"locale"`
	UrlShare      string  `json:"urlShare"`
	Lang          string  `json:"lang"`
	ArticleID     uint    `json:"articleId"`
	RatingComment *string `json:"ratingComment"`
	Meta          *Meta   `json:"meta"`
}

type Meta struct {
	Title        string  `json:"title"`
	Keyword      string  `json:"keyword"`
	Description  string  `json:"description"`
	CanonicalURL *string `json:"canonicalURL"`
}

func (article *Article) AddIdArticle(id uint) {
	article.ArticleID = id
}

func (article *Article) AddAuthorArticle(author *Author) {
	article.Author = author
}

func (article *Article) AddImageArticle(image *ImageArticle) {
	article.ImageArticle = image
}

func (article *Article) AddTranslationArticle(translation *Translation) {
	article.Translation = translation
}

func (article *Article) GetArticle() {

	jsonData, err := json.MarshalIndent(article, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func main() {

	imageArticle := `{
			"image": "https://api.duniagames.co.id/api/content/upload/file/9874800061688113932.jpg",
			"imageDesktop": "https://api.duniagames.co.id/api/content/upload/file/9074251731688113932.jpg",
			"imageHighlight": "https://api.duniagames.co.id/api/content/upload/file/9842838701688113932.jpg"
	}`
	author := `{
			"fullName":"haris13",
			"profilePicture":"https:\/\/api.duniagames.co.id\/api\/user\/upload\/image\/223816331649748676.jpg",
			"userName":"Haris Firmansyah"
	}`

	var imageData ImageArticle

	err := json.Unmarshal([]byte(imageArticle), &imageData)
	if err != nil {
		fmt.Println("Error parsing translation JSON1:", err)

	}

	var authorData Author
	err = json.Unmarshal([]byte(author), &authorData)
	if err != nil {
		fmt.Println("Error parsing translation JSON2:", err)

	}

	var translationData Translation
	err = json.Unmarshal([]byte(translation), &translationData)
	if err != nil {
		fmt.Println("Error parsing translation JSON3:", err)

	}
	var newArticle Article

	newArticle.AddIdArticle(23928)
	newArticle.AddAuthorArticle(&authorData)
	newArticle.AddImageArticle(&imageData)
	newArticle.AddTranslationArticle(&translationData)

	newArticle.GetArticle()

	translation := `{
		"title":"Benedetta Mobile Legends (ML) Hero: Skill, Build dan Kelebihan",
		"slug":"hero-benedetta-mobile-legends",
		"excerpt":"Penasaran ingin tahu tentang hero Benedetta Mobile Legends (ML)? Lihat artikel Dunia Games dan pelajari tentang skills, build, dan kelebihannya.",
		"content":"<p><span style=\"font-weight: 400;\">Benedetta adalah seorang hero<\/span><em><span style=\"font-weight: 400;\"> Assassin<\/span><\/em><span style=\"font-weight: 400;\"> di <\/span><em><span style=\"font-weight: 400;\">Mobile Legends<\/span><\/em><span style=\"font-weight: 400;\"> dengan kemampuan yang menakutkan di <\/span><em><span style=\"font-weight: 400;\">Land of Dawn<\/span><\/em><span style=\"font-weight: 400;\">. Benedetta mampu menghasilkan <\/span><em><span style=\"font-weight: 400;\">burst damage<\/span><\/em><span style=\"font-weight: 400;\"> yang tinggi dengan area yang besar, membuat musuh merasa kesulitan saat berhadapan dengan Benedetta ketika <\/span><em><span style=\"font-weight: 400;\">team fight<\/span><\/em><span style=\"font-weight: 400;\">. Penasaran ingin tahu tentang hero Benedetta Mobile Legends (ML)? Lihat artikel Dunia Games dan pelajari tentang skills, build, dan kelebihannya.<\/span><\/p>\r\n<h2><strong>Kelebihan Benedetta Hero Mobile Legends<\/strong><\/h2>\r\n<p><strong><img src=\"https:\/\/api.duniagames.co.id\/api\/content\/upload\/file\/12416826841609743104.jpg\" alt=\"Patch Notes 1.5.44 Mobile Legends: Masih Over Power, Benedetta Dapatkan  Nerf! | Dunia Games\" \/><\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Berikut kelebihan Benedetta hero <\/span><em><span style=\"font-weight: 400;\">Mobile Legends<\/span><\/em><span style=\"font-weight: 400;\">.<\/span><\/p>\r\n<p><strong style=\"font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">1. Cocok Menempati Banyak Role<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Dengan kelebihan yang dimiliki, Benedetta bisa menempati banyak role dengan sama baiknya. Benedetta bisa ditempatkan di jungle, EXP lane, roam, dan Gold lane. <\/span><\/p>\r\n<p><strong style=\"font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">2. Mobilitas Tinggi<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Benedetta punya skill pasif dan aktif yang membuatnya memiliki mobilitas tinggi di<\/span><em><span style=\"font-weight: 400;\"> Land of Dawn<\/span><\/em><span style=\"font-weight: 400;\">. Dengan memanfaatkan efek blink, Benedetta menjadi hero yang sangat lincah.<\/span><\/p>\r\n<p><strong style=\"font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">3. Cooldown Skill Singkat<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Dengan <\/span><em><span style=\"font-weight: 400;\">cooldown skill<\/span><\/em><span style=\"font-weight: 400;\"> yang cukup singkat, Benedetta bisa melakukan spam skill dan pasif.<\/span><\/p>\r\n<p><strong style=\"font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">4. Melakukan Split Push<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Benedetta bisa melakukan <\/span><em><span style=\"font-weight: 400;\">clear wave minion<\/span><\/em><span style=\"font-weight: 400;\"> dengan mudah. Sebab tipe skill Benedetta adalah <\/span><em><span style=\"font-weight: 400;\">Area of Effect<\/span><\/em><span style=\"font-weight: 400;\">. Dengan <\/span><em><span style=\"font-weight: 400;\">full item attack<\/span><\/em><span style=\"font-weight: 400;\">, Benedetta juga bisa leluasa untuk <\/span><em><span style=\"font-weight: 400;\">split push<\/span><\/em><span style=\"font-weight: 400;\"> dan meruntuhkan turret lawan.<\/span><\/p>\r\n<p><strong style=\"font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\">5. Menggocek Musuh<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Skill<\/span><em><span style=\"font-weight: 400;\"> blink<\/span><\/em><span style=\"font-weight: 400;\"> Benedetta bisa digunakan untuk menggocek musuh. Ditambah kemampuan <\/span><em><span style=\"font-weight: 400;\">fast hand<\/span><\/em><span style=\"font-weight: 400;\"> sang pilot, Benedetta bisa menjadi hero yang tak terhentikan di <\/span><em><span style=\"font-weight: 400;\">Land of Dawn<\/span><\/em><span style=\"font-weight: 400;\">.<\/span><\/p>\r\n<h2><strong>Skill Hero Benedetta ML<\/strong><\/h2>\r\n<p><strong><img src=\"https:\/\/api.duniagames.co.id\/api\/content\/upload\/file\/10109530761607572550.jpg\" alt=\"5 Assassin Mobile Legends Terbaik di Desember 2020, Natalia Masih  Mematikan! | Dunia Games\" \/><\/strong><\/p>\r\n<p><strong>Pasif - Elapsed Daytime<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Saat menahan tombol <\/span><em><span style=\"font-weight: 400;\">basic attack<\/span><\/em><span style=\"font-weight: 400;\">, Benedetta memasuki kondisi <\/span><em><span style=\"font-weight: 400;\">Swordout<\/span><\/em><span style=\"font-weight: 400;\"> dan terus mengumpulkan <\/span><em><span style=\"font-weight: 400;\">Sword Intent<\/span><\/em><span style=\"font-weight: 400;\">. Jika melepaskan tombol <\/span><em><span style=\"font-weight: 400;\">basic attack<\/span><\/em><span style=\"font-weight: 400;\"> setelah <\/span><em><span style=\"font-weight: 400;\">Sword Intent<\/span><\/em><span style=\"font-weight: 400;\"> terisi penuh, Benedetta akan menggunakan <\/span><em><span style=\"font-weight: 400;\">Swordout Slash<\/span><\/em><span style=\"font-weight: 400;\"> ke arah yang ditentukan dan melakukan dash ke depan, memberikan <em>physical damage<\/em> (diidentifikasi sebagai <\/span><em><span style=\"font-weight: 400;\">skill damage<\/span><\/em><span style=\"font-weight: 400;\">) kepada musuh di jalur.&nbsp;<\/span><\/p>\r\n<p><span style=\"font-weight: 400;\">Ketika Benedetta sedang dalam keadaan <\/span><em><span style=\"font-weight: 400;\">Swordout<\/span><\/em><span style=\"font-weight: 400;\"> dan <\/span><em><span style=\"font-weight: 400;\">Sword Intent<\/span><\/em><span style=\"font-weight: 400;\"> terisi penuh dengan menahan tombol <\/span><em><span style=\"font-weight: 400;\">basic attack<\/span><\/em><span style=\"font-weight: 400;\">, dia dapat melepaskan <\/span><em><span style=\"font-weight: 400;\">Swordout Slash<\/span><\/em><span style=\"font-weight: 400;\">. Benedetta juga bisa mendapatkan <\/span><em><span style=\"font-weight: 400;\">Sword Intent<\/span><\/em><span style=\"font-weight: 400;\"> ketika memberikan damage melalui <\/span><em><span style=\"font-weight: 400;\">basic attack<\/span><\/em><span style=\"font-weight: 400;\"> dan skill.<\/span><\/p>\r\n<p><strong>Skill 1 - Phantom Slash<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Benedetta mundur dengan cepat dan meninggalkan bayangan di depan. Setelah beberapa saat, bayangannya menebas ke depan di area berbentuk kipas, memberikan physical damage. Sementara itu, Benedetta juga melakukan dash ke depan untuk menebas dan memberikan physical damage. (Jika target terkena bayangannya, damage yang ditimbulkan oleh Benedetta sendiri akan meningkat).<\/span><\/p>\r\n<p><strong>Skill 2 - An Eye for An Eye<\/strong><span style=\"font-weight: 400;\">&nbsp;<\/span><\/p>\r\n<p><span style=\"font-weight: 400;\">Benedetta mengangkat senjatanya untuk bertahan, mendapatkan immunity terhadap <\/span><em><span style=\"font-weight: 400;\">crowd control<\/span><\/em><span style=\"font-weight: 400;\"> dan memblokir damage 1 kali dari sumber apapun selama jangka waktu tertentu. Setelah pertahanan, dia menusuk ke arah yang ditentukan, memberikan <\/span><em><span style=\"font-weight: 400;\">physical damage<\/span><\/em><span style=\"font-weight: 400;\"> dan memperlambat target. Jika Benedetta berhasil menahan efek <\/span><em><span style=\"font-weight: 400;\">crowd control<\/span><\/em><span style=\"font-weight: 400;\"> selama durasi pertahanan, ia akan memberikan stun pada target bukan efek slow.<\/span><\/p>\r\n<p><strong>Ultimate - Alecto: Final Blow<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Benedetta mengepalkan <\/span><em><span style=\"font-weight: 400;\">Alecto <\/span><\/em><span style=\"font-weight: 400;\">dan memotong ke depan setelah jeda singkat (invisible selama jeda tersebut), memberikan efek slow kepada semua target di jalan sebesar 70% selama 0,8 detik.&nbsp;<\/span><\/p>\r\n<p><span style=\"font-weight: 400;\">Setelah melakukan dash ke depan, Benedetta meledakkan Sword Intent di jalan, menahan musuh dengan Sword Intent selama 2,5 detik, memberikan physical damage dan memperlambat musuh sebesar 20% setiap 0,2 detik.<\/span><\/p>\r\n<p><strong>Baca Juga:<\/strong><\/p>\r\n<ul>\r\n<li aria-level=\"1\"><strong><a href=\"https:\/\/duniagames.co.id\/discover\/article\/hero-untuk-push-rank-ke-mythic\">10 Rekomendasi Hero Solo Rank, Cocok Buat Push Rank ke Mythic Mobile Legends!<\/a><\/strong><\/li>\r\n<li aria-level=\"1\"><a style=\"font-weight: bold; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\" href=\"https:\/\/duniagames.co.id\/discover\/article\/3-hero-mage-yang-dilarang-di-rank-mobile-legends\">3 Hero Mage yang Dilarang di Rank Mobile Legends, Berani Pick Auto Lose Streak<\/a><\/li>\r\n<li aria-level=\"1\"><a style=\"font-weight: bold; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;\" href=\"https:\/\/duniagames.co.id\/discover\/article\/daftar-skin-lancelot\">Daftar Skin Lancelot Terbaik Mobile Legends yang Pernah Dikeluarkan Moonton<\/a><\/li>\r\n<\/ul>\r\n<h2><strong>Build Item Benedetta ML<\/strong><\/h2>\r\n<p><strong><img src=\"https:\/\/api.duniagames.co.id\/api\/content\/upload\/file\/2673043271613617789.jpg\" alt=\"Prediksi 5 Hero yang Mungkin Bakal Laris di MPL ID Season 7 | Dunia Games\" \/><\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Inilah rekomendasi pro untuk build item Benedetta ML 2023.<\/span><\/p>\r\n<p><strong>1. Warrior Boots&nbsp;<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Item pertama yang harus dibeli Benedetta adalah sepatu. Sepatu ini memberikan tambahan <\/span><em><span style=\"font-weight: 400;\">movement speed <\/span><\/em><span style=\"font-weight: 400;\">dan<\/span><em><span style=\"font-weight: 400;\"> physical defense<\/span><\/em><span style=\"font-weight: 400;\">.&nbsp;<\/span><\/p>\r\n<p><strong>2. Blade of Heptaseas<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Item attack ini meningkatkan atribut DPS Benedetta. Selain itu, item ini memiliki pasif unik bernama Ambush, yang mana basic attack Benedetta akan memberikan efek slow (dengan syarat Benedetta tidak menerima damage selama 5 detik).&nbsp;<\/span><\/p>\r\n<p><strong>3. Bloodlust Axe<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Item ini meningkatkan <\/span><em><span style=\"font-weight: 400;\">physical attack<\/span><\/em><span style=\"font-weight: 400;\"> dan mengurangi <\/span><em><span style=\"font-weight: 400;\">cooldown <\/span><\/em><span style=\"font-weight: 400;\">skill.<\/span><\/p>\r\n<p><strong>4. Blade of Despair<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Item ini memberikan tambahan <\/span><em><span style=\"font-weight: 400;\">physical attack <\/span><\/em><span style=\"font-weight: 400;\">dan <\/span><em><span style=\"font-weight: 400;\">movement speed<\/span><\/em><span style=\"font-weight: 400;\">.<\/span><\/p>\r\n<p><strong>5. Athena&rsquo;s Shield<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Item defense ini meningkatkan <\/span><em><span style=\"font-weight: 400;\">HP Max, Magical Defense, <\/span><\/em><span style=\"font-weight: 400;\">dan<\/span><em><span style=\"font-weight: 400;\"> HP Regen<\/span><\/em><span style=\"font-weight: 400;\">.<\/span><\/p>\r\n<p><strong>6. War Axe<\/strong><\/p>\r\n<p><span style=\"font-weight: 400;\">Item ini dapat meningkatkan damage Benedetta. Semakin lama Benedetta berada di game tanpa tereliminasi, membuatnya semakin kuat.<\/span><\/p>\r\n<p><span style=\"font-weight: 400;\">Rekomendasi <\/span><strong><em>Build Emblem Benedetta: Assassin<\/em><\/strong><span style=\"font-weight: 400;\">, dengan <\/span><strong><em>battle spell Flicker<\/em><\/strong><span style=\"font-weight: 400;\"> atau <\/span><strong><em>Retribution<\/em><\/strong><span style=\"font-weight: 400;\">.<\/span><\/p>\r\n<h2><strong>FAQ tentang Hero Benedetta ML<\/strong><\/h2>\r\n<p><strong><img src=\"https:\/\/asset.kompas.com\/crops\/bOUHxHMU7TFw4KxBqw-46y9ji2o=\/315x236:924x642\/750x500\/data\/photo\/2020\/11\/06\/5fa4f33de5da8.jpg\" alt=\"Daftar Hero Assassin Mobile Legends Terbaik Versi Dunia Games | Dunia Games\" \/><\/strong><\/p>\r\n<h3><strong>1. Apa role Benedetta di ML?&nbsp;<\/strong><\/h3>\r\n<p><span style=\"font-weight: 400;\">Benedetta adalah hero assassin yang cocok mengisi posisi EXP Lane dan Jungle. Benedetta juga bisa jadi roamer karena lincah. Bahkan Benedetta mampu melengkapi posisi gold lane karena bisa mengumpulkan gold dengan cepat.<\/span><\/p>\r\n<h3><strong>2. Siapa yang bisa mengalahkan Benedetta?&nbsp;<\/strong><\/h3>\r\n<p><span style=\"font-weight: 400;\">Gloo dan Phoveus bisa mengalahkan Benedetta. Jadi, pastikan untuk ban kedua hero itu jika ingin menggunakan Benedetta.<\/span><\/p>\r\n<h3><strong>3. Berapa tingkat kemenangan Benedetta?&nbsp;<\/strong><\/h3>\r\n<p><em><span style=\"font-weight: 400;\">Win rate<\/span><\/em><span style=\"font-weight: 400;\"> Benedetta sempat mencapai 100% ketika dipakai Antimage di MPL ID S7. Namun, di MPL ID S8, <\/span><em><span style=\"font-weight: 400;\">win rate<\/span><\/em><span style=\"font-weight: 400;\"> Benedetta hanya 50%, tetapi termasuk tinggi dan menjadi kunci kemenangan.<\/span><\/p>\r\n<h3><strong>4. Apa kelemahan Benedetta?<\/strong><\/h3>\r\n<p><span style=\"font-weight: 400;\">Kelemahan Benedetta adalah sulit untuk digunakan bagi pemain pemula. Sebab butuh <\/span><em><span style=\"font-weight: 400;\">fast hand <\/span><\/em><span style=\"font-weight: 400;\">dan mekanik tinggi untuk memainkan hero ini. Selain itu, Benedetta juga punya skill set yang rumit. Jadi, kamu butuh waktu dan usaha ekstra untuk mempelajarinya.<\/span><\/p>\r\n<p><span style=\"font-weight: 400;\">Nah, itulah informasi lengkap tentang <\/span><em><span style=\"font-weight: 400;\">build<\/span><\/em><span style=\"font-weight: 400;\">, kekuatan dan kelemahan Benedetta <\/span><em><span style=\"font-weight: 400;\">Mobile Legends (ML)<\/span><\/em><span style=\"font-weight: 400;\">.<\/span> <span style=\"font-weight: 400;\">Biar mabar makin seru, jangan lupa <\/span><a href=\"https:\/\/duniagames.co.id\/top-up\/item\/mobile-legends\"><span style=\"font-weight: 400;\">top up mobile legends<\/span><\/a><span style=\"font-weight: 400;\"> dulu.<\/span><\/p>\r\n<p><a href=\"https:\/\/duniagames.co.id\/top-up\"><span style=\"font-weight: 400;\">Top-up Murah, banyak untungnya, hanya di Dunia Games<\/span><\/a><\/p>\r\n<p style=\"text-align: center;\"><strong>Baca Juga &gt;&gt;<\/strong> <a href=\"https:\/\/duniagames.co.id\/discover\/article\/freya-skin-special-beach-sweetheart\"><strong>Kamu Fans Berat Freya? Simak Review Freya Skin Special Beach Sweetheart Berikut!<\/strong><\/a><\/p>",
		"locale":"id",
		"urlShare":"https:\/\/duniagames.co.id\/discover\/article\/hero-benedetta-mobile-legends",
		"lang":"id",
		"articleId":23928,
		"ratingComment":null,
		"meta":{
		   "title":"Benedetta Mobile Legends (ML) Hero: Skill, Build dan Kelebihan",
		   "keyword":"Benedetta",
		   "description":"Penasaran ingin tahu tentang hero Benedetta Mobile Legends (ML)? Lihat artikel Dunia Games dan pelajari tentang skills, build, dan kelebihannya.",
		   "canonicalURL":null
		}
}`

}
