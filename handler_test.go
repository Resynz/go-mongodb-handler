/**
 * @Author: Resynz
 * @Date: 2020/3/5 16:55
 */
package go_mongodb_handler

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}
	if err = handler.DbClient.Disconnect(nil); err != nil {
		t.Error(err)
	}
}

func TestObjectIdHex(t *testing.T) {
	id, err := ObjectIdHex("5e60b05cc9e7546ab55ae736")
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestHandler_FindById(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}
	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	var r result
	err = handler.FindById(&r, "5e60b05cc9e7546ab55ae736", "info")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestHandler_FindOne(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}
	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	var r result
	err = handler.FindOne(&r, &bson.M{"serial_no": "dsfdsfds"}, "info")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestHandler_FindAll(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}
	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	var r []result
	err = handler.FindAll(&r, &bson.M{"serial_no": "dsfdsfds"}, "info")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestHandler_InsertOne(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	insertId, err := handler.InsertOne(&result{
		SerialNo: "dsfdsfds",
		Status:   "pending",
	}, "info")

	if err != nil {
		t.Error(err)
	}

	t.Log("insertId:", insertId)
}

func TestHandler_InsertMany(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	docs := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		res := &result{
			SerialNo: fmt.Sprintf("insertMany_%d", i),
			Status:   "pending",
		}
		docs[i] = res
	}

	ids, err := handler.InsertMany(docs, "info")
	if err != nil {
		t.Error(err)
	}

	t.Log("insertMany id list:", ids)
}

func TestHandler_UpdateOne(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	count, err := handler.UpdateOne(&bson.D{{"$set", &result{
		SerialNo: "update_1",
		Status:   "running",
	}}}, &bson.M{"serial_no": "hhha_4"}, "info")

	if err != nil {
		t.Error(err)
	}

	t.Log("update one count:", count)
}

func TestHandler_UpdateMany(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	type result struct {
		SerialNo string `bson:"serial_no"`
		Status   string `bson:"status"`
	}

	count, err := handler.UpdateMany(&bson.D{{"$set", &result{
		SerialNo: "hhha_5",
		Status:   "done",
	}}}, &bson.M{"serial_no": "update_333"}, "info")

	if err != nil {
		t.Error(err)
	}

	t.Log("update many count:", count)
}

func TestHandler_DeleteById(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	count, err := handler.DeleteById("5e61e54e1097cfe50c8d7f9e", "info")

	if err != nil {
		t.Error(err)
	}

	t.Log("delete by id count:", count)
}

func TestHandler_DeleteOne(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	count, err := handler.DeleteOne(&bson.M{"serial_no": "hhha_5"}, "info")

	if err != nil {
		t.Error(err)
	}

	t.Log("delete one count:", count)
}

func TestHandler_DeleteMany(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	count, err := handler.DeleteMany(&bson.M{"serial_no": "hhha_0"}, "info")

	if err != nil {
		t.Error(err)
	}

	t.Log("delete many count:", count)
}

func TestHandler_Count(t *testing.T) {
	handler, err := NewHandler(&Config{
		Host:     "localhost",
		Port:     27017,
		UserName: "tester",
		Password: "123456",
		Database: "test",
		Options:  nil,
	})
	if err != nil {
		t.Error(err)
	}

	count, err := handler.Count(&bson.M{"status": "pending"}, "info")

	if err != nil {
		t.Error(err)
	}
	t.Log("count:", count)
}
