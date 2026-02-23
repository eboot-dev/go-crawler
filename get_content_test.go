package main

import (
		"testing"
		"log"
		"net/url"
		"reflect"
		"strings"
	)

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

	log.Println("Test getH1FromHTML")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputBody)
			
			if actual != tc.expected {
				t.Errorf("\tTest %s FAIL\n\t\t expected: %v, actual: %v", tc.name, tc.expected, actual)
				return
			}
			log.Printf("\tTest %s PASSED\n",tc.name)
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
			name:     "2 [Empty Tag/Only spaces]",
			inputBody: "<html><body><h1>Test Title</h1><p>  </p></body></html>",
			expected: "",
		},
		{
			name:     "3 [Empty]",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected: "",
		},
        // add more test cases here
	}

	log.Println("Test getFirstParagraphFromHTML")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			
			if actual != tc.expected {
				t.Errorf("\tTest %s FAIL\n\t\t expected: '%v', actual: '%v'", tc.name, tc.expected, actual)
				return
			}
			log.Printf("\tTest %s PASSED\n",tc.name)
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
		},{
			name:     "4 [Multple Main]",
			inputBody: `<html><body>
	<p>Outside paragraph.</p>
	<main>
		<p>Main paragraph.</p>
		<p>A paragraph.</p>
	</main>
	<main>
		<p>Main2 paragraph.</p>
		<p>A paragraph 2.</p>
	</main>
</body></html>`,
			expected: "Main paragraph.",
		},
        // add more test cases here
	}

	log.Println("Test getFirstMainParagraphFromHTML")
	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstMainParagraphFromHTML(tc.inputBody)
			
			if actual != tc.expected {
				t.Errorf("\tTest %s FAIL:\n\t\t expected: '%v', actual: '%v'", tc.name, tc.expected, actual)
				return
			}
			
			log.Printf("\tTest %s PASSED\n",tc.name)
		})
		
	}
	
}

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="https://blog.boot.dev"><span>Boot.dev</span></a></body></html>`
	
	log.Println("Test getURLsFromHTML - Absolute")
  baseURL, err := url.Parse(inputURL)
  if err != nil {
      t.Errorf("\tcouldn't parse input URL: %v", err)
      return
  }

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("\tunexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\texpected %v, got %v", expected, actual)
		return
	}
	log.Printf("\tTest %s PASSED\n","TestGetURLsFromHTMLAbsolute")
}

func TestGetImagesFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="/logo.png" alt="Logo"></body></html>`
	log.Println("Test getImagesFromHTML - Relative")
    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("\tunexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\texpected %v, got %v", expected, actual)
		return
	}
	log.Printf("\tTest %s PASSED\n","TestGetImagesFromHTMLRelative")
}


func TestGetURLsFromHTMLFull(t *testing.T) {
	log.Println("Test getURLsFromHTML - Full")
	cases := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://blog.boot.dev">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "relative URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
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
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev</span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)

			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("\tTest %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("\tTest %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("\tTest %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\tTest %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tc.name, tc.expected, actual)
				return
			}
			log.Printf("\tTest %s PASSED\n","TestGetURLsFromHTMLFull")
		})
	}
}