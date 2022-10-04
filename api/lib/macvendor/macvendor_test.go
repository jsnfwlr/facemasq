package macvendor

import (
	"testing"
	"time"
)

type VendorTestSet struct {
	MAC            string
	ExpectedVendor string
	ExpectedError  string
}

func TestMacVendor(t *testing.T) {
	vendorTestSet := []VendorTestSet{
		{
			MAC:            "BE:82:E0:2E:24:A6",
			ExpectedVendor: "",
			ExpectedError:  "could not determine vendor of `BE:82:E0:2E:24:A6`",
		},
		{
			MAC:            "E8:A7:30:68:5B:86",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
		{
			MAC:            "E8:A7:30:68",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
	}
	for i := range vendorTestSet {
		vendor, err := Lookup(vendorTestSet[i].MAC)
		if err != nil && err.Error() != vendorTestSet[i].ExpectedError {
			t.Error(err)
			continue
		}
		if vendor != vendorTestSet[i].ExpectedVendor {
			t.Errorf("Vendor was expected to be `%s` - got `%s`", vendorTestSet[i].ExpectedVendor, vendor)
		}
		if i < len(vendorTestSet)-1 {
			time.Sleep(2 * time.Second)
		}
	}
}
