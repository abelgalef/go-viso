package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abelgalef/go-viso/pkg/models"
	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Service interface {
	Close() error
	GenerateTableSchema() error
	GetTables() []*models.Table
}

type DBconfig struct {
	User, Password, Host, Port, DbName, DbType string
}

type dbService struct {
	DB      *sql.DB
	Tables  []*models.Table
	Details DBconfig
}

func NewDBService(user, password, host, port, dbName, dbType string) (Service, error) {
	var dsn string

	switch dbType {
	case "mysql":
		cfg := mysql.Config{
			User:      user,
			Passwd:    password,
			Net:       "tcp",
			Addr:      host + ":" + port,
			DBName:    dbName,
			ParseTime: true,
		}
		dsn = cfg.FormatDSN()

	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	default:
		return nil, fmt.Errorf("unsupported database: %s", dbType)
	}

	db, err := sql.Open(dbType, dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &dbService{DB: db, Tables: make([]*models.Table, 5), Details: DBconfig{User: user, Host: host, Port: port, Password: password, DbName: dbName}}, nil
}

func (d *dbService) Close() error {
	if err := d.DB.Close(); err != nil {
		return err
	}
	return nil
}

func (d *dbService) GenerateTableSchema() error {
	d.Tables = make([]*models.Table, 0)

	rows, err := d.DB.Query("SELECT TABLE_NAME, TABLE_TYPE, TABLE_ROWS, CREATE_TIME, UPDATE_TIME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA= ?", d.Details.DbName)
	if err != nil {
		return fmt.Errorf("couldn't generate table schema: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var tables models.Table
		if err := rows.Scan(&tables.Name, &tables.TType, &tables.Rows, &tables.CreatedAt, &tables.UpdatedAt); err != nil {
			return fmt.Errorf("couldn't generate table schema: %s", err.Error())
		}

		fmt.Println(tables)

		if err := d.setTableDiscriptors(&tables); err != nil {
			return err
		}

		d.Tables = append(d.Tables, &tables)
	}

	return nil
}

func (d *dbService) setTableDiscriptors(tables *models.Table) error {
	rows, err := d.DB.Query("describe " + tables.Name)
	if err != nil {
		return fmt.Errorf("couldn't get table schema: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var schema models.TableSchema
		if err := rows.Scan(&schema.Field, &schema.Type, &schema.Null, &schema.Key, &schema.DefaultData, &schema.Extra); err != nil {
			return fmt.Errorf("couldn't get table schema: %s", err.Error())
		}

		tables.Schema = append(tables.Schema, &schema)
	}

	return nil
}

func (d *dbService) GetTables() []*models.Table {
	return d.Tables
}
