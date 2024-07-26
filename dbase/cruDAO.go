package dbase

import (
	"context"
	"some_code/apperror"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VacsDAO struct {
	c *mongo.Collection
}

func NewVacsDAO(ctx context.Context, client * mongo.Client, collection string) (*VacsDAO, error) {
	dao := &VacsDAO{
		c: client.Database("core").Collection("shortUrls"),
	}
	if err := dao.createIndices(ctx); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *VacsDAO) createIndices(ctx context.Context) error {
	_, err := dao.c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "expireAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	})
	return err
}

func (dao *VacsDAO) InsertVacs(ctx context.Context, vacs []*Vac) error {

	IMvacs := make([]interface{}, len(vacs))
	for i, vac := range vacs {
		IMvacs[i] = vac
	}
	_, err := dao.c.InsertMany(ctx, IMvacs)
	return err
}

func (dao *VacsDAO) FindbyLang(ctx context.Context, lang string) (*Vac, error) {
	filter := bson.D{{Key: "lang", Value: lang}}
	var vac Vac
	err := dao.c.FindOne(ctx, filter).Decode(&vac)
	switch {
	case err == nil:
		return &vac, nil
	case err == mongo.ErrNoDocuments:
		return nil, apperror.ErrNotFound
	default:
		return nil, err
	}
}

