package main

import (
		"testing"
	)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove last slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove multiple slashes",
			inputURL: "https://blog.boot.dev/path///",
			expected: "blog.boot.dev/path",
		},
        // add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}


func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputBody      string
		expected      string
	}{
		{
			name:     "1",
			inputBody: `<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>`,
			expected: "Welcome to Boot.dev",
		},
		{
			name:     "2",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "3 [Empty]",
			inputBody: "Test Title",
			expected: "",
		},
        // add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputBody)
			
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputBody      string
		expected      string
	}{
		{
			name:     "1",
			inputBody: `<html><body>
	<p>Outside paragraph.</p>
	<main>
		<p>Main paragraph.</p>
	</main>
</body></html>`,
			expected: "Outside paragraph.",
		},
		{
			name:     "2",
			inputBody: "<html><body><h1>Test Title</h1><p>Main paragraph.</p></body></html>",
			expected: "Main paragraph.",
		},
		{
			name:     "3 [Empty]",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected: "",
		},
        // add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL:\n\t expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
	
}
func TestGetFirstMainParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputBody      string
		expected      string
	}{
		{
			name:     "1",
			inputBody: `<html><body>
	<p>Outside paragraph.</p>
	<main>
		<p>Main paragraph.</p>
		<p>Other paragraph.</p>
		<p>Last paragraph.</p>
	</main>
</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name:     "2 [No Main]",
			inputBody: "<html><body><h1>Test Title</h1><p>Main paragraph.</p></body></html>",
			expected: "",
		},
		{
			name:     "3 [No p in Main]",
			inputBody: "<html><body><h1>Test Title</h1><main><h2>ciao</h2></main></body></html>",
			expected: "",
		},
        // add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstMainParagraphFromHTML(tc.inputBody)
			
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL:\n\t expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
	
}