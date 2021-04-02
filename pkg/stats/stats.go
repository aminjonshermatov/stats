package stats

import "github.com/aminjonshermatov/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	count := int(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}

		sum += payment.Amount
		count++
	}

	return sum / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category != category {
			continue
		}

		if payment.Status == types.StatusFail {
			continue
		}

		sum += payment.Amount
	}

	return sum
}