package province

import (
	"database/sql"
	"os"
)

type Repository struct {
}

func (r *Repository) FindAllProvinces() ([]Province, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM provinces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var name string
	var nameEN string
	var slug string
	var code int
	var regionID int
	var lat float32
	var lng float32

	var provinces []Province
	for rows.Next() {
		err = rows.Scan(&id, &name, &nameEN, &slug, &code, &regionID, &lat, &lng)
		if err != nil {
			return nil, err
		}

		p := Province{id, name, nameEN, slug, code, regionID, lat, lng}
		provinces = append(provinces, p)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *Repository) FindProvince(pid int64) (Province, error) {
	db, err := getDB()
	if err != nil {
		return Province{}, err
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM provinces WHERE id = ?", pid)
	if err != nil {
		return Province{}, err
	}
	defer row.Close()

	var id int
	var name string
	var nameEN string
	var slug string
	var code int
	var regionID int
	var lat float32
	var lng float32

	row.Next()
	err = row.Err()
	if err != nil {
		return Province{}, err
	}
	err = row.Scan(&id, &name, &nameEN, &slug, &code, &regionID, &lat, &lng)
	if err != nil {
		return Province{}, err
	}
	p := Province{id, name, nameEN, slug, code, regionID, lat, lng}
	return p, nil
}

func (r *Repository) FindDistrictsByProvince(pid int64) ([]District, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM districts WHERE province_id = ?", pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var name string
	var nameEN string
	var code int
	var provinceID int
	var provinceCode int
	var lat float32
	var lng float32

	var districts []District
	for rows.Next() {
		err = rows.Scan(&id, &name, &nameEN, &code, &provinceID, &provinceCode, &lat, &lng)
		if err != nil {
			return nil, err
		}
		d := District{id, name, nameEN, code, provinceID, provinceCode, lat, lng}
		districts = append(districts, d)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return districts, nil
}

func (r *Repository) FindDistrict(pid int64, did int64) (District, error) {
	db, err := getDB()
	if err != nil {
		return District{}, err
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM districts WHERE province_id = ? AND id = ?", pid, did)
	if err != nil {
		return District{}, err
	}
	defer row.Close()

	var id int
	var name string
	var nameEN string
	var code int
	var provinceID int
	var provinceCode int
	var lat float32
	var lng float32

	row.Next()
	err = row.Err()
	if err != nil {
		return District{}, err
	}
	err = row.Scan(&id, &name, &nameEN, &code, &provinceID, &provinceCode, &lat, &lng)
	if err != nil {
		return District{}, err
	}
	d := District{id, name, nameEN, code, provinceID, provinceCode, lat, lng}
	return d, nil
}

func (r *Repository) FindSubdistrictsByProvinceDistrict(pid int64, did int64) ([]Subdistrict, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM subdistricts WHERE province_id = ? AND district_id = ?", pid, did)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var name string
	var nameEN string
	var code int
	var postcode string
	var distID int
	var distCode int
	var provID int
	var provCode int
	var lat float32
	var lng float32

	var subdistricts []Subdistrict
	for rows.Next() {
		err = rows.Scan(&id, &name, &nameEN, &code, &postcode, &distID, &distCode, &provID, &provCode, &lat, &lng)
		if err != nil {
			return nil, err
		}
		s := Subdistrict{id, name, nameEN, code, postcode, distID, distCode, provID, provCode, lat, lng}
		subdistricts = append(subdistricts, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func (r *Repository) FindSubdistrict(pid int64, did int64, sid int64) (Subdistrict, error) {
	db, err := getDB()
	if err != nil {
		return Subdistrict{}, err
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM subdistricts WHERE province_id = ? AND district_id = ? AND id = ?", pid, did, sid)
	if err != nil {
		return Subdistrict{}, err
	}
	defer row.Close()

	var id int
	var name string
	var nameEN string
	var code int
	var postcode string
	var distID int
	var distCode int
	var provID int
	var provCode int
	var lat float32
	var lng float32

	row.Next()
	err = row.Err()
	if err != nil {
		return Subdistrict{}, err
	}
	err = row.Scan(&id, &name, &nameEN, &code, &postcode, &distID, &distCode, &provID, &provCode, &lat, &lng)
	if err != nil {
		return Subdistrict{}, err
	}
	s := Subdistrict{id, name, nameEN, code, postcode, distID, distCode, provID, provCode, lat, lng}
	return s, nil
}

func (r *Repository) FindSubdistrictsByPostcode(p string) ([]Subdistrict, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM subdistricts WHERE postcode = ?", p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var name string
	var nameEN string
	var code int
	var postcode string
	var distID int
	var distCode int
	var provID int
	var provCode int
	var lat float32
	var lng float32

	var subdistricts []Subdistrict
	for rows.Next() {
		err = rows.Scan(&id, &name, &nameEN, &code, &postcode, &distID, &distCode, &provID, &provCode, &lat, &lng)
		if err != nil {
			return nil, err
		}
		s := Subdistrict{id, name, nameEN, code, postcode, distID, distCode, provID, provCode, lat, lng}
		subdistricts = append(subdistricts, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func getDB() (*sql.DB, error) {
	var prefix string
	if len(os.Args) < 2 {
		prefix = "./province/"
	} else {
		prefix = os.Args[1]
	}

	db, err := sql.Open("sqlite3", prefix+"provinces.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
