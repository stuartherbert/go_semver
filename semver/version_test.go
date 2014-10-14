package semver

import (
    "fmt"
    "testing"
)

type VersionExpectedResult struct {
    lhs      string
    rhs      string
    expected int
}

// ========================================================================
//
// Compare two versions using the Compare method
//
// ------------------------------------------------------------------------

func TestCanCompareTwoVersions(t *testing.T) {
    // our list of things to compare
    var toCompareList = []VersionExpectedResult{
        // things that should be the same
        VersionExpectedResult{"0.1", "0.1", COMP_EQUAL},
        VersionExpectedResult{"0.1", "0.1.0", COMP_EQUAL},
        VersionExpectedResult{"1.0", "1.0", COMP_EQUAL},
        VersionExpectedResult{"1.0", "1.0.0", COMP_EQUAL},
        VersionExpectedResult{"1.1", "1.1", COMP_EQUAL},
        VersionExpectedResult{"1.1", "1.1.0", COMP_EQUAL},
        VersionExpectedResult{"0.0.1", "0.0.1", COMP_EQUAL},
        VersionExpectedResult{"0.0.1-alpha-1", "0.0.1-alpha-1", COMP_EQUAL},

        // things that should be larger on the RHS
        VersionExpectedResult{"1.0", "1.1", COMP_LARGER},
        VersionExpectedResult{"1.0.0", "1.1.0", COMP_LARGER},
        VersionExpectedResult{"1.1.0", "1.1.1", COMP_LARGER},
        VersionExpectedResult{"1.1.0", "1.11.0", COMP_LARGER},
        VersionExpectedResult{"0.0.1", "0.0.2", COMP_LARGER},
        VersionExpectedResult{"0.0.1-alpha-1", "0.0.1-alpha-2", COMP_LARGER},

        // things that should be smaller on the RHS
        VersionExpectedResult{"0.0.2", "0.0.1", COMP_SMALLER},
        VersionExpectedResult{"1.1.0", "1.0.0", COMP_SMALLER},
        VersionExpectedResult{"1.1.1", "1.1.0", COMP_SMALLER},
        VersionExpectedResult{"11.0.0", "1.1.0", COMP_SMALLER},
        VersionExpectedResult{"0.0.1-alpha-2", "0.0.1-alpha-1", COMP_SMALLER},

        // things that make no sense to compare
        VersionExpectedResult{"0.0.1-alpha-1", "0.0.2-alpha-2", COMP_APPLES_AND_ORANGES},
        VersionExpectedResult{"0.0.1-alpha-1", "0.0.2-beta-1", COMP_APPLES_AND_ORANGES},
        VersionExpectedResult{"0.0.1", "0.0.2-beta-1", COMP_APPLES_AND_ORANGES},
        VersionExpectedResult{"0.0.1-beta-1", "0.0.2", COMP_APPLES_AND_ORANGES},
    }

    for _, toCompare := range toCompareList {
        // compile the lhs
        lhs, err := ParseVersion(toCompare.lhs)
        if err != nil {
            t.Error(err)
            return
        }
        rhs, err := ParseVersion(toCompare.rhs)
        if err != nil {
            t.Error(err)
            return
        }
        actual := lhs.Compare(&rhs)

        // what happened?
        if actual != toCompare.expected {
            fmt.Println("lhs: %s; rhs: %s", toCompare.lhs, toCompare.rhs)
            t.Errorf("expected: %d; actual: %d", toCompare.expected, actual)
            return
        }
    }
}
