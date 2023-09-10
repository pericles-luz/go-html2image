package html2pdf

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-html2image/internal/common"
)

type HTML2PDF struct {
	common.FileBase
	marginTop    string
	PDFDirectory string
}

func NewHTML2PDF() *HTML2PDF {
	result := &HTML2PDF{}
	result.SetFileType(".pdf")
	return result
}

func (h *HTML2PDF) GeneratePDF() error {
	if h.GetSource() == "" {
		return errors.New(common.NO_PDF_SOURCE)
	}
	if h.GetDestinationPath() == "" {
		return errors.New(common.NO_PDF_DESTINATION)
	}
	return h.GeneratePdfWithExec()
}

func (h *HTML2PDF) GeneratePdfWithExec() error {
	err := os.WriteFile(h.GetDestinationPath()+".html", []byte(h.GetSource()), 0666)
	if err != nil {
		return err
	}
	file := "/usr/bin/wkhtmltopdf"
	if !utils.FileExists(file) {
		file = "/usr/local/bin/wkhtmltopdf"
	}
	if !utils.FileExists(file) {
		return errors.New("wkhtmltopdf não encontrado")
	}
	err = exec.Command(
		file,
		"--footer-font-size", "7",
		"--footer-center", "[page]/[topage]",
		"--margin-top", h.getMarginTop(),
		"--encoding", "utf-8",
		h.GetDestinationPath()+".html", h.GetDestinationPath()+h.GetFileType()).Run()
	if err != nil {
		return err
	}
	err = os.Remove(h.GetDestinationPath() + ".html")
	if err != nil {
		return err
	}
	stat, err := os.Stat(h.GetDestinationPath() + h.GetFileType())
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		// apaga o arquivo gerado
		err = os.Remove(h.GetDestinationPath() + h.GetFileType())
		if err != nil {
			return err
		}
		return errors.New("pdf gerado com tamanho 0")
	}
	return nil
}

func (h *HTML2PDF) SetDestination(destinationFile string) {
	destinationFile = strings.ReplaceAll(destinationFile, "..", "")
	destinationFile = strings.ReplaceAll(destinationFile, "/", "")
	destinationFile = strings.ReplaceAll(destinationFile, h.GetFileType(), "")
	h.SetDestinationPath(h.getPDFDirectory() + string(filepath.Separator) + destinationFile)
}

func (h *HTML2PDF) SetMarginTop(marginTop string) {
	h.marginTop = marginTop
}

func (h *HTML2PDF) SetPDFDirectory(pdfDirectory string) {
	h.PDFDirectory = pdfDirectory
}

func (h *HTML2PDF) getMarginTop() string {
	if h.marginTop == "" {
		return "1cm"
	}
	return h.marginTop
}

func (h *HTML2PDF) getPDFDirectory() string {
	if h.PDFDirectory != "" {
		return h.PDFDirectory
	}
	path := utils.GetBaseDirectory("pdf")
	return path
}
