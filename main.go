package main

import (
	"fmt"
	"snowflake-demo-package/backends"
	"snowflake-demo-package/tables"
)

func main() {
	fields := tables.NewFactsAppEngagement().Fields
	query := tables.Select(fields.AppTitle.WithAlias("App_Name"),
		fields.Foregroundduration.Avg().WithAlias("Avg_Time_in_App"),
		fields.Panelistid.CountDistinct().WithAlias("User_Count"),
		fields.Starttimestamp.Day().WithAlias("Day_of_Week")).
		Filter(fields.AppTitle.Like("%Chime%").
			And(fields.Startmarket.Like("%Wilmington%"))).
		OrderBy(
			backends.Ordering{
				Field:     "User_Count",
				Direction: "DESC",
			}).
		Limit(10)

	compiled, err := query.Compile()

	if err != nil {
		println(fmt.Sprintf("error compiling query: %v", err))
	}

	println(fmt.Sprintf("%v", compiled))

	executed, err := query.Execute()

	if err != nil {
		println(fmt.Sprintf("error executing query: %v", err))
	}

	println(fmt.Sprintf("%v", executed))
}
