package tool

import "fmt"

var RootDir = "store_image"

func FormatDate(year, month int) string {
	return fmt.Sprintf("%d年%d月", year, month)
}
