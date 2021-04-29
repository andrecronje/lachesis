package toml

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
)

func ParseFile(fileName string) (*ast.Table, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	res, err := toml.Parse(content)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func WriteTo(t *ast.Table, w io.Writer) error {
	return writeTo("", t, w)
}

func writeTo(parent string, t *ast.Table, w io.Writer) (err error) {
	var fullName string
	if parent != "" {
		fullName = parent + "." + t.Name
	} else {
		fullName = t.Name
	}

	if fullName != "" {
		switch t.Type {
		case ast.TableTypeNormal:
			_, err = fmt.Fprintf(w, "\n[%s]\n", fullName)
		case ast.TableTypeArray:
			_, err = fmt.Fprintf(w, "\n[[%s]]\n", fullName)
		}
	}
	if err != nil {
		return
	}

	for _, v := range t.Fields {
		switch x := v.(type) {
		case *ast.KeyValue:
			fmt.Fprintf(w, "%s = %s\n", x.Key, x.Value.Source())
		default:
			// skip
		}
		if err != nil {
			return
		}
	}

	for _, v := range t.Fields {
		switch x := v.(type) {
		case *ast.KeyValue:
			// skip
		case *ast.Table:
			err = writeTo(fullName, x, w)
			if err != nil {
				return
			}
		case []*ast.Table:
			for _, f := range x {
				err = writeTo(fullName, f, w)
				if err != nil {
					return
				}
			}
		default:
			panic(x)
		}
	}

	return
}
