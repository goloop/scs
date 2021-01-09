package scs

import "testing"

// TestNew tests New function.
func TestNew(t *testing.T) {
	var test = struct {
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
	pascal, err := New(Pascal, test.example)
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

// TestObjToCamel tests ToCamel method of the object.
func TestObjToCamel(t *testing.T) {
	var test = struct {
		example string
		camel   string
		pascal  string
	}{
		example: "http 2 https convertor",
		camel:   "http2HTTPSConvertor",
		pascal:  "HTTP2HTTPSConvertor",
	}

	pascal, _ := New(Pascal, test.example)
	camel := pascal.ToCamel()
	if r := camel.Value(); r != test.camel {
		t.Errorf("test for Pascal To Camel is failed, "+
			"expected %s but %s", test.camel, r)
	}
}

// TestObjToKebab tests ToKebab method of the object.
func TestObjToKebab(t *testing.T) {
	var test = struct {
		example string
		kebab   string
		snake   string
	}{
		example: "http 2 https convertor",
		kebab:   "http-2-https-convertor",
		snake:   "http_2_https_convertor",
	}

	snake, _ := New(Snake, test.example)
	kebab := snake.ToKebab()
	if r := kebab.Value(); r != test.kebab {
		t.Errorf("test for Snake To Kebab is failed, "+
			"expected %s but %s", test.kebab, r)
	}
}

// TestObjToPascal tests ToPascal method of the object.
func TestObjToPascal(t *testing.T) {
	var test = struct {
		example string
		camel   string
		pascal  string
	}{
		example: "http 2 https convertor",
		camel:   "http2HTTPSConvertor",
		pascal:  "HTTP2HTTPSConvertor",
	}

	camel, _ := New(Camel, test.example)
	pascal := camel.ToPascal()
	if r := pascal.Value(); r != test.pascal {
		t.Errorf("test for Camle To Pascal is failed, "+
			"expected %s but %s", test.pascal, r)
	}
}

// TestObjToSnake tests ToSnake method of the object.
func TestObjToSnake(t *testing.T) {
	var test = struct {
		example string
		kebab   string
		snake   string
	}{
		example: "http 2 https convertor",
		kebab:   "http-2-https-convertor",
		snake:   "http_2_https_convertor",
	}

	kebab, _ := New(Kebab, test.example)
	snake := kebab.ToSnake()
	if r := snake.Value(); r != test.snake {
		t.Errorf("test for Kebab To Snake is failed, "+
			"expected %s but %s", test.snake, r)
	}
}
