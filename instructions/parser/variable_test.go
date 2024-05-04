package parser

import (
	"esorun/variable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVariable(t *testing.T) {
	var descriptor *variable.Descriptor
	var err error

	descriptor, err = ParseVariable([]byte("@normalVar"))
	assert.Equal(t, "normalVar", descriptor.Name)
	assert.False(t, descriptor.IsInternal)
	assert.Equal(t, variable.DescriptorVersion(variable.VersionBasic), descriptor.Version)
	assert.Nil(t, descriptor.InnerDescriptor)
	assert.Nil(t, err)

	descriptor, err = ParseVariable([]byte("@@normalVar"))
	assert.Equal(t, "normalVar", descriptor.Name)
	assert.True(t, descriptor.IsInternal)
	assert.Equal(t, variable.DescriptorVersion(variable.VersionBasic), descriptor.Version)
	assert.Nil(t, descriptor.InnerDescriptor)
	assert.Nil(t, err)

	descriptor, err = ParseVariable([]byte("@arrVar[@normalVar]"))
	assert.Equal(t, "arrVar", descriptor.Name)
	assert.False(t, descriptor.IsInternal)
	assert.Equal(t, variable.DescriptorVersion(variable.VersionArrayDic), descriptor.Version)
	assert.Equal(t, "normalVar", descriptor.InnerDescriptor.Name)
	assert.False(t, descriptor.InnerDescriptor.IsInternal)
	assert.Equal(t, variable.DescriptorVersion(variable.VersionBasic), descriptor.InnerDescriptor.Version)
	assert.Nil(t, err)

	descriptor, err = ParseVariable([]byte("@@arrVar[@@normalVar]"))
	assert.Equal(t, "arrVar", descriptor.Name)
	assert.True(t, descriptor.IsInternal)
	assert.Equal(t, variable.DescriptorVersion(variable.VersionArrayDic), descriptor.Version)
	assert.Equal(t, "normalVar", descriptor.InnerDescriptor.Name)
	assert.True(t, descriptor.InnerDescriptor.IsInternal)
	assert.Equal(t, variable.DescriptorVersion(variable.VersionBasic), descriptor.InnerDescriptor.Version)
	assert.Nil(t, err)

	descriptor, err = ParseVariable([]byte("normalVar"))
	assert.Nil(t, descriptor)
	assert.Error(t, err)
}
