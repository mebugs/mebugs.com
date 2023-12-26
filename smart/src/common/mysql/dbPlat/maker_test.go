package dbPlat

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	temp = `
package dbPlat

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
${imports}
)

// ${tableStruct} ${tableComment}
type ${tableStruct} struct {
${tableComments}}

// ${tableStruct}Table ${tableComment}泛型造器
var ${tableStruct}Table actuator.Table[${tableStruct}]

// DataBase 实现指定数据库
func (t ${tableStruct}) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t ${tableStruct}) TableName() string {
	return "${table}"
}

`
)

var typeMap = map[string]string{
	"bigint":   "uint64",
	"int":      "uint16",
	"smallint": "uint8",
	"varchar":  "string",
	"datetime": "*time.Time",
}
var commonColumn = map[string]bool{"mark": true, "status": true, "create_at": true, "update_at": true}

func TestDBPlatMaker(t *testing.T) {
	InitPlatFromDb()
	// 查询计划生成表
	//tabels := []string{"account"}
	// 查询创建指令

	path := "src/common/mysql/dbPlat/"

	platDb.Exec("use information_schema;")
	var table, tableStruct, tableFile, tableComment string
	table = "login_record"
	platDb.Raw(fmt.Sprintf("select `table_comment` from `tables` where `table_schema` = 'smart' AND `table_name`='%s'", table)).Scan(&tableComment)
	if tableComment == "" {
		t.Logf("Can not Find %s", table)
		return
	}
	tableStruct = toFixStr(table, true)
	tableFile = toFixStr(table, false) + ".go"
	var importHave, commonHave bool
	var aL, bL int
	columnList := make([][3]string, 0)
	rows, _ := platDb.Raw(fmt.Sprintf("select `column_name`,`data_type`,`column_comment` from `columns` where `table_schema` = 'smart' AND `table_name`='%s' ORDER BY ORDINAL_POSITION", table)).Rows()
	// 组装字段
	for rows.Next() {
		var a, b, c string
		rows.Scan(&a, &b, &c)
		// 公共字段
		if commonColumn[a] {
			commonHave = true
			if aL < 6 {
				aL = 6
			}
		} else {
			if len(a) > aL {
				aL = len(a)
			}
			if b == "datetime" {
				importHave = true
			}
			bs := typeMap[b]
			if len(bs) > bL {
				bL = len(bs)
			}
			columnList = append(columnList, [3]string{toFixStr(a, true), bs, c})
		}
	}
	// 组装结构体
	var sb strings.Builder
	for _, item := range columnList {
		tab1 := strings.Repeat(" ", aL-len(item[0])+1)
		tab2 := strings.Repeat(" ", bL-len(item[1])+1)
		sb.WriteString(fmt.Sprintf("\t%s%s%s%s // %s\n", item[0], tab1, item[1], tab2, item[2]))
	}
	if commonHave {
		sb.WriteString("\tCommon\n")
	}
	tableComments := sb.String()
	imports := ""
	if importHave {
		imports = "\t\"time\""
	}
	code := strings.ReplaceAll(temp, "${imports}", imports)
	code = strings.ReplaceAll(code, "${tableStruct}", tableStruct)
	code = strings.ReplaceAll(code, "${tableComment}", tableComment)
	code = strings.ReplaceAll(code, "${tableComments}", tableComments)
	code = strings.ReplaceAll(code, "${table}", table)
	os.WriteFile(path+tableFile, []byte(code), 0777)
}

func toFixStr(s string, pre bool) string {
	// 首字母大写
	ss := s
	if pre {
		ss = strings.ToUpper(s[0:1]) + s[1:]
	}
	sn := ""
	// 下换线转换 _x => X
	ci := -1
	for i, si := range ss {
		if string(si) == "_" {
			ci = i + 1
		} else {
			if i == ci {
				sn = sn + strings.ToUpper(string(si))
			} else {
				sn = sn + string(si)
			}
		}
	}
	return sn
}
