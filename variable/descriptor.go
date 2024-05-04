package variable

type DescriptorVersion int

const (
	VersionBasic = iota
	VersionArray
	VersionDic
)

type Descriptor struct {
	name       string
	version    DescriptorVersion
	isInternal bool
	descriptor *Descriptor
}
