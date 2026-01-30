package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsCh := make(chan result)
	for _, url := range urls {
		go func() {
			resultsCh <- result{url, wc(url)}
		}()
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsCh
		results[r.string] = r.bool
	}
	return results
}
