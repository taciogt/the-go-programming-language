package main

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"reflect"
	"strings"
	"testing"
)

func TestElementByID(t *testing.T) {
	type args struct {
		doc string
		id  string
	}
	tests := []struct {
		name string
		args args
		want html.Node
	}{{
		name: "empty html",
		args: args{
			doc: "<html></html>",
		},
		want: html.Node{},
	}, {
		name: "html with id to find",
		args: args{
			doc: "<html><body><h1 id=\"some-id\">The Page Header</h1></body></html>",
			id:  "some-id",
		},
		want: html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.H1,
			Data:     "h1",
			Attr: []html.Attribute{html.Attribute{
				Key: "id",
				Val: "some-id",
			}},
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(tt.args.doc))
			if err != nil {
				t.Error(err)
				return
			}
			got := ElementByID(doc, tt.args.id)
			simplifiedGot := html.Node{
				Type:      got.Type,
				DataAtom:  got.DataAtom,
				Data:      got.Data,
				Namespace: got.Namespace,
				Attr:      got.Attr,
			}
			if !reflect.DeepEqual(simplifiedGot, tt.want) {
				t.Errorf("ElementByID() = \n%+v, want \n%+v", simplifiedGot, tt.want)
			}
		})
	}
}
