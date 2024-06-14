package utils

import (
	"strings"
)

type HuntSound struct {
	Name     string
	Distance string
}

var GunIndex = map[string]string{
	"Berthier Mle 1892":         "Berthier Mle 1892",
	"马枪":                        "Berthier Mle 1892",
	"1892":                      "Berthier Mle 1892",
	"Bornheim No. 3":            "Bornheim No. 3",
	"三号手枪":                      "Bornheim No. 3",
	"Bow":                       "Bow",
	"弓":                         "Bow",
	"caldwell 92 new army":      "Caldwell 92 New army",
	"92新军左":                     "Caldwell 92 New army",
	"新军左":                       "Caldwell 92 New army",
	"caldwell pax":              "Caldwell Pax",
	"Pax手枪":                     "Caldwell Pax",
	"pax":                       "Caldwell Pax",
	"caldwell rival 78":         "Caldwell Rival 78",
	"78喷":                       "Caldwell Rival 78",
	"Crossbow":                  "Crossbow",
	"弩":                         "Crossbow",
	"十字弩":                       "Crossbow",
	"Crown & King Auto-5":       "Crown & King Auto-5",
	"皇冠喷":                       "Crown & King Auto-5",
	"Dolch 96":                  "Dolch 96",
	"多尔西":                       "Dolch 96",
	"Flare Pistol":              "Flare Pistol",
	"信号枪":                       "Flare Pistol",
	"LeMat Mark II":             "LeMat Mark II",
	"马克2":                       "LeMat Mark II",
	"Lebel 1886":                "Lebel 1886",
	"乐贝":                        "Lebel 1886",
	"Martini-Henry IC1":         "Martini-Henry IC1",
	"马提尼":                       "Martini-Henry IC1",
	"Mosin-Nagant M1891":        "Mosin-Nagant M1891",
	"莫辛":                        "Mosin-Nagant M1891",
	"莫辛纳甘":                      "Mosin-Nagant M1891",
	"Nagant M1895 Officer":      "Nagant M1895 Officer",
	"officer左轮":                 "Nagant M1895 Officer",
	"公务员手枪":                     "Nagant M1895 Officer",
	"Nagant M1895 Silencer":     "Nagant M1895 Silencer",
	"纳甘消音左轮":                    "Nagant M1895 Silencer",
	"Nagant消音":                  "Nagant M1895 Silencer",
	"nitro express":             "Nitro Express Rifle",
	"猎象":                        "Nitro Express Rifle",
	"Quad Derringer":            "Quad Derringer",
	"四管袖珍手枪":                    "Quad Derringer",
	"Romero 77":                 "Romero 77",
	"77喷":                       "Romero 77",
	"Scottfied Model 3":         "Scottfied Model 3",
	"3号手枪":                      "Scottfied Model 3",
	"Sparks LRR":                "Sparks LRR",
	"LRR":                       "Sparks LRR",
	"乐融融":                       "Sparks LRR",
	"Sparks LRR Silencer":       "Sparks LRR Silencer",
	"消音LRR":                     "Sparks LRR Silencer",
	"Sparks Pistol":             "Sparks Pistol",
	"LRR手枪":                     "Sparks Pistol",
	"Specter 1882":              "Specter 1882",
	"幽灵喷":                       "Specter 1882",
	"Springfield 1866":          "Springfield 1866",
	"春田":                        "Springfield 1866",
	"Vetterli 71 Karabiner":     "Vetterli 71 Karabiner",
	"维特利":                       "Vetterli 71 Karabiner",
	"Windfield 1893 Slate":      "Windfield 1893 Slate",
	"石板喷":                       "Windfield 1893 Slate",
	"Winfield 1887 Terminus":    "Winfield 1887 Terminus",
	"杠杆喷":                       "Winfield 1887 Terminus",
	"Winfield M1873C":           "Winfield M1873C",
	"温菲":                        "Winfield M1873C",
	"Winfield M1873C Silencer":  "Winfield M1873C Silencer",
	"消音温菲":                      "Winfield M1873C Silencer",
	"Winfield M1876 Centennial": "Winfield M1876 Centennial",
	"温菲1876":                    "Winfield M1876 Centennial",
	"温菲百年庆典":                    "Winfield M1876 Centennial",
}

func GetHuntSound(sound HuntSound) string {
	baseURL := "http://hunt.kamille.ovh/audio/"
	distance := sound.Distance

	switch GunIndex[sound.Name] {
	case "Berthier Mle 1892":
		return baseURL + "Berthier%20Mle%201892-" + distance2mp3(distance)
	case "Bornheim No. 3":
		return baseURL + "Bornheim%20No.%203-" + distance2mp3(distance)
	case "Bow":
		return baseURL + "Bow-" + distance2mp3(distance)
	case "Caldwell 92 New army":
		return baseURL + "Caldwell%2092%20New%20army-" + distance2mp3(distance)
	case "Caldwell Conversion Pistol":
		return baseURL + "Caldwell%20Conversion%20Pistol-" + distance2mp3(distance)
	case "Caldwell Conversion Uppercut":
		return baseURL + "Caldwell%20Conversion%20Uppercut-" + distance2mp3(distance)
	case "Caldwell Pax":
		return baseURL + "Caldwell%20Pax-" + distance2mp3(distance)
	case "Caldwell Rival 78":
		return baseURL + "Caldwell%20Rival%2078-" + distance2mp3(distance)
	case "Crossbow":
		return baseURL + "Crossbow-" + distance2mp3(distance)
	case "Crown & King Auto-5":
		return baseURL + "Crown%20&%20King%20Auto-5-" + distance2mp3(distance)
	case "Dolch 96":
		return baseURL + "Dolch%2096-" + distance2mp3(distance)
	case "Flare Pistol":
		return baseURL + "Flare%20Pistol-" + distance2mp3(distance)
	case "LeMat Mark II":
		return baseURL + "LeMat%20Mark%20II-" + distance2mp3(distance)
	case "Lebel 1886":
		return baseURL + "Lebel%201886-" + distance2mp3(distance)
	case "Martini-Henry IC1":
		return baseURL + "Martini-Henry%20IC1-" + distance2mp3(distance)
	case "Mosin-Nagant M1891":
		return baseURL + "Mosin-Nagant%20M1891-" + distance2mp3(distance)
	case "Mosin-Nagant M1891 Avtomat":
		return baseURL + "Mosin-Nagant%20M1891%20Avtomat-" + distance2mp3(distance)
	case "Nagant M1895 Officer":
		return baseURL + "Nagant%20M1895%20Officer-" + distance2mp3(distance)
	case "Nagant M1895 Silencer":
		return baseURL + "Nagant%20M1895%20Silencer-" + distance2mp3(distance)
	case "Nitro Express Rifle":
		return baseURL + "Nitro%20Express%20Rifle-" + distance2mp3(distance)
	case "Quad Derringer":
		return baseURL + "Quad%20Derringer-" + distance2mp3(distance)
	case "Romero 77":
		return baseURL + "Romero%2077-" + distance2mp3(distance)
	case "Scottfied Model 3":
		return baseURL + "Scottfied%20Model 3-" + distance2mp3(distance)
	case "Sparks LRR":
		return baseURL + "Sparks%20LRR-" + distance2mp3(distance)
	case "Sparks LRR Silencer":
		return baseURL + "Sparks%20LRR%20Silencer-" + distance2mp3(distance)
	case "Sparks Pistol":
		return baseURL + "Sparks%20Pistol-" + distance2mp3(distance)
	case "Specter 1882":
		return baseURL + "Specter%201882-" + distance2mp3(distance)
	case "Springfield 1866":
		return baseURL + "Springfield%201866-" + distance2mp3(distance)
	case "Vetterli 71 Karabiner":
		return baseURL + "Vetterli%2071%20Karabiner-" + distance2mp3(distance)
	case "Vetterli 71 Karabiner Silencer":
		return baseURL + "Vetterli%2071%20Karabiner%20Silencer-" + distance2mp3(distance)
	case "Windfield 1893 Slate":
		return baseURL + "Windfield%201893%20Slate-" + distance2mp3(distance)
	case "Winfield 1887 Terminus":
		return baseURL + "Winfield%201887%20Terminus-" + distance2mp3(distance)
	case "Winfield M1873C":
		return baseURL + "Winfield%20M1873C-" + distance2mp3(distance)
	case "Winfield M1873C Silencer":
		return baseURL + "Winfield%20M1873C%20Silencer-" + distance2mp3(distance)
	case "Winfield M1876 Centennial":
		return baseURL + "Winfield%20M1876%20Centennial-" + distance2mp3(distance)
	default:
		return baseURL + "Lebel%201886-" + "06.mp3"
	}
}

func distance2mp3(distance string) string {
	// 生成距离映射
	switch distance {
	case "1000m":
		return "01.mp3"
	case "500m":
		return "02,mp3"
	case "350m":
		return "03.mp3"
	case "200m":
		return "04.mp3"
	case "100m":
		return "05.mp3"
	case "50m":
		return "06.mp3"
	case "20m":
		return "07.mp3"
	case "5m":
		return "08.mp3"
	case "0m":
		return "09.mp3"
	default:
		return "06.mp3"
	}
}

// 获取枪目录
func GetGunIndex() string {
	uniqueMap := make(map[string]bool)
	uniqueValues := []string{}

	for _, value := range GunIndex {
		if !uniqueMap[value] {
			uniqueMap[value] = true
			uniqueValues = append(uniqueValues, value)
		}
	}
	result := strings.Join(uniqueValues, "\n")
	return result
}
