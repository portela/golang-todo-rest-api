package concurrency

type WebsiteChecker func(string) bool

// As we don't need either value to be named, each of them is anonymous
// within the struct; this can be useful in when it's hard to know what
// to name a value.
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// Anonymous functions:
		// -they can be executed at the same time that they're declared - this is what the ()
		// -they maintain access to the lexical scope they are defined in - all the variables
		//  that are available at the point when you declare the anonymous function are also
		//  available in the body of the function.
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
