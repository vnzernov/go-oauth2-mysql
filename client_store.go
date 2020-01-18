package mysql

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/gorp.v2"
	"gopkg.in/oauth2.v3"

	// "gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

// ClientStoreItem data item
type ClientStoreItem struct {
	ID     string `db:"id,primarykey"`
	Secret string `db:"secret"`
	Domain string `db:"domain"`
	Data   string `db:"data"`
}

// NewClientDefaultStore create mysql store instance
func NewClientDefaultStore(config *Config) *ClientStore {
	return NewClientStore(config, "", 0)
}

// NewClientStore create mysql store instance,
// config mysql configuration,
// tableName table name (default oauth2_client),
// GC time interval (in seconds, default 600)
func NewClientStore(config *Config, tableName string, gcInterval int) *ClientStore {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.MaxLifetime)

	return NewClientStoreWithDB(db, tableName, gcInterval)
}

// NewClientStoreWithDB create mysql store instance,
// db sql.DB,
// tableName table name (default oauth2_token),
// GC time interval (in seconds, default 600)
func NewClientStoreWithDB(db *sql.DB, tableName string, gcInterval int) *ClientStore {
	clientStore := &ClientStore{
		db:        &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Encoding: "UTF8", Engine: "MyISAM"}},
		tableName: "oauth2_client",
		stdout:    os.Stderr,
	}
	if tableName != "" {
		clientStore.tableName = tableName
	}

	interval := 600
	if gcInterval > 0 {
		interval = gcInterval
	}
	clientStore.ticker = time.NewTicker(time.Second * time.Duration(interval))

	clientStore.db.AddTableWithName(ClientStoreItem{}, clientStore.tableName)
	//table.AddIndex("idx_id", "Btree", []string{"id"})

	err := clientStore.db.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}
	//clientStore.db.CreateIndex()
	return clientStore
}

// ClientStore mysql client store
type ClientStore struct {
	tableName string
	db        *gorp.DbMap
	stdout    io.Writer
	ticker    *time.Ticker
}

// SetStdout set error output
func (s *ClientStore) SetStdout(stdout io.Writer) *ClientStore {
	s.stdout = stdout
	return s
}

// Close close the store
func (s *ClientStore) Close() {
	s.ticker.Stop()
	s.db.Db.Close()
}

func (s *ClientStore) errorf(format string, args ...interface{}) {
	if s.stdout != nil {
		buf := fmt.Sprintf("[OAUTH2-MYSQL-ERROR]: "+format, args...)
		s.stdout.Write([]byte(buf))
	}
}

func (s *ClientStore) toClientInfo(data string) (oauth2.ClientInfo, error) {
	var cm models.Client
	err := jsoniter.Unmarshal([]byte(data), &cm)
	return &cm, err
}

// Create creates and stores the new client information
func (s *ClientStore) Create(info oauth2.ClientInfo) error {
	data, err := jsoniter.Marshal(info)
	if err != nil {
		return err
	}
	err = s.db.Insert(&ClientStoreItem{ID: info.GetID(), Secret: info.GetSecret(), Domain: info.GetDomain(), Data: string(data)})
	if err != nil {
		return err
	}
	return nil
}

// Create creates and stores the new client information
func (s *ClientStore) Delete(id string) error {
	_, err := s.db.Delete(&ClientStoreItem{ID: id})
	if err != nil {
		return err
	}
	return nil
}

// GetByID retrieves and returns client information by id
func (s *ClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=? LIMIT 1", s.tableName)
	var item ClientStoreItem
	err := s.db.SelectOne(&item, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return s.toClientInfo(item.Data)
}
