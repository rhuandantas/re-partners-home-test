package usecases

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mock_cache "github.com/rhuandantas/re-partners-home-test/test/mock/repository/cache"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Store Item size", func() {
	var (
		storePackSizeUsecase StorePackSize
		mockCtrl             *gomock.Controller
		memcacheClient       *mock_cache.MockMemcacheClient
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		memcacheClient = mock_cache.NewMockMemcacheClient(mockCtrl)
		storePackSizeUsecase = NewStorePackSize(memcacheClient)
	})

	Context("When storing pack sizes", func() {
		It("should return nil", func() {
			memcacheClient.EXPECT().Set("pack-size", []byte("250,500,1000"), gomock.Any()).Return(nil)
			err := storePackSizeUsecase.Execute([]int{250, 500, 1000})
			Expect(err).To(BeNil())
		})
		When("get an error from storing pack sizes", func() {
			It("should return error", func() {
				memcacheClient.EXPECT().Set("pack-size", []byte("250,500,1000"), gomock.Any()).Return(errors.New("error storing pack sizes"))
				err := storePackSizeUsecase.Execute([]int{250, 500, 1000})
				Expect(err.Error()).To(Equal("error storing pack sizes"))
			})
		})
	})
})
