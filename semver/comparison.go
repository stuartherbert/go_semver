package semver

import (
    "fmt"
    "strings"
)

// errors returned when a version does not match an expression
var (
    ErrDifferentMajorVersions   = fmt.Errorf("major version numbers are different")
    ErrDifferentMinorVersions   = fmt.Errorf("minor version numbers are different")
    ErrDifferentPatchLevel      = fmt.Errorf("patchlevels are different")
    ErrDifferentStabilityLevels = fmt.Errorf("stability levels are different")
    ErrDifferentReleaseNumbers  = fmt.Errorf("release numbers are different")
    ErrIncomparable             = fmt.Errorf("LHS and RHS are incomparable")
    ErrUnknownOperator          = fmt.Errorf("unknown operator; cannot compare")
    ErrMajorVersionTooSmall     = fmt.Errorf("major version number is too small")
    ErrMinorVersionTooSmall     = fmt.Errorf("minor version number is too small")
    ErrPatchLevelTooSmall       = fmt.Errorf("patchlevel is too small")
    ErrReleaseNumberTooSmall    = fmt.Errorf("release number is too small")
    ErrMajorVersionTooLarge     = fmt.Errorf("major version number is too large")
    ErrMinorVersionTooLarge     = fmt.Errorf("minor version number is too large")
    ErrPatchLevelTooLarge       = fmt.Errorf("patchlevel is too large")
    ErrReleaseNumberTooLarge    = fmt.Errorf("release number is too large")
    ErrUnstableVersion          = fmt.Errorf("unexpected unstable version received")
    ErrStableVersion            = fmt.Errorf("unexpected stable version received")
    ErrOlderUnstableVersion     = fmt.Errorf("older unstable version")
    ErrNewerStableVersion       = fmt.Errorf("newer stable version")
)

// Matches checks to see if 'version' matches the expression that we have
// already parsed.
//
// this is a convenience method around 'MatchesVersion', to avoid parsing
// the 'version' string yourself first
//
// returns 'true' if the version matches the expression in 'lhs'
// returns 'false' plus one of the Err* values if the version does not
// match
func (lhs *VersionExpression) Matches(version string) (bool, error) {
    // we need to turn our raw string into a comparison struct first
    rhs, err := ParseVersion(version)
    if err != nil {
        return false, err
    }

    return lhs.MatchesVersion(&rhs)
}

// MatchesVersion checks to see if 'version' matches the expression that
// we have already parsed.
//
// returns 'true' if the version matches the expression in 'lhs'
// returns 'false' plus one of the Err* values if the version does not
// match
func (lhs *VersionExpression) MatchesVersion(rhs *SemVersion) (bool, error) {
    switch lhs.Operator {
    case OP_EQUALS:
        return lhs.matchesEquals(rhs)

    case OP_GT_EQUALS:
        return lhs.matchesGreaterThanOrEqualTo(rhs)

    case OP_LT_EQUALS:
        return lhs.matchesLessThanOrEqualTo(rhs)

    case OP_TILDE:
        return lhs.matchesCompatibleWith(rhs)
    }

    // if we get here, then we do not recognise the operator
    return false, ErrUnknownOperator
}

func (lhs *VersionExpression) matchesEquals(rhs *SemVersion) (bool, error) {
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

func (lhs *VersionExpression) matchesGreaterThanOrEqualTo(rhs *SemVersion) (bool, error) {
    // are we checking a stable or an unstable release?
    if lhs.Version.Stability == "" {
        return lhs.matchesGreaterThanOrEqualToStable(rhs)
    }

    // we are checking an unstable release
    return lhs.matchesGreaterThanOrEqualToUnstable(rhs)
}

func (lhs *VersionExpression) matchesGreaterThanOrEqualToStable(rhs *SemVersion) (bool, error) {
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

func (lhs *VersionExpression) matchesGreaterThanOrEqualToUnstable(rhs *SemVersion) (bool, error) {
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

func (lhs *VersionExpression) matchesLessThanOrEqualTo(rhs *SemVersion) (bool, error) {
    // are we checking a stable or an unstable release?
    if lhs.Version.Stability == "" {
        return lhs.matchesLessThanOrEqualToStable(rhs)
    }

    // we are checking an unstable release
    return lhs.matchesLessThanOrEqualToUnstable(rhs)
}

func (lhs *VersionExpression) matchesLessThanOrEqualToStable(rhs *SemVersion) (bool, error) {
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

func (lhs *VersionExpression) matchesLessThanOrEqualToUnstable(rhs *SemVersion) (bool, error) {
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

func (lhs *VersionExpression) matchesCompatibleWith(rhs *SemVersion) (bool, error) {
    if lhs.Version.Stability == "" {
        return lhs.matchesCompatibleWithStable(rhs)
    }

    // we are checking an unstable release
    return lhs.matchesCompatibleWithUnstable(rhs)
}

func (lhs *VersionExpression) matchesCompatibleWithStable(rhs *SemVersion) (bool, error) {
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

func (lhs *VersionExpression) matchesCompatibleWithUnstable(rhs *SemVersion) (bool, error) {
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
