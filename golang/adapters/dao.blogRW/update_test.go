package blogRW_test

import (
	blogRW "clean_architecture/golang/adapters/dao.blogRW"
	"clean_architecture/golang/testData"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRw_happyUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := blogRW.New(db)
	blog := testData.Blog()

	mock.ExpectExec(regexp.QuoteMeta(blogRW.UpdateSql)).
		WithArgs(blog.ID.Value, blog.Title.Value, blog.Body.Value, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 5))

	// モック化されたDBを用いてテスト対象関数を実行
	if _, err = rw.Update(blog.ID.Value, blog); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRw_unHappyUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := blogRW.New(db)
	blog := testData.Blog()

	mock.ExpectExec(regexp.QuoteMeta(blogRW.UpdateSql)).
		WithArgs(blog.ID.Value, blog.Title.Value, blog.Body.Value, AnyTime{}).
		WillReturnError(fmt.Errorf("some error"))

	// モック化されたDBを用いてテスト対象関数を実行
	if _, err = rw.Update(blog.ID.Value, blog); err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
