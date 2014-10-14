package semver

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

type Comparison struct {
    Operator int        // which operator are we using?
    Version  SemVersion // which version is specified?
}

const OP_EQUALS = 0
const OP_GT_EQUALS = 1
const OP_LT_EQUALS = 2
const OP_TILDE = 3
const OP_AT = 4

var OpList = []string{"=", ">=", "<=", "~", "@"}

var versionRegexes [4]*regexp.Regexp

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

func Parse(raw string) (Comparison, error) {
    // do we have an operator?
    op, offset, err := startsWithOperator(raw)
    if err != nil {
        return Comparison{}, err
    }

    // do we have a semantically-correct version number too?
    version, err := parseVersionWithOffset(raw, offset)
    if err != nil {
        return Comparison{}, err
    }

    parsed := Comparison{op, version}
    return parsed, nil
}

func startsWithOperator(raw string) (int, int, error) {
    for i, opToEval := range OpList {
        if strings.HasPrefix(raw, opToEval) {
            return i, len(opToEval), nil
        }
    }

    // if we get here, then we cannot decode the string
    return -1, -1, fmt.Errorf("unrecognised operator")
}

func ParseVersion(raw string) (SemVersion, error) {
    return parseVersionWithOffset(raw, 0)
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
