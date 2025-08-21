package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Llala/simplecat/util"
)

type Store interface {
	Querier
	ParseText(ctx context.Context, arg SourceUnitParams) error
	FormTextTx(ctx context.Context, arg TranslationUnitFormParams) (FormTextTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type SourceUnitParams struct {
	ApplicationID int64  `json:"application_id"`
	Text          string `json:"text"`
}

type TranslationUnitParams struct {
	ApplicationID int64 `json:"application_id"`
	SourceUnitID  int64 `json:"source_unit_id"`
}

func (store *SQLStore) ParseText(ctx context.Context, arg SourceUnitParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// var err error
		ParsedTextArr := util.ParseTextUtil(arg.Text)
		for _, parsed := range ParsedTextArr {
			src_unit, err := q.CreateSourceUnit(ctx, CreateSourceUnitParams{
				ApplicationID: int32(arg.ApplicationID),
				Text: sql.NullString{
					String: parsed,
					Valid:  true,
				},
			})
			if err != nil {
				return err
			}

			translationUnit, err := q.CreateTranslationUnit(ctx, CreateTranslationUnitParams{
				ApplicationID: int32(arg.ApplicationID),
				SourceUnitID:  int32(src_unit.ID),
			})
			if err != nil {
				return err
			}

			_, err = q.UpdateSourceUnit(ctx, UpdateSourceUnitParams{
				ID: src_unit.ID,
				TranslationUnitID: sql.NullInt32{
					Int32: int32(translationUnit.ID),
					Valid: true,
				},
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

type TranslationUnitFormParams struct {
	ApplicationID int64 `form:"application_id" binding:"required"`
}

type FormTextTxResult struct {
	Application Application `json:"application"`
}

func (store *SQLStore) FormTextTx(ctx context.Context, arg TranslationUnitFormParams) (FormTextTxResult, error) {
	var result FormTextTxResult
	err := store.execTx(ctx, func(q *Queries) error {

		translationList, err := q.ListSourceUnitJoinNoLimit(ctx, int32(arg.ApplicationID))
		if err != nil {
			return err
		}

		resultTranslation := ""
		for _, translation := range translationList {
			textUnit := ""
			if translation.TranslationText.String == "" {
				textUnit = translation.SourceText.String

			} else {
				textUnit = translation.TranslationText.String
			}

			resultTranslation = resultTranslation + textUnit + ". "

		}
		resultTranslation = strings.TrimSpace(resultTranslation)

		arg2 := UpdateApplicationParams{
			ID: arg.ApplicationID,
			TranslationText: sql.NullString{
				String: resultTranslation,
				Valid:  true,
			},
		}

		result.Application, err = q.UpdateApplication(ctx, arg2)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
