package utils

import (
	"encoding/json"
	"fmt"
	models "hello/Models"
	"os"
	"reflect"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/goyesql/v2"
	goyesqlx "github.com/knadh/goyesql/v2/sqlx"
)

type Cache struct {
	cacheMap sync.Map
}

func (c *Cache) Set(key string, value any) {
	c.cacheMap.Store(key, value)
}
func (c *Cache) Get(key string) (any, bool) {
	val, ok := c.cacheMap.Load(key)
	if ok {
		return val, ok
	} else {
		return "", false
	}
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsZero()
}

func StructsToMap(in interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)
	var resMap = make(map[string]any)

	for k, v := range inInterface {
		resMap[k] = v
	}
	return resMap
}

func ReadQueries(sqlFile string) goyesql.Queries {
	qB, err := os.ReadFile(sqlFile) // For read access.
	if err != nil {
		panic("error reading SQL file")
	}
	fmt.Println(string(qB))
	qMap, err := goyesql.ParseBytes(qB)
	if err != nil {
		panic("error parsing SQL queries")
	}

	return qMap
}

func PrepareQueries(qMap goyesql.Queries, db *sqlx.DB) *models.Queries {
	// Scan and prepare all queries.
	var q models.Queries
	if err := goyesqlx.ScanToStruct(&q, qMap, db.Unsafe()); err != nil {
		panic("error preparing SQL queries")
	}
	return &q
}

func InstallSchema(schemaFile string, db *sqlx.DB) error {
	qB, err := os.ReadFile(schemaFile) // For read access.
	if err != nil {
		panic("error reading SQL file")
	}

	fmt.Println(string(qB))
	_, errDB := db.Exec(string(qB))
	return errDB
}
