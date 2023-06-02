package sqlite

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lea-video/go-lea-music/db"
	"github.com/lea-video/go-lea-music/utility"
)

//go:embed sqlite_scripts/*.sql
var embeddedScripts embed.FS

type SQLiteDB struct {
	db *sql.DB
}

func InitSQLite(filename string) (db.GenericDB, error) {
	db, err := openOrCreateDatabase(filename)
	if err != nil {
		return nil, err
	}

	return &SQLiteDB{db: db}, nil
}

func runUpgradeScripts(db *sql.DB, currentVersion int) error {
	nextV := currentVersion + 1
	for {
		data, err := embeddedScripts.ReadFile(fmt.Sprintf("sqlite_scripts/sql_upgrade_v%03d.sql", nextV))
		if errors.Is(err, os.ErrNotExist) {
			break
		}
		if err != nil {
			return err
		}

		_, err = db.Exec(string(data))
		if err != nil {
			log.Printf("Failed to execute upgrade script to v%d\nError: %v", nextV, err)
			return err
		}

		log.Printf("Executed upgrade script to v%d", nextV)
		nextV += 1
	}

	return nil
}

func queryVersion(db *sql.DB) (int, error) {
	// Query the version from the info table
	var rawVersion string
	err := db.QueryRow("SELECT val FROM info WHERE key=\"version\" LIMIT 1").Scan(&rawVersion)
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(rawVersion)
}

func openOrCreateDatabase(filename string) (*sql.DB, error) {
	// Open DB
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	// if the file does not yet exist we simply create it and are done
	if !utility.FileExists(filename) {
		init_script, err := embeddedScripts.ReadFile("sqlite_scripts/sql_init_tables.sql")
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(string(init_script))
		if err != nil {
			return nil, err
		}
		// we could exit now, but for good measure we still run the upgrade scripts
	}

	// If the file does exist we have to care about upgrading the database
	version, err := queryVersion(db)
	if err != nil {
		return nil, err
	}

	// Run table upgrade scripts
	err = runUpgradeScripts(db, version)

	return db, err
}
