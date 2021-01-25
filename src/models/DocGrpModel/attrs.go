package DocGrpModel

type DocGrpModelAttrFunc func(d *DocGrpImpl)

type DocGrpModelAttrFuncs []DocGrpModelAttrFunc

func WithGroupID(id int) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.GroupID = id
	}
}

func WithGroupName(groupName string) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.GroupName = groupName
	}
}

func (this DocGrpModelAttrFuncs) Apply(d *DocGrpImpl) {
	for _, f := range this {
		f(d)
	}
}
