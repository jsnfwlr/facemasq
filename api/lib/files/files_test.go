package files

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"facemasq/lib/password"
)

func TestWriteOut(t *testing.T) {
	tests, err := GetDir("tests")
	if err != nil {
		t.Error(`Getdir("tests") failed`)
	}
	dstFile := fmt.Sprintf("%s/test.file", tests)
	err = WriteOut(dstFile, "Testing source file")
	if err != nil {
		t.Errorf("error while writing to %s: %v", dstFile, err)
	}

	dstPerm := fmt.Sprintf("%s/perm/test.perm", tests)
	err = WriteOut(dstPerm, "Testing source file")
	if err == nil {
		t.Errorf("should not be able to write to %s", dstPerm)
	}

	dstFail := fmt.Sprintf("%s/fail/test.file", tests)
	err = WriteOut(dstFail, "Testing source file")
	if err == nil {
		t.Errorf("should not be able to write to %s", dstFail)
	}
}

func TestFileExists(t *testing.T) {
	tests, err := GetDir("tests")
	if err != nil {
		t.Error(`Getdir("tests") failed`)
	}
	srcFile := fmt.Sprintf("%s/test.file", tests)
	srcFail := fmt.Sprintf("%s/test.fail", tests)
	srcPerm := fmt.Sprintf("%s/perm/test.perm", tests)
	if FileExists(srcFail) {
		t.Errorf("%s should not exist", srcFail)
	}
	if !FileExists(srcFile) {
		t.Errorf("%s should exist", srcFile)
	}
	if FileExists(srcPerm) {
		t.Errorf("%s should error ", srcPerm)
	}
}

func TestCopy(t *testing.T) {
	tests, err := GetDir("tests")
	if err != nil {
		t.Error(`Getdir("tests") failed`)
	}

	srcFile := fmt.Sprintf("%s/test.file", tests)
	dstFile := fmt.Sprintf("%s/copy.file", tests)
	dstFail := fmt.Sprintf("%s/test.copy", strings.Replace(tests, "tests", "failures", -1))
	srcFail := fmt.Sprintf("%s/test.fail", tests)
	srcPerm := fmt.Sprintf("%s/test.perm", tests)
	srcDir := tests
	srcLink := fmt.Sprintf("%s/test.link", tests)
	dstLink := fmt.Sprintf("%s/copy.link", tests)

	_, err = Copy(srcFile, dstFile)
	if err != nil {
		t.Errorf("Could not copy %s to %s: %v", srcFile, dstFile, err)
	}

	_, err = Copy(srcFile, dstFail)
	if err == nil {
		t.Errorf("Should not be able to copy %s to folder that doesn't exist (%s)", srcFile, dstFail)
	}

	_, err = Copy(srcFail, dstFile)
	if err == nil {
		t.Errorf("Should not be able to copy non-existent file (%s) to %s", srcFail, dstFile)
	}

	_, err = Copy(srcDir, dstFile)
	if err == nil {
		t.Errorf("Should not be able to copy directory (%s) to %s: %v", srcDir, dstFile, err)
	}

	_, err = Copy(srcPerm, dstFile)
	if err == nil {
		t.Errorf("Should not be able to copy file (%s) to %s", srcPerm, dstFile)
	}

	_, err = Copy(srcLink, dstLink)
	if err != nil {
		t.Errorf("Should be able to copy link %s to %s: %v", srcLink, dstLink, err)
	}
}

func TestGetAppRoot(t *testing.T) {
	rootDir, err := GetAppRoot()
	if err != nil {
		t.Error("How can we not find the current working directory?")
	}
	newDir := fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, rootDir, "deleteme")
	_ = os.Mkdir(newDir, 0o777)

	os.Chdir(newDir)
	subDir, err := GetAppRoot()
	if err != nil {
		t.Error("How can we not find the current working directory?")
	}
	if subDir == rootDir {
		t.Errorf("%s and %s should be different", subDir, rootDir)
	}
	err = os.Remove(subDir)
	if err != nil {
		t.Errorf("should be able to remove %s", subDir)
	}

	testDir, err := GetAppRoot()
	if err == nil {
		t.Errorf("Should not be able to get %s as AppRoot", testDir)
	}
	os.Chdir(rootDir)

	mode = "forced"
	forceDir, err := GetAppRoot()
	if err != nil {
		t.Errorf("This should force the rootDir to `/app`: %v", err)
	}
	if forceDir != "/app" {
		t.Errorf("This should force the rootDir to `/app` not %s", forceDir)
	}
}

func TestGetDir(t *testing.T) {
	_, err := GetDir("data")
	if err != nil {
		t.Errorf(`Getdir("data") failed: %v`, err)
	}

	_, err = GetDir("../data")
	if err == nil {
		t.Error(`Getdir("../data") should have failed`)
	}

	// Use a semi-random string to test this one, so we don't trip over any folders that actually exist on the filesystem
	randomDir, _ := password.HashPassword(time.Now().Format("2006-01-02 15:04:05"))
	mode = "forced"
	dir, err := GetDir(randomDir)
	if err == nil {
		t.Errorf("Getdir(`%s`) should have thrown an error saying `%s`could not be found", randomDir, dir)
	}
}
