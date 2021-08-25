package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-track-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-track-api/internal/repo TrackRepo
//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozonva/ova-track-api/internal/saver Saver
