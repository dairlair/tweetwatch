package storage

import (
	pb "github.com/dairlair/twitwatch/pkg/api/v1"
)

// AddStream inserts stream into database
func (storage *Storage) AddStream(stream *pb.Stream) (id int64, err error) {

	const addStreamSQL = `
		INSERT INTO stream (
			track
		) VALUES (
			$1
		) RETURNING stream_id
	`

	tx, err := storage.connPool.Begin()
	if err != nil {
		return 0, pgError(err)
	}

	if err := tx.QueryRow(addStreamSQL, stream.Track).Scan(&id); err != nil {
		tx.Rollback()
		return 0, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}

// GetStreams returns
func (storage *Storage) GetStreams() (streams []*pb.Stream, err error) {
	const getStreamsSQL = `
		SELECT 
			stream_id
			, track
		FROM
			stream
	`

	rows, err := storage.connPool.Query(getStreamsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stream pb.Stream
		if err := rows.Scan(
			&stream.Id,
			&stream.Track,
		); err != nil {
			return nil, err
		}
		streams = append(streams, &stream)
	}

}
