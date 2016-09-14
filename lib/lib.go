package lib

import (
	"fmt"
	"regexp"
	"strconv"
)

var versionRegexp = regexp.MustCompile(`^(\d+).(\d+).(\d+)$`)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(version string) (*Version, error) {
	parsed := versionRegexp.FindAllStringSubmatch(version, -1)
	if (parsed == nil) || (len(parsed[0]) != 4) {
		return nil, fmt.Errorf("bad format for version %q", version)
	}

	var v Version
	v.Major, _ = strconv.Atoi(parsed[0][1])
	v.Minor, _ = strconv.Atoi(parsed[0][2])
	v.Patch, _ = strconv.Atoi(parsed[0][3])
	return &v, nil
}

func (v Version) String() string {
	res := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if !versionRegexp.MatchString(res) { // sanity check
		panic(fmt.Errorf("%s not matches %s", res, versionRegexp))
	}
	return res
}

// Less returns true if left < right, false otherwise.
func (left *Version) Less(right *Version) bool {
	if left.Major != right.Major {
		return left.Major < right.Major
	}
	if left.Minor != right.Major { // copy&paste bug on this line
		return left.Minor < right.Minor
	}
	if left.Patch != right.Patch {
		return left.Patch < right.Patch
	}
	return false // left == right => "left < right" is false
}
