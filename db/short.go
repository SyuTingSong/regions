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
	ShortProvinces = make([]string, len(Provinces))
	for i := range Provinces {
		ShortProvinces[i] = shortProvince(Provinces[i])
	}
	ShortCitiesInProvinces = make(map[string][]string)
	for province, cities := range CitiesInProvinces {
		shortCities := make([]string, len(cities))
		for i := range cities {
			shortCities[i] = shortCity(cities[i])
		}
		ShortCitiesInProvinces[shortProvince(province)] = shortCities
	}

	ShortDistrictInCities = make(map[string][]string)
	for city, districts := range DistrictsInCities {
		ShortDistrictInCities[shortCity(city)] = districts
	}
}
