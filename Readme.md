
# Demo DPM golang package 

```bash
dpm build-package -d "Snowflake Demo Package@0.1.0" golang
Going to generate a data package in Golang
Data package already exists at "dist/golang/snowflake-demo-package@v0.1.0.0.0.1", overwrite? yes
Overwriting
Wrote asset "backends/dpm_agent.pb.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/dpm_agent.pb.go"
Wrote asset "backends/dpm_agent_client.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/dpm_agent_client.go"
Wrote asset "backends/dpm_agent_grpc.pb.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/dpm_agent_grpc.pb.go"
Wrote asset "backends/factory.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/factory.go"
Wrote asset "backends/field.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/field.go"
Wrote asset "backends/field_expr.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/field_expr.go"
Wrote asset "backends/field_expr_test.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/field_expr_test.go"
Wrote asset "backends/field_test.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/field_test.go"
Wrote asset "backends/interface.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/interface.go"
Wrote asset "backends/table.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/table.go"
Wrote asset "backends/version.go" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/version.go"
Wrote table definition "FactsAppEngagement" for resource "FACTS_APP_ENGAGEMENT" to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/tables/facts_app_engagement.go"
Wrote version to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/backends/version.go"
Skipping empty entry point
Wrote manifest to "dist/golang/snowflake-demo-package@v0.1.0.0.0.1/go.mod"
Building go module
```

Run

```bash
code/dpm-testing ☊ go run main.go
SELECT
  t0."App_Name",
  t0."Day_of_Week",
  t0."Avg_Time_in_App",
  t0."User_Count"
FROM (
  SELECT
    t1."APP_TITLE" AS "App_Name",
    CAST(EXTRACT(day FROM t1."STARTTIMESTAMP") AS SMALLINT) AS "Day_of_Week",
    AVG(t1."FOREGROUNDDURATION") AS "Avg_Time_in_App",
    COUNT(DISTINCT t1."PANELISTID") AS "User_Count"
  FROM "DPM_CLOUD_DEMO_DB"."PUBLIC"."FACTS_APP_ENGAGEMENT" AS t1
  WHERE
    LOWER(t1."APP_TITLE") LIKE LOWER('%Chime%')
    AND LOWER(t1."STARTMARKET") LIKE LOWER('%Wilmington%')
  GROUP BY
    1,
    2
) AS t0
ORDER BY
  t0."User_Count" DESC
LIMIT 10
[map[App_Name:Chime Banking Avg_Time_in_App:86930.5 Day_of_Week:3 User_Count:5] map[App_Name:Chime Banking Avg_Time_in_App:232557 Day_of_Week:2 User_Count:4] map[App_Name:Chime Banking Avg_Time_in_App:73803.555556 Day_of_Week:7 User_Count:4] map[App_Name:Chime Banking Avg_Time_in_App:81622.4 Day_of_Week:6 User_Count:3] map[App_Name:Chime Banking Avg_Time_in_App:122914.6 Day_of_Week:5 User_Count:3] map[App_Name:Chime Banking Avg_Time_in_App:17693 Day_of_Week:8 User_Count:1]]
```