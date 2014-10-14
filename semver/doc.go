// Package semver is a version string parsing and comparison library
//
// Versions
//
// The semver package allows you to parse software version strings of
// the form:
//
//    X.Y.Z-stability-R
//
// where:
//
//     'X' is the major version number
//     'Y' is the minor version number
//     'Z' is the patch level
//     'stability' indicates what kind of unstable release this is
//     'R' is the release number for unstable releases
//
// Notes:
//
//     The 'Z' is optional (defaults to 0 if missing)
//     A version string is assumed 'stable' if the stability is missing
//
// Example version numbers include:
//
//     0.1
//     1.0-alpha-1
//     1.0.0-alpha-1
//     1.0
//     1.0.0
//     1.1.0-SNAPSHOT-20141013
//
// Comparisons
//
// The semver package also includes support for comparing two version
// strings against each other, using the following operators:
//
//     =  : requires exact match
//     >= : any version that's greater than or equal to
//     <= : any version that's less than or equal to
//     ~  : any version that's greater than or equal to, and has the same
//          major version number
//
// For example:
//
//     =1.3 : matches '1.3' and '1.3.0'
//     =1.3.0 : matches '1.3' and '1.3.0'
//     =1.3.0-alpha-1 : matches '1.3-alpha-1' and '1.3.0-alpha-1'
//
//     >=1.3 : matches '1.3', '1.3.0', '1.3.1', '1.4.0', '2.0.0' and so on
//     >=1.3.0 : matches '1.3', '1.3.0', '1.3.1', '1.4.0', '2.0.0' and
//               so on
//     >=1.3.0-alpha-1 : matches '1.3-alpha-1', '1.3.0-alpha-1',
//                      '1.3.0-alpha-2' and so on
//
//     <=1.3 : matches '1.3', '1.3.0', '1.2.99999', '1.1.0', '0.9' and
//             so on
//     <=1.3.0 : matches '1.3', '1.3.0', '1.2.99999', '1.1.0', '0.9' and
//               so on
//     <=1.3.0-alpha-2 : matches '1.3-alpha-2', '1.3.0-alpha-2',
//                       '1.3.0-alpha-1' only
//
//     ~1.3 : matches '1.3', '1.3.0', '1.3.1', '1.4.0', and all newer
//            stable 1.x releases
//     ~1.3.0 : matches '1.3', '1.3.0', '1.3.1', '1.4.0', and all newer
//              stable 1.x releases
//     ~1.3.0-alpha-1 : matches '1.3-alpha-1', '1.3.0-alpha-1' and all
//                      newer alpha releases of 1.3.0
//
// You can only compare like releases with like:
//
//     stable releases can only be compared with other stable releases
//
//     any unstable release can only be compared with another release with
//     the same stability level AND same X.Y.Z
//
// The semver API returns meaningful errors when a comparison fails,
// explaining exactly why two version strings are different or can't be
// compared.
package semver
