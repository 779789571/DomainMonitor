package core

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Database_init(root_path string) *sql.DB {
	sqldb_path := root_path + "/results/result.db"
	if CheckFileExisted(sqldb_path) {
		//load database
		db, err := sql.Open("sqlite3", sqldb_path)
		if err != nil {
			Errorf("Can't open db file:%s %s\n", sqldb_path, err)
			os.Exit(-1)
		}
		Infof(" Loading database successfully: %s\n", sqldb_path)

		return db
	} else {
		//creat database
		db, err := sql.Open("sqlite3", sqldb_path)
		if err != nil {
			Errorf("Can't create db file:%s %s\n", sqldb_path, err)
			os.Exit(-1)
		}

		CreatTable(db)
		Infof(" Creating database successfully: %s\n", sqldb_path)
		return db
	}
}
func CreatTable(db *sql.DB) {

	table := `CREATE TABLE IF NOT EXISTS domains(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				domain TEXT NOT NULL,
				subdomain TEXT NOT NULL UNIQUE ,
				http_status INT,
				new INT,
				ip TEXT,
				resource TEXT,
				create_time TEXT,
				scan_status TEXT
	)`
	result, err := db.Exec(table)
	Debugf("create table:%s\n", result)
	CheckError(err)
}
func SaveData(db *sql.DB, subdomain Subdomain) {
	time := GetTimeNow()
	stmt, err := db.Prepare("INSERT INTO domains(domain,subdomain,new,resource,create_time,http_status)  values(?, ?,?,?,?,?)")
	CheckError(err)
	res, err1 := stmt.Exec(subdomain.Domain, subdomain.SubdomainName, subdomain.New, subdomain.Resource, time, subdomain.Http_status)
	if err1 != nil {
		Errorf("Wrong when excuting INSERT INTO domains " + subdomain.Domain + "," + subdomain.SubdomainName + " \n")
	}

	if res != nil {
		//Debugf("res:%s\n", res)
		//打印新发现的子域名
		Infof(" found new subdomain! [%s]\n", subdomain.SubdomainName)
	}
}

//获取一个域名的数据库信息
func GetSubdomainList(db *sql.DB, domain string) []string {
	rows, err := db.Query("SELECT subdomain FROM domains where domain = ?", domain)
	CheckError(err)
	defer rows.Close()
	var subdomain string
	var DbSubdomainList []string
	for rows.Next() {
		rows.Scan(&subdomain)
		//Debugf("sql...%s\n", subdomain)
		DbSubdomainList = append(DbSubdomainList, subdomain)
	}

	return DbSubdomainList
}

//更新数据库
func UpdateData(db *sql.DB, subdomain Subdomain) {
	time := GetTimeNow()
	stmt, err := db.Prepare("UPDATE domains SET new = ?,create_time=? WHERE subdomain = ?")
	CheckError(err)
	res, err1 := stmt.Exec(subdomain.New, time, subdomain.SubdomainName)
	if err1 != nil {
		Errorf("Wrong when excuting UPDATE domains " + subdomain.Domain + "," + subdomain.SubdomainName + " \n")
	}
	if res != nil {
		//Debugf("res:%s\n", res)
	}
}

//数据库里有没有新发现的域名，有的话返回true，无则返回false
func IfNewSubdomainFound(db *sql.DB) bool {
	rows, err := db.Query("SELECT MAX(new) FROM domains ")
	CheckError(err)
	defer rows.Close()
	var new int
	for rows.Next() {
		rows.Scan(&new)
	}
	if new == 1 {
		return true
	}
	return false
}

//返回的数据切片形式为 [pass.xiami.com|xiami.com]
func SelectNewSubdomain(db *sql.DB) []string {
	rows, err := db.Query("SELECT subdomain,domain,http_status FROM domains where new = 1 ")
	CheckError(err)
	defer rows.Close()
	var subdomain string
	var domain string
	var http_status int
	var new_subdomain_list []string
	for rows.Next() {
		rows.Scan(&subdomain, &domain, &http_status)
		new_subdomain := subdomain + "|" + domain + "|" + strconv.Itoa(http_status)
		new_subdomain_list = append(new_subdomain_list, new_subdomain)
	}
	return new_subdomain_list
}

//更新所以new字段为0
func UpdateAllNewToOld(db *sql.DB,domain string)  bool{
	stmt, err := db.Prepare("UPDATE domains SET new =0 WHERE domain=?")
	CheckError(err)
	res, err1 := stmt.Exec(domain)
	affect, err := res.RowsAffected()
	if err1 != nil{
		Debugf("update new parma false")
		return false
	}
	Infof("%s 更新了 %d 个域名结果",domain,affect)
	return true
}
