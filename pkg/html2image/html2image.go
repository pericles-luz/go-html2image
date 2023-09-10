package html2image

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-html2image/internal/common"
)

const (
	NO_IMAGE_SOURCE      = "sem imagem de origem"
	NO_IMAGE_DESTINATION = "sem imagem de destino"
)

type HTML2Image struct {
	common.FileBase
	imageDirectory string
	screenWidth    uint64
	imageType      string
	useExec        bool
}

func NewHTML2Imge() *HTML2Image {
	result := &HTML2Image{}
	result.useExec = true
	result.SetFileType("png")
	result.screenWidth = 640
	return result
}

func (h *HTML2Image) GenerateImage() error {
	if h.GetSource() == "" {
		return errors.New(NO_IMAGE_SOURCE)
	}
	if h.GetDestinationPath() == "" {
		return errors.New(NO_IMAGE_DESTINATION)
	}
	return h.generateImageWithExec()
}

func (h *HTML2Image) generateImageWithExec() error {
	err := os.WriteFile(h.GetDestinationPath()+".html", []byte(h.GetSource()), 0666)
	if err != nil {
		return err
	}
	log.Println("/usr/local/bin/wkhtmltoimage", "--format", h.getImageType(), "--width", "640", "--quality", "70", h.GetDestinationPath()+".html", h.GetDestinationPath())
	err = exec.Command("/usr/local/bin/wkhtmltoimage", "--format", h.getImageType(), "--width", "640", "--quality", "70", h.GetDestinationPath()+".html", h.GetDestinationPath()).Run()
	if err != nil {
		return err
	}
	err = os.Remove(h.GetDestinationPath() + ".html")
	if err != nil {
		return err
	}
	stat, err := os.Stat(h.GetDestinationPath())
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		// apaga o arquivo gerado
		err = os.Remove(h.GetDestinationPath())
		if err != nil {
			return err
		}
		return errors.New("imagem gerada com tamanho 0")
	}
	return nil
}

func (h *HTML2Image) SetDestination(destinationFile string) {
	destinationFile = strings.ReplaceAll(destinationFile, "..", "")
	destinationFile = strings.ReplaceAll(destinationFile, "/", "")
	destinationFile = strings.ReplaceAll(destinationFile, "."+h.getImageType(), "")
	h.SetDestinationPath(h.getImageDirectory() + string(filepath.Separator) + destinationFile + "." + h.getImageType())
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

func (h *HTML2Image) SetImageDirectory(imageDirectory string) {
	h.imageDirectory = imageDirectory
}

func (h *HTML2Image) getImageDirectory() string {
	if h.imageDirectory != "" {
		return h.imageDirectory
	}
	return utils.GetBaseDirectory("images")
}

func (h *HTML2Image) GetDestination() string {
	return h.GetDestinationPath()
}

func (i *HTML2Image) SetUseExec(useExec bool) {
	i.useExec = useExec
}
