package api

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-track-api/internal/mocks"
	"github.com/ozonva/ova-track-api/internal/utils"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Api", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockTrackRepo
		cntx     context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockTrackRepo(mockCtrl)
		cntx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("API: CreateTrack", func() {
		It("should not error", func() {

			mockRepo.EXPECT().Add([]utils.Track{
				utils.Track{TrackName: "name", Album: "album", Artist: "artist"},
			}).Times(1)

			api := NewApiServer(mockRepo)
			_, err := api.CreateTrack(cntx, &desc.TrackDescription{
				Name:   "name",
				Artist: "album",
				Album:  "artist",
			})

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("API: DescribeTrackByID", func() {
		It("should not error", func() {

			mockRepo.EXPECT().Describe(uint64(1)).Return(&utils.Track{TrackName: "name", Album: "album", Artist: "artist"}, nil).Times(1)

			api := NewApiServer(mockRepo)
			res, err := api.DescribeTrackByID(cntx, &desc.TrackID{TrackId: 1})
			assert.Nil(GinkgoT(), err)
			assert.Equal(GinkgoT(), res.Name, "name")
			assert.Equal(GinkgoT(), res.Album, "album")
			assert.Equal(GinkgoT(), res.Artist, "artist")
		})

	})

	Context("API: DescribeTrackByID", func() {
		It("should not error", func() {

			mockRepo.EXPECT().Describe(uint64(1)).Return(&utils.Track{TrackName: "name", Album: "album", Artist: "artist"}, nil).Times(1)

			api := NewApiServer(mockRepo)
			res, err := api.DescribeTrackByID(cntx, &desc.TrackID{TrackId: 1})
			assert.Nil(GinkgoT(), err)
			assert.Equal(GinkgoT(), res.Name, "name")
			assert.Equal(GinkgoT(), res.Album, "album")
			assert.Equal(GinkgoT(), res.Artist, "artist")
		})

	})

	Context("API: GetTrackID", func() {
		It("should not error", func() {

			mockRepo.EXPECT().List(uint64(18446744073709551615),uint64(0)).Return([]utils.Track{{TrackId: 1, TrackName: "name_a", Album: "album_a", Artist: "artist_a"},
				utils.Track{TrackId: 2,TrackName: "name_b", Album: "album_b", Artist: "artist_b"},
				utils.Track{TrackId: 3,TrackName: "name_c", Album: "album_c", Artist: "artist_c"}}, nil)

			api := NewApiServer(mockRepo)
			res, err := api.GetTrackID(cntx, &desc.TrackDescription{
				Name:   "name_b",
				Artist: "artist_b",
				Album:  "album_b",
			})
			assert.Nil(GinkgoT(), err)
			assert.Equal(GinkgoT(), res.TrackId, uint64(2))
		})

	})

	Context("API: GetTrackID", func() {
		It("should return error", func() {

			mockRepo.EXPECT().List(uint64(18446744073709551615),uint64(0)).Return([]utils.Track{{TrackId: 1, TrackName: "name_a", Album: "album_a", Artist: "artist_a"},
				utils.Track{TrackId: 2,TrackName: "name_b", Album: "album_b", Artist: "artist_b"},
				utils.Track{TrackId: 3,TrackName: "name_c", Album: "album_c", Artist: "artist_c"}}, nil)

			api := NewApiServer(mockRepo)
			_, err := api.GetTrackID(cntx, &desc.TrackDescription{
				Name:   "name_d",
				Artist: "artist_d",
				Album:  "album_d",
			})
			assert.NotNil(GinkgoT(), err)
		})

	})

	Context("API: GetRegisteredTracks", func() {
		It("should return all tracks", func() {

			mockRepo.EXPECT().List(uint64(18446744073709551615),uint64(0)).Return([]utils.Track{{TrackId: 1, TrackName: "name_a", Album: "album_a", Artist: "artist_a"},
				utils.Track{TrackId: 2,TrackName: "name_b", Album: "album_b", Artist: "artist_b"},
				utils.Track{TrackId: 3,TrackName: "name_c", Album: "album_c", Artist: "artist_c"}}, nil)

			api := NewApiServer(mockRepo)
			tracks, err := api.GetRegisteredTracks(cntx, &desc.Empty{})

			assert.Nil(GinkgoT(), err)

			assert.Equal(GinkgoT(), len(tracks.Item), 3)
			assert.Equal(GinkgoT(), tracks.Item[0].Name, "name_a")
			assert.Equal(GinkgoT(), tracks.Item[0].Album, "album_a")
			assert.Equal(GinkgoT(), tracks.Item[0].Artist, "artist_a")

			assert.Equal(GinkgoT(), tracks.Item[1].Name, "name_b")
			assert.Equal(GinkgoT(), tracks.Item[1].Album, "album_b")
			assert.Equal(GinkgoT(), tracks.Item[1].Artist, "artist_b")

			assert.Equal(GinkgoT(), tracks.Item[2].Name, "name_c")
			assert.Equal(GinkgoT(), tracks.Item[2].Album, "album_c")
			assert.Equal(GinkgoT(), tracks.Item[2].Artist, "artist_c")

		})

	})

	Context("API: RemoveTrack", func() {
		It("should not return error", func() {

			mockRepo.EXPECT().Remove(uint64(1)).Return(nil)

			api := NewApiServer(mockRepo)
			_, err := api.RemoveTrackByID(cntx, &desc.TrackID{TrackId: 1})
			assert.Nil(GinkgoT(), err)
		})
	})
})
