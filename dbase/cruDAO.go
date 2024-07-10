package dbase

import (
	"context"
	"some_code/apperror"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VacsDAO struct {
	c *mongo.Collection
}

func NewVacsDAO(ctx context.Context, client * mongo.Client, collection string) (*VacsDAO, error) {
	return &VacsDAO{
		c: client.Database("vacpars").Collection(collection),
	}, nil
}

func (dao *VacsDAO) InsertVacs(ctx context.Context, vacs []*Vac) error{

	IMvacs := make([]interface{}, len(vacs))
	for i, vac := range vacs {
		IMvacs[i] = vac
	}
	_, err := dao.c.InsertMany(ctx, IMvacs)
	return err
}

func (dao *VacsDAO) FindbyLang(ctx context.Context, lang string)(*Vac, error){
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