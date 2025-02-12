package usecases

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mock_cache "github.com/rhuandantas/re-partners-home-test/test/mock/repository/cache"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Pack Items", func() {
	var (
		packItemUsecase PackItem
		mockCtrl        *gomock.Controller
		memcacheClient  *mock_cache.MockMemcacheClient
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		memcacheClient = mock_cache.NewMockMemcacheClient(mockCtrl)
		packItemUsecase = NewPackItem(memcacheClient)
	})

	Context("When packing items", func() {
		It("should return the correct pack sizes", func() {
			memcacheClient.EXPECT().Get("pack-size").Return([]byte("250,500,1000"), nil)
			packs, err := packItemUsecase.Execute(1000)
			Expect(err).To(BeNil())
			Expect(packs[1000]).To(Equal(1))
		})
		When("pack sizes are not found", func() {
			It("should return an error", func() {
				memcacheClient.EXPECT().Get("pack-size").Return([]byte(""), nil)
				_, err := packItemUsecase.Execute(1000)
				Expect(err.Error()).To(Equal("pack sizes not found. please add pack sizes"))
			})
		})
		When("error getting pack sizes", func() {
			It("should return an error", func() {
				memcacheClient.EXPECT().Get("pack-size").Return(nil, errors.New("error getting pack sizes"))
				_, err := packItemUsecase.Execute(1000)
				Expect(err.Error()).To(Equal("error getting pack sizes"))
			})
		})
		When("error parsing pack sizes", func() {
			It("should return an error", func() {
				memcacheClient.EXPECT().Get("pack-size").Return([]byte("250,500,1000,invalid"), nil)
				_, err := packItemUsecase.Execute(1000)
				Expect(err.Error()).To(Equal("error parsing pack sizes"))
			})
		})

		DescribeTable("Packing items",
			func(a int, expected map[int]int) {
				memcacheClient.EXPECT().Get("pack-size").Return([]byte("250,500,1000,2000,5000"), nil)
				result, _ := packItemUsecase.Execute(a)
				for k, v := range expected {
					Expect(result).To(HaveKeyWithValue(k, v))
				}
			},

			Entry("Packing 5000", 5000, map[int]int{5000: 1}),
			Entry("Packing 1000", 1000, map[int]int{1000: 1}),
			Entry("Packing 10000", 10000, map[int]int{5000: 2}),
			Entry("Packing 10001", 10001, map[int]int{5000: 2, 250: 1}),
			Entry("Packing 13000", 13000, map[int]int{5000: 2, 2000: 1, 1000: 1}),
			Entry("Packing 13501", 13501, map[int]int{5000: 2, 2000: 1, 1000: 1, 500: 1, 250: 1}),
			Entry("Packing 13752", 13752, map[int]int{5000: 2, 2000: 1, 1000: 1, 500: 1, 250: 2}),
		)
	})
})
