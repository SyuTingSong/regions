package db

func shortProvince(province string) string {
	var runes = []rune(province)
	var l = len(runes)
	if runes[l-1] == '省' {
		return string(runes[0 : l-1])
	}
	if runes[l-1] == '市' || runes[l-1] == '区' {
		return string(runes[0:2])
	}
	return province
}

func shortCity(city string) string {
	var runes = []rune(city)
	var l = len(runes)
	if runes[l-1] == '市' || runes[l-1] == '盟' {
		return string(runes[0 : l-1])
	}
	return city
}

func InitShort() {
	if ShortProvinces != nil {
		return
	}
	ShortProvinces = make(map[string][]string)
	for province, cities := range Provinces {
		shortCities := make([]string, len(cities))
		for i := range cities {
			shortCities[i] = shortCity(cities[i])
		}
		ShortProvinces[shortProvince(province)] = shortCities
	}

	ShortCities = make(map[string][]string)
	for city, districts := range Cities {
		ShortCities[shortCity(city)] = districts
	}
}
