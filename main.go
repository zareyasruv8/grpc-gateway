package main

import (
	"fmt"
	"sort"
)

// ExtractFieldMaskPaths recursively traverses a map to generate leaf-level FieldMask paths.
func ExtractFieldMaskPaths(prefix string, val map[string]interface{}) []string {
	var paths []string
	for k, v := range val {
		path := k
		if prefix != "" {
			path = prefix + "." + k
		}

		if nestedMap, ok := v.(map[string]interface{}); ok && len(nestedMap) > 0 {
			paths = append(paths, ExtractFieldMaskPaths(path, nestedMap)...)
		} else {
			paths = append(paths, path)
		}
	}
	sort.Strings(paths)
	return paths
}

func main() {
	// Example usage demonstrating the fix
	payload := map[string]interface{}{
		"foo": "bar",
		"parent": map[string]interface{}{
			"child_a": "value",
		},
	}
	paths := ExtractFieldMaskPaths("", payload)
	fmt.Printf("Generated paths: %v\n", paths)
}