package postgres

import (
	"database/sql"
	pd "my_mod/generated/bus"
)

type Bussrepo struct {
	Db *sql.DB
}

func NewBussRepo(db *sql.DB) *Bussrepo {
	return &Bussrepo{Db: db}
}

func (b *Bussrepo) BusSchedule(p *pd.BusRequest) (*pd.BusScheduleResponse, error) {
	rows, err := b.Db.Query("select b.bekat from  transport as t join bekat as b on t.TRANSPORT_NUMBER=b.bus_number where  t.TRANSPORT_NUMBER=$1", p.BusNumber)
	if err != nil {
		return nil, err
	}
	var bekat string
	resp := make([]string, 0)

	for rows.Next() {
		err = rows.Scan(&bekat)

		if err != nil {
			return nil, err
		}
		resp = append(resp, bekat)
	}
	return &pd.BusScheduleResponse{Schule: resp}, nil
}

func (b *Bussrepo) TraBusLocations(p *pd.BusRequest) (*pd.BusLocationResponse, error) {
	var resc pd.BusLocationResponse

	err := b.Db.QueryRow(`
		select 
			transport_number,
			location 
			from 
				transport 
			where 
				transport_number=$1
	`, p.BusNumber).Scan(&resc.BusNumber, &resc.Locatsion)
	return &resc, err

}

func (b *Bussrepo) GetReportTrafficJam(p *pd.TrafficJamReportRequest) (*pd.TrafficJamReportResponse, error) {

	var resp pd.TrafficJamReportResponse

	err := b.Db.QueryRow(`
		select 
			condation 
		from transport 
		where 
			location=$1 and discription=$2
		`, p.Locatsion, p.Discription).Scan(&resp.Condition)
	return &resp, err
}
