// Given a year, report if it is a leap year.

package tasks

// Lapyear check is year is leap
func Lapyear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
