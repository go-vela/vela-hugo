// SPDX-License-Identifier: Apache-2.0

package main

// Build represents the plugin configuration for Hugo information.
type Build struct {
	// hostname (and path) to the root, e.g. http://spf13.com/
	BaseURL string
	// include content marked as draft
	Draft bool
	// include expired content
	Expired bool
	// include content with publishdate in the future
	Future bool
}
