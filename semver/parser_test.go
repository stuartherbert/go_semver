package semver

import (
    "testing"
)

// ========================================================================
//
// Tests for ParseVersion()
//
// ------------------------------------------------------------------------

func TestCanParseMajorMinor(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      1,
        Minor:      3,
        PatchLevel: 0,
        Stability:  "",
        Release:    0,
    }

    // perform the test
    actual, err := ParseVersion("1.3")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseMajorMinorPatchlevel(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      1,
        Minor:      3,
        PatchLevel: 6,
        Stability:  "",
        Release:    0,
    }

    // perform the test
    actual, err := ParseVersion("1.3.6")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseLargeMajorMinorPatchlevel(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      100,
        Minor:      365,
        PatchLevel: 699,
        Stability:  "",
        Release:    0,
    }

    // perform the test
    actual, err := ParseVersion("100.365.699")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseMajorMinorUnstableRelease(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      1,
        Minor:      3,
        PatchLevel: 0,
        Stability:  "alpha",
        Release:    1,
    }

    // perform the test
    actual, err := ParseVersion("1.3-alpha-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseMajorMinorPatchlevelUnstableRelease(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      1,
        Minor:      3,
        PatchLevel: 6,
        Stability:  "alpha",
        Release:    1,
    }

    // perform the test
    actual, err := ParseVersion("1.3.6-alpha-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestStabilityLevelCaseIsPreserved(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      1,
        Minor:      3,
        PatchLevel: 6,
        Stability:  "SNAPSHOT",
        Release:    1,
    }

    // perform the test
    actual, err := ParseVersion("1.3.6-SNAPSHOT-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestStabilityLevelCanIncludeAnUnderscore(t *testing.T) {
    // what result do we expect?
    expected := SemVersion{
        Major:      1,
        Minor:      3,
        PatchLevel: 6,
        Stability:  "alpha_romeo",
        Release:    1,
    }

    // perform the test
    actual, err := ParseVersion("1.3.6-alpha_romeo-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

// ========================================================================
//
// Tests for Parse() with =
//
// ------------------------------------------------------------------------

func TestCanParseEqualsOperatorWithMajorMinor(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 0,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse("=1.3")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseEqualsOperatorWithMajorMinorPatchlevel(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse("=1.3.6")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseEqualsOperatorWithMajorMinorPatchlevelUnstableRelease(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "alpha",
            Release:    1,
        },
    }

    // perform the test
    actual, err := Parse("=1.3.6-alpha-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

// ========================================================================
//
// Tests for Parse() with >=
//
// ------------------------------------------------------------------------

func TestCanParseGreaterThanOrEqualToOperatorWithMajorMinor(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_GT_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 0,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse(">=1.3")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanGreaterThanOrEqualToOperatorWithMajorMinorPatchlevel(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_GT_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse(">=1.3.6")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseGreaterThanOrEqualToOperatorWithMajorMinorPatchlevelUnstableRelease(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_GT_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "alpha",
            Release:    1,
        },
    }

    // perform the test
    actual, err := Parse(">=1.3.6-alpha-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

// ========================================================================
//
// Tests for Parse() with <=
//
// ------------------------------------------------------------------------

func TestCanParseLessThanOrEqualToOperatorWithMajorMinor(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_LT_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 0,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse("<=1.3")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseLessThanOrEqualToOperatorWithMajorMinorPatchlevel(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_LT_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse("<=1.3.6")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseLessThanOrEqualToOperatorWithMajorMinorPatchlevelUnstableRelease(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_LT_EQUALS,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "alpha",
            Release:    1,
        },
    }

    // perform the test
    actual, err := Parse("<=1.3.6-alpha-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

// ========================================================================
//
// Tests for Parse() with ~
//
// ------------------------------------------------------------------------

func TestCanParseTildeOperatorWithMajorMinor(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_TILDE,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 0,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse("~1.3")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseTildeOperatorWithMajorMinorPatchlevel(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_TILDE,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "",
            Release:    0,
        },
    }

    // perform the test
    actual, err := Parse("~1.3.6")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

func TestCanParseTildeOperatorWithMajorMinorPatchlevelUnstableRelease(t *testing.T) {
    // what result do we expect?
    expected := Comparison{
        Operator: OP_TILDE,
        Version: SemVersion{
            Major:      1,
            Minor:      3,
            PatchLevel: 6,
            Stability:  "alpha",
            Release:    1,
        },
    }

    // perform the test
    actual, err := Parse("~1.3.6-alpha-1")

    // was an error returned?
    if err != nil {
        t.Error(err)
        return
    }

    // did we get back what we expected?
    if actual != expected {
        t.Errorf("Expected %d, received %d", expected, actual)
        return
    }
}

// ========================================================================
//
// Tests for Parse() with @
//
// ------------------------------------------------------------------------
