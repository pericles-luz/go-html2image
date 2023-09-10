package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/pericles-luz/go-base/pkg/utils"
)

type HTML2PDF struct {
	source          string
	destinationPath string
	marginTop       string
}

func (h *HTML2PDF) SetSource(source string) {
	h.source = source
}

func (h *HTML2PDF) GeneratePdfWithExec() error {
	err := os.WriteFile(h.destinationPath+".html", []byte(h.source), 0666)
	if err != nil {
		sentry.CaptureException(err)
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
		h.destinationPath+".html", h.destinationPath).Run()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	err = os.Remove(h.destinationPath + ".html")
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	stat, err := os.Stat(h.destinationPath)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	if stat.Size() == 0 {
		// apaga o arquivo gerado
		err = os.Remove(h.destinationPath)
		if err != nil {
			sentry.CaptureException(err)
			return err
		}
		return errors.New("pdf gerado com tamanho 0")
	}
	return nil
}

func (h *HTML2PDF) SetDestination(destinationFile string) {
	destinationFile = strings.ReplaceAll(destinationFile, "..", "")
	destinationFile = strings.ReplaceAll(destinationFile, "/", "")
	destinationFile = strings.ReplaceAll(destinationFile, ".pdf", "")
	h.destinationPath = h.getPdfDirectory() + string(filepath.Separator) + destinationFile + ".pdf"
}

func (h *HTML2PDF) SetMarginTop(marginTop string) {
	h.marginTop = marginTop
}

func (h *HTML2PDF) getMarginTop() string {
	if h.marginTop == "" {
		return "1cm"
	}
	return h.marginTop
}

func (h *HTML2PDF) getPdfDirectory() string {
	return ""
	// path := GetBaseDirectory("pdf")
	// return path
}
