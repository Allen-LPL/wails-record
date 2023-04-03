package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cast"
	"log"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.initDB()
	a.createTable()
}

// initDB initializes the sqlite database
func (a *App) initDB() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	a.db = db
}

// createTable creates the data table if not exists
func (a *App) createTable() {
	sqlStmt := `
	create table if not exists data (
		date text not null,
		contesta text not null,
		contestb text not null,
		typea text not null,
		typeb text not null,
		wina real not null,
		winb real not null,
		drawa real not null,
		drawb real not null,
		losea real not null,
		loseb real not null,
		constraint pk primary key (date, contesta, contestb)
	  );
	  `
	_, err := a.db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// List returns record list
func (a *App) List(name string) interface{} {
	return ""
}

// Data is the structure of the data to be stored or retrieved
//
//	type Data struct {
//		Date     string  `json:"date"`
//		Contesta string  `json:"contesta"`
//		Contestb string  `json:"contestb"`
//		Type     string  `json:"type"`
//		Win      float64 `json:"win"`
//		Draw     float64 `json:"draw"`
//		Lose     float64 `json:"lose"`
//	}
type Data struct {
	Date     string `json:"date"`
	Contesta string `json:"contesta"`
	Contestb string `json:"contestb"`
	TypeA    string `json:"typea"`
	TypeB    string `json:"typeb"`
	WinA     string `json:"wina"`
	WinB     string `json:"winb"`
	DrawA    string `json:"drawa"`
	DrawB    string `json:"drawb"`
	LoseA    string `json:"losea"`
	LoseB    string `json:"loseb"`
}

type ReqData struct {
	Date     string `json:"date"`
	Contesta string `json:"contesta"`
	Contestb string `json:"contestb"`
	TypeA    string `json:"typea"`
	TypeB    string `json:"typeb"`
	WinA     string `json:"wina"`
	WinB     string `json:"winb"`
	DrawA    string `json:"drawa"`
	DrawB    string `json:"drawb"`
	LoseA    string `json:"losea"`
	LoseB    string `json:"loseb"`
}

// SaveData saves a data record to the database
func (a *App) SaveData(data *Data) H {
	fmt.Printf("SaveData: %v", data)

	// checkout all params is not null
	if data.Date == "" || data.Contesta == "" || data.Contestb == "" || data.TypeA == "" || data.TypeB == "" || data.WinA == "" || data.WinB == "" || data.DrawA == "" || data.DrawB == "" || data.LoseA == "" || data.LoseB == "" {
		return M{
			"code": -1,
			"msg":  "ERROR : 参数异常",
		}
	}

	wina := cast.ToFloat64(data.WinA)
	winb := cast.ToFloat64(data.WinB)
	drawa := cast.ToFloat64(data.DrawA)
	drawb := cast.ToFloat64(data.DrawB)
	losea := cast.ToFloat64(data.LoseA)
	loseb := cast.ToFloat64(data.LoseB)

	//Determine whether the data is unique by date and contesta and contestb
	var count int
	err := a.db.QueryRow("select count(*) from data where `date`=? and contesta=? and contestb=?", data.Date, data.Contesta, data.Contestb).Scan(&count)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	if count > 0 {
		return M{
			"code": -1,
			"msg":  "ERROR : 数据已存在",
		}
	}

	// insert data type A
	stmt, err := a.db.Prepare("insert into data(date, contesta, contestb, typea, wina, drawa, losea, typeb, winb, drawb, loseb) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	res, err := stmt.Exec(data.Date, data.Contesta, data.Contestb, data.TypeA, wina, drawa, losea, data.TypeB, winb, drawb, loseb)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	fmt.Println("SaveData:", rowCount)

	return M{
		"code": 200,
		"msg":  "success",
	}
}

// updateData saves a data record to the database
func (a *App) UpdateData(data *Data) H {
	fmt.Printf("SaveData: %v", data)

	// checkout all params is not null
	if data.Date == "" || data.Contesta == "" || data.Contestb == "" || data.TypeA == "" || data.TypeB == "" || data.WinA == "" || data.WinB == "" || data.DrawA == "" || data.DrawB == "" || data.LoseA == "" || data.LoseB == "" {
		return M{
			"code": -1,
			"msg":  "ERROR : 参数异常",
		}
	}

	wina := cast.ToFloat64(data.WinA)
	winb := cast.ToFloat64(data.WinB)
	drawa := cast.ToFloat64(data.DrawA)
	drawb := cast.ToFloat64(data.DrawB)
	losea := cast.ToFloat64(data.LoseA)
	loseb := cast.ToFloat64(data.LoseB)

	//Determine whether the data is not unique by date and contesta and contestb
	var count int
	err := a.db.QueryRow("select count(*) from data where `date`=? and contesta=? and contestb=?", data.Date, data.Contesta, data.Contestb).Scan(&count)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	if count <= 0 {
		return M{
			"code": -1,
			"msg":  "ERROR : 数据不存在",
		}
	}

	// update data by date and contesta and contestb
	stmt, err := a.db.Prepare("update data set wina=?, drawa=?, losea=?, winb=?, drawb=?, loseb=? where `date`=? and contesta=? and contestb=?")
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	res, err := stmt.Exec(wina, drawa, losea, winb, drawb, loseb, data.Date, data.Contesta, data.Contestb)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	fmt.Println("UpdateData:", rowCount)

	return M{
		"code": 200,
		"msg":  "success",
	}
}

// GetDataByDate gets a list of data records by date from the database
func (a *App) GetDataByDate(date []string, page int, pageSize int) H {
	fmt.Printf("dataList: %v", date)
	if len(date) == 0 {
		return a.GetDataList(page, pageSize)
	}
	offset := (page - 1) * pageSize

	sqlScript := fmt.Sprintf("select `date`, contesta, contestb, typea, wina, drawa, losea, typeb, winb, drawb, loseb from data where `date` = '%s' ORDER BY `date` DESC limit %d,%d ", date[0], offset, pageSize)
	if len(date) > 1 || date[1] != date[0] {
		sqlScript = fmt.Sprintf("select `date`, contesta, contestb, typea, wina, drawa, losea, typeb, winb, drawb, loseb from data where `date` between '%s' and '%s' ORDER BY `date` DESC limit %d,%d ", date[0], date[1], offset, pageSize)
	}

	fmt.Println("sqlScript:", sqlScript)
	rows, err := a.db.Query(sqlScript)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}

	// total
	count := 0
	if len(date) > 1 || date[1] != date[0] {
		err = a.db.QueryRow("select count(*) from data where `date` between ? and ?", date[0], date[1]).Scan(&count)
	} else {
		err = a.db.QueryRow("select count(*) from data where `date` = ?", date[0]).Scan(&count)
	}
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}

	defer rows.Close()
	var result []*Data
	for rows.Next() {
		var d Data
		err = rows.Scan(&d.Date, &d.Contesta, &d.Contestb, &d.TypeA, &d.WinA, &d.DrawA, &d.LoseA, &d.TypeB, &d.WinB, &d.DrawB, &d.LoseB)
		if err != nil {
			return M{
				"code": -1,
				"msg":  "ERROR : " + err.Error(),
			}
		}
		result = append(result, &d)
	}
	err = rows.Err()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}

	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code":  200,
		"data":  result,
		"total": count,
	}
}

// GetDataList gets a list of data records list from the database
func (a *App) GetDataList(page, pageSize int) H {
	fmt.Sprintf("page: %d, pageSize: %d \n", page, pageSize)
	offset := (page - 1) * pageSize

	rows, err := a.db.Query(fmt.Sprintf("select `date`, contesta, contestb, typea, wina, drawa, losea, typeb, winb, drawb, loseb from data ORDER BY `date` DESC limit %d,%d ", offset, pageSize))
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}

	defer rows.Close()
	var result []*Data
	for rows.Next() {
		var d Data
		err = rows.Scan(&d.Date, &d.Contesta, &d.Contestb, &d.TypeA, &d.WinA, &d.DrawA, &d.LoseA, &d.TypeB, &d.WinB, &d.DrawB, &d.LoseB)
		if err != nil {
			return M{
				"code": -1,
				"msg":  "ERROR : " + err.Error(),
			}
		}
		result = append(result, &d)
	}
	err = rows.Err()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}

	// total
	row := a.db.QueryRow("select count(*) from data")
	var count int
	err = row.Scan(&count)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}

	return M{
		"code":  200,
		"data":  result,
		"total": count,
	}
}

// DeleteData deletes a data record from the database by date and contesta and contestb
func (a *App) DeleteData(date string, contesta string, contestb string) H {
	// 验证参数
	if date == "" || contesta == "" || contestb == "" {
		return M{
			"code": -1,
			"msg":  "ERROR : 参数异常",
		}
	}

	stmt, err := a.db.Prepare("delete from data where `date`=? and contesta=? and contestb=?")
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	defer stmt.Close()
	res, err := stmt.Exec(date, contesta, contestb)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	fmt.Println("DeleteData:", rowCount)
	return M{
		"code": 200,
		"msg":  "success",
	}
}
