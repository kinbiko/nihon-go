package nihon

var hiragana = map[string]kana{

	"a":   {val: "あ", skip: 1},
	"i":   {val: "い", skip: 1},
	"u":   {val: "う", skip: 1},
	"e":   {val: "え", skip: 1},
	"o":   {val: "お", skip: 1},
	"ka":  {val: "か", skip: 2},
	"ki":  {val: "き", skip: 2},
	"ku":  {val: "く", skip: 2},
	"ke":  {val: "け", skip: 2},
	"ko":  {val: "こ", skip: 2},
	"sa":  {val: "さ", skip: 2},
	"shi": {val: "し", skip: 3},
	"su":  {val: "す", skip: 2},
	"se":  {val: "せ", skip: 2},
	"so":  {val: "そ", skip: 2},
	"ta":  {val: "た", skip: 2},
	"chi": {val: "ち", skip: 3},
	"tsu": {val: "つ", skip: 3},
	"te":  {val: "て", skip: 2},
	"to":  {val: "と", skip: 2},
	"na":  {val: "な", skip: 2},
	"ni":  {val: "に", skip: 2},
	"nu":  {val: "ぬ", skip: 2},
	"ne":  {val: "ね", skip: 2},
	"no":  {val: "の", skip: 2},
	"ha":  {val: "は", skip: 2},
	"hi":  {val: "ひ", skip: 2},
	"hu":  {val: "ふ", skip: 2},
	"fu":  {val: "ふ", skip: 2},
	"he":  {val: "へ", skip: 2},
	"ho":  {val: "ほ", skip: 2},
	"ma":  {val: "ま", skip: 2},
	"mi":  {val: "み", skip: 2},
	"mu":  {val: "む", skip: 2},
	"me":  {val: "め", skip: 2},
	"mo":  {val: "も", skip: 2},
	"ya":  {val: "や", skip: 2},
	"yu":  {val: "ゆ", skip: 2},
	"yo":  {val: "よ", skip: 2},
	"ra":  {val: "ら", skip: 2},
	"ri":  {val: "り", skip: 2},
	"ru":  {val: "る", skip: 2},
	"re":  {val: "れ", skip: 2},
	"ro":  {val: "ろ", skip: 2},
	"wa":  {val: "わ", skip: 2},
	"wi":  {val: "ゐ", skip: 2},
	"wo":  {val: "を", skip: 2},
	"we":  {val: "ゑ", skip: 2},
	"n":   {val: "ん", skip: 1},

	"ga":  {val: "が", skip: 2},
	"gi":  {val: "ぎ", skip: 2},
	"gu":  {val: "ぐ", skip: 2},
	"ge":  {val: "げ", skip: 2},
	"go":  {val: "ご", skip: 2},
	"za":  {val: "ざ", skip: 2},
	"ji":  {val: "じ", skip: 2},
	"zu":  {val: "ず", skip: 2},
	"ze":  {val: "ぜ", skip: 2},
	"zo":  {val: "ぞ", skip: 2},
	"da":  {val: "だ", skip: 2},
	"di":  {val: "ぢ", skip: 2},
	"dzu": {val: "づ", skip: 3},
	"de":  {val: "で", skip: 2},
	"do":  {val: "ど", skip: 2},
	"ba":  {val: "ば", skip: 2},
	"bi":  {val: "び", skip: 2},
	"bu":  {val: "ぶ", skip: 2},
	"be":  {val: "べ", skip: 2},
	"bo":  {val: "ぼ", skip: 2},
	"pa":  {val: "ぱ", skip: 2},
	"pi":  {val: "ぴ", skip: 2},
	"pu":  {val: "ぷ", skip: 2},
	"pe":  {val: "ぺ", skip: 2},
	"po":  {val: "ぽ", skip: 2},

	"kya": {val: "きゃ", skip: 3},
	"kyu": {val: "きゅ", skip: 3},
	"kyo": {val: "きょ", skip: 3},
	"gya": {val: "ぎゃ", skip: 3},
	"gyu": {val: "ぎゅ", skip: 3},
	"gyo": {val: "ぎょ", skip: 3},
	"sha": {val: "しゃ", skip: 3},
	"shu": {val: "しゅ", skip: 3},
	"sho": {val: "しょ", skip: 3},
	"ja":  {val: "じゃ", skip: 2},
	"ju":  {val: "じゅ", skip: 2},
	"jo":  {val: "じょ", skip: 2},
	"cha": {val: "ちゃ", skip: 3},
	"cho": {val: "ちょ", skip: 3},
	"chu": {val: "ちゅ", skip: 3},
	"nya": {val: "にゃ", skip: 3},
	"nyu": {val: "にゅ", skip: 3},
	"nyo": {val: "にょ", skip: 3},
	"hya": {val: "ひゃ", skip: 3},
	"hyu": {val: "ひゅ", skip: 3},
	"hyo": {val: "ひょ", skip: 3},
	"bya": {val: "びゃ", skip: 3},
	"byu": {val: "びゅ", skip: 3},
	"byo": {val: "びょ", skip: 3},
	"pya": {val: "ぴゃ", skip: 3},
	"pyu": {val: "ぴゅ", skip: 3},
	"pyo": {val: "ぴょ", skip: 3},
	"mya": {val: "みゃ", skip: 3},
	"myu": {val: "みゅ", skip: 3},
	"myo": {val: "みょ", skip: 3},
	"rya": {val: "りゃ", skip: 3},
	"ryu": {val: "りゅ", skip: 3},
	"ryo": {val: "りょ", skip: 3},

	"t":    {val: "っ", skip: 1},
	"kka":  {val: "っか", skip: 3},
	"kki":  {val: "っき", skip: 3},
	"kku":  {val: "っく", skip: 3},
	"kke":  {val: "っけ", skip: 3},
	"kko":  {val: "っこ", skip: 3},
	"kkya": {val: "っきゃ", skip: 4},
	"kkyu": {val: "っきゅ", skip: 4},
	"kkyo": {val: "っきょ", skip: 4},
	"ssa":  {val: "っさ", skip: 3},
	"sshi": {val: "っし", skip: 4},
	"ssu":  {val: "っす", skip: 3},
	"sse":  {val: "っせ", skip: 3},
	"sso":  {val: "っそ", skip: 3},
	"ssha": {val: "っしゃ", skip: 4},
	"sshu": {val: "っしゅ", skip: 4},
	"ssho": {val: "っしょ", skip: 4},
	"tta":  {val: "った", skip: 3},
	"cchi": {val: "っち", skip: 4},
	"ttsu": {val: "っつ", skip: 4},
	"tte":  {val: "って", skip: 3},
	"tto":  {val: "っと", skip: 3},
	"ccha": {val: "っちゃ", skip: 4},
	"cchu": {val: "っちゅ", skip: 4},
	"ccho": {val: "っちょ", skip: 4},
	"ppa":  {val: "っぱ", skip: 3},
	"ppi":  {val: "っぴ", skip: 3},
	"ppu":  {val: "っぷ", skip: 3},
	"ppe":  {val: "っぺ", skip: 3},
	"ppo":  {val: "っぽ", skip: 3},
	"ppya": {val: "っぴゃ", skip: 4},
	"ppyu": {val: "っぴゅ", skip: 4},
	"ppyo": {val: "っぴょ", skip: 4},
}

var katakana = map[string]kana{

	"A":    {val: "ア", skip: 1},
	"AA":   {val: "アー", skip: 2},
	"I":    {val: "イ", skip: 1},
	"II":   {val: "イー", skip: 2},
	"U":    {val: "ウ", skip: 1},
	"UU":   {val: "ウー", skip: 2},
	"E":    {val: "エ", skip: 1},
	"EE":   {val: "エー", skip: 2},
	"EI":   {val: "エー", skip: 2},
	"O":    {val: "オ", skip: 1},
	"OO":   {val: "オー", skip: 2},
	"OU":   {val: "オー", skip: 2},
	"KA":   {val: "カ", skip: 2},
	"KAA":  {val: "カー", skip: 3},
	"KI":   {val: "キ", skip: 2},
	"KII":  {val: "キー", skip: 3},
	"KU":   {val: "ク", skip: 2},
	"KUU":  {val: "クー", skip: 3},
	"KE":   {val: "ケ", skip: 2},
	"KEE":  {val: "ケー", skip: 3},
	"KEI":  {val: "ケー", skip: 3},
	"KO":   {val: "コ", skip: 2},
	"KOO":  {val: "コー", skip: 3},
	"KOU":  {val: "コー", skip: 3},
	"SA":   {val: "サ", skip: 2},
	"SAA":  {val: "サー", skip: 3},
	"SHI":  {val: "シ", skip: 3},
	"SHII": {val: "シー", skip: 4},
	"SU":   {val: "ス", skip: 2},
	"SUU":  {val: "スー", skip: 3},
	"SE":   {val: "セ", skip: 2},
	"SEE":  {val: "セー", skip: 3},
	"SEI":  {val: "セー", skip: 3},
	"SO":   {val: "ソ", skip: 2},
	"SOO":  {val: "ソー", skip: 3},
	"SOU":  {val: "ソー", skip: 3},
	"TA":   {val: "タ", skip: 2},
	"TAA":  {val: "ター", skip: 3},
	"CHI":  {val: "チ", skip: 3},
	"CHII": {val: "チー", skip: 4},
	"TSU":  {val: "ツ", skip: 3},
	"TSUU": {val: "ツー", skip: 4},
	"TE":   {val: "テ", skip: 2},
	"TEE":  {val: "テー", skip: 3},
	"TEI":  {val: "テー", skip: 3},
	"TO":   {val: "ト", skip: 2},
	"TOO":  {val: "トー", skip: 3},
	"TOU":  {val: "トー", skip: 3},
	"NA":   {val: "ナ", skip: 2},
	"NAA":  {val: "ナー", skip: 3},
	"NI":   {val: "ニ", skip: 2},
	"NII":  {val: "ニー", skip: 3},
	"NU":   {val: "ヌ", skip: 2},
	"NUU":  {val: "ヌー", skip: 3},
	"NE":   {val: "ネ", skip: 2},
	"NEE":  {val: "ネー", skip: 3},
	"NEI":  {val: "ネー", skip: 3},
	"NO":   {val: "ノ", skip: 2},
	"NOO":  {val: "ノー", skip: 3},
	"NOU":  {val: "ノー", skip: 3},
	"HA":   {val: "ハ", skip: 2},
	"HAA":  {val: "ハー", skip: 3},
	"HI":   {val: "ヒ", skip: 2},
	"HII":  {val: "ヒー", skip: 3},
	"HU":   {val: "フ", skip: 2},
	"HUU":  {val: "フー", skip: 3},
	"FU":   {val: "フ", skip: 2},
	"FUU":  {val: "フー", skip: 3},
	"HE":   {val: "ヘ", skip: 2},
	"HEE":  {val: "ヘー", skip: 3},
	"HEI":  {val: "ヘー", skip: 3},
	"HO":   {val: "ホ", skip: 2},
	"HOO":  {val: "ホー", skip: 3},
	"HOU":  {val: "ホー", skip: 3},
	"MA":   {val: "マ", skip: 2},
	"MAA":  {val: "マー", skip: 3},
	"MI":   {val: "ミ", skip: 2},
	"MII":  {val: "ミー", skip: 3},
	"MU":   {val: "ム", skip: 2},
	"MUU":  {val: "ムー", skip: 3},
	"ME":   {val: "メ", skip: 2},
	"MEE":  {val: "メー", skip: 3},
	"MEI":  {val: "メー", skip: 3},
	"MO":   {val: "モ", skip: 2},
	"MOO":  {val: "モー", skip: 3},
	"MOU":  {val: "モー", skip: 3},
	"YA":   {val: "ヤ", skip: 2},
	"YAA":  {val: "ヤー", skip: 3},
	"YU":   {val: "ユ", skip: 2},
	"YUU":  {val: "ユー", skip: 3},
	"YO":   {val: "ヨ", skip: 2},
	"YOO":  {val: "ヨー", skip: 3},
	"YOU":  {val: "ヨー", skip: 3},
	"RA":   {val: "ラ", skip: 2},
	"RAA":  {val: "ラー", skip: 3},
	"RI":   {val: "リ", skip: 2},
	"RII":  {val: "リー", skip: 3},
	"RU":   {val: "ル", skip: 2},
	"RUU":  {val: "ルー", skip: 3},
	"RE":   {val: "レ", skip: 2},
	"REE":  {val: "レー", skip: 3},
	"REI":  {val: "レー", skip: 3},
	"RO":   {val: "ロ", skip: 2},
	"ROO":  {val: "ロー", skip: 3},
	"ROU":  {val: "ロー", skip: 3},
	"WA":   {val: "ワ", skip: 2},
	"WAA":  {val: "ワー", skip: 3},
	"WI":   {val: "ヰ", skip: 2},
	"WII":  {val: "ヰー", skip: 3},
	"WO":   {val: "ヲ", skip: 2},
	"WOO":  {val: "ヲー", skip: 3},
	"WOU":  {val: "ヲー", skip: 3},
	"WE":   {val: "ヱ", skip: 2},
	"WEE":  {val: "ヱー", skip: 3},
	"WEI":  {val: "ヱー", skip: 3},
	"N":    {val: "ン", skip: 1},

	"GA":   {val: "ガ", skip: 2},
	"GAA":  {val: "ガー", skip: 3},
	"GI":   {val: "ギ", skip: 2},
	"GII":  {val: "ギー", skip: 3},
	"GU":   {val: "グ", skip: 2},
	"GUU":  {val: "グー", skip: 3},
	"GE":   {val: "ゲ", skip: 2},
	"GEE":  {val: "ゲー", skip: 3},
	"GEI":  {val: "ゲー", skip: 3},
	"GO":   {val: "ゴ", skip: 2},
	"GOO":  {val: "ゴー", skip: 3},
	"GOU":  {val: "ゴー", skip: 3},
	"ZA":   {val: "ザ", skip: 2},
	"ZAA":  {val: "ザー", skip: 3},
	"JI":   {val: "ジ", skip: 2},
	"JII":  {val: "ジー", skip: 3},
	"ZU":   {val: "ズ", skip: 2},
	"ZUU":  {val: "ズー", skip: 3},
	"ZE":   {val: "ゼ", skip: 2},
	"ZEE":  {val: "ゼー", skip: 3},
	"ZEI":  {val: "ゼー", skip: 3},
	"ZO":   {val: "ゾ", skip: 2},
	"ZOO":  {val: "ゾー", skip: 3},
	"ZOU":  {val: "ゾー", skip: 3},
	"DA":   {val: "ダ", skip: 2},
	"DAA":  {val: "ダー", skip: 3},
	"DI":   {val: "ヂ", skip: 2},
	"DII":  {val: "ヂー", skip: 3},
	"DZU":  {val: "ヅ", skip: 3},
	"DZUU": {val: "ヅー", skip: 4},
	"DE":   {val: "デ", skip: 2},
	"DEE":  {val: "デー", skip: 3},
	"DEI":  {val: "デー", skip: 3},
	"DO":   {val: "ド", skip: 2},
	"DOO":  {val: "ドー", skip: 3},
	"DOU":  {val: "ドー", skip: 3},
	"BA":   {val: "バ", skip: 2},
	"BAA":  {val: "バー", skip: 3},
	"BI":   {val: "ビ", skip: 2},
	"BII":  {val: "ビー", skip: 3},
	"BU":   {val: "ブ", skip: 2},
	"BUU":  {val: "ブー", skip: 3},
	"BE":   {val: "ベ", skip: 2},
	"BEE":  {val: "ベー", skip: 3},
	"BEI":  {val: "ベー", skip: 3},
	"BO":   {val: "ボ", skip: 2},
	"BOO":  {val: "ボー", skip: 3},
	"BOU":  {val: "ボー", skip: 3},
	"PA":   {val: "パ", skip: 2},
	"PAA":  {val: "パー", skip: 3},
	"PI":   {val: "ピ", skip: 2},
	"PII":  {val: "ピー", skip: 3},
	"PU":   {val: "プ", skip: 2},
	"PUU":  {val: "プー", skip: 3},
	"PE":   {val: "ペ", skip: 2},
	"PEE":  {val: "ペー", skip: 3},
	"PEI":  {val: "ペー", skip: 3},
	"PO":   {val: "ポ", skip: 2},
	"POO":  {val: "ポー", skip: 3},
	"POU":  {val: "ポー", skip: 3},

	"KYA":  {val: "キャ", skip: 3},
	"KYAA": {val: "キャー", skip: 4},
	"KYU":  {val: "キュ", skip: 3},
	"KYUU": {val: "キュー", skip: 4},
	"KYO":  {val: "キョ", skip: 3},
	"KYOO": {val: "キョー", skip: 4},
	"KYOU": {val: "キョー", skip: 4},
	"GYA":  {val: "ギャ", skip: 3},
	"GYAA": {val: "ギャー", skip: 4},
	"GYU":  {val: "ギュ", skip: 3},
	"GYUU": {val: "ギュー", skip: 4},
	"GYO":  {val: "ギョ", skip: 3},
	"GYOO": {val: "ギョー", skip: 4},
	"GYOU": {val: "ギョー", skip: 4},
	"SHA":  {val: "シャ", skip: 3},
	"SHAA": {val: "シャー", skip: 4},
	"SHU":  {val: "シュ", skip: 3},
	"SHUU": {val: "シュー", skip: 4},
	"SHO":  {val: "ショ", skip: 3},
	"SHOO": {val: "ショー", skip: 4},
	"SHOU": {val: "ショー", skip: 4},
	"JA":   {val: "ジャ", skip: 2},
	"JAA":  {val: "ジャー", skip: 3},
	"JU":   {val: "ジュ", skip: 2},
	"JUU":  {val: "ジュー", skip: 3},
	"JO":   {val: "ジョ", skip: 2},
	"JOO":  {val: "ジョー", skip: 3},
	"JOU":  {val: "ジョー", skip: 3},
	"CHA":  {val: "チャ", skip: 3},
	"CHAA": {val: "チャー", skip: 4},
	"CHO":  {val: "チョ", skip: 3},
	"CHOO": {val: "チョー", skip: 4},
	"CHOU": {val: "チョー", skip: 4},
	"CHU":  {val: "チュ", skip: 3},
	"CHUU": {val: "チュー", skip: 4},
	"NYA":  {val: "ニャ", skip: 3},
	"NYAA": {val: "ニャー", skip: 4},
	"NYU":  {val: "ニュ", skip: 3},
	"NYUU": {val: "ニュー", skip: 4},
	"NYO":  {val: "ニョ", skip: 3},
	"NYOO": {val: "ニョー", skip: 4},
	"NYOU": {val: "ニョー", skip: 4},
	"HYA":  {val: "ヒャ", skip: 3},
	"HYAA": {val: "ヒャー", skip: 4},
	"HYU":  {val: "ヒュ", skip: 3},
	"HYUU": {val: "ヒュー", skip: 4},
	"HYO":  {val: "ヒョ", skip: 3},
	"HYOO": {val: "ヒョー", skip: 4},
	"HYOU": {val: "ヒョー", skip: 4},
	"BYA":  {val: "ビャ", skip: 3},
	"BYAA": {val: "ビャー", skip: 4},
	"BYU":  {val: "ビュ", skip: 3},
	"BYUU": {val: "ビュー", skip: 4},
	"BYO":  {val: "ビョ", skip: 3},
	"BYOO": {val: "ビョー", skip: 4},
	"BYOU": {val: "ビョー", skip: 4},
	"PYA":  {val: "ピャ", skip: 3},
	"PYAA": {val: "ピャー", skip: 4},
	"PYU":  {val: "ピュ", skip: 3},
	"PYUU": {val: "ピュー", skip: 4},
	"PYO":  {val: "ピョ", skip: 3},
	"PYOO": {val: "ピョー", skip: 4},
	"PYOU": {val: "ピョー", skip: 4},
	"MYA":  {val: "ミャ", skip: 3},
	"MYAA": {val: "ミャー", skip: 4},
	"MYU":  {val: "ミュ", skip: 3},
	"MYUU": {val: "ミュー", skip: 4},
	"MYO":  {val: "ミョ", skip: 3},
	"MYOO": {val: "ミョー", skip: 4},
	"MYOU": {val: "ミョー", skip: 4},
	"RYA":  {val: "リャ", skip: 3},
	"RYAA": {val: "リャー", skip: 4},
	"RYU":  {val: "リュ", skip: 3},
	"RYUU": {val: "リュー", skip: 4},
	"RYO":  {val: "リョ", skip: 3},
	"RYOO": {val: "リョー", skip: 4},
	"RYOU": {val: "リョー", skip: 4},

	"T":     {val: "ッ", skip: 1},
	"KKA":   {val: "ッカ", skip: 3},
	"KKAA":  {val: "ッカー", skip: 4},
	"KKI":   {val: "ッキ", skip: 3},
	"KKII":  {val: "ッキー", skip: 4},
	"KKU":   {val: "ック", skip: 3},
	"KKUU":  {val: "ックー", skip: 4},
	"KKE":   {val: "ッケ", skip: 3},
	"KKEE":  {val: "ッケー", skip: 4},
	"KKEI":  {val: "ッケー", skip: 4},
	"KKO":   {val: "ッコ", skip: 3},
	"KKOO":  {val: "ッコー", skip: 4},
	"KKOU":  {val: "ッコー", skip: 4},
	"KKYA":  {val: "ッキャ", skip: 4},
	"KKYAA": {val: "ッキャー", skip: 5},
	"KKYU":  {val: "ッキュ", skip: 4},
	"KKYUU": {val: "ッキュー", skip: 5},
	"KKYO":  {val: "ッキョ", skip: 4},
	"KKYOO": {val: "ッキョー", skip: 5},
	"KKYOU": {val: "ッキョー", skip: 5},
	"SSA":   {val: "ッサ", skip: 3},
	"SSAA":  {val: "ッサー", skip: 4},
	"SSHI":  {val: "ッシ", skip: 4},
	"SSHII": {val: "ッシー", skip: 5},
	"SSU":   {val: "ッス", skip: 3},
	"SSUU":  {val: "ッスー", skip: 4},
	"SSE":   {val: "ッセ", skip: 3},
	"SSEE":  {val: "ッセー", skip: 4},
	"SSEI":  {val: "ッセー", skip: 4},
	"SSO":   {val: "ッソ", skip: 3},
	"SSOO":  {val: "ッソー", skip: 4},
	"SSOU":  {val: "ッソー", skip: 4},
	"SSHA":  {val: "ッシャ", skip: 4},
	"SSHAA": {val: "ッシャー", skip: 5},
	"SSHU":  {val: "ッシュ", skip: 4},
	"SSHUU": {val: "ッシュー", skip: 5},
	"SSHO":  {val: "ッショ", skip: 4},
	"SSHOO": {val: "ッショー", skip: 5},
	"SSHOU": {val: "ッショー", skip: 5},
	"TTA":   {val: "ッタ", skip: 3},
	"TTAA":  {val: "ッター", skip: 4},
	"CCHI":  {val: "ッチ", skip: 4},
	"CCHII": {val: "ッチー", skip: 5},
	"TTSU":  {val: "ッツ", skip: 4},
	"TTSUU": {val: "ッツー", skip: 5},
	"TTE":   {val: "ッテ", skip: 3},
	"TTEE":  {val: "ッテー", skip: 4},
	"TTEI":  {val: "ッテー", skip: 4},
	"TTO":   {val: "ット", skip: 3},
	"TTOO":  {val: "ットー", skip: 4},
	"TTOU":  {val: "ットー", skip: 4},
	"CCHA":  {val: "ッチャ", skip: 4},
	"CCHAA": {val: "ッチャー", skip: 5},
	"CCHU":  {val: "ッチュ", skip: 4},
	"CCHUU": {val: "ッチュー", skip: 5},
	"CCHO":  {val: "ッチョ", skip: 4},
	"CCHOO": {val: "ッチョー", skip: 5},
	"CCHOU": {val: "ッチョー", skip: 5},
	"PPA":   {val: "ッパ", skip: 3},
	"PPAA":  {val: "ッパー", skip: 4},
	"PPI":   {val: "ッピ", skip: 3},
	"PPII":  {val: "ッピー", skip: 4},
	"PPU":   {val: "ップ", skip: 3},
	"PPUU":  {val: "ップー", skip: 4},
	"PPE":   {val: "ッペ", skip: 3},
	"PPEE":  {val: "ッペー", skip: 4},
	"PPEI":  {val: "ッペー", skip: 4},
	"PPO":   {val: "ッポ", skip: 3},
	"PPOO":  {val: "ッポー", skip: 4},
	"PPOU":  {val: "ッポー", skip: 4},
	"PPYA":  {val: "ッピャ", skip: 4},
	"PPYAA": {val: "ッピャー", skip: 5},
	"PPYU":  {val: "ッピュ", skip: 4},
	"PPYUU": {val: "ッピュー", skip: 5},
	"PPYO":  {val: "ッピョ", skip: 4},
	"PPYOO": {val: "ッピョー", skip: 5},
	"PPYOU": {val: "ッピョー", skip: 5},
}