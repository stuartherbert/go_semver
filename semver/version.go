package semver

import (
    "strings"
)

// SemVersion holds the structure of a version, in the form
//
// X.Y.Z-<stability>R
type SemVersion struct {
    Major      int    // X
    Minor      int    // Y
    PatchLevel int    // Z
    Stability  string // stability
    Release    int    // R
}

const COMP_SMALLER = 0
const COMP_EQUAL = 1
const COMP_LARGER = 2
const COMP_APPLES_AND_ORANGES = 3

func (lhs *SemVersion) Compare(rhs *SemVersion) int {
    // are both sides comparable at all?
    if lhs.Stability != rhs.Stability {
        return COMP_APPLES_AND_ORANGES
    }

    if lhs.Stability == "" {
        return lhs.CompareStable(rhs)
    } else {
        return lhs.CompareUnstable(rhs)
    }
}

func (lhs *SemVersion) CompareStable(rhs *SemVersion) int {
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

func (lhs *SemVersion) CompareUnstable(rhs *SemVersion) int {
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
