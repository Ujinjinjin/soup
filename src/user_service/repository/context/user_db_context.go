package context

import (
	"context"
	"github.com/jackc/pgx/v4"
	//"github.com/jmoiron/sqlx"
	"log"
)

type UserDbContext struct {

}

type DbTestTable struct {
	id int
	name string
}

func RowToDbTestTable(row pgx.Row) (DbTestTable, error){
	var testTable DbTestTable
	err := row.Scan(&testTable.id, &testTable.name)
	if err != nil {
		return testTable, err
	}
	return testTable, nil
}

func NewUserDbContext() *UserDbContext {
	return &UserDbContext{

	}
}

func (s *UserDbContext) Dispose() {

}

func (s *UserDbContext) TestSingle() (string, error) {
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), "host=localhost user=local_user password=qwer1234 database=local_db")
	if err != nil {
		log.Printf("ERROR: %s", err)
		return "", err
	}
	defer conn.Close(context.Background())

	testTable, err := RowToDbTestTable(conn.QueryRow(context.Background(), "select id, name from test_table where id=1"))
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}

	return testTable.name, nil
}

func (s *UserDbContext) TestQuery() (string, error) {
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), "host=localhost user=local_user password=qwer1234 database=local_db")
	if err != nil {
		log.Printf("ERROR: %s", err)
		return "", err
	}
	defer conn.Close(context.Background())

	testTable, err := conn.Query(context.Background(), "select id, name from test_table")
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}

	var result []DbTestTable

	for testTable.Next() {
		var id int
		var name string
		testTable.Scan(&id, &name)
		result = append(result, DbTestTable{
			id:   id,
			name: name,
		})
	}

	log.Printf("result: %v", result)

	return "testTable.name", nil
}