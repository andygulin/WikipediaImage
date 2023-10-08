package tool

import "fmt"

func FormatDate(year, month int) string {
	return fmt.Sprintf("%d年%d月", year, month)
}
