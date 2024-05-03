package value

type Type interface {
	GetName() TypeName
	IsValid() bool
}

type BaseType struct {
	Name TypeName
}

func (st BaseType) GetName() TypeName {
	return st.Name
}

func (st BaseType) IsValid() bool {
	return st.Name != TypeNameArray && st.Name != TypeNameDictionary
}

type ArrayType struct {
	BaseType
	itemType Type
}

func (at ArrayType) getName() string {
	return string(at.Name)
}

func (at ArrayType) IsValid() bool {
	return at.Name == TypeNameArray && at.itemType.IsValid()
}

type DictionaryType struct {
	BaseType
	itemType Type
}

func (dt DictionaryType) getName() string {
	return string(dt.Name)
}

func (dt DictionaryType) IsValid() bool {
	return dt.Name == TypeNameDictionary &&
		dt.itemType.IsValid()
}
