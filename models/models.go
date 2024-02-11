package models

/*
========================
Public structs
========================
*/
type LocationResponse struct {
	Name  string  `json:"name"`
	Local local   `json:"local"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

type WeatherResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []list `json:"list"`
	City    city   `json:"city"`
}

/*
========================
Private structs
========================
*/
type local struct {
	Ascii string `json:"ascii"`
}

type list struct {
	Dt      int       `json:"dt"`
	Main    main      `json:"main"`
	Weather []weather `json:"weather"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}
type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

type city struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Coord    coord  `json:"coord"`
	Country  string `json:"country"`
	Timezone int    `json:"timezone"`
	Sunrise  int    `json:"sunrise"`
	Sunset   int    `json:"sunset"`
}
type coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
