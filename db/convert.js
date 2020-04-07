import district3 from "./district3.js";

function shortProvince(province) {
	if (province.endsWith("市") || province.endsWith("区")) {
		return province.substr(0, 2);
	}
	if (province.endsWith("省")) {
		return province.substr(0, province.length - 1);
	}
	return province;
}

function shortCity(city) {
	if (city.endsWith("市") || city.endsWith("盟")) {
		return city.substr(0, city.length - 1);
	}
	return city;
}


const citiesInProvince = {};
const districtInCity = {};
district3.forEach(province => {
	if (["北京市", "上海市", "天津市", "重庆市", "香港特别行政区", "澳门特别行政区"]
		.indexOf(province.name) !== -1) {
		citiesInProvince[province.name] = [province.name];
		districtInCity[province.name] = province.children.map(({ name }) => name);
	} else {
		citiesInProvince[province.name] = province.children.map(city => {
			districtInCity[city.name] = city.children && city.children.length ?
				city.children.map(({ name }) => name) : [city.name];
			return city.name;
		});
	}
});

console.log("package db\n");
console.log("var Provinces = map[string][]string {");
Object.keys(citiesInProvince).forEach(province => {
	console.log(`\t"${province}": {`);
	citiesInProvince[province].forEach(city => {
		console.log(`\t\t"${city}",`);
	});
	console.log("\t},");
});
console.log("}");
console.log("var Cities = map[string][]string {");
Object.keys(districtInCity).forEach(city => {
	console.log(`\t"${city}": {`);
	districtInCity[city].forEach(district => {
		console.log(`\t\t"${district}",`);
	});
	console.log("\t},");
});
console.log("}");
console.log("var ShortProvinces map[string][]string");
console.log("var ShortCities map[string][]string");
