package instructions

type DescriptorVersion int

const (
	VersionBasic = iota
	VersionArrayDic
)

type VariableDescriptor struct {
	Name            string
	Version         DescriptorVersion
	IsInternal      bool
	InnerDescriptor *VariableDescriptor
}
