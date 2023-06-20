package html2image

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-easy-html-template/pkg/easy_html_template"
)

const (
	NO_IMAGE_SOURCE      = "sem imagem de origem"
	NO_IMAGE_DESTINATION = "sem imagem de destino"
)

type HTML2Image struct {
	source          string
	destinationPath string
	screenWidth     uint64
	imageType       string
	useExec         bool
}

func New() *HTML2Image {
	HTML2Image := &HTML2Image{}
	HTML2Image.useExec = true
	HTML2Image.imageType = "png"
	HTML2Image.screenWidth = 640
	return HTML2Image
}

type HTML2ImageInterface interface {
	GenerateImage() error
	SetSource(string)
	SetDestination(string)
	SetScreenWidth(uint64)
	GetDestination() string
}

func (h *HTML2Image) GenerateImage() error {
	if h.source == "" {
		return errors.New(NO_IMAGE_SOURCE)
	}
	if h.destinationPath == "" {
		return errors.New(NO_IMAGE_DESTINATION)
	}
	return h.generateImageWithExec()
}

func (h *HTML2Image) generateImageWithExec() error {
	err := os.WriteFile(h.destinationPath+".html", []byte(h.source), 0666)
	if err != nil {
		return err
	}
	log.Println("/usr/local/bin/wkhtmltoimage", "--format", h.getImageType(), "--width", "640", "--quality", "70", h.destinationPath+".html", h.destinationPath)
	err = exec.Command("/usr/local/bin/wkhtmltoimage", "--format", h.getImageType(), "--width", "640", "--quality", "70", h.destinationPath+".html", h.destinationPath).Run()
	if err != nil {
		return err
	}
	err = os.Remove(h.destinationPath + ".html")
	if err != nil {
		return err
	}
	stat, err := os.Stat(h.destinationPath)
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		// apaga o arquivo gerado
		err = os.Remove(h.destinationPath)
		if err != nil {
			return err
		}
		return errors.New("imagem gerada com tamanho 0")
	}
	return nil
}

func (h *HTML2Image) SetSource(source string) {
	h.source = source
}

func (h *HTML2Image) SetDestination(destinationFile string) {
	destinationFile = strings.ReplaceAll(destinationFile, "..", "")
	destinationFile = strings.ReplaceAll(destinationFile, "/", "")
	destinationFile = strings.ReplaceAll(destinationFile, "."+h.getImageType(), "")
	h.destinationPath = h.getImageDirectory() + string(filepath.Separator) + destinationFile + "." + h.getImageType()
}

func (h *HTML2Image) SetScreenWidth(screenWidth uint64) {
	h.screenWidth = screenWidth
}

func (h *HTML2Image) getImageType() string {
	if h.imageType == "" {
		return "png"
	}
	return h.imageType
}

func (h *HTML2Image) getImageDirectory() string {
	return utils.GetBaseDirectory("images")
}

func (h *HTML2Image) GetDestination() string {
	return h.destinationPath
}

func (i *HTML2Image) SetUseExec(useExec bool) {
	i.useExec = useExec
}

func (h *HTML2Image) LoadDynamicTemplate(templatePath string, assets, data map[string]string) error {
	tpl, err := easy_html_template.LoadDynamicTemplateWithAssets(templatePath, assets, data)
	if err != nil {
		return err
	}

	h.SetSource(tpl)
	return nil
}
