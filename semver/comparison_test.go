package semver

import (
    "fmt"
    "testing"
)

type ExpectedError struct {
    lhs string
    rhs string
    err error
}

// ========================================================================
//
// Compare two versions using the equals operator
//
// ------------------------------------------------------------------------

func TestCanMatchUsingEquals(t *testing.T) {
    // what result do we expect?
    expected := true

    // our list of strings to match
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered equivalent
    var toMatch = [][2]string{
        [2]string{"=1.3", "1.3"},
        [2]string{"=1.3", "1.3.0"},
        [2]string{"=1.3.0", "1.3.0"},
        [2]string{"=1.3.1", "1.3.1"},
        [2]string{"=2.6-alpha-1", "2.6-alpha-1"},
        [2]string{"=2.6-alpha-1", "2.6.0-alpha-1"},
        [2]string{"=2.5.99-alpha-1", "2.5.99-alpha-1"},
        [2]string{"=2.5.99-ALPHA-1", "2.5.99-ALPHA-1"},
        [2]string{"=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141013"},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet[0])
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet[1])

        // was an error returned?
        if err != nil {
            fmt.Println(lhs)
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != expected {
            t.Errorf("Expected %d, received %d", expected, actual)
            return
        }
    }
}

func TestCannotMatchUsingEquals(t *testing.T) {
    // our list of strings to compare
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered non-equivalent
    var toMatch = []ExpectedError{
        ExpectedError{"=1.3", "11.3", ErrDifferentMajorVersions},
        ExpectedError{"=1.3", "1.33", ErrDifferentMinorVersions},
        ExpectedError{"=1.3.0", "11.3.0", ErrDifferentMajorVersions},
        ExpectedError{"=1.3.0", "1.33.0", ErrDifferentMinorVersions},
        ExpectedError{"=1.3.1", "1.3.2", ErrDifferentPatchLevel},
        ExpectedError{"=2.6-alpha-1", "2.6-alpha-2", ErrDifferentReleaseNumbers},
        ExpectedError{"=2.6-alpha-1", "2.6.0-alpha-2", ErrDifferentReleaseNumbers},
        ExpectedError{"=2.6-alpha-1", "2.6.0-beta-2", ErrDifferentStabilityLevels},
        ExpectedError{"=2.5.99-alpha-1", "2.5.99-alpha-2", ErrDifferentReleaseNumbers},
        ExpectedError{"=2.5.99-ALPHA-1", "2.5.99-ALPHA-2", ErrDifferentReleaseNumbers},
        ExpectedError{"=2.5.99-ALPHA-1", "2.5.99-BETA-2", ErrDifferentStabilityLevels},
        ExpectedError{"=2.5.99-ALPHA-1", "2.5.99", ErrDifferentStabilityLevels},
        ExpectedError{"=2.5.99", "2.5.99-BETA-2", ErrDifferentStabilityLevels},
        ExpectedError{"=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141012", ErrDifferentReleaseNumbers},
        ExpectedError{"=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141014", ErrDifferentReleaseNumbers},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet.lhs)
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet.rhs)

        // was an error returned?
        if err != matchSet.err {
            fmt.Println(lhs)
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != false {
            t.Errorf("Expected %d, received %d", false, actual)
            return
        }
    }
}

// ========================================================================
//
// Compare two versions using the greater than or equals operator
//
// ------------------------------------------------------------------------

func TestCanMatchUsingGreaterThanOrEquals(t *testing.T) {
    // what result do we expect?
    expected := true

    // our list of strings to match
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered equivalent
    var toMatch = [][2]string{
        // this first set is the same that we use in the 'equals' operator
        // test above ... they're all expected to pass too
        [2]string{">=1.3", "1.3"},
        [2]string{">=1.3", "1.3.0"},
        [2]string{">=1.3.0", "1.3.0"},
        [2]string{">=1.3.1", "1.3.1"},
        [2]string{">=2.6-alpha-1", "2.6-alpha-1"},
        [2]string{">=2.6-alpha-1", "2.6.0-alpha-1"},
        [2]string{">=2.5.99-alpha-1", "2.5.99-alpha-1"},
        [2]string{">=2.5.99-ALPHA-1", "2.5.99-ALPHA-1"},
        [2]string{">=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141013"},

        // this second set should only work when used with the '>=' operator
        [2]string{">=1.3", "2.0"},
        [2]string{">=1.3", "2.0.0"},
        [2]string{">=1.3", "2.1.0"},
        [2]string{">=1.3", "2.1.1"},
        [2]string{">=1.3", "1.4"},
        [2]string{">=1.3", "1.3.1"},
        [2]string{">=1.3.0", "2.0"},
        [2]string{">=1.3.0", "2.0.0"},
        [2]string{">=1.3.0", "2.1.0"},
        [2]string{">=1.3.0", "2.1.1"},
        [2]string{">=1.3.0", "1.4"},
        [2]string{">=1.3.0", "1.3.1"},
        [2]string{">=1.3-alpha-1", "1.3-alpha-2"},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet[0])
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet[1])

        // was an error returned?
        if err != nil {
            fmt.Println(lhs)
            fmt.Println(matchSet[1])
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != expected {
            t.Errorf("Expected %d, received %d", expected, actual)
            return
        }
    }
}

func TestCannotMatchUsingGreaterThanOrEquals(t *testing.T) {
    // our list of strings to compare
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered non-equivalent
    var toMatch = []ExpectedError{
        ExpectedError{">=1.3", "0.3", ErrMajorVersionTooSmall},
        ExpectedError{">=1.3", "1.2", ErrMinorVersionTooSmall},
        ExpectedError{">=1.3.0", "0.3.0", ErrMajorVersionTooSmall},
        ExpectedError{">=1.3.0", "1.2.0", ErrMinorVersionTooSmall},
        ExpectedError{">=1.3.1", "1.3.0", ErrPatchLevelTooSmall},
        ExpectedError{">=2.6-alpha-2", "2.6-alpha-1", ErrReleaseNumberTooSmall},
        ExpectedError{">=2.6-alpha-2", "2.6.0-alpha-1", ErrReleaseNumberTooSmall},
        ExpectedError{">=2.6-alpha-1", "2.6.0-beta-2", ErrDifferentStabilityLevels},
        ExpectedError{">=2.6-alpha-1", "2.6.0", ErrDifferentStabilityLevels},
        ExpectedError{">=2.6", "2.6.0-beta-2", ErrDifferentStabilityLevels},
        ExpectedError{">=2.5.99-alpha-2", "2.5.99-alpha-1", ErrReleaseNumberTooSmall},
        ExpectedError{">=2.5.99-ALPHA-2", "2.5.99-ALPHA-1", ErrReleaseNumberTooSmall},
        ExpectedError{">=2.5.99-ALPHA-1", "2.5.99-BETA-2", ErrDifferentStabilityLevels},
        ExpectedError{">=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141012", ErrReleaseNumberTooSmall},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet.lhs)
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet.rhs)

        // was an error returned?
        if err != matchSet.err {
            fmt.Println(matchSet.lhs)
            fmt.Println(matchSet.rhs)
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != false {
            t.Errorf("Expected %d, received %d", false, actual)
            return
        }
    }
}

// ========================================================================
//
// Compare two versions using the less than or equals operator
//
// ------------------------------------------------------------------------

func TestCanMatchUsingLessThanOrEquals(t *testing.T) {
    // what result do we expect?
    expected := true

    // our list of strings to match
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered equivalent
    var toMatch = [][2]string{
        // this first set is the same that we use in the 'equals' operator
        // test above ... they're all expected to pass too
        [2]string{"<=1.3", "1.3"},
        [2]string{"<=1.3", "1.3.0"},
        [2]string{"<=1.3.0", "1.3.0"},
        [2]string{"<=1.3.1", "1.3.1"},
        [2]string{"<=2.6-alpha-1", "2.6-alpha-1"},
        [2]string{"<=2.6-alpha-1", "2.6.0-alpha-1"},
        [2]string{"<=2.5.99-alpha-1", "2.5.99-alpha-1"},
        [2]string{"<=2.5.99-ALPHA-1", "2.5.99-ALPHA-1"},
        [2]string{"<=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141013"},

        // this second set should only work when used with the '<=' operator
        [2]string{"<=1.3", "1.2"},
        [2]string{"<=1.3", "1.2.0"},
        [2]string{"<=1.3", "1.2.99"},
        [2]string{"<=1.3", "0.9"},
        [2]string{"<=1.3", "0.9.999"},
        [2]string{"<=1.3.0", "1.2"},
        [2]string{"<=1.3.0", "1.2.0"},
        [2]string{"<=1.3.0", "1.2.99"},
        [2]string{"<=1.3.0", "0.9"},
        [2]string{"<=1.3.0", "0.9.99"},
        [2]string{"<=1.3-alpha-2", "1.3-alpha-1"},
        [2]string{"<=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141012"},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet[0])
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet[1])

        // was an error returned?
        if err != nil {
            fmt.Println(lhs)
            fmt.Println(matchSet[1])
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != expected {
            t.Errorf("Expected %d, received %d", expected, actual)
            return
        }
    }
}

func TestCannotMatchUsingLessThanOrEquals(t *testing.T) {
    // our list of strings to compare
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered non-equivalent
    var toMatch = []ExpectedError{
        ExpectedError{"<=1.3", "2.0", ErrMajorVersionTooLarge},
        ExpectedError{"<=1.3", "1.4", ErrMinorVersionTooLarge},
        ExpectedError{"<=1.3.0", "2.3.0", ErrMajorVersionTooLarge},
        ExpectedError{"<=1.3.0", "1.4.0", ErrMinorVersionTooLarge},
        ExpectedError{"<=1.3.1", "1.3.2", ErrPatchLevelTooLarge},
        ExpectedError{"<=2.6-alpha-1", "2.6-alpha-2", ErrReleaseNumberTooLarge},
        ExpectedError{"<=2.6-alpha-1", "2.6.0-alpha-2", ErrReleaseNumberTooLarge},
        ExpectedError{"<=2.6-alpha-1", "2.6.0-beta-2", ErrDifferentStabilityLevels},
        ExpectedError{"<=2.6-alpha-1", "2.6.0", ErrDifferentStabilityLevels},
        ExpectedError{"<=2.6", "2.6.0-beta-2", ErrDifferentStabilityLevels},
        ExpectedError{"<=2.5.99-alpha-1", "2.5.99-alpha-2", ErrReleaseNumberTooLarge},
        ExpectedError{"<=2.5.99-ALPHA-1", "2.5.99-ALPHA-2", ErrReleaseNumberTooLarge},
        ExpectedError{"<=2.5.99-ALPHA-1", "2.5.99-BETA-2", ErrDifferentStabilityLevels},
        ExpectedError{"<=2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141014", ErrReleaseNumberTooLarge},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet.lhs)
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet.rhs)

        // was an error returned?
        if err != matchSet.err {
            fmt.Println(matchSet.lhs)
            fmt.Println(matchSet.rhs)
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != false {
            t.Errorf("Expected %d, received %d", false, actual)
            return
        }
    }
}

// ========================================================================
//
// Compare two versions using the tilde operator
//
// ------------------------------------------------------------------------

func TestCanMatchUsingTilde(t *testing.T) {
    // what result do we expect?
    expected := true

    // our list of strings to match
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered equivalent
    var toMatch = [][2]string{
        // this first set is the same that we use in the 'equals' operator
        // test above ... they're all expected to pass too
        [2]string{"~1.3", "1.3"},
        [2]string{"~1.3", "1.3.0"},
        [2]string{"~1.3.0", "1.3.0"},
        [2]string{"~1.3.1", "1.3.1"},
        [2]string{"~2.6-alpha-1", "2.6-alpha-1"},
        [2]string{"~2.6-alpha-1", "2.6.0-alpha-1"},
        [2]string{"~2.5.99-alpha-1", "2.5.99-alpha-1"},
        [2]string{"~2.5.99-ALPHA-1", "2.5.99-ALPHA-1"},
        [2]string{"~2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141013"},

        // this second set should only work when used with the '~' operator
        [2]string{"~1.3", "1.3.1"},
        [2]string{"~1.3", "1.4.0"},
        [2]string{"~1.3", "1.9.99"},
        [2]string{"~1.3.0", "1.3.1"},
        [2]string{"~1.3.0", "1.4.0"},
        [2]string{"~1.3.0", "1.9.99"},
        [2]string{"~1.3-alpha-2", "1.3-alpha-3"},
        [2]string{"~2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141014"},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet[0])
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet[1])

        // was an error returned?
        if err != nil {
            fmt.Println(lhs)
            fmt.Println(matchSet[1])
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != expected {
            t.Errorf("Expected %d, received %d", expected, actual)
            return
        }
    }
}

func TestCannotMatchUsingTilde(t *testing.T) {
    // our list of strings to compare
    //
    // LHS contains the operator
    // RHS contains only a version number to compare against
    //
    // all of these pairs should be considered non-equivalent
    var toMatch = []ExpectedError{
        ExpectedError{"~1.3", "2.0", ErrDifferentMajorVersions},
        ExpectedError{"~1.3", "1.2", ErrMinorVersionTooSmall},
        ExpectedError{"~1.3.0", "2.3.0", ErrDifferentMajorVersions},
        ExpectedError{"~1.3.0", "1.2.0", ErrMinorVersionTooSmall},
        ExpectedError{"~1.3.1", "1.3.0", ErrPatchLevelTooSmall},
        ExpectedError{"~2.6-alpha-2", "2.6-alpha-1", ErrReleaseNumberTooSmall},
        ExpectedError{"~2.6-alpha-2", "2.6.0-alpha-1", ErrReleaseNumberTooSmall},
        ExpectedError{"~2.6-alpha-1", "2.6.0-beta-2", ErrDifferentStabilityLevels},
        ExpectedError{"~2.6-alpha-1", "2.6.0", ErrDifferentStabilityLevels},
        ExpectedError{"~2.6", "2.6.1-alpha-1", ErrDifferentStabilityLevels},
        ExpectedError{"~2.5.99-alpha-2", "2.5.99-alpha-1", ErrReleaseNumberTooSmall},
        ExpectedError{"~2.5.99-ALPHA-2", "2.5.99-ALPHA-1", ErrReleaseNumberTooSmall},
        ExpectedError{"~2.5.99-ALPHA-1", "2.5.99-BETA-2", ErrDifferentStabilityLevels},
        ExpectedError{"~2.5.99-RC-1", "2.5.99", ErrDifferentStabilityLevels},
        ExpectedError{"~2.5.99-SNAPSHOT-20141013", "2.5.99-SNAPSHOT-20141012", ErrReleaseNumberTooSmall},
    }

    for _, matchSet := range toMatch {
        // perform the test
        lhs, err := Parse(matchSet.lhs)
        if err != nil {
            t.Error(err)
            return
        }
        actual, err := lhs.Matches(matchSet.rhs)

        // was an error returned?
        if err != matchSet.err {
            fmt.Println(matchSet.lhs)
            fmt.Println(matchSet.rhs)
            t.Error(err)
            return
        }

        // did we get back what we expected?
        if actual != false {
            t.Errorf("Expected %d, received %d", false, actual)
            return
        }
    }
}
