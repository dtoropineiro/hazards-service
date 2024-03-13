package configs

import (
	"context"
	"github.com/Kamva/mgm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hazards-emergencies-go/models"
	"os"
	"strings"
	"time"
)

const firesCollection = "fires"

var dbName string
var mongoUri string

func initOld() {

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(mongoUri))
	if err != nil {
		logrus.Error(err)
	}
}

func Connect() *mongo.Client {
	dbName = viper.GetString("mongodb.database")
	mongoUri = viper.GetString("mongodb.uri")
	mongoPassword := os.Getenv("MONGODB_PASSWORD")
	mongoUri = strings.Replace(mongoUri, "__MONGODB_PASSWORD__", mongoPassword, 1)
	viper.Set("mongodb.uri", mongoUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options.Database()
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		logrus.Error(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Connected to MongoDB")
	return client
}

func GetFiresCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database(dbName).Collection(firesCollection)
	return collection
}

func GetCollectionPaginated(ctx context.Context, client *mongo.Client, limit, page int) []models.Fire {
	collection := client.Database(dbName).Collection(firesCollection)
	curr, err := collection.Find(ctx, bson.D{{}}, GetMongoPaginate(limit, page).GetPaginatedOpts())
	result := make([]models.Fire, 0)
	if err != nil {
		logrus.Error(err)
	}
	for curr.Next(ctx) {
		var fire models.Fire
		if err := curr.Decode(&fire); err != nil {
			logrus.Info(err)
		}

		result = append(result, fire)
	}
	return result
}

type mongoPaginate struct {
	limit int64
	page  int64
}

func GetMongoPaginate(limit, page int) *mongoPaginate {
	return &mongoPaginate{
		limit: int64(limit),
		page:  int64(page),
	}
}

func (mp *mongoPaginate) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := mp.page*mp.limit - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}
