package main

import (
	"fmt"
	"log"
	"time"

	"github.com/flosch/pongo2"
)

var tempTpl *pongo2.Template

func main() {
	var err error
	tempTpl, err = pongo2.FromString("学习{{planned_hours}}小时，您已参加{{accumulated_time}}次学习")
	if err != nil {
		panic(err)
	}
	content, err := eventTplExc(map[string]interface{}{
		"planned_hours":    "20",   // 计划学时
		"accumulated_time": "50.0", // 累计学时
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("模板替换成功", content)
}

func eventTplExc(context pongo2.Context) (s string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	return tempTpl.Execute(context)
}

func filterFmtDate(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	layout := param.String()
	if layout == "" {
		return nil, &pongo2.Error{
			Sender:    "filter:parse_date",
			OrigError: errs.New("filter input argument is required"),
		}
	}
	t, parseErr := time.Parse(layout, in.Interface().(string))
	if parseErr != nil {
		return nil, &pongo2.Error{
			Sender:    "filter:fmt_date",
			OrigError: errs.Wrapc(parseErr, "filter input argument must can be parsed to type 'time.Time' by param argument"),
		}
	}
	return pongo2.AsValue(t.Format("2006-01-02")), nil
}
