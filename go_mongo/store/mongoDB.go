package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DB   *mongo.Database
	Coll map[string]*mongo.Collection
}

// Collection 连接表
func (m *MongoDB) Collection(table string) {
	m.Coll[table] = m.DB.Collection(table)
}

// AddOne 插入一条
// bson.D{{"name", "Alice"}}
// doc := make(map[string]interface{})
// doc["title"] = "test"
// doc["content"] = "this is a test"
func (m *MongoDB) AddOne(ctx context.Context, table string, doc interface{}) (*mongo.InsertOneResult, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	return m.Coll[table].InsertOne(ctx, doc)
}

// AddMore 插入多条
// docs := []interface{}{
// bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}},
// bson.D{{"title", "Showcasing a Blossoming Binary"}, {"text", "Binary data, safely stored with GridFS. Bucket the data"}},
// }
func (m *MongoDB) AddMore(ctx context.Context, table string, docs []interface{}) (*mongo.InsertManyResult, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	return m.Coll[table].InsertMany(ctx, docs)
}

// UpdateOne 更新单条
// filter := bson.D{{"title", "test"}}
// update := bson.D{{"$set", bson.D{{"avg_rating", 4.5}}}}
func (m *MongoDB) UpdateOne(ctx context.Context, table string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	return m.Coll[table].UpdateOne(ctx, filter, update)
}

/*
db.testdata.insert([
   { "name" : "小A", "identify": "aaa","age": 1, "group_identify": "ops"},
   { "name" : "小B", "identify": "bbb","age": 2, "group_identify": "ops"},
   { "name" : "小C", "identify": "ccc","age": 3, "group_identify": "ops"},
   { "name" : "小D", "identify": "ddd","age": 4, "group_identify": "test"},
   { "name" : "小E", "identify": "eee","age": 5, "group_identify": "test"},
   { "name" : "小F", "identify": "fff","age": 6, "group_identify": "test"}
])
*/

// UpdateMany 更新多条
func (m *MongoDB) UpdateMany(ctx context.Context, table string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	return m.Coll[table].UpdateMany(ctx, filter, update)
}

// FindOne 查询单条
func (m *MongoDB) FindOne(ctx context.Context, table string, filter interface{}) (*bson.M, error) {
	var result bson.M
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	err := m.Coll[table].FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindAsID ID查询
func (m *MongoDB) FindAsID(ctx context.Context, table, objID string) (*bson.M, error) {
	var result bson.M
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	id, err := primitive.ObjectIDFromHex(objID)
	if err != nil {
		return nil, err
	}
	err = m.Coll[table].FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	return &result, err
}

type Article struct {
	Title string
	Text  string
}

// FindLike 模糊查询
// findOptions := options.Find()

//filter := bson.D{}
//filter = append(filter, bson.E{
//Key:   "title",
//Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + "a" + ".*", Options: "i"}}}) //i 表示不区分大小写

func (m *MongoDB) FindLike(ctx context.Context, table string, findOptions *options.FindOptions, filter interface{}) ([]*Article, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	cus, err := m.Coll[table].Find(ctx, filter, findOptions)
	if err != nil {

	}
	defer func(cus *mongo.Cursor, ctx context.Context) {
		err := cus.Close(ctx)
		if err != nil {
			return
		}
	}(cus, ctx)
	list := make([]*Article, 0)
	for cus.Next(ctx) {
		article := new(Article)
		err = cus.Decode(&article)
		if err != nil {
			fmt.Printf("decode failed: %v\n", err)
			return nil, err
		}
		list = append(list, article)
	}
	return list, err
}

// FindMany 查询多条
// filter := bson.D{{"age", bson.D{{"$lte", 3}}}}
// 查询年龄小于等于3的，这里特别有意思，能够使用$lte这种方法，类似这样的，
// MongoDB还提供了很多其他的查询方法，比如$gt等等
func (m *MongoDB) FindMany(ctx context.Context, table string, filter interface{}) ([]bson.M, error) {
	/*
		插入测试数据
		db.testdata.insert([
			{ "name" : "小A", "identify": "aaa","age":1},
			{ "name" : "小B", "identify": "bbb","age":2},
			{ "name" : "小C", "identify": "ccc","age":3},
			{ "name" : "小D", "identify": "ddd","age":4},
			{ "name" : "小E", "identify": "eee","age":5},
			{ "name" : "小F", "identify": "fff","age":6},
		])
	*/
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	cursor, err := m.Coll[table].Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	//for _, result := range results {
	//	v, err := encoder.Encode(result, encoder.SortMapKeys)
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	}
	//	fmt.Println(string(v))
	//}
	return results, nil
}

// Replace 替换单条
// filter := bson.D{{"identify", "aaa"}}
// replacement := bson.D{{"name", "小A-r"}, {"identify", "aaa"}, {"content", "this is aaa replace"}}
func (m *MongoDB) Replace(ctx context.Context, table string, filter interface{}, replaceMent interface{}) (*mongo.UpdateResult, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	return m.Coll[table].ReplaceOne(ctx, filter, replaceMent)
}

// DeleteOne 删除单条
// filter := bson.D{{"identify", "aaa"}}
func (m *MongoDB) DeleteOne(ctx context.Context, table string, filter interface{}) error {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	_, err := m.Coll[table].DeleteOne(ctx, filter)
	return err
}

/*
db.testdata.insert([
   { "name" : "小A", "identify": "aaa","age": 1, "group_identify": "ops"},
   { "name" : "小B", "identify": "bbb","age": 2, "group_identify": "ops"},
   { "name" : "小C", "identify": "ccc","age": 3, "group_identify": "ops"},
   { "name" : "小D", "identify": "ddd","age": 4, "group_identify": "test"},
   { "name" : "小E", "identify": "eee","age": 5, "group_identify": "test"},
   { "name" : "小F", "identify": "fff","age": 6, "group_identify": "test"}
])
*/

// DeleteMany 删除多条
// filter := bson.D{{"age", bson.D{{"$gt", 3}}}}
func (m *MongoDB) DeleteMany(ctx context.Context, table string, filter interface{}) error {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	_, err := m.Coll[table].DeleteMany(ctx, filter)
	return err
}

// Count 汇总
// EstimatedDocumentCount()：获得集合中文档数量的近似值
// CountDocuments()：获得集合中文档的确切数量
// filter := bson.D{{"group_identify", "test"}}
func (m *MongoDB) Count(ctx context.Context, table string, filter interface{}) (int64, int64, error) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	estCount, estCountErr := m.Coll[table].EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		return 0, 0, estCountErr
	}
	count, err := m.Coll[table].CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, 0, err
	}
	return estCount, count, nil
}

// Aggregate 关联
func (m *MongoDB) Aggregate(ctx context.Context, table, from string) {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	query := []bson.M{{
		"$lookup": bson.M{
			"from":         from,
			"localField":   "identify",
			"foreignField": "groupIdentify",
			"as":           "output",
		}}}

	/*
		过滤
			query := []bson.M{{
				"$lookup": bson.M{
					"from":         "user",
					"localField":   "identify",
					"foreignField": "groupIdentify",
					"as":           "output",
				}},
				{"$match": bson.M{"identify": "yunweizu"}},
			}
	*/
	coll := m.Coll[table]
	cur, err := coll.Aggregate(ctx, query)
	if err != nil {
		fmt.Printf("aggregate failed:%v\n", err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		// 当数据没有映射到结构体时，可以通过map查询
		one := make(map[string]interface{})
		err := cur.Decode(&one)

		if err != nil {
			fmt.Printf("%v\n", err)
		}
		//v, err := encoder.Encode(one, encoder.SortMapKeys)
		//if err != nil {
		//	fmt.Printf("%v\n", err)
		//}
		//fmt.Println(string(v))
	}
}

// UpdateOneField 添加字段
// $push 在MongoDB中会重复添加
// addToSet 等于$push 但不会重复添加
// $ db.test.update({ "test" : "test" },{ $addToSet: { label_list: "1" } })
func (m *MongoDB) UpdateOneField(ctx context.Context, table string, objID string) error {
	objid, err := primitive.ObjectIDFromHex("62159551120b25bd2c801b09")
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objid}
	updata := bson.M{"$push": bson.M{"link_data": bson.M{"field_identify": "1", "model_data_id": "5"}}}

	// 添加多条
	/*
		linkData := []map[string]string{
				{
					"field_identify": "eryajf_guanliandd",
					"model_data_id":  "6215aaf220ea934fb727096c",
				},
				{
					"field_identify": "eryajf_guanliandd",
					"model_data_id":  "6215aaf220ea934fbaaaaaaa",
				},
			}
			updata := bson.M{"$push": bson.M{"link_data": bson.M{"$each": linkData, "$position": 0}}}
	*/
	_, err = m.Coll[table].UpdateOne(ctx, filter, updata)
	return err
}

// DeleteOneField 删除字段

func (m *MongoDB) DeleteOneField(ctx context.Context, table string, objID string) error {
	if m.Coll[table] == nil {
		m.Collection(table)
	}
	//objid, err := primitive.ObjectIDFromHex("62159551120b25bd2c801b09")
	objid, err := primitive.ObjectIDFromHex(objID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objid}
	updata := bson.M{"$pull": bson.M{"link_data": bson.M{"field_identify": "1", "model_data_id": "5"}}}

	/*
		删除多条

		linkData := []map[string]string{
				{
					"field_identify": "eryajf_guanliandd",
					"model_data_id":  "6215aaf220ea934fb727096c",
				},
				{
					"field_identify": "eryajf_guanliandd",
					"model_data_id":  "6215aaf220ea934fbaaaaaaa",
				},
			}
			updata := bson.M{"$pullAll": bson.M{"link_data_testa": linkData}}
	*/

	_, err = m.Coll[table].UpdateOne(ctx, filter, updata)
	if err != nil {
		return err
	}
	return nil
}
