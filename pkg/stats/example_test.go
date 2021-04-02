package stats

import (
	"fmt"
	"github.com/aminjonshermatov/bank/pkg/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			ID: 		1,
			Amount: 	1_000_00,
			Category: 	"Foo",
		},
		{
			ID: 		2,
			Amount: 	2_000_00,
			Category: 	"Bar",
		},
		{
			ID: 		3,
			Amount: 	3_000_00,
			Category: 	"Baz",
		},
	}))

	// Output: 200000
}

func ExampleTotalInCategory() {
	fmt.Println(TotalInCategory([]types.Payment{
		{
			ID: 		1,
			Amount: 	10_000_00,
			Category: 	"Shopping",
		},
		{
			ID: 		2,
			Amount: 	1_000_00,
			Category: 	"Foods",
		},
		{
			ID: 		3,
			Amount: 	50_000_00,
			Category: 	"Gadgets",
		},
		{
			ID: 		4,
			Amount: 	5_000_00,
			Category: 	"Shopping",
		},
	}, types.Category("Shopping")))

	// Output: 1500000
}