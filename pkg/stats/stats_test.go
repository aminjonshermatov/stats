package stats

import (
	"github.com/aminjonshermatov/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 8_000_00,
		"food": 2_000_00,
		"fun": 	5_000_00,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestCategoriesAvg_success(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 2_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
		{ID: 6, Category: "fun", Amount: 3_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
		"food": 2_000_00,
		"fun": 4_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_nil(t *testing.T) {
	var first map[types.Category]types.Money
	var second map[types.Category]types.Money

	result := PeriodsDynamic(first, second)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestPeriodsDynamic_empty(t *testing.T) {
	first := map[types.Category]types.Money{}
	second := map[types.Category]types.Money{}

	result := PeriodsDynamic(first, second)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestPeriodsDynamic_onTwoPartsExistPositive(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_onTwoPartsExistNegative(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_notExistOnSecond(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_notExistOnFirst(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}
	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}