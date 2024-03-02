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
	baseURL := "https://hunt.kamille.ovh/audio/"
	distance := sound.Distance

	switch GunIndex[sound.Name] {
	case "Berthier Mle 1892":
		return baseURL + "Berthier Mle 1892-" + distance2mp3(distance)
	case "Bornheim No. 3":
		return baseURL + "Bornheim No. 3-" + distance2mp3(distance)
	case "Bow":
		return baseURL + "Bow-" + distance2mp3(distance)
	case "Caldwell 92 New army":
		return baseURL + "Caldwell 92 New army-" + distance2mp3(distance)
	case "Caldwell Conversion Pistol":
		return baseURL + "Caldwell Conversion Pistol-" + distance2mp3(distance)
	case "Caldwell Conversion Uppercut":
		return baseURL + "Caldwell Conversion Uppercut-" + distance2mp3(distance)
	case "Caldwell Pax":
		return baseURL + "Caldwell Pax-" + distance2mp3(distance)
	case "Caldwell Rival 78":
		return baseURL + "Caldwell Rival 78-" + distance2mp3(distance)
	case "Crossbow":
		return baseURL + "Crossbow-" + distance2mp3(distance)
	case "Crown & King Auto-5":
		return baseURL + "Crown & King Auto-5-" + distance2mp3(distance)
	case "Dolch 96":
		return baseURL + "Dolch 96-" + distance2mp3(distance)
	case "Flare Pistol":
		return baseURL + "Flare Pistol-" + distance2mp3(distance)
	case "LeMat Mark II":
		return baseURL + "LeMat Mark II-" + distance2mp3(distance)
	case "Lebel 1886":
		return baseURL + "Lebel 1886-" + distance2mp3(distance)
	case "Martini-Henry IC1":
		return baseURL + "Martini-Henry IC1-" + distance2mp3(distance)
	case "Mosin-Nagant M1891":
		return baseURL + "Mosin-Nagant M1891-" + distance2mp3(distance)
	case "Mosin-Nagant M1891 Avtomat":
		return baseURL + "Mosin-Nagant M1891 Avtomat-" + distance2mp3(distance)
	case "Nagant M1895 Officer":
		return baseURL + "Nagant M1895 Officer-" + distance2mp3(distance)
	case "Nagant M1895 Silencer":
		return baseURL + "Nagant M1895 Silencer-" + distance2mp3(distance)
	case "Nitro Express Rifle":
		return baseURL + "Nitro Express Rifle-" + distance2mp3(distance)
	case "Quad Derringer":
		return baseURL + "Quad Derringer-" + distance2mp3(distance)
	case "Romero 77":
		return baseURL + "Romero 77-" + distance2mp3(distance)
	case "Scottfied Model 3":
		return baseURL + "Scottfied Model 3-" + distance2mp3(distance)
	case "Sparks LRR":
		return baseURL + "Sparks LRR-" + distance2mp3(distance)
	case "Sparks LRR Silencer":
		return baseURL + "Sparks LRR Silencer-" + distance2mp3(distance)
	case "Sparks Pistol":
		return baseURL + "Sparks Pistol-" + distance2mp3(distance)
	case "Specter 1882":
		return baseURL + "Specter 1882-" + distance2mp3(distance)
	case "Springfield 1866":
		return baseURL + "Springfield 1866-" + distance2mp3(distance)
	case "Vetterli 71 Karabiner":
		return baseURL + "Vetterli 71 Karabiner-" + distance2mp3(distance)
	case "Vetterli 71 Karabiner Silencer":
		return baseURL + "Vetterli 71 Karabiner Silencer-" + distance2mp3(distance)
	case "Windfield 1893 Slate":
		return baseURL + "Windfield 1893 Slate-" + distance2mp3(distance)
	case "Winfield 1887 Terminus":
		return baseURL + "Winfield 1887 Terminus-" + distance2mp3(distance)
	case "Winfield M1873C":
		return baseURL + "Winfield M1873C-" + distance2mp3(distance)
	case "Winfield M1873C Silencer":
		return baseURL + "Winfield M1873C Silencer-" + distance2mp3(distance)
	case "Winfield M1876 Centennial":
		return baseURL + "Winfield M1876 Centennial-" + distance2mp3(distance)
	default:
		return baseURL + "Lebel 1886-" + "06.mp3"
	}
}

func distance2mp3(distance string) string {
	// 生成距离映射
	switch distance {
	case "1000":
		return "01.mp3"
	case "500":
		return "02,mp3"
	case "350":
		return "03.mp3"
	case "200":
		return "04.mp3"
	case "100":
		return "05.mp3"
	case "50":
		return "06.mp3"
	case "20":
		return "07.mp3"
	case "5":
		return "08.mp3"
	case "0":
		return "09.mp3"
	default:
		return "06.mp3"
	}
}

// 获取枪目录
func GetIndex() string {
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
