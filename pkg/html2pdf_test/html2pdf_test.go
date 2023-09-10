package html2pdf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-html2image/pkg/html2pdf"
	"github.com/stretchr/testify/require"
)

func TestPdfGenerationWithExecFromString(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skipping test on github")
	}
	html2PDF := html2pdf.NewHTML2PDF()
	html2PDF.SetSource("<html><body><h1>Teste de execução</h1></body></html>")
	html2PDF.SetDestination("teste.pdf")
	err := html2PDF.GeneratePDF()
	require.NoError(t, err)
}
