package templ

import (
	"github.com/a-h/lexical/parse"
)

// Element.

// Element open tag.
type elementOpenTag struct {
	Name       string
	Attributes []Attribute
}

func newElementOpenTagParser() elementOpenTagParser {
	return elementOpenTagParser{}
}

type elementOpenTagParser struct {
}

func (p elementOpenTagParser) asElementOpenTag(parts []interface{}) (result interface{}, ok bool) {
	return elementOpenTag{
		Name:       parts[1].(string),
		Attributes: parts[2].([]Attribute),
	}, true
}

func (p elementOpenTagParser) Parse(pi parse.Input) parse.Result {
	return parse.All(p.asElementOpenTag,
		parse.Rune('<'),
		elementNameParser,
		newAttributesParser().Parse,
		parse.Optional(parse.WithStringConcatCombiner, whitespaceParser),
		parse.Rune('>'),
	)(pi)
}

// Element close tag.
type elementCloseTag struct {
	Name string
}

func asElementCloseTag(parts []interface{}) (result interface{}, ok bool) {
	return elementCloseTag{
		Name: parts[1].(string),
	}, true
}

var elementCloseTagParser = parse.All(asElementCloseTag,
	parse.String("</"),
	elementNameParser,
	parse.Rune('>'),
)

// Attribute name.
var attributeNameParser = parse.StringUntil(parse.Rune('='))

// Constant attribute.
var attributeConstantValueParser = parse.StringUntil(parse.Rune('"'))

func newConstantAttributeParser() constantAttributeParser {
	return constantAttributeParser{}
}

type constantAttributeParser struct {
}

func (p constantAttributeParser) asConstantAttribute(parts []interface{}) (result interface{}, ok bool) {
	return ConstantAttribute{
		Name:  parts[1].(string),
		Value: parts[4].(string),
	}, true
}

func (p constantAttributeParser) Parse(pi parse.Input) parse.Result {
	return parse.All(p.asConstantAttribute,
		whitespaceParser,
		attributeNameParser,
		parse.Rune('='),
		parse.Rune('"'),
		attributeConstantValueParser,
		parse.Rune('"'),
	)(pi)
}

// ExpressionAttribute.
func newExpressionAttributeParser() expressionAttributeParser {
	return expressionAttributeParser{}
}

type expressionAttributeParser struct {
}

func (p expressionAttributeParser) asExpressionAttribute(parts []interface{}) (result interface{}, ok bool) {
	return ExpressionAttribute{
		Name:  parts[1].(string),
		Value: parts[3].(StringExpression),
	}, true
}

func (p expressionAttributeParser) Parse(pi parse.Input) parse.Result {
	return parse.All(p.asExpressionAttribute,
		whitespaceParser,
		attributeNameParser,
		parse.Rune('='),
		newStringExpressionParser().Parse,
	)(pi)
}

// Attributes.
func newAttributesParser() attributesParser {
	return attributesParser{}
}

type attributesParser struct {
}

func (p attributesParser) asAttributeArray(parts []interface{}) (result interface{}, ok bool) {
	op := make([]Attribute, len(parts))
	for i := 0; i < len(parts); i++ {
		switch v := parts[i].(type) {
		case ConstantAttribute:
			op[i] = v
		case ExpressionAttribute:
			op[i] = v
		}
	}
	return op, true
}

func (p attributesParser) Parse(pi parse.Input) parse.Result {
	return parse.Many(p.asAttributeArray, 0, 255,
		parse.Or(
			newExpressionAttributeParser().Parse,
			newConstantAttributeParser().Parse,
		),
	)(pi)
}

// Element name.
var lowerAZ = "abcdefghijklmnopqrstuvwxyz"
var upperAZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers = "0123456789"
var elementNameParser = parse.Then(parse.WithStringConcatCombiner,
	parse.RuneIn(lowerAZ),
	parse.Many(parse.WithStringConcatCombiner, 0, 15, parse.RuneIn(lowerAZ+upperAZ+numbers)),
)

// Element.
func newElementOpenCloseParser() elementOpenCloseParser {
	return elementOpenCloseParser{}
}

type elementOpenCloseParser struct {
	SourceRangeToItemLookup SourceMap
}

func (p elementOpenCloseParser) asElement(parts []interface{}) (result interface{}, ok bool) {
	e := Element{
		Name:       parts[0].(elementOpenTag).Name,
		Attributes: parts[0].(elementOpenTag).Attributes,
	}
	if arr, isArray := parts[1].([]Node); isArray {
		e.Children = append(e.Children, arr...)
	}
	return e, true
}

func (p elementOpenCloseParser) asChildren(parts []interface{}) (result interface{}, ok bool) {
	return parts, true
}

func (p elementOpenCloseParser) Parse(pi parse.Input) parse.Result {
	//TODO: Don't use parse.All, check that the close tag name matches the start name!
	return parse.All(p.asElement,
		newElementOpenTagParser().Parse,
		newTemplateNodeParser().Parse,
		elementCloseTagParser,
	)(pi)
}

// Element self-closing tag.
func newElementSelfClosingParser() elementSelfClosingParser {
	return elementSelfClosingParser{}
}

type elementSelfClosingParser struct {
	SourceRangeToItemLookup SourceMap
}

func (p elementSelfClosingParser) asElement(parts []interface{}) (result interface{}, ok bool) {
	return Element{
		Name:       parts[1].(string),
		Attributes: parts[2].([]Attribute),
	}, true
}

func (p elementSelfClosingParser) Parse(pi parse.Input) parse.Result {
	return parse.All(p.asElement,
		parse.Rune('<'),
		elementNameParser,
		newAttributesParser().Parse,
		parse.Optional(parse.WithStringConcatCombiner, whitespaceParser),
		parse.String("/>"),
	)(pi)
}

// Element
func newElementParser() elementParser {
	return elementParser{}
}

type elementParser struct {
}

func (p elementParser) Parse(pi parse.Input) parse.Result {
	return parse.Or(
		newElementSelfClosingParser().Parse,
		newElementOpenCloseParser().Parse,
	)(pi)
}
