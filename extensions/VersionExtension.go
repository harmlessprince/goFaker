package extensions

type VersionExtension interface {
	Semver(params ...bool)
}
