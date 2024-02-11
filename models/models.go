package models

type Location struct {
	Name  string `json:"name"`
	Local struct {
		Ascii string `json:"ascii"`
	} `json:"local_names"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Weather struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
	City struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country  string `json:"country"`
		Timezone int    `json:"timezone"`
		Sunrise  int    `json:"sunrise"`
		Sunset   int    `json:"sunset"`
	} `json:"city"`
}
