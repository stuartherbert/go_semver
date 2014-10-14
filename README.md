# go-semver

A simple library to compare two semantic version numbers against each other. Written for Go.

[![GoDoc](https://godoc.org/github.com/stuartherbert/go_semver/server?status.svg)](http://godoc.org/github.com/stuartherbert/go_semver/semver)
[![Build Status](https://travis-ci.org/stuartherbert/go_semver.svg?branch=master)](https://travis-ci.org/stuartherbert/go_semver)

## Semantic Version Strings

A semantic version string is of the form:

    X.Y.Z-<stability>-R

where:

* __X__ is the major version number. This is always incremented when backwards-compatibility is broken, or when you've added major new features.
* __Y__ is the minor version number. When this is incremented, Z gets reset back to 0.
* __Z__ is the patch level. This is incremented by 1 when you put out a bugfix release.
* __<stability>__ is one of: alpha, beta, pre, snapshot, dev, rc. They are used to indicate experimental releases that have been made to share work in progress. 'Stable' releases do not include a 'stability' level in the version number.
* __R__ is the release number used to tell different unstable releases apart. It is an integer.

### Stability Levels

Some package managers (notably RPM) treat the 'stability' (or 'release' field in RPM terms) simply as an ASCII string, with no knowledge of what the field means.  In practice, this is unhelpful when trying to upgrade from one version of a package to another.

When making an unstable release, you can actually use any [A-Za-z0-9_] string you like in the 'stability' field. We recommend that you use one of:

* snapshot, dev
* alpha
* beta
* rc, pre

Go-semver preserves case, and internally treats the 'stability' level as case-insensitive when making comparisons.

## Branch Names / Commit IDs

Go-semver also supports branch names and commit IDs (such as treeish used in Github) as a special case. These can only be used with the special '@' comparison operator.

## Comparison Operators

Go-semver understands the following comparison operators:

* =X.Y.Z - only exact version
* >=X[.Y[.Z]] - all versions from X.[Y[.Z]]
* <=X[.Y[.Z]] - highest version below X[.Y[.Z]]
* ~X[.Y] - equivalent to '>= X[.Y], <X+1.0'
* @<branch|commit_id> - only the branch or commit_id specified

## Comparison Of Unstable Releases

Go-semver supports:

* comparison between two releases with the same stability level that share the same X.Y.Z version
* comparison between an unstable release and a stable release that share the same major release number (X value)

Any other releases with _different_ stability levels are considered incomparable, because Go-semver doesn't assign any meaning to the contents of the 'stability' string.  There are just too many values in real-world use for Go-semver to try and put 'stability' levels in any sort of order.

For example, you can compare:

* '1.0.0-alpha-1', '1.0.0-alpha-5', and '1.0.0' successfully

But you can't compare:

* '1.0.0-alpha-1' and '1.0.0-beta-1' - returns ErrDifferentUnstable
