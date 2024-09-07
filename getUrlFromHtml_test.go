package main

import (
	"reflect"
	"testing"
)

func TestGetUrlFromHtml(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no URLs in HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<p>No links here</p>
	</body>
</html>
`,
			expected: []string{},
		},
		{
			name:     "relative URL with missing base",
			inputURL: "",
			inputBody: `
<html>
	<body>
		<a href="/relative/path">Link</a>
	</body>
</html>
`,
			expected: []string{"/relative/path"},
		},
		{
			name:     "malformed HTML with valid link",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one"><span>Boot.dev</span></a
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "multiple relative links",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a href="/path/one">Link One</a>
		<a href="/path/two">Link Two</a>
		<a href="/path/three">Link Three</a>
	</body>
</html>
`,
			expected: []string{
				"https://example.com/path/one",
				"https://example.com/path/two",
				"https://example.com/path/three",
			},
		},
		{
			name:     "href attribute missing",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a>Missing href</a>
	</body>
</html>
`,
			expected: []string{},
		},
		{
			name:      "empty body",
			inputURL:  "https://example.com",
			inputBody: "",
			expected:  []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			// Normalize nil slice to empty slice for comparison
			if actual == nil {
				actual = []string{}
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
