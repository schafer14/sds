package mongo_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"github.com/segmentio/ksuid"

	"github.com/schafer14/sds"
	mongoStorage "github.com/schafer14/sds/mongo"
	mtest "github.com/schafer14/sds/mongo/test"
	"github.com/schafer14/sds/test"
)

type entity struct {
	ID    string `bson:"_id"`
	Field string
}

func (entity *entity) GetID() string {
	return entity.ID
}

func (entity *entity) String() string {
	return entity.ID
}

func TestMongoDB(t *testing.T) {

	if testing.Short() {
		t.Skip("docker tests do not run in short mode")
	}

	t.Parallel()
	ctx := context.Background()
	is := is.New(t)

	container := mtest.SetupDatabase(t)
	t.Cleanup(func() { mtest.TeardownDatabase(t, container) })
	db, err := mtest.DatabaseTest(t, container)
	is.NoErr(err)

	coll := db.Collection("test_" + ksuid.New().String())

	store, err := mongoStorage.New[*entity](coll)
	is.NoErr(err)

	test.DoesItWork(t, ctx, store, func(id string) error {
		return store.Save(ctx, &entity{
			ID:    id,
			Field: id,
		})
	})

}

func TestMongoDBDataStructure(t *testing.T) {

	if testing.Short() {
		t.Skip("docker tests do not run in short mode")
	}

	t.Parallel()
	ctx := context.Background()
	is := is.New(t)

	container := mtest.SetupDatabase(t)
	t.Cleanup(func() { mtest.TeardownDatabase(t, container) })
	db, err := mtest.DatabaseTest(t, container)
	is.NoErr(err)
	coll := db.Collection("test_" + ksuid.New().String())

	store, err := mongoStorage.New[*entity](coll)
	is.NoErr(err)

	err = store.Save(ctx, &entity{
		ID:    "abc",
		Field: "123",
	})
	is.NoErr(err)

	res, err := store.Find(ctx, "abc")
	is.NoErr(err)

	is.Equal(res.Field, "123")
	is.Equal(res.ID, "abc")

	lots, curs, err := store.Query(ctx, sds.WithLimit(42))
	is.NoErr(err)
	is.Equal(curs, nil)
	is.Equal(len(lots), 1)
	is.Equal(lots[0].Field, "123")
	is.Equal(lots[0].ID, "abc")
}
