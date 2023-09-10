package common

import "github.com/pericles-luz/go-easy-html-template/pkg/easy_html_template"

const (
	NO_PDF_SOURCE      = "sem html de origem"
	NO_PDF_DESTINATION = "sem path de destino"
)

type FileBase struct {
	source          string
	destinationPath string
	fileType        string
}

func (f *FileBase) SetSource(source string) {
	f.source = source
}

func (f *FileBase) SetDestinationPath(destinationPath string) {
	f.destinationPath = destinationPath
}

func (f *FileBase) SetFileType(fileType string) {
	f.fileType = fileType
}

func (f *FileBase) GetSource() string {
	return f.source
}

func (f *FileBase) GetDestinationPath() string {
	return f.destinationPath
}

func (f *FileBase) GetFileType() string {
	return f.fileType
}

func (f *FileBase) LoadDynamicTemplate(templatePath string, assets, data map[string]string) error {
	tpl, err := easy_html_template.LoadDynamicTemplateWithAssets(templatePath, assets, data)
	if err != nil {
		return err
	}
	f.SetSource(tpl)
	return nil
}
