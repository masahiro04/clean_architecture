package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/domains"
	"clean_architecture/golang/testData"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTest(name string) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"test-db",
		5432,
		"postgresql",
		"postgresql",
		"test-api",
	)

	txdb.Register(name, "postgres", conn)
	dialector := postgres.New(
		postgres.Config{
			DriverName: name,
			DSN:        conn,
		},
	)
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db, nil
}

func Prepare(name string, seeds []interface{}) (*gorm.DB, error) {
	db, err := NewTest(name)
	if err != nil {
		return nil, err
	}

	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}

// TODO: 単体だと動くけど、NewTest/Prepareは一度しか動かない。parallelだけでなく、db connectionそもそもの問題のようなきも。。。。
func TestGetAll(t *testing.T) {

	blogs := testData.Blogs(5)

	seeds := []interface{}{
		&blogDao.BlogDto{
			ID:    blogs.Value[0].ID.Value,
			Title: blogs.Value[0].Title.Value,
			Body:  blogs.Value[0].Body.Value,
		},
		&blogDao.BlogDto{
			ID:    blogs.Value[1].ID.Value,
			Title: blogs.Value[1].Title.Value,
			Body:  blogs.Value[1].Body.Value,
		},
		&blogDao.BlogDto{
			ID:    blogs.Value[2].ID.Value,
			Title: blogs.Value[2].Title.Value,
			Body:  blogs.Value[2].Body.Value,
		},
		&blogDao.BlogDto{
			ID:    blogs.Value[3].ID.Value,
			Title: blogs.Value[3].Title.Value,
			Body:  blogs.Value[3].Body.Value,
		},
	}

	db, err := Prepare("users_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := blogDao.New(db)

	tests := []struct {
		name      string
		length    int
		wantError error
	}{
		{
			name:      "Get blogs",
			length:    blogs.Size(),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		// NOTE(okubo): ttにtt入れるとparallelのscopeの問題を回避できるので、一旦そのままで実装してます
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// TODO(okubo): parallelの方が圧倒的に早いけど、goroutineの影響？で
			// db.Close()のタイミング合わないので、一旦は並行処理は断念
			// t.Parallel()

			b, err := rw.GetAll()

			if tt.wantError == nil {
				assert.NoError(t, err)
				assert.Equal(t, b.Size(), tt.length)
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestGetById(t *testing.T) {

	blog := testData.Blog()
	dummyBlog := testData.BlogWithID(100) // 100でnot foud起こす
	seeds := []interface{}{
		&blogDao.BlogDto{
			ID:    blog.ID.Value,
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}

	db, err := Prepare("user_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := blogDao.New(db)

	tests := []struct {
		name      string
		blog      *domains.Blog
		wantError error
	}{
		{
			name:      "Get a blog",
			blog:      &blog,
			wantError: nil,
		},
		{
			name:      "Return not found error",
			blog:      &dummyBlog,
			wantError: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// TODO(okubo): parallelの方が圧倒的に早いけど、goroutineの影響？で
			// db.Close()のタイミング合わないので、一旦は並行処理は断念
			// t.Parallel()

			_, err := rw.GetById(tt.blog.ID.Value)
			if tt.wantError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
