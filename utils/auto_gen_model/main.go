package not_main

// import (
// 	"fmt"

// 	"github.com/gohouse/converter"
//
// )

// func main() {
// 	c := config.Config
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", c.DBUser, c.DBPassWord, c.DBHost, c.DBName)

// 	err := converter.NewTable2Struct().
// 		Config(&converter.T2tConfig{
// 			SeperatFile: false,
// 			TagToLower:  true,
// 			UcFirstOnly: false,
// 		}).
// 		//EnableJsonTag(true).
// 		PackageName("schema").
// 		TagKey("json").
// 		Table("report_class_summary_scores").
// 		RealNameMethod("TableName").
// 		SavePath("generated_test.go").
// 		Dsn(dsn).
// 		Run()
// 	fmt.Println(err)
// }
