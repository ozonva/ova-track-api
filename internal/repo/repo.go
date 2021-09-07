package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/ozonva/ova-track-api/internal/utils"
)

type TrackRepo interface {
	Add([]utils.Track) error
	List(limit, offset uint64) ([]utils.Track, error)
	Describe(id uint64) (*utils.Track, error)
	Remove(id uint64) error
}

type PrintRepo struct{}

func (pr PrintRepo) Add(tracks []utils.Track) error {
	fmt.Println("Add called", tracks)
	return errors.New("")
}
func (pr PrintRepo) List(limit, offset uint64) ([]utils.Track, error) {
	fmt.Printf("")
	return nil, nil
}
func (pr PrintRepo) Describe(id uint64) (*utils.Track, error) {
	fmt.Printf("")
	return nil, nil
}

func (pr PrintRepo) Remove(id uint64) error {
	fmt.Printf("")
	return nil
}

type SQLTrackRepo struct {
	db     *sql.DB
	nextId uint64
	inited bool
}

func NewSQLTrackRepo(pdb *sql.DB) *SQLTrackRepo {
	repo := SQLTrackRepo{pdb, utils.InitialTrackId, false}
	repo.initRepo()
	return &repo
}

func (sql *SQLTrackRepo) initRepo() {
	if !sql.inited {
		tracks, _ := sql.List(utils.MaxTrackId, 0)
		for _, v := range tracks {
			if v.TrackId > sql.nextId {
				sql.nextId = v.TrackId
			}
		}
		sql.nextId++
		sql.inited = true
	}
}

func (sql *SQLTrackRepo) Add(tracks []utils.Track) error {
	if sql.nextId == utils.MaxTrackId {
		return errors.New("cant add anymore tracks")
	}
	builder := squirrel.
		Insert("tracks").
		Columns("id", "name", "album", "artist")

	for _, track := range tracks {
		builder = builder.Values(track.TrackId, track.TrackName, track.Album, track.Artist)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = sql.db.Exec(query, args...)
	return err
}
func (sql *SQLTrackRepo) List(limit, offset uint64) ([]utils.Track, error) {

	var tracks utils.Track
	result := make([]utils.Track, 0, limit)

	query, args, err := squirrel.
		Select("*").
		From("tracks").
		OrderBy("id DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := sql.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&tracks); err != nil {
			return nil, err
		}
		result = append(result, tracks)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
func (sql *SQLTrackRepo) Describe(id uint64) (*utils.Track, error) {
	var res utils.Track
	query, args, err := squirrel.
		Update("tracks").
		Set("name", res.TrackName).
		Set("album", res.Album).
		Set("artist", res.Artist).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	result, err := sql.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return nil, errors.New("nothing to get")
	}
	return &res, nil
}

func (sql *SQLTrackRepo) Remove(id uint64) error {
	query, args, err := squirrel.
		Delete("tracks").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = sql.db.Exec(query, args...)
	return err
}
