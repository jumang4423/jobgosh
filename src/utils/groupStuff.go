package utils

import types "jobgosh/src/typesAndVars"

// add group workspace
func AddAsGroupMember(_Group string, directoryName string, _allGroups []types.Group) {

	for i, v := range _allGroups {

		if v.Group == _Group {
			// add
			_allGroups[i].WorkSpace = append(v.WorkSpace, directoryName)
		}
	}

	SaveGroupJson(_allGroups)
}

// find group in group.json
func FindGroup(directoryName string) string {

	var group string = ""
	var _allGroups []types.Group
	_allGroups = LoadGroupJson()

	for _, v := range _allGroups {

		// find workspace
		for _, r := range v.WorkSpace {
			if r == directoryName {
				group = v.Group
			}
		}

	}

	return group
}