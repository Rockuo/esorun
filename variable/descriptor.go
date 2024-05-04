package variable

type DescriptorVersion int

const (
	VersionBasic = iota
	VersionArrayDic
)

type Descriptor struct {
	Name            string
	Version         DescriptorVersion
	IsInternal      bool
	InnerDescriptor *Descriptor
}
