package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// AddStream inserts stream into database
func (storage *Storage) AddTopic(topic entity.TopicInterface) (result entity.TopicInterface, err error) {
	return nil, nil
	//const addTopicSQL = `
	//	INSERT INTO stream (
	//		track
	//	) VALUES (
	//		$1
	//	) RETURNING stream_id
	//`
	//
	//tx, err := storage.connPool.Begin()
	//if err != nil {
	//	return nil, pgError(err)
	//}
	//defer func() {
	//	if err != nil {
	//		pgRollback(tx)
	//	}
	//}()
	//
	//var id int64
	//if err := tx.QueryRow(addStreamSQL, stream.GetTrack()).Scan(&id); err != nil {
	//	return nil, pgError(err)
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	return nil, pgError(err)
	//}
	//
	//result = entity.NewStream(id, stream.GetTrack())
	//
	//return result, nil
}