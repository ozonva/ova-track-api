package saver

import (

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-track-api/internal/mocks"
	"github.com/ozonva/ova-track-api/internal/utils"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl *gomock.Controller
		mockFlusher *mocks.MockFlusher
		testSaver BufferedSaver
	)

	var (
		tra, _ = utils.New(0,"foo", "a", "a")
		trb, _ = utils.New(1,"foo", "a", "b")
		trc, _ = utils.New(2,"foo", "a", "c")
		trd, _ = utils.New(3,"foo", "a", "d")
		tre, _ = utils.New(4,"foo", "a", "e")
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockCtrl)
		testSaver = NewBufferSaver(mockFlusher)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Adding track to storage", func() {
		When("Write success", func() {

			//-------------------------------------------------------------
			//-------------------------------------------------------------

			Context("Should repeat flush buffered if is not flushed", func() {
				BeforeEach(func() {
					gomock.InOrder(
						//mockFlusher.EXPECT().Flush([]utils.Track{tra,trb}).Return([]utils.Track{tra,trb}).Times(1),
						mockFlusher.EXPECT().Flush([]utils.Track{tra,trb}).Return([]utils.Track{tra,trb}).Times(1),
						mockFlusher.EXPECT().Flush([]utils.Track{tra,trb,trc,trd}).Return([]utils.Track{tra,trb,trc,trd}).Times(1),
					)
				})
				It("Should return nil", func() {
					testSaver.SaveToBuffer([]utils.Track{tra,trb})
					testSaver.SaveToBuffer([]utils.Track{trc,trd})
					testSaver.SaveToBuffer([]utils.Track{tre})

				})
			})

			//-------------------------------------------------------------
			//-------------------------------------------------------------

			Context("Should work normally", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockFlusher.EXPECT().Flush([]utils.Track{tra,trb}).Return(nil).Times(1),
						mockFlusher.EXPECT().Flush([]utils.Track{trc,trd}).Return(nil).Times(1))
				})
				It("Should return nil", func() {
					testSaver.SaveToBuffer([]utils.Track{tra,trb})
					testSaver.SaveToBuffer([]utils.Track{trc,trd})
					testSaver.SaveToBuffer([]utils.Track{tre})
				})
			})
			//-------------------------------------------------------------
			//-------------------------------------------------------------

			Context("Force flush", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockFlusher.EXPECT().Flush([]utils.Track{tra,trb}).Return([]utils.Track{tra,trb}).Times(1),
					)
				})
				It("Should return nil", func() {
					testSaver.SaveToBuffer([]utils.Track{tra,trb})
					testSaver.FlushBuffer()
				})
			})

			//-------------------------------------------------------------
			//-------------------------------------------------------------

		})


	})
})