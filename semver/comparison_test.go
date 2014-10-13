package semver

import (
    "fmt"
    "testing"
)

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
    // LHS contains the equals operator
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
    // LHS contains the equals operator
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
