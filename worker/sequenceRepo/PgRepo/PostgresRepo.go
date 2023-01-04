package PgRepo

import (
	"PP/worker/genericMath"
	"PP/worker/sequenceRepo"
	"database/sql"
	_ "encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(cfg Config) (*PostgresDB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres connect error : (%v)", err)
	}
	fmt.Println(db)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

func InitDB() (*PostgresDB, error) {
	viper.AddConfigPath("../config") //what?
	viper.SetConfigName("configDB")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error in reading config: %v", err)
		return nil, err
	}
	db, err := NewPostgresDB(Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error creating db: %v \n", err)
		return nil, err
	}
	return db, nil
}

func (db *PostgresDB) GetSequence(title string) (*genericMath.FloatSequence, error) {
	var a []byte
	err := db.db.QueryRow(
		"SELECT data FROM seq.storage WHERE title = $1 ",
		title,
	).Scan(&a)
	if err != nil {
		log.Printf("err in query to db for %s : %v", title, err)
		return nil, fmt.Errorf("no \" %s \" in database", title)
	}
	ans, err := sequenceRepo.GetSequenceJson(a)
	if err != nil {
		return nil, fmt.Errorf("check user_id exists query row failed: %w", err)
	}
	return ans, nil

}

func ParseSequence(a []uint8) (*genericMath.FloatSequence, error) {
	seq := genericMath.FloatSequence{}
	a = a[1 : len(a)-1]
	line := string(a)
	data := strings.Split(line, ",")
	for _, snum := range data {
		f, _ := strconv.ParseFloat(snum, 64)
		seq.Append(f)
	}

	return &seq, nil
}

func (db *PostgresDB) AddSequence(title string, data []byte) error {
	err := db.db.QueryRow(
		"INSERT INTO seq.storage (title, data) VALUES ($1, $2)",
		title, data,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
	}
	return nil

}
