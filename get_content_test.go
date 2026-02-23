package main

import (
		"testing"
		"log"
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