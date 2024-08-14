package shared

type DrawStyle int

const (
	RegularStyle DrawStyle = iota
	DrawnStyle
)

type BackgroundFormat int

const (
	SolidBackground BackgroundFormat = iota
	ImageBackground
)

type SkeletonType int

const (
	RegularSkeletonType SkeletonType = iota
	GridSkeletonType
)

type Platform int

const (
	PlatformPC Platform = iota
	PlatformPS
	PlatformXBOX
)
