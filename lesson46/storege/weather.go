package postgres

import (
	"database/sql"
	pb "my_mod/generated/wheather"
)

type Weatherrepo struct {
	Db *sql.DB
}

func NewWeatherRepo(db *sql.DB) *Weatherrepo {
	return &Weatherrepo{Db: db}
}

func (w *Weatherrepo) GetWeather(p *pb.CurrentWeatherRequest) (*pb.CurrentWeatherRespons, error) {

	var res pb.CurrentWeatherRespons

	err := w.Db.QueryRow("select temperature,humidity,wind from weather where lacation=$1", p.Location).Scan(&res.Temprature, &res.Humidity, &res.Wind)

	return &res, err
}

func (w *Weatherrepo) WeatherForecast(p *pb.WeatherForecastRequest) (*pb.WeatherForecastRespons, error) {

	rows, err := w.Db.Query("select date,temprature,condition from weather where locations=$1 limit $2", p.Location, p.Day)

	if err != nil {
		return nil, err
	}

	var forecest pb.WeatherForecast
	var resp []*pb.WeatherForecast
	for rows.Next() {
		err = rows.Scan(&forecest.Date, &forecest.Temperature, &forecest.Condition)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &forecest)
	}
	return &pb.WeatherForecastRespons{Forecasts: resp}, nil
}

func (w *Weatherrepo) IsTraffic(p *pb.IsTrafficRequest) (*pb.IsTraficRespons, error) {
	var resp pb.IsTraficRespons
	err := w.Db.QueryRow(`
		select 
			condation 
		from 
			weather 
		where 
			location=$1 and date=$2
	`, p.Location, p.Date).Scan(&resp)
	return &resp, err
}
