package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Llala/simplecat/util"

	"github.com/stretchr/testify/require"
)

func createRandomApplication(t *testing.T) Application {
	name := util.RandomString(25)
	text := util.RandomText(120, 3, 14)

	arg := CreateApplicationParams{
		Name:       name,
		SourceText: text,
	}

	application, err := testQueries.CreateApplication(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, application)

	require.Equal(t, arg.Name, application.Name)
	require.Equal(t, arg.SourceText, application.SourceText)

	require.NotZero(t, application.ID)
	require.NotZero(t, application.CreatedAt)

	return application
}

func TestCreateApplication(t *testing.T) {
	createRandomApplication(t)
}

func TestGetApplication(t *testing.T) {
	application1 := createRandomApplication(t)
	application2, err := testQueries.GetApplication(context.Background(), application1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, application2)

	require.Equal(t, application1.ID, application2.ID)
	require.Equal(t, application1.Name, application2.Name)
	require.Equal(t, application1.SourceText, application2.SourceText)
	require.WithinDuration(t, application1.CreatedAt, application2.CreatedAt, time.Second)
}

func TestUpdateApplication(t *testing.T) {
	application1 := createRandomApplication(t)

	arg := UpdateApplicationParams{
		ID:         application1.ID,
		SourceText: util.RandomText(120, 4, 20),
	}

	application2, err := testQueries.UpdateApplication(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, application2)

	require.Equal(t, application1.ID, application2.ID)
	require.Equal(t, application1.Name, application2.Name)
	require.Equal(t, arg.SourceText, application2.SourceText)
	require.WithinDuration(t, application1.CreatedAt, application2.CreatedAt, time.Second)

}

func TestDeleteApplication(t *testing.T) {
	application1 := createRandomApplication(t)
	err := testQueries.DeleteApplication(context.Background(), application1.ID)
	require.NoError(t, err)

	application2, err := testQueries.GetApplication(context.Background(), application1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, application2)
}

func TestListApplications(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomApplication(t)
	}
	arg := ListApplicationsParams{
		Limit:  5,
		Offset: 0,
	}

	applications, err := testQueries.ListApplications(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, applications)
}
