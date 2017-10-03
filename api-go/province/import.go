package province

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Import() {
	var prefix string
	if len(os.Args) < 2 {
		prefix = "../province/"
	} else {
		prefix = os.Args[1]
	}

	db, err := sql.Open("sqlite3", prefix+"provinces.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		log.Fatal(err)
	}

	var provinces []Province
	var districts []District
	var subdistricts []Subdistrict

	err = extractProvinces(prefix, &provinces)
	if err != nil {
		log.Fatal(err)
	}

	err = insertProvinces(db, provinces)
	if err != nil {
		log.Fatal(err)
	}

	err = extractDistricts(prefix, &districts)
	if err != nil {
		log.Fatal(err)
	}

	err = insertDistricts(db, districts)
	if err != nil {
		log.Fatal(err)
	}

	err = extractSubdistricts(prefix, &subdistricts)
	if err != nil {
		log.Fatal(err)
	}

	err = insertSubdistricts(db, subdistricts)
	if err != nil {
		log.Fatal(err)
	}
}

func createTables(db *sql.DB) error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS provinces (id INTEGER NOT NULL PRIMARY KEY,
																				name TEXT,
																				name_en TEXT,
																				slug TEXT,
																				code INTEGER,
																				region_id INTEGER,
																				lat FLOAT,
																				lng FLOAT);

	CREATE UNIQUE INDEX IF NOT EXISTS idx_provinces_id ON provinces (id);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_provinces_slug ON provinces (slug);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_provinces_code ON provinces (code);
	CREATE INDEX IF NOT EXISTS idx_provinces_region_id ON provinces (region_id);

	CREATE TABLE IF NOT EXISTS districts (id INTEGER NOT NULL PRIMARY KEY,
																				name TEXT,
																				name_en TEXT,
																				code INTEGER,
																				province_id INTEGER,
																				province_code INTEGER,
																				lat FLOAT,
																				lng FLOAT);

	CREATE UNIQUE INDEX IF NOT EXISTS idx_districts_id ON districts (id);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_districts_code ON districts (code);
	CREATE INDEX IF NOT EXISTS idx_districts_province_id ON districts (province_id);
	CREATE INDEX IF NOT EXISTS idx_districts_province_code ON districts (province_code);
	CREATE INDEX IF NOT EXISTS idx_districts_province_id_id ON districts (province_id, id);

	CREATE TABLE IF NOT EXISTS subdistricts (id INTEGER NOT NULL PRIMARY KEY,
																					name TEXT,
																					name_en TEXT,
																					code INTEGER,
																					postcode TEXT,
																					district_id INTEGER,
																					district_code INTEGER,
																					province_id INTEGER,
																					province_code INTEGER,
																					lat FLOAT,
																					lng FLOAT);

	CREATE UNIQUE INDEX IF NOT EXISTS idx_subdistricts_id ON subdistricts (id);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_subdistricts_code ON subdistricts (code);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_postcode ON subdistricts (postcode);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_district_id ON subdistricts (district_id);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_district_code ON subdistricts (district_code);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_province_id ON subdistricts (province_id);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_province_code ON subdistricts (province_code);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_province_id_district_id ON subdistricts (province_id, district_id);
	CREATE INDEX IF NOT EXISTS idx_subdistricts_province_id_district_id_id ON subdistricts (province_id, district_id, id);
	`
	_, err := db.Exec(sqlStmt)
	return err
}

func extractProvinces(prefix string, p *[]Province) error {
	file, err := ioutil.ReadFile(prefix + "provinces.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &p)
	if err != nil {
		return err
	}
	return nil
}

func extractDistricts(prefix string, d *[]District) error {
	file, err := ioutil.ReadFile(prefix + "districts.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &d)
	if err != nil {
		return err
	}
	return nil
}

func extractSubdistricts(prefix string, s *[]Subdistrict) error {
	file, err := ioutil.ReadFile(prefix + "subdistricts.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &s)
	if err != nil {
		return err
	}
	return nil
}

func insertProvinces(db *sql.DB, provinces []Province) error {
	for _, p := range provinces {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO provinces VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(p.ID, p.Name, p.NameEN, p.Slug, p.Code, p.RegionID, p.Lat, p.Lng)
		if err != nil {
			return err
		}
		tx.Commit()
	}
	return nil
}

func insertDistricts(db *sql.DB, districts []District) error {
	for _, d := range districts {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO districts VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(d.ID, d.Name, d.NameEN, d.Code, d.ProvinceID, d.ProvinceCode, d.Lat, d.Lng)
		if err != nil {
			return err
		}
		tx.Commit()
	}
	return nil
}

func insertSubdistricts(db *sql.DB, subdistricts []Subdistrict) error {
	for _, s := range subdistricts {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO subdistricts VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(s.ID, s.Name, s.NameEN, s.Code, s.Postcode, s.DistrictID, s.DistrictCode, s.ProvinceID, 0, s.Lat, s.Lng)
		if err != nil {
			return err
		}
		tx.Commit()
	}
	return nil
}
