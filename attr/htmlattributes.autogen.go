//go:generate go run generate.go

// Package attr defines markup to create HTML attributes supported by Facebook React.
//
// Created from "HTML Attributes" as defined by Facebook in
// https://facebook.github.io/react/docs/tags-and-attributes.html

package attr

import "github.com/bep/gr"

// Accept creates an HTML attribute for 'accept'.
func Accept(v string) gr.Modifier {
	return gr.Prop("accept", v)
}

// AcceptCharset creates an HTML attribute for 'acceptCharset'.
func AcceptCharset(v string) gr.Modifier {
	return gr.Prop("acceptCharset", v)
}

// AccessKey creates an HTML attribute for 'accessKey'.
func AccessKey(v string) gr.Modifier {
	return gr.Prop("accessKey", v)
}

// Action creates an HTML attribute for 'action'.
func Action(v string) gr.Modifier {
	return gr.Prop("action", v)
}

// AllowFullScreen creates an HTML attribute for 'allowFullScreen'.
func AllowFullScreen(v string) gr.Modifier {
	return gr.Prop("allowFullScreen", v)
}

// AllowTransparency creates an HTML attribute for 'allowTransparency'.
func AllowTransparency(v string) gr.Modifier {
	return gr.Prop("allowTransparency", v)
}

// Alt creates an HTML attribute for 'alt'.
func Alt(v string) gr.Modifier {
	return gr.Prop("alt", v)
}

// Async creates an HTML attribute for 'async'.
func Async(v string) gr.Modifier {
	return gr.Prop("async", v)
}

// AutoComplete creates an HTML attribute for 'autoComplete'.
func AutoComplete(v string) gr.Modifier {
	return gr.Prop("autoComplete", v)
}

// AutoFocus creates an HTML attribute for 'autoFocus'.
func AutoFocus(v string) gr.Modifier {
	return gr.Prop("autoFocus", v)
}

// AutoPlay creates an HTML attribute for 'autoPlay'.
func AutoPlay(v string) gr.Modifier {
	return gr.Prop("autoPlay", v)
}

// Capture creates an HTML attribute for 'capture'.
func Capture(v string) gr.Modifier {
	return gr.Prop("capture", v)
}

// CellPadding creates an HTML attribute for 'cellPadding'.
func CellPadding(v string) gr.Modifier {
	return gr.Prop("cellPadding", v)
}

// CellSpacing creates an HTML attribute for 'cellSpacing'.
func CellSpacing(v string) gr.Modifier {
	return gr.Prop("cellSpacing", v)
}

// Challenge creates an HTML attribute for 'challenge'.
func Challenge(v string) gr.Modifier {
	return gr.Prop("challenge", v)
}

// CharSet creates an HTML attribute for 'charSet'.
func CharSet(v string) gr.Modifier {
	return gr.Prop("charSet", v)
}

// Checked creates an HTML attribute for 'checked'.
func Checked(v string) gr.Modifier {
	return gr.Prop("checked", v)
}

// Cite creates an HTML attribute for 'cite'.
func Cite(v string) gr.Modifier {
	return gr.Prop("cite", v)
}

// ClassID creates an HTML attribute for 'classID'.
func ClassID(v string) gr.Modifier {
	return gr.Prop("classID", v)
}

// ClassName creates an HTML attribute for 'className'.
func ClassName(v string) gr.Modifier {
	return gr.Prop("className", v)
}

// ColSpan creates an HTML attribute for 'colSpan'.
func ColSpan(v string) gr.Modifier {
	return gr.Prop("colSpan", v)
}

// Cols creates an HTML attribute for 'cols'.
func Cols(v string) gr.Modifier {
	return gr.Prop("cols", v)
}

// Content creates an HTML attribute for 'content'.
func Content(v string) gr.Modifier {
	return gr.Prop("content", v)
}

// ContentEditable creates an HTML attribute for 'contentEditable'.
func ContentEditable(v string) gr.Modifier {
	return gr.Prop("contentEditable", v)
}

// ContextMenu creates an HTML attribute for 'contextMenu'.
func ContextMenu(v string) gr.Modifier {
	return gr.Prop("contextMenu", v)
}

// Controls creates an HTML attribute for 'controls'.
func Controls(v string) gr.Modifier {
	return gr.Prop("controls", v)
}

// Coords creates an HTML attribute for 'coords'.
func Coords(v string) gr.Modifier {
	return gr.Prop("coords", v)
}

// CrossOrigin creates an HTML attribute for 'crossOrigin'.
func CrossOrigin(v string) gr.Modifier {
	return gr.Prop("crossOrigin", v)
}

// Data creates an HTML attribute for 'data'.
func Data(v string) gr.Modifier {
	return gr.Prop("data", v)
}

// DateTime creates an HTML attribute for 'dateTime'.
func DateTime(v string) gr.Modifier {
	return gr.Prop("dateTime", v)
}

// Default creates an HTML attribute for 'default'.
func Default(v string) gr.Modifier {
	return gr.Prop("default", v)
}

// Defer creates an HTML attribute for 'defer'.
func Defer(v string) gr.Modifier {
	return gr.Prop("defer", v)
}

// Dir creates an HTML attribute for 'dir'.
func Dir(v string) gr.Modifier {
	return gr.Prop("dir", v)
}

// Disabled creates an HTML attribute for 'disabled'.
func Disabled(v string) gr.Modifier {
	return gr.Prop("disabled", v)
}

// Download creates an HTML attribute for 'download'.
func Download(v string) gr.Modifier {
	return gr.Prop("download", v)
}

// Draggable creates an HTML attribute for 'draggable'.
func Draggable(v string) gr.Modifier {
	return gr.Prop("draggable", v)
}

// EncType creates an HTML attribute for 'encType'.
func EncType(v string) gr.Modifier {
	return gr.Prop("encType", v)
}

// Form creates an HTML attribute for 'form'.
func Form(v string) gr.Modifier {
	return gr.Prop("form", v)
}

// FormAction creates an HTML attribute for 'formAction'.
func FormAction(v string) gr.Modifier {
	return gr.Prop("formAction", v)
}

// FormEncType creates an HTML attribute for 'formEncType'.
func FormEncType(v string) gr.Modifier {
	return gr.Prop("formEncType", v)
}

// FormMethod creates an HTML attribute for 'formMethod'.
func FormMethod(v string) gr.Modifier {
	return gr.Prop("formMethod", v)
}

// FormNoValidate creates an HTML attribute for 'formNoValidate'.
func FormNoValidate(v string) gr.Modifier {
	return gr.Prop("formNoValidate", v)
}

// FormTarget creates an HTML attribute for 'formTarget'.
func FormTarget(v string) gr.Modifier {
	return gr.Prop("formTarget", v)
}

// FrameBorder creates an HTML attribute for 'frameBorder'.
func FrameBorder(v string) gr.Modifier {
	return gr.Prop("frameBorder", v)
}

// Headers creates an HTML attribute for 'headers'.
func Headers(v string) gr.Modifier {
	return gr.Prop("headers", v)
}

// Height creates an HTML attribute for 'height'.
func Height(v string) gr.Modifier {
	return gr.Prop("height", v)
}

// Hidden creates an HTML attribute for 'hidden'.
func Hidden(v string) gr.Modifier {
	return gr.Prop("hidden", v)
}

// High creates an HTML attribute for 'high'.
func High(v string) gr.Modifier {
	return gr.Prop("high", v)
}

// HRef creates an HTML attribute for 'href'.
func HRef(v string) gr.Modifier {
	return gr.Prop("href", v)
}

// HRefLang creates an HTML attribute for 'hrefLang'.
func HRefLang(v string) gr.Modifier {
	return gr.Prop("hrefLang", v)
}

// HTMLFor creates an HTML attribute for 'htmlFor'.
func HTMLFor(v string) gr.Modifier {
	return gr.Prop("htmlFor", v)
}

// HTTPEquiv creates an HTML attribute for 'httpEquiv'.
func HTTPEquiv(v string) gr.Modifier {
	return gr.Prop("httpEquiv", v)
}

// Icon creates an HTML attribute for 'icon'.
func Icon(v string) gr.Modifier {
	return gr.Prop("icon", v)
}

// ID creates an HTML attribute for 'id'.
func ID(v string) gr.Modifier {
	return gr.Prop("id", v)
}

// InputMode creates an HTML attribute for 'inputMode'.
func InputMode(v string) gr.Modifier {
	return gr.Prop("inputMode", v)
}

// Integrity creates an HTML attribute for 'integrity'.
func Integrity(v string) gr.Modifier {
	return gr.Prop("integrity", v)
}

// Is creates an HTML attribute for 'is'.
func Is(v string) gr.Modifier {
	return gr.Prop("is", v)
}

// KeyParams creates an HTML attribute for 'keyParams'.
func KeyParams(v string) gr.Modifier {
	return gr.Prop("keyParams", v)
}

// KeyType creates an HTML attribute for 'keyType'.
func KeyType(v string) gr.Modifier {
	return gr.Prop("keyType", v)
}

// Kind creates an HTML attribute for 'kind'.
func Kind(v string) gr.Modifier {
	return gr.Prop("kind", v)
}

// Label creates an HTML attribute for 'label'.
func Label(v string) gr.Modifier {
	return gr.Prop("label", v)
}

// Lang creates an HTML attribute for 'lang'.
func Lang(v string) gr.Modifier {
	return gr.Prop("lang", v)
}

// List creates an HTML attribute for 'list'.
func List(v string) gr.Modifier {
	return gr.Prop("list", v)
}

// Loop creates an HTML attribute for 'loop'.
func Loop(v string) gr.Modifier {
	return gr.Prop("loop", v)
}

// Low creates an HTML attribute for 'low'.
func Low(v string) gr.Modifier {
	return gr.Prop("low", v)
}

// Manifest creates an HTML attribute for 'manifest'.
func Manifest(v string) gr.Modifier {
	return gr.Prop("manifest", v)
}

// MarginHeight creates an HTML attribute for 'marginHeight'.
func MarginHeight(v string) gr.Modifier {
	return gr.Prop("marginHeight", v)
}

// MarginWidth creates an HTML attribute for 'marginWidth'.
func MarginWidth(v string) gr.Modifier {
	return gr.Prop("marginWidth", v)
}

// Max creates an HTML attribute for 'max'.
func Max(v string) gr.Modifier {
	return gr.Prop("max", v)
}

// MaxLength creates an HTML attribute for 'maxLength'.
func MaxLength(v string) gr.Modifier {
	return gr.Prop("maxLength", v)
}

// Media creates an HTML attribute for 'media'.
func Media(v string) gr.Modifier {
	return gr.Prop("media", v)
}

// MediaGroup creates an HTML attribute for 'mediaGroup'.
func MediaGroup(v string) gr.Modifier {
	return gr.Prop("mediaGroup", v)
}

// Method creates an HTML attribute for 'method'.
func Method(v string) gr.Modifier {
	return gr.Prop("method", v)
}

// Min creates an HTML attribute for 'min'.
func Min(v string) gr.Modifier {
	return gr.Prop("min", v)
}

// MinLength creates an HTML attribute for 'minLength'.
func MinLength(v string) gr.Modifier {
	return gr.Prop("minLength", v)
}

// Multiple creates an HTML attribute for 'multiple'.
func Multiple(v string) gr.Modifier {
	return gr.Prop("multiple", v)
}

// Muted creates an HTML attribute for 'muted'.
func Muted(v string) gr.Modifier {
	return gr.Prop("muted", v)
}

// Name creates an HTML attribute for 'name'.
func Name(v string) gr.Modifier {
	return gr.Prop("name", v)
}

// NoValidate creates an HTML attribute for 'noValidate'.
func NoValidate(v string) gr.Modifier {
	return gr.Prop("noValidate", v)
}

// Nonce creates an HTML attribute for 'nonce'.
func Nonce(v string) gr.Modifier {
	return gr.Prop("nonce", v)
}

// Open creates an HTML attribute for 'open'.
func Open(v string) gr.Modifier {
	return gr.Prop("open", v)
}

// Optimum creates an HTML attribute for 'optimum'.
func Optimum(v string) gr.Modifier {
	return gr.Prop("optimum", v)
}

// Pattern creates an HTML attribute for 'pattern'.
func Pattern(v string) gr.Modifier {
	return gr.Prop("pattern", v)
}

// Placeholder creates an HTML attribute for 'placeholder'.
func Placeholder(v string) gr.Modifier {
	return gr.Prop("placeholder", v)
}

// Poster creates an HTML attribute for 'poster'.
func Poster(v string) gr.Modifier {
	return gr.Prop("poster", v)
}

// Preload creates an HTML attribute for 'preload'.
func Preload(v string) gr.Modifier {
	return gr.Prop("preload", v)
}

// Profile creates an HTML attribute for 'profile'.
func Profile(v string) gr.Modifier {
	return gr.Prop("profile", v)
}

// RadioGroup creates an HTML attribute for 'radioGroup'.
func RadioGroup(v string) gr.Modifier {
	return gr.Prop("radioGroup", v)
}

// ReadOnly creates an HTML attribute for 'readOnly'.
func ReadOnly(v string) gr.Modifier {
	return gr.Prop("readOnly", v)
}

// Rel creates an HTML attribute for 'rel'.
func Rel(v string) gr.Modifier {
	return gr.Prop("rel", v)
}

// Required creates an HTML attribute for 'required'.
func Required(v string) gr.Modifier {
	return gr.Prop("required", v)
}

// Reversed creates an HTML attribute for 'reversed'.
func Reversed(v string) gr.Modifier {
	return gr.Prop("reversed", v)
}

// Role creates an HTML attribute for 'role'.
func Role(v string) gr.Modifier {
	return gr.Prop("role", v)
}

// RowSpan creates an HTML attribute for 'rowSpan'.
func RowSpan(v string) gr.Modifier {
	return gr.Prop("rowSpan", v)
}

// Rows creates an HTML attribute for 'rows'.
func Rows(v string) gr.Modifier {
	return gr.Prop("rows", v)
}

// Sandbox creates an HTML attribute for 'sandbox'.
func Sandbox(v string) gr.Modifier {
	return gr.Prop("sandbox", v)
}

// Scope creates an HTML attribute for 'scope'.
func Scope(v string) gr.Modifier {
	return gr.Prop("scope", v)
}

// Scoped creates an HTML attribute for 'scoped'.
func Scoped(v string) gr.Modifier {
	return gr.Prop("scoped", v)
}

// Scrolling creates an HTML attribute for 'scrolling'.
func Scrolling(v string) gr.Modifier {
	return gr.Prop("scrolling", v)
}

// Seamless creates an HTML attribute for 'seamless'.
func Seamless(v string) gr.Modifier {
	return gr.Prop("seamless", v)
}

// Selected creates an HTML attribute for 'selected'.
func Selected(v string) gr.Modifier {
	return gr.Prop("selected", v)
}

// Shape creates an HTML attribute for 'shape'.
func Shape(v string) gr.Modifier {
	return gr.Prop("shape", v)
}

// Size creates an HTML attribute for 'size'.
func Size(v string) gr.Modifier {
	return gr.Prop("size", v)
}

// Sizes creates an HTML attribute for 'sizes'.
func Sizes(v string) gr.Modifier {
	return gr.Prop("sizes", v)
}

// Span creates an HTML attribute for 'span'.
func Span(v string) gr.Modifier {
	return gr.Prop("span", v)
}

// SpellCheck creates an HTML attribute for 'spellCheck'.
func SpellCheck(v string) gr.Modifier {
	return gr.Prop("spellCheck", v)
}

// Src creates an HTML attribute for 'src'.
func Src(v string) gr.Modifier {
	return gr.Prop("src", v)
}

// SrcDoc creates an HTML attribute for 'srcDoc'.
func SrcDoc(v string) gr.Modifier {
	return gr.Prop("srcDoc", v)
}

// SrcLang creates an HTML attribute for 'srcLang'.
func SrcLang(v string) gr.Modifier {
	return gr.Prop("srcLang", v)
}

// SrcSet creates an HTML attribute for 'srcSet'.
func SrcSet(v string) gr.Modifier {
	return gr.Prop("srcSet", v)
}

// Start creates an HTML attribute for 'start'.
func Start(v string) gr.Modifier {
	return gr.Prop("start", v)
}

// Step creates an HTML attribute for 'step'.
func Step(v string) gr.Modifier {
	return gr.Prop("step", v)
}

// Style creates an HTML attribute for 'style'.
func Style(v string) gr.Modifier {
	return gr.Prop("style", v)
}

// Summary creates an HTML attribute for 'summary'.
func Summary(v string) gr.Modifier {
	return gr.Prop("summary", v)
}

// TabIndex creates an HTML attribute for 'tabIndex'.
func TabIndex(v string) gr.Modifier {
	return gr.Prop("tabIndex", v)
}

// Target creates an HTML attribute for 'target'.
func Target(v string) gr.Modifier {
	return gr.Prop("target", v)
}

// Title creates an HTML attribute for 'title'.
func Title(v string) gr.Modifier {
	return gr.Prop("title", v)
}

// Type creates an HTML attribute for 'type'.
func Type(v string) gr.Modifier {
	return gr.Prop("type", v)
}

// UseMap creates an HTML attribute for 'useMap'.
func UseMap(v string) gr.Modifier {
	return gr.Prop("useMap", v)
}

// Value creates an HTML attribute for 'value'.
func Value(v string) gr.Modifier {
	return gr.Prop("value", v)
}

// Width creates an HTML attribute for 'width'.
func Width(v string) gr.Modifier {
	return gr.Prop("width", v)
}

// Wmode creates an HTML attribute for 'wmode'.
func Wmode(v string) gr.Modifier {
	return gr.Prop("wmode", v)
}

// Wrap creates an HTML attribute for 'wrap'.
func Wrap(v string) gr.Modifier {
	return gr.Prop("wrap", v)
}

// About creates an HTML attribute for 'about'.
func About(v string) gr.Modifier {
	return gr.Prop("about", v)
}

// Datatype creates an HTML attribute for 'datatype'.
func Datatype(v string) gr.Modifier {
	return gr.Prop("datatype", v)
}

// Inlist creates an HTML attribute for 'inlist'.
func Inlist(v string) gr.Modifier {
	return gr.Prop("inlist", v)
}

// Prefix creates an HTML attribute for 'prefix'.
func Prefix(v string) gr.Modifier {
	return gr.Prop("prefix", v)
}

// Property creates an HTML attribute for 'property'.
func Property(v string) gr.Modifier {
	return gr.Prop("property", v)
}

// Resource creates an HTML attribute for 'resource'.
func Resource(v string) gr.Modifier {
	return gr.Prop("resource", v)
}

// Typeof creates an HTML attribute for 'typeof'.
func Typeof(v string) gr.Modifier {
	return gr.Prop("typeof", v)
}

// Vocab creates an HTML attribute for 'vocab'.
func Vocab(v string) gr.Modifier {
	return gr.Prop("vocab", v)
}

// AutoCapitalize creates an HTML attribute for 'autoCapitalize'.
func AutoCapitalize(v string) gr.Modifier {
	return gr.Prop("autoCapitalize", v)
}

// AutoCorrect creates an HTML attribute for 'autoCorrect'.
func AutoCorrect(v string) gr.Modifier {
	return gr.Prop("autoCorrect", v)
}

// Color creates an HTML attribute for 'color'.
func Color(v string) gr.Modifier {
	return gr.Prop("color", v)
}

// ItemProp creates an HTML attribute for 'itemProp'.
func ItemProp(v string) gr.Modifier {
	return gr.Prop("itemProp", v)
}

// Security creates an HTML attribute for 'security'.
func Security(v string) gr.Modifier {
	return gr.Prop("security", v)
}

// Unselectable creates an HTML attribute for 'unselectable'.
func Unselectable(v string) gr.Modifier {
	return gr.Prop("unselectable", v)
}

// Results creates an HTML attribute for 'results'.
func Results(v string) gr.Modifier {
	return gr.Prop("results", v)
}

// AutoSave creates an HTML attribute for 'autoSave'.
func AutoSave(v string) gr.Modifier {
	return gr.Prop("autoSave", v)
}