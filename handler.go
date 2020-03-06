/**
 * @Author: Resynz
 * @Date: 2020/3/5 16:14
 */
package go_mongodb_handler

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Handler struct {
	DbClient *mongo.Client
	DbConfig *Config
}

func ObjectIdHex(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func generateApplyURI(conf *Config) string {
	base := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conf.UserName, conf.Password, conf.Host, conf.Port, conf.Database)
	lo := len(conf.Options)
	if lo == 0 {
		return base
	}
	ops := make([]string, lo)
	for i, v := range conf.Options {
		ops[i] = fmt.Sprintf("%s=%s", v.Name, v.Value)
	}

	return fmt.Sprintf("%s?%s", base, strings.Join(ops, "&"))
}

func NewHandler(config *Config) (*Handler, error) {
	var err error
	handler := &Handler{
		DbClient: nil,
		DbConfig: config,
	}
	op := options.Client().ApplyURI(generateApplyURI(config))
	handler.DbClient, err = mongo.Connect(nil, op)
	if err != nil {
		return nil, err
	}

	err = handler.DbClient.Ping(nil, nil)
	return handler, err
}

func (h *Handler) FindById(v interface{}, id, name string) error {
	oid, err := ObjectIdHex(id)
	if err != nil {
		return err
	}
	return h.FindOne(v, &bson.M{"_id": oid}, name)
}

func (h *Handler) FindOne(v, filter interface{}, name string) error {
	return h.DbClient.Database(h.DbConfig.Database).Collection(name).FindOne(nil, filter).Decode(v)
}

func (h *Handler) FindAll(v, filter interface{}, name string) error {
	c, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).Find(nil, filter)
	if err != nil {
		return err
	}

	return c.All(nil, v)
}

func (h *Handler) InsertOne(v interface{}, name string) (string, error) {
	result, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).InsertOne(nil, v)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (h *Handler) InsertMany(documents []interface{}, name string) ([]string, error) {
	res, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).InsertMany(nil, documents)
	if err != nil {
		return nil, err
	}
	results := make([]string, len(res.InsertedIDs))
	for i, v := range res.InsertedIDs {
		results[i] = v.(primitive.ObjectID).Hex()
	}
	return results, nil
}

func (h *Handler) UpdateOne(v, filter interface{}, name string) (int64, error) {
	res, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).UpdateOne(nil, filter, v)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func (h *Handler) UpdateMany(v, filter interface{}, name string) (int64, error) {
	res, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).UpdateMany(nil, filter, v)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func (h *Handler) DeleteById(id, name string) (int64, error) {
	oid, err := ObjectIdHex(id)
	if err != nil {
		return 0, err
	}
	return h.DeleteOne(&bson.M{"_id": oid}, name)
}

func (h *Handler) DeleteOne(filter interface{}, name string) (int64, error) {
	res, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).DeleteOne(nil, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (h *Handler) DeleteMany(filter interface{}, name string) (int64, error) {
	res, err := h.DbClient.Database(h.DbConfig.Database).Collection(name).DeleteMany(nil, filter)
	if err != nil {
		return 0, err

	}
	return res.DeletedCount, nil
}

func (h *Handler) Count(filter interface{}, name string) (int64, error) {
	return h.DbClient.Database(h.DbConfig.Database).Collection(name).CountDocuments(nil, filter)
}
