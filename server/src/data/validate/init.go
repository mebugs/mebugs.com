package validate

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"reflect"
	"runtime"
	"siteOl.com/stone/server/src/utils/log"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var v *validator.Validate                        // TODO 自定义校验器使用
var transLang = []string{"zh", "en"}             // 支持更多语言请添加
var transMap = make(map[string]ut.Translator, 0) // 校验错误取这个Map进行错误翻译

// init
func init() {
	validate, err := transInit()
	if err != nil {
		log.ErrorF("transInitValidate Fail")
		return
	}
	v = validate
	// TODO 注册自定义校验
	// CustomBindingValidator()
}

// 初始化语言国际化
func transInit() (validate *validator.Validate, err error) {
	if vObj, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(zhT, zhT, enT)
		validate = vObj
		// 初始化全部翻译对象
		for _, local := range transLang {
			trans, o := uni.GetTranslator(local)
			if !o {
				err = errors.New("GetTranslator Err")
				return
			}
			// 注册全部翻译器
			switch local {
			case "en":
				err = enTrans.RegisterDefaultTranslations(validate, trans)
			case "zh":
				err = zhTrans.RegisterDefaultTranslations(validate, trans)
			default:
				err = zhTrans.RegisterDefaultTranslations(validate, trans)
			}
			transMap[local] = trans
		}
	}
	return
}

// 获取函数名
func fnName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	dotIndex := strings.LastIndex(fullName, ".")
	return fullName[dotIndex+1:]
}

func registry(fn validator.Func) {
	funcName := fnName(fn)
	err := v.RegisterValidation(funcName, fn)
	if err != nil {
		return
	}
}

// CustomBindingValidator 注册自定义校验
func CustomBindingValidator() {
	registry(maxLength)
}
