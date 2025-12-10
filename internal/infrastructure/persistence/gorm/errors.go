package persistence_gorm

import (
	"errors"
	apperrors "go-boilerplate/internal/pkg/errors"
	"strings"

	"gorm.io/gorm"
)

func HandleDBError(err error, resourceName string) error {
	if err == nil {
		return nil
	}

	// レコードが見つからない場合
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.NewNotFoundError(
			resourceName+" not found",
			err,
		)
	}

	// 一意制約違反の場合
	if isDuplicateError(err) {
		return apperrors.NewAlreadyExistsError(
			resourceName+" already exists",
			err,
		)
	}

	// 外部キー制約違反の場合
	if isForeignKeyError(err) {
		return apperrors.NewValidationError(
			"invalid reference in "+resourceName,
			err,
		)
	}

	// その他のデータベースエラー
	return apperrors.NewInternalError(
		"database error occurred while processing "+resourceName,
		err,
	)
}

// データベースのユニーク制約違反エラーかどうかを判定します
func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return strings.Contains(errMsg, "Error 1062: Duplicate entry") || // MySQL
		strings.Contains(errMsg, "duplicate key value violates unique constraint") // PostgreSQL
}

// isForeignKeyError は外部キー制約違反エラーを判定します
func isForeignKeyError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return strings.Contains(errMsg, "Error 1452: Cannot add or update a child row") || // MySQL
		strings.Contains(errMsg, "violates foreign key constraint") // PostgreSQL
}