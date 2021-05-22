package kama

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/sanity-io/litter"
)

func dump(val reflect.Value, compact bool, hideZeroValues bool, indentLevel int) string {
	/*
		`compact` indicates whether the struct should be laid in one line or not
		`hideZeroValues` indicates whether to show zeroValued vars
		`indentLevel` is the number of spaces from the left-most side of the termninal for struct names
	*/
	iType := val.Type()
	maxL := 720

	if iType == nil {
		// TODO: handle this better
		return "Nil NotImplemented"
	}
	indentLevel = indentLevel + 1

	switch iType.Kind() {
	case reflect.String:
		return dumpString(val, compact, hideZeroValues)
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64:
		return fmt.Sprint(val)
	case reflect.Struct:
		// the reason we are doing this is because sanity-io/litter has no way to compact
		// arrays/slices/maps that are inside structs.
		// This logic can be discarded if sanity-io/litter implements similar.
		// see: https://github.com/sanity-io/litter/pull/43
		fromPtr := false
		return dumpStruct(val, fromPtr, compact, hideZeroValues, indentLevel)
	case reflect.Ptr:
		v := val.Elem()
		if v.IsValid() {
			if v.Type().Kind() == reflect.Struct {
				// the reason we are doing this is because sanity-io/litter has no way to compact
				// arrays/slices/maps that are inside structs.
				// This logic can be discarded if sanity-io/litter implements similar.
				// see: https://github.com/sanity-io/litter/pull/43
				fromPtr := true
				return dumpStruct(v, fromPtr, compact, hideZeroValues, indentLevel)
			}
		}
	case reflect.Array,
		reflect.Slice:
		// In future we could restrict compaction only to arrays/slices/maps that are of primitive(basic) types
		// see: https://github.com/sanity-io/litter/pull/43
		cpt := true
		hideZeroValues := true
		return dumpSlice(val, cpt, hideZeroValues, indentLevel)
	case reflect.Map:
		// In future we could restrict compaction only to arrays/slices/maps that are of primitive(basic) types
		// see: https://github.com/sanity-io/litter/pull/43
		maxL = 50
	default:
		return fmt.Sprintf("%v NotImplemented", iType.Kind())
	}

	x := 9
	if x < 5 {
		sq := litter.Options{
			StripPackageNames: true,
			HidePrivateFields: true,
			HideZeroValues:    false,
			FieldExclusions:   regexp.MustCompile(`^(XXX_.*)$`), // XXX_ is a prefix of fields generated by protoc-gen-go
			Separator:         " "}

		s := sq.Sdump(val)
		if len(s) <= maxL {
			maxL = len(s)
			return s[:maxL]
		}
	}

	_ = maxL
	return "NotImplemented (note:Went outside swith.)"
}

func dumpString(v reflect.Value, compact bool, hideZeroValues bool) string {
	//dumps strings
	maxL := 50

	numEntries := v.Len()
	constraint := int(math.Min(float64(numEntries), float64(maxL))) + 2 // the `+2` is important so that the final quote `"` at end of string is not cut off
	s := fmt.Sprintf("%#v", v)[:constraint]

	if numEntries > constraint {
		remainder := numEntries - constraint
		s = s + fmt.Sprintf(" ...<%d more redacted>..", remainder)
	}
	if s == "" {
		s = `""`
	}

	return s
}

func dumpStruct(v reflect.Value, fromPtr bool, compact bool, hideZeroValues bool, indentLevel int) string {
	/*
		`fromPtr` indicates whether the struct is a value or a pointer; `T{}` vs `&T{}`
		`compact` indicates whether the struct should be laid in one line or not
		`hideZeroValues` indicates whether to show zeroValued fields
		`indentLevel` is the number of spaces from the left-most side of the termninal for struct names
	*/
	// This logic is only required until similar logic is implemented in sanity-io/litter
	// see:
	// - https://github.com/sanity-io/litter/issues/34
	// - https://github.com/sanity-io/litter/pull/43

	typeName := v.Type().Name()
	if fromPtr {
		typeName = "&" + typeName
	}

	sep := "\n"
	if compact {
		sep = ""
	}
	fieldNameSep := strings.Repeat("  ", indentLevel)
	if compact {
		fieldNameSep = ""
	}

	vt := v.Type()
	s := fmt.Sprintf("%s{%s", typeName, sep)

	numFields := v.NumField()
	for i := 0; i < numFields; i++ {
		vtf := vt.Field(i)
		fieldd := v.Field(i)
		if unicode.IsUpper(rune(vtf.Name[0])) {
			if hideZeroValues && isZeroValue(fieldd) {
				continue
			} else {
				val := dump(fieldd, compact, hideZeroValues, indentLevel)
				s = s + fieldNameSep + vtf.Name + ": " + val + fmt.Sprintf(",%s", sep)
			}
		}
	}
	s = s + "}"
	return s
}

func dumpSlice(v reflect.Value, compact bool, hideZeroValues bool, indentLevel int) string {
	//dumps slices & arrays
	maxL := 10
	numEntries := v.Len()
	constraint := int(math.Min(float64(numEntries), float64(maxL)))
	typeName := v.Type().String()

	s := typeName + "{"
	for i := 0; i < constraint; i++ {
		elm := v.Index(i) // todo: call dump on this
		s = s + dump(elm, compact, hideZeroValues, indentLevel) + ","
	}
	if numEntries > constraint {
		remainder := numEntries - constraint
		s = s + fmt.Sprintf(" ...<%d more redacted>..", remainder)
	}
	s = s + "}"
	return s
}

func isPointerValue(v reflect.Value) bool {
	// Taken from https://github.com/sanity-io/litter/blob/v1.5.1/util.go
	// under the MIT license;
	// https://github.com/sanity-io/litter/blob/v1.5.1/LICENSE
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return true
	}
	return false
}

func isZeroValue(v reflect.Value) bool {
	// Taken from https://github.com/sanity-io/litter/blob/v1.5.1/util.go
	// under the MIT license;
	// https://github.com/sanity-io/litter/blob/v1.5.1/LICENSE
	return (isPointerValue(v) && v.IsNil()) ||
		(v.IsValid() && v.CanInterface() && reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()))
}
