package metadata

import (
	"testing"

	"github.com/web_test_launcher/util/bazel"
)

const (
	allFields      = "__main__/metadata/testdata/all-fields.json"
	chromeLinux    = "__main__/metadata/testdata/chrome-linux.json"
	androidBrowser = "__main__/metadata/testdata/android-browser-gingerbread-nexus-s.json"
	fakeBrowser    = "__main__/metadata/testdata/android-browser-gingerbread-nexus-s_faked_on_chrome-linux.json"
)

func TestFromFile(t *testing.T) {
	f, err := bazel.Runfile(allFields)
	if err != nil {
		t.Fatal(err)
	}
	file, err := FromFile(f)
	if err != nil {
		t.Fatal(err)
	}

	expected := Metadata{
		FormFactor:      "PHONE",
		BrowserName:     "chrome",
		Environment:     "chromeos",
		BrowserLabel:    "//testing/web/browsers:figaro",
		TestLabel:       "//testing/web/launcher:tests",
		CropScreenshots: true,
		RecordVideo:     "always",
	}

	if !equals(expected, file) {
		t.Errorf("Got %+v, expected %+v", file, expected)
	}
}

func TestMergeFromFile(t *testing.T) {
	f1, err := bazel.Runfile(chromeLinux)
	if err != nil {
		t.Fatal(err)
	}
	cl, err := FromFile(f1)
	if err != nil {
		t.Fatal(err)
	}

	f2, err := bazel.Runfile(androidBrowser)
	if err != nil {
		t.Fatal(err)
	}
	ab, err := FromFile(f2)
	if err != nil {
		t.Fatal(err)
	}

	f3, err := bazel.Runfile(fakeBrowser)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := FromFile(f3)
	if err != nil {
		t.Fatal(err)
	}

	merged := Merge(cl, ab)

	if !equals(merged, fb) {
		t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", cl, ab, merged, fb)
	}
}

func TestMerge(t *testing.T) {
	type testCase struct {
		input1 Metadata
		input2 Metadata
		result Metadata
	}

	testCases := []testCase{
		testCase{
			Metadata{FormFactor: "PHONE"},
			Metadata{FormFactor: "TABLET"},
			Metadata{FormFactor: "TABLET"},
		},
		testCase{
			Metadata{FormFactor: "PHONE"},
			Metadata{FormFactor: ""},
			Metadata{FormFactor: "PHONE"},
		},
		testCase{
			Metadata{BrowserName: "chrome"},
			Metadata{BrowserName: "firefox"},
			Metadata{BrowserName: "firefox"},
		},
		testCase{
			Metadata{BrowserName: "chrome"},
			Metadata{BrowserName: ""},
			Metadata{BrowserName: "chrome"},
		},
		testCase{
			Metadata{Environment: "linux"},
			Metadata{Environment: "android"},
			Metadata{Environment: "android"},
		},
		testCase{
			Metadata{Environment: "linux"},
			Metadata{Environment: ""},
			Metadata{Environment: "linux"},
		},
		testCase{
			Metadata{BrowserLabel: "//testing/web/browsers:figaro"},
			Metadata{BrowserLabel: "//testing/web/browsers:murphy"},
			Metadata{BrowserLabel: "//testing/web/browsers:murphy"},
		},
		testCase{
			Metadata{BrowserLabel: "//testing/web/browsers:figaro"},
			Metadata{BrowserLabel: ""},
			Metadata{BrowserLabel: "//testing/web/browsers:figaro"},
		},
		testCase{
			Metadata{TestLabel: "//testing/web/browsers:figaro"},
			Metadata{TestLabel: "//testing/web/browsers:murphy"},
			Metadata{TestLabel: "//testing/web/browsers:murphy"},
		},
		testCase{
			Metadata{TestLabel: "//testing/web/browsers:figaro"},
			Metadata{TestLabel: ""},
			Metadata{TestLabel: "//testing/web/browsers:figaro"},
		},
		testCase{
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: false},
			Metadata{CropScreenshots: false},
		},
		testCase{
			Metadata{CropScreenshots: false},
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: true},
		},
		testCase{
			Metadata{CropScreenshots: false},
			Metadata{CropScreenshots: nil},
			Metadata{CropScreenshots: false},
		},
		testCase{
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: nil},
			Metadata{CropScreenshots: true},
		},
	}

	for _, tc := range testCases {
		a := Merge(tc.input1, tc.input2)
		if !equals(a, tc.result) {
			t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, a, tc.result)
		}
	}
}

func equals(e, a Metadata) bool {
	return jsonEquals(e.Capabilities, a.Capabilities) &&
		e.FormFactor == a.FormFactor &&
		e.BrowserName == a.BrowserName &&
		e.Environment == a.Environment &&
		e.TestLabel == a.TestLabel &&
		e.CropScreenshots == a.CropScreenshots &&
		e.RecordVideo == a.RecordVideo
}

func jsonEquals(e, v interface{}) bool {
	switch te := e.(type) {
	case []interface{}:
		tv, ok := v.([]interface{})
		return ok && sliceEquals(te, tv)
	case map[string]interface{}:
		tv, ok := v.(map[string]interface{})
		return ok && mapEquals(te, tv)
	default:
		return e == v
	}
}

func sliceEquals(e, v []interface{}) bool {
	if len(e) != len(v) {
		return false
	}
	for i := 0; i < len(e); i++ {
		if !jsonEquals(e[i], v[i]) {
			return false
		}
	}
	return true
}

func mapEquals(e, v map[string]interface{}) bool {
	if len(e) != len(v) {
		return false
	}
	for ek, ev := range e {
		if vv, ok := v[ek]; !ok || !jsonEquals(ev, vv) {
			return false
		}
	}
	return true
}
