// iteration package will take in a string and return it repeated 5 times.
package iteration

// Repeat() takes in a string and returns it repeated 5 times.
func Repeat(letter string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += letter
	}
	return repeated
}
