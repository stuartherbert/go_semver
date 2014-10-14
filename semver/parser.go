package semver

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

// VersionExpression holds the result of a parsed version expression
//
// create one by calling:
//
//     exp = semver.ParseExpression("<operator><version>")
type VersionExpression struct {
    Operator int        // which operator are we using?
    Version  SemVersion // which version is specified?
}

// value of VersionExpression.Operator when the expression requires an
// exact match
const OP_EQUALS = 0

// value of VersionExpression.Operator when the expression requires a
// version that is greater than or equal to
const OP_GT_EQUALS = 1

// value of VersionExpression.Operator when the expression requires a
// version that is less than or equal to
const OP_LT_EQUALS = 2

// value of VersionExpression.Operator when the expression requires a
// version that is both compatible (ie same major version) AND greater
// than or equal to
const OP_TILDE = 3

// value of VersionExpression.Operator when the expression requires a
// non-version string of some kind (such as a commit_id or a branch name)
//
// currently unsupported
const OP_AT = 4

// value of VersionExpression.Operator when the expression requires
// a version that does NOT equal
const OP_NOT_EQUALS = 5

// a list of supported operators
var opList = []string{"=", ">=", "<=", "~", "@", "!="}

// holds our compiled regexes, so that we don't have to compile them
// more than once
var versionRegexes [4]*regexp.Regexp

// compiles all of the regexes that we need to parse version strings
func init() {
    // here are the regex's that will match version strings
    var regexes [4]string
    regexes[0] = "(?P<Major>[0-9]+)\\.(?P<Minor>[0-9]+)\\.(?P<Patchlevel>[0-9]+)-(?P<Stability>[^-]+)-(?P<Release>[0-9]+)" // x.y.z-<stability>-r
    regexes[1] = "(?P<Major>[0-9]+)\\.(?P<Minor>[0-9]+)\\.(?P<Patchlevel>[0-9]+)"                                          // x.y.z
    regexes[2] = "(?P<Major>[0-9]+)\\.(?P<Minor>[0-9]+)-(?P<Stability>[^-]+)-(?P<Release>[0-9]+)"                          // x.y-<stability>-r
    regexes[3] = "(?P<Major>[0-9]+)\\.(?P<Minor>[0-9]+)$"                                                                  // x.y

    // compile the regexes to use later
    for i, regex := range regexes {
        versionRegexes[i] = regexp.MustCompile(regex)
    }
}

// ParseExpression converts a version expression string into a
// VersionExpression struct.
//
// Takes an expression of the form:
//
//     <OPERATOR><version-string>
//
// and turns it into a VersionExpression struct
func ParseExpression(exp string) (VersionExpression, error) {
    // do we have an operator?
    op, offset, err := startsWithOperator(exp)
    if err != nil {
        return VersionExpression{}, err
    }

    // do we have a semantically-correct version number too?
    version, err := parseVersionWithOffset(exp, offset)
    if err != nil {
        return VersionExpression{}, err
    }

    parsed := VersionExpression{op, version}
    return parsed, nil
}

func startsWithOperator(raw string) (int, int, error) {
    for i, opToEval := range opList {
        if strings.HasPrefix(raw, opToEval) {
            return i, len(opToEval), nil
        }
    }

    // if we get here, then we cannot decode the string
    return -1, -1, fmt.Errorf("unrecognised operator")
}

// ParseVersion takes a version string and turns it into a SemVersion
// struct.
//
// Takes any of these strings:
//
//     X.Y
//     X.Y.Z
//     X.Y-<stability>-R
//     X.Y.Z-<stability>-R
//
// and turns it into a SemVersion struct
func ParseVersion(version string) (SemVersion, error) {
    return parseVersionWithOffset(version, 0)
}

func parseVersionWithOffset(raw string, offset int) (SemVersion, error) {
    for _, re := range versionRegexes {
        matches := re.FindStringSubmatch(raw)
        if len(matches) == 0 {
            continue
        }

        // store the named results
        capture := make(map[string]string)
        for i, name := range re.SubexpNames() {
            if i == 0 || name == "" {
                continue
            }
            capture[name] = matches[i]
        }

        // build our return value
        version := SemVersion{}

        // Major and Minor versions are present in all of the regexes
        // that we use
        version.Major, _ = strconv.Atoi(capture["Major"])
        version.Minor, _ = strconv.Atoi(capture["Minor"])

        // the remaining elements are optional
        if capture["Patchlevel"] != "" {
            version.PatchLevel, _ = strconv.Atoi(capture["Patchlevel"])
        }
        if capture["Stability"] != "" {
            version.Stability = capture["Stability"]
        }
        if capture["Release"] != "" {
            version.Release, _ = strconv.Atoi(capture["Release"])
        }

        return version, nil
    }

    return SemVersion{}, fmt.Errorf("don't know how to interpret matches yet")
}
