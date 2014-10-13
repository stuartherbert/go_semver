package semver

// SemVersion holds the structure of a version, in the form
//
// X.Y.Z-<stability>R
type SemVersion struct {
    Major      int    // X
    Minor      int    // Y
    PatchLevel int    // Z
    Stability  string // stability
    Release    int    // R
}
