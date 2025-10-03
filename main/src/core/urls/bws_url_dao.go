package urls

import (
	ses_database "bws_microservice_url/main/src/config/database"
	"bws_microservice_url/main/src/entity"
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BwsProjectDao struct {
	collection *mongo.Collection
}

func (t *BwsProjectDao) Constructor1() *BwsProjectDao {
	t.collection = ses_database.BwsProjectCollection()
	return t
}

/**
 * Insert project
 */
func (t *BwsProjectDao) insertUrl(url entity.BwsUrlTrack) (err error) {
	_, err = t.collection.InsertOne(context.TODO(), url)
	return
}

/**
 * get urls tracks
 */
func (t *BwsProjectDao) getUrlTracks(idCustomer primitive.ObjectID) (urlTracks []entity.BwsUrlTrack, err error) {
	filter := bson.D{{"idCustomer", idCustomer}}
	cursor, err := t.collection.Find(context.TODO(), filter)
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Error().Err(err).Msg("error to close cursor")
		}
	}()
	urlTracks = make([]entity.BwsUrlTrack, 0)
	err = cursor.All(context.TODO(), &urlTracks)
	return
}

/**
 * Get coupon by code and project name
 */
func (t *BwsProjectDao) getUrlTrack(url string) (exist bool, urlTrack entity.BwsUrlTrack, err error) {
	filter := bson.D{
		{"url", url},
	}
	singleResult := t.collection.FindOne(context.TODO(), filter)
	if err = singleResult.Decode(&urlTrack); errors.Is(err, mongo.ErrNoDocuments) {
		return false, urlTrack, nil
	}
	if err = singleResult.Decode(&urlTrack); err != nil {
		return false, urlTrack, err
	}
	return true, urlTrack, err
}
