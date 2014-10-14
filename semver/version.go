package semver

import (
    "strings"
)

// SemVersion holds the structure of a version, in the form
//
//     X.Y.Z-<stability>-R
//
// where:
//
//     SemVersion.Major holds X
//     SemVersion.Minor holds Y
//     SemVersion.PatchLevel holds Z
//     SemVersion.Stability holds <stability> (blank == 'stable')
//     SemVersion.Release holds the unstable release number
type SemVersion struct {
    Major      int    // X
    Minor      int    // Y
    PatchLevel int    // Z
    Stability  string // stability
    Release    int    // R
}

// returned by SemVersion.Compare when 'rhs' is smaller
const COMP_SMALLER = 0

// returned by SemVersion.Compare when both version numbers are the same
const COMP_EQUAL = 1

// returned by SemVersion.Compare when 'rhs' is larger
const COMP_LARGER = 2

// returned by SemVersion.Compare when the version numbers can't be
// compared (e.g. different stability levels)
const COMP_APPLES_AND_ORANGES = 3

// returned by SemVersion.CompareString when the 'version' can't be
// parsed into a SemVersion struct
const COMP_PARSE_ERROR = -1

// CompareVersions compares two version strings and tells you whether
// one is larger, smaller or the same as the other.
//
// Compare two version strings, returning a COMP_* constant to indicate
// how the right hand side (rhs) compares to the left hand side (lhs)
//
// e.g.
//
//     1.3, 1.4 == COMP_LARGER
//     1.4, 1.3 == COMP_SMALLER
//     1.3, 1.3 == COMP_EQUAL
//
// this is a convenience wrapper around SemVersion.Compare()
func CompareVersions(lhs string, rhs string) (int, error) {
    lhsVersion, err := ParseVersion(lhs)
    if err != nil {
        return COMP_PARSE_ERROR, err
    }

    rhsVersion, err := ParseVersion(rhs)
    if err != nil {
        return COMP_PARSE_ERROR, err
    }

    return lhsVersion.Compare(&rhsVersion), nil
}

// CompareString parses and compares a version string against a SemVersion
// struct.
//
// convenience wrapper around SemVersion.Compare when you don't want to
// parse a version string yourself first
func (lhs *SemVersion) CompareString(version string) (int, error) {
    rhs, err := ParseVersion(version)
    if err != nil {
        return COMP_PARSE_ERROR, err
    }

    return lhs.Compare(&rhs), nil
}

// Compare compares two SemVersion structs against each other.
//
// compares two versions, and returns a COMP_* constant to tell you whether
// the right hand side is larger, smaller, or the same as the left hand
// side
func (lhs *SemVersion) Compare(rhs *SemVersion) int {
    // are both sides comparable at all?
    if lhs.Stability != rhs.Stability {
        return COMP_APPLES_AND_ORANGES
    }

    // which set of rules do we need to use?
    if lhs.Stability == "" {
        // one side is stable - we require the other side to be too
        return lhs.compareStable(rhs)
    }

    // one side is unstable - we require the other side to be too
    return lhs.compareUnstable(rhs)
}

func (lhs *SemVersion) compareStable(rhs *SemVersion) int {
    if lhs.Stability != "" || rhs.Stability != "" {
        return COMP_APPLES_AND_ORANGES
    }

    if lhs.Major < rhs.Major {
        return COMP_LARGER
    }
    if lhs.Major > rhs.Major {
        return COMP_SMALLER
    }

    if lhs.Minor < rhs.Minor {
        return COMP_LARGER
    }
    if lhs.Minor > rhs.Minor {
        return COMP_SMALLER
    }

    if lhs.PatchLevel < rhs.PatchLevel {
        return COMP_LARGER
    }
    if lhs.PatchLevel > rhs.PatchLevel {
        return COMP_SMALLER
    }

    return COMP_EQUAL
}

func (lhs *SemVersion) compareUnstable(rhs *SemVersion) int {
    if lhs.Stability == "" || strings.ToLower(rhs.Stability) != strings.ToLower(rhs.Stability) {
        return COMP_APPLES_AND_ORANGES
    }

    if lhs.Major != rhs.Major {
        return COMP_APPLES_AND_ORANGES
    }

    if lhs.Minor != rhs.Minor {
        return COMP_APPLES_AND_ORANGES
    }

    if lhs.PatchLevel != rhs.PatchLevel {
        return COMP_APPLES_AND_ORANGES
    }

    if lhs.Release < rhs.Release {
        return COMP_LARGER
    }

    if lhs.Release > rhs.Release {
        return COMP_SMALLER
    }

    return COMP_EQUAL
}
