package flusher


import (
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-track-api/internal/mocks"
	"github.com/ozonva/ova-track-api/internal/utils"
)
var _ = Describe("Flusher", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockTrackRepo
	)
	const chunkSize = 2
	var (
		tra, _ = utils.New(0,"foo", "a", "a")
		trb, _ = utils.New(1,"foo", "a", "b")
		trc, _ = utils.New(2,"foo", "a", "c")
		trd, _ = utils.New(3,"foo", "a", "d")
		tre, _ = utils.New(4,"foo", "a", "e")

		testFlusher ChunkFlusher
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockTrackRepo(mockCtrl)
		testFlusher = ChunkFlusher{chunkSize, mockRepo}
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Adding track to storage", func() {
		When("Write success", func() {

			//-------------------------------------------------------------
			//-------------------------------------------------------------

			Context("Write less than chunkSize", func() {
				BeforeEach(func() {
					mockRepo.EXPECT().Add([]utils.Track{tra}).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					Expect(testFlusher.Flush([]utils.Track{tra})).To(BeNil())
				})
			})

			//-------------------------------------------------------------
			//-------------------------------------------------------------

			Context("Write chunkSize", func() {

				BeforeEach(func() {
					mockRepo.EXPECT().Add([]utils.Track{tra,trb}).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					Expect(testFlusher.Flush([]utils.Track{tra,trb})).To(BeNil())
				})
			})

			//-------------------------------------------------------------
			//-------------------------------------------------------------

  			Context("Write more than chunkSize", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().Add([]utils.Track{tra,trb}).Return(nil).Times(1),
						mockRepo.EXPECT().Add([]utils.Track{trc,trd}).Return(nil).Times(1),
						mockRepo.EXPECT().Add([]utils.Track{tre}).Return(nil).Times(1),
					)
				})
				It("Should return nil", func() {
					Expect(testFlusher.Flush([]utils.Track{tra,trb,trc,trd,tre})).To(BeNil())
				})
			})

			//-------------------------------------------------------------
			//-------------------------------------------------------------

		})

		When("Database disconnected", func() {
			err := fmt.Errorf("failed")

			Context("All data", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().Add([]utils.Track{tra,trb}).Return(err).Times(1),
						mockRepo.EXPECT().Add([]utils.Track{trc,trd}).Return(err).Times(1),
						mockRepo.EXPECT().Add([]utils.Track{tre}).Return(err).Times(1),
					)
				})
				It("Should return all failed data", func() {
					Expect(testFlusher.Flush([]utils.Track{tra,trb,trc,trd,tre})).To(Equal([]utils.Track{tra,trb,trc,trd,tre}))
				})
			})
			Context("Add second chunk failed", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().Add([]utils.Track{tra,trb}).Return(nil).Times(1),
						mockRepo.EXPECT().Add([]utils.Track{trc,trd}).Return(err).Times(1),
						mockRepo.EXPECT().Add([]utils.Track{tre}).Return(nil).Times(1),
					)
				})
				It("Should return second chunk", func() {
					Expect(testFlusher.Flush([]utils.Track{tra,trb,trc,trd,tre})).To(Equal([]utils.Track{trc,trd}))
				})
			})
		})
	})
})