// Package capabilities performs operations on maps representing WebDriver capabilities.
package capabilities

// Merge takes two JSON objects, and merges them.
//
// The resulting object will have all of the keys in the two input objects.
// For each key that is in both obejcts:
//   - if both objects have objects for values, then the result object will have
//     a value resulting from recursively calling Merge.
//   - if both objects have lists for values, then the result object will have
//     a value resulting from concatenating the two lists.
//   - Otherwise the result object will have the value from the second object.
func Merge(m1, m2 map[string]interface{}) map[string]interface{} {
	if len(m1) == 0 {
		return m2
	}
	if len(m2) == 0 {
		return m1
	}
	nm := make(map[string]interface{})
	for k, v := range m1 {
		nm[k] = v
	}
	for k, v2 := range m2 {
		nm[k] = mergeValues(nm[k], v2)
	}
	return nm
}

func mergeValues(j1, j2 interface{}) interface{} {
	switch t1 := j1.(type) {
	case map[string]interface{}:
		if t2, ok := j2.(map[string]interface{}); ok {
			return Merge(t1, t2)
		}
	case []interface{}:
		if t2, ok := j2.([]interface{}); ok {
			return mergeLists(t1, t2)
		}
	}
	return j2
}

func mergeLists(m1, m2 []interface{}) []interface{} {
	if len(m1) == 0 {
		return m2
	}
	if len(m2) == 0 {
		return m1
	}
	nl := make([]interface{}, 0, len(m1)+len(m2))
	nl = append(nl, m1...)
	nl = append(nl, m2...)
	return nl
}
