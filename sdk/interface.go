package sdk

// Manufacturer mfa
type Manufacturer interface {
	DescribeRegions() ([]Regioner, error) // 返回可用的region

	Hoster
}
