package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type BasicHs struct {
	NameSingle, NamePlural string
}

var packages = []*BasicHs{
	{"company", "companies"},
	{"deal", "deals"},
	{"product", "products"},
	{"contact", "contacts"},
	{"lead", "leads"},
	{"line_item", "line_items"},
}

const BASE_PATH = "/crm/v3/objects"

func main() {

	var files []*File

	for _, hs := range packages {
		f := NewFile("hubspot")

		f.Comment("Generated. Don't change.")
		f.Line()

		f.ImportName("fmt", "fmt")

		f.Comment(fmt.Sprintf("GET %s", path(hs.NamePlural)))
		methodBase(f).Id(getName(hs.NamePlural)).Params(
			Id("p").Id("*ListParams"),
		).Params(
			Id("*PaginatedResponse"),
			Error(),
		).Block(
			pathVar(hs.NamePlural),
			Return(Id("c").Dot("list").Call(Id("path"), Id("p"))),
		).Line()

		f.Comment(fmt.Sprintf("GET %s/{id}", path(hs.NamePlural)))
		methodBase(f).Id(getName(hs.NameSingle)).Params(
			Id("id").String(),
			Id("p").Id("*ReadParams"),
		).Params(
			Id("*HsObject"),
			Error(),
		).Block(
			idPathVar(hs.NamePlural),
			Return(Id("c").Dot("get").Call(Id("path"), Id("p"))),
		).Line()

		f.Comment(fmt.Sprintf("POST %s", path(hs.NamePlural)))
		methodBase(f).Id(createName(hs.NameSingle)).Params(
			Id("body").Id("*CreateBody"),
		).Params(
			Id("*HsObject"),
			Error(),
		).Block(
			pathVar(hs.NamePlural),
			Return(Id("c").Dot("create").Call(Id("path"), Id("body"))),
		).Line()

		f.Comment(fmt.Sprintf("PATCH %s/{id}", path(hs.NamePlural)))
		methodBase(f).Id(updateName(hs.NameSingle)).Params(
			Id("id").String(),
			Id("properties").Id("map[string]string"),
		).Params(
			Id("*HsObject"),
			Error(),
		).Block(
			pathVar(hs.NamePlural),
			Return(Id("c").Dot("update").Call(Id("path"), Id("properties"))),
		).Line()

		f.Comment(fmt.Sprintf("DELETE %s/{id}", path(hs.NamePlural)))
		methodBase(f).Id(deleteName(hs.NameSingle)).Params(
			Id("id").String(),
		).Error().Block(
			idPathVar(hs.NamePlural),
			Return(Id("c").Dot("delete").Call(Id("path"))),
		).Line()

		buf := bytes.Buffer{}
		f.Render(&buf)

		os.WriteFile(fmt.Sprintf("%s_gen.go", hs.NameSingle), buf.Bytes(), 0646)
		files = append(files, f)
	}

}

func path(name string) string {
	return fmt.Sprintf("%s/%s", BASE_PATH, name)
}
func idPath(name string) string {
	return path(name) + "/%s"
}

// return a path variables a string
func pathVar(name string) *Statement {
	path := path(name)

	c := Id("path").
		Op(":=").Lit(path)
	return c
}

// return a path variable with fmt.Sprintf for adding id
func idPathVar(name string) *Statement {
	path := idPath(name)

	c := Id("path").
		Op(":=").
		Qual("fmt", "Sprintf").
		Call(Lit(path), Id("id"))

	return c
}

func methodBase(f *File) *Statement {
	return f.Func().Params(Id("c").Id("*Client"))
}

func getName(name string) string {
	return fmt.Sprintf("Get%s", capitalize(toCamelCase(name)))
}

func createName(name string) string {
	return fmt.Sprintf("Create%s", capitalize(toCamelCase(name)))
}

func updateName(name string) string {
	return fmt.Sprintf("Update%s", capitalize(toCamelCase(name)))
}

func deleteName(name string) string {
	return fmt.Sprintf("Delete%s", capitalize(toCamelCase(name)))
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	sr := []rune(s)
	sr[0] = []rune(strings.ToUpper(string(sr[0])))[0]

	return string(sr)
}

func toCamelCase(input string) string {
	out := []rune{}

	var prevChar rune
	for _, char := range strings.ToLower(input) {
		if strings.ContainsRune("-_", char) {
			prevChar = char
			continue
		}

		if strings.ContainsRune("-_", prevChar) {
			r := []rune(strings.ToUpper(string(char)))[0]
			out = append(out, r)
		} else {
			out = append(out, char)
		}

		prevChar = char
	}
	_ = prevChar

	return string(out)
}
