import district3 from "./district3.js";

const provinces = [];
const citiesInProvince = {};
const districtInCity = {};
district3.forEach(province => {
	provinces.push(province.name);
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
console.log("var Provinces = []string{");
provinces.forEach(province => {
	console.log(`\t"${province}",`);
});
console.log("}");
console.log("var CitiesInProvinces = map[string][]string {");
Object.keys(citiesInProvince).forEach(province => {
	console.log(`\t"${province}": {`);
	citiesInProvince[province].forEach(city => {
		console.log(`\t\t"${city}",`);
	});
	console.log("\t},");
});
console.log("}");
console.log("var DistrictsInCities = map[string][]string {");
Object.keys(districtInCity).forEach(city => {
	console.log(`\t"${city}": {`);
	districtInCity[city].forEach(district => {
		console.log(`\t\t"${district}",`);
	});
	console.log("\t},");
});
console.log("}");

console.log("var ShortProvinces []string");
console.log("var ShortCitiesInProvinces map[string][]string");
console.log("var ShortDistrictInCities map[string][]string");
