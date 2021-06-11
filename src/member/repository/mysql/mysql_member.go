package mysql

import (
	"context"
	"database/sql"

	"github.com/sendaljpt/subscription-service/src/domain"
	"github.com/sirupsen/logrus"
)

// "context"
// "database/sql"
// "fmt"

type mysqlMemberRepository struct {
	Conn *sql.DB
}

func NewMysqlMemberRepository(Conn *sql.DB) domain.MemberRepository {
	return &mysqlMemberRepository{Conn}
}

func (m *mysqlMemberRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Member, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Member, 0)
	for rows.Next() {
		mem := domain.Member{}
		err = rows.Scan(
			&mem.Id,
			&mem.Name,
			&mem.Email,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, mem)
	}

	return result, nil

}
func (m *mysqlMemberRepository) Fetch(ctx context.Context) (res []domain.Member, err error) {
	query := `SELECT id, name, email FROM member`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}
