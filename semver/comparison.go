package semver

import (
    "fmt"
    "strings"
)

var (
    ErrDifferentMajorVersions   = fmt.Errorf("Major version numbers are different")
    ErrDifferentMinorVersions   = fmt.Errorf("Minor version numbers are different")
    ErrDifferentPatchLevel      = fmt.Errorf("Patchlevels are different")
    ErrDifferentStabilityLevels = fmt.Errorf("Stability levels are different")
    ErrDifferentReleaseNumbers  = fmt.Errorf("Release numbers are different")
    ErrIncomparable             = fmt.Errorf("LHS and RHS are incomparable")
    ErrUnknownOperator          = fmt.Errorf("Unknown operator; cannot compare")
    ErrMajorVersionTooSmall     = fmt.Errorf("Major version number is too small")
    ErrMinorVersionTooSmall     = fmt.Errorf("Minor version number is too small")
    ErrPatchLevelTooSmall       = fmt.Errorf("Patchlevel is too small")
    ErrReleaseNumberTooSmall    = fmt.Errorf("Release number is too small")
    ErrMajorVersionTooLarge     = fmt.Errorf("Major version number is too large")
    ErrMinorVersionTooLarge     = fmt.Errorf("Minor version number is too large")
    ErrPatchLevelTooLarge       = fmt.Errorf("Patchlevel is too large")
    ErrReleaseNumberTooLarge    = fmt.Errorf("Release number is too large")
    ErrUnstableVersion          = fmt.Errorf("Unexpected unstable version received")
    ErrStableVersion            = fmt.Errorf("Unexpected stable version received")
    ErrOlderUnstableVersion     = fmt.Errorf("Older unstable version")
    ErrNewerStableVersion       = fmt.Errorf("Newer stable version")
)

func (lhs *Comparison) Matches(raw string) (bool, error) {
    // we need to turn our raw string into a comparison struct first
    rhs, err := ParseVersion(raw)
    if err != nil {
        return false, err
    }

    switch lhs.Operator {
    case OP_EQUALS:
        return lhs.MatchesEquals(&rhs)

    case OP_GT_EQUALS:
        return lhs.MatchesGreaterThanOrEqualTo(&rhs)

    case OP_LT_EQUALS:
        return lhs.MatchesLessThanOrEqualTo(&rhs)

    case OP_TILDE:
        return lhs.MatchesCompatibleWith(&rhs)
    }

    // if we get here, then we do not recognise the operator
    return false, ErrUnknownOperator
}

func (lhs *Comparison) MatchesEquals(rhs *SemVersion) (bool, error) {
    // the left hand side needs to be the same as the right hand side
    //
    // this 'if' comparison might look long-winded, but it has the
    // advantage of not allocating any memory at all
    if lhs.Version.Major != rhs.Major {
        return false, ErrDifferentMajorVersions
    }
    if lhs.Version.Minor != rhs.Minor {
        return false, ErrDifferentMinorVersions
    }
    if lhs.Version.PatchLevel != rhs.PatchLevel {
        return false, ErrDifferentPatchLevel
    }
    if strings.ToLower(lhs.Version.Stability) != strings.ToLower(rhs.Stability) {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Release != rhs.Release {
        return false, ErrDifferentReleaseNumbers
    }

    // if we get to here, then we're happy that everything will work
    return true, nil
}

func (lhs *Comparison) MatchesGreaterThanOrEqualTo(rhs *SemVersion) (bool, error) {
    // are we checking a stable or an unstable release?
    if lhs.Version.Stability == "" {
        return lhs.MatchesGreaterThanOrEqualToStable(rhs)
    } else {
        return lhs.MatchesGreaterThanOrEqualToUnstable(rhs)
    }
}

func (lhs *Comparison) MatchesGreaterThanOrEqualToStable(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability != rhs.Stability {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Stability != "" || rhs.Stability != "" {
        return false, ErrUnstableVersion
    }

    // now it's just a straight-forward check of each of the numerical fields
    // in turn
    if rhs.Major < lhs.Version.Major {
        return false, ErrMajorVersionTooSmall
    }
    if rhs.Major > lhs.Version.Major {
        return true, nil
    }

    // at this point, lhs.X == rhs.X
    if rhs.Minor < lhs.Version.Minor {
        return false, ErrMinorVersionTooSmall
    }
    if rhs.Minor > lhs.Version.Minor {
        return true, nil
    }

    // at this point, lhs.X.Y = rhs.X.Y
    if rhs.PatchLevel < lhs.Version.PatchLevel {
        return false, ErrPatchLevelTooSmall
    }

    // if we get here, then we're good
    //
    // we're a stable version string, so there is no stability level
    // to check at all
    return true, nil
}

func (lhs *Comparison) MatchesGreaterThanOrEqualToUnstable(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability != rhs.Stability {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Stability == "" || rhs.Stability == "" {
        return false, ErrStableVersion
    }

    // now it's just a straight-forward check of each of the numerical fields
    // in turn
    if lhs.Version.Major != rhs.Major {
        return false, ErrDifferentMajorVersions
    }
    if lhs.Version.Minor != rhs.Minor {
        return false, ErrDifferentMinorVersions
    }
    if lhs.Version.PatchLevel != rhs.PatchLevel {
        return false, ErrDifferentPatchLevel
    }

    // we are an unstable release
    if rhs.Release < lhs.Version.Release {
        return false, ErrReleaseNumberTooSmall
    }

    // if we get here, then we're good
    return true, nil
}

func (lhs *Comparison) MatchesLessThanOrEqualTo(rhs *SemVersion) (bool, error) {
    // are we checking a stable or an unstable release?
    if lhs.Version.Stability == "" {
        return lhs.MatchesLessThanOrEqualToStable(rhs)
    } else {
        return lhs.MatchesLessThanOrEqualToUnstable(rhs)
    }
}

func (lhs *Comparison) MatchesLessThanOrEqualToStable(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability != rhs.Stability {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Stability != "" || rhs.Stability != "" {
        return false, ErrUnstableVersion
    }

    // now it's just a straight-forward check of each of the numerical fields
    // in turn
    if lhs.Version.Major < rhs.Major {
        return false, ErrMajorVersionTooLarge
    }
    if lhs.Version.Major > rhs.Major {
        return true, nil
    }

    // at this point, lhs.X == rhs.X
    if lhs.Version.Minor < rhs.Minor {
        return false, ErrMinorVersionTooLarge
    }
    if lhs.Version.Minor > rhs.Minor {
        return true, nil
    }

    // at this point, lhs.X.Y = rhs.X.Y
    if rhs.PatchLevel > lhs.Version.PatchLevel {
        return false, ErrPatchLevelTooLarge
    }

    // if we get here, then we're good
    //
    // we're a stable version string, so there is no stability level
    // to check at all
    return true, nil
}

func (lhs *Comparison) MatchesLessThanOrEqualToUnstable(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability != rhs.Stability {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Stability == "" || rhs.Stability == "" {
        return false, ErrStableVersion
    }

    // now it's just a straight-forward check of each of the numerical fields
    // in turn
    if lhs.Version.Major != rhs.Major {
        return false, ErrDifferentMajorVersions
    }
    if lhs.Version.Minor != rhs.Minor {
        return false, ErrDifferentMinorVersions
    }
    if lhs.Version.PatchLevel != rhs.PatchLevel {
        return false, ErrDifferentPatchLevel
    }

    // we are an unstable release
    if lhs.Version.Release < rhs.Release {
        return false, ErrReleaseNumberTooLarge
    }

    // if we get here, then we're good
    return true, nil
}

func (lhs *Comparison) MatchesCompatibleWith(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability == "" {
        return lhs.MatchesCompatibleWithStable(rhs)
    } else {
        return lhs.MatchesCompatibleWithUnstable(rhs)
    }
}

func (lhs *Comparison) MatchesCompatibleWithStable(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability != rhs.Stability {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Stability != "" || rhs.Stability != "" {
        return false, ErrUnstableVersion
    }

    if lhs.Version.Major != rhs.Major {
        return false, ErrDifferentMajorVersions
    }

    if rhs.Minor < lhs.Version.Minor {
        return false, ErrMinorVersionTooSmall
    }
    if rhs.Minor > lhs.Version.Minor {
        return true, nil
    }

    if rhs.PatchLevel < lhs.Version.PatchLevel {
        return false, ErrPatchLevelTooSmall
    }

    return true, nil
}

func (lhs *Comparison) MatchesCompatibleWithUnstable(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability != rhs.Stability {
        return false, ErrDifferentStabilityLevels
    }
    if lhs.Version.Stability == "" || rhs.Stability == "" {
        return false, ErrStableVersion
    }

    if lhs.Version.Major != rhs.Major {
        return false, ErrDifferentMajorVersions
    }

    if lhs.Version.Minor != rhs.Minor {
        return false, ErrDifferentMinorVersions
    }

    if lhs.Version.PatchLevel != rhs.PatchLevel {
        return false, ErrDifferentPatchLevel
    }

    if rhs.Release < lhs.Version.Release {
        return false, ErrReleaseNumberTooSmall
    }

    return true, nil
}
