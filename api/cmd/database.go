package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-pg/pg/v9"
)

// DBLogger can be used to log SQL query strings on their way to or from
// the postgres database. To use this, put the following line in your code
// before where you want to see the query go-pg uses:
// s.pgdb.AddQueryHook(database.DBLogger{})
type DBLogger struct{}

// BeforeQuery is invoked before the query is run.
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	fmt.Println(q.FormattedQuery())
	return ctx, nil
}

// AfterQuery is invoked after the query is run.
func (d DBLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	return nil
}

func setup(connStr string) (*pg.DB, error) {
	// Using a custom parser of the conn string because go-pg doesn't support sslmode=verify-full nor sslrootcert
	pgoptions, err := parseOptionsFromURL(connStr)
	if err != nil {
		log.Println("parsing connection string error:", err)
		return nil, err
	}

	pgoptions.ReadTimeout = 5 * time.Second
	// we want an extremely long write timeout for the initial load of old records
	pgoptions.WriteTimeout = 5 * time.Second
	db := pg.Connect(pgoptions)

	var n int
	_, err = db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		log.Println("QueryOne error:", err)
		return db, err
	}

	// this string slice op is safe even if no @ is present in connStr;
	// if no @, strings.Index returns -1 (which we then +1 to get back to 0)
	// if @ present, we get only the bit after it, if no @, we get the whole string.
	log.Println("Connected to postgres database:", connStr[strings.Index(connStr, "@")+1:])
	return db, nil
}

func parseOptionsFromURL(connString string) (*pg.Options, error) {
	connURL, err := url.Parse(connString)
	if err != nil {
		return nil, fmt.Errorf("database url parse error: %s", err.Error())
	}

	pgoptions := &pg.Options{
		Addr:     connURL.Host,
		Database: strings.TrimPrefix(connURL.Path, "/"),
	}

	pgoptions.User, pgoptions.Password = parseUserInfo(connURL.User)

	queryParams := connURL.Query()
	sslmode := queryParams.Get("sslmode")
	sslrootcert := queryParams.Get("sslrootcert")

	if (len(sslmode) == 0 || sslmode == "disable") || len(sslrootcert) == 0 {
		return pgoptions, nil
	}

	// assuming & trying to replicate verify-full
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         connURL.Hostname(),
	}

	tlsConfig.RootCAs = x509.NewCertPool()

	cert, err := ioutil.ReadFile(filepath.Clean(sslrootcert))
	if err != nil {
		log.Println("sslrootcert error:", err)
		return nil, err
	}

	if !tlsConfig.RootCAs.AppendCertsFromPEM(cert) {
		return nil, errors.New("couldn't parse pem in sslrootcert")
	}

	pgoptions.TLSConfig = tlsConfig

	return pgoptions, nil
}

func parseUserInfo(u *url.Userinfo) (user, password string) {
	if u == nil {
		return "postgres", ""
	}

	user = u.Username()
	password, _ = u.Password()
	return
}
