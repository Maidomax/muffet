package muffet

type checkerOptions struct {
	fetcherOptions
	BufferSize int
	FollowRobotsTxt,
	FollowSitemapXML,
	FollowURLParams,
	SkipTLSVerification bool
}
