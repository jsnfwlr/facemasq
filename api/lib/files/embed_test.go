package files

/*
func getLocalFile() ([]byte, error) {
	libFilesTemplates, err := GetDir("api/lib/files/templates")
	if err != nil {
		return []byte{}, err
	}
	return os.ReadFile(fmt.Sprintf("%s/%s", libFilesTemplates, "dnsmasqExportedFileHeader.tmpl"))
}

func TestGetEmbeddedFileContents(t *testing.T) {
	contents1, err := getLocalFile()
	if err != nil {
		t.Error("Can't read template from filesystem")
	}

	contents2, err := GetEmbeddedFileContents("templates/dnsmasqExportedFileHeader.tmpl")
	if err != nil {
		t.Error("Can't read embedded template")
	}

	if string(contents1) != string(contents2) {
		t.Errorf("Two values don't match\nFile:\n%s<<EOF\nEmbed:\n%s<<EOF", string(contents1), string(contents2))
	}
}

func TestGetEmbeddedFile(t *testing.T) {
	fsFile, err := GetEmbeddedFile("templates/dnsmasqExportedFileHeader.tmpl")
	if err != nil {
		t.Error("Can't get embedded templates/dnsmasqExportedFileHeader.tmpl")
	}
	info, err := fsFile.Stat()
	if err != nil {
		t.Error("Can't get file info for templates/dnsmasqExportedFileHeader.tmpl")
	}
	if info.Name() != "dnsmasqExportedFileHeader.tmpl" {
		t.Errorf("%s does not match dnsmasqExportedFileHeader", info.Name())
	}
}

func TestGetEmbeddedFileSystem(t *testing.T) {
	contents1, err := getLocalFile()
	if err != nil {
		t.Error("Can't read template from filesystem")
	}

	embedFS := GetEmbeddedFileSystem()
	contents2, err := fs.ReadFile(embedFS, "templates/dnsmasqExportedFileHeader.tmpl")
	if err != nil {
		t.Error("Can't read embedded template")
	}

	if string(contents1) != string(contents2) {
		t.Errorf("Two values don't match\nFile:\n%s<<EOF\nEmbed:\n%s<<EOF", string(contents1), string(contents2))
	}
}
*/
