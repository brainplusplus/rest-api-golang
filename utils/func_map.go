package utils

import (
	"fmt"
	"github.com/leekchan/accounting"
	"html/template"
	"math/big"
	"simple-ecommerce/types"
	"strconv"
	"time"
)

func GetAllFuncMap() template.FuncMap {
	return template.FuncMap{
		"json_date_to_string":   funcMapJsonDateToString,
		"json_date_format":      funcMapJsonDateFormat,
		"time_to_string":        funcMapTimeToString,
		"time_format":           funcMapTimeFormat,
		"display_default_money": funcMapDisplayDefaultMoney,
		"display_money":         funcMapDisplayMoney,
	}
}

func funcMapJsonDateToString(val types.JsonDate) string {
	return val.ToTimeString()
}

func funcMapJsonDateFormat(val types.JsonDate, format string) string {
	return val.Format(format)
}

func funcMapTimeToString(val time.Time) string {
	return time.Time(val).Format("2006-01-02")
}

func funcMapTimeFormat(val time.Time, format string) string {
	return val.Format(format)
}

func funcMapDisplayDefaultMoney(money interface{}) template.HTML {
	return funcMapDisplayMoney(money, "Rp.", 0, ".", ",")
}

func funcMapDisplayMoney(money interface{}, symbol string, precision int, thousand string, decimal string) template.HTML {
	val := ""
	var moneyFloat float64
	moneyStr := fmt.Sprintf("%v", money)
	//log.Info("displayMoney")
	//log.Info(moneyStr)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	bigVal := big.NewFloat(moneyFloat) //new(big.Float).SetPrec(2).SetString(moneyStr)
	//log.Info(bigVal)
	ac := accounting.Accounting{Symbol: symbol, Precision: precision, Thousand: thousand, Decimal: decimal}
	moneyFloat, _ = bigVal.Float64()
	val = ac.FormatMoney(moneyFloat)
	//log.Info(val)
	//log.Info("end displayMoney")
	return template.HTML(val)
}
