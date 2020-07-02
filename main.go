package muffet

import (
	"fmt"
	"io"
	"os"
)

// func main() {
// 	s, err := command(os.Args[1:], os.Stdout)

// 	if err != nil {
// 		fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}

// 	os.Exit(s)
// }

func Crawl(url string) (map[string][]string, error) {
	s, err := command([]string{url}, os.Stdout)

	if err != nil {
		fprintln(os.Stderr, err)
	}

	return s, nil
}

func command(ss []string, w io.Writer) (map[string][]string, error) {
	args, err := getArguments(ss)
	resultMap := make(map[string][]string)

	if err != nil {
		return resultMap, err
	}

	c, err := newChecker(args.URL, checkerOptions{
		fetcherOptions{
			2, //args.Concurrency,
			args.ExcludedPatterns,
			args.Headers,
			args.IgnoreFragments,
			args.FollowURLParams,
			args.MaxRedirections,
			args.Timeout,
			args.OnePageOnly,
		},
		9999, //args.BufferSize,
		args.FollowRobotsTxt,
		args.FollowSitemapXML,
		args.FollowURLParams,
		args.SkipTLSVerification,
	})

	if err != nil {
		return resultMap, err
	}

	go c.Check()

	for r := range c.Results() {
		if !r.OK() || args.Verbose {
			// fprintln(w, r.String(args.Verbose))
			resultMap[r.url] = r.errorMessages
		}
	}

	return resultMap, nil
}

func fprintln(w io.Writer, xs ...interface{}) {
	if _, err := fmt.Fprintln(w, xs...); err != nil {
		panic(err)
	}
}
