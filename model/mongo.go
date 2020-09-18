package model

import (
	"blogBackend/config"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	UserCol    *mongo.Collection
	ArticleCol *mongo.Collection
	CommentCol *mongo.Collection
)

const (
	USER    = "user"
	ARTICLE = "article"
	COMMENT = "comment"
)

func init() {

	client, err := mongo.NewClient(options.Client().
		ApplyURI(config.Config.Database.Mongo.DatabaseAddress))
	if err != nil {
		log.Fatal(err)
	}

	second := time.Duration(config.Config.Database.Mongo.ClientExpTime) * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	blogDatabase := client.Database(config.Config.Database.Mongo.DatabaseName)
	UserCol = blogDatabase.Collection(config.Config.Database.Mongo.UserCollection)
	ArticleCol = blogDatabase.Collection(config.Config.Database.Mongo.ArticleCollection)
	CommentCol = blogDatabase.Collection(config.Config.Database.Mongo.CommentCollection)
}

// just support one tag search
// if want to find by key except primer key, fill the keyName
// the return interface is []User or []Article or []Comment
func FindAll(colName string, keyName string, keyValue string, tag ...string) (result interface{}, err error) {

	var filter bson.M
	if tag != nil && tag[0] != "" {
		if keyName != "" && keyValue != "" {
			filter = bson.M{"article_tag": tag[0], keyName: keyValue}
		}
		filter = bson.M{"article_tag": tag[0]}
	} else {
		if keyName != "" && keyValue != "" {
			filter = bson.M{keyName: keyValue}
		}
	}

	switch colName {
	case USER:
		{
			cursor, err := UserCol.Find(context.TODO(), filter)
			if err != nil {
				return nil, err
			}

			var results []User
			if err = cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}

			return results, nil
		}

	case ARTICLE:
		{
			cursor, err := ArticleCol.Find(context.TODO(), filter)
			if err != nil {
				return nil, err
			}

			var results []Article
			if err = cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}

			return results, nil
		}

	case COMMENT:
		{
			cursor, err := CommentCol.Find(context.TODO(), filter)
			if err != nil {
				return nil, err
			}

			var results []Comment
			if err = cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}

			return results, nil
		}
	}

	return result, err
}

// the result interface is  User or Article or Comment  module struct
func FindById(colName string, id string) (result interface{}, err error) {

	switch colName {
	case USER:
		{
			var results User

			filter := bson.M{"user_id": id}
			err = UserCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil
		}

	case ARTICLE:
		{
			var results Article

			filter := bson.M{"article_id": id}
			err = ArticleCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil
		}

	case COMMENT:
		{
			var results Comment

			filter := bson.M{"comment_id": id}
			err = CommentCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil
		}
	}

	return result, err

}

func UpdateOne(colName string, content interface{}) (err error) {

	opts := options.Update().SetUpsert(true)

	switch colName {
	case USER:
		filter := bson.M{"user_id": content.(User).UserId}
		update := bson.M{"$set": content.(User)}
		_, err = UserCol.UpdateOne(context.TODO(), filter, update, opts)

	case ARTICLE:
		filter := bson.M{"article_id": content.(Article).ArticleId}
		update := bson.M{"$set": content.(Article)}
		_, err = ArticleCol.UpdateOne(context.TODO(), filter, update, opts)

	case COMMENT:
		filter := bson.M{"comment_id": content.(Comment).CommentId}
		update := bson.M{"$set": content.(Comment)}
		_, err = CommentCol.UpdateOne(context.TODO(), filter, update, opts)

	}

	if err != nil {
		return err
	}

	return nil
}

func Save(colName string, content interface{}) (err error) {

	switch colName {
	case USER:
		_, err = UserCol.InsertOne(context.TODO(), *content.(*User))

	case ARTICLE:
		_, err = ArticleCol.InsertOne(context.TODO(), *content.(*Article))

	case COMMENT:
		_, err = CommentCol.InsertOne(context.TODO(), *content.(*Comment))
	}

	if err != nil {
		return err
	}

	return nil
}

func DeleteById(colName string, id string) error {
	var (
		filter bson.M
		col    *mongo.Collection
	)

	switch colName {
	case USER:
		filter = bson.M{"user_id": id}
		col = UserCol

	case ARTICLE:
		filter = bson.M{"article_id": id}
		col = ArticleCol

	case COMMENT:
		filter = bson.M{"comment_id": id}
		col = CommentCol
	}

	if col == nil {
		return errors.New("In delete,  switch error")
	}

	_, err := col.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil

}
