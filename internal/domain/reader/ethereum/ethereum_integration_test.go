package ethereum

import (
	"testing"

	"github.com/mrbryside/internal/domain/reader"
	"github.com/stretchr/testify/assert"
)

func TestAddTransactions(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "Test integration fetch and add transaction success",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			// Arrange
			eCli := NewEthereumClient()
			rAgg := reader.NewReader([]string{"0x28c6c06298d514db089934071355e5743bf21d60"}, 17065470, 17065471)

			// Act
			err := eCli.FetchTransaction(rAgg)

			// Assert
			for _, tran := range rAgg.TransactionsVo() {
				assert.NotNil(t, tran.Value)
				assert.NotNil(t, tran.Block)
			}
			assert.Nil(t, err)
		})
	}
}
