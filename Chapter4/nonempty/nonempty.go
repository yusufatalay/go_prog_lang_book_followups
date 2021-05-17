// Nonempty is an example of an in-placeslice algorithm
// NOTE this function modifies the given array
package main

func nonempty(strings []string) []string {

	i := 0
	for _, s = range strings {

		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings
}
