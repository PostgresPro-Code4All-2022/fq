// Code below generated from scalar_gen.go.tmpl
package decode

import (
	"errors"
	"fmt"

	"github.com/wader/fq/pkg/bitio"
)

// Type BitBuf

// ActualBitBuf asserts actual value is a BitBuf and returns it
func (s Scalar) ActualBitBuf() *bitio.Buffer {
	v, ok := s.Actual.(*bitio.Buffer)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Actual %v as *bitio.Buffer", s.Actual))
	}
	return v
}

// SymBitBuf asserts symbolic value is a BitBuf and returns it
func (s Scalar) SymBitBuf() *bitio.Buffer {
	v, ok := s.Sym.(*bitio.Buffer)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Sym %v as *bitio.Buffer", s.Sym))
	}
	return v
}

// FieldBitBufScalarFn adds a field, calls scalar functions and returns actual value as a BitBuf
func (d *D) FieldBitBufScalarFn(name string, fn func(d *D) Scalar, sfns ...ScalarFn) *bitio.Buffer {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d), nil }, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "BitBuf", Pos: d.Pos()})
	}
	return v.ActualBitBuf()
}

// FieldBitBufFn adds a field, calls *bitio.Buffer decode function and returns actual value as a BitBuf
func (d *D) FieldBitBufFn(name string, fn func(d *D) *bitio.Buffer, sfns ...ScalarFn) *bitio.Buffer {
	return d.FieldBitBufScalarFn(name, func(d *D) Scalar { return Scalar{Actual: fn(d)} }, sfns...)
}

// TryFieldBitBufScalarFn tries to add a field, calls scalar functions and returns actual value as a BitBuf
func (d *D) TryFieldBitBufScalarFn(name string, fn func(d *D) (Scalar, error), sfns ...ScalarFn) (*bitio.Buffer, error) {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d) }, sfns...)
	if err != nil {
		return nil, err
	}
	return v.ActualBitBuf(), err
}

// TryFieldBitBufFn tries to add a field, calls *bitio.Buffer decode function and returns actual value as a BitBuf
func (d *D) TryFieldBitBufFn(name string, fn func(d *D) (*bitio.Buffer, error), sfns ...ScalarFn) (*bitio.Buffer, error) {
	return d.TryFieldBitBufScalarFn(name, func(d *D) (Scalar, error) {
		v, err := fn(d)
		return Scalar{Actual: v}, err
	}, sfns...)
}

// Type Bool

// ActualBool asserts actual value is a Bool and returns it
func (s Scalar) ActualBool() bool {
	v, ok := s.Actual.(bool)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Actual %v as bool", s.Actual))
	}
	return v
}

// SymBool asserts symbolic value is a Bool and returns it
func (s Scalar) SymBool() bool {
	v, ok := s.Sym.(bool)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Sym %v as bool", s.Sym))
	}
	return v
}

// FieldBoolScalarFn adds a field, calls scalar functions and returns actual value as a Bool
func (d *D) FieldBoolScalarFn(name string, fn func(d *D) Scalar, sfns ...ScalarFn) bool {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d), nil }, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "Bool", Pos: d.Pos()})
	}
	return v.ActualBool()
}

// FieldBoolFn adds a field, calls bool decode function and returns actual value as a Bool
func (d *D) FieldBoolFn(name string, fn func(d *D) bool, sfns ...ScalarFn) bool {
	return d.FieldBoolScalarFn(name, func(d *D) Scalar { return Scalar{Actual: fn(d)} }, sfns...)
}

// TryFieldBoolScalarFn tries to add a field, calls scalar functions and returns actual value as a Bool
func (d *D) TryFieldBoolScalarFn(name string, fn func(d *D) (Scalar, error), sfns ...ScalarFn) (bool, error) {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d) }, sfns...)
	if err != nil {
		return false, err
	}
	return v.ActualBool(), err
}

// TryFieldBoolFn tries to add a field, calls bool decode function and returns actual value as a Bool
func (d *D) TryFieldBoolFn(name string, fn func(d *D) (bool, error), sfns ...ScalarFn) (bool, error) {
	return d.TryFieldBoolScalarFn(name, func(d *D) (Scalar, error) {
		v, err := fn(d)
		return Scalar{Actual: v}, err
	}, sfns...)
}

// Type F

// ActualF asserts actual value is a F and returns it
func (s Scalar) ActualF() float64 {
	v, ok := s.Actual.(float64)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Actual %v as float64", s.Actual))
	}
	return v
}

// SymF asserts symbolic value is a F and returns it
func (s Scalar) SymF() float64 {
	v, ok := s.Sym.(float64)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Sym %v as float64", s.Sym))
	}
	return v
}

// FieldFScalarFn adds a field, calls scalar functions and returns actual value as a F
func (d *D) FieldFScalarFn(name string, fn func(d *D) Scalar, sfns ...ScalarFn) float64 {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d), nil }, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "F", Pos: d.Pos()})
	}
	return v.ActualF()
}

// FieldFFn adds a field, calls float64 decode function and returns actual value as a F
func (d *D) FieldFFn(name string, fn func(d *D) float64, sfns ...ScalarFn) float64 {
	return d.FieldFScalarFn(name, func(d *D) Scalar { return Scalar{Actual: fn(d)} }, sfns...)
}

// TryFieldFScalarFn tries to add a field, calls scalar functions and returns actual value as a F
func (d *D) TryFieldFScalarFn(name string, fn func(d *D) (Scalar, error), sfns ...ScalarFn) (float64, error) {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d) }, sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualF(), err
}

// TryFieldFFn tries to add a field, calls float64 decode function and returns actual value as a F
func (d *D) TryFieldFFn(name string, fn func(d *D) (float64, error), sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFScalarFn(name, func(d *D) (Scalar, error) {
		v, err := fn(d)
		return Scalar{Actual: v}, err
	}, sfns...)
}

// Type S

// ActualS asserts actual value is a S and returns it
func (s Scalar) ActualS() int64 {
	v, ok := s.Actual.(int64)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Actual %v as int64", s.Actual))
	}
	return v
}

// SymS asserts symbolic value is a S and returns it
func (s Scalar) SymS() int64 {
	v, ok := s.Sym.(int64)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Sym %v as int64", s.Sym))
	}
	return v
}

// FieldSScalarFn adds a field, calls scalar functions and returns actual value as a S
func (d *D) FieldSScalarFn(name string, fn func(d *D) Scalar, sfns ...ScalarFn) int64 {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d), nil }, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "S", Pos: d.Pos()})
	}
	return v.ActualS()
}

// FieldSFn adds a field, calls int64 decode function and returns actual value as a S
func (d *D) FieldSFn(name string, fn func(d *D) int64, sfns ...ScalarFn) int64 {
	return d.FieldSScalarFn(name, func(d *D) Scalar { return Scalar{Actual: fn(d)} }, sfns...)
}

// TryFieldSScalarFn tries to add a field, calls scalar functions and returns actual value as a S
func (d *D) TryFieldSScalarFn(name string, fn func(d *D) (Scalar, error), sfns ...ScalarFn) (int64, error) {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d) }, sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualS(), err
}

// TryFieldSFn tries to add a field, calls int64 decode function and returns actual value as a S
func (d *D) TryFieldSFn(name string, fn func(d *D) (int64, error), sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSScalarFn(name, func(d *D) (Scalar, error) {
		v, err := fn(d)
		return Scalar{Actual: v}, err
	}, sfns...)
}

// Type Str

// ActualStr asserts actual value is a Str and returns it
func (s Scalar) ActualStr() string {
	v, ok := s.Actual.(string)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Actual %v as string", s.Actual))
	}
	return v
}

// SymStr asserts symbolic value is a Str and returns it
func (s Scalar) SymStr() string {
	v, ok := s.Sym.(string)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Sym %v as string", s.Sym))
	}
	return v
}

// FieldStrScalarFn adds a field, calls scalar functions and returns actual value as a Str
func (d *D) FieldStrScalarFn(name string, fn func(d *D) Scalar, sfns ...ScalarFn) string {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d), nil }, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "Str", Pos: d.Pos()})
	}
	return v.ActualStr()
}

// FieldStrFn adds a field, calls string decode function and returns actual value as a Str
func (d *D) FieldStrFn(name string, fn func(d *D) string, sfns ...ScalarFn) string {
	return d.FieldStrScalarFn(name, func(d *D) Scalar { return Scalar{Actual: fn(d)} }, sfns...)
}

// TryFieldStrScalarFn tries to add a field, calls scalar functions and returns actual value as a Str
func (d *D) TryFieldStrScalarFn(name string, fn func(d *D) (Scalar, error), sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d) }, sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// TryFieldStrFn tries to add a field, calls string decode function and returns actual value as a Str
func (d *D) TryFieldStrFn(name string, fn func(d *D) (string, error), sfns ...ScalarFn) (string, error) {
	return d.TryFieldStrScalarFn(name, func(d *D) (Scalar, error) {
		v, err := fn(d)
		return Scalar{Actual: v}, err
	}, sfns...)
}

// Type U

// ActualU asserts actual value is a U and returns it
func (s Scalar) ActualU() uint64 {
	v, ok := s.Actual.(uint64)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Actual %v as uint64", s.Actual))
	}
	return v
}

// SymU asserts symbolic value is a U and returns it
func (s Scalar) SymU() uint64 {
	v, ok := s.Sym.(uint64)
	if !ok {
		panic(fmt.Sprintf("failed to type assert s.Sym %v as uint64", s.Sym))
	}
	return v
}

// FieldUScalarFn adds a field, calls scalar functions and returns actual value as a U
func (d *D) FieldUScalarFn(name string, fn func(d *D) Scalar, sfns ...ScalarFn) uint64 {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d), nil }, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "U", Pos: d.Pos()})
	}
	return v.ActualU()
}

// FieldUFn adds a field, calls uint64 decode function and returns actual value as a U
func (d *D) FieldUFn(name string, fn func(d *D) uint64, sfns ...ScalarFn) uint64 {
	return d.FieldUScalarFn(name, func(d *D) Scalar { return Scalar{Actual: fn(d)} }, sfns...)
}

// TryFieldUScalarFn tries to add a field, calls scalar functions and returns actual value as a U
func (d *D) TryFieldUScalarFn(name string, fn func(d *D) (Scalar, error), sfns ...ScalarFn) (uint64, error) {
	v, err := d.TryFieldScalar(name, func(_ Scalar) (Scalar, error) { return fn(d) }, sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualU(), err
}

// TryFieldUFn tries to add a field, calls uint64 decode function and returns actual value as a U
func (d *D) TryFieldUFn(name string, fn func(d *D) (uint64, error), sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUScalarFn(name, func(d *D) (Scalar, error) {
		v, err := fn(d)
		return Scalar{Actual: v}, err
	}, sfns...)
}

// Validate/Assert Bool

func (d *D) assertBool(assert bool, vs ...bool) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualBool()
		for _, b := range vs {
			if a == b {
				s.Description = "valid"
				return s, nil
			}
		}
		s.Description = "invalid"
		if assert && !d.Options.Force {
			return s, errors.New("failed to assert Bool")
		}
		return s, nil
	}
}

// AssertBool asserts that actual value is one of given bool values
func (d *D) AssertBool(vs ...bool) func(s Scalar) (Scalar, error) {
	return d.assertBool(true, vs...)
}

// ValidateBool validates that actual value is one of given bool values
func (d *D) ValidateBool(vs ...bool) func(s Scalar) (Scalar, error) {
	return d.assertBool(false, vs...)
}

// Validate/Assert F

func (d *D) assertF(assert bool, vs ...float64) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualF()
		for _, b := range vs {
			if a == b {
				s.Description = "valid"
				return s, nil
			}
		}
		s.Description = "invalid"
		if assert && !d.Options.Force {
			return s, errors.New("failed to assert F")
		}
		return s, nil
	}
}

// AssertF asserts that actual value is one of given float64 values
func (d *D) AssertF(vs ...float64) func(s Scalar) (Scalar, error) {
	return d.assertF(true, vs...)
}

// ValidateF validates that actual value is one of given float64 values
func (d *D) ValidateF(vs ...float64) func(s Scalar) (Scalar, error) {
	return d.assertF(false, vs...)
}

// Validate/Assert S

func (d *D) assertS(assert bool, vs ...int64) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualS()
		for _, b := range vs {
			if a == b {
				s.Description = "valid"
				return s, nil
			}
		}
		s.Description = "invalid"
		if assert && !d.Options.Force {
			return s, errors.New("failed to assert S")
		}
		return s, nil
	}
}

// AssertS asserts that actual value is one of given int64 values
func (d *D) AssertS(vs ...int64) func(s Scalar) (Scalar, error) {
	return d.assertS(true, vs...)
}

// ValidateS validates that actual value is one of given int64 values
func (d *D) ValidateS(vs ...int64) func(s Scalar) (Scalar, error) {
	return d.assertS(false, vs...)
}

// Validate/Assert Str

func (d *D) assertStr(assert bool, vs ...string) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualStr()
		for _, b := range vs {
			if a == b {
				s.Description = "valid"
				return s, nil
			}
		}
		s.Description = "invalid"
		if assert && !d.Options.Force {
			return s, errors.New("failed to assert Str")
		}
		return s, nil
	}
}

// AssertStr asserts that actual value is one of given string values
func (d *D) AssertStr(vs ...string) func(s Scalar) (Scalar, error) {
	return d.assertStr(true, vs...)
}

// ValidateStr validates that actual value is one of given string values
func (d *D) ValidateStr(vs ...string) func(s Scalar) (Scalar, error) {
	return d.assertStr(false, vs...)
}

// Validate/Assert U

func (d *D) assertU(assert bool, vs ...uint64) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualU()
		for _, b := range vs {
			if a == b {
				s.Description = "valid"
				return s, nil
			}
		}
		s.Description = "invalid"
		if assert && !d.Options.Force {
			return s, errors.New("failed to assert U")
		}
		return s, nil
	}
}

// AssertU asserts that actual value is one of given uint64 values
func (d *D) AssertU(vs ...uint64) func(s Scalar) (Scalar, error) {
	return d.assertU(true, vs...)
}

// ValidateU validates that actual value is one of given uint64 values
func (d *D) ValidateU(vs ...uint64) func(s Scalar) (Scalar, error) {
	return d.assertU(false, vs...)
}

// Map Bool -> Scalar

type BoolToScalar map[bool]Scalar

// MapBoolToScalar maps Bool to a Scalar
func (d *D) MapBoolToScalar(m BoolToScalar) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualBool()
		if ns, ok := m[a]; ok {
			ns.Actual = a
			s = ns
		}
		return s, nil
	}
}

// Map Bool -> Bool

type BoolToBool map[bool]bool

// MapBoolToBoolSym maps Bool to a symolic Bool
func (d *D) MapBoolToBoolSym(m BoolToBool) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualBool()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Bool -> F

type BoolToF map[bool]float64

// MapBoolToFSym maps Bool to a symolic F
func (d *D) MapBoolToFSym(m BoolToF) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualBool()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Bool -> S

type BoolToS map[bool]int64

// MapBoolToSSym maps Bool to a symolic S
func (d *D) MapBoolToSSym(m BoolToS) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualBool()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Bool -> Str

type BoolToStr map[bool]string

// MapBoolToStrSym maps Bool to a symolic Str
func (d *D) MapBoolToStrSym(m BoolToStr) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualBool()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Bool -> U

type BoolToU map[bool]uint64

// MapBoolToUSym maps Bool to a symolic U
func (d *D) MapBoolToUSym(m BoolToU) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualBool()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map S -> Scalar

type SToScalar map[int64]Scalar

// MapSToScalar maps S to a Scalar
func (d *D) MapSToScalar(m SToScalar) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualS()
		if ns, ok := m[a]; ok {
			ns.Actual = a
			s = ns
		}
		return s, nil
	}
}

// Map S -> Bool

type SToBool map[int64]bool

// MapSToBoolSym maps S to a symolic Bool
func (d *D) MapSToBoolSym(m SToBool) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualS()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map S -> F

type SToF map[int64]float64

// MapSToFSym maps S to a symolic F
func (d *D) MapSToFSym(m SToF) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualS()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map S -> S

type SToS map[int64]int64

// MapSToSSym maps S to a symolic S
func (d *D) MapSToSSym(m SToS) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualS()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map S -> Str

type SToStr map[int64]string

// MapSToStrSym maps S to a symolic Str
func (d *D) MapSToStrSym(m SToStr) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualS()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map S -> U

type SToU map[int64]uint64

// MapSToUSym maps S to a symolic U
func (d *D) MapSToUSym(m SToU) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualS()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Str -> Scalar

type StrToScalar map[string]Scalar

// MapStrToScalar maps Str to a Scalar
func (d *D) MapStrToScalar(m StrToScalar) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualStr()
		if ns, ok := m[a]; ok {
			ns.Actual = a
			s = ns
		}
		return s, nil
	}
}

// Map Str -> Bool

type StrToBool map[string]bool

// MapStrToBoolSym maps Str to a symolic Bool
func (d *D) MapStrToBoolSym(m StrToBool) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualStr()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Str -> F

type StrToF map[string]float64

// MapStrToFSym maps Str to a symolic F
func (d *D) MapStrToFSym(m StrToF) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualStr()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Str -> S

type StrToS map[string]int64

// MapStrToSSym maps Str to a symolic S
func (d *D) MapStrToSSym(m StrToS) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualStr()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Str -> Str

type StrToStr map[string]string

// MapStrToStrSym maps Str to a symolic Str
func (d *D) MapStrToStrSym(m StrToStr) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualStr()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map Str -> U

type StrToU map[string]uint64

// MapStrToUSym maps Str to a symolic U
func (d *D) MapStrToUSym(m StrToU) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualStr()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map U -> Scalar

type UToScalar map[uint64]Scalar

// MapUToScalar maps U to a Scalar
func (d *D) MapUToScalar(m UToScalar) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		a := s.ActualU()
		if ns, ok := m[a]; ok {
			ns.Actual = a
			s = ns
		}
		return s, nil
	}
}

// Map U -> Bool

type UToBool map[uint64]bool

// MapUToBoolSym maps U to a symolic Bool
func (d *D) MapUToBoolSym(m UToBool) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualU()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map U -> F

type UToF map[uint64]float64

// MapUToFSym maps U to a symolic F
func (d *D) MapUToFSym(m UToF) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualU()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map U -> S

type UToS map[uint64]int64

// MapUToSSym maps U to a symolic S
func (d *D) MapUToSSym(m UToS) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualU()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map U -> Str

type UToStr map[uint64]string

// MapUToStrSym maps U to a symolic Str
func (d *D) MapUToStrSym(m UToStr) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualU()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Map U -> U

type UToU map[uint64]uint64

// MapUToUSym maps U to a symolic U
func (d *D) MapUToUSym(m UToU) func(s Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		if t, ok := m[s.ActualU()]; ok {
			s.Sym = t
		}
		return s, nil
	}
}

// Reader RawLen

// TryRawLen tries to read nBits raw bits
func (d *D) TryRawLen(nBits int64) (*bitio.Buffer, error) { return d.tryBitBuf(nBits) }

func (d *D) ScalarRawLen(nBits int64) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryBitBuf(nBits)
		s.Actual = v
		return s, err
	}
}

// RawLen reads nBits raw bits
func (d *D) RawLen(nBits int64) *bitio.Buffer {
	v, err := d.tryBitBuf(nBits)
	if err != nil {
		panic(IOError{Err: err, Op: "RawLen", Pos: d.Pos()})
	}
	return v
}

// TryFieldRawLen tries to add a field and read nBits raw bits
func (d *D) TryFieldRawLen(name string, nBits int64, sfns ...ScalarFn) (*bitio.Buffer, error) {
	v, err := d.TryFieldScalar(name, d.ScalarRawLen(nBits), sfns...)
	if err != nil {
		return nil, err
	}
	return v.ActualBitBuf(), err
}

// FieldRawLen adds a field and reads nBits raw bits
func (d *D) FieldRawLen(name string, nBits int64, sfns ...ScalarFn) *bitio.Buffer {
	v, err := d.TryFieldRawLen(name, nBits, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "RawLen", Pos: d.Pos()})
	}
	return v
}

// Reader Bool

// TryBool tries to read 1 bit boolean
func (d *D) TryBool() (bool, error) { return d.tryBool() }

func (d *D) ScalarBool() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryBool()
		s.Actual = v
		return s, err
	}
}

// Bool reads 1 bit boolean
func (d *D) Bool() bool {
	v, err := d.tryBool()
	if err != nil {
		panic(IOError{Err: err, Op: "Bool", Pos: d.Pos()})
	}
	return v
}

// TryFieldBool tries to add a field and read 1 bit boolean
func (d *D) TryFieldBool(name string, sfns ...ScalarFn) (bool, error) {
	return d.TryFieldBoolFn(name, (*D).TryBool, sfns...)
}

// FieldBool adds a field and reads 1 bit boolean
func (d *D) FieldBool(name string, sfns ...ScalarFn) bool {
	return d.FieldBoolFn(name, (*D).Bool, sfns...)
}

// Reader U

// TryU tries to read nBits bits unsigned integer in current endian
func (d *D) TryU(nBits int) (uint64, error) { return d.tryUE(nBits, d.Endian) }

func (d *D) ScalarU(nBits int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(nBits, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U reads nBits bits unsigned integer in current endian
func (d *D) U(nBits int) uint64 {
	v, err := d.tryUE(nBits, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U", Pos: d.Pos()})
	}
	return v
}

// TryFieldU tries to add a field and read nBits bits unsigned integer in current endian
func (d *D) TryFieldU(name string, nBits int, sfns ...ScalarFn) (uint64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarU(nBits), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualU(), err
}

// FieldU adds a field and reads nBits bits unsigned integer in current endian
func (d *D) FieldU(name string, nBits int, sfns ...ScalarFn) uint64 {
	v, err := d.TryFieldU(name, nBits, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "U", Pos: d.Pos()})
	}
	return v
}

// Reader UE

// TryUE tries to read nBits unsigned integer in specified endian
func (d *D) TryUE(nBits int, endian Endian) (uint64, error) { return d.tryUE(nBits, endian) }

func (d *D) ScalarUE(nBits int, endian Endian) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(nBits, endian)
		s.Actual = v
		return s, err
	}
}

// UE reads nBits unsigned integer in specified endian
func (d *D) UE(nBits int, endian Endian) uint64 {
	v, err := d.tryUE(nBits, endian)
	if err != nil {
		panic(IOError{Err: err, Op: "UE", Pos: d.Pos()})
	}
	return v
}

// TryFieldUE tries to add a field and read nBits unsigned integer in specified endian
func (d *D) TryFieldUE(name string, nBits int, endian Endian, sfns ...ScalarFn) (uint64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUE(nBits, endian), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualU(), err
}

// FieldUE adds a field and reads nBits unsigned integer in specified endian
func (d *D) FieldUE(name string, nBits int, endian Endian, sfns ...ScalarFn) uint64 {
	v, err := d.TryFieldUE(name, nBits, endian, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UE", Pos: d.Pos()})
	}
	return v
}

// Reader U1

// TryU1 tries to read 1 bit unsigned integer in current endian
func (d *D) TryU1() (uint64, error) { return d.tryUE(1, d.Endian) }

func (d *D) ScalarU1() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(1, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U1 reads 1 bit unsigned integer in current endian
func (d *D) U1() uint64 {
	v, err := d.tryUE(1, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U1", Pos: d.Pos()})
	}
	return v
}

// TryFieldU1 tries to add a field and read 1 bit unsigned integer in current endian
func (d *D) TryFieldU1(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU1, sfns...)
}

// FieldU1 adds a field and reads 1 bit unsigned integer in current endian
func (d *D) FieldU1(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U1, sfns...)
}

// Reader U2

// TryU2 tries to read 2 bit unsigned integer in current endian
func (d *D) TryU2() (uint64, error) { return d.tryUE(2, d.Endian) }

func (d *D) ScalarU2() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(2, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U2 reads 2 bit unsigned integer in current endian
func (d *D) U2() uint64 {
	v, err := d.tryUE(2, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U2", Pos: d.Pos()})
	}
	return v
}

// TryFieldU2 tries to add a field and read 2 bit unsigned integer in current endian
func (d *D) TryFieldU2(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU2, sfns...)
}

// FieldU2 adds a field and reads 2 bit unsigned integer in current endian
func (d *D) FieldU2(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U2, sfns...)
}

// Reader U3

// TryU3 tries to read 3 bit unsigned integer in current endian
func (d *D) TryU3() (uint64, error) { return d.tryUE(3, d.Endian) }

func (d *D) ScalarU3() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(3, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U3 reads 3 bit unsigned integer in current endian
func (d *D) U3() uint64 {
	v, err := d.tryUE(3, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U3", Pos: d.Pos()})
	}
	return v
}

// TryFieldU3 tries to add a field and read 3 bit unsigned integer in current endian
func (d *D) TryFieldU3(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU3, sfns...)
}

// FieldU3 adds a field and reads 3 bit unsigned integer in current endian
func (d *D) FieldU3(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U3, sfns...)
}

// Reader U4

// TryU4 tries to read 4 bit unsigned integer in current endian
func (d *D) TryU4() (uint64, error) { return d.tryUE(4, d.Endian) }

func (d *D) ScalarU4() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(4, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U4 reads 4 bit unsigned integer in current endian
func (d *D) U4() uint64 {
	v, err := d.tryUE(4, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U4", Pos: d.Pos()})
	}
	return v
}

// TryFieldU4 tries to add a field and read 4 bit unsigned integer in current endian
func (d *D) TryFieldU4(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU4, sfns...)
}

// FieldU4 adds a field and reads 4 bit unsigned integer in current endian
func (d *D) FieldU4(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U4, sfns...)
}

// Reader U5

// TryU5 tries to read 5 bit unsigned integer in current endian
func (d *D) TryU5() (uint64, error) { return d.tryUE(5, d.Endian) }

func (d *D) ScalarU5() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(5, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U5 reads 5 bit unsigned integer in current endian
func (d *D) U5() uint64 {
	v, err := d.tryUE(5, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U5", Pos: d.Pos()})
	}
	return v
}

// TryFieldU5 tries to add a field and read 5 bit unsigned integer in current endian
func (d *D) TryFieldU5(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU5, sfns...)
}

// FieldU5 adds a field and reads 5 bit unsigned integer in current endian
func (d *D) FieldU5(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U5, sfns...)
}

// Reader U6

// TryU6 tries to read 6 bit unsigned integer in current endian
func (d *D) TryU6() (uint64, error) { return d.tryUE(6, d.Endian) }

func (d *D) ScalarU6() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(6, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U6 reads 6 bit unsigned integer in current endian
func (d *D) U6() uint64 {
	v, err := d.tryUE(6, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U6", Pos: d.Pos()})
	}
	return v
}

// TryFieldU6 tries to add a field and read 6 bit unsigned integer in current endian
func (d *D) TryFieldU6(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU6, sfns...)
}

// FieldU6 adds a field and reads 6 bit unsigned integer in current endian
func (d *D) FieldU6(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U6, sfns...)
}

// Reader U7

// TryU7 tries to read 7 bit unsigned integer in current endian
func (d *D) TryU7() (uint64, error) { return d.tryUE(7, d.Endian) }

func (d *D) ScalarU7() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(7, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U7 reads 7 bit unsigned integer in current endian
func (d *D) U7() uint64 {
	v, err := d.tryUE(7, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U7", Pos: d.Pos()})
	}
	return v
}

// TryFieldU7 tries to add a field and read 7 bit unsigned integer in current endian
func (d *D) TryFieldU7(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU7, sfns...)
}

// FieldU7 adds a field and reads 7 bit unsigned integer in current endian
func (d *D) FieldU7(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U7, sfns...)
}

// Reader U8

// TryU8 tries to read 8 bit unsigned integer in current endian
func (d *D) TryU8() (uint64, error) { return d.tryUE(8, d.Endian) }

func (d *D) ScalarU8() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(8, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U8 reads 8 bit unsigned integer in current endian
func (d *D) U8() uint64 {
	v, err := d.tryUE(8, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U8", Pos: d.Pos()})
	}
	return v
}

// TryFieldU8 tries to add a field and read 8 bit unsigned integer in current endian
func (d *D) TryFieldU8(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU8, sfns...)
}

// FieldU8 adds a field and reads 8 bit unsigned integer in current endian
func (d *D) FieldU8(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U8, sfns...)
}

// Reader U9

// TryU9 tries to read 9 bit unsigned integer in current endian
func (d *D) TryU9() (uint64, error) { return d.tryUE(9, d.Endian) }

func (d *D) ScalarU9() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(9, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U9 reads 9 bit unsigned integer in current endian
func (d *D) U9() uint64 {
	v, err := d.tryUE(9, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U9", Pos: d.Pos()})
	}
	return v
}

// TryFieldU9 tries to add a field and read 9 bit unsigned integer in current endian
func (d *D) TryFieldU9(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU9, sfns...)
}

// FieldU9 adds a field and reads 9 bit unsigned integer in current endian
func (d *D) FieldU9(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U9, sfns...)
}

// Reader U10

// TryU10 tries to read 10 bit unsigned integer in current endian
func (d *D) TryU10() (uint64, error) { return d.tryUE(10, d.Endian) }

func (d *D) ScalarU10() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(10, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U10 reads 10 bit unsigned integer in current endian
func (d *D) U10() uint64 {
	v, err := d.tryUE(10, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U10", Pos: d.Pos()})
	}
	return v
}

// TryFieldU10 tries to add a field and read 10 bit unsigned integer in current endian
func (d *D) TryFieldU10(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU10, sfns...)
}

// FieldU10 adds a field and reads 10 bit unsigned integer in current endian
func (d *D) FieldU10(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U10, sfns...)
}

// Reader U11

// TryU11 tries to read 11 bit unsigned integer in current endian
func (d *D) TryU11() (uint64, error) { return d.tryUE(11, d.Endian) }

func (d *D) ScalarU11() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(11, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U11 reads 11 bit unsigned integer in current endian
func (d *D) U11() uint64 {
	v, err := d.tryUE(11, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U11", Pos: d.Pos()})
	}
	return v
}

// TryFieldU11 tries to add a field and read 11 bit unsigned integer in current endian
func (d *D) TryFieldU11(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU11, sfns...)
}

// FieldU11 adds a field and reads 11 bit unsigned integer in current endian
func (d *D) FieldU11(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U11, sfns...)
}

// Reader U12

// TryU12 tries to read 12 bit unsigned integer in current endian
func (d *D) TryU12() (uint64, error) { return d.tryUE(12, d.Endian) }

func (d *D) ScalarU12() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(12, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U12 reads 12 bit unsigned integer in current endian
func (d *D) U12() uint64 {
	v, err := d.tryUE(12, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U12", Pos: d.Pos()})
	}
	return v
}

// TryFieldU12 tries to add a field and read 12 bit unsigned integer in current endian
func (d *D) TryFieldU12(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU12, sfns...)
}

// FieldU12 adds a field and reads 12 bit unsigned integer in current endian
func (d *D) FieldU12(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U12, sfns...)
}

// Reader U13

// TryU13 tries to read 13 bit unsigned integer in current endian
func (d *D) TryU13() (uint64, error) { return d.tryUE(13, d.Endian) }

func (d *D) ScalarU13() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(13, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U13 reads 13 bit unsigned integer in current endian
func (d *D) U13() uint64 {
	v, err := d.tryUE(13, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U13", Pos: d.Pos()})
	}
	return v
}

// TryFieldU13 tries to add a field and read 13 bit unsigned integer in current endian
func (d *D) TryFieldU13(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU13, sfns...)
}

// FieldU13 adds a field and reads 13 bit unsigned integer in current endian
func (d *D) FieldU13(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U13, sfns...)
}

// Reader U14

// TryU14 tries to read 14 bit unsigned integer in current endian
func (d *D) TryU14() (uint64, error) { return d.tryUE(14, d.Endian) }

func (d *D) ScalarU14() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(14, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U14 reads 14 bit unsigned integer in current endian
func (d *D) U14() uint64 {
	v, err := d.tryUE(14, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U14", Pos: d.Pos()})
	}
	return v
}

// TryFieldU14 tries to add a field and read 14 bit unsigned integer in current endian
func (d *D) TryFieldU14(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU14, sfns...)
}

// FieldU14 adds a field and reads 14 bit unsigned integer in current endian
func (d *D) FieldU14(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U14, sfns...)
}

// Reader U15

// TryU15 tries to read 15 bit unsigned integer in current endian
func (d *D) TryU15() (uint64, error) { return d.tryUE(15, d.Endian) }

func (d *D) ScalarU15() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(15, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U15 reads 15 bit unsigned integer in current endian
func (d *D) U15() uint64 {
	v, err := d.tryUE(15, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U15", Pos: d.Pos()})
	}
	return v
}

// TryFieldU15 tries to add a field and read 15 bit unsigned integer in current endian
func (d *D) TryFieldU15(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU15, sfns...)
}

// FieldU15 adds a field and reads 15 bit unsigned integer in current endian
func (d *D) FieldU15(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U15, sfns...)
}

// Reader U16

// TryU16 tries to read 16 bit unsigned integer in current endian
func (d *D) TryU16() (uint64, error) { return d.tryUE(16, d.Endian) }

func (d *D) ScalarU16() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(16, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U16 reads 16 bit unsigned integer in current endian
func (d *D) U16() uint64 {
	v, err := d.tryUE(16, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U16", Pos: d.Pos()})
	}
	return v
}

// TryFieldU16 tries to add a field and read 16 bit unsigned integer in current endian
func (d *D) TryFieldU16(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU16, sfns...)
}

// FieldU16 adds a field and reads 16 bit unsigned integer in current endian
func (d *D) FieldU16(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U16, sfns...)
}

// Reader U17

// TryU17 tries to read 17 bit unsigned integer in current endian
func (d *D) TryU17() (uint64, error) { return d.tryUE(17, d.Endian) }

func (d *D) ScalarU17() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(17, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U17 reads 17 bit unsigned integer in current endian
func (d *D) U17() uint64 {
	v, err := d.tryUE(17, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U17", Pos: d.Pos()})
	}
	return v
}

// TryFieldU17 tries to add a field and read 17 bit unsigned integer in current endian
func (d *D) TryFieldU17(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU17, sfns...)
}

// FieldU17 adds a field and reads 17 bit unsigned integer in current endian
func (d *D) FieldU17(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U17, sfns...)
}

// Reader U18

// TryU18 tries to read 18 bit unsigned integer in current endian
func (d *D) TryU18() (uint64, error) { return d.tryUE(18, d.Endian) }

func (d *D) ScalarU18() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(18, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U18 reads 18 bit unsigned integer in current endian
func (d *D) U18() uint64 {
	v, err := d.tryUE(18, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U18", Pos: d.Pos()})
	}
	return v
}

// TryFieldU18 tries to add a field and read 18 bit unsigned integer in current endian
func (d *D) TryFieldU18(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU18, sfns...)
}

// FieldU18 adds a field and reads 18 bit unsigned integer in current endian
func (d *D) FieldU18(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U18, sfns...)
}

// Reader U19

// TryU19 tries to read 19 bit unsigned integer in current endian
func (d *D) TryU19() (uint64, error) { return d.tryUE(19, d.Endian) }

func (d *D) ScalarU19() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(19, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U19 reads 19 bit unsigned integer in current endian
func (d *D) U19() uint64 {
	v, err := d.tryUE(19, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U19", Pos: d.Pos()})
	}
	return v
}

// TryFieldU19 tries to add a field and read 19 bit unsigned integer in current endian
func (d *D) TryFieldU19(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU19, sfns...)
}

// FieldU19 adds a field and reads 19 bit unsigned integer in current endian
func (d *D) FieldU19(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U19, sfns...)
}

// Reader U20

// TryU20 tries to read 20 bit unsigned integer in current endian
func (d *D) TryU20() (uint64, error) { return d.tryUE(20, d.Endian) }

func (d *D) ScalarU20() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(20, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U20 reads 20 bit unsigned integer in current endian
func (d *D) U20() uint64 {
	v, err := d.tryUE(20, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U20", Pos: d.Pos()})
	}
	return v
}

// TryFieldU20 tries to add a field and read 20 bit unsigned integer in current endian
func (d *D) TryFieldU20(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU20, sfns...)
}

// FieldU20 adds a field and reads 20 bit unsigned integer in current endian
func (d *D) FieldU20(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U20, sfns...)
}

// Reader U21

// TryU21 tries to read 21 bit unsigned integer in current endian
func (d *D) TryU21() (uint64, error) { return d.tryUE(21, d.Endian) }

func (d *D) ScalarU21() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(21, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U21 reads 21 bit unsigned integer in current endian
func (d *D) U21() uint64 {
	v, err := d.tryUE(21, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U21", Pos: d.Pos()})
	}
	return v
}

// TryFieldU21 tries to add a field and read 21 bit unsigned integer in current endian
func (d *D) TryFieldU21(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU21, sfns...)
}

// FieldU21 adds a field and reads 21 bit unsigned integer in current endian
func (d *D) FieldU21(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U21, sfns...)
}

// Reader U22

// TryU22 tries to read 22 bit unsigned integer in current endian
func (d *D) TryU22() (uint64, error) { return d.tryUE(22, d.Endian) }

func (d *D) ScalarU22() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(22, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U22 reads 22 bit unsigned integer in current endian
func (d *D) U22() uint64 {
	v, err := d.tryUE(22, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U22", Pos: d.Pos()})
	}
	return v
}

// TryFieldU22 tries to add a field and read 22 bit unsigned integer in current endian
func (d *D) TryFieldU22(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU22, sfns...)
}

// FieldU22 adds a field and reads 22 bit unsigned integer in current endian
func (d *D) FieldU22(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U22, sfns...)
}

// Reader U23

// TryU23 tries to read 23 bit unsigned integer in current endian
func (d *D) TryU23() (uint64, error) { return d.tryUE(23, d.Endian) }

func (d *D) ScalarU23() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(23, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U23 reads 23 bit unsigned integer in current endian
func (d *D) U23() uint64 {
	v, err := d.tryUE(23, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U23", Pos: d.Pos()})
	}
	return v
}

// TryFieldU23 tries to add a field and read 23 bit unsigned integer in current endian
func (d *D) TryFieldU23(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU23, sfns...)
}

// FieldU23 adds a field and reads 23 bit unsigned integer in current endian
func (d *D) FieldU23(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U23, sfns...)
}

// Reader U24

// TryU24 tries to read 24 bit unsigned integer in current endian
func (d *D) TryU24() (uint64, error) { return d.tryUE(24, d.Endian) }

func (d *D) ScalarU24() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(24, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U24 reads 24 bit unsigned integer in current endian
func (d *D) U24() uint64 {
	v, err := d.tryUE(24, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U24", Pos: d.Pos()})
	}
	return v
}

// TryFieldU24 tries to add a field and read 24 bit unsigned integer in current endian
func (d *D) TryFieldU24(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU24, sfns...)
}

// FieldU24 adds a field and reads 24 bit unsigned integer in current endian
func (d *D) FieldU24(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U24, sfns...)
}

// Reader U25

// TryU25 tries to read 25 bit unsigned integer in current endian
func (d *D) TryU25() (uint64, error) { return d.tryUE(25, d.Endian) }

func (d *D) ScalarU25() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(25, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U25 reads 25 bit unsigned integer in current endian
func (d *D) U25() uint64 {
	v, err := d.tryUE(25, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U25", Pos: d.Pos()})
	}
	return v
}

// TryFieldU25 tries to add a field and read 25 bit unsigned integer in current endian
func (d *D) TryFieldU25(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU25, sfns...)
}

// FieldU25 adds a field and reads 25 bit unsigned integer in current endian
func (d *D) FieldU25(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U25, sfns...)
}

// Reader U26

// TryU26 tries to read 26 bit unsigned integer in current endian
func (d *D) TryU26() (uint64, error) { return d.tryUE(26, d.Endian) }

func (d *D) ScalarU26() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(26, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U26 reads 26 bit unsigned integer in current endian
func (d *D) U26() uint64 {
	v, err := d.tryUE(26, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U26", Pos: d.Pos()})
	}
	return v
}

// TryFieldU26 tries to add a field and read 26 bit unsigned integer in current endian
func (d *D) TryFieldU26(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU26, sfns...)
}

// FieldU26 adds a field and reads 26 bit unsigned integer in current endian
func (d *D) FieldU26(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U26, sfns...)
}

// Reader U27

// TryU27 tries to read 27 bit unsigned integer in current endian
func (d *D) TryU27() (uint64, error) { return d.tryUE(27, d.Endian) }

func (d *D) ScalarU27() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(27, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U27 reads 27 bit unsigned integer in current endian
func (d *D) U27() uint64 {
	v, err := d.tryUE(27, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U27", Pos: d.Pos()})
	}
	return v
}

// TryFieldU27 tries to add a field and read 27 bit unsigned integer in current endian
func (d *D) TryFieldU27(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU27, sfns...)
}

// FieldU27 adds a field and reads 27 bit unsigned integer in current endian
func (d *D) FieldU27(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U27, sfns...)
}

// Reader U28

// TryU28 tries to read 28 bit unsigned integer in current endian
func (d *D) TryU28() (uint64, error) { return d.tryUE(28, d.Endian) }

func (d *D) ScalarU28() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(28, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U28 reads 28 bit unsigned integer in current endian
func (d *D) U28() uint64 {
	v, err := d.tryUE(28, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U28", Pos: d.Pos()})
	}
	return v
}

// TryFieldU28 tries to add a field and read 28 bit unsigned integer in current endian
func (d *D) TryFieldU28(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU28, sfns...)
}

// FieldU28 adds a field and reads 28 bit unsigned integer in current endian
func (d *D) FieldU28(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U28, sfns...)
}

// Reader U29

// TryU29 tries to read 29 bit unsigned integer in current endian
func (d *D) TryU29() (uint64, error) { return d.tryUE(29, d.Endian) }

func (d *D) ScalarU29() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(29, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U29 reads 29 bit unsigned integer in current endian
func (d *D) U29() uint64 {
	v, err := d.tryUE(29, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U29", Pos: d.Pos()})
	}
	return v
}

// TryFieldU29 tries to add a field and read 29 bit unsigned integer in current endian
func (d *D) TryFieldU29(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU29, sfns...)
}

// FieldU29 adds a field and reads 29 bit unsigned integer in current endian
func (d *D) FieldU29(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U29, sfns...)
}

// Reader U30

// TryU30 tries to read 30 bit unsigned integer in current endian
func (d *D) TryU30() (uint64, error) { return d.tryUE(30, d.Endian) }

func (d *D) ScalarU30() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(30, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U30 reads 30 bit unsigned integer in current endian
func (d *D) U30() uint64 {
	v, err := d.tryUE(30, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U30", Pos: d.Pos()})
	}
	return v
}

// TryFieldU30 tries to add a field and read 30 bit unsigned integer in current endian
func (d *D) TryFieldU30(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU30, sfns...)
}

// FieldU30 adds a field and reads 30 bit unsigned integer in current endian
func (d *D) FieldU30(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U30, sfns...)
}

// Reader U31

// TryU31 tries to read 31 bit unsigned integer in current endian
func (d *D) TryU31() (uint64, error) { return d.tryUE(31, d.Endian) }

func (d *D) ScalarU31() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(31, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U31 reads 31 bit unsigned integer in current endian
func (d *D) U31() uint64 {
	v, err := d.tryUE(31, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U31", Pos: d.Pos()})
	}
	return v
}

// TryFieldU31 tries to add a field and read 31 bit unsigned integer in current endian
func (d *D) TryFieldU31(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU31, sfns...)
}

// FieldU31 adds a field and reads 31 bit unsigned integer in current endian
func (d *D) FieldU31(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U31, sfns...)
}

// Reader U32

// TryU32 tries to read 32 bit unsigned integer in current endian
func (d *D) TryU32() (uint64, error) { return d.tryUE(32, d.Endian) }

func (d *D) ScalarU32() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(32, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U32 reads 32 bit unsigned integer in current endian
func (d *D) U32() uint64 {
	v, err := d.tryUE(32, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U32", Pos: d.Pos()})
	}
	return v
}

// TryFieldU32 tries to add a field and read 32 bit unsigned integer in current endian
func (d *D) TryFieldU32(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU32, sfns...)
}

// FieldU32 adds a field and reads 32 bit unsigned integer in current endian
func (d *D) FieldU32(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U32, sfns...)
}

// Reader U33

// TryU33 tries to read 33 bit unsigned integer in current endian
func (d *D) TryU33() (uint64, error) { return d.tryUE(33, d.Endian) }

func (d *D) ScalarU33() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(33, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U33 reads 33 bit unsigned integer in current endian
func (d *D) U33() uint64 {
	v, err := d.tryUE(33, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U33", Pos: d.Pos()})
	}
	return v
}

// TryFieldU33 tries to add a field and read 33 bit unsigned integer in current endian
func (d *D) TryFieldU33(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU33, sfns...)
}

// FieldU33 adds a field and reads 33 bit unsigned integer in current endian
func (d *D) FieldU33(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U33, sfns...)
}

// Reader U34

// TryU34 tries to read 34 bit unsigned integer in current endian
func (d *D) TryU34() (uint64, error) { return d.tryUE(34, d.Endian) }

func (d *D) ScalarU34() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(34, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U34 reads 34 bit unsigned integer in current endian
func (d *D) U34() uint64 {
	v, err := d.tryUE(34, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U34", Pos: d.Pos()})
	}
	return v
}

// TryFieldU34 tries to add a field and read 34 bit unsigned integer in current endian
func (d *D) TryFieldU34(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU34, sfns...)
}

// FieldU34 adds a field and reads 34 bit unsigned integer in current endian
func (d *D) FieldU34(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U34, sfns...)
}

// Reader U35

// TryU35 tries to read 35 bit unsigned integer in current endian
func (d *D) TryU35() (uint64, error) { return d.tryUE(35, d.Endian) }

func (d *D) ScalarU35() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(35, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U35 reads 35 bit unsigned integer in current endian
func (d *D) U35() uint64 {
	v, err := d.tryUE(35, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U35", Pos: d.Pos()})
	}
	return v
}

// TryFieldU35 tries to add a field and read 35 bit unsigned integer in current endian
func (d *D) TryFieldU35(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU35, sfns...)
}

// FieldU35 adds a field and reads 35 bit unsigned integer in current endian
func (d *D) FieldU35(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U35, sfns...)
}

// Reader U36

// TryU36 tries to read 36 bit unsigned integer in current endian
func (d *D) TryU36() (uint64, error) { return d.tryUE(36, d.Endian) }

func (d *D) ScalarU36() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(36, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U36 reads 36 bit unsigned integer in current endian
func (d *D) U36() uint64 {
	v, err := d.tryUE(36, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U36", Pos: d.Pos()})
	}
	return v
}

// TryFieldU36 tries to add a field and read 36 bit unsigned integer in current endian
func (d *D) TryFieldU36(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU36, sfns...)
}

// FieldU36 adds a field and reads 36 bit unsigned integer in current endian
func (d *D) FieldU36(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U36, sfns...)
}

// Reader U37

// TryU37 tries to read 37 bit unsigned integer in current endian
func (d *D) TryU37() (uint64, error) { return d.tryUE(37, d.Endian) }

func (d *D) ScalarU37() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(37, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U37 reads 37 bit unsigned integer in current endian
func (d *D) U37() uint64 {
	v, err := d.tryUE(37, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U37", Pos: d.Pos()})
	}
	return v
}

// TryFieldU37 tries to add a field and read 37 bit unsigned integer in current endian
func (d *D) TryFieldU37(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU37, sfns...)
}

// FieldU37 adds a field and reads 37 bit unsigned integer in current endian
func (d *D) FieldU37(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U37, sfns...)
}

// Reader U38

// TryU38 tries to read 38 bit unsigned integer in current endian
func (d *D) TryU38() (uint64, error) { return d.tryUE(38, d.Endian) }

func (d *D) ScalarU38() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(38, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U38 reads 38 bit unsigned integer in current endian
func (d *D) U38() uint64 {
	v, err := d.tryUE(38, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U38", Pos: d.Pos()})
	}
	return v
}

// TryFieldU38 tries to add a field and read 38 bit unsigned integer in current endian
func (d *D) TryFieldU38(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU38, sfns...)
}

// FieldU38 adds a field and reads 38 bit unsigned integer in current endian
func (d *D) FieldU38(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U38, sfns...)
}

// Reader U39

// TryU39 tries to read 39 bit unsigned integer in current endian
func (d *D) TryU39() (uint64, error) { return d.tryUE(39, d.Endian) }

func (d *D) ScalarU39() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(39, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U39 reads 39 bit unsigned integer in current endian
func (d *D) U39() uint64 {
	v, err := d.tryUE(39, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U39", Pos: d.Pos()})
	}
	return v
}

// TryFieldU39 tries to add a field and read 39 bit unsigned integer in current endian
func (d *D) TryFieldU39(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU39, sfns...)
}

// FieldU39 adds a field and reads 39 bit unsigned integer in current endian
func (d *D) FieldU39(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U39, sfns...)
}

// Reader U40

// TryU40 tries to read 40 bit unsigned integer in current endian
func (d *D) TryU40() (uint64, error) { return d.tryUE(40, d.Endian) }

func (d *D) ScalarU40() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(40, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U40 reads 40 bit unsigned integer in current endian
func (d *D) U40() uint64 {
	v, err := d.tryUE(40, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U40", Pos: d.Pos()})
	}
	return v
}

// TryFieldU40 tries to add a field and read 40 bit unsigned integer in current endian
func (d *D) TryFieldU40(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU40, sfns...)
}

// FieldU40 adds a field and reads 40 bit unsigned integer in current endian
func (d *D) FieldU40(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U40, sfns...)
}

// Reader U41

// TryU41 tries to read 41 bit unsigned integer in current endian
func (d *D) TryU41() (uint64, error) { return d.tryUE(41, d.Endian) }

func (d *D) ScalarU41() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(41, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U41 reads 41 bit unsigned integer in current endian
func (d *D) U41() uint64 {
	v, err := d.tryUE(41, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U41", Pos: d.Pos()})
	}
	return v
}

// TryFieldU41 tries to add a field and read 41 bit unsigned integer in current endian
func (d *D) TryFieldU41(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU41, sfns...)
}

// FieldU41 adds a field and reads 41 bit unsigned integer in current endian
func (d *D) FieldU41(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U41, sfns...)
}

// Reader U42

// TryU42 tries to read 42 bit unsigned integer in current endian
func (d *D) TryU42() (uint64, error) { return d.tryUE(42, d.Endian) }

func (d *D) ScalarU42() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(42, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U42 reads 42 bit unsigned integer in current endian
func (d *D) U42() uint64 {
	v, err := d.tryUE(42, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U42", Pos: d.Pos()})
	}
	return v
}

// TryFieldU42 tries to add a field and read 42 bit unsigned integer in current endian
func (d *D) TryFieldU42(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU42, sfns...)
}

// FieldU42 adds a field and reads 42 bit unsigned integer in current endian
func (d *D) FieldU42(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U42, sfns...)
}

// Reader U43

// TryU43 tries to read 43 bit unsigned integer in current endian
func (d *D) TryU43() (uint64, error) { return d.tryUE(43, d.Endian) }

func (d *D) ScalarU43() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(43, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U43 reads 43 bit unsigned integer in current endian
func (d *D) U43() uint64 {
	v, err := d.tryUE(43, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U43", Pos: d.Pos()})
	}
	return v
}

// TryFieldU43 tries to add a field and read 43 bit unsigned integer in current endian
func (d *D) TryFieldU43(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU43, sfns...)
}

// FieldU43 adds a field and reads 43 bit unsigned integer in current endian
func (d *D) FieldU43(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U43, sfns...)
}

// Reader U44

// TryU44 tries to read 44 bit unsigned integer in current endian
func (d *D) TryU44() (uint64, error) { return d.tryUE(44, d.Endian) }

func (d *D) ScalarU44() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(44, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U44 reads 44 bit unsigned integer in current endian
func (d *D) U44() uint64 {
	v, err := d.tryUE(44, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U44", Pos: d.Pos()})
	}
	return v
}

// TryFieldU44 tries to add a field and read 44 bit unsigned integer in current endian
func (d *D) TryFieldU44(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU44, sfns...)
}

// FieldU44 adds a field and reads 44 bit unsigned integer in current endian
func (d *D) FieldU44(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U44, sfns...)
}

// Reader U45

// TryU45 tries to read 45 bit unsigned integer in current endian
func (d *D) TryU45() (uint64, error) { return d.tryUE(45, d.Endian) }

func (d *D) ScalarU45() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(45, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U45 reads 45 bit unsigned integer in current endian
func (d *D) U45() uint64 {
	v, err := d.tryUE(45, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U45", Pos: d.Pos()})
	}
	return v
}

// TryFieldU45 tries to add a field and read 45 bit unsigned integer in current endian
func (d *D) TryFieldU45(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU45, sfns...)
}

// FieldU45 adds a field and reads 45 bit unsigned integer in current endian
func (d *D) FieldU45(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U45, sfns...)
}

// Reader U46

// TryU46 tries to read 46 bit unsigned integer in current endian
func (d *D) TryU46() (uint64, error) { return d.tryUE(46, d.Endian) }

func (d *D) ScalarU46() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(46, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U46 reads 46 bit unsigned integer in current endian
func (d *D) U46() uint64 {
	v, err := d.tryUE(46, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U46", Pos: d.Pos()})
	}
	return v
}

// TryFieldU46 tries to add a field and read 46 bit unsigned integer in current endian
func (d *D) TryFieldU46(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU46, sfns...)
}

// FieldU46 adds a field and reads 46 bit unsigned integer in current endian
func (d *D) FieldU46(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U46, sfns...)
}

// Reader U47

// TryU47 tries to read 47 bit unsigned integer in current endian
func (d *D) TryU47() (uint64, error) { return d.tryUE(47, d.Endian) }

func (d *D) ScalarU47() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(47, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U47 reads 47 bit unsigned integer in current endian
func (d *D) U47() uint64 {
	v, err := d.tryUE(47, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U47", Pos: d.Pos()})
	}
	return v
}

// TryFieldU47 tries to add a field and read 47 bit unsigned integer in current endian
func (d *D) TryFieldU47(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU47, sfns...)
}

// FieldU47 adds a field and reads 47 bit unsigned integer in current endian
func (d *D) FieldU47(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U47, sfns...)
}

// Reader U48

// TryU48 tries to read 48 bit unsigned integer in current endian
func (d *D) TryU48() (uint64, error) { return d.tryUE(48, d.Endian) }

func (d *D) ScalarU48() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(48, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U48 reads 48 bit unsigned integer in current endian
func (d *D) U48() uint64 {
	v, err := d.tryUE(48, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U48", Pos: d.Pos()})
	}
	return v
}

// TryFieldU48 tries to add a field and read 48 bit unsigned integer in current endian
func (d *D) TryFieldU48(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU48, sfns...)
}

// FieldU48 adds a field and reads 48 bit unsigned integer in current endian
func (d *D) FieldU48(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U48, sfns...)
}

// Reader U49

// TryU49 tries to read 49 bit unsigned integer in current endian
func (d *D) TryU49() (uint64, error) { return d.tryUE(49, d.Endian) }

func (d *D) ScalarU49() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(49, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U49 reads 49 bit unsigned integer in current endian
func (d *D) U49() uint64 {
	v, err := d.tryUE(49, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U49", Pos: d.Pos()})
	}
	return v
}

// TryFieldU49 tries to add a field and read 49 bit unsigned integer in current endian
func (d *D) TryFieldU49(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU49, sfns...)
}

// FieldU49 adds a field and reads 49 bit unsigned integer in current endian
func (d *D) FieldU49(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U49, sfns...)
}

// Reader U50

// TryU50 tries to read 50 bit unsigned integer in current endian
func (d *D) TryU50() (uint64, error) { return d.tryUE(50, d.Endian) }

func (d *D) ScalarU50() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(50, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U50 reads 50 bit unsigned integer in current endian
func (d *D) U50() uint64 {
	v, err := d.tryUE(50, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U50", Pos: d.Pos()})
	}
	return v
}

// TryFieldU50 tries to add a field and read 50 bit unsigned integer in current endian
func (d *D) TryFieldU50(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU50, sfns...)
}

// FieldU50 adds a field and reads 50 bit unsigned integer in current endian
func (d *D) FieldU50(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U50, sfns...)
}

// Reader U51

// TryU51 tries to read 51 bit unsigned integer in current endian
func (d *D) TryU51() (uint64, error) { return d.tryUE(51, d.Endian) }

func (d *D) ScalarU51() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(51, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U51 reads 51 bit unsigned integer in current endian
func (d *D) U51() uint64 {
	v, err := d.tryUE(51, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U51", Pos: d.Pos()})
	}
	return v
}

// TryFieldU51 tries to add a field and read 51 bit unsigned integer in current endian
func (d *D) TryFieldU51(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU51, sfns...)
}

// FieldU51 adds a field and reads 51 bit unsigned integer in current endian
func (d *D) FieldU51(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U51, sfns...)
}

// Reader U52

// TryU52 tries to read 52 bit unsigned integer in current endian
func (d *D) TryU52() (uint64, error) { return d.tryUE(52, d.Endian) }

func (d *D) ScalarU52() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(52, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U52 reads 52 bit unsigned integer in current endian
func (d *D) U52() uint64 {
	v, err := d.tryUE(52, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U52", Pos: d.Pos()})
	}
	return v
}

// TryFieldU52 tries to add a field and read 52 bit unsigned integer in current endian
func (d *D) TryFieldU52(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU52, sfns...)
}

// FieldU52 adds a field and reads 52 bit unsigned integer in current endian
func (d *D) FieldU52(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U52, sfns...)
}

// Reader U53

// TryU53 tries to read 53 bit unsigned integer in current endian
func (d *D) TryU53() (uint64, error) { return d.tryUE(53, d.Endian) }

func (d *D) ScalarU53() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(53, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U53 reads 53 bit unsigned integer in current endian
func (d *D) U53() uint64 {
	v, err := d.tryUE(53, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U53", Pos: d.Pos()})
	}
	return v
}

// TryFieldU53 tries to add a field and read 53 bit unsigned integer in current endian
func (d *D) TryFieldU53(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU53, sfns...)
}

// FieldU53 adds a field and reads 53 bit unsigned integer in current endian
func (d *D) FieldU53(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U53, sfns...)
}

// Reader U54

// TryU54 tries to read 54 bit unsigned integer in current endian
func (d *D) TryU54() (uint64, error) { return d.tryUE(54, d.Endian) }

func (d *D) ScalarU54() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(54, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U54 reads 54 bit unsigned integer in current endian
func (d *D) U54() uint64 {
	v, err := d.tryUE(54, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U54", Pos: d.Pos()})
	}
	return v
}

// TryFieldU54 tries to add a field and read 54 bit unsigned integer in current endian
func (d *D) TryFieldU54(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU54, sfns...)
}

// FieldU54 adds a field and reads 54 bit unsigned integer in current endian
func (d *D) FieldU54(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U54, sfns...)
}

// Reader U55

// TryU55 tries to read 55 bit unsigned integer in current endian
func (d *D) TryU55() (uint64, error) { return d.tryUE(55, d.Endian) }

func (d *D) ScalarU55() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(55, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U55 reads 55 bit unsigned integer in current endian
func (d *D) U55() uint64 {
	v, err := d.tryUE(55, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U55", Pos: d.Pos()})
	}
	return v
}

// TryFieldU55 tries to add a field and read 55 bit unsigned integer in current endian
func (d *D) TryFieldU55(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU55, sfns...)
}

// FieldU55 adds a field and reads 55 bit unsigned integer in current endian
func (d *D) FieldU55(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U55, sfns...)
}

// Reader U56

// TryU56 tries to read 56 bit unsigned integer in current endian
func (d *D) TryU56() (uint64, error) { return d.tryUE(56, d.Endian) }

func (d *D) ScalarU56() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(56, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U56 reads 56 bit unsigned integer in current endian
func (d *D) U56() uint64 {
	v, err := d.tryUE(56, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U56", Pos: d.Pos()})
	}
	return v
}

// TryFieldU56 tries to add a field and read 56 bit unsigned integer in current endian
func (d *D) TryFieldU56(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU56, sfns...)
}

// FieldU56 adds a field and reads 56 bit unsigned integer in current endian
func (d *D) FieldU56(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U56, sfns...)
}

// Reader U57

// TryU57 tries to read 57 bit unsigned integer in current endian
func (d *D) TryU57() (uint64, error) { return d.tryUE(57, d.Endian) }

func (d *D) ScalarU57() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(57, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U57 reads 57 bit unsigned integer in current endian
func (d *D) U57() uint64 {
	v, err := d.tryUE(57, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U57", Pos: d.Pos()})
	}
	return v
}

// TryFieldU57 tries to add a field and read 57 bit unsigned integer in current endian
func (d *D) TryFieldU57(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU57, sfns...)
}

// FieldU57 adds a field and reads 57 bit unsigned integer in current endian
func (d *D) FieldU57(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U57, sfns...)
}

// Reader U58

// TryU58 tries to read 58 bit unsigned integer in current endian
func (d *D) TryU58() (uint64, error) { return d.tryUE(58, d.Endian) }

func (d *D) ScalarU58() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(58, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U58 reads 58 bit unsigned integer in current endian
func (d *D) U58() uint64 {
	v, err := d.tryUE(58, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U58", Pos: d.Pos()})
	}
	return v
}

// TryFieldU58 tries to add a field and read 58 bit unsigned integer in current endian
func (d *D) TryFieldU58(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU58, sfns...)
}

// FieldU58 adds a field and reads 58 bit unsigned integer in current endian
func (d *D) FieldU58(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U58, sfns...)
}

// Reader U59

// TryU59 tries to read 59 bit unsigned integer in current endian
func (d *D) TryU59() (uint64, error) { return d.tryUE(59, d.Endian) }

func (d *D) ScalarU59() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(59, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U59 reads 59 bit unsigned integer in current endian
func (d *D) U59() uint64 {
	v, err := d.tryUE(59, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U59", Pos: d.Pos()})
	}
	return v
}

// TryFieldU59 tries to add a field and read 59 bit unsigned integer in current endian
func (d *D) TryFieldU59(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU59, sfns...)
}

// FieldU59 adds a field and reads 59 bit unsigned integer in current endian
func (d *D) FieldU59(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U59, sfns...)
}

// Reader U60

// TryU60 tries to read 60 bit unsigned integer in current endian
func (d *D) TryU60() (uint64, error) { return d.tryUE(60, d.Endian) }

func (d *D) ScalarU60() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(60, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U60 reads 60 bit unsigned integer in current endian
func (d *D) U60() uint64 {
	v, err := d.tryUE(60, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U60", Pos: d.Pos()})
	}
	return v
}

// TryFieldU60 tries to add a field and read 60 bit unsigned integer in current endian
func (d *D) TryFieldU60(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU60, sfns...)
}

// FieldU60 adds a field and reads 60 bit unsigned integer in current endian
func (d *D) FieldU60(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U60, sfns...)
}

// Reader U61

// TryU61 tries to read 61 bit unsigned integer in current endian
func (d *D) TryU61() (uint64, error) { return d.tryUE(61, d.Endian) }

func (d *D) ScalarU61() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(61, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U61 reads 61 bit unsigned integer in current endian
func (d *D) U61() uint64 {
	v, err := d.tryUE(61, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U61", Pos: d.Pos()})
	}
	return v
}

// TryFieldU61 tries to add a field and read 61 bit unsigned integer in current endian
func (d *D) TryFieldU61(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU61, sfns...)
}

// FieldU61 adds a field and reads 61 bit unsigned integer in current endian
func (d *D) FieldU61(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U61, sfns...)
}

// Reader U62

// TryU62 tries to read 62 bit unsigned integer in current endian
func (d *D) TryU62() (uint64, error) { return d.tryUE(62, d.Endian) }

func (d *D) ScalarU62() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(62, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U62 reads 62 bit unsigned integer in current endian
func (d *D) U62() uint64 {
	v, err := d.tryUE(62, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U62", Pos: d.Pos()})
	}
	return v
}

// TryFieldU62 tries to add a field and read 62 bit unsigned integer in current endian
func (d *D) TryFieldU62(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU62, sfns...)
}

// FieldU62 adds a field and reads 62 bit unsigned integer in current endian
func (d *D) FieldU62(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U62, sfns...)
}

// Reader U63

// TryU63 tries to read 63 bit unsigned integer in current endian
func (d *D) TryU63() (uint64, error) { return d.tryUE(63, d.Endian) }

func (d *D) ScalarU63() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(63, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U63 reads 63 bit unsigned integer in current endian
func (d *D) U63() uint64 {
	v, err := d.tryUE(63, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U63", Pos: d.Pos()})
	}
	return v
}

// TryFieldU63 tries to add a field and read 63 bit unsigned integer in current endian
func (d *D) TryFieldU63(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU63, sfns...)
}

// FieldU63 adds a field and reads 63 bit unsigned integer in current endian
func (d *D) FieldU63(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U63, sfns...)
}

// Reader U64

// TryU64 tries to read 64 bit unsigned integer in current endian
func (d *D) TryU64() (uint64, error) { return d.tryUE(64, d.Endian) }

func (d *D) ScalarU64() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(64, d.Endian)
		s.Actual = v
		return s, err
	}
}

// U64 reads 64 bit unsigned integer in current endian
func (d *D) U64() uint64 {
	v, err := d.tryUE(64, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "U64", Pos: d.Pos()})
	}
	return v
}

// TryFieldU64 tries to add a field and read 64 bit unsigned integer in current endian
func (d *D) TryFieldU64(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU64, sfns...)
}

// FieldU64 adds a field and reads 64 bit unsigned integer in current endian
func (d *D) FieldU64(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U64, sfns...)
}

// Reader U8LE

// TryU8LE tries to read 8 bit unsigned integer in little-endian
func (d *D) TryU8LE() (uint64, error) { return d.tryUE(8, LittleEndian) }

func (d *D) ScalarU8LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(8, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U8LE reads 8 bit unsigned integer in little-endian
func (d *D) U8LE() uint64 {
	v, err := d.tryUE(8, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U8LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU8LE tries to add a field and read 8 bit unsigned integer in little-endian
func (d *D) TryFieldU8LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU8LE, sfns...)
}

// FieldU8LE adds a field and reads 8 bit unsigned integer in little-endian
func (d *D) FieldU8LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U8LE, sfns...)
}

// Reader U9LE

// TryU9LE tries to read 9 bit unsigned integer in little-endian
func (d *D) TryU9LE() (uint64, error) { return d.tryUE(9, LittleEndian) }

func (d *D) ScalarU9LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(9, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U9LE reads 9 bit unsigned integer in little-endian
func (d *D) U9LE() uint64 {
	v, err := d.tryUE(9, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U9LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU9LE tries to add a field and read 9 bit unsigned integer in little-endian
func (d *D) TryFieldU9LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU9LE, sfns...)
}

// FieldU9LE adds a field and reads 9 bit unsigned integer in little-endian
func (d *D) FieldU9LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U9LE, sfns...)
}

// Reader U10LE

// TryU10LE tries to read 10 bit unsigned integer in little-endian
func (d *D) TryU10LE() (uint64, error) { return d.tryUE(10, LittleEndian) }

func (d *D) ScalarU10LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(10, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U10LE reads 10 bit unsigned integer in little-endian
func (d *D) U10LE() uint64 {
	v, err := d.tryUE(10, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U10LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU10LE tries to add a field and read 10 bit unsigned integer in little-endian
func (d *D) TryFieldU10LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU10LE, sfns...)
}

// FieldU10LE adds a field and reads 10 bit unsigned integer in little-endian
func (d *D) FieldU10LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U10LE, sfns...)
}

// Reader U11LE

// TryU11LE tries to read 11 bit unsigned integer in little-endian
func (d *D) TryU11LE() (uint64, error) { return d.tryUE(11, LittleEndian) }

func (d *D) ScalarU11LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(11, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U11LE reads 11 bit unsigned integer in little-endian
func (d *D) U11LE() uint64 {
	v, err := d.tryUE(11, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U11LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU11LE tries to add a field and read 11 bit unsigned integer in little-endian
func (d *D) TryFieldU11LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU11LE, sfns...)
}

// FieldU11LE adds a field and reads 11 bit unsigned integer in little-endian
func (d *D) FieldU11LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U11LE, sfns...)
}

// Reader U12LE

// TryU12LE tries to read 12 bit unsigned integer in little-endian
func (d *D) TryU12LE() (uint64, error) { return d.tryUE(12, LittleEndian) }

func (d *D) ScalarU12LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(12, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U12LE reads 12 bit unsigned integer in little-endian
func (d *D) U12LE() uint64 {
	v, err := d.tryUE(12, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U12LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU12LE tries to add a field and read 12 bit unsigned integer in little-endian
func (d *D) TryFieldU12LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU12LE, sfns...)
}

// FieldU12LE adds a field and reads 12 bit unsigned integer in little-endian
func (d *D) FieldU12LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U12LE, sfns...)
}

// Reader U13LE

// TryU13LE tries to read 13 bit unsigned integer in little-endian
func (d *D) TryU13LE() (uint64, error) { return d.tryUE(13, LittleEndian) }

func (d *D) ScalarU13LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(13, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U13LE reads 13 bit unsigned integer in little-endian
func (d *D) U13LE() uint64 {
	v, err := d.tryUE(13, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U13LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU13LE tries to add a field and read 13 bit unsigned integer in little-endian
func (d *D) TryFieldU13LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU13LE, sfns...)
}

// FieldU13LE adds a field and reads 13 bit unsigned integer in little-endian
func (d *D) FieldU13LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U13LE, sfns...)
}

// Reader U14LE

// TryU14LE tries to read 14 bit unsigned integer in little-endian
func (d *D) TryU14LE() (uint64, error) { return d.tryUE(14, LittleEndian) }

func (d *D) ScalarU14LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(14, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U14LE reads 14 bit unsigned integer in little-endian
func (d *D) U14LE() uint64 {
	v, err := d.tryUE(14, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U14LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU14LE tries to add a field and read 14 bit unsigned integer in little-endian
func (d *D) TryFieldU14LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU14LE, sfns...)
}

// FieldU14LE adds a field and reads 14 bit unsigned integer in little-endian
func (d *D) FieldU14LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U14LE, sfns...)
}

// Reader U15LE

// TryU15LE tries to read 15 bit unsigned integer in little-endian
func (d *D) TryU15LE() (uint64, error) { return d.tryUE(15, LittleEndian) }

func (d *D) ScalarU15LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(15, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U15LE reads 15 bit unsigned integer in little-endian
func (d *D) U15LE() uint64 {
	v, err := d.tryUE(15, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U15LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU15LE tries to add a field and read 15 bit unsigned integer in little-endian
func (d *D) TryFieldU15LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU15LE, sfns...)
}

// FieldU15LE adds a field and reads 15 bit unsigned integer in little-endian
func (d *D) FieldU15LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U15LE, sfns...)
}

// Reader U16LE

// TryU16LE tries to read 16 bit unsigned integer in little-endian
func (d *D) TryU16LE() (uint64, error) { return d.tryUE(16, LittleEndian) }

func (d *D) ScalarU16LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(16, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U16LE reads 16 bit unsigned integer in little-endian
func (d *D) U16LE() uint64 {
	v, err := d.tryUE(16, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U16LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU16LE tries to add a field and read 16 bit unsigned integer in little-endian
func (d *D) TryFieldU16LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU16LE, sfns...)
}

// FieldU16LE adds a field and reads 16 bit unsigned integer in little-endian
func (d *D) FieldU16LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U16LE, sfns...)
}

// Reader U17LE

// TryU17LE tries to read 17 bit unsigned integer in little-endian
func (d *D) TryU17LE() (uint64, error) { return d.tryUE(17, LittleEndian) }

func (d *D) ScalarU17LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(17, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U17LE reads 17 bit unsigned integer in little-endian
func (d *D) U17LE() uint64 {
	v, err := d.tryUE(17, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U17LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU17LE tries to add a field and read 17 bit unsigned integer in little-endian
func (d *D) TryFieldU17LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU17LE, sfns...)
}

// FieldU17LE adds a field and reads 17 bit unsigned integer in little-endian
func (d *D) FieldU17LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U17LE, sfns...)
}

// Reader U18LE

// TryU18LE tries to read 18 bit unsigned integer in little-endian
func (d *D) TryU18LE() (uint64, error) { return d.tryUE(18, LittleEndian) }

func (d *D) ScalarU18LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(18, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U18LE reads 18 bit unsigned integer in little-endian
func (d *D) U18LE() uint64 {
	v, err := d.tryUE(18, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U18LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU18LE tries to add a field and read 18 bit unsigned integer in little-endian
func (d *D) TryFieldU18LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU18LE, sfns...)
}

// FieldU18LE adds a field and reads 18 bit unsigned integer in little-endian
func (d *D) FieldU18LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U18LE, sfns...)
}

// Reader U19LE

// TryU19LE tries to read 19 bit unsigned integer in little-endian
func (d *D) TryU19LE() (uint64, error) { return d.tryUE(19, LittleEndian) }

func (d *D) ScalarU19LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(19, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U19LE reads 19 bit unsigned integer in little-endian
func (d *D) U19LE() uint64 {
	v, err := d.tryUE(19, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U19LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU19LE tries to add a field and read 19 bit unsigned integer in little-endian
func (d *D) TryFieldU19LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU19LE, sfns...)
}

// FieldU19LE adds a field and reads 19 bit unsigned integer in little-endian
func (d *D) FieldU19LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U19LE, sfns...)
}

// Reader U20LE

// TryU20LE tries to read 20 bit unsigned integer in little-endian
func (d *D) TryU20LE() (uint64, error) { return d.tryUE(20, LittleEndian) }

func (d *D) ScalarU20LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(20, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U20LE reads 20 bit unsigned integer in little-endian
func (d *D) U20LE() uint64 {
	v, err := d.tryUE(20, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U20LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU20LE tries to add a field and read 20 bit unsigned integer in little-endian
func (d *D) TryFieldU20LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU20LE, sfns...)
}

// FieldU20LE adds a field and reads 20 bit unsigned integer in little-endian
func (d *D) FieldU20LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U20LE, sfns...)
}

// Reader U21LE

// TryU21LE tries to read 21 bit unsigned integer in little-endian
func (d *D) TryU21LE() (uint64, error) { return d.tryUE(21, LittleEndian) }

func (d *D) ScalarU21LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(21, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U21LE reads 21 bit unsigned integer in little-endian
func (d *D) U21LE() uint64 {
	v, err := d.tryUE(21, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U21LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU21LE tries to add a field and read 21 bit unsigned integer in little-endian
func (d *D) TryFieldU21LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU21LE, sfns...)
}

// FieldU21LE adds a field and reads 21 bit unsigned integer in little-endian
func (d *D) FieldU21LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U21LE, sfns...)
}

// Reader U22LE

// TryU22LE tries to read 22 bit unsigned integer in little-endian
func (d *D) TryU22LE() (uint64, error) { return d.tryUE(22, LittleEndian) }

func (d *D) ScalarU22LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(22, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U22LE reads 22 bit unsigned integer in little-endian
func (d *D) U22LE() uint64 {
	v, err := d.tryUE(22, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U22LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU22LE tries to add a field and read 22 bit unsigned integer in little-endian
func (d *D) TryFieldU22LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU22LE, sfns...)
}

// FieldU22LE adds a field and reads 22 bit unsigned integer in little-endian
func (d *D) FieldU22LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U22LE, sfns...)
}

// Reader U23LE

// TryU23LE tries to read 23 bit unsigned integer in little-endian
func (d *D) TryU23LE() (uint64, error) { return d.tryUE(23, LittleEndian) }

func (d *D) ScalarU23LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(23, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U23LE reads 23 bit unsigned integer in little-endian
func (d *D) U23LE() uint64 {
	v, err := d.tryUE(23, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U23LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU23LE tries to add a field and read 23 bit unsigned integer in little-endian
func (d *D) TryFieldU23LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU23LE, sfns...)
}

// FieldU23LE adds a field and reads 23 bit unsigned integer in little-endian
func (d *D) FieldU23LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U23LE, sfns...)
}

// Reader U24LE

// TryU24LE tries to read 24 bit unsigned integer in little-endian
func (d *D) TryU24LE() (uint64, error) { return d.tryUE(24, LittleEndian) }

func (d *D) ScalarU24LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(24, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U24LE reads 24 bit unsigned integer in little-endian
func (d *D) U24LE() uint64 {
	v, err := d.tryUE(24, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U24LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU24LE tries to add a field and read 24 bit unsigned integer in little-endian
func (d *D) TryFieldU24LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU24LE, sfns...)
}

// FieldU24LE adds a field and reads 24 bit unsigned integer in little-endian
func (d *D) FieldU24LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U24LE, sfns...)
}

// Reader U25LE

// TryU25LE tries to read 25 bit unsigned integer in little-endian
func (d *D) TryU25LE() (uint64, error) { return d.tryUE(25, LittleEndian) }

func (d *D) ScalarU25LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(25, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U25LE reads 25 bit unsigned integer in little-endian
func (d *D) U25LE() uint64 {
	v, err := d.tryUE(25, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U25LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU25LE tries to add a field and read 25 bit unsigned integer in little-endian
func (d *D) TryFieldU25LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU25LE, sfns...)
}

// FieldU25LE adds a field and reads 25 bit unsigned integer in little-endian
func (d *D) FieldU25LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U25LE, sfns...)
}

// Reader U26LE

// TryU26LE tries to read 26 bit unsigned integer in little-endian
func (d *D) TryU26LE() (uint64, error) { return d.tryUE(26, LittleEndian) }

func (d *D) ScalarU26LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(26, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U26LE reads 26 bit unsigned integer in little-endian
func (d *D) U26LE() uint64 {
	v, err := d.tryUE(26, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U26LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU26LE tries to add a field and read 26 bit unsigned integer in little-endian
func (d *D) TryFieldU26LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU26LE, sfns...)
}

// FieldU26LE adds a field and reads 26 bit unsigned integer in little-endian
func (d *D) FieldU26LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U26LE, sfns...)
}

// Reader U27LE

// TryU27LE tries to read 27 bit unsigned integer in little-endian
func (d *D) TryU27LE() (uint64, error) { return d.tryUE(27, LittleEndian) }

func (d *D) ScalarU27LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(27, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U27LE reads 27 bit unsigned integer in little-endian
func (d *D) U27LE() uint64 {
	v, err := d.tryUE(27, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U27LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU27LE tries to add a field and read 27 bit unsigned integer in little-endian
func (d *D) TryFieldU27LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU27LE, sfns...)
}

// FieldU27LE adds a field and reads 27 bit unsigned integer in little-endian
func (d *D) FieldU27LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U27LE, sfns...)
}

// Reader U28LE

// TryU28LE tries to read 28 bit unsigned integer in little-endian
func (d *D) TryU28LE() (uint64, error) { return d.tryUE(28, LittleEndian) }

func (d *D) ScalarU28LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(28, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U28LE reads 28 bit unsigned integer in little-endian
func (d *D) U28LE() uint64 {
	v, err := d.tryUE(28, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U28LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU28LE tries to add a field and read 28 bit unsigned integer in little-endian
func (d *D) TryFieldU28LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU28LE, sfns...)
}

// FieldU28LE adds a field and reads 28 bit unsigned integer in little-endian
func (d *D) FieldU28LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U28LE, sfns...)
}

// Reader U29LE

// TryU29LE tries to read 29 bit unsigned integer in little-endian
func (d *D) TryU29LE() (uint64, error) { return d.tryUE(29, LittleEndian) }

func (d *D) ScalarU29LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(29, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U29LE reads 29 bit unsigned integer in little-endian
func (d *D) U29LE() uint64 {
	v, err := d.tryUE(29, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U29LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU29LE tries to add a field and read 29 bit unsigned integer in little-endian
func (d *D) TryFieldU29LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU29LE, sfns...)
}

// FieldU29LE adds a field and reads 29 bit unsigned integer in little-endian
func (d *D) FieldU29LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U29LE, sfns...)
}

// Reader U30LE

// TryU30LE tries to read 30 bit unsigned integer in little-endian
func (d *D) TryU30LE() (uint64, error) { return d.tryUE(30, LittleEndian) }

func (d *D) ScalarU30LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(30, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U30LE reads 30 bit unsigned integer in little-endian
func (d *D) U30LE() uint64 {
	v, err := d.tryUE(30, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U30LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU30LE tries to add a field and read 30 bit unsigned integer in little-endian
func (d *D) TryFieldU30LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU30LE, sfns...)
}

// FieldU30LE adds a field and reads 30 bit unsigned integer in little-endian
func (d *D) FieldU30LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U30LE, sfns...)
}

// Reader U31LE

// TryU31LE tries to read 31 bit unsigned integer in little-endian
func (d *D) TryU31LE() (uint64, error) { return d.tryUE(31, LittleEndian) }

func (d *D) ScalarU31LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(31, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U31LE reads 31 bit unsigned integer in little-endian
func (d *D) U31LE() uint64 {
	v, err := d.tryUE(31, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U31LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU31LE tries to add a field and read 31 bit unsigned integer in little-endian
func (d *D) TryFieldU31LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU31LE, sfns...)
}

// FieldU31LE adds a field and reads 31 bit unsigned integer in little-endian
func (d *D) FieldU31LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U31LE, sfns...)
}

// Reader U32LE

// TryU32LE tries to read 32 bit unsigned integer in little-endian
func (d *D) TryU32LE() (uint64, error) { return d.tryUE(32, LittleEndian) }

func (d *D) ScalarU32LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(32, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U32LE reads 32 bit unsigned integer in little-endian
func (d *D) U32LE() uint64 {
	v, err := d.tryUE(32, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U32LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU32LE tries to add a field and read 32 bit unsigned integer in little-endian
func (d *D) TryFieldU32LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU32LE, sfns...)
}

// FieldU32LE adds a field and reads 32 bit unsigned integer in little-endian
func (d *D) FieldU32LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U32LE, sfns...)
}

// Reader U33LE

// TryU33LE tries to read 33 bit unsigned integer in little-endian
func (d *D) TryU33LE() (uint64, error) { return d.tryUE(33, LittleEndian) }

func (d *D) ScalarU33LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(33, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U33LE reads 33 bit unsigned integer in little-endian
func (d *D) U33LE() uint64 {
	v, err := d.tryUE(33, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U33LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU33LE tries to add a field and read 33 bit unsigned integer in little-endian
func (d *D) TryFieldU33LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU33LE, sfns...)
}

// FieldU33LE adds a field and reads 33 bit unsigned integer in little-endian
func (d *D) FieldU33LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U33LE, sfns...)
}

// Reader U34LE

// TryU34LE tries to read 34 bit unsigned integer in little-endian
func (d *D) TryU34LE() (uint64, error) { return d.tryUE(34, LittleEndian) }

func (d *D) ScalarU34LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(34, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U34LE reads 34 bit unsigned integer in little-endian
func (d *D) U34LE() uint64 {
	v, err := d.tryUE(34, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U34LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU34LE tries to add a field and read 34 bit unsigned integer in little-endian
func (d *D) TryFieldU34LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU34LE, sfns...)
}

// FieldU34LE adds a field and reads 34 bit unsigned integer in little-endian
func (d *D) FieldU34LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U34LE, sfns...)
}

// Reader U35LE

// TryU35LE tries to read 35 bit unsigned integer in little-endian
func (d *D) TryU35LE() (uint64, error) { return d.tryUE(35, LittleEndian) }

func (d *D) ScalarU35LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(35, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U35LE reads 35 bit unsigned integer in little-endian
func (d *D) U35LE() uint64 {
	v, err := d.tryUE(35, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U35LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU35LE tries to add a field and read 35 bit unsigned integer in little-endian
func (d *D) TryFieldU35LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU35LE, sfns...)
}

// FieldU35LE adds a field and reads 35 bit unsigned integer in little-endian
func (d *D) FieldU35LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U35LE, sfns...)
}

// Reader U36LE

// TryU36LE tries to read 36 bit unsigned integer in little-endian
func (d *D) TryU36LE() (uint64, error) { return d.tryUE(36, LittleEndian) }

func (d *D) ScalarU36LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(36, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U36LE reads 36 bit unsigned integer in little-endian
func (d *D) U36LE() uint64 {
	v, err := d.tryUE(36, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U36LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU36LE tries to add a field and read 36 bit unsigned integer in little-endian
func (d *D) TryFieldU36LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU36LE, sfns...)
}

// FieldU36LE adds a field and reads 36 bit unsigned integer in little-endian
func (d *D) FieldU36LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U36LE, sfns...)
}

// Reader U37LE

// TryU37LE tries to read 37 bit unsigned integer in little-endian
func (d *D) TryU37LE() (uint64, error) { return d.tryUE(37, LittleEndian) }

func (d *D) ScalarU37LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(37, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U37LE reads 37 bit unsigned integer in little-endian
func (d *D) U37LE() uint64 {
	v, err := d.tryUE(37, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U37LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU37LE tries to add a field and read 37 bit unsigned integer in little-endian
func (d *D) TryFieldU37LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU37LE, sfns...)
}

// FieldU37LE adds a field and reads 37 bit unsigned integer in little-endian
func (d *D) FieldU37LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U37LE, sfns...)
}

// Reader U38LE

// TryU38LE tries to read 38 bit unsigned integer in little-endian
func (d *D) TryU38LE() (uint64, error) { return d.tryUE(38, LittleEndian) }

func (d *D) ScalarU38LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(38, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U38LE reads 38 bit unsigned integer in little-endian
func (d *D) U38LE() uint64 {
	v, err := d.tryUE(38, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U38LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU38LE tries to add a field and read 38 bit unsigned integer in little-endian
func (d *D) TryFieldU38LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU38LE, sfns...)
}

// FieldU38LE adds a field and reads 38 bit unsigned integer in little-endian
func (d *D) FieldU38LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U38LE, sfns...)
}

// Reader U39LE

// TryU39LE tries to read 39 bit unsigned integer in little-endian
func (d *D) TryU39LE() (uint64, error) { return d.tryUE(39, LittleEndian) }

func (d *D) ScalarU39LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(39, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U39LE reads 39 bit unsigned integer in little-endian
func (d *D) U39LE() uint64 {
	v, err := d.tryUE(39, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U39LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU39LE tries to add a field and read 39 bit unsigned integer in little-endian
func (d *D) TryFieldU39LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU39LE, sfns...)
}

// FieldU39LE adds a field and reads 39 bit unsigned integer in little-endian
func (d *D) FieldU39LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U39LE, sfns...)
}

// Reader U40LE

// TryU40LE tries to read 40 bit unsigned integer in little-endian
func (d *D) TryU40LE() (uint64, error) { return d.tryUE(40, LittleEndian) }

func (d *D) ScalarU40LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(40, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U40LE reads 40 bit unsigned integer in little-endian
func (d *D) U40LE() uint64 {
	v, err := d.tryUE(40, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U40LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU40LE tries to add a field and read 40 bit unsigned integer in little-endian
func (d *D) TryFieldU40LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU40LE, sfns...)
}

// FieldU40LE adds a field and reads 40 bit unsigned integer in little-endian
func (d *D) FieldU40LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U40LE, sfns...)
}

// Reader U41LE

// TryU41LE tries to read 41 bit unsigned integer in little-endian
func (d *D) TryU41LE() (uint64, error) { return d.tryUE(41, LittleEndian) }

func (d *D) ScalarU41LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(41, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U41LE reads 41 bit unsigned integer in little-endian
func (d *D) U41LE() uint64 {
	v, err := d.tryUE(41, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U41LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU41LE tries to add a field and read 41 bit unsigned integer in little-endian
func (d *D) TryFieldU41LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU41LE, sfns...)
}

// FieldU41LE adds a field and reads 41 bit unsigned integer in little-endian
func (d *D) FieldU41LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U41LE, sfns...)
}

// Reader U42LE

// TryU42LE tries to read 42 bit unsigned integer in little-endian
func (d *D) TryU42LE() (uint64, error) { return d.tryUE(42, LittleEndian) }

func (d *D) ScalarU42LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(42, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U42LE reads 42 bit unsigned integer in little-endian
func (d *D) U42LE() uint64 {
	v, err := d.tryUE(42, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U42LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU42LE tries to add a field and read 42 bit unsigned integer in little-endian
func (d *D) TryFieldU42LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU42LE, sfns...)
}

// FieldU42LE adds a field and reads 42 bit unsigned integer in little-endian
func (d *D) FieldU42LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U42LE, sfns...)
}

// Reader U43LE

// TryU43LE tries to read 43 bit unsigned integer in little-endian
func (d *D) TryU43LE() (uint64, error) { return d.tryUE(43, LittleEndian) }

func (d *D) ScalarU43LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(43, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U43LE reads 43 bit unsigned integer in little-endian
func (d *D) U43LE() uint64 {
	v, err := d.tryUE(43, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U43LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU43LE tries to add a field and read 43 bit unsigned integer in little-endian
func (d *D) TryFieldU43LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU43LE, sfns...)
}

// FieldU43LE adds a field and reads 43 bit unsigned integer in little-endian
func (d *D) FieldU43LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U43LE, sfns...)
}

// Reader U44LE

// TryU44LE tries to read 44 bit unsigned integer in little-endian
func (d *D) TryU44LE() (uint64, error) { return d.tryUE(44, LittleEndian) }

func (d *D) ScalarU44LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(44, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U44LE reads 44 bit unsigned integer in little-endian
func (d *D) U44LE() uint64 {
	v, err := d.tryUE(44, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U44LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU44LE tries to add a field and read 44 bit unsigned integer in little-endian
func (d *D) TryFieldU44LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU44LE, sfns...)
}

// FieldU44LE adds a field and reads 44 bit unsigned integer in little-endian
func (d *D) FieldU44LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U44LE, sfns...)
}

// Reader U45LE

// TryU45LE tries to read 45 bit unsigned integer in little-endian
func (d *D) TryU45LE() (uint64, error) { return d.tryUE(45, LittleEndian) }

func (d *D) ScalarU45LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(45, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U45LE reads 45 bit unsigned integer in little-endian
func (d *D) U45LE() uint64 {
	v, err := d.tryUE(45, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U45LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU45LE tries to add a field and read 45 bit unsigned integer in little-endian
func (d *D) TryFieldU45LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU45LE, sfns...)
}

// FieldU45LE adds a field and reads 45 bit unsigned integer in little-endian
func (d *D) FieldU45LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U45LE, sfns...)
}

// Reader U46LE

// TryU46LE tries to read 46 bit unsigned integer in little-endian
func (d *D) TryU46LE() (uint64, error) { return d.tryUE(46, LittleEndian) }

func (d *D) ScalarU46LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(46, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U46LE reads 46 bit unsigned integer in little-endian
func (d *D) U46LE() uint64 {
	v, err := d.tryUE(46, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U46LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU46LE tries to add a field and read 46 bit unsigned integer in little-endian
func (d *D) TryFieldU46LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU46LE, sfns...)
}

// FieldU46LE adds a field and reads 46 bit unsigned integer in little-endian
func (d *D) FieldU46LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U46LE, sfns...)
}

// Reader U47LE

// TryU47LE tries to read 47 bit unsigned integer in little-endian
func (d *D) TryU47LE() (uint64, error) { return d.tryUE(47, LittleEndian) }

func (d *D) ScalarU47LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(47, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U47LE reads 47 bit unsigned integer in little-endian
func (d *D) U47LE() uint64 {
	v, err := d.tryUE(47, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U47LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU47LE tries to add a field and read 47 bit unsigned integer in little-endian
func (d *D) TryFieldU47LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU47LE, sfns...)
}

// FieldU47LE adds a field and reads 47 bit unsigned integer in little-endian
func (d *D) FieldU47LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U47LE, sfns...)
}

// Reader U48LE

// TryU48LE tries to read 48 bit unsigned integer in little-endian
func (d *D) TryU48LE() (uint64, error) { return d.tryUE(48, LittleEndian) }

func (d *D) ScalarU48LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(48, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U48LE reads 48 bit unsigned integer in little-endian
func (d *D) U48LE() uint64 {
	v, err := d.tryUE(48, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U48LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU48LE tries to add a field and read 48 bit unsigned integer in little-endian
func (d *D) TryFieldU48LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU48LE, sfns...)
}

// FieldU48LE adds a field and reads 48 bit unsigned integer in little-endian
func (d *D) FieldU48LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U48LE, sfns...)
}

// Reader U49LE

// TryU49LE tries to read 49 bit unsigned integer in little-endian
func (d *D) TryU49LE() (uint64, error) { return d.tryUE(49, LittleEndian) }

func (d *D) ScalarU49LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(49, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U49LE reads 49 bit unsigned integer in little-endian
func (d *D) U49LE() uint64 {
	v, err := d.tryUE(49, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U49LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU49LE tries to add a field and read 49 bit unsigned integer in little-endian
func (d *D) TryFieldU49LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU49LE, sfns...)
}

// FieldU49LE adds a field and reads 49 bit unsigned integer in little-endian
func (d *D) FieldU49LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U49LE, sfns...)
}

// Reader U50LE

// TryU50LE tries to read 50 bit unsigned integer in little-endian
func (d *D) TryU50LE() (uint64, error) { return d.tryUE(50, LittleEndian) }

func (d *D) ScalarU50LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(50, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U50LE reads 50 bit unsigned integer in little-endian
func (d *D) U50LE() uint64 {
	v, err := d.tryUE(50, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U50LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU50LE tries to add a field and read 50 bit unsigned integer in little-endian
func (d *D) TryFieldU50LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU50LE, sfns...)
}

// FieldU50LE adds a field and reads 50 bit unsigned integer in little-endian
func (d *D) FieldU50LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U50LE, sfns...)
}

// Reader U51LE

// TryU51LE tries to read 51 bit unsigned integer in little-endian
func (d *D) TryU51LE() (uint64, error) { return d.tryUE(51, LittleEndian) }

func (d *D) ScalarU51LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(51, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U51LE reads 51 bit unsigned integer in little-endian
func (d *D) U51LE() uint64 {
	v, err := d.tryUE(51, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U51LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU51LE tries to add a field and read 51 bit unsigned integer in little-endian
func (d *D) TryFieldU51LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU51LE, sfns...)
}

// FieldU51LE adds a field and reads 51 bit unsigned integer in little-endian
func (d *D) FieldU51LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U51LE, sfns...)
}

// Reader U52LE

// TryU52LE tries to read 52 bit unsigned integer in little-endian
func (d *D) TryU52LE() (uint64, error) { return d.tryUE(52, LittleEndian) }

func (d *D) ScalarU52LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(52, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U52LE reads 52 bit unsigned integer in little-endian
func (d *D) U52LE() uint64 {
	v, err := d.tryUE(52, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U52LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU52LE tries to add a field and read 52 bit unsigned integer in little-endian
func (d *D) TryFieldU52LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU52LE, sfns...)
}

// FieldU52LE adds a field and reads 52 bit unsigned integer in little-endian
func (d *D) FieldU52LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U52LE, sfns...)
}

// Reader U53LE

// TryU53LE tries to read 53 bit unsigned integer in little-endian
func (d *D) TryU53LE() (uint64, error) { return d.tryUE(53, LittleEndian) }

func (d *D) ScalarU53LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(53, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U53LE reads 53 bit unsigned integer in little-endian
func (d *D) U53LE() uint64 {
	v, err := d.tryUE(53, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U53LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU53LE tries to add a field and read 53 bit unsigned integer in little-endian
func (d *D) TryFieldU53LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU53LE, sfns...)
}

// FieldU53LE adds a field and reads 53 bit unsigned integer in little-endian
func (d *D) FieldU53LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U53LE, sfns...)
}

// Reader U54LE

// TryU54LE tries to read 54 bit unsigned integer in little-endian
func (d *D) TryU54LE() (uint64, error) { return d.tryUE(54, LittleEndian) }

func (d *D) ScalarU54LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(54, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U54LE reads 54 bit unsigned integer in little-endian
func (d *D) U54LE() uint64 {
	v, err := d.tryUE(54, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U54LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU54LE tries to add a field and read 54 bit unsigned integer in little-endian
func (d *D) TryFieldU54LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU54LE, sfns...)
}

// FieldU54LE adds a field and reads 54 bit unsigned integer in little-endian
func (d *D) FieldU54LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U54LE, sfns...)
}

// Reader U55LE

// TryU55LE tries to read 55 bit unsigned integer in little-endian
func (d *D) TryU55LE() (uint64, error) { return d.tryUE(55, LittleEndian) }

func (d *D) ScalarU55LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(55, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U55LE reads 55 bit unsigned integer in little-endian
func (d *D) U55LE() uint64 {
	v, err := d.tryUE(55, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U55LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU55LE tries to add a field and read 55 bit unsigned integer in little-endian
func (d *D) TryFieldU55LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU55LE, sfns...)
}

// FieldU55LE adds a field and reads 55 bit unsigned integer in little-endian
func (d *D) FieldU55LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U55LE, sfns...)
}

// Reader U56LE

// TryU56LE tries to read 56 bit unsigned integer in little-endian
func (d *D) TryU56LE() (uint64, error) { return d.tryUE(56, LittleEndian) }

func (d *D) ScalarU56LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(56, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U56LE reads 56 bit unsigned integer in little-endian
func (d *D) U56LE() uint64 {
	v, err := d.tryUE(56, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U56LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU56LE tries to add a field and read 56 bit unsigned integer in little-endian
func (d *D) TryFieldU56LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU56LE, sfns...)
}

// FieldU56LE adds a field and reads 56 bit unsigned integer in little-endian
func (d *D) FieldU56LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U56LE, sfns...)
}

// Reader U57LE

// TryU57LE tries to read 57 bit unsigned integer in little-endian
func (d *D) TryU57LE() (uint64, error) { return d.tryUE(57, LittleEndian) }

func (d *D) ScalarU57LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(57, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U57LE reads 57 bit unsigned integer in little-endian
func (d *D) U57LE() uint64 {
	v, err := d.tryUE(57, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U57LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU57LE tries to add a field and read 57 bit unsigned integer in little-endian
func (d *D) TryFieldU57LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU57LE, sfns...)
}

// FieldU57LE adds a field and reads 57 bit unsigned integer in little-endian
func (d *D) FieldU57LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U57LE, sfns...)
}

// Reader U58LE

// TryU58LE tries to read 58 bit unsigned integer in little-endian
func (d *D) TryU58LE() (uint64, error) { return d.tryUE(58, LittleEndian) }

func (d *D) ScalarU58LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(58, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U58LE reads 58 bit unsigned integer in little-endian
func (d *D) U58LE() uint64 {
	v, err := d.tryUE(58, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U58LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU58LE tries to add a field and read 58 bit unsigned integer in little-endian
func (d *D) TryFieldU58LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU58LE, sfns...)
}

// FieldU58LE adds a field and reads 58 bit unsigned integer in little-endian
func (d *D) FieldU58LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U58LE, sfns...)
}

// Reader U59LE

// TryU59LE tries to read 59 bit unsigned integer in little-endian
func (d *D) TryU59LE() (uint64, error) { return d.tryUE(59, LittleEndian) }

func (d *D) ScalarU59LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(59, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U59LE reads 59 bit unsigned integer in little-endian
func (d *D) U59LE() uint64 {
	v, err := d.tryUE(59, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U59LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU59LE tries to add a field and read 59 bit unsigned integer in little-endian
func (d *D) TryFieldU59LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU59LE, sfns...)
}

// FieldU59LE adds a field and reads 59 bit unsigned integer in little-endian
func (d *D) FieldU59LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U59LE, sfns...)
}

// Reader U60LE

// TryU60LE tries to read 60 bit unsigned integer in little-endian
func (d *D) TryU60LE() (uint64, error) { return d.tryUE(60, LittleEndian) }

func (d *D) ScalarU60LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(60, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U60LE reads 60 bit unsigned integer in little-endian
func (d *D) U60LE() uint64 {
	v, err := d.tryUE(60, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U60LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU60LE tries to add a field and read 60 bit unsigned integer in little-endian
func (d *D) TryFieldU60LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU60LE, sfns...)
}

// FieldU60LE adds a field and reads 60 bit unsigned integer in little-endian
func (d *D) FieldU60LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U60LE, sfns...)
}

// Reader U61LE

// TryU61LE tries to read 61 bit unsigned integer in little-endian
func (d *D) TryU61LE() (uint64, error) { return d.tryUE(61, LittleEndian) }

func (d *D) ScalarU61LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(61, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U61LE reads 61 bit unsigned integer in little-endian
func (d *D) U61LE() uint64 {
	v, err := d.tryUE(61, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U61LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU61LE tries to add a field and read 61 bit unsigned integer in little-endian
func (d *D) TryFieldU61LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU61LE, sfns...)
}

// FieldU61LE adds a field and reads 61 bit unsigned integer in little-endian
func (d *D) FieldU61LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U61LE, sfns...)
}

// Reader U62LE

// TryU62LE tries to read 62 bit unsigned integer in little-endian
func (d *D) TryU62LE() (uint64, error) { return d.tryUE(62, LittleEndian) }

func (d *D) ScalarU62LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(62, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U62LE reads 62 bit unsigned integer in little-endian
func (d *D) U62LE() uint64 {
	v, err := d.tryUE(62, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U62LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU62LE tries to add a field and read 62 bit unsigned integer in little-endian
func (d *D) TryFieldU62LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU62LE, sfns...)
}

// FieldU62LE adds a field and reads 62 bit unsigned integer in little-endian
func (d *D) FieldU62LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U62LE, sfns...)
}

// Reader U63LE

// TryU63LE tries to read 63 bit unsigned integer in little-endian
func (d *D) TryU63LE() (uint64, error) { return d.tryUE(63, LittleEndian) }

func (d *D) ScalarU63LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(63, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U63LE reads 63 bit unsigned integer in little-endian
func (d *D) U63LE() uint64 {
	v, err := d.tryUE(63, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U63LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU63LE tries to add a field and read 63 bit unsigned integer in little-endian
func (d *D) TryFieldU63LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU63LE, sfns...)
}

// FieldU63LE adds a field and reads 63 bit unsigned integer in little-endian
func (d *D) FieldU63LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U63LE, sfns...)
}

// Reader U64LE

// TryU64LE tries to read 64 bit unsigned integer in little-endian
func (d *D) TryU64LE() (uint64, error) { return d.tryUE(64, LittleEndian) }

func (d *D) ScalarU64LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(64, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// U64LE reads 64 bit unsigned integer in little-endian
func (d *D) U64LE() uint64 {
	v, err := d.tryUE(64, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U64LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU64LE tries to add a field and read 64 bit unsigned integer in little-endian
func (d *D) TryFieldU64LE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU64LE, sfns...)
}

// FieldU64LE adds a field and reads 64 bit unsigned integer in little-endian
func (d *D) FieldU64LE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U64LE, sfns...)
}

// Reader U8BE

// TryU8BE tries to read 8 bit unsigned integer in big-endian
func (d *D) TryU8BE() (uint64, error) { return d.tryUE(8, BigEndian) }

func (d *D) ScalarU8BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(8, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U8BE reads 8 bit unsigned integer in big-endian
func (d *D) U8BE() uint64 {
	v, err := d.tryUE(8, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U8BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU8BE tries to add a field and read 8 bit unsigned integer in big-endian
func (d *D) TryFieldU8BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU8BE, sfns...)
}

// FieldU8BE adds a field and reads 8 bit unsigned integer in big-endian
func (d *D) FieldU8BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U8BE, sfns...)
}

// Reader U9BE

// TryU9BE tries to read 9 bit unsigned integer in big-endian
func (d *D) TryU9BE() (uint64, error) { return d.tryUE(9, BigEndian) }

func (d *D) ScalarU9BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(9, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U9BE reads 9 bit unsigned integer in big-endian
func (d *D) U9BE() uint64 {
	v, err := d.tryUE(9, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U9BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU9BE tries to add a field and read 9 bit unsigned integer in big-endian
func (d *D) TryFieldU9BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU9BE, sfns...)
}

// FieldU9BE adds a field and reads 9 bit unsigned integer in big-endian
func (d *D) FieldU9BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U9BE, sfns...)
}

// Reader U10BE

// TryU10BE tries to read 10 bit unsigned integer in big-endian
func (d *D) TryU10BE() (uint64, error) { return d.tryUE(10, BigEndian) }

func (d *D) ScalarU10BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(10, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U10BE reads 10 bit unsigned integer in big-endian
func (d *D) U10BE() uint64 {
	v, err := d.tryUE(10, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U10BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU10BE tries to add a field and read 10 bit unsigned integer in big-endian
func (d *D) TryFieldU10BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU10BE, sfns...)
}

// FieldU10BE adds a field and reads 10 bit unsigned integer in big-endian
func (d *D) FieldU10BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U10BE, sfns...)
}

// Reader U11BE

// TryU11BE tries to read 11 bit unsigned integer in big-endian
func (d *D) TryU11BE() (uint64, error) { return d.tryUE(11, BigEndian) }

func (d *D) ScalarU11BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(11, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U11BE reads 11 bit unsigned integer in big-endian
func (d *D) U11BE() uint64 {
	v, err := d.tryUE(11, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U11BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU11BE tries to add a field and read 11 bit unsigned integer in big-endian
func (d *D) TryFieldU11BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU11BE, sfns...)
}

// FieldU11BE adds a field and reads 11 bit unsigned integer in big-endian
func (d *D) FieldU11BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U11BE, sfns...)
}

// Reader U12BE

// TryU12BE tries to read 12 bit unsigned integer in big-endian
func (d *D) TryU12BE() (uint64, error) { return d.tryUE(12, BigEndian) }

func (d *D) ScalarU12BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(12, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U12BE reads 12 bit unsigned integer in big-endian
func (d *D) U12BE() uint64 {
	v, err := d.tryUE(12, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U12BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU12BE tries to add a field and read 12 bit unsigned integer in big-endian
func (d *D) TryFieldU12BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU12BE, sfns...)
}

// FieldU12BE adds a field and reads 12 bit unsigned integer in big-endian
func (d *D) FieldU12BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U12BE, sfns...)
}

// Reader U13BE

// TryU13BE tries to read 13 bit unsigned integer in big-endian
func (d *D) TryU13BE() (uint64, error) { return d.tryUE(13, BigEndian) }

func (d *D) ScalarU13BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(13, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U13BE reads 13 bit unsigned integer in big-endian
func (d *D) U13BE() uint64 {
	v, err := d.tryUE(13, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U13BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU13BE tries to add a field and read 13 bit unsigned integer in big-endian
func (d *D) TryFieldU13BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU13BE, sfns...)
}

// FieldU13BE adds a field and reads 13 bit unsigned integer in big-endian
func (d *D) FieldU13BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U13BE, sfns...)
}

// Reader U14BE

// TryU14BE tries to read 14 bit unsigned integer in big-endian
func (d *D) TryU14BE() (uint64, error) { return d.tryUE(14, BigEndian) }

func (d *D) ScalarU14BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(14, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U14BE reads 14 bit unsigned integer in big-endian
func (d *D) U14BE() uint64 {
	v, err := d.tryUE(14, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U14BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU14BE tries to add a field and read 14 bit unsigned integer in big-endian
func (d *D) TryFieldU14BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU14BE, sfns...)
}

// FieldU14BE adds a field and reads 14 bit unsigned integer in big-endian
func (d *D) FieldU14BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U14BE, sfns...)
}

// Reader U15BE

// TryU15BE tries to read 15 bit unsigned integer in big-endian
func (d *D) TryU15BE() (uint64, error) { return d.tryUE(15, BigEndian) }

func (d *D) ScalarU15BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(15, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U15BE reads 15 bit unsigned integer in big-endian
func (d *D) U15BE() uint64 {
	v, err := d.tryUE(15, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U15BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU15BE tries to add a field and read 15 bit unsigned integer in big-endian
func (d *D) TryFieldU15BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU15BE, sfns...)
}

// FieldU15BE adds a field and reads 15 bit unsigned integer in big-endian
func (d *D) FieldU15BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U15BE, sfns...)
}

// Reader U16BE

// TryU16BE tries to read 16 bit unsigned integer in big-endian
func (d *D) TryU16BE() (uint64, error) { return d.tryUE(16, BigEndian) }

func (d *D) ScalarU16BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(16, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U16BE reads 16 bit unsigned integer in big-endian
func (d *D) U16BE() uint64 {
	v, err := d.tryUE(16, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U16BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU16BE tries to add a field and read 16 bit unsigned integer in big-endian
func (d *D) TryFieldU16BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU16BE, sfns...)
}

// FieldU16BE adds a field and reads 16 bit unsigned integer in big-endian
func (d *D) FieldU16BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U16BE, sfns...)
}

// Reader U17BE

// TryU17BE tries to read 17 bit unsigned integer in big-endian
func (d *D) TryU17BE() (uint64, error) { return d.tryUE(17, BigEndian) }

func (d *D) ScalarU17BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(17, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U17BE reads 17 bit unsigned integer in big-endian
func (d *D) U17BE() uint64 {
	v, err := d.tryUE(17, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U17BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU17BE tries to add a field and read 17 bit unsigned integer in big-endian
func (d *D) TryFieldU17BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU17BE, sfns...)
}

// FieldU17BE adds a field and reads 17 bit unsigned integer in big-endian
func (d *D) FieldU17BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U17BE, sfns...)
}

// Reader U18BE

// TryU18BE tries to read 18 bit unsigned integer in big-endian
func (d *D) TryU18BE() (uint64, error) { return d.tryUE(18, BigEndian) }

func (d *D) ScalarU18BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(18, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U18BE reads 18 bit unsigned integer in big-endian
func (d *D) U18BE() uint64 {
	v, err := d.tryUE(18, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U18BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU18BE tries to add a field and read 18 bit unsigned integer in big-endian
func (d *D) TryFieldU18BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU18BE, sfns...)
}

// FieldU18BE adds a field and reads 18 bit unsigned integer in big-endian
func (d *D) FieldU18BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U18BE, sfns...)
}

// Reader U19BE

// TryU19BE tries to read 19 bit unsigned integer in big-endian
func (d *D) TryU19BE() (uint64, error) { return d.tryUE(19, BigEndian) }

func (d *D) ScalarU19BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(19, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U19BE reads 19 bit unsigned integer in big-endian
func (d *D) U19BE() uint64 {
	v, err := d.tryUE(19, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U19BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU19BE tries to add a field and read 19 bit unsigned integer in big-endian
func (d *D) TryFieldU19BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU19BE, sfns...)
}

// FieldU19BE adds a field and reads 19 bit unsigned integer in big-endian
func (d *D) FieldU19BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U19BE, sfns...)
}

// Reader U20BE

// TryU20BE tries to read 20 bit unsigned integer in big-endian
func (d *D) TryU20BE() (uint64, error) { return d.tryUE(20, BigEndian) }

func (d *D) ScalarU20BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(20, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U20BE reads 20 bit unsigned integer in big-endian
func (d *D) U20BE() uint64 {
	v, err := d.tryUE(20, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U20BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU20BE tries to add a field and read 20 bit unsigned integer in big-endian
func (d *D) TryFieldU20BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU20BE, sfns...)
}

// FieldU20BE adds a field and reads 20 bit unsigned integer in big-endian
func (d *D) FieldU20BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U20BE, sfns...)
}

// Reader U21BE

// TryU21BE tries to read 21 bit unsigned integer in big-endian
func (d *D) TryU21BE() (uint64, error) { return d.tryUE(21, BigEndian) }

func (d *D) ScalarU21BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(21, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U21BE reads 21 bit unsigned integer in big-endian
func (d *D) U21BE() uint64 {
	v, err := d.tryUE(21, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U21BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU21BE tries to add a field and read 21 bit unsigned integer in big-endian
func (d *D) TryFieldU21BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU21BE, sfns...)
}

// FieldU21BE adds a field and reads 21 bit unsigned integer in big-endian
func (d *D) FieldU21BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U21BE, sfns...)
}

// Reader U22BE

// TryU22BE tries to read 22 bit unsigned integer in big-endian
func (d *D) TryU22BE() (uint64, error) { return d.tryUE(22, BigEndian) }

func (d *D) ScalarU22BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(22, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U22BE reads 22 bit unsigned integer in big-endian
func (d *D) U22BE() uint64 {
	v, err := d.tryUE(22, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U22BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU22BE tries to add a field and read 22 bit unsigned integer in big-endian
func (d *D) TryFieldU22BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU22BE, sfns...)
}

// FieldU22BE adds a field and reads 22 bit unsigned integer in big-endian
func (d *D) FieldU22BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U22BE, sfns...)
}

// Reader U23BE

// TryU23BE tries to read 23 bit unsigned integer in big-endian
func (d *D) TryU23BE() (uint64, error) { return d.tryUE(23, BigEndian) }

func (d *D) ScalarU23BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(23, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U23BE reads 23 bit unsigned integer in big-endian
func (d *D) U23BE() uint64 {
	v, err := d.tryUE(23, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U23BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU23BE tries to add a field and read 23 bit unsigned integer in big-endian
func (d *D) TryFieldU23BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU23BE, sfns...)
}

// FieldU23BE adds a field and reads 23 bit unsigned integer in big-endian
func (d *D) FieldU23BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U23BE, sfns...)
}

// Reader U24BE

// TryU24BE tries to read 24 bit unsigned integer in big-endian
func (d *D) TryU24BE() (uint64, error) { return d.tryUE(24, BigEndian) }

func (d *D) ScalarU24BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(24, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U24BE reads 24 bit unsigned integer in big-endian
func (d *D) U24BE() uint64 {
	v, err := d.tryUE(24, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U24BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU24BE tries to add a field and read 24 bit unsigned integer in big-endian
func (d *D) TryFieldU24BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU24BE, sfns...)
}

// FieldU24BE adds a field and reads 24 bit unsigned integer in big-endian
func (d *D) FieldU24BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U24BE, sfns...)
}

// Reader U25BE

// TryU25BE tries to read 25 bit unsigned integer in big-endian
func (d *D) TryU25BE() (uint64, error) { return d.tryUE(25, BigEndian) }

func (d *D) ScalarU25BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(25, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U25BE reads 25 bit unsigned integer in big-endian
func (d *D) U25BE() uint64 {
	v, err := d.tryUE(25, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U25BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU25BE tries to add a field and read 25 bit unsigned integer in big-endian
func (d *D) TryFieldU25BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU25BE, sfns...)
}

// FieldU25BE adds a field and reads 25 bit unsigned integer in big-endian
func (d *D) FieldU25BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U25BE, sfns...)
}

// Reader U26BE

// TryU26BE tries to read 26 bit unsigned integer in big-endian
func (d *D) TryU26BE() (uint64, error) { return d.tryUE(26, BigEndian) }

func (d *D) ScalarU26BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(26, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U26BE reads 26 bit unsigned integer in big-endian
func (d *D) U26BE() uint64 {
	v, err := d.tryUE(26, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U26BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU26BE tries to add a field and read 26 bit unsigned integer in big-endian
func (d *D) TryFieldU26BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU26BE, sfns...)
}

// FieldU26BE adds a field and reads 26 bit unsigned integer in big-endian
func (d *D) FieldU26BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U26BE, sfns...)
}

// Reader U27BE

// TryU27BE tries to read 27 bit unsigned integer in big-endian
func (d *D) TryU27BE() (uint64, error) { return d.tryUE(27, BigEndian) }

func (d *D) ScalarU27BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(27, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U27BE reads 27 bit unsigned integer in big-endian
func (d *D) U27BE() uint64 {
	v, err := d.tryUE(27, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U27BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU27BE tries to add a field and read 27 bit unsigned integer in big-endian
func (d *D) TryFieldU27BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU27BE, sfns...)
}

// FieldU27BE adds a field and reads 27 bit unsigned integer in big-endian
func (d *D) FieldU27BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U27BE, sfns...)
}

// Reader U28BE

// TryU28BE tries to read 28 bit unsigned integer in big-endian
func (d *D) TryU28BE() (uint64, error) { return d.tryUE(28, BigEndian) }

func (d *D) ScalarU28BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(28, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U28BE reads 28 bit unsigned integer in big-endian
func (d *D) U28BE() uint64 {
	v, err := d.tryUE(28, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U28BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU28BE tries to add a field and read 28 bit unsigned integer in big-endian
func (d *D) TryFieldU28BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU28BE, sfns...)
}

// FieldU28BE adds a field and reads 28 bit unsigned integer in big-endian
func (d *D) FieldU28BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U28BE, sfns...)
}

// Reader U29BE

// TryU29BE tries to read 29 bit unsigned integer in big-endian
func (d *D) TryU29BE() (uint64, error) { return d.tryUE(29, BigEndian) }

func (d *D) ScalarU29BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(29, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U29BE reads 29 bit unsigned integer in big-endian
func (d *D) U29BE() uint64 {
	v, err := d.tryUE(29, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U29BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU29BE tries to add a field and read 29 bit unsigned integer in big-endian
func (d *D) TryFieldU29BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU29BE, sfns...)
}

// FieldU29BE adds a field and reads 29 bit unsigned integer in big-endian
func (d *D) FieldU29BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U29BE, sfns...)
}

// Reader U30BE

// TryU30BE tries to read 30 bit unsigned integer in big-endian
func (d *D) TryU30BE() (uint64, error) { return d.tryUE(30, BigEndian) }

func (d *D) ScalarU30BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(30, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U30BE reads 30 bit unsigned integer in big-endian
func (d *D) U30BE() uint64 {
	v, err := d.tryUE(30, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U30BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU30BE tries to add a field and read 30 bit unsigned integer in big-endian
func (d *D) TryFieldU30BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU30BE, sfns...)
}

// FieldU30BE adds a field and reads 30 bit unsigned integer in big-endian
func (d *D) FieldU30BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U30BE, sfns...)
}

// Reader U31BE

// TryU31BE tries to read 31 bit unsigned integer in big-endian
func (d *D) TryU31BE() (uint64, error) { return d.tryUE(31, BigEndian) }

func (d *D) ScalarU31BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(31, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U31BE reads 31 bit unsigned integer in big-endian
func (d *D) U31BE() uint64 {
	v, err := d.tryUE(31, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U31BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU31BE tries to add a field and read 31 bit unsigned integer in big-endian
func (d *D) TryFieldU31BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU31BE, sfns...)
}

// FieldU31BE adds a field and reads 31 bit unsigned integer in big-endian
func (d *D) FieldU31BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U31BE, sfns...)
}

// Reader U32BE

// TryU32BE tries to read 32 bit unsigned integer in big-endian
func (d *D) TryU32BE() (uint64, error) { return d.tryUE(32, BigEndian) }

func (d *D) ScalarU32BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(32, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U32BE reads 32 bit unsigned integer in big-endian
func (d *D) U32BE() uint64 {
	v, err := d.tryUE(32, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U32BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU32BE tries to add a field and read 32 bit unsigned integer in big-endian
func (d *D) TryFieldU32BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU32BE, sfns...)
}

// FieldU32BE adds a field and reads 32 bit unsigned integer in big-endian
func (d *D) FieldU32BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U32BE, sfns...)
}

// Reader U33BE

// TryU33BE tries to read 33 bit unsigned integer in big-endian
func (d *D) TryU33BE() (uint64, error) { return d.tryUE(33, BigEndian) }

func (d *D) ScalarU33BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(33, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U33BE reads 33 bit unsigned integer in big-endian
func (d *D) U33BE() uint64 {
	v, err := d.tryUE(33, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U33BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU33BE tries to add a field and read 33 bit unsigned integer in big-endian
func (d *D) TryFieldU33BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU33BE, sfns...)
}

// FieldU33BE adds a field and reads 33 bit unsigned integer in big-endian
func (d *D) FieldU33BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U33BE, sfns...)
}

// Reader U34BE

// TryU34BE tries to read 34 bit unsigned integer in big-endian
func (d *D) TryU34BE() (uint64, error) { return d.tryUE(34, BigEndian) }

func (d *D) ScalarU34BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(34, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U34BE reads 34 bit unsigned integer in big-endian
func (d *D) U34BE() uint64 {
	v, err := d.tryUE(34, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U34BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU34BE tries to add a field and read 34 bit unsigned integer in big-endian
func (d *D) TryFieldU34BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU34BE, sfns...)
}

// FieldU34BE adds a field and reads 34 bit unsigned integer in big-endian
func (d *D) FieldU34BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U34BE, sfns...)
}

// Reader U35BE

// TryU35BE tries to read 35 bit unsigned integer in big-endian
func (d *D) TryU35BE() (uint64, error) { return d.tryUE(35, BigEndian) }

func (d *D) ScalarU35BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(35, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U35BE reads 35 bit unsigned integer in big-endian
func (d *D) U35BE() uint64 {
	v, err := d.tryUE(35, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U35BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU35BE tries to add a field and read 35 bit unsigned integer in big-endian
func (d *D) TryFieldU35BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU35BE, sfns...)
}

// FieldU35BE adds a field and reads 35 bit unsigned integer in big-endian
func (d *D) FieldU35BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U35BE, sfns...)
}

// Reader U36BE

// TryU36BE tries to read 36 bit unsigned integer in big-endian
func (d *D) TryU36BE() (uint64, error) { return d.tryUE(36, BigEndian) }

func (d *D) ScalarU36BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(36, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U36BE reads 36 bit unsigned integer in big-endian
func (d *D) U36BE() uint64 {
	v, err := d.tryUE(36, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U36BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU36BE tries to add a field and read 36 bit unsigned integer in big-endian
func (d *D) TryFieldU36BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU36BE, sfns...)
}

// FieldU36BE adds a field and reads 36 bit unsigned integer in big-endian
func (d *D) FieldU36BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U36BE, sfns...)
}

// Reader U37BE

// TryU37BE tries to read 37 bit unsigned integer in big-endian
func (d *D) TryU37BE() (uint64, error) { return d.tryUE(37, BigEndian) }

func (d *D) ScalarU37BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(37, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U37BE reads 37 bit unsigned integer in big-endian
func (d *D) U37BE() uint64 {
	v, err := d.tryUE(37, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U37BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU37BE tries to add a field and read 37 bit unsigned integer in big-endian
func (d *D) TryFieldU37BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU37BE, sfns...)
}

// FieldU37BE adds a field and reads 37 bit unsigned integer in big-endian
func (d *D) FieldU37BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U37BE, sfns...)
}

// Reader U38BE

// TryU38BE tries to read 38 bit unsigned integer in big-endian
func (d *D) TryU38BE() (uint64, error) { return d.tryUE(38, BigEndian) }

func (d *D) ScalarU38BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(38, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U38BE reads 38 bit unsigned integer in big-endian
func (d *D) U38BE() uint64 {
	v, err := d.tryUE(38, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U38BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU38BE tries to add a field and read 38 bit unsigned integer in big-endian
func (d *D) TryFieldU38BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU38BE, sfns...)
}

// FieldU38BE adds a field and reads 38 bit unsigned integer in big-endian
func (d *D) FieldU38BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U38BE, sfns...)
}

// Reader U39BE

// TryU39BE tries to read 39 bit unsigned integer in big-endian
func (d *D) TryU39BE() (uint64, error) { return d.tryUE(39, BigEndian) }

func (d *D) ScalarU39BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(39, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U39BE reads 39 bit unsigned integer in big-endian
func (d *D) U39BE() uint64 {
	v, err := d.tryUE(39, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U39BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU39BE tries to add a field and read 39 bit unsigned integer in big-endian
func (d *D) TryFieldU39BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU39BE, sfns...)
}

// FieldU39BE adds a field and reads 39 bit unsigned integer in big-endian
func (d *D) FieldU39BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U39BE, sfns...)
}

// Reader U40BE

// TryU40BE tries to read 40 bit unsigned integer in big-endian
func (d *D) TryU40BE() (uint64, error) { return d.tryUE(40, BigEndian) }

func (d *D) ScalarU40BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(40, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U40BE reads 40 bit unsigned integer in big-endian
func (d *D) U40BE() uint64 {
	v, err := d.tryUE(40, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U40BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU40BE tries to add a field and read 40 bit unsigned integer in big-endian
func (d *D) TryFieldU40BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU40BE, sfns...)
}

// FieldU40BE adds a field and reads 40 bit unsigned integer in big-endian
func (d *D) FieldU40BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U40BE, sfns...)
}

// Reader U41BE

// TryU41BE tries to read 41 bit unsigned integer in big-endian
func (d *D) TryU41BE() (uint64, error) { return d.tryUE(41, BigEndian) }

func (d *D) ScalarU41BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(41, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U41BE reads 41 bit unsigned integer in big-endian
func (d *D) U41BE() uint64 {
	v, err := d.tryUE(41, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U41BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU41BE tries to add a field and read 41 bit unsigned integer in big-endian
func (d *D) TryFieldU41BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU41BE, sfns...)
}

// FieldU41BE adds a field and reads 41 bit unsigned integer in big-endian
func (d *D) FieldU41BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U41BE, sfns...)
}

// Reader U42BE

// TryU42BE tries to read 42 bit unsigned integer in big-endian
func (d *D) TryU42BE() (uint64, error) { return d.tryUE(42, BigEndian) }

func (d *D) ScalarU42BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(42, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U42BE reads 42 bit unsigned integer in big-endian
func (d *D) U42BE() uint64 {
	v, err := d.tryUE(42, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U42BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU42BE tries to add a field and read 42 bit unsigned integer in big-endian
func (d *D) TryFieldU42BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU42BE, sfns...)
}

// FieldU42BE adds a field and reads 42 bit unsigned integer in big-endian
func (d *D) FieldU42BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U42BE, sfns...)
}

// Reader U43BE

// TryU43BE tries to read 43 bit unsigned integer in big-endian
func (d *D) TryU43BE() (uint64, error) { return d.tryUE(43, BigEndian) }

func (d *D) ScalarU43BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(43, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U43BE reads 43 bit unsigned integer in big-endian
func (d *D) U43BE() uint64 {
	v, err := d.tryUE(43, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U43BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU43BE tries to add a field and read 43 bit unsigned integer in big-endian
func (d *D) TryFieldU43BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU43BE, sfns...)
}

// FieldU43BE adds a field and reads 43 bit unsigned integer in big-endian
func (d *D) FieldU43BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U43BE, sfns...)
}

// Reader U44BE

// TryU44BE tries to read 44 bit unsigned integer in big-endian
func (d *D) TryU44BE() (uint64, error) { return d.tryUE(44, BigEndian) }

func (d *D) ScalarU44BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(44, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U44BE reads 44 bit unsigned integer in big-endian
func (d *D) U44BE() uint64 {
	v, err := d.tryUE(44, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U44BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU44BE tries to add a field and read 44 bit unsigned integer in big-endian
func (d *D) TryFieldU44BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU44BE, sfns...)
}

// FieldU44BE adds a field and reads 44 bit unsigned integer in big-endian
func (d *D) FieldU44BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U44BE, sfns...)
}

// Reader U45BE

// TryU45BE tries to read 45 bit unsigned integer in big-endian
func (d *D) TryU45BE() (uint64, error) { return d.tryUE(45, BigEndian) }

func (d *D) ScalarU45BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(45, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U45BE reads 45 bit unsigned integer in big-endian
func (d *D) U45BE() uint64 {
	v, err := d.tryUE(45, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U45BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU45BE tries to add a field and read 45 bit unsigned integer in big-endian
func (d *D) TryFieldU45BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU45BE, sfns...)
}

// FieldU45BE adds a field and reads 45 bit unsigned integer in big-endian
func (d *D) FieldU45BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U45BE, sfns...)
}

// Reader U46BE

// TryU46BE tries to read 46 bit unsigned integer in big-endian
func (d *D) TryU46BE() (uint64, error) { return d.tryUE(46, BigEndian) }

func (d *D) ScalarU46BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(46, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U46BE reads 46 bit unsigned integer in big-endian
func (d *D) U46BE() uint64 {
	v, err := d.tryUE(46, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U46BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU46BE tries to add a field and read 46 bit unsigned integer in big-endian
func (d *D) TryFieldU46BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU46BE, sfns...)
}

// FieldU46BE adds a field and reads 46 bit unsigned integer in big-endian
func (d *D) FieldU46BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U46BE, sfns...)
}

// Reader U47BE

// TryU47BE tries to read 47 bit unsigned integer in big-endian
func (d *D) TryU47BE() (uint64, error) { return d.tryUE(47, BigEndian) }

func (d *D) ScalarU47BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(47, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U47BE reads 47 bit unsigned integer in big-endian
func (d *D) U47BE() uint64 {
	v, err := d.tryUE(47, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U47BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU47BE tries to add a field and read 47 bit unsigned integer in big-endian
func (d *D) TryFieldU47BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU47BE, sfns...)
}

// FieldU47BE adds a field and reads 47 bit unsigned integer in big-endian
func (d *D) FieldU47BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U47BE, sfns...)
}

// Reader U48BE

// TryU48BE tries to read 48 bit unsigned integer in big-endian
func (d *D) TryU48BE() (uint64, error) { return d.tryUE(48, BigEndian) }

func (d *D) ScalarU48BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(48, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U48BE reads 48 bit unsigned integer in big-endian
func (d *D) U48BE() uint64 {
	v, err := d.tryUE(48, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U48BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU48BE tries to add a field and read 48 bit unsigned integer in big-endian
func (d *D) TryFieldU48BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU48BE, sfns...)
}

// FieldU48BE adds a field and reads 48 bit unsigned integer in big-endian
func (d *D) FieldU48BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U48BE, sfns...)
}

// Reader U49BE

// TryU49BE tries to read 49 bit unsigned integer in big-endian
func (d *D) TryU49BE() (uint64, error) { return d.tryUE(49, BigEndian) }

func (d *D) ScalarU49BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(49, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U49BE reads 49 bit unsigned integer in big-endian
func (d *D) U49BE() uint64 {
	v, err := d.tryUE(49, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U49BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU49BE tries to add a field and read 49 bit unsigned integer in big-endian
func (d *D) TryFieldU49BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU49BE, sfns...)
}

// FieldU49BE adds a field and reads 49 bit unsigned integer in big-endian
func (d *D) FieldU49BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U49BE, sfns...)
}

// Reader U50BE

// TryU50BE tries to read 50 bit unsigned integer in big-endian
func (d *D) TryU50BE() (uint64, error) { return d.tryUE(50, BigEndian) }

func (d *D) ScalarU50BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(50, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U50BE reads 50 bit unsigned integer in big-endian
func (d *D) U50BE() uint64 {
	v, err := d.tryUE(50, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U50BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU50BE tries to add a field and read 50 bit unsigned integer in big-endian
func (d *D) TryFieldU50BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU50BE, sfns...)
}

// FieldU50BE adds a field and reads 50 bit unsigned integer in big-endian
func (d *D) FieldU50BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U50BE, sfns...)
}

// Reader U51BE

// TryU51BE tries to read 51 bit unsigned integer in big-endian
func (d *D) TryU51BE() (uint64, error) { return d.tryUE(51, BigEndian) }

func (d *D) ScalarU51BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(51, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U51BE reads 51 bit unsigned integer in big-endian
func (d *D) U51BE() uint64 {
	v, err := d.tryUE(51, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U51BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU51BE tries to add a field and read 51 bit unsigned integer in big-endian
func (d *D) TryFieldU51BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU51BE, sfns...)
}

// FieldU51BE adds a field and reads 51 bit unsigned integer in big-endian
func (d *D) FieldU51BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U51BE, sfns...)
}

// Reader U52BE

// TryU52BE tries to read 52 bit unsigned integer in big-endian
func (d *D) TryU52BE() (uint64, error) { return d.tryUE(52, BigEndian) }

func (d *D) ScalarU52BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(52, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U52BE reads 52 bit unsigned integer in big-endian
func (d *D) U52BE() uint64 {
	v, err := d.tryUE(52, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U52BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU52BE tries to add a field and read 52 bit unsigned integer in big-endian
func (d *D) TryFieldU52BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU52BE, sfns...)
}

// FieldU52BE adds a field and reads 52 bit unsigned integer in big-endian
func (d *D) FieldU52BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U52BE, sfns...)
}

// Reader U53BE

// TryU53BE tries to read 53 bit unsigned integer in big-endian
func (d *D) TryU53BE() (uint64, error) { return d.tryUE(53, BigEndian) }

func (d *D) ScalarU53BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(53, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U53BE reads 53 bit unsigned integer in big-endian
func (d *D) U53BE() uint64 {
	v, err := d.tryUE(53, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U53BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU53BE tries to add a field and read 53 bit unsigned integer in big-endian
func (d *D) TryFieldU53BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU53BE, sfns...)
}

// FieldU53BE adds a field and reads 53 bit unsigned integer in big-endian
func (d *D) FieldU53BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U53BE, sfns...)
}

// Reader U54BE

// TryU54BE tries to read 54 bit unsigned integer in big-endian
func (d *D) TryU54BE() (uint64, error) { return d.tryUE(54, BigEndian) }

func (d *D) ScalarU54BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(54, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U54BE reads 54 bit unsigned integer in big-endian
func (d *D) U54BE() uint64 {
	v, err := d.tryUE(54, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U54BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU54BE tries to add a field and read 54 bit unsigned integer in big-endian
func (d *D) TryFieldU54BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU54BE, sfns...)
}

// FieldU54BE adds a field and reads 54 bit unsigned integer in big-endian
func (d *D) FieldU54BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U54BE, sfns...)
}

// Reader U55BE

// TryU55BE tries to read 55 bit unsigned integer in big-endian
func (d *D) TryU55BE() (uint64, error) { return d.tryUE(55, BigEndian) }

func (d *D) ScalarU55BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(55, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U55BE reads 55 bit unsigned integer in big-endian
func (d *D) U55BE() uint64 {
	v, err := d.tryUE(55, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U55BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU55BE tries to add a field and read 55 bit unsigned integer in big-endian
func (d *D) TryFieldU55BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU55BE, sfns...)
}

// FieldU55BE adds a field and reads 55 bit unsigned integer in big-endian
func (d *D) FieldU55BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U55BE, sfns...)
}

// Reader U56BE

// TryU56BE tries to read 56 bit unsigned integer in big-endian
func (d *D) TryU56BE() (uint64, error) { return d.tryUE(56, BigEndian) }

func (d *D) ScalarU56BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(56, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U56BE reads 56 bit unsigned integer in big-endian
func (d *D) U56BE() uint64 {
	v, err := d.tryUE(56, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U56BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU56BE tries to add a field and read 56 bit unsigned integer in big-endian
func (d *D) TryFieldU56BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU56BE, sfns...)
}

// FieldU56BE adds a field and reads 56 bit unsigned integer in big-endian
func (d *D) FieldU56BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U56BE, sfns...)
}

// Reader U57BE

// TryU57BE tries to read 57 bit unsigned integer in big-endian
func (d *D) TryU57BE() (uint64, error) { return d.tryUE(57, BigEndian) }

func (d *D) ScalarU57BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(57, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U57BE reads 57 bit unsigned integer in big-endian
func (d *D) U57BE() uint64 {
	v, err := d.tryUE(57, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U57BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU57BE tries to add a field and read 57 bit unsigned integer in big-endian
func (d *D) TryFieldU57BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU57BE, sfns...)
}

// FieldU57BE adds a field and reads 57 bit unsigned integer in big-endian
func (d *D) FieldU57BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U57BE, sfns...)
}

// Reader U58BE

// TryU58BE tries to read 58 bit unsigned integer in big-endian
func (d *D) TryU58BE() (uint64, error) { return d.tryUE(58, BigEndian) }

func (d *D) ScalarU58BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(58, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U58BE reads 58 bit unsigned integer in big-endian
func (d *D) U58BE() uint64 {
	v, err := d.tryUE(58, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U58BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU58BE tries to add a field and read 58 bit unsigned integer in big-endian
func (d *D) TryFieldU58BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU58BE, sfns...)
}

// FieldU58BE adds a field and reads 58 bit unsigned integer in big-endian
func (d *D) FieldU58BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U58BE, sfns...)
}

// Reader U59BE

// TryU59BE tries to read 59 bit unsigned integer in big-endian
func (d *D) TryU59BE() (uint64, error) { return d.tryUE(59, BigEndian) }

func (d *D) ScalarU59BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(59, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U59BE reads 59 bit unsigned integer in big-endian
func (d *D) U59BE() uint64 {
	v, err := d.tryUE(59, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U59BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU59BE tries to add a field and read 59 bit unsigned integer in big-endian
func (d *D) TryFieldU59BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU59BE, sfns...)
}

// FieldU59BE adds a field and reads 59 bit unsigned integer in big-endian
func (d *D) FieldU59BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U59BE, sfns...)
}

// Reader U60BE

// TryU60BE tries to read 60 bit unsigned integer in big-endian
func (d *D) TryU60BE() (uint64, error) { return d.tryUE(60, BigEndian) }

func (d *D) ScalarU60BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(60, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U60BE reads 60 bit unsigned integer in big-endian
func (d *D) U60BE() uint64 {
	v, err := d.tryUE(60, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U60BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU60BE tries to add a field and read 60 bit unsigned integer in big-endian
func (d *D) TryFieldU60BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU60BE, sfns...)
}

// FieldU60BE adds a field and reads 60 bit unsigned integer in big-endian
func (d *D) FieldU60BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U60BE, sfns...)
}

// Reader U61BE

// TryU61BE tries to read 61 bit unsigned integer in big-endian
func (d *D) TryU61BE() (uint64, error) { return d.tryUE(61, BigEndian) }

func (d *D) ScalarU61BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(61, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U61BE reads 61 bit unsigned integer in big-endian
func (d *D) U61BE() uint64 {
	v, err := d.tryUE(61, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U61BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU61BE tries to add a field and read 61 bit unsigned integer in big-endian
func (d *D) TryFieldU61BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU61BE, sfns...)
}

// FieldU61BE adds a field and reads 61 bit unsigned integer in big-endian
func (d *D) FieldU61BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U61BE, sfns...)
}

// Reader U62BE

// TryU62BE tries to read 62 bit unsigned integer in big-endian
func (d *D) TryU62BE() (uint64, error) { return d.tryUE(62, BigEndian) }

func (d *D) ScalarU62BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(62, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U62BE reads 62 bit unsigned integer in big-endian
func (d *D) U62BE() uint64 {
	v, err := d.tryUE(62, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U62BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU62BE tries to add a field and read 62 bit unsigned integer in big-endian
func (d *D) TryFieldU62BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU62BE, sfns...)
}

// FieldU62BE adds a field and reads 62 bit unsigned integer in big-endian
func (d *D) FieldU62BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U62BE, sfns...)
}

// Reader U63BE

// TryU63BE tries to read 63 bit unsigned integer in big-endian
func (d *D) TryU63BE() (uint64, error) { return d.tryUE(63, BigEndian) }

func (d *D) ScalarU63BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(63, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U63BE reads 63 bit unsigned integer in big-endian
func (d *D) U63BE() uint64 {
	v, err := d.tryUE(63, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U63BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU63BE tries to add a field and read 63 bit unsigned integer in big-endian
func (d *D) TryFieldU63BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU63BE, sfns...)
}

// FieldU63BE adds a field and reads 63 bit unsigned integer in big-endian
func (d *D) FieldU63BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U63BE, sfns...)
}

// Reader U64BE

// TryU64BE tries to read 64 bit unsigned integer in big-endian
func (d *D) TryU64BE() (uint64, error) { return d.tryUE(64, BigEndian) }

func (d *D) ScalarU64BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUE(64, BigEndian)
		s.Actual = v
		return s, err
	}
}

// U64BE reads 64 bit unsigned integer in big-endian
func (d *D) U64BE() uint64 {
	v, err := d.tryUE(64, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "U64BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldU64BE tries to add a field and read 64 bit unsigned integer in big-endian
func (d *D) TryFieldU64BE(name string, sfns ...ScalarFn) (uint64, error) {
	return d.TryFieldUFn(name, (*D).TryU64BE, sfns...)
}

// FieldU64BE adds a field and reads 64 bit unsigned integer in big-endian
func (d *D) FieldU64BE(name string, sfns ...ScalarFn) uint64 {
	return d.FieldUFn(name, (*D).U64BE, sfns...)
}

// Reader S

// TryS tries to read nBits bits signed integer in current endian
func (d *D) TryS(nBits int) (int64, error) { return d.trySE(nBits, d.Endian) }

func (d *D) ScalarS(nBits int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(nBits, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S reads nBits bits signed integer in current endian
func (d *D) S(nBits int) int64 {
	v, err := d.trySE(nBits, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S", Pos: d.Pos()})
	}
	return v
}

// TryFieldS tries to add a field and read nBits bits signed integer in current endian
func (d *D) TryFieldS(name string, nBits int, sfns ...ScalarFn) (int64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarS(nBits), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualS(), err
}

// FieldS adds a field and reads nBits bits signed integer in current endian
func (d *D) FieldS(name string, nBits int, sfns ...ScalarFn) int64 {
	v, err := d.TryFieldS(name, nBits, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "S", Pos: d.Pos()})
	}
	return v
}

// Reader SE

// TrySE tries to read nBits signed integer in specified endian
func (d *D) TrySE(nBits int, endian Endian) (int64, error) { return d.trySE(nBits, endian) }

func (d *D) ScalarSE(nBits int, endian Endian) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(nBits, endian)
		s.Actual = v
		return s, err
	}
}

// SE reads nBits signed integer in specified endian
func (d *D) SE(nBits int, endian Endian) int64 {
	v, err := d.trySE(nBits, endian)
	if err != nil {
		panic(IOError{Err: err, Op: "SE", Pos: d.Pos()})
	}
	return v
}

// TryFieldSE tries to add a field and read nBits signed integer in specified endian
func (d *D) TryFieldSE(name string, nBits int, endian Endian, sfns ...ScalarFn) (int64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarSE(nBits, endian), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualS(), err
}

// FieldSE adds a field and reads nBits signed integer in specified endian
func (d *D) FieldSE(name string, nBits int, endian Endian, sfns ...ScalarFn) int64 {
	v, err := d.TryFieldSE(name, nBits, endian, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "SE", Pos: d.Pos()})
	}
	return v
}

// Reader S1

// TryS1 tries to read 1 bit signed integer in current endian
func (d *D) TryS1() (int64, error) { return d.trySE(1, d.Endian) }

func (d *D) ScalarS1() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(1, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S1 reads 1 bit signed integer in current endian
func (d *D) S1() int64 {
	v, err := d.trySE(1, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S1", Pos: d.Pos()})
	}
	return v
}

// TryFieldS1 tries to add a field and read 1 bit signed integer in current endian
func (d *D) TryFieldS1(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS1, sfns...)
}

// FieldS1 adds a field and reads 1 bit signed integer in current endian
func (d *D) FieldS1(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S1, sfns...)
}

// Reader S2

// TryS2 tries to read 2 bit signed integer in current endian
func (d *D) TryS2() (int64, error) { return d.trySE(2, d.Endian) }

func (d *D) ScalarS2() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(2, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S2 reads 2 bit signed integer in current endian
func (d *D) S2() int64 {
	v, err := d.trySE(2, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S2", Pos: d.Pos()})
	}
	return v
}

// TryFieldS2 tries to add a field and read 2 bit signed integer in current endian
func (d *D) TryFieldS2(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS2, sfns...)
}

// FieldS2 adds a field and reads 2 bit signed integer in current endian
func (d *D) FieldS2(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S2, sfns...)
}

// Reader S3

// TryS3 tries to read 3 bit signed integer in current endian
func (d *D) TryS3() (int64, error) { return d.trySE(3, d.Endian) }

func (d *D) ScalarS3() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(3, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S3 reads 3 bit signed integer in current endian
func (d *D) S3() int64 {
	v, err := d.trySE(3, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S3", Pos: d.Pos()})
	}
	return v
}

// TryFieldS3 tries to add a field and read 3 bit signed integer in current endian
func (d *D) TryFieldS3(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS3, sfns...)
}

// FieldS3 adds a field and reads 3 bit signed integer in current endian
func (d *D) FieldS3(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S3, sfns...)
}

// Reader S4

// TryS4 tries to read 4 bit signed integer in current endian
func (d *D) TryS4() (int64, error) { return d.trySE(4, d.Endian) }

func (d *D) ScalarS4() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(4, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S4 reads 4 bit signed integer in current endian
func (d *D) S4() int64 {
	v, err := d.trySE(4, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S4", Pos: d.Pos()})
	}
	return v
}

// TryFieldS4 tries to add a field and read 4 bit signed integer in current endian
func (d *D) TryFieldS4(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS4, sfns...)
}

// FieldS4 adds a field and reads 4 bit signed integer in current endian
func (d *D) FieldS4(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S4, sfns...)
}

// Reader S5

// TryS5 tries to read 5 bit signed integer in current endian
func (d *D) TryS5() (int64, error) { return d.trySE(5, d.Endian) }

func (d *D) ScalarS5() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(5, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S5 reads 5 bit signed integer in current endian
func (d *D) S5() int64 {
	v, err := d.trySE(5, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S5", Pos: d.Pos()})
	}
	return v
}

// TryFieldS5 tries to add a field and read 5 bit signed integer in current endian
func (d *D) TryFieldS5(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS5, sfns...)
}

// FieldS5 adds a field and reads 5 bit signed integer in current endian
func (d *D) FieldS5(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S5, sfns...)
}

// Reader S6

// TryS6 tries to read 6 bit signed integer in current endian
func (d *D) TryS6() (int64, error) { return d.trySE(6, d.Endian) }

func (d *D) ScalarS6() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(6, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S6 reads 6 bit signed integer in current endian
func (d *D) S6() int64 {
	v, err := d.trySE(6, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S6", Pos: d.Pos()})
	}
	return v
}

// TryFieldS6 tries to add a field and read 6 bit signed integer in current endian
func (d *D) TryFieldS6(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS6, sfns...)
}

// FieldS6 adds a field and reads 6 bit signed integer in current endian
func (d *D) FieldS6(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S6, sfns...)
}

// Reader S7

// TryS7 tries to read 7 bit signed integer in current endian
func (d *D) TryS7() (int64, error) { return d.trySE(7, d.Endian) }

func (d *D) ScalarS7() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(7, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S7 reads 7 bit signed integer in current endian
func (d *D) S7() int64 {
	v, err := d.trySE(7, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S7", Pos: d.Pos()})
	}
	return v
}

// TryFieldS7 tries to add a field and read 7 bit signed integer in current endian
func (d *D) TryFieldS7(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS7, sfns...)
}

// FieldS7 adds a field and reads 7 bit signed integer in current endian
func (d *D) FieldS7(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S7, sfns...)
}

// Reader S8

// TryS8 tries to read 8 bit signed integer in current endian
func (d *D) TryS8() (int64, error) { return d.trySE(8, d.Endian) }

func (d *D) ScalarS8() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(8, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S8 reads 8 bit signed integer in current endian
func (d *D) S8() int64 {
	v, err := d.trySE(8, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S8", Pos: d.Pos()})
	}
	return v
}

// TryFieldS8 tries to add a field and read 8 bit signed integer in current endian
func (d *D) TryFieldS8(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS8, sfns...)
}

// FieldS8 adds a field and reads 8 bit signed integer in current endian
func (d *D) FieldS8(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S8, sfns...)
}

// Reader S9

// TryS9 tries to read 9 bit signed integer in current endian
func (d *D) TryS9() (int64, error) { return d.trySE(9, d.Endian) }

func (d *D) ScalarS9() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(9, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S9 reads 9 bit signed integer in current endian
func (d *D) S9() int64 {
	v, err := d.trySE(9, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S9", Pos: d.Pos()})
	}
	return v
}

// TryFieldS9 tries to add a field and read 9 bit signed integer in current endian
func (d *D) TryFieldS9(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS9, sfns...)
}

// FieldS9 adds a field and reads 9 bit signed integer in current endian
func (d *D) FieldS9(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S9, sfns...)
}

// Reader S10

// TryS10 tries to read 10 bit signed integer in current endian
func (d *D) TryS10() (int64, error) { return d.trySE(10, d.Endian) }

func (d *D) ScalarS10() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(10, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S10 reads 10 bit signed integer in current endian
func (d *D) S10() int64 {
	v, err := d.trySE(10, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S10", Pos: d.Pos()})
	}
	return v
}

// TryFieldS10 tries to add a field and read 10 bit signed integer in current endian
func (d *D) TryFieldS10(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS10, sfns...)
}

// FieldS10 adds a field and reads 10 bit signed integer in current endian
func (d *D) FieldS10(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S10, sfns...)
}

// Reader S11

// TryS11 tries to read 11 bit signed integer in current endian
func (d *D) TryS11() (int64, error) { return d.trySE(11, d.Endian) }

func (d *D) ScalarS11() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(11, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S11 reads 11 bit signed integer in current endian
func (d *D) S11() int64 {
	v, err := d.trySE(11, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S11", Pos: d.Pos()})
	}
	return v
}

// TryFieldS11 tries to add a field and read 11 bit signed integer in current endian
func (d *D) TryFieldS11(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS11, sfns...)
}

// FieldS11 adds a field and reads 11 bit signed integer in current endian
func (d *D) FieldS11(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S11, sfns...)
}

// Reader S12

// TryS12 tries to read 12 bit signed integer in current endian
func (d *D) TryS12() (int64, error) { return d.trySE(12, d.Endian) }

func (d *D) ScalarS12() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(12, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S12 reads 12 bit signed integer in current endian
func (d *D) S12() int64 {
	v, err := d.trySE(12, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S12", Pos: d.Pos()})
	}
	return v
}

// TryFieldS12 tries to add a field and read 12 bit signed integer in current endian
func (d *D) TryFieldS12(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS12, sfns...)
}

// FieldS12 adds a field and reads 12 bit signed integer in current endian
func (d *D) FieldS12(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S12, sfns...)
}

// Reader S13

// TryS13 tries to read 13 bit signed integer in current endian
func (d *D) TryS13() (int64, error) { return d.trySE(13, d.Endian) }

func (d *D) ScalarS13() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(13, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S13 reads 13 bit signed integer in current endian
func (d *D) S13() int64 {
	v, err := d.trySE(13, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S13", Pos: d.Pos()})
	}
	return v
}

// TryFieldS13 tries to add a field and read 13 bit signed integer in current endian
func (d *D) TryFieldS13(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS13, sfns...)
}

// FieldS13 adds a field and reads 13 bit signed integer in current endian
func (d *D) FieldS13(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S13, sfns...)
}

// Reader S14

// TryS14 tries to read 14 bit signed integer in current endian
func (d *D) TryS14() (int64, error) { return d.trySE(14, d.Endian) }

func (d *D) ScalarS14() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(14, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S14 reads 14 bit signed integer in current endian
func (d *D) S14() int64 {
	v, err := d.trySE(14, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S14", Pos: d.Pos()})
	}
	return v
}

// TryFieldS14 tries to add a field and read 14 bit signed integer in current endian
func (d *D) TryFieldS14(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS14, sfns...)
}

// FieldS14 adds a field and reads 14 bit signed integer in current endian
func (d *D) FieldS14(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S14, sfns...)
}

// Reader S15

// TryS15 tries to read 15 bit signed integer in current endian
func (d *D) TryS15() (int64, error) { return d.trySE(15, d.Endian) }

func (d *D) ScalarS15() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(15, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S15 reads 15 bit signed integer in current endian
func (d *D) S15() int64 {
	v, err := d.trySE(15, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S15", Pos: d.Pos()})
	}
	return v
}

// TryFieldS15 tries to add a field and read 15 bit signed integer in current endian
func (d *D) TryFieldS15(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS15, sfns...)
}

// FieldS15 adds a field and reads 15 bit signed integer in current endian
func (d *D) FieldS15(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S15, sfns...)
}

// Reader S16

// TryS16 tries to read 16 bit signed integer in current endian
func (d *D) TryS16() (int64, error) { return d.trySE(16, d.Endian) }

func (d *D) ScalarS16() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(16, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S16 reads 16 bit signed integer in current endian
func (d *D) S16() int64 {
	v, err := d.trySE(16, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S16", Pos: d.Pos()})
	}
	return v
}

// TryFieldS16 tries to add a field and read 16 bit signed integer in current endian
func (d *D) TryFieldS16(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS16, sfns...)
}

// FieldS16 adds a field and reads 16 bit signed integer in current endian
func (d *D) FieldS16(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S16, sfns...)
}

// Reader S17

// TryS17 tries to read 17 bit signed integer in current endian
func (d *D) TryS17() (int64, error) { return d.trySE(17, d.Endian) }

func (d *D) ScalarS17() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(17, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S17 reads 17 bit signed integer in current endian
func (d *D) S17() int64 {
	v, err := d.trySE(17, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S17", Pos: d.Pos()})
	}
	return v
}

// TryFieldS17 tries to add a field and read 17 bit signed integer in current endian
func (d *D) TryFieldS17(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS17, sfns...)
}

// FieldS17 adds a field and reads 17 bit signed integer in current endian
func (d *D) FieldS17(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S17, sfns...)
}

// Reader S18

// TryS18 tries to read 18 bit signed integer in current endian
func (d *D) TryS18() (int64, error) { return d.trySE(18, d.Endian) }

func (d *D) ScalarS18() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(18, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S18 reads 18 bit signed integer in current endian
func (d *D) S18() int64 {
	v, err := d.trySE(18, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S18", Pos: d.Pos()})
	}
	return v
}

// TryFieldS18 tries to add a field and read 18 bit signed integer in current endian
func (d *D) TryFieldS18(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS18, sfns...)
}

// FieldS18 adds a field and reads 18 bit signed integer in current endian
func (d *D) FieldS18(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S18, sfns...)
}

// Reader S19

// TryS19 tries to read 19 bit signed integer in current endian
func (d *D) TryS19() (int64, error) { return d.trySE(19, d.Endian) }

func (d *D) ScalarS19() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(19, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S19 reads 19 bit signed integer in current endian
func (d *D) S19() int64 {
	v, err := d.trySE(19, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S19", Pos: d.Pos()})
	}
	return v
}

// TryFieldS19 tries to add a field and read 19 bit signed integer in current endian
func (d *D) TryFieldS19(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS19, sfns...)
}

// FieldS19 adds a field and reads 19 bit signed integer in current endian
func (d *D) FieldS19(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S19, sfns...)
}

// Reader S20

// TryS20 tries to read 20 bit signed integer in current endian
func (d *D) TryS20() (int64, error) { return d.trySE(20, d.Endian) }

func (d *D) ScalarS20() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(20, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S20 reads 20 bit signed integer in current endian
func (d *D) S20() int64 {
	v, err := d.trySE(20, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S20", Pos: d.Pos()})
	}
	return v
}

// TryFieldS20 tries to add a field and read 20 bit signed integer in current endian
func (d *D) TryFieldS20(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS20, sfns...)
}

// FieldS20 adds a field and reads 20 bit signed integer in current endian
func (d *D) FieldS20(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S20, sfns...)
}

// Reader S21

// TryS21 tries to read 21 bit signed integer in current endian
func (d *D) TryS21() (int64, error) { return d.trySE(21, d.Endian) }

func (d *D) ScalarS21() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(21, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S21 reads 21 bit signed integer in current endian
func (d *D) S21() int64 {
	v, err := d.trySE(21, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S21", Pos: d.Pos()})
	}
	return v
}

// TryFieldS21 tries to add a field and read 21 bit signed integer in current endian
func (d *D) TryFieldS21(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS21, sfns...)
}

// FieldS21 adds a field and reads 21 bit signed integer in current endian
func (d *D) FieldS21(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S21, sfns...)
}

// Reader S22

// TryS22 tries to read 22 bit signed integer in current endian
func (d *D) TryS22() (int64, error) { return d.trySE(22, d.Endian) }

func (d *D) ScalarS22() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(22, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S22 reads 22 bit signed integer in current endian
func (d *D) S22() int64 {
	v, err := d.trySE(22, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S22", Pos: d.Pos()})
	}
	return v
}

// TryFieldS22 tries to add a field and read 22 bit signed integer in current endian
func (d *D) TryFieldS22(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS22, sfns...)
}

// FieldS22 adds a field and reads 22 bit signed integer in current endian
func (d *D) FieldS22(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S22, sfns...)
}

// Reader S23

// TryS23 tries to read 23 bit signed integer in current endian
func (d *D) TryS23() (int64, error) { return d.trySE(23, d.Endian) }

func (d *D) ScalarS23() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(23, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S23 reads 23 bit signed integer in current endian
func (d *D) S23() int64 {
	v, err := d.trySE(23, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S23", Pos: d.Pos()})
	}
	return v
}

// TryFieldS23 tries to add a field and read 23 bit signed integer in current endian
func (d *D) TryFieldS23(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS23, sfns...)
}

// FieldS23 adds a field and reads 23 bit signed integer in current endian
func (d *D) FieldS23(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S23, sfns...)
}

// Reader S24

// TryS24 tries to read 24 bit signed integer in current endian
func (d *D) TryS24() (int64, error) { return d.trySE(24, d.Endian) }

func (d *D) ScalarS24() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(24, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S24 reads 24 bit signed integer in current endian
func (d *D) S24() int64 {
	v, err := d.trySE(24, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S24", Pos: d.Pos()})
	}
	return v
}

// TryFieldS24 tries to add a field and read 24 bit signed integer in current endian
func (d *D) TryFieldS24(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS24, sfns...)
}

// FieldS24 adds a field and reads 24 bit signed integer in current endian
func (d *D) FieldS24(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S24, sfns...)
}

// Reader S25

// TryS25 tries to read 25 bit signed integer in current endian
func (d *D) TryS25() (int64, error) { return d.trySE(25, d.Endian) }

func (d *D) ScalarS25() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(25, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S25 reads 25 bit signed integer in current endian
func (d *D) S25() int64 {
	v, err := d.trySE(25, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S25", Pos: d.Pos()})
	}
	return v
}

// TryFieldS25 tries to add a field and read 25 bit signed integer in current endian
func (d *D) TryFieldS25(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS25, sfns...)
}

// FieldS25 adds a field and reads 25 bit signed integer in current endian
func (d *D) FieldS25(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S25, sfns...)
}

// Reader S26

// TryS26 tries to read 26 bit signed integer in current endian
func (d *D) TryS26() (int64, error) { return d.trySE(26, d.Endian) }

func (d *D) ScalarS26() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(26, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S26 reads 26 bit signed integer in current endian
func (d *D) S26() int64 {
	v, err := d.trySE(26, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S26", Pos: d.Pos()})
	}
	return v
}

// TryFieldS26 tries to add a field and read 26 bit signed integer in current endian
func (d *D) TryFieldS26(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS26, sfns...)
}

// FieldS26 adds a field and reads 26 bit signed integer in current endian
func (d *D) FieldS26(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S26, sfns...)
}

// Reader S27

// TryS27 tries to read 27 bit signed integer in current endian
func (d *D) TryS27() (int64, error) { return d.trySE(27, d.Endian) }

func (d *D) ScalarS27() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(27, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S27 reads 27 bit signed integer in current endian
func (d *D) S27() int64 {
	v, err := d.trySE(27, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S27", Pos: d.Pos()})
	}
	return v
}

// TryFieldS27 tries to add a field and read 27 bit signed integer in current endian
func (d *D) TryFieldS27(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS27, sfns...)
}

// FieldS27 adds a field and reads 27 bit signed integer in current endian
func (d *D) FieldS27(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S27, sfns...)
}

// Reader S28

// TryS28 tries to read 28 bit signed integer in current endian
func (d *D) TryS28() (int64, error) { return d.trySE(28, d.Endian) }

func (d *D) ScalarS28() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(28, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S28 reads 28 bit signed integer in current endian
func (d *D) S28() int64 {
	v, err := d.trySE(28, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S28", Pos: d.Pos()})
	}
	return v
}

// TryFieldS28 tries to add a field and read 28 bit signed integer in current endian
func (d *D) TryFieldS28(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS28, sfns...)
}

// FieldS28 adds a field and reads 28 bit signed integer in current endian
func (d *D) FieldS28(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S28, sfns...)
}

// Reader S29

// TryS29 tries to read 29 bit signed integer in current endian
func (d *D) TryS29() (int64, error) { return d.trySE(29, d.Endian) }

func (d *D) ScalarS29() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(29, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S29 reads 29 bit signed integer in current endian
func (d *D) S29() int64 {
	v, err := d.trySE(29, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S29", Pos: d.Pos()})
	}
	return v
}

// TryFieldS29 tries to add a field and read 29 bit signed integer in current endian
func (d *D) TryFieldS29(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS29, sfns...)
}

// FieldS29 adds a field and reads 29 bit signed integer in current endian
func (d *D) FieldS29(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S29, sfns...)
}

// Reader S30

// TryS30 tries to read 30 bit signed integer in current endian
func (d *D) TryS30() (int64, error) { return d.trySE(30, d.Endian) }

func (d *D) ScalarS30() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(30, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S30 reads 30 bit signed integer in current endian
func (d *D) S30() int64 {
	v, err := d.trySE(30, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S30", Pos: d.Pos()})
	}
	return v
}

// TryFieldS30 tries to add a field and read 30 bit signed integer in current endian
func (d *D) TryFieldS30(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS30, sfns...)
}

// FieldS30 adds a field and reads 30 bit signed integer in current endian
func (d *D) FieldS30(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S30, sfns...)
}

// Reader S31

// TryS31 tries to read 31 bit signed integer in current endian
func (d *D) TryS31() (int64, error) { return d.trySE(31, d.Endian) }

func (d *D) ScalarS31() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(31, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S31 reads 31 bit signed integer in current endian
func (d *D) S31() int64 {
	v, err := d.trySE(31, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S31", Pos: d.Pos()})
	}
	return v
}

// TryFieldS31 tries to add a field and read 31 bit signed integer in current endian
func (d *D) TryFieldS31(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS31, sfns...)
}

// FieldS31 adds a field and reads 31 bit signed integer in current endian
func (d *D) FieldS31(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S31, sfns...)
}

// Reader S32

// TryS32 tries to read 32 bit signed integer in current endian
func (d *D) TryS32() (int64, error) { return d.trySE(32, d.Endian) }

func (d *D) ScalarS32() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(32, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S32 reads 32 bit signed integer in current endian
func (d *D) S32() int64 {
	v, err := d.trySE(32, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S32", Pos: d.Pos()})
	}
	return v
}

// TryFieldS32 tries to add a field and read 32 bit signed integer in current endian
func (d *D) TryFieldS32(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS32, sfns...)
}

// FieldS32 adds a field and reads 32 bit signed integer in current endian
func (d *D) FieldS32(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S32, sfns...)
}

// Reader S33

// TryS33 tries to read 33 bit signed integer in current endian
func (d *D) TryS33() (int64, error) { return d.trySE(33, d.Endian) }

func (d *D) ScalarS33() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(33, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S33 reads 33 bit signed integer in current endian
func (d *D) S33() int64 {
	v, err := d.trySE(33, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S33", Pos: d.Pos()})
	}
	return v
}

// TryFieldS33 tries to add a field and read 33 bit signed integer in current endian
func (d *D) TryFieldS33(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS33, sfns...)
}

// FieldS33 adds a field and reads 33 bit signed integer in current endian
func (d *D) FieldS33(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S33, sfns...)
}

// Reader S34

// TryS34 tries to read 34 bit signed integer in current endian
func (d *D) TryS34() (int64, error) { return d.trySE(34, d.Endian) }

func (d *D) ScalarS34() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(34, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S34 reads 34 bit signed integer in current endian
func (d *D) S34() int64 {
	v, err := d.trySE(34, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S34", Pos: d.Pos()})
	}
	return v
}

// TryFieldS34 tries to add a field and read 34 bit signed integer in current endian
func (d *D) TryFieldS34(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS34, sfns...)
}

// FieldS34 adds a field and reads 34 bit signed integer in current endian
func (d *D) FieldS34(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S34, sfns...)
}

// Reader S35

// TryS35 tries to read 35 bit signed integer in current endian
func (d *D) TryS35() (int64, error) { return d.trySE(35, d.Endian) }

func (d *D) ScalarS35() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(35, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S35 reads 35 bit signed integer in current endian
func (d *D) S35() int64 {
	v, err := d.trySE(35, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S35", Pos: d.Pos()})
	}
	return v
}

// TryFieldS35 tries to add a field and read 35 bit signed integer in current endian
func (d *D) TryFieldS35(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS35, sfns...)
}

// FieldS35 adds a field and reads 35 bit signed integer in current endian
func (d *D) FieldS35(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S35, sfns...)
}

// Reader S36

// TryS36 tries to read 36 bit signed integer in current endian
func (d *D) TryS36() (int64, error) { return d.trySE(36, d.Endian) }

func (d *D) ScalarS36() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(36, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S36 reads 36 bit signed integer in current endian
func (d *D) S36() int64 {
	v, err := d.trySE(36, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S36", Pos: d.Pos()})
	}
	return v
}

// TryFieldS36 tries to add a field and read 36 bit signed integer in current endian
func (d *D) TryFieldS36(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS36, sfns...)
}

// FieldS36 adds a field and reads 36 bit signed integer in current endian
func (d *D) FieldS36(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S36, sfns...)
}

// Reader S37

// TryS37 tries to read 37 bit signed integer in current endian
func (d *D) TryS37() (int64, error) { return d.trySE(37, d.Endian) }

func (d *D) ScalarS37() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(37, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S37 reads 37 bit signed integer in current endian
func (d *D) S37() int64 {
	v, err := d.trySE(37, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S37", Pos: d.Pos()})
	}
	return v
}

// TryFieldS37 tries to add a field and read 37 bit signed integer in current endian
func (d *D) TryFieldS37(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS37, sfns...)
}

// FieldS37 adds a field and reads 37 bit signed integer in current endian
func (d *D) FieldS37(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S37, sfns...)
}

// Reader S38

// TryS38 tries to read 38 bit signed integer in current endian
func (d *D) TryS38() (int64, error) { return d.trySE(38, d.Endian) }

func (d *D) ScalarS38() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(38, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S38 reads 38 bit signed integer in current endian
func (d *D) S38() int64 {
	v, err := d.trySE(38, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S38", Pos: d.Pos()})
	}
	return v
}

// TryFieldS38 tries to add a field and read 38 bit signed integer in current endian
func (d *D) TryFieldS38(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS38, sfns...)
}

// FieldS38 adds a field and reads 38 bit signed integer in current endian
func (d *D) FieldS38(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S38, sfns...)
}

// Reader S39

// TryS39 tries to read 39 bit signed integer in current endian
func (d *D) TryS39() (int64, error) { return d.trySE(39, d.Endian) }

func (d *D) ScalarS39() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(39, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S39 reads 39 bit signed integer in current endian
func (d *D) S39() int64 {
	v, err := d.trySE(39, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S39", Pos: d.Pos()})
	}
	return v
}

// TryFieldS39 tries to add a field and read 39 bit signed integer in current endian
func (d *D) TryFieldS39(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS39, sfns...)
}

// FieldS39 adds a field and reads 39 bit signed integer in current endian
func (d *D) FieldS39(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S39, sfns...)
}

// Reader S40

// TryS40 tries to read 40 bit signed integer in current endian
func (d *D) TryS40() (int64, error) { return d.trySE(40, d.Endian) }

func (d *D) ScalarS40() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(40, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S40 reads 40 bit signed integer in current endian
func (d *D) S40() int64 {
	v, err := d.trySE(40, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S40", Pos: d.Pos()})
	}
	return v
}

// TryFieldS40 tries to add a field and read 40 bit signed integer in current endian
func (d *D) TryFieldS40(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS40, sfns...)
}

// FieldS40 adds a field and reads 40 bit signed integer in current endian
func (d *D) FieldS40(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S40, sfns...)
}

// Reader S41

// TryS41 tries to read 41 bit signed integer in current endian
func (d *D) TryS41() (int64, error) { return d.trySE(41, d.Endian) }

func (d *D) ScalarS41() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(41, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S41 reads 41 bit signed integer in current endian
func (d *D) S41() int64 {
	v, err := d.trySE(41, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S41", Pos: d.Pos()})
	}
	return v
}

// TryFieldS41 tries to add a field and read 41 bit signed integer in current endian
func (d *D) TryFieldS41(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS41, sfns...)
}

// FieldS41 adds a field and reads 41 bit signed integer in current endian
func (d *D) FieldS41(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S41, sfns...)
}

// Reader S42

// TryS42 tries to read 42 bit signed integer in current endian
func (d *D) TryS42() (int64, error) { return d.trySE(42, d.Endian) }

func (d *D) ScalarS42() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(42, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S42 reads 42 bit signed integer in current endian
func (d *D) S42() int64 {
	v, err := d.trySE(42, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S42", Pos: d.Pos()})
	}
	return v
}

// TryFieldS42 tries to add a field and read 42 bit signed integer in current endian
func (d *D) TryFieldS42(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS42, sfns...)
}

// FieldS42 adds a field and reads 42 bit signed integer in current endian
func (d *D) FieldS42(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S42, sfns...)
}

// Reader S43

// TryS43 tries to read 43 bit signed integer in current endian
func (d *D) TryS43() (int64, error) { return d.trySE(43, d.Endian) }

func (d *D) ScalarS43() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(43, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S43 reads 43 bit signed integer in current endian
func (d *D) S43() int64 {
	v, err := d.trySE(43, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S43", Pos: d.Pos()})
	}
	return v
}

// TryFieldS43 tries to add a field and read 43 bit signed integer in current endian
func (d *D) TryFieldS43(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS43, sfns...)
}

// FieldS43 adds a field and reads 43 bit signed integer in current endian
func (d *D) FieldS43(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S43, sfns...)
}

// Reader S44

// TryS44 tries to read 44 bit signed integer in current endian
func (d *D) TryS44() (int64, error) { return d.trySE(44, d.Endian) }

func (d *D) ScalarS44() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(44, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S44 reads 44 bit signed integer in current endian
func (d *D) S44() int64 {
	v, err := d.trySE(44, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S44", Pos: d.Pos()})
	}
	return v
}

// TryFieldS44 tries to add a field and read 44 bit signed integer in current endian
func (d *D) TryFieldS44(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS44, sfns...)
}

// FieldS44 adds a field and reads 44 bit signed integer in current endian
func (d *D) FieldS44(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S44, sfns...)
}

// Reader S45

// TryS45 tries to read 45 bit signed integer in current endian
func (d *D) TryS45() (int64, error) { return d.trySE(45, d.Endian) }

func (d *D) ScalarS45() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(45, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S45 reads 45 bit signed integer in current endian
func (d *D) S45() int64 {
	v, err := d.trySE(45, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S45", Pos: d.Pos()})
	}
	return v
}

// TryFieldS45 tries to add a field and read 45 bit signed integer in current endian
func (d *D) TryFieldS45(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS45, sfns...)
}

// FieldS45 adds a field and reads 45 bit signed integer in current endian
func (d *D) FieldS45(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S45, sfns...)
}

// Reader S46

// TryS46 tries to read 46 bit signed integer in current endian
func (d *D) TryS46() (int64, error) { return d.trySE(46, d.Endian) }

func (d *D) ScalarS46() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(46, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S46 reads 46 bit signed integer in current endian
func (d *D) S46() int64 {
	v, err := d.trySE(46, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S46", Pos: d.Pos()})
	}
	return v
}

// TryFieldS46 tries to add a field and read 46 bit signed integer in current endian
func (d *D) TryFieldS46(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS46, sfns...)
}

// FieldS46 adds a field and reads 46 bit signed integer in current endian
func (d *D) FieldS46(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S46, sfns...)
}

// Reader S47

// TryS47 tries to read 47 bit signed integer in current endian
func (d *D) TryS47() (int64, error) { return d.trySE(47, d.Endian) }

func (d *D) ScalarS47() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(47, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S47 reads 47 bit signed integer in current endian
func (d *D) S47() int64 {
	v, err := d.trySE(47, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S47", Pos: d.Pos()})
	}
	return v
}

// TryFieldS47 tries to add a field and read 47 bit signed integer in current endian
func (d *D) TryFieldS47(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS47, sfns...)
}

// FieldS47 adds a field and reads 47 bit signed integer in current endian
func (d *D) FieldS47(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S47, sfns...)
}

// Reader S48

// TryS48 tries to read 48 bit signed integer in current endian
func (d *D) TryS48() (int64, error) { return d.trySE(48, d.Endian) }

func (d *D) ScalarS48() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(48, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S48 reads 48 bit signed integer in current endian
func (d *D) S48() int64 {
	v, err := d.trySE(48, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S48", Pos: d.Pos()})
	}
	return v
}

// TryFieldS48 tries to add a field and read 48 bit signed integer in current endian
func (d *D) TryFieldS48(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS48, sfns...)
}

// FieldS48 adds a field and reads 48 bit signed integer in current endian
func (d *D) FieldS48(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S48, sfns...)
}

// Reader S49

// TryS49 tries to read 49 bit signed integer in current endian
func (d *D) TryS49() (int64, error) { return d.trySE(49, d.Endian) }

func (d *D) ScalarS49() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(49, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S49 reads 49 bit signed integer in current endian
func (d *D) S49() int64 {
	v, err := d.trySE(49, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S49", Pos: d.Pos()})
	}
	return v
}

// TryFieldS49 tries to add a field and read 49 bit signed integer in current endian
func (d *D) TryFieldS49(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS49, sfns...)
}

// FieldS49 adds a field and reads 49 bit signed integer in current endian
func (d *D) FieldS49(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S49, sfns...)
}

// Reader S50

// TryS50 tries to read 50 bit signed integer in current endian
func (d *D) TryS50() (int64, error) { return d.trySE(50, d.Endian) }

func (d *D) ScalarS50() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(50, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S50 reads 50 bit signed integer in current endian
func (d *D) S50() int64 {
	v, err := d.trySE(50, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S50", Pos: d.Pos()})
	}
	return v
}

// TryFieldS50 tries to add a field and read 50 bit signed integer in current endian
func (d *D) TryFieldS50(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS50, sfns...)
}

// FieldS50 adds a field and reads 50 bit signed integer in current endian
func (d *D) FieldS50(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S50, sfns...)
}

// Reader S51

// TryS51 tries to read 51 bit signed integer in current endian
func (d *D) TryS51() (int64, error) { return d.trySE(51, d.Endian) }

func (d *D) ScalarS51() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(51, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S51 reads 51 bit signed integer in current endian
func (d *D) S51() int64 {
	v, err := d.trySE(51, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S51", Pos: d.Pos()})
	}
	return v
}

// TryFieldS51 tries to add a field and read 51 bit signed integer in current endian
func (d *D) TryFieldS51(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS51, sfns...)
}

// FieldS51 adds a field and reads 51 bit signed integer in current endian
func (d *D) FieldS51(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S51, sfns...)
}

// Reader S52

// TryS52 tries to read 52 bit signed integer in current endian
func (d *D) TryS52() (int64, error) { return d.trySE(52, d.Endian) }

func (d *D) ScalarS52() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(52, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S52 reads 52 bit signed integer in current endian
func (d *D) S52() int64 {
	v, err := d.trySE(52, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S52", Pos: d.Pos()})
	}
	return v
}

// TryFieldS52 tries to add a field and read 52 bit signed integer in current endian
func (d *D) TryFieldS52(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS52, sfns...)
}

// FieldS52 adds a field and reads 52 bit signed integer in current endian
func (d *D) FieldS52(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S52, sfns...)
}

// Reader S53

// TryS53 tries to read 53 bit signed integer in current endian
func (d *D) TryS53() (int64, error) { return d.trySE(53, d.Endian) }

func (d *D) ScalarS53() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(53, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S53 reads 53 bit signed integer in current endian
func (d *D) S53() int64 {
	v, err := d.trySE(53, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S53", Pos: d.Pos()})
	}
	return v
}

// TryFieldS53 tries to add a field and read 53 bit signed integer in current endian
func (d *D) TryFieldS53(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS53, sfns...)
}

// FieldS53 adds a field and reads 53 bit signed integer in current endian
func (d *D) FieldS53(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S53, sfns...)
}

// Reader S54

// TryS54 tries to read 54 bit signed integer in current endian
func (d *D) TryS54() (int64, error) { return d.trySE(54, d.Endian) }

func (d *D) ScalarS54() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(54, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S54 reads 54 bit signed integer in current endian
func (d *D) S54() int64 {
	v, err := d.trySE(54, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S54", Pos: d.Pos()})
	}
	return v
}

// TryFieldS54 tries to add a field and read 54 bit signed integer in current endian
func (d *D) TryFieldS54(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS54, sfns...)
}

// FieldS54 adds a field and reads 54 bit signed integer in current endian
func (d *D) FieldS54(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S54, sfns...)
}

// Reader S55

// TryS55 tries to read 55 bit signed integer in current endian
func (d *D) TryS55() (int64, error) { return d.trySE(55, d.Endian) }

func (d *D) ScalarS55() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(55, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S55 reads 55 bit signed integer in current endian
func (d *D) S55() int64 {
	v, err := d.trySE(55, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S55", Pos: d.Pos()})
	}
	return v
}

// TryFieldS55 tries to add a field and read 55 bit signed integer in current endian
func (d *D) TryFieldS55(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS55, sfns...)
}

// FieldS55 adds a field and reads 55 bit signed integer in current endian
func (d *D) FieldS55(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S55, sfns...)
}

// Reader S56

// TryS56 tries to read 56 bit signed integer in current endian
func (d *D) TryS56() (int64, error) { return d.trySE(56, d.Endian) }

func (d *D) ScalarS56() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(56, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S56 reads 56 bit signed integer in current endian
func (d *D) S56() int64 {
	v, err := d.trySE(56, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S56", Pos: d.Pos()})
	}
	return v
}

// TryFieldS56 tries to add a field and read 56 bit signed integer in current endian
func (d *D) TryFieldS56(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS56, sfns...)
}

// FieldS56 adds a field and reads 56 bit signed integer in current endian
func (d *D) FieldS56(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S56, sfns...)
}

// Reader S57

// TryS57 tries to read 57 bit signed integer in current endian
func (d *D) TryS57() (int64, error) { return d.trySE(57, d.Endian) }

func (d *D) ScalarS57() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(57, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S57 reads 57 bit signed integer in current endian
func (d *D) S57() int64 {
	v, err := d.trySE(57, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S57", Pos: d.Pos()})
	}
	return v
}

// TryFieldS57 tries to add a field and read 57 bit signed integer in current endian
func (d *D) TryFieldS57(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS57, sfns...)
}

// FieldS57 adds a field and reads 57 bit signed integer in current endian
func (d *D) FieldS57(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S57, sfns...)
}

// Reader S58

// TryS58 tries to read 58 bit signed integer in current endian
func (d *D) TryS58() (int64, error) { return d.trySE(58, d.Endian) }

func (d *D) ScalarS58() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(58, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S58 reads 58 bit signed integer in current endian
func (d *D) S58() int64 {
	v, err := d.trySE(58, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S58", Pos: d.Pos()})
	}
	return v
}

// TryFieldS58 tries to add a field and read 58 bit signed integer in current endian
func (d *D) TryFieldS58(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS58, sfns...)
}

// FieldS58 adds a field and reads 58 bit signed integer in current endian
func (d *D) FieldS58(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S58, sfns...)
}

// Reader S59

// TryS59 tries to read 59 bit signed integer in current endian
func (d *D) TryS59() (int64, error) { return d.trySE(59, d.Endian) }

func (d *D) ScalarS59() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(59, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S59 reads 59 bit signed integer in current endian
func (d *D) S59() int64 {
	v, err := d.trySE(59, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S59", Pos: d.Pos()})
	}
	return v
}

// TryFieldS59 tries to add a field and read 59 bit signed integer in current endian
func (d *D) TryFieldS59(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS59, sfns...)
}

// FieldS59 adds a field and reads 59 bit signed integer in current endian
func (d *D) FieldS59(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S59, sfns...)
}

// Reader S60

// TryS60 tries to read 60 bit signed integer in current endian
func (d *D) TryS60() (int64, error) { return d.trySE(60, d.Endian) }

func (d *D) ScalarS60() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(60, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S60 reads 60 bit signed integer in current endian
func (d *D) S60() int64 {
	v, err := d.trySE(60, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S60", Pos: d.Pos()})
	}
	return v
}

// TryFieldS60 tries to add a field and read 60 bit signed integer in current endian
func (d *D) TryFieldS60(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS60, sfns...)
}

// FieldS60 adds a field and reads 60 bit signed integer in current endian
func (d *D) FieldS60(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S60, sfns...)
}

// Reader S61

// TryS61 tries to read 61 bit signed integer in current endian
func (d *D) TryS61() (int64, error) { return d.trySE(61, d.Endian) }

func (d *D) ScalarS61() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(61, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S61 reads 61 bit signed integer in current endian
func (d *D) S61() int64 {
	v, err := d.trySE(61, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S61", Pos: d.Pos()})
	}
	return v
}

// TryFieldS61 tries to add a field and read 61 bit signed integer in current endian
func (d *D) TryFieldS61(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS61, sfns...)
}

// FieldS61 adds a field and reads 61 bit signed integer in current endian
func (d *D) FieldS61(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S61, sfns...)
}

// Reader S62

// TryS62 tries to read 62 bit signed integer in current endian
func (d *D) TryS62() (int64, error) { return d.trySE(62, d.Endian) }

func (d *D) ScalarS62() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(62, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S62 reads 62 bit signed integer in current endian
func (d *D) S62() int64 {
	v, err := d.trySE(62, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S62", Pos: d.Pos()})
	}
	return v
}

// TryFieldS62 tries to add a field and read 62 bit signed integer in current endian
func (d *D) TryFieldS62(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS62, sfns...)
}

// FieldS62 adds a field and reads 62 bit signed integer in current endian
func (d *D) FieldS62(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S62, sfns...)
}

// Reader S63

// TryS63 tries to read 63 bit signed integer in current endian
func (d *D) TryS63() (int64, error) { return d.trySE(63, d.Endian) }

func (d *D) ScalarS63() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(63, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S63 reads 63 bit signed integer in current endian
func (d *D) S63() int64 {
	v, err := d.trySE(63, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S63", Pos: d.Pos()})
	}
	return v
}

// TryFieldS63 tries to add a field and read 63 bit signed integer in current endian
func (d *D) TryFieldS63(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS63, sfns...)
}

// FieldS63 adds a field and reads 63 bit signed integer in current endian
func (d *D) FieldS63(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S63, sfns...)
}

// Reader S64

// TryS64 tries to read 64 bit signed integer in current endian
func (d *D) TryS64() (int64, error) { return d.trySE(64, d.Endian) }

func (d *D) ScalarS64() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(64, d.Endian)
		s.Actual = v
		return s, err
	}
}

// S64 reads 64 bit signed integer in current endian
func (d *D) S64() int64 {
	v, err := d.trySE(64, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "S64", Pos: d.Pos()})
	}
	return v
}

// TryFieldS64 tries to add a field and read 64 bit signed integer in current endian
func (d *D) TryFieldS64(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS64, sfns...)
}

// FieldS64 adds a field and reads 64 bit signed integer in current endian
func (d *D) FieldS64(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S64, sfns...)
}

// Reader S8LE

// TryS8LE tries to read 8 bit signed integer in little-endian
func (d *D) TryS8LE() (int64, error) { return d.trySE(8, LittleEndian) }

func (d *D) ScalarS8LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(8, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S8LE reads 8 bit signed integer in little-endian
func (d *D) S8LE() int64 {
	v, err := d.trySE(8, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S8LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS8LE tries to add a field and read 8 bit signed integer in little-endian
func (d *D) TryFieldS8LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS8LE, sfns...)
}

// FieldS8LE adds a field and reads 8 bit signed integer in little-endian
func (d *D) FieldS8LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S8LE, sfns...)
}

// Reader S9LE

// TryS9LE tries to read 9 bit signed integer in little-endian
func (d *D) TryS9LE() (int64, error) { return d.trySE(9, LittleEndian) }

func (d *D) ScalarS9LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(9, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S9LE reads 9 bit signed integer in little-endian
func (d *D) S9LE() int64 {
	v, err := d.trySE(9, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S9LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS9LE tries to add a field and read 9 bit signed integer in little-endian
func (d *D) TryFieldS9LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS9LE, sfns...)
}

// FieldS9LE adds a field and reads 9 bit signed integer in little-endian
func (d *D) FieldS9LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S9LE, sfns...)
}

// Reader S10LE

// TryS10LE tries to read 10 bit signed integer in little-endian
func (d *D) TryS10LE() (int64, error) { return d.trySE(10, LittleEndian) }

func (d *D) ScalarS10LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(10, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S10LE reads 10 bit signed integer in little-endian
func (d *D) S10LE() int64 {
	v, err := d.trySE(10, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S10LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS10LE tries to add a field and read 10 bit signed integer in little-endian
func (d *D) TryFieldS10LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS10LE, sfns...)
}

// FieldS10LE adds a field and reads 10 bit signed integer in little-endian
func (d *D) FieldS10LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S10LE, sfns...)
}

// Reader S11LE

// TryS11LE tries to read 11 bit signed integer in little-endian
func (d *D) TryS11LE() (int64, error) { return d.trySE(11, LittleEndian) }

func (d *D) ScalarS11LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(11, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S11LE reads 11 bit signed integer in little-endian
func (d *D) S11LE() int64 {
	v, err := d.trySE(11, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S11LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS11LE tries to add a field and read 11 bit signed integer in little-endian
func (d *D) TryFieldS11LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS11LE, sfns...)
}

// FieldS11LE adds a field and reads 11 bit signed integer in little-endian
func (d *D) FieldS11LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S11LE, sfns...)
}

// Reader S12LE

// TryS12LE tries to read 12 bit signed integer in little-endian
func (d *D) TryS12LE() (int64, error) { return d.trySE(12, LittleEndian) }

func (d *D) ScalarS12LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(12, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S12LE reads 12 bit signed integer in little-endian
func (d *D) S12LE() int64 {
	v, err := d.trySE(12, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S12LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS12LE tries to add a field and read 12 bit signed integer in little-endian
func (d *D) TryFieldS12LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS12LE, sfns...)
}

// FieldS12LE adds a field and reads 12 bit signed integer in little-endian
func (d *D) FieldS12LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S12LE, sfns...)
}

// Reader S13LE

// TryS13LE tries to read 13 bit signed integer in little-endian
func (d *D) TryS13LE() (int64, error) { return d.trySE(13, LittleEndian) }

func (d *D) ScalarS13LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(13, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S13LE reads 13 bit signed integer in little-endian
func (d *D) S13LE() int64 {
	v, err := d.trySE(13, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S13LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS13LE tries to add a field and read 13 bit signed integer in little-endian
func (d *D) TryFieldS13LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS13LE, sfns...)
}

// FieldS13LE adds a field and reads 13 bit signed integer in little-endian
func (d *D) FieldS13LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S13LE, sfns...)
}

// Reader S14LE

// TryS14LE tries to read 14 bit signed integer in little-endian
func (d *D) TryS14LE() (int64, error) { return d.trySE(14, LittleEndian) }

func (d *D) ScalarS14LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(14, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S14LE reads 14 bit signed integer in little-endian
func (d *D) S14LE() int64 {
	v, err := d.trySE(14, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S14LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS14LE tries to add a field and read 14 bit signed integer in little-endian
func (d *D) TryFieldS14LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS14LE, sfns...)
}

// FieldS14LE adds a field and reads 14 bit signed integer in little-endian
func (d *D) FieldS14LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S14LE, sfns...)
}

// Reader S15LE

// TryS15LE tries to read 15 bit signed integer in little-endian
func (d *D) TryS15LE() (int64, error) { return d.trySE(15, LittleEndian) }

func (d *D) ScalarS15LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(15, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S15LE reads 15 bit signed integer in little-endian
func (d *D) S15LE() int64 {
	v, err := d.trySE(15, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S15LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS15LE tries to add a field and read 15 bit signed integer in little-endian
func (d *D) TryFieldS15LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS15LE, sfns...)
}

// FieldS15LE adds a field and reads 15 bit signed integer in little-endian
func (d *D) FieldS15LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S15LE, sfns...)
}

// Reader S16LE

// TryS16LE tries to read 16 bit signed integer in little-endian
func (d *D) TryS16LE() (int64, error) { return d.trySE(16, LittleEndian) }

func (d *D) ScalarS16LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(16, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S16LE reads 16 bit signed integer in little-endian
func (d *D) S16LE() int64 {
	v, err := d.trySE(16, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S16LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS16LE tries to add a field and read 16 bit signed integer in little-endian
func (d *D) TryFieldS16LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS16LE, sfns...)
}

// FieldS16LE adds a field and reads 16 bit signed integer in little-endian
func (d *D) FieldS16LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S16LE, sfns...)
}

// Reader S17LE

// TryS17LE tries to read 17 bit signed integer in little-endian
func (d *D) TryS17LE() (int64, error) { return d.trySE(17, LittleEndian) }

func (d *D) ScalarS17LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(17, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S17LE reads 17 bit signed integer in little-endian
func (d *D) S17LE() int64 {
	v, err := d.trySE(17, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S17LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS17LE tries to add a field and read 17 bit signed integer in little-endian
func (d *D) TryFieldS17LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS17LE, sfns...)
}

// FieldS17LE adds a field and reads 17 bit signed integer in little-endian
func (d *D) FieldS17LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S17LE, sfns...)
}

// Reader S18LE

// TryS18LE tries to read 18 bit signed integer in little-endian
func (d *D) TryS18LE() (int64, error) { return d.trySE(18, LittleEndian) }

func (d *D) ScalarS18LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(18, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S18LE reads 18 bit signed integer in little-endian
func (d *D) S18LE() int64 {
	v, err := d.trySE(18, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S18LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS18LE tries to add a field and read 18 bit signed integer in little-endian
func (d *D) TryFieldS18LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS18LE, sfns...)
}

// FieldS18LE adds a field and reads 18 bit signed integer in little-endian
func (d *D) FieldS18LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S18LE, sfns...)
}

// Reader S19LE

// TryS19LE tries to read 19 bit signed integer in little-endian
func (d *D) TryS19LE() (int64, error) { return d.trySE(19, LittleEndian) }

func (d *D) ScalarS19LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(19, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S19LE reads 19 bit signed integer in little-endian
func (d *D) S19LE() int64 {
	v, err := d.trySE(19, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S19LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS19LE tries to add a field and read 19 bit signed integer in little-endian
func (d *D) TryFieldS19LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS19LE, sfns...)
}

// FieldS19LE adds a field and reads 19 bit signed integer in little-endian
func (d *D) FieldS19LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S19LE, sfns...)
}

// Reader S20LE

// TryS20LE tries to read 20 bit signed integer in little-endian
func (d *D) TryS20LE() (int64, error) { return d.trySE(20, LittleEndian) }

func (d *D) ScalarS20LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(20, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S20LE reads 20 bit signed integer in little-endian
func (d *D) S20LE() int64 {
	v, err := d.trySE(20, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S20LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS20LE tries to add a field and read 20 bit signed integer in little-endian
func (d *D) TryFieldS20LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS20LE, sfns...)
}

// FieldS20LE adds a field and reads 20 bit signed integer in little-endian
func (d *D) FieldS20LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S20LE, sfns...)
}

// Reader S21LE

// TryS21LE tries to read 21 bit signed integer in little-endian
func (d *D) TryS21LE() (int64, error) { return d.trySE(21, LittleEndian) }

func (d *D) ScalarS21LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(21, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S21LE reads 21 bit signed integer in little-endian
func (d *D) S21LE() int64 {
	v, err := d.trySE(21, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S21LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS21LE tries to add a field and read 21 bit signed integer in little-endian
func (d *D) TryFieldS21LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS21LE, sfns...)
}

// FieldS21LE adds a field and reads 21 bit signed integer in little-endian
func (d *D) FieldS21LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S21LE, sfns...)
}

// Reader S22LE

// TryS22LE tries to read 22 bit signed integer in little-endian
func (d *D) TryS22LE() (int64, error) { return d.trySE(22, LittleEndian) }

func (d *D) ScalarS22LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(22, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S22LE reads 22 bit signed integer in little-endian
func (d *D) S22LE() int64 {
	v, err := d.trySE(22, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S22LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS22LE tries to add a field and read 22 bit signed integer in little-endian
func (d *D) TryFieldS22LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS22LE, sfns...)
}

// FieldS22LE adds a field and reads 22 bit signed integer in little-endian
func (d *D) FieldS22LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S22LE, sfns...)
}

// Reader S23LE

// TryS23LE tries to read 23 bit signed integer in little-endian
func (d *D) TryS23LE() (int64, error) { return d.trySE(23, LittleEndian) }

func (d *D) ScalarS23LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(23, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S23LE reads 23 bit signed integer in little-endian
func (d *D) S23LE() int64 {
	v, err := d.trySE(23, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S23LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS23LE tries to add a field and read 23 bit signed integer in little-endian
func (d *D) TryFieldS23LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS23LE, sfns...)
}

// FieldS23LE adds a field and reads 23 bit signed integer in little-endian
func (d *D) FieldS23LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S23LE, sfns...)
}

// Reader S24LE

// TryS24LE tries to read 24 bit signed integer in little-endian
func (d *D) TryS24LE() (int64, error) { return d.trySE(24, LittleEndian) }

func (d *D) ScalarS24LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(24, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S24LE reads 24 bit signed integer in little-endian
func (d *D) S24LE() int64 {
	v, err := d.trySE(24, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S24LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS24LE tries to add a field and read 24 bit signed integer in little-endian
func (d *D) TryFieldS24LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS24LE, sfns...)
}

// FieldS24LE adds a field and reads 24 bit signed integer in little-endian
func (d *D) FieldS24LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S24LE, sfns...)
}

// Reader S25LE

// TryS25LE tries to read 25 bit signed integer in little-endian
func (d *D) TryS25LE() (int64, error) { return d.trySE(25, LittleEndian) }

func (d *D) ScalarS25LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(25, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S25LE reads 25 bit signed integer in little-endian
func (d *D) S25LE() int64 {
	v, err := d.trySE(25, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S25LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS25LE tries to add a field and read 25 bit signed integer in little-endian
func (d *D) TryFieldS25LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS25LE, sfns...)
}

// FieldS25LE adds a field and reads 25 bit signed integer in little-endian
func (d *D) FieldS25LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S25LE, sfns...)
}

// Reader S26LE

// TryS26LE tries to read 26 bit signed integer in little-endian
func (d *D) TryS26LE() (int64, error) { return d.trySE(26, LittleEndian) }

func (d *D) ScalarS26LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(26, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S26LE reads 26 bit signed integer in little-endian
func (d *D) S26LE() int64 {
	v, err := d.trySE(26, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S26LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS26LE tries to add a field and read 26 bit signed integer in little-endian
func (d *D) TryFieldS26LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS26LE, sfns...)
}

// FieldS26LE adds a field and reads 26 bit signed integer in little-endian
func (d *D) FieldS26LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S26LE, sfns...)
}

// Reader S27LE

// TryS27LE tries to read 27 bit signed integer in little-endian
func (d *D) TryS27LE() (int64, error) { return d.trySE(27, LittleEndian) }

func (d *D) ScalarS27LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(27, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S27LE reads 27 bit signed integer in little-endian
func (d *D) S27LE() int64 {
	v, err := d.trySE(27, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S27LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS27LE tries to add a field and read 27 bit signed integer in little-endian
func (d *D) TryFieldS27LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS27LE, sfns...)
}

// FieldS27LE adds a field and reads 27 bit signed integer in little-endian
func (d *D) FieldS27LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S27LE, sfns...)
}

// Reader S28LE

// TryS28LE tries to read 28 bit signed integer in little-endian
func (d *D) TryS28LE() (int64, error) { return d.trySE(28, LittleEndian) }

func (d *D) ScalarS28LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(28, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S28LE reads 28 bit signed integer in little-endian
func (d *D) S28LE() int64 {
	v, err := d.trySE(28, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S28LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS28LE tries to add a field and read 28 bit signed integer in little-endian
func (d *D) TryFieldS28LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS28LE, sfns...)
}

// FieldS28LE adds a field and reads 28 bit signed integer in little-endian
func (d *D) FieldS28LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S28LE, sfns...)
}

// Reader S29LE

// TryS29LE tries to read 29 bit signed integer in little-endian
func (d *D) TryS29LE() (int64, error) { return d.trySE(29, LittleEndian) }

func (d *D) ScalarS29LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(29, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S29LE reads 29 bit signed integer in little-endian
func (d *D) S29LE() int64 {
	v, err := d.trySE(29, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S29LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS29LE tries to add a field and read 29 bit signed integer in little-endian
func (d *D) TryFieldS29LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS29LE, sfns...)
}

// FieldS29LE adds a field and reads 29 bit signed integer in little-endian
func (d *D) FieldS29LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S29LE, sfns...)
}

// Reader S30LE

// TryS30LE tries to read 30 bit signed integer in little-endian
func (d *D) TryS30LE() (int64, error) { return d.trySE(30, LittleEndian) }

func (d *D) ScalarS30LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(30, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S30LE reads 30 bit signed integer in little-endian
func (d *D) S30LE() int64 {
	v, err := d.trySE(30, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S30LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS30LE tries to add a field and read 30 bit signed integer in little-endian
func (d *D) TryFieldS30LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS30LE, sfns...)
}

// FieldS30LE adds a field and reads 30 bit signed integer in little-endian
func (d *D) FieldS30LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S30LE, sfns...)
}

// Reader S31LE

// TryS31LE tries to read 31 bit signed integer in little-endian
func (d *D) TryS31LE() (int64, error) { return d.trySE(31, LittleEndian) }

func (d *D) ScalarS31LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(31, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S31LE reads 31 bit signed integer in little-endian
func (d *D) S31LE() int64 {
	v, err := d.trySE(31, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S31LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS31LE tries to add a field and read 31 bit signed integer in little-endian
func (d *D) TryFieldS31LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS31LE, sfns...)
}

// FieldS31LE adds a field and reads 31 bit signed integer in little-endian
func (d *D) FieldS31LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S31LE, sfns...)
}

// Reader S32LE

// TryS32LE tries to read 32 bit signed integer in little-endian
func (d *D) TryS32LE() (int64, error) { return d.trySE(32, LittleEndian) }

func (d *D) ScalarS32LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(32, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S32LE reads 32 bit signed integer in little-endian
func (d *D) S32LE() int64 {
	v, err := d.trySE(32, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S32LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS32LE tries to add a field and read 32 bit signed integer in little-endian
func (d *D) TryFieldS32LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS32LE, sfns...)
}

// FieldS32LE adds a field and reads 32 bit signed integer in little-endian
func (d *D) FieldS32LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S32LE, sfns...)
}

// Reader S33LE

// TryS33LE tries to read 33 bit signed integer in little-endian
func (d *D) TryS33LE() (int64, error) { return d.trySE(33, LittleEndian) }

func (d *D) ScalarS33LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(33, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S33LE reads 33 bit signed integer in little-endian
func (d *D) S33LE() int64 {
	v, err := d.trySE(33, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S33LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS33LE tries to add a field and read 33 bit signed integer in little-endian
func (d *D) TryFieldS33LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS33LE, sfns...)
}

// FieldS33LE adds a field and reads 33 bit signed integer in little-endian
func (d *D) FieldS33LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S33LE, sfns...)
}

// Reader S34LE

// TryS34LE tries to read 34 bit signed integer in little-endian
func (d *D) TryS34LE() (int64, error) { return d.trySE(34, LittleEndian) }

func (d *D) ScalarS34LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(34, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S34LE reads 34 bit signed integer in little-endian
func (d *D) S34LE() int64 {
	v, err := d.trySE(34, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S34LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS34LE tries to add a field and read 34 bit signed integer in little-endian
func (d *D) TryFieldS34LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS34LE, sfns...)
}

// FieldS34LE adds a field and reads 34 bit signed integer in little-endian
func (d *D) FieldS34LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S34LE, sfns...)
}

// Reader S35LE

// TryS35LE tries to read 35 bit signed integer in little-endian
func (d *D) TryS35LE() (int64, error) { return d.trySE(35, LittleEndian) }

func (d *D) ScalarS35LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(35, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S35LE reads 35 bit signed integer in little-endian
func (d *D) S35LE() int64 {
	v, err := d.trySE(35, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S35LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS35LE tries to add a field and read 35 bit signed integer in little-endian
func (d *D) TryFieldS35LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS35LE, sfns...)
}

// FieldS35LE adds a field and reads 35 bit signed integer in little-endian
func (d *D) FieldS35LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S35LE, sfns...)
}

// Reader S36LE

// TryS36LE tries to read 36 bit signed integer in little-endian
func (d *D) TryS36LE() (int64, error) { return d.trySE(36, LittleEndian) }

func (d *D) ScalarS36LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(36, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S36LE reads 36 bit signed integer in little-endian
func (d *D) S36LE() int64 {
	v, err := d.trySE(36, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S36LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS36LE tries to add a field and read 36 bit signed integer in little-endian
func (d *D) TryFieldS36LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS36LE, sfns...)
}

// FieldS36LE adds a field and reads 36 bit signed integer in little-endian
func (d *D) FieldS36LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S36LE, sfns...)
}

// Reader S37LE

// TryS37LE tries to read 37 bit signed integer in little-endian
func (d *D) TryS37LE() (int64, error) { return d.trySE(37, LittleEndian) }

func (d *D) ScalarS37LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(37, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S37LE reads 37 bit signed integer in little-endian
func (d *D) S37LE() int64 {
	v, err := d.trySE(37, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S37LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS37LE tries to add a field and read 37 bit signed integer in little-endian
func (d *D) TryFieldS37LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS37LE, sfns...)
}

// FieldS37LE adds a field and reads 37 bit signed integer in little-endian
func (d *D) FieldS37LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S37LE, sfns...)
}

// Reader S38LE

// TryS38LE tries to read 38 bit signed integer in little-endian
func (d *D) TryS38LE() (int64, error) { return d.trySE(38, LittleEndian) }

func (d *D) ScalarS38LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(38, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S38LE reads 38 bit signed integer in little-endian
func (d *D) S38LE() int64 {
	v, err := d.trySE(38, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S38LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS38LE tries to add a field and read 38 bit signed integer in little-endian
func (d *D) TryFieldS38LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS38LE, sfns...)
}

// FieldS38LE adds a field and reads 38 bit signed integer in little-endian
func (d *D) FieldS38LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S38LE, sfns...)
}

// Reader S39LE

// TryS39LE tries to read 39 bit signed integer in little-endian
func (d *D) TryS39LE() (int64, error) { return d.trySE(39, LittleEndian) }

func (d *D) ScalarS39LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(39, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S39LE reads 39 bit signed integer in little-endian
func (d *D) S39LE() int64 {
	v, err := d.trySE(39, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S39LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS39LE tries to add a field and read 39 bit signed integer in little-endian
func (d *D) TryFieldS39LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS39LE, sfns...)
}

// FieldS39LE adds a field and reads 39 bit signed integer in little-endian
func (d *D) FieldS39LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S39LE, sfns...)
}

// Reader S40LE

// TryS40LE tries to read 40 bit signed integer in little-endian
func (d *D) TryS40LE() (int64, error) { return d.trySE(40, LittleEndian) }

func (d *D) ScalarS40LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(40, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S40LE reads 40 bit signed integer in little-endian
func (d *D) S40LE() int64 {
	v, err := d.trySE(40, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S40LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS40LE tries to add a field and read 40 bit signed integer in little-endian
func (d *D) TryFieldS40LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS40LE, sfns...)
}

// FieldS40LE adds a field and reads 40 bit signed integer in little-endian
func (d *D) FieldS40LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S40LE, sfns...)
}

// Reader S41LE

// TryS41LE tries to read 41 bit signed integer in little-endian
func (d *D) TryS41LE() (int64, error) { return d.trySE(41, LittleEndian) }

func (d *D) ScalarS41LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(41, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S41LE reads 41 bit signed integer in little-endian
func (d *D) S41LE() int64 {
	v, err := d.trySE(41, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S41LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS41LE tries to add a field and read 41 bit signed integer in little-endian
func (d *D) TryFieldS41LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS41LE, sfns...)
}

// FieldS41LE adds a field and reads 41 bit signed integer in little-endian
func (d *D) FieldS41LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S41LE, sfns...)
}

// Reader S42LE

// TryS42LE tries to read 42 bit signed integer in little-endian
func (d *D) TryS42LE() (int64, error) { return d.trySE(42, LittleEndian) }

func (d *D) ScalarS42LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(42, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S42LE reads 42 bit signed integer in little-endian
func (d *D) S42LE() int64 {
	v, err := d.trySE(42, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S42LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS42LE tries to add a field and read 42 bit signed integer in little-endian
func (d *D) TryFieldS42LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS42LE, sfns...)
}

// FieldS42LE adds a field and reads 42 bit signed integer in little-endian
func (d *D) FieldS42LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S42LE, sfns...)
}

// Reader S43LE

// TryS43LE tries to read 43 bit signed integer in little-endian
func (d *D) TryS43LE() (int64, error) { return d.trySE(43, LittleEndian) }

func (d *D) ScalarS43LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(43, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S43LE reads 43 bit signed integer in little-endian
func (d *D) S43LE() int64 {
	v, err := d.trySE(43, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S43LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS43LE tries to add a field and read 43 bit signed integer in little-endian
func (d *D) TryFieldS43LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS43LE, sfns...)
}

// FieldS43LE adds a field and reads 43 bit signed integer in little-endian
func (d *D) FieldS43LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S43LE, sfns...)
}

// Reader S44LE

// TryS44LE tries to read 44 bit signed integer in little-endian
func (d *D) TryS44LE() (int64, error) { return d.trySE(44, LittleEndian) }

func (d *D) ScalarS44LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(44, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S44LE reads 44 bit signed integer in little-endian
func (d *D) S44LE() int64 {
	v, err := d.trySE(44, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S44LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS44LE tries to add a field and read 44 bit signed integer in little-endian
func (d *D) TryFieldS44LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS44LE, sfns...)
}

// FieldS44LE adds a field and reads 44 bit signed integer in little-endian
func (d *D) FieldS44LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S44LE, sfns...)
}

// Reader S45LE

// TryS45LE tries to read 45 bit signed integer in little-endian
func (d *D) TryS45LE() (int64, error) { return d.trySE(45, LittleEndian) }

func (d *D) ScalarS45LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(45, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S45LE reads 45 bit signed integer in little-endian
func (d *D) S45LE() int64 {
	v, err := d.trySE(45, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S45LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS45LE tries to add a field and read 45 bit signed integer in little-endian
func (d *D) TryFieldS45LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS45LE, sfns...)
}

// FieldS45LE adds a field and reads 45 bit signed integer in little-endian
func (d *D) FieldS45LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S45LE, sfns...)
}

// Reader S46LE

// TryS46LE tries to read 46 bit signed integer in little-endian
func (d *D) TryS46LE() (int64, error) { return d.trySE(46, LittleEndian) }

func (d *D) ScalarS46LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(46, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S46LE reads 46 bit signed integer in little-endian
func (d *D) S46LE() int64 {
	v, err := d.trySE(46, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S46LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS46LE tries to add a field and read 46 bit signed integer in little-endian
func (d *D) TryFieldS46LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS46LE, sfns...)
}

// FieldS46LE adds a field and reads 46 bit signed integer in little-endian
func (d *D) FieldS46LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S46LE, sfns...)
}

// Reader S47LE

// TryS47LE tries to read 47 bit signed integer in little-endian
func (d *D) TryS47LE() (int64, error) { return d.trySE(47, LittleEndian) }

func (d *D) ScalarS47LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(47, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S47LE reads 47 bit signed integer in little-endian
func (d *D) S47LE() int64 {
	v, err := d.trySE(47, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S47LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS47LE tries to add a field and read 47 bit signed integer in little-endian
func (d *D) TryFieldS47LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS47LE, sfns...)
}

// FieldS47LE adds a field and reads 47 bit signed integer in little-endian
func (d *D) FieldS47LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S47LE, sfns...)
}

// Reader S48LE

// TryS48LE tries to read 48 bit signed integer in little-endian
func (d *D) TryS48LE() (int64, error) { return d.trySE(48, LittleEndian) }

func (d *D) ScalarS48LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(48, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S48LE reads 48 bit signed integer in little-endian
func (d *D) S48LE() int64 {
	v, err := d.trySE(48, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S48LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS48LE tries to add a field and read 48 bit signed integer in little-endian
func (d *D) TryFieldS48LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS48LE, sfns...)
}

// FieldS48LE adds a field and reads 48 bit signed integer in little-endian
func (d *D) FieldS48LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S48LE, sfns...)
}

// Reader S49LE

// TryS49LE tries to read 49 bit signed integer in little-endian
func (d *D) TryS49LE() (int64, error) { return d.trySE(49, LittleEndian) }

func (d *D) ScalarS49LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(49, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S49LE reads 49 bit signed integer in little-endian
func (d *D) S49LE() int64 {
	v, err := d.trySE(49, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S49LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS49LE tries to add a field and read 49 bit signed integer in little-endian
func (d *D) TryFieldS49LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS49LE, sfns...)
}

// FieldS49LE adds a field and reads 49 bit signed integer in little-endian
func (d *D) FieldS49LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S49LE, sfns...)
}

// Reader S50LE

// TryS50LE tries to read 50 bit signed integer in little-endian
func (d *D) TryS50LE() (int64, error) { return d.trySE(50, LittleEndian) }

func (d *D) ScalarS50LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(50, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S50LE reads 50 bit signed integer in little-endian
func (d *D) S50LE() int64 {
	v, err := d.trySE(50, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S50LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS50LE tries to add a field and read 50 bit signed integer in little-endian
func (d *D) TryFieldS50LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS50LE, sfns...)
}

// FieldS50LE adds a field and reads 50 bit signed integer in little-endian
func (d *D) FieldS50LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S50LE, sfns...)
}

// Reader S51LE

// TryS51LE tries to read 51 bit signed integer in little-endian
func (d *D) TryS51LE() (int64, error) { return d.trySE(51, LittleEndian) }

func (d *D) ScalarS51LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(51, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S51LE reads 51 bit signed integer in little-endian
func (d *D) S51LE() int64 {
	v, err := d.trySE(51, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S51LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS51LE tries to add a field and read 51 bit signed integer in little-endian
func (d *D) TryFieldS51LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS51LE, sfns...)
}

// FieldS51LE adds a field and reads 51 bit signed integer in little-endian
func (d *D) FieldS51LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S51LE, sfns...)
}

// Reader S52LE

// TryS52LE tries to read 52 bit signed integer in little-endian
func (d *D) TryS52LE() (int64, error) { return d.trySE(52, LittleEndian) }

func (d *D) ScalarS52LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(52, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S52LE reads 52 bit signed integer in little-endian
func (d *D) S52LE() int64 {
	v, err := d.trySE(52, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S52LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS52LE tries to add a field and read 52 bit signed integer in little-endian
func (d *D) TryFieldS52LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS52LE, sfns...)
}

// FieldS52LE adds a field and reads 52 bit signed integer in little-endian
func (d *D) FieldS52LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S52LE, sfns...)
}

// Reader S53LE

// TryS53LE tries to read 53 bit signed integer in little-endian
func (d *D) TryS53LE() (int64, error) { return d.trySE(53, LittleEndian) }

func (d *D) ScalarS53LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(53, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S53LE reads 53 bit signed integer in little-endian
func (d *D) S53LE() int64 {
	v, err := d.trySE(53, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S53LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS53LE tries to add a field and read 53 bit signed integer in little-endian
func (d *D) TryFieldS53LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS53LE, sfns...)
}

// FieldS53LE adds a field and reads 53 bit signed integer in little-endian
func (d *D) FieldS53LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S53LE, sfns...)
}

// Reader S54LE

// TryS54LE tries to read 54 bit signed integer in little-endian
func (d *D) TryS54LE() (int64, error) { return d.trySE(54, LittleEndian) }

func (d *D) ScalarS54LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(54, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S54LE reads 54 bit signed integer in little-endian
func (d *D) S54LE() int64 {
	v, err := d.trySE(54, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S54LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS54LE tries to add a field and read 54 bit signed integer in little-endian
func (d *D) TryFieldS54LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS54LE, sfns...)
}

// FieldS54LE adds a field and reads 54 bit signed integer in little-endian
func (d *D) FieldS54LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S54LE, sfns...)
}

// Reader S55LE

// TryS55LE tries to read 55 bit signed integer in little-endian
func (d *D) TryS55LE() (int64, error) { return d.trySE(55, LittleEndian) }

func (d *D) ScalarS55LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(55, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S55LE reads 55 bit signed integer in little-endian
func (d *D) S55LE() int64 {
	v, err := d.trySE(55, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S55LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS55LE tries to add a field and read 55 bit signed integer in little-endian
func (d *D) TryFieldS55LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS55LE, sfns...)
}

// FieldS55LE adds a field and reads 55 bit signed integer in little-endian
func (d *D) FieldS55LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S55LE, sfns...)
}

// Reader S56LE

// TryS56LE tries to read 56 bit signed integer in little-endian
func (d *D) TryS56LE() (int64, error) { return d.trySE(56, LittleEndian) }

func (d *D) ScalarS56LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(56, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S56LE reads 56 bit signed integer in little-endian
func (d *D) S56LE() int64 {
	v, err := d.trySE(56, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S56LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS56LE tries to add a field and read 56 bit signed integer in little-endian
func (d *D) TryFieldS56LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS56LE, sfns...)
}

// FieldS56LE adds a field and reads 56 bit signed integer in little-endian
func (d *D) FieldS56LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S56LE, sfns...)
}

// Reader S57LE

// TryS57LE tries to read 57 bit signed integer in little-endian
func (d *D) TryS57LE() (int64, error) { return d.trySE(57, LittleEndian) }

func (d *D) ScalarS57LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(57, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S57LE reads 57 bit signed integer in little-endian
func (d *D) S57LE() int64 {
	v, err := d.trySE(57, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S57LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS57LE tries to add a field and read 57 bit signed integer in little-endian
func (d *D) TryFieldS57LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS57LE, sfns...)
}

// FieldS57LE adds a field and reads 57 bit signed integer in little-endian
func (d *D) FieldS57LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S57LE, sfns...)
}

// Reader S58LE

// TryS58LE tries to read 58 bit signed integer in little-endian
func (d *D) TryS58LE() (int64, error) { return d.trySE(58, LittleEndian) }

func (d *D) ScalarS58LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(58, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S58LE reads 58 bit signed integer in little-endian
func (d *D) S58LE() int64 {
	v, err := d.trySE(58, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S58LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS58LE tries to add a field and read 58 bit signed integer in little-endian
func (d *D) TryFieldS58LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS58LE, sfns...)
}

// FieldS58LE adds a field and reads 58 bit signed integer in little-endian
func (d *D) FieldS58LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S58LE, sfns...)
}

// Reader S59LE

// TryS59LE tries to read 59 bit signed integer in little-endian
func (d *D) TryS59LE() (int64, error) { return d.trySE(59, LittleEndian) }

func (d *D) ScalarS59LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(59, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S59LE reads 59 bit signed integer in little-endian
func (d *D) S59LE() int64 {
	v, err := d.trySE(59, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S59LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS59LE tries to add a field and read 59 bit signed integer in little-endian
func (d *D) TryFieldS59LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS59LE, sfns...)
}

// FieldS59LE adds a field and reads 59 bit signed integer in little-endian
func (d *D) FieldS59LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S59LE, sfns...)
}

// Reader S60LE

// TryS60LE tries to read 60 bit signed integer in little-endian
func (d *D) TryS60LE() (int64, error) { return d.trySE(60, LittleEndian) }

func (d *D) ScalarS60LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(60, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S60LE reads 60 bit signed integer in little-endian
func (d *D) S60LE() int64 {
	v, err := d.trySE(60, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S60LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS60LE tries to add a field and read 60 bit signed integer in little-endian
func (d *D) TryFieldS60LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS60LE, sfns...)
}

// FieldS60LE adds a field and reads 60 bit signed integer in little-endian
func (d *D) FieldS60LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S60LE, sfns...)
}

// Reader S61LE

// TryS61LE tries to read 61 bit signed integer in little-endian
func (d *D) TryS61LE() (int64, error) { return d.trySE(61, LittleEndian) }

func (d *D) ScalarS61LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(61, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S61LE reads 61 bit signed integer in little-endian
func (d *D) S61LE() int64 {
	v, err := d.trySE(61, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S61LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS61LE tries to add a field and read 61 bit signed integer in little-endian
func (d *D) TryFieldS61LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS61LE, sfns...)
}

// FieldS61LE adds a field and reads 61 bit signed integer in little-endian
func (d *D) FieldS61LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S61LE, sfns...)
}

// Reader S62LE

// TryS62LE tries to read 62 bit signed integer in little-endian
func (d *D) TryS62LE() (int64, error) { return d.trySE(62, LittleEndian) }

func (d *D) ScalarS62LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(62, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S62LE reads 62 bit signed integer in little-endian
func (d *D) S62LE() int64 {
	v, err := d.trySE(62, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S62LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS62LE tries to add a field and read 62 bit signed integer in little-endian
func (d *D) TryFieldS62LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS62LE, sfns...)
}

// FieldS62LE adds a field and reads 62 bit signed integer in little-endian
func (d *D) FieldS62LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S62LE, sfns...)
}

// Reader S63LE

// TryS63LE tries to read 63 bit signed integer in little-endian
func (d *D) TryS63LE() (int64, error) { return d.trySE(63, LittleEndian) }

func (d *D) ScalarS63LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(63, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S63LE reads 63 bit signed integer in little-endian
func (d *D) S63LE() int64 {
	v, err := d.trySE(63, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S63LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS63LE tries to add a field and read 63 bit signed integer in little-endian
func (d *D) TryFieldS63LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS63LE, sfns...)
}

// FieldS63LE adds a field and reads 63 bit signed integer in little-endian
func (d *D) FieldS63LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S63LE, sfns...)
}

// Reader S64LE

// TryS64LE tries to read 64 bit signed integer in little-endian
func (d *D) TryS64LE() (int64, error) { return d.trySE(64, LittleEndian) }

func (d *D) ScalarS64LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(64, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// S64LE reads 64 bit signed integer in little-endian
func (d *D) S64LE() int64 {
	v, err := d.trySE(64, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S64LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS64LE tries to add a field and read 64 bit signed integer in little-endian
func (d *D) TryFieldS64LE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS64LE, sfns...)
}

// FieldS64LE adds a field and reads 64 bit signed integer in little-endian
func (d *D) FieldS64LE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S64LE, sfns...)
}

// Reader S8BE

// TryS8BE tries to read 8 bit signed integer in big-endian
func (d *D) TryS8BE() (int64, error) { return d.trySE(8, BigEndian) }

func (d *D) ScalarS8BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(8, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S8BE reads 8 bit signed integer in big-endian
func (d *D) S8BE() int64 {
	v, err := d.trySE(8, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S8BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS8BE tries to add a field and read 8 bit signed integer in big-endian
func (d *D) TryFieldS8BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS8BE, sfns...)
}

// FieldS8BE adds a field and reads 8 bit signed integer in big-endian
func (d *D) FieldS8BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S8BE, sfns...)
}

// Reader S9BE

// TryS9BE tries to read 9 bit signed integer in big-endian
func (d *D) TryS9BE() (int64, error) { return d.trySE(9, BigEndian) }

func (d *D) ScalarS9BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(9, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S9BE reads 9 bit signed integer in big-endian
func (d *D) S9BE() int64 {
	v, err := d.trySE(9, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S9BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS9BE tries to add a field and read 9 bit signed integer in big-endian
func (d *D) TryFieldS9BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS9BE, sfns...)
}

// FieldS9BE adds a field and reads 9 bit signed integer in big-endian
func (d *D) FieldS9BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S9BE, sfns...)
}

// Reader S10BE

// TryS10BE tries to read 10 bit signed integer in big-endian
func (d *D) TryS10BE() (int64, error) { return d.trySE(10, BigEndian) }

func (d *D) ScalarS10BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(10, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S10BE reads 10 bit signed integer in big-endian
func (d *D) S10BE() int64 {
	v, err := d.trySE(10, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S10BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS10BE tries to add a field and read 10 bit signed integer in big-endian
func (d *D) TryFieldS10BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS10BE, sfns...)
}

// FieldS10BE adds a field and reads 10 bit signed integer in big-endian
func (d *D) FieldS10BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S10BE, sfns...)
}

// Reader S11BE

// TryS11BE tries to read 11 bit signed integer in big-endian
func (d *D) TryS11BE() (int64, error) { return d.trySE(11, BigEndian) }

func (d *D) ScalarS11BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(11, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S11BE reads 11 bit signed integer in big-endian
func (d *D) S11BE() int64 {
	v, err := d.trySE(11, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S11BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS11BE tries to add a field and read 11 bit signed integer in big-endian
func (d *D) TryFieldS11BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS11BE, sfns...)
}

// FieldS11BE adds a field and reads 11 bit signed integer in big-endian
func (d *D) FieldS11BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S11BE, sfns...)
}

// Reader S12BE

// TryS12BE tries to read 12 bit signed integer in big-endian
func (d *D) TryS12BE() (int64, error) { return d.trySE(12, BigEndian) }

func (d *D) ScalarS12BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(12, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S12BE reads 12 bit signed integer in big-endian
func (d *D) S12BE() int64 {
	v, err := d.trySE(12, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S12BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS12BE tries to add a field and read 12 bit signed integer in big-endian
func (d *D) TryFieldS12BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS12BE, sfns...)
}

// FieldS12BE adds a field and reads 12 bit signed integer in big-endian
func (d *D) FieldS12BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S12BE, sfns...)
}

// Reader S13BE

// TryS13BE tries to read 13 bit signed integer in big-endian
func (d *D) TryS13BE() (int64, error) { return d.trySE(13, BigEndian) }

func (d *D) ScalarS13BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(13, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S13BE reads 13 bit signed integer in big-endian
func (d *D) S13BE() int64 {
	v, err := d.trySE(13, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S13BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS13BE tries to add a field and read 13 bit signed integer in big-endian
func (d *D) TryFieldS13BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS13BE, sfns...)
}

// FieldS13BE adds a field and reads 13 bit signed integer in big-endian
func (d *D) FieldS13BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S13BE, sfns...)
}

// Reader S14BE

// TryS14BE tries to read 14 bit signed integer in big-endian
func (d *D) TryS14BE() (int64, error) { return d.trySE(14, BigEndian) }

func (d *D) ScalarS14BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(14, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S14BE reads 14 bit signed integer in big-endian
func (d *D) S14BE() int64 {
	v, err := d.trySE(14, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S14BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS14BE tries to add a field and read 14 bit signed integer in big-endian
func (d *D) TryFieldS14BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS14BE, sfns...)
}

// FieldS14BE adds a field and reads 14 bit signed integer in big-endian
func (d *D) FieldS14BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S14BE, sfns...)
}

// Reader S15BE

// TryS15BE tries to read 15 bit signed integer in big-endian
func (d *D) TryS15BE() (int64, error) { return d.trySE(15, BigEndian) }

func (d *D) ScalarS15BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(15, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S15BE reads 15 bit signed integer in big-endian
func (d *D) S15BE() int64 {
	v, err := d.trySE(15, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S15BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS15BE tries to add a field and read 15 bit signed integer in big-endian
func (d *D) TryFieldS15BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS15BE, sfns...)
}

// FieldS15BE adds a field and reads 15 bit signed integer in big-endian
func (d *D) FieldS15BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S15BE, sfns...)
}

// Reader S16BE

// TryS16BE tries to read 16 bit signed integer in big-endian
func (d *D) TryS16BE() (int64, error) { return d.trySE(16, BigEndian) }

func (d *D) ScalarS16BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(16, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S16BE reads 16 bit signed integer in big-endian
func (d *D) S16BE() int64 {
	v, err := d.trySE(16, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S16BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS16BE tries to add a field and read 16 bit signed integer in big-endian
func (d *D) TryFieldS16BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS16BE, sfns...)
}

// FieldS16BE adds a field and reads 16 bit signed integer in big-endian
func (d *D) FieldS16BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S16BE, sfns...)
}

// Reader S17BE

// TryS17BE tries to read 17 bit signed integer in big-endian
func (d *D) TryS17BE() (int64, error) { return d.trySE(17, BigEndian) }

func (d *D) ScalarS17BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(17, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S17BE reads 17 bit signed integer in big-endian
func (d *D) S17BE() int64 {
	v, err := d.trySE(17, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S17BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS17BE tries to add a field and read 17 bit signed integer in big-endian
func (d *D) TryFieldS17BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS17BE, sfns...)
}

// FieldS17BE adds a field and reads 17 bit signed integer in big-endian
func (d *D) FieldS17BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S17BE, sfns...)
}

// Reader S18BE

// TryS18BE tries to read 18 bit signed integer in big-endian
func (d *D) TryS18BE() (int64, error) { return d.trySE(18, BigEndian) }

func (d *D) ScalarS18BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(18, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S18BE reads 18 bit signed integer in big-endian
func (d *D) S18BE() int64 {
	v, err := d.trySE(18, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S18BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS18BE tries to add a field and read 18 bit signed integer in big-endian
func (d *D) TryFieldS18BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS18BE, sfns...)
}

// FieldS18BE adds a field and reads 18 bit signed integer in big-endian
func (d *D) FieldS18BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S18BE, sfns...)
}

// Reader S19BE

// TryS19BE tries to read 19 bit signed integer in big-endian
func (d *D) TryS19BE() (int64, error) { return d.trySE(19, BigEndian) }

func (d *D) ScalarS19BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(19, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S19BE reads 19 bit signed integer in big-endian
func (d *D) S19BE() int64 {
	v, err := d.trySE(19, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S19BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS19BE tries to add a field and read 19 bit signed integer in big-endian
func (d *D) TryFieldS19BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS19BE, sfns...)
}

// FieldS19BE adds a field and reads 19 bit signed integer in big-endian
func (d *D) FieldS19BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S19BE, sfns...)
}

// Reader S20BE

// TryS20BE tries to read 20 bit signed integer in big-endian
func (d *D) TryS20BE() (int64, error) { return d.trySE(20, BigEndian) }

func (d *D) ScalarS20BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(20, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S20BE reads 20 bit signed integer in big-endian
func (d *D) S20BE() int64 {
	v, err := d.trySE(20, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S20BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS20BE tries to add a field and read 20 bit signed integer in big-endian
func (d *D) TryFieldS20BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS20BE, sfns...)
}

// FieldS20BE adds a field and reads 20 bit signed integer in big-endian
func (d *D) FieldS20BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S20BE, sfns...)
}

// Reader S21BE

// TryS21BE tries to read 21 bit signed integer in big-endian
func (d *D) TryS21BE() (int64, error) { return d.trySE(21, BigEndian) }

func (d *D) ScalarS21BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(21, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S21BE reads 21 bit signed integer in big-endian
func (d *D) S21BE() int64 {
	v, err := d.trySE(21, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S21BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS21BE tries to add a field and read 21 bit signed integer in big-endian
func (d *D) TryFieldS21BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS21BE, sfns...)
}

// FieldS21BE adds a field and reads 21 bit signed integer in big-endian
func (d *D) FieldS21BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S21BE, sfns...)
}

// Reader S22BE

// TryS22BE tries to read 22 bit signed integer in big-endian
func (d *D) TryS22BE() (int64, error) { return d.trySE(22, BigEndian) }

func (d *D) ScalarS22BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(22, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S22BE reads 22 bit signed integer in big-endian
func (d *D) S22BE() int64 {
	v, err := d.trySE(22, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S22BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS22BE tries to add a field and read 22 bit signed integer in big-endian
func (d *D) TryFieldS22BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS22BE, sfns...)
}

// FieldS22BE adds a field and reads 22 bit signed integer in big-endian
func (d *D) FieldS22BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S22BE, sfns...)
}

// Reader S23BE

// TryS23BE tries to read 23 bit signed integer in big-endian
func (d *D) TryS23BE() (int64, error) { return d.trySE(23, BigEndian) }

func (d *D) ScalarS23BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(23, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S23BE reads 23 bit signed integer in big-endian
func (d *D) S23BE() int64 {
	v, err := d.trySE(23, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S23BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS23BE tries to add a field and read 23 bit signed integer in big-endian
func (d *D) TryFieldS23BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS23BE, sfns...)
}

// FieldS23BE adds a field and reads 23 bit signed integer in big-endian
func (d *D) FieldS23BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S23BE, sfns...)
}

// Reader S24BE

// TryS24BE tries to read 24 bit signed integer in big-endian
func (d *D) TryS24BE() (int64, error) { return d.trySE(24, BigEndian) }

func (d *D) ScalarS24BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(24, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S24BE reads 24 bit signed integer in big-endian
func (d *D) S24BE() int64 {
	v, err := d.trySE(24, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S24BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS24BE tries to add a field and read 24 bit signed integer in big-endian
func (d *D) TryFieldS24BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS24BE, sfns...)
}

// FieldS24BE adds a field and reads 24 bit signed integer in big-endian
func (d *D) FieldS24BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S24BE, sfns...)
}

// Reader S25BE

// TryS25BE tries to read 25 bit signed integer in big-endian
func (d *D) TryS25BE() (int64, error) { return d.trySE(25, BigEndian) }

func (d *D) ScalarS25BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(25, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S25BE reads 25 bit signed integer in big-endian
func (d *D) S25BE() int64 {
	v, err := d.trySE(25, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S25BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS25BE tries to add a field and read 25 bit signed integer in big-endian
func (d *D) TryFieldS25BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS25BE, sfns...)
}

// FieldS25BE adds a field and reads 25 bit signed integer in big-endian
func (d *D) FieldS25BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S25BE, sfns...)
}

// Reader S26BE

// TryS26BE tries to read 26 bit signed integer in big-endian
func (d *D) TryS26BE() (int64, error) { return d.trySE(26, BigEndian) }

func (d *D) ScalarS26BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(26, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S26BE reads 26 bit signed integer in big-endian
func (d *D) S26BE() int64 {
	v, err := d.trySE(26, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S26BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS26BE tries to add a field and read 26 bit signed integer in big-endian
func (d *D) TryFieldS26BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS26BE, sfns...)
}

// FieldS26BE adds a field and reads 26 bit signed integer in big-endian
func (d *D) FieldS26BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S26BE, sfns...)
}

// Reader S27BE

// TryS27BE tries to read 27 bit signed integer in big-endian
func (d *D) TryS27BE() (int64, error) { return d.trySE(27, BigEndian) }

func (d *D) ScalarS27BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(27, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S27BE reads 27 bit signed integer in big-endian
func (d *D) S27BE() int64 {
	v, err := d.trySE(27, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S27BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS27BE tries to add a field and read 27 bit signed integer in big-endian
func (d *D) TryFieldS27BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS27BE, sfns...)
}

// FieldS27BE adds a field and reads 27 bit signed integer in big-endian
func (d *D) FieldS27BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S27BE, sfns...)
}

// Reader S28BE

// TryS28BE tries to read 28 bit signed integer in big-endian
func (d *D) TryS28BE() (int64, error) { return d.trySE(28, BigEndian) }

func (d *D) ScalarS28BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(28, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S28BE reads 28 bit signed integer in big-endian
func (d *D) S28BE() int64 {
	v, err := d.trySE(28, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S28BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS28BE tries to add a field and read 28 bit signed integer in big-endian
func (d *D) TryFieldS28BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS28BE, sfns...)
}

// FieldS28BE adds a field and reads 28 bit signed integer in big-endian
func (d *D) FieldS28BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S28BE, sfns...)
}

// Reader S29BE

// TryS29BE tries to read 29 bit signed integer in big-endian
func (d *D) TryS29BE() (int64, error) { return d.trySE(29, BigEndian) }

func (d *D) ScalarS29BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(29, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S29BE reads 29 bit signed integer in big-endian
func (d *D) S29BE() int64 {
	v, err := d.trySE(29, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S29BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS29BE tries to add a field and read 29 bit signed integer in big-endian
func (d *D) TryFieldS29BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS29BE, sfns...)
}

// FieldS29BE adds a field and reads 29 bit signed integer in big-endian
func (d *D) FieldS29BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S29BE, sfns...)
}

// Reader S30BE

// TryS30BE tries to read 30 bit signed integer in big-endian
func (d *D) TryS30BE() (int64, error) { return d.trySE(30, BigEndian) }

func (d *D) ScalarS30BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(30, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S30BE reads 30 bit signed integer in big-endian
func (d *D) S30BE() int64 {
	v, err := d.trySE(30, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S30BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS30BE tries to add a field and read 30 bit signed integer in big-endian
func (d *D) TryFieldS30BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS30BE, sfns...)
}

// FieldS30BE adds a field and reads 30 bit signed integer in big-endian
func (d *D) FieldS30BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S30BE, sfns...)
}

// Reader S31BE

// TryS31BE tries to read 31 bit signed integer in big-endian
func (d *D) TryS31BE() (int64, error) { return d.trySE(31, BigEndian) }

func (d *D) ScalarS31BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(31, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S31BE reads 31 bit signed integer in big-endian
func (d *D) S31BE() int64 {
	v, err := d.trySE(31, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S31BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS31BE tries to add a field and read 31 bit signed integer in big-endian
func (d *D) TryFieldS31BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS31BE, sfns...)
}

// FieldS31BE adds a field and reads 31 bit signed integer in big-endian
func (d *D) FieldS31BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S31BE, sfns...)
}

// Reader S32BE

// TryS32BE tries to read 32 bit signed integer in big-endian
func (d *D) TryS32BE() (int64, error) { return d.trySE(32, BigEndian) }

func (d *D) ScalarS32BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(32, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S32BE reads 32 bit signed integer in big-endian
func (d *D) S32BE() int64 {
	v, err := d.trySE(32, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S32BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS32BE tries to add a field and read 32 bit signed integer in big-endian
func (d *D) TryFieldS32BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS32BE, sfns...)
}

// FieldS32BE adds a field and reads 32 bit signed integer in big-endian
func (d *D) FieldS32BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S32BE, sfns...)
}

// Reader S33BE

// TryS33BE tries to read 33 bit signed integer in big-endian
func (d *D) TryS33BE() (int64, error) { return d.trySE(33, BigEndian) }

func (d *D) ScalarS33BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(33, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S33BE reads 33 bit signed integer in big-endian
func (d *D) S33BE() int64 {
	v, err := d.trySE(33, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S33BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS33BE tries to add a field and read 33 bit signed integer in big-endian
func (d *D) TryFieldS33BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS33BE, sfns...)
}

// FieldS33BE adds a field and reads 33 bit signed integer in big-endian
func (d *D) FieldS33BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S33BE, sfns...)
}

// Reader S34BE

// TryS34BE tries to read 34 bit signed integer in big-endian
func (d *D) TryS34BE() (int64, error) { return d.trySE(34, BigEndian) }

func (d *D) ScalarS34BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(34, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S34BE reads 34 bit signed integer in big-endian
func (d *D) S34BE() int64 {
	v, err := d.trySE(34, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S34BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS34BE tries to add a field and read 34 bit signed integer in big-endian
func (d *D) TryFieldS34BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS34BE, sfns...)
}

// FieldS34BE adds a field and reads 34 bit signed integer in big-endian
func (d *D) FieldS34BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S34BE, sfns...)
}

// Reader S35BE

// TryS35BE tries to read 35 bit signed integer in big-endian
func (d *D) TryS35BE() (int64, error) { return d.trySE(35, BigEndian) }

func (d *D) ScalarS35BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(35, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S35BE reads 35 bit signed integer in big-endian
func (d *D) S35BE() int64 {
	v, err := d.trySE(35, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S35BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS35BE tries to add a field and read 35 bit signed integer in big-endian
func (d *D) TryFieldS35BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS35BE, sfns...)
}

// FieldS35BE adds a field and reads 35 bit signed integer in big-endian
func (d *D) FieldS35BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S35BE, sfns...)
}

// Reader S36BE

// TryS36BE tries to read 36 bit signed integer in big-endian
func (d *D) TryS36BE() (int64, error) { return d.trySE(36, BigEndian) }

func (d *D) ScalarS36BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(36, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S36BE reads 36 bit signed integer in big-endian
func (d *D) S36BE() int64 {
	v, err := d.trySE(36, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S36BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS36BE tries to add a field and read 36 bit signed integer in big-endian
func (d *D) TryFieldS36BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS36BE, sfns...)
}

// FieldS36BE adds a field and reads 36 bit signed integer in big-endian
func (d *D) FieldS36BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S36BE, sfns...)
}

// Reader S37BE

// TryS37BE tries to read 37 bit signed integer in big-endian
func (d *D) TryS37BE() (int64, error) { return d.trySE(37, BigEndian) }

func (d *D) ScalarS37BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(37, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S37BE reads 37 bit signed integer in big-endian
func (d *D) S37BE() int64 {
	v, err := d.trySE(37, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S37BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS37BE tries to add a field and read 37 bit signed integer in big-endian
func (d *D) TryFieldS37BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS37BE, sfns...)
}

// FieldS37BE adds a field and reads 37 bit signed integer in big-endian
func (d *D) FieldS37BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S37BE, sfns...)
}

// Reader S38BE

// TryS38BE tries to read 38 bit signed integer in big-endian
func (d *D) TryS38BE() (int64, error) { return d.trySE(38, BigEndian) }

func (d *D) ScalarS38BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(38, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S38BE reads 38 bit signed integer in big-endian
func (d *D) S38BE() int64 {
	v, err := d.trySE(38, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S38BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS38BE tries to add a field and read 38 bit signed integer in big-endian
func (d *D) TryFieldS38BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS38BE, sfns...)
}

// FieldS38BE adds a field and reads 38 bit signed integer in big-endian
func (d *D) FieldS38BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S38BE, sfns...)
}

// Reader S39BE

// TryS39BE tries to read 39 bit signed integer in big-endian
func (d *D) TryS39BE() (int64, error) { return d.trySE(39, BigEndian) }

func (d *D) ScalarS39BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(39, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S39BE reads 39 bit signed integer in big-endian
func (d *D) S39BE() int64 {
	v, err := d.trySE(39, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S39BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS39BE tries to add a field and read 39 bit signed integer in big-endian
func (d *D) TryFieldS39BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS39BE, sfns...)
}

// FieldS39BE adds a field and reads 39 bit signed integer in big-endian
func (d *D) FieldS39BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S39BE, sfns...)
}

// Reader S40BE

// TryS40BE tries to read 40 bit signed integer in big-endian
func (d *D) TryS40BE() (int64, error) { return d.trySE(40, BigEndian) }

func (d *D) ScalarS40BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(40, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S40BE reads 40 bit signed integer in big-endian
func (d *D) S40BE() int64 {
	v, err := d.trySE(40, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S40BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS40BE tries to add a field and read 40 bit signed integer in big-endian
func (d *D) TryFieldS40BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS40BE, sfns...)
}

// FieldS40BE adds a field and reads 40 bit signed integer in big-endian
func (d *D) FieldS40BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S40BE, sfns...)
}

// Reader S41BE

// TryS41BE tries to read 41 bit signed integer in big-endian
func (d *D) TryS41BE() (int64, error) { return d.trySE(41, BigEndian) }

func (d *D) ScalarS41BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(41, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S41BE reads 41 bit signed integer in big-endian
func (d *D) S41BE() int64 {
	v, err := d.trySE(41, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S41BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS41BE tries to add a field and read 41 bit signed integer in big-endian
func (d *D) TryFieldS41BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS41BE, sfns...)
}

// FieldS41BE adds a field and reads 41 bit signed integer in big-endian
func (d *D) FieldS41BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S41BE, sfns...)
}

// Reader S42BE

// TryS42BE tries to read 42 bit signed integer in big-endian
func (d *D) TryS42BE() (int64, error) { return d.trySE(42, BigEndian) }

func (d *D) ScalarS42BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(42, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S42BE reads 42 bit signed integer in big-endian
func (d *D) S42BE() int64 {
	v, err := d.trySE(42, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S42BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS42BE tries to add a field and read 42 bit signed integer in big-endian
func (d *D) TryFieldS42BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS42BE, sfns...)
}

// FieldS42BE adds a field and reads 42 bit signed integer in big-endian
func (d *D) FieldS42BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S42BE, sfns...)
}

// Reader S43BE

// TryS43BE tries to read 43 bit signed integer in big-endian
func (d *D) TryS43BE() (int64, error) { return d.trySE(43, BigEndian) }

func (d *D) ScalarS43BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(43, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S43BE reads 43 bit signed integer in big-endian
func (d *D) S43BE() int64 {
	v, err := d.trySE(43, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S43BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS43BE tries to add a field and read 43 bit signed integer in big-endian
func (d *D) TryFieldS43BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS43BE, sfns...)
}

// FieldS43BE adds a field and reads 43 bit signed integer in big-endian
func (d *D) FieldS43BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S43BE, sfns...)
}

// Reader S44BE

// TryS44BE tries to read 44 bit signed integer in big-endian
func (d *D) TryS44BE() (int64, error) { return d.trySE(44, BigEndian) }

func (d *D) ScalarS44BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(44, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S44BE reads 44 bit signed integer in big-endian
func (d *D) S44BE() int64 {
	v, err := d.trySE(44, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S44BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS44BE tries to add a field and read 44 bit signed integer in big-endian
func (d *D) TryFieldS44BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS44BE, sfns...)
}

// FieldS44BE adds a field and reads 44 bit signed integer in big-endian
func (d *D) FieldS44BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S44BE, sfns...)
}

// Reader S45BE

// TryS45BE tries to read 45 bit signed integer in big-endian
func (d *D) TryS45BE() (int64, error) { return d.trySE(45, BigEndian) }

func (d *D) ScalarS45BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(45, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S45BE reads 45 bit signed integer in big-endian
func (d *D) S45BE() int64 {
	v, err := d.trySE(45, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S45BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS45BE tries to add a field and read 45 bit signed integer in big-endian
func (d *D) TryFieldS45BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS45BE, sfns...)
}

// FieldS45BE adds a field and reads 45 bit signed integer in big-endian
func (d *D) FieldS45BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S45BE, sfns...)
}

// Reader S46BE

// TryS46BE tries to read 46 bit signed integer in big-endian
func (d *D) TryS46BE() (int64, error) { return d.trySE(46, BigEndian) }

func (d *D) ScalarS46BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(46, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S46BE reads 46 bit signed integer in big-endian
func (d *D) S46BE() int64 {
	v, err := d.trySE(46, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S46BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS46BE tries to add a field and read 46 bit signed integer in big-endian
func (d *D) TryFieldS46BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS46BE, sfns...)
}

// FieldS46BE adds a field and reads 46 bit signed integer in big-endian
func (d *D) FieldS46BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S46BE, sfns...)
}

// Reader S47BE

// TryS47BE tries to read 47 bit signed integer in big-endian
func (d *D) TryS47BE() (int64, error) { return d.trySE(47, BigEndian) }

func (d *D) ScalarS47BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(47, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S47BE reads 47 bit signed integer in big-endian
func (d *D) S47BE() int64 {
	v, err := d.trySE(47, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S47BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS47BE tries to add a field and read 47 bit signed integer in big-endian
func (d *D) TryFieldS47BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS47BE, sfns...)
}

// FieldS47BE adds a field and reads 47 bit signed integer in big-endian
func (d *D) FieldS47BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S47BE, sfns...)
}

// Reader S48BE

// TryS48BE tries to read 48 bit signed integer in big-endian
func (d *D) TryS48BE() (int64, error) { return d.trySE(48, BigEndian) }

func (d *D) ScalarS48BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(48, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S48BE reads 48 bit signed integer in big-endian
func (d *D) S48BE() int64 {
	v, err := d.trySE(48, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S48BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS48BE tries to add a field and read 48 bit signed integer in big-endian
func (d *D) TryFieldS48BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS48BE, sfns...)
}

// FieldS48BE adds a field and reads 48 bit signed integer in big-endian
func (d *D) FieldS48BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S48BE, sfns...)
}

// Reader S49BE

// TryS49BE tries to read 49 bit signed integer in big-endian
func (d *D) TryS49BE() (int64, error) { return d.trySE(49, BigEndian) }

func (d *D) ScalarS49BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(49, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S49BE reads 49 bit signed integer in big-endian
func (d *D) S49BE() int64 {
	v, err := d.trySE(49, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S49BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS49BE tries to add a field and read 49 bit signed integer in big-endian
func (d *D) TryFieldS49BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS49BE, sfns...)
}

// FieldS49BE adds a field and reads 49 bit signed integer in big-endian
func (d *D) FieldS49BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S49BE, sfns...)
}

// Reader S50BE

// TryS50BE tries to read 50 bit signed integer in big-endian
func (d *D) TryS50BE() (int64, error) { return d.trySE(50, BigEndian) }

func (d *D) ScalarS50BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(50, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S50BE reads 50 bit signed integer in big-endian
func (d *D) S50BE() int64 {
	v, err := d.trySE(50, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S50BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS50BE tries to add a field and read 50 bit signed integer in big-endian
func (d *D) TryFieldS50BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS50BE, sfns...)
}

// FieldS50BE adds a field and reads 50 bit signed integer in big-endian
func (d *D) FieldS50BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S50BE, sfns...)
}

// Reader S51BE

// TryS51BE tries to read 51 bit signed integer in big-endian
func (d *D) TryS51BE() (int64, error) { return d.trySE(51, BigEndian) }

func (d *D) ScalarS51BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(51, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S51BE reads 51 bit signed integer in big-endian
func (d *D) S51BE() int64 {
	v, err := d.trySE(51, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S51BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS51BE tries to add a field and read 51 bit signed integer in big-endian
func (d *D) TryFieldS51BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS51BE, sfns...)
}

// FieldS51BE adds a field and reads 51 bit signed integer in big-endian
func (d *D) FieldS51BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S51BE, sfns...)
}

// Reader S52BE

// TryS52BE tries to read 52 bit signed integer in big-endian
func (d *D) TryS52BE() (int64, error) { return d.trySE(52, BigEndian) }

func (d *D) ScalarS52BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(52, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S52BE reads 52 bit signed integer in big-endian
func (d *D) S52BE() int64 {
	v, err := d.trySE(52, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S52BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS52BE tries to add a field and read 52 bit signed integer in big-endian
func (d *D) TryFieldS52BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS52BE, sfns...)
}

// FieldS52BE adds a field and reads 52 bit signed integer in big-endian
func (d *D) FieldS52BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S52BE, sfns...)
}

// Reader S53BE

// TryS53BE tries to read 53 bit signed integer in big-endian
func (d *D) TryS53BE() (int64, error) { return d.trySE(53, BigEndian) }

func (d *D) ScalarS53BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(53, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S53BE reads 53 bit signed integer in big-endian
func (d *D) S53BE() int64 {
	v, err := d.trySE(53, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S53BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS53BE tries to add a field and read 53 bit signed integer in big-endian
func (d *D) TryFieldS53BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS53BE, sfns...)
}

// FieldS53BE adds a field and reads 53 bit signed integer in big-endian
func (d *D) FieldS53BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S53BE, sfns...)
}

// Reader S54BE

// TryS54BE tries to read 54 bit signed integer in big-endian
func (d *D) TryS54BE() (int64, error) { return d.trySE(54, BigEndian) }

func (d *D) ScalarS54BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(54, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S54BE reads 54 bit signed integer in big-endian
func (d *D) S54BE() int64 {
	v, err := d.trySE(54, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S54BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS54BE tries to add a field and read 54 bit signed integer in big-endian
func (d *D) TryFieldS54BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS54BE, sfns...)
}

// FieldS54BE adds a field and reads 54 bit signed integer in big-endian
func (d *D) FieldS54BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S54BE, sfns...)
}

// Reader S55BE

// TryS55BE tries to read 55 bit signed integer in big-endian
func (d *D) TryS55BE() (int64, error) { return d.trySE(55, BigEndian) }

func (d *D) ScalarS55BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(55, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S55BE reads 55 bit signed integer in big-endian
func (d *D) S55BE() int64 {
	v, err := d.trySE(55, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S55BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS55BE tries to add a field and read 55 bit signed integer in big-endian
func (d *D) TryFieldS55BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS55BE, sfns...)
}

// FieldS55BE adds a field and reads 55 bit signed integer in big-endian
func (d *D) FieldS55BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S55BE, sfns...)
}

// Reader S56BE

// TryS56BE tries to read 56 bit signed integer in big-endian
func (d *D) TryS56BE() (int64, error) { return d.trySE(56, BigEndian) }

func (d *D) ScalarS56BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(56, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S56BE reads 56 bit signed integer in big-endian
func (d *D) S56BE() int64 {
	v, err := d.trySE(56, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S56BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS56BE tries to add a field and read 56 bit signed integer in big-endian
func (d *D) TryFieldS56BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS56BE, sfns...)
}

// FieldS56BE adds a field and reads 56 bit signed integer in big-endian
func (d *D) FieldS56BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S56BE, sfns...)
}

// Reader S57BE

// TryS57BE tries to read 57 bit signed integer in big-endian
func (d *D) TryS57BE() (int64, error) { return d.trySE(57, BigEndian) }

func (d *D) ScalarS57BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(57, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S57BE reads 57 bit signed integer in big-endian
func (d *D) S57BE() int64 {
	v, err := d.trySE(57, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S57BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS57BE tries to add a field and read 57 bit signed integer in big-endian
func (d *D) TryFieldS57BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS57BE, sfns...)
}

// FieldS57BE adds a field and reads 57 bit signed integer in big-endian
func (d *D) FieldS57BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S57BE, sfns...)
}

// Reader S58BE

// TryS58BE tries to read 58 bit signed integer in big-endian
func (d *D) TryS58BE() (int64, error) { return d.trySE(58, BigEndian) }

func (d *D) ScalarS58BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(58, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S58BE reads 58 bit signed integer in big-endian
func (d *D) S58BE() int64 {
	v, err := d.trySE(58, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S58BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS58BE tries to add a field and read 58 bit signed integer in big-endian
func (d *D) TryFieldS58BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS58BE, sfns...)
}

// FieldS58BE adds a field and reads 58 bit signed integer in big-endian
func (d *D) FieldS58BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S58BE, sfns...)
}

// Reader S59BE

// TryS59BE tries to read 59 bit signed integer in big-endian
func (d *D) TryS59BE() (int64, error) { return d.trySE(59, BigEndian) }

func (d *D) ScalarS59BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(59, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S59BE reads 59 bit signed integer in big-endian
func (d *D) S59BE() int64 {
	v, err := d.trySE(59, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S59BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS59BE tries to add a field and read 59 bit signed integer in big-endian
func (d *D) TryFieldS59BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS59BE, sfns...)
}

// FieldS59BE adds a field and reads 59 bit signed integer in big-endian
func (d *D) FieldS59BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S59BE, sfns...)
}

// Reader S60BE

// TryS60BE tries to read 60 bit signed integer in big-endian
func (d *D) TryS60BE() (int64, error) { return d.trySE(60, BigEndian) }

func (d *D) ScalarS60BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(60, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S60BE reads 60 bit signed integer in big-endian
func (d *D) S60BE() int64 {
	v, err := d.trySE(60, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S60BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS60BE tries to add a field and read 60 bit signed integer in big-endian
func (d *D) TryFieldS60BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS60BE, sfns...)
}

// FieldS60BE adds a field and reads 60 bit signed integer in big-endian
func (d *D) FieldS60BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S60BE, sfns...)
}

// Reader S61BE

// TryS61BE tries to read 61 bit signed integer in big-endian
func (d *D) TryS61BE() (int64, error) { return d.trySE(61, BigEndian) }

func (d *D) ScalarS61BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(61, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S61BE reads 61 bit signed integer in big-endian
func (d *D) S61BE() int64 {
	v, err := d.trySE(61, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S61BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS61BE tries to add a field and read 61 bit signed integer in big-endian
func (d *D) TryFieldS61BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS61BE, sfns...)
}

// FieldS61BE adds a field and reads 61 bit signed integer in big-endian
func (d *D) FieldS61BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S61BE, sfns...)
}

// Reader S62BE

// TryS62BE tries to read 62 bit signed integer in big-endian
func (d *D) TryS62BE() (int64, error) { return d.trySE(62, BigEndian) }

func (d *D) ScalarS62BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(62, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S62BE reads 62 bit signed integer in big-endian
func (d *D) S62BE() int64 {
	v, err := d.trySE(62, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S62BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS62BE tries to add a field and read 62 bit signed integer in big-endian
func (d *D) TryFieldS62BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS62BE, sfns...)
}

// FieldS62BE adds a field and reads 62 bit signed integer in big-endian
func (d *D) FieldS62BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S62BE, sfns...)
}

// Reader S63BE

// TryS63BE tries to read 63 bit signed integer in big-endian
func (d *D) TryS63BE() (int64, error) { return d.trySE(63, BigEndian) }

func (d *D) ScalarS63BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(63, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S63BE reads 63 bit signed integer in big-endian
func (d *D) S63BE() int64 {
	v, err := d.trySE(63, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S63BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS63BE tries to add a field and read 63 bit signed integer in big-endian
func (d *D) TryFieldS63BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS63BE, sfns...)
}

// FieldS63BE adds a field and reads 63 bit signed integer in big-endian
func (d *D) FieldS63BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S63BE, sfns...)
}

// Reader S64BE

// TryS64BE tries to read 64 bit signed integer in big-endian
func (d *D) TryS64BE() (int64, error) { return d.trySE(64, BigEndian) }

func (d *D) ScalarS64BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.trySE(64, BigEndian)
		s.Actual = v
		return s, err
	}
}

// S64BE reads 64 bit signed integer in big-endian
func (d *D) S64BE() int64 {
	v, err := d.trySE(64, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "S64BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldS64BE tries to add a field and read 64 bit signed integer in big-endian
func (d *D) TryFieldS64BE(name string, sfns ...ScalarFn) (int64, error) {
	return d.TryFieldSFn(name, (*D).TryS64BE, sfns...)
}

// FieldS64BE adds a field and reads 64 bit signed integer in big-endian
func (d *D) FieldS64BE(name string, sfns ...ScalarFn) int64 {
	return d.FieldSFn(name, (*D).S64BE, sfns...)
}

// Reader F

// TryF tries to read nBit IEEE 754 float in current endian
func (d *D) TryF(nBits int) (float64, error) { return d.tryFE(nBits, d.Endian) }

func (d *D) ScalarF(nBits int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(nBits, d.Endian)
		s.Actual = v
		return s, err
	}
}

// F reads nBit IEEE 754 float in current endian
func (d *D) F(nBits int) float64 {
	v, err := d.tryFE(nBits, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "F", Pos: d.Pos()})
	}
	return v
}

// TryFieldF tries to add a field and read nBit IEEE 754 float in current endian
func (d *D) TryFieldF(name string, nBits int, sfns ...ScalarFn) (float64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarF(nBits), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualF(), err
}

// FieldF adds a field and reads nBit IEEE 754 float in current endian
func (d *D) FieldF(name string, nBits int, sfns ...ScalarFn) float64 {
	v, err := d.TryFieldF(name, nBits, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "F", Pos: d.Pos()})
	}
	return v
}

// Reader FE

// TryFE tries to read nBit IEEE 754 float in specified endian
func (d *D) TryFE(nBits int, endian Endian) (float64, error) { return d.tryFE(nBits, endian) }

func (d *D) ScalarFE(nBits int, endian Endian) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(nBits, endian)
		s.Actual = v
		return s, err
	}
}

// FE reads nBit IEEE 754 float in specified endian
func (d *D) FE(nBits int, endian Endian) float64 {
	v, err := d.tryFE(nBits, endian)
	if err != nil {
		panic(IOError{Err: err, Op: "FE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFE tries to add a field and read nBit IEEE 754 float in specified endian
func (d *D) TryFieldFE(name string, nBits int, endian Endian, sfns ...ScalarFn) (float64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarFE(nBits, endian), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualF(), err
}

// FieldFE adds a field and reads nBit IEEE 754 float in specified endian
func (d *D) FieldFE(name string, nBits int, endian Endian, sfns ...ScalarFn) float64 {
	v, err := d.TryFieldFE(name, nBits, endian, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "FE", Pos: d.Pos()})
	}
	return v
}

// Reader F16

// TryF16 tries to read 16 bit IEEE 754 float in current endian
func (d *D) TryF16() (float64, error) { return d.tryFE(16, d.Endian) }

func (d *D) ScalarF16() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(16, d.Endian)
		s.Actual = v
		return s, err
	}
}

// F16 reads 16 bit IEEE 754 float in current endian
func (d *D) F16() float64 {
	v, err := d.tryFE(16, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "F16", Pos: d.Pos()})
	}
	return v
}

// TryFieldF16 tries to add a field and read 16 bit IEEE 754 float in current endian
func (d *D) TryFieldF16(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF16, sfns...)
}

// FieldF16 adds a field and reads 16 bit IEEE 754 float in current endian
func (d *D) FieldF16(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F16, sfns...)
}

// Reader F32

// TryF32 tries to read 32 bit IEEE 754 float in current endian
func (d *D) TryF32() (float64, error) { return d.tryFE(32, d.Endian) }

func (d *D) ScalarF32() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(32, d.Endian)
		s.Actual = v
		return s, err
	}
}

// F32 reads 32 bit IEEE 754 float in current endian
func (d *D) F32() float64 {
	v, err := d.tryFE(32, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "F32", Pos: d.Pos()})
	}
	return v
}

// TryFieldF32 tries to add a field and read 32 bit IEEE 754 float in current endian
func (d *D) TryFieldF32(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF32, sfns...)
}

// FieldF32 adds a field and reads 32 bit IEEE 754 float in current endian
func (d *D) FieldF32(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F32, sfns...)
}

// Reader F64

// TryF64 tries to read 64 bit IEEE 754 float in current endian
func (d *D) TryF64() (float64, error) { return d.tryFE(64, d.Endian) }

func (d *D) ScalarF64() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(64, d.Endian)
		s.Actual = v
		return s, err
	}
}

// F64 reads 64 bit IEEE 754 float in current endian
func (d *D) F64() float64 {
	v, err := d.tryFE(64, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "F64", Pos: d.Pos()})
	}
	return v
}

// TryFieldF64 tries to add a field and read 64 bit IEEE 754 float in current endian
func (d *D) TryFieldF64(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF64, sfns...)
}

// FieldF64 adds a field and reads 64 bit IEEE 754 float in current endian
func (d *D) FieldF64(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F64, sfns...)
}

// Reader F16LE

// TryF16LE tries to read 16 bit IEEE 754 float in little-endian
func (d *D) TryF16LE() (float64, error) { return d.tryFE(16, LittleEndian) }

func (d *D) ScalarF16LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(16, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// F16LE reads 16 bit IEEE 754 float in little-endian
func (d *D) F16LE() float64 {
	v, err := d.tryFE(16, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "F16LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldF16LE tries to add a field and read 16 bit IEEE 754 float in little-endian
func (d *D) TryFieldF16LE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF16LE, sfns...)
}

// FieldF16LE adds a field and reads 16 bit IEEE 754 float in little-endian
func (d *D) FieldF16LE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F16LE, sfns...)
}

// Reader F32LE

// TryF32LE tries to read 32 bit IEEE 754 float in little-endian
func (d *D) TryF32LE() (float64, error) { return d.tryFE(32, LittleEndian) }

func (d *D) ScalarF32LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(32, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// F32LE reads 32 bit IEEE 754 float in little-endian
func (d *D) F32LE() float64 {
	v, err := d.tryFE(32, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "F32LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldF32LE tries to add a field and read 32 bit IEEE 754 float in little-endian
func (d *D) TryFieldF32LE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF32LE, sfns...)
}

// FieldF32LE adds a field and reads 32 bit IEEE 754 float in little-endian
func (d *D) FieldF32LE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F32LE, sfns...)
}

// Reader F64LE

// TryF64LE tries to read 64 bit IEEE 754 float in little-endian
func (d *D) TryF64LE() (float64, error) { return d.tryFE(64, LittleEndian) }

func (d *D) ScalarF64LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(64, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// F64LE reads 64 bit IEEE 754 float in little-endian
func (d *D) F64LE() float64 {
	v, err := d.tryFE(64, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "F64LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldF64LE tries to add a field and read 64 bit IEEE 754 float in little-endian
func (d *D) TryFieldF64LE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF64LE, sfns...)
}

// FieldF64LE adds a field and reads 64 bit IEEE 754 float in little-endian
func (d *D) FieldF64LE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F64LE, sfns...)
}

// Reader F16BE

// TryF16BE tries to read 16 bit IEEE 754 float in big-endian
func (d *D) TryF16BE() (float64, error) { return d.tryFE(16, BigEndian) }

func (d *D) ScalarF16BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(16, BigEndian)
		s.Actual = v
		return s, err
	}
}

// F16BE reads 16 bit IEEE 754 float in big-endian
func (d *D) F16BE() float64 {
	v, err := d.tryFE(16, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "F16BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldF16BE tries to add a field and read 16 bit IEEE 754 float in big-endian
func (d *D) TryFieldF16BE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF16BE, sfns...)
}

// FieldF16BE adds a field and reads 16 bit IEEE 754 float in big-endian
func (d *D) FieldF16BE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F16BE, sfns...)
}

// Reader F32BE

// TryF32BE tries to read 32 bit IEEE 754 float in big-endian
func (d *D) TryF32BE() (float64, error) { return d.tryFE(32, BigEndian) }

func (d *D) ScalarF32BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(32, BigEndian)
		s.Actual = v
		return s, err
	}
}

// F32BE reads 32 bit IEEE 754 float in big-endian
func (d *D) F32BE() float64 {
	v, err := d.tryFE(32, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "F32BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldF32BE tries to add a field and read 32 bit IEEE 754 float in big-endian
func (d *D) TryFieldF32BE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF32BE, sfns...)
}

// FieldF32BE adds a field and reads 32 bit IEEE 754 float in big-endian
func (d *D) FieldF32BE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F32BE, sfns...)
}

// Reader F64BE

// TryF64BE tries to read 64 bit IEEE 754 float in big-endian
func (d *D) TryF64BE() (float64, error) { return d.tryFE(64, BigEndian) }

func (d *D) ScalarF64BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFE(64, BigEndian)
		s.Actual = v
		return s, err
	}
}

// F64BE reads 64 bit IEEE 754 float in big-endian
func (d *D) F64BE() float64 {
	v, err := d.tryFE(64, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "F64BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldF64BE tries to add a field and read 64 bit IEEE 754 float in big-endian
func (d *D) TryFieldF64BE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryF64BE, sfns...)
}

// FieldF64BE adds a field and reads 64 bit IEEE 754 float in big-endian
func (d *D) FieldF64BE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).F64BE, sfns...)
}

// Reader FP

// TryFP tries to read nBits fixed-point number in current endian
func (d *D) TryFP(nBits int, fBits int64) (float64, error) { return d.tryFPE(nBits, fBits, d.Endian) }

func (d *D) ScalarFP(nBits int, fBits int64) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(nBits, fBits, d.Endian)
		s.Actual = v
		return s, err
	}
}

// FP reads nBits fixed-point number in current endian
func (d *D) FP(nBits int, fBits int64) float64 {
	v, err := d.tryFPE(nBits, fBits, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP tries to add a field and read nBits fixed-point number in current endian
func (d *D) TryFieldFP(name string, nBits int, fBits int64, sfns ...ScalarFn) (float64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarFP(nBits, fBits), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualF(), err
}

// FieldFP adds a field and reads nBits fixed-point number in current endian
func (d *D) FieldFP(name string, nBits int, fBits int64, sfns ...ScalarFn) float64 {
	v, err := d.TryFieldFP(name, nBits, fBits, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "FP", Pos: d.Pos()})
	}
	return v
}

// Reader FPE

// TryFPE tries to read nBits fixed-point number in specified endian
func (d *D) TryFPE(nBits int, fBits int64, endian Endian) (float64, error) {
	return d.tryFPE(nBits, fBits, endian)
}

func (d *D) ScalarFPE(nBits int, fBits int64, endian Endian) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(nBits, fBits, endian)
		s.Actual = v
		return s, err
	}
}

// FPE reads nBits fixed-point number in specified endian
func (d *D) FPE(nBits int, fBits int64, endian Endian) float64 {
	v, err := d.tryFPE(nBits, fBits, endian)
	if err != nil {
		panic(IOError{Err: err, Op: "FPE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFPE tries to add a field and read nBits fixed-point number in specified endian
func (d *D) TryFieldFPE(name string, nBits int, fBits int64, endian Endian, sfns ...ScalarFn) (float64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarFPE(nBits, fBits, endian), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualF(), err
}

// FieldFPE adds a field and reads nBits fixed-point number in specified endian
func (d *D) FieldFPE(name string, nBits int, fBits int64, endian Endian, sfns ...ScalarFn) float64 {
	v, err := d.TryFieldFPE(name, nBits, fBits, endian, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "FPE", Pos: d.Pos()})
	}
	return v
}

// Reader FP16

// TryFP16 tries to read 16 bit fixed-point number in current endian
func (d *D) TryFP16() (float64, error) { return d.tryFPE(16, 8, d.Endian) }

func (d *D) ScalarFP16() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(16, 8, d.Endian)
		s.Actual = v
		return s, err
	}
}

// FP16 reads 16 bit fixed-point number in current endian
func (d *D) FP16() float64 {
	v, err := d.tryFPE(16, 8, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP16", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP16 tries to add a field and read 16 bit fixed-point number in current endian
func (d *D) TryFieldFP16(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP16, sfns...)
}

// FieldFP16 adds a field and reads 16 bit fixed-point number in current endian
func (d *D) FieldFP16(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP16, sfns...)
}

// Reader FP32

// TryFP32 tries to read 32 bit fixed-point number in current endian
func (d *D) TryFP32() (float64, error) { return d.tryFPE(32, 16, d.Endian) }

func (d *D) ScalarFP32() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(32, 16, d.Endian)
		s.Actual = v
		return s, err
	}
}

// FP32 reads 32 bit fixed-point number in current endian
func (d *D) FP32() float64 {
	v, err := d.tryFPE(32, 16, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP32", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP32 tries to add a field and read 32 bit fixed-point number in current endian
func (d *D) TryFieldFP32(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP32, sfns...)
}

// FieldFP32 adds a field and reads 32 bit fixed-point number in current endian
func (d *D) FieldFP32(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP32, sfns...)
}

// Reader FP64

// TryFP64 tries to read 64 bit fixed-point number in current endian
func (d *D) TryFP64() (float64, error) { return d.tryFPE(64, 32, d.Endian) }

func (d *D) ScalarFP64() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(64, 32, d.Endian)
		s.Actual = v
		return s, err
	}
}

// FP64 reads 64 bit fixed-point number in current endian
func (d *D) FP64() float64 {
	v, err := d.tryFPE(64, 32, d.Endian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP64", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP64 tries to add a field and read 64 bit fixed-point number in current endian
func (d *D) TryFieldFP64(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP64, sfns...)
}

// FieldFP64 adds a field and reads 64 bit fixed-point number in current endian
func (d *D) FieldFP64(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP64, sfns...)
}

// Reader FP16LE

// TryFP16LE tries to read 16 bit fixed-point number in little-endian
func (d *D) TryFP16LE() (float64, error) { return d.tryFPE(16, 8, LittleEndian) }

func (d *D) ScalarFP16LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(16, 8, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// FP16LE reads 16 bit fixed-point number in little-endian
func (d *D) FP16LE() float64 {
	v, err := d.tryFPE(16, 8, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP16LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP16LE tries to add a field and read 16 bit fixed-point number in little-endian
func (d *D) TryFieldFP16LE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP16LE, sfns...)
}

// FieldFP16LE adds a field and reads 16 bit fixed-point number in little-endian
func (d *D) FieldFP16LE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP16LE, sfns...)
}

// Reader FP32LE

// TryFP32LE tries to read 32 bit fixed-point number in little-endian
func (d *D) TryFP32LE() (float64, error) { return d.tryFPE(32, 16, LittleEndian) }

func (d *D) ScalarFP32LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(32, 16, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// FP32LE reads 32 bit fixed-point number in little-endian
func (d *D) FP32LE() float64 {
	v, err := d.tryFPE(32, 16, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP32LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP32LE tries to add a field and read 32 bit fixed-point number in little-endian
func (d *D) TryFieldFP32LE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP32LE, sfns...)
}

// FieldFP32LE adds a field and reads 32 bit fixed-point number in little-endian
func (d *D) FieldFP32LE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP32LE, sfns...)
}

// Reader FP64LE

// TryFP64LE tries to read 64 bit fixed-point number in little-endian
func (d *D) TryFP64LE() (float64, error) { return d.tryFPE(64, 32, LittleEndian) }

func (d *D) ScalarFP64LE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(64, 32, LittleEndian)
		s.Actual = v
		return s, err
	}
}

// FP64LE reads 64 bit fixed-point number in little-endian
func (d *D) FP64LE() float64 {
	v, err := d.tryFPE(64, 32, LittleEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP64LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP64LE tries to add a field and read 64 bit fixed-point number in little-endian
func (d *D) TryFieldFP64LE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP64LE, sfns...)
}

// FieldFP64LE adds a field and reads 64 bit fixed-point number in little-endian
func (d *D) FieldFP64LE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP64LE, sfns...)
}

// Reader FP16BE

// TryFP16BE tries to read 16 bit fixed-point number in big-endian
func (d *D) TryFP16BE() (float64, error) { return d.tryFPE(16, 8, BigEndian) }

func (d *D) ScalarFP16BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(16, 8, BigEndian)
		s.Actual = v
		return s, err
	}
}

// FP16BE reads 16 bit fixed-point number in big-endian
func (d *D) FP16BE() float64 {
	v, err := d.tryFPE(16, 8, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP16BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP16BE tries to add a field and read 16 bit fixed-point number in big-endian
func (d *D) TryFieldFP16BE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP16BE, sfns...)
}

// FieldFP16BE adds a field and reads 16 bit fixed-point number in big-endian
func (d *D) FieldFP16BE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP16BE, sfns...)
}

// Reader FP32BE

// TryFP32BE tries to read 32 bit fixed-point number in big-endian
func (d *D) TryFP32BE() (float64, error) { return d.tryFPE(32, 16, BigEndian) }

func (d *D) ScalarFP32BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(32, 16, BigEndian)
		s.Actual = v
		return s, err
	}
}

// FP32BE reads 32 bit fixed-point number in big-endian
func (d *D) FP32BE() float64 {
	v, err := d.tryFPE(32, 16, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP32BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP32BE tries to add a field and read 32 bit fixed-point number in big-endian
func (d *D) TryFieldFP32BE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP32BE, sfns...)
}

// FieldFP32BE adds a field and reads 32 bit fixed-point number in big-endian
func (d *D) FieldFP32BE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP32BE, sfns...)
}

// Reader FP64BE

// TryFP64BE tries to read 64 bit fixed-point number in big-endian
func (d *D) TryFP64BE() (float64, error) { return d.tryFPE(64, 32, BigEndian) }

func (d *D) ScalarFP64BE() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryFPE(64, 32, BigEndian)
		s.Actual = v
		return s, err
	}
}

// FP64BE reads 64 bit fixed-point number in big-endian
func (d *D) FP64BE() float64 {
	v, err := d.tryFPE(64, 32, BigEndian)
	if err != nil {
		panic(IOError{Err: err, Op: "FP64BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldFP64BE tries to add a field and read 64 bit fixed-point number in big-endian
func (d *D) TryFieldFP64BE(name string, sfns ...ScalarFn) (float64, error) {
	return d.TryFieldFFn(name, (*D).TryFP64BE, sfns...)
}

// FieldFP64BE adds a field and reads 64 bit fixed-point number in big-endian
func (d *D) FieldFP64BE(name string, sfns ...ScalarFn) float64 {
	return d.FieldFFn(name, (*D).FP64BE, sfns...)
}

// Reader Unary

// TryUnary tries to read unary integer using ov as "one" value
func (d *D) TryUnary(ov uint64) (uint64, error) { return d.tryUnary(ov) }

func (d *D) ScalarUnary(ov uint64) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryUnary(ov)
		s.Actual = v
		return s, err
	}
}

// Unary reads unary integer using ov as "one" value
func (d *D) Unary(ov uint64) uint64 {
	v, err := d.tryUnary(ov)
	if err != nil {
		panic(IOError{Err: err, Op: "Unary", Pos: d.Pos()})
	}
	return v
}

// TryFieldUnary tries to add a field and read unary integer using ov as "one" value
func (d *D) TryFieldUnary(name string, ov uint64, sfns ...ScalarFn) (uint64, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUnary(ov), sfns...)
	if err != nil {
		return 0, err
	}
	return v.ActualU(), err
}

// FieldUnary adds a field and reads unary integer using ov as "one" value
func (d *D) FieldUnary(name string, ov uint64, sfns ...ScalarFn) uint64 {
	v, err := d.TryFieldUnary(name, ov, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "Unary", Pos: d.Pos()})
	}
	return v
}

// Reader UTF8

// TryUTF8 tries to read nBytes bytes UTF8 string
func (d *D) TryUTF8(nBytes int) (string, error) { return d.tryText(nBytes, UTF8BOM) }

func (d *D) ScalarUTF8(nBytes int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryText(nBytes, UTF8BOM)
		s.Actual = v
		return s, err
	}
}

// UTF8 reads nBytes bytes UTF8 string
func (d *D) UTF8(nBytes int) string {
	v, err := d.tryText(nBytes, UTF8BOM)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF8", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF8 tries to add a field and read nBytes bytes UTF8 string
func (d *D) TryFieldUTF8(name string, nBytes int, sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUTF8(nBytes), sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// FieldUTF8 adds a field and reads nBytes bytes UTF8 string
func (d *D) FieldUTF8(name string, nBytes int, sfns ...ScalarFn) string {
	v, err := d.TryFieldUTF8(name, nBytes, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UTF8", Pos: d.Pos()})
	}
	return v
}

// Reader UTF16

// TryUTF16 tries to read nBytes bytes UTF16 string, default big-endian and accepts BOM
func (d *D) TryUTF16(nBytes int) (string, error) { return d.tryText(nBytes, UTF16BOM) }

func (d *D) ScalarUTF16(nBytes int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryText(nBytes, UTF16BOM)
		s.Actual = v
		return s, err
	}
}

// UTF16 reads nBytes bytes UTF16 string, default big-endian and accepts BOM
func (d *D) UTF16(nBytes int) string {
	v, err := d.tryText(nBytes, UTF16BOM)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF16", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF16 tries to add a field and read nBytes bytes UTF16 string, default big-endian and accepts BOM
func (d *D) TryFieldUTF16(name string, nBytes int, sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUTF16(nBytes), sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// FieldUTF16 adds a field and reads nBytes bytes UTF16 string, default big-endian and accepts BOM
func (d *D) FieldUTF16(name string, nBytes int, sfns ...ScalarFn) string {
	v, err := d.TryFieldUTF16(name, nBytes, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UTF16", Pos: d.Pos()})
	}
	return v
}

// Reader UTF16LE

// TryUTF16LE tries to read nBytes bytes UTF16 little-endian string
func (d *D) TryUTF16LE(nBytes int) (string, error) { return d.tryText(nBytes, UTF16LE) }

func (d *D) ScalarUTF16LE(nBytes int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryText(nBytes, UTF16LE)
		s.Actual = v
		return s, err
	}
}

// UTF16LE reads nBytes bytes UTF16 little-endian string
func (d *D) UTF16LE(nBytes int) string {
	v, err := d.tryText(nBytes, UTF16LE)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF16LE", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF16LE tries to add a field and read nBytes bytes UTF16 little-endian string
func (d *D) TryFieldUTF16LE(name string, nBytes int, sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUTF16LE(nBytes), sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// FieldUTF16LE adds a field and reads nBytes bytes UTF16 little-endian string
func (d *D) FieldUTF16LE(name string, nBytes int, sfns ...ScalarFn) string {
	v, err := d.TryFieldUTF16LE(name, nBytes, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UTF16LE", Pos: d.Pos()})
	}
	return v
}

// Reader UTF16BE

// TryUTF16BE tries to read nBytes bytes UTF16 big-endian string
func (d *D) TryUTF16BE(nBytes int) (string, error) { return d.tryText(nBytes, UTF16BE) }

func (d *D) ScalarUTF16BE(nBytes int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryText(nBytes, UTF16BE)
		s.Actual = v
		return s, err
	}
}

// UTF16BE reads nBytes bytes UTF16 big-endian string
func (d *D) UTF16BE(nBytes int) string {
	v, err := d.tryText(nBytes, UTF16BE)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF16BE", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF16BE tries to add a field and read nBytes bytes UTF16 big-endian string
func (d *D) TryFieldUTF16BE(name string, nBytes int, sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUTF16BE(nBytes), sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// FieldUTF16BE adds a field and reads nBytes bytes UTF16 big-endian string
func (d *D) FieldUTF16BE(name string, nBytes int, sfns ...ScalarFn) string {
	v, err := d.TryFieldUTF16BE(name, nBytes, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UTF16BE", Pos: d.Pos()})
	}
	return v
}

// Reader UTF8ShortString

// TryUTF8ShortString tries to read one byte length fixed UTF8 string
func (d *D) TryUTF8ShortString() (string, error) { return d.tryTextLenPrefixed(8, -1, UTF8BOM) }

func (d *D) ScalarUTF8ShortString() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryTextLenPrefixed(8, -1, UTF8BOM)
		s.Actual = v
		return s, err
	}
}

// UTF8ShortString reads one byte length fixed UTF8 string
func (d *D) UTF8ShortString() string {
	v, err := d.tryTextLenPrefixed(8, -1, UTF8BOM)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF8ShortString", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF8ShortString tries to add a field and read one byte length fixed UTF8 string
func (d *D) TryFieldUTF8ShortString(name string, sfns ...ScalarFn) (string, error) {
	return d.TryFieldStrFn(name, (*D).TryUTF8ShortString, sfns...)
}

// FieldUTF8ShortString adds a field and reads one byte length fixed UTF8 string
func (d *D) FieldUTF8ShortString(name string, sfns ...ScalarFn) string {
	return d.FieldStrFn(name, (*D).UTF8ShortString, sfns...)
}

// Reader UTF8ShortStringFixedLen

// TryUTF8ShortStringFixedLen tries to read fixedBytes bytes long one byte length prefixed UTF8 string
func (d *D) TryUTF8ShortStringFixedLen(fixedBytes int) (string, error) {
	return d.tryTextLenPrefixed(8, fixedBytes, UTF8BOM)
}

func (d *D) ScalarUTF8ShortStringFixedLen(fixedBytes int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryTextLenPrefixed(8, fixedBytes, UTF8BOM)
		s.Actual = v
		return s, err
	}
}

// UTF8ShortStringFixedLen reads fixedBytes bytes long one byte length prefixed UTF8 string
func (d *D) UTF8ShortStringFixedLen(fixedBytes int) string {
	v, err := d.tryTextLenPrefixed(8, fixedBytes, UTF8BOM)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF8ShortStringFixedLen", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF8ShortStringFixedLen tries to add a field and read fixedBytes bytes long one byte length prefixed UTF8 string
func (d *D) TryFieldUTF8ShortStringFixedLen(name string, fixedBytes int, sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUTF8ShortStringFixedLen(fixedBytes), sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// FieldUTF8ShortStringFixedLen adds a field and reads fixedBytes bytes long one byte length prefixed UTF8 string
func (d *D) FieldUTF8ShortStringFixedLen(name string, fixedBytes int, sfns ...ScalarFn) string {
	v, err := d.TryFieldUTF8ShortStringFixedLen(name, fixedBytes, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UTF8ShortStringFixedLen", Pos: d.Pos()})
	}
	return v
}

// Reader UTF8Null

// TryUTF8Null tries to read null terminated UTF8 string
func (d *D) TryUTF8Null() (string, error) { return d.tryTextNull(1, UTF8BOM) }

func (d *D) ScalarUTF8Null() func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryTextNull(1, UTF8BOM)
		s.Actual = v
		return s, err
	}
}

// UTF8Null reads null terminated UTF8 string
func (d *D) UTF8Null() string {
	v, err := d.tryTextNull(1, UTF8BOM)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF8Null", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF8Null tries to add a field and read null terminated UTF8 string
func (d *D) TryFieldUTF8Null(name string, sfns ...ScalarFn) (string, error) {
	return d.TryFieldStrFn(name, (*D).TryUTF8Null, sfns...)
}

// FieldUTF8Null adds a field and reads null terminated UTF8 string
func (d *D) FieldUTF8Null(name string, sfns ...ScalarFn) string {
	return d.FieldStrFn(name, (*D).UTF8Null, sfns...)
}

// Reader UTF8NullFixedLen

// TryUTF8NullFixedLen tries to read fixedBytes bytes long null terminated UTF8 string
func (d *D) TryUTF8NullFixedLen(fixedBytes int) (string, error) {
	return d.tryTextNullLen(fixedBytes, UTF8BOM)
}

func (d *D) ScalarUTF8NullFixedLen(fixedBytes int) func(Scalar) (Scalar, error) {
	return func(s Scalar) (Scalar, error) {
		v, err := d.tryTextNullLen(fixedBytes, UTF8BOM)
		s.Actual = v
		return s, err
	}
}

// UTF8NullFixedLen reads fixedBytes bytes long null terminated UTF8 string
func (d *D) UTF8NullFixedLen(fixedBytes int) string {
	v, err := d.tryTextNullLen(fixedBytes, UTF8BOM)
	if err != nil {
		panic(IOError{Err: err, Op: "UTF8NullFixedLen", Pos: d.Pos()})
	}
	return v
}

// TryFieldUTF8NullFixedLen tries to add a field and read fixedBytes bytes long null terminated UTF8 string
func (d *D) TryFieldUTF8NullFixedLen(name string, fixedBytes int, sfns ...ScalarFn) (string, error) {
	v, err := d.TryFieldScalar(name, d.ScalarUTF8NullFixedLen(fixedBytes), sfns...)
	if err != nil {
		return "", err
	}
	return v.ActualStr(), err
}

// FieldUTF8NullFixedLen adds a field and reads fixedBytes bytes long null terminated UTF8 string
func (d *D) FieldUTF8NullFixedLen(name string, fixedBytes int, sfns ...ScalarFn) string {
	v, err := d.TryFieldUTF8NullFixedLen(name, fixedBytes, sfns...)
	if err != nil {
		panic(IOError{Err: err, Name: name, Op: "UTF8NullFixedLen", Pos: d.Pos()})
	}
	return v
}
