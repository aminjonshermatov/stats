package stats

import "github.com/aminjonshermatov/bank/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		sum += payment.Amount
	}

	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category != category {
			continue
		}

		sum += payment.Amount
	}

	return sum
}