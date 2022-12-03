package macvendor

import (
	"fmt"
	"testing"
	"time"
)

type VendorTestSet struct {
	MAC            string
	ExpectedVendor string
	ExpectedError  string
}

func TestMacVendorObeyLimit(t *testing.T) {
	vendorTestSet := []VendorTestSet{

		{
			MAC:            "E8:A7:30:67:5B:86",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
		{
			MAC:            "E8:A7:30:67:5B",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
		{
			MAC:            "E8:A7:30:67",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
		{
			MAC:            "E8:A7:30",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
		{
			MAC:            "1C:61:B4:97:D1:55",
			ExpectedVendor: "TP-Link Corporation Limited",
			ExpectedError:  "",
		},
		{
			MAC:            "1C:61:B4:97:D1",
			ExpectedVendor: "TP-Link Corporation Limited",
			ExpectedError:  "",
		},
		{
			MAC:            "1C:61:B4:97",
			ExpectedVendor: "TP-Link Corporation Limited",
			ExpectedError:  "",
		},
		{
			MAC:            "1C:61:B4",
			ExpectedVendor: "TP-Link Corporation Limited",
			ExpectedError:  "",
		},
		{
			MAC:            "BE:82:E0:27:24:A6",
			ExpectedVendor: "",
			ExpectedError:  "could not determine vendor of `BE:82:E0:27:24:A6`: Not Found",
		},
	}
	for i := range vendorTestSet {
		vendor, err := Lookup(vendorTestSet[i].MAC)
		if err != nil && err.Error() != vendorTestSet[i].ExpectedError {
			t.Error(err)
		} else if vendor != vendorTestSet[i].ExpectedVendor {
			t.Errorf("Vendor was expected to be `%s` - got `%s`", vendorTestSet[i].ExpectedVendor, vendor)
		}
		if i < len(vendorTestSet)-1 {
			time.Sleep(2500 * time.Millisecond)
		}
	}
}

func TestMacVendorBreakLimit(t *testing.T) {
	time.Sleep(2500 * time.Millisecond)
	vendorTestSet := []VendorTestSet{

		{
			MAC:            "E8:A7:30:67:5B:86",
			ExpectedVendor: "Apple, Inc.",
			ExpectedError:  "",
		},
		{
			MAC:            "E8:A7:30:67:5B",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `E8:A7:30:67:5B`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "E8:A7:30:67",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `E8:A7:30:67`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "E8:A7:30",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `E8:A7:30`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "1C:61:B4:97:D1:55",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `1C:61:B4:97:D1:55`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "1C:61:B4:97:D1",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `1C:61:B4:97:D1`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "1C:61:B4:97",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `1C:61:B4:97`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "1C:61:B4",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `1C:61:B4`: Too Many Requests - locked for %f seconds", LockTime),
		},
		{
			MAC:            "BE:82:E0:27:24:A6",
			ExpectedVendor: "",
			ExpectedError:  fmt.Sprintf("could not determine vendor of `BE:82:E0:27:24:A6`: Too Many Requests - locked for %f seconds", LockTime),
		},
	}

	for i := range vendorTestSet {
		vendor, err := Lookup(vendorTestSet[i].MAC)
		if err != nil && err.Error() != vendorTestSet[i].ExpectedError {
			t.Error(err)
		} else if vendor != vendorTestSet[i].ExpectedVendor {
			t.Errorf("Vendor was expected to be `%s` - got `%s`", vendorTestSet[i].ExpectedVendor, vendor)
		}
	}
	time.Sleep(2500 * time.Millisecond)
	vendor, err := Lookup(vendorTestSet[0].MAC)
	if err != nil && err.Error() != vendorTestSet[0].ExpectedError {
		t.Error(err)
	} else if vendor != vendorTestSet[0].ExpectedVendor {
		t.Errorf("Vendor was expected to be `%s` - got `%s`", vendorTestSet[0].ExpectedVendor, vendor)
	}
}
