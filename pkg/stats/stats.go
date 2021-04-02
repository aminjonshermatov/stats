package stats

import "github.com/aminjonshermatov/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	count := 0

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

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	paymentCountOfCategory := map[types.Category]int{}
	categoriesAmount := map[types.Category]types.Money{}

	for _, payment := range payments {
		categoriesAmount[payment.Category] += payment.Amount
		paymentCountOfCategory[payment.Category]++
	}

	for category := range categoriesAmount {
		categoriesAmount[category] /= types.Money(paymentCountOfCategory[category])
	}

	return categoriesAmount
}