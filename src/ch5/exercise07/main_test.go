package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHtmlPrettyPrinter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		wantW string
	}{{
		name:  "empty html",
		input: "<html></html>",
		wantW: "<html>\n" +
			"  <head>\n" +
			"  </head>\n" +
			"  <body>\n" +
			"  </body>\n" +
			"</html>\n",
	}, {
		name:  "html with content",
		input: "<html><body><h1>The Page Header</h1></body></html>",
		wantW: "<html>\n" +
			"  <head>\n" +
			"  </head>\n" +
			"  <body>\n" +
			"    <h1>The Page Header</h1>\n" +
			"  </body>\n" +
			"</html>\n",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			r := strings.NewReader(tt.input)
			HtmlPrettyPrinter(r, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("HtmlPrettyPrinter() = \n%v\n, want = \n%v", gotW, tt.wantW)

			}
		})
	}
}
