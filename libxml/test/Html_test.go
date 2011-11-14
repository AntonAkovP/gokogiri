package test

import (
	"libxml"
	"libxml/help"
	"testing"
	"strings"
)

func TestHtmlSimpleParse(t *testing.T) {
	doc := libxml.HtmlParseString("<html><head /><body /></html>")
	if doc.Size() != 1 {
		t.Error("Incorrect size")
	}
	// Doctype gets returned as the first child!
	htmlTag := doc.First().Next()
	if htmlTag.Size() != 2 {
		print(htmlTag.Name())
		t.Error("Two tags are inside of <html>")
	}
	doc.Free()
	help.XmlCleanUpParser()
	if help.XmlMemoryAllocation() != 0 {
		t.Errorf("Memeory leaks %d!!!", help.XmlMemoryAllocation())
		help.XmlMemoryLeakReport()
	}
}

func TestHtmlCDataTag(t *testing.T) {
	doc := libxml.HtmlParseString(LoadFile("docs/script.html"))
	if doc.Size() != 1 {
		t.Error("Incorrect size")
	}
	scriptTag := doc.RootElement().FirstElement().FirstElement()
	if scriptTag.Name() != "script" {
		t.Error("Should have selected the script tag")
	}
	content := scriptTag.Content()
	scriptTag.SetContent(content)
	doc.Free()
	help.XmlCleanUpParser()
	if help.XmlMemoryAllocation() != 0 {
		t.Errorf("Memeory leaks %d!!!", help.XmlMemoryAllocation())
		help.XmlMemoryLeakReport()
	}
}

func TestHtmlEmptyDoc(t *testing.T) {
	doc := libxml.HtmlParseString("")
	if !strings.Contains(doc.DumpHTML(), "<!DOCTYPE") {
		t.Error("Should have actually made a doc")
	}
	doc.Free()
	help.XmlCleanUpParser()
	if help.XmlMemoryAllocation() != 0 {
		t.Errorf("Memeory leaks %d!!!", help.XmlMemoryAllocation())
		help.XmlMemoryLeakReport()
	}
}
