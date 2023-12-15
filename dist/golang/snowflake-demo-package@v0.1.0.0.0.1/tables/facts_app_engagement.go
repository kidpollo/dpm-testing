
package tables

import (
	"fmt"
	"sync"

	"snowflake-demo-package/backends"
)

type FactsAppEngagementFields struct {
	Panelistid backends.StringField
	Age backends.Field
	Gender backends.StringField
	Ethnicity backends.StringField
	Factid backends.StringField
	Devicemakemodel backends.StringField
	Operator backends.StringField
	Startmarket backends.StringField
	Startinbuilding backends.StringField
	Startzipcode backends.StringField
	Starttimestamp backends.DateTimeField
	Startlongitude backends.Field
	Startlatitude backends.Field
	Appname backends.StringField
	AppTitle backends.StringField
	Foregroundduration backends.Field
	Foregroundendtime backends.DateTimeField
	Foregroundstarttime backends.DateTimeField
	Screenofftime backends.DateTimeField
	Screenonduration backends.Field
	Screenontime backends.DateTimeField
	Visibleduration backends.Field
	Visibleendtime backends.DateTimeField
	Visiblestarttime backends.DateTimeField
	Isp backends.StringField
	StartDataconntech backends.StringField
	Year backends.Field
	Month backends.Field
	Day backends.Field
}

type FactsAppEngagement struct {
	Table_ *backends.Table
	Fields *FactsAppEngagementFields
}

var (
	instance *FactsAppEngagement
	once     sync.Once
)

func asTableFields(faf *FactsAppEngagementFields) []backends.Expr {
	return []backends.Expr{
        &faf.Panelistid,
	&faf.Age,
	&faf.Gender,
	&faf.Ethnicity,
	&faf.Factid,
	&faf.Devicemakemodel,
	&faf.Operator,
	&faf.Startmarket,
	&faf.Startinbuilding,
	&faf.Startzipcode,
	&faf.Starttimestamp,
	&faf.Startlongitude,
	&faf.Startlatitude,
	&faf.Appname,
	&faf.AppTitle,
	&faf.Foregroundduration,
	&faf.Foregroundendtime,
	&faf.Foregroundstarttime,
	&faf.Screenofftime,
	&faf.Screenonduration,
	&faf.Screenontime,
	&faf.Visibleduration,
	&faf.Visibleendtime,
	&faf.Visiblestarttime,
	&faf.Isp,
	&faf.StartDataconntech,
	&faf.Year,
	&faf.Month,
	&faf.Day,
	}
}

func NewFactsAppEngagement() *FactsAppEngagement {
	once.Do(func() {
		fields := &FactsAppEngagementFields{
            Panelistid: *backends.NewStringField("PANELISTID"),
	Age: *backends.NewField("AGE"),
	Gender: *backends.NewStringField("GENDER"),
	Ethnicity: *backends.NewStringField("ETHNICITY"),
	Factid: *backends.NewStringField("FACTID"),
	Devicemakemodel: *backends.NewStringField("DEVICEMAKEMODEL"),
	Operator: *backends.NewStringField("OPERATOR"),
	Startmarket: *backends.NewStringField("STARTMARKET"),
	Startinbuilding: *backends.NewStringField("STARTINBUILDING"),
	Startzipcode: *backends.NewStringField("STARTZIPCODE"),
	Starttimestamp: *backends.NewDateTimeField("STARTTIMESTAMP"),
	Startlongitude: *backends.NewField("STARTLONGITUDE"),
	Startlatitude: *backends.NewField("STARTLATITUDE"),
	Appname: *backends.NewStringField("APPNAME"),
	AppTitle: *backends.NewStringField("APP_TITLE"),
	Foregroundduration: *backends.NewField("FOREGROUNDDURATION"),
	Foregroundendtime: *backends.NewDateTimeField("FOREGROUNDENDTIME"),
	Foregroundstarttime: *backends.NewDateTimeField("FOREGROUNDSTARTTIME"),
	Screenofftime: *backends.NewDateTimeField("SCREENOFFTIME"),
	Screenonduration: *backends.NewField("SCREENONDURATION"),
	Screenontime: *backends.NewDateTimeField("SCREENONTIME"),
	Visibleduration: *backends.NewField("VISIBLEDURATION"),
	Visibleendtime: *backends.NewDateTimeField("VISIBLEENDTIME"),
	Visiblestarttime: *backends.NewDateTimeField("VISIBLESTARTTIME"),
	Isp: *backends.NewStringField("ISP"),
	StartDataconntech: *backends.NewStringField("START_DATACONNTECH"),
	Year: *backends.NewField("YEAR"),
	Month: *backends.NewField("MONTH"),
	Day: *backends.NewField("DAY"),
		}
		newTable, err := backends.NewTable(
			"4c5a11b5-e302-4c76-8ae0-e254c19bb581",
			"snowflake-demo-package",
			"0.1.0",
			"FACTS_APP_ENGAGEMENT",
			asTableFields(fields),
			nil,
			"https://example.snowflakecomputing.com", // Source path
			nil,
			nil,
			nil,
			1,
		)
		if err != nil {
			// SHould this crash or log?
			println(fmt.Sprintf("error creating NewTable: %v", err))
		}
		instance = &FactsAppEngagement{
			Fields: fields,
			Table_: newTable,
		}
	})
	return instance
}

func (f *FactsAppEngagement) Select(selection ...interface{}) *backends.Table {
	return f.Table_.Select(selection...)
}

func Select(selection ...interface{}) *backends.Table {
	return NewFactsAppEngagement().Select(selection...)
}
