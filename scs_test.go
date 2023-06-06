package scs

import "testing"

// TestVersion tests Version function.
func TestVersion(t *testing.T) {
	if v := Version(); v != "v"+version {
		t.Error("incorrect  Version function")
	}
}

// TestNew tests New function.
func TestNew(t *testing.T) {
	test := struct {
		example string
		camel   string
		kebab   string
		pascal  string
		snake   string
	}{
		example: "http 2 https convertor",
		camel:   "http2HTTPSConvertor",
		kebab:   "http-2-https-convertor",
		pascal:  "HTTP2HTTPSConvertor",
		snake:   "http_2_https_convertor",
	}

	// Incorrect
	wrong, err := New(10)
	if err == nil {
		t.Error("there must be an error")
	}

	if wrong.do("a-b-c") != "a-b-c" {
		t.Error("incorrect do method")
	}

	// Camel
	camel, err := New(Camel, test.example)
	if err != nil {
		t.Error(err)
	}

	if r := camel.Value(); r != test.camel {
		t.Errorf("test for New(Camel) is failed, "+
			"expected %s but %s", test.camel, r)
	}

	// Kebab
	kebab, err := New(Kebab, test.example)
	if err != nil {
		t.Error(err)
	}

	if r := kebab.Value(); r != test.kebab {
		t.Errorf("test for New(Kebab) is failed, "+
			"expected %s but %s", test.kebab, r)
	}

	// Pascal
	pascal, err := New(Pascal)
	pascal.Eat(test.example)
	if err != nil {
		t.Error(err)
	}

	if r := pascal.Value(); r != test.pascal {
		t.Errorf("test for New(Pascal) is failed, "+
			"expected %s but %s", test.pascal, r)
	}

	// Snake
	snake, err := New(Snake, test.example)
	if err != nil {
		t.Error(err)
	}

	if r := snake.Value(); r != test.snake {
		t.Errorf("test for New(Snake) is failed, "+
			"expected %s but %s", test.snake, r)
	}
}

// TestObjIsValid tests IsValid method of the object.
func TestObjIsValid(t *testing.T) {
	// Incorrect
	wrongNew, err := New(0, "hello world")
	if err == nil {
		t.Error("there must be an error")
	}

	if wrongNew.IsValid() {
		t.Error("test for IsValid is failed, " +
			"expected false but true")
	}

	wrongRaw := &StringCaseStyle{}
	if wrongRaw.IsValid() {
		t.Error("test for IsValid is failed, " +
			"expected false but true")
	}

	// Valid
	obj, err := New(Camel, "hello world")
	if err != nil {
		t.Error(err)
	}

	if !obj.IsValid() {
		t.Error("test for IsValid is failed, " +
			"expected true but false")
	}
}

// TestObjPascalCopyTo tests Pascal -> CopyTo* method of the object.
func TestObjPascalCopyTo(t *testing.T) {
	basic, _ := New(Pascal, "http 2 https convertor")
	tests := []struct {
		fn     func() (*StringCaseStyle, error)
		result string
	}{
		{basic.CopyToCamel, "http2HTTPSConvertor"},
		{basic.CopyToKebab, "http-2-https-convertor"},
		{basic.CopyToPascal, "HTTP2HTTPSConvertor"},
		{basic.CopyToSnake, "http_2_https_convertor"},
	}

	for _, test := range tests {
		obj, err := test.fn()
		if err != nil {
			t.Error(err)
		}

		if r := obj.Value(); r != test.result {
			t.Errorf("expected %s but %s", test.result, r)
		}
	}
}

// TestObjPascalToCamel tests Pascal -> ToPascal method of the object.
func TestObjPascalToCamel(t *testing.T) {
	expected := "http2HTTPSConvertor"

	obj, _ := New(Pascal, "http 2 https convertor")
	err := obj.ToCamel()
	if err != nil {
		t.Error(err)
	}

	if r := obj.Value(); r != expected {
		t.Errorf("expected %s but %s", expected, r)
	}
}

// TestObjSnakeCopyTo tests Snake -> CopyTo* method of the object.
func TestObjSnakeCopyTo(t *testing.T) {
	basic, _ := New(Snake, "http 2 https convertor")
	tests := []struct {
		fn     func() (*StringCaseStyle, error)
		result string
	}{
		{basic.CopyToCamel, "http2HTTPSConvertor"},
		{basic.CopyToKebab, "http-2-https-convertor"},
		{basic.CopyToPascal, "HTTP2HTTPSConvertor"},
		{basic.CopyToSnake, "http_2_https_convertor"},
	}

	for _, test := range tests {
		obj, err := test.fn()
		if err != nil {
			t.Error(err)
		}

		if r := obj.Value(); r != test.result {
			t.Errorf("expected %s but %s", test.result, r)
		}
	}
}

// TestObjSnakeToKebab tests Snake -> ToKebab method of the object.
func TestObjSnakeToKebab(t *testing.T) {
	expected := "http-2-https-convertor"

	obj, _ := New(Snake, "http 2 https convertor")
	err := obj.ToKebab()
	if err != nil {
		t.Error(err)
	}

	if r := obj.Value(); r != expected {
		t.Errorf("expected %s but %s", expected, r)
	}
}

// TestObjCamelToCopy tests Camel -> CopyTo* method of the object.
func TestObjCamelCopyTo(t *testing.T) {
	basic, _ := New(Camel, "http 2 https convertor")
	tests := []struct {
		fn     func() (*StringCaseStyle, error)
		result string
	}{
		{basic.CopyToCamel, "http2HTTPSConvertor"},
		{basic.CopyToKebab, "http-2-https-convertor"},
		{basic.CopyToPascal, "HTTP2HTTPSConvertor"},
		{basic.CopyToSnake, "http_2_https_convertor"},
	}

	for _, test := range tests {
		obj, err := test.fn()
		if err != nil {
			t.Error(err)
		}

		if r := obj.Value(); r != test.result {
			t.Errorf("expected %s but %s", test.result, r)
		}
	}
}

// TestObjCamleToPascal tests Camel -> ToPascal method of the object.
func TestObjCamelToPascal(t *testing.T) {
	expected := "HTTP2HTTPSConvertor"

	obj, _ := New(Camel, "http 2 https convertor")
	err := obj.ToPascal()
	if err != nil {
		t.Error(err)
	}

	if r := obj.Value(); r != expected {
		t.Errorf("expected %s but %s", expected, r)
	}
}

// TestObjKebabCopyTo tests Kebab -> CopyTo method of the object.
func TestObjKebabCopyTo(t *testing.T) {
	basic, _ := New(Kebab, "http 2 https convertor")
	tests := []struct {
		fn     func() (*StringCaseStyle, error)
		result string
	}{
		{basic.CopyToCamel, "http2HTTPSConvertor"},
		{basic.CopyToKebab, "http-2-https-convertor"},
		{basic.CopyToPascal, "HTTP2HTTPSConvertor"},
		{basic.CopyToSnake, "http_2_https_convertor"},
	}

	for _, test := range tests {
		obj, err := test.fn()
		if err != nil {
			t.Error(err)
		}

		if r := obj.Value(); r != test.result {
			t.Errorf("expected %s but %s", test.result, r)
		}
	}
}

// TestObjKebabToSnake tests Kebab -> ToSnake method of the object.
func TestObjKebabToSnake(t *testing.T) {
	expected := "http_2_https_convertor"

	obj, _ := New(Kebab, "http 2 https convertor")
	err := obj.ToSnake()
	if err != nil {
		t.Error(err)
	}

	if r := obj.Value(); r != expected {
		t.Errorf("expected %s but %s", expected, r)
	}
}

// TestObjIsCamel tests IsCamel method of the object.
func TestObjIsCamel(t *testing.T) {
	camel, _ := New(Camel)
	if !camel.IsCamel() {
		t.Error("test for IsCamel() is failed, expected true but false")
	}
}

// TestObjIsKebab tests IsKebab method of the object.
func TestObjIsKebab(t *testing.T) {
	kebab, _ := New(Kebab)
	if !kebab.IsKebab() {
		t.Error("test for IsKebab() is failed, expected true but false")
	}
}

// TestObjIsPascal tests IsPascal method of the object.
func TestObjIsPascal(t *testing.T) {
	pascal, _ := New(Pascal)
	if !pascal.IsPascal() {
		t.Error("test for IsPascal() is failed, expected true but false")
	}
}

// TestObjIsSnake tests IsSnake method of the object.
func TestObjIsSnake(t *testing.T) {
	snake, _ := New(Snake)
	if !snake.IsSnake() {
		t.Error("test for IsSnake() is failed, expected true but false")
	}
}

// TestObjEat tests Set method of the object.
func TestObjEat(t *testing.T) {
	snake, _ := New(Snake)
	if snake.Eat("hello world") != "hello_world" {
		t.Error("conversion failed")
	}
}

// TestObjSet tests Set method of the object.
func TestObjSet(t *testing.T) {
	snake, _ := New(Snake)
	if snake.Set("hello world").Value() != "hello_world" {
		t.Error("conversion failed")
	}
}
