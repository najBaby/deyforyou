package templates

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/facebook/ent/schema/field"

	"github.com/go-openapi/inflect"
)

var (
	// Funcs are the predefined template
	// functions used by the codegen.
	Funcs = template.FuncMap{
		"add":         add,
		"append":      reflect.Append,
		"appends":     reflect.AppendSlice,
		"order":       order,
		"camel":       camel,
		"snake":       snake,
		"pascal":      pascal,
		"xrange":      xrange,
		"plural":      plural,
		"aggregate":   aggregate,
		"primitives":  primitives,
		"singular":    rules.Singularize,
		"quote":       strconv.Quote,
		"base":        filepath.Base,
		"keys":        keys,
		"join":        join,
		"lower":       strings.ToLower,
		"upper":       strings.ToUpper,
		"hasField":    hasField,
		"indirect":    indirect,
		"hasPrefix":   strings.HasPrefix,
		"hasSuffix":   strings.HasSuffix,
		"trimPackage": trimPackage,
		"split":       strings.Split,
		"tagLookup":   tagLookup,
		"toString":    toString,
		"dict":        dict,
		"get":         get,
		"set":         set,
		"unset":       unset,
		"hasKey":      hasKey,
		"list":        list,
		"fail":        fail,
		"replace":     strings.ReplaceAll,
	}
	rules    = ruleset()
	acronyms = make(map[string]struct{})
)

// xrange generates a slice of len n.
func xrange(n int) (a []int) {
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return
}

// plural a name.
func plural(name string) string {
	p := rules.Pluralize(name)
	if p == name {
		p += "Slice"
	}
	return p
}

func isSeparator(r rune) bool {
	return r == '_' || r == '-'
}

func pascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

// pascal converts the given name into a PascalCase.
//
//	user_info 	=> UserInfo
//	full_name 	=> FullName
//	user_id   	=> UserID
//	full-admin	=> FullAdmin
//
func pascal(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	return pascalWords(words)
}

// camel converts the given name into a camelCase.
//
//	user_info  => userInfo
//	full_name  => fullName
//	user_id    => userID
//	full-admin => fullAdmin
//
func camel(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	if len(words) == 1 {
		return strings.ToLower(words[0])
	}
	return strings.ToLower(words[0]) + pascalWords(words[1:])
}

// snake converts the given struct or field name into a snake_case.
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
//
func snake(s string) string {
	var (
		j int
		b strings.Builder
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// Put '_' if it is not a start or end of a word, current letter is uppercase,
		// and previous is lowercase (cases like: "UserInfo"), or next letter is also
		// a lowercase and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) {
			if unicode.IsLower(rune(s[i-1])) ||
				j != i-1 && unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1])) {
				j = i
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// add calculates summarize list of variables.
func add(xs ...int) (n int) {
	for _, x := range xs {
		n += x
	}
	return
}

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// Add common initialisms from golint and more.
	for _, w := range []string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB",
		"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TCP",
		"TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM",
		"XML", "XMPP", "XSRF", "XSS",
	} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}

// order returns a map of sort orders.
// The key is the function name, and the value its database keyword.
func order() map[string]string {
	return map[string]string{
		"asc":  "incr",
		"desc": "decr",
	}
}

// aggregate returns a map between all agg-functions and if they accept a field name as a parameter or not.
func aggregate() map[string]bool {
	return map[string]bool{
		"min":   true,
		"max":   true,
		"sum":   true,
		"mean":  true,
		"count": false,
	}
}

// keys returns the given map keys.
func keys(v reflect.Value) ([]string, error) {
	if k := v.Type().Kind(); k != reflect.Map {
		return nil, fmt.Errorf("expect map for keys, got: %s", k)
	}
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		keys[i] = v.String()
	}
	sort.Strings(keys)
	return keys, nil
}

// primitives returns all primitives types.
func primitives() []string {
	return []string{field.TypeString.String(), field.TypeInt.String(), field.TypeFloat64.String(), field.TypeBool.String()}
}

// join is a wrapper around strings.Join to provide consistent output.
func join(a []string, sep string) string {
	sort.Strings(a)
	return strings.Join(a, sep)
}

// hasField determines if a struct has a field with the given name.
func hasField(v interface{}, name string) bool {
	vr := reflect.Indirect(reflect.ValueOf(v))
	return vr.FieldByName(name).IsValid()
}

// trimPackage trims the package name from the given identifier.
func trimPackage(ident, pkg string) string {
	return strings.TrimPrefix(ident, pkg+".")
}

// indirect returns the item at the end of indirection.
func indirect(v reflect.Value) reflect.Value {
	for ; v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface; v = v.Elem() {
	}
	return v
}

// tagLookup returns the value associated with key in the tag string.
func tagLookup(tag, key string) string {
	v, _ := reflect.StructTag(tag).Lookup(key)
	return v
}

// toString converts `v` to a string.
func toString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprint(v)
	}
}

// dict creates a dictionary from a list of pairs.
func dict(v ...interface{}) map[string]interface{} {
	lenv := len(v)
	dict := make(map[string]interface{}, lenv/2)
	for i := 0; i < lenv; i += 2 {
		key := toString(v[i])
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}

// get the value from the dict for key.
func get(d map[string]interface{}, key string) interface{} {
	if val, ok := d[key]; ok {
		return val
	}
	return ""
}

// set adds a value to the dict for key.
func set(d map[string]interface{}, key string, value interface{}) map[string]interface{} {
	d[key] = value
	return d
}

// unset removes a key from the dict.
func unset(d map[string]interface{}, key string) map[string]interface{} {
	delete(d, key)
	return d
}

// hasKey tests whether a key is found in dict.
func hasKey(d map[string]interface{}, key string) bool {
	_, ok := d[key]
	return ok
}

// list creates a list from values.
func list(v ...interface{}) []interface{} {
	return v
}

// fail unconditionally returns an empty string and an error with the specified text.
func fail(msg string) (string, error) {
	return "", errors.New(msg)
}
