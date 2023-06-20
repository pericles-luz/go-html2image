package html2image_test

import (
	"go-html2image/pkg/html2image"
	"os"
	"testing"
	"time"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestImageGeneration(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skipping test on github")
	}
	ig := html2image.New()
	ig.SetSource(htmlSource())
	ig.SetDestination("teste.png")
	ig.SetScreenWidth(640)
	require.NoError(t, ig.GenerateImage())
}

func TestHTML2ImageLoadDynamicTemplate(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skipping test on github")
	}
	assets := make(map[string]string)
	assets["FONT_Nunito_01"] = utils.GetBaseDirectory("templates/assets/fonts") + "/XRXV3I6Li01BKofIO-aBTMnFcQIG.woff2"
	assets["FONT_Nunito_02"] = utils.GetBaseDirectory("templates/assets/fonts") + "/XRXV3I6Li01BKofINeaBTMnFcQ.woff2"
	assets["PNG_ConnectSimulacaoBase"] = utils.GetBaseDirectory("templates/assets/images") + "/connect-simulacao-base.png"
	assets["PNG_ConnectSimulacaoLogo"] = utils.GetBaseDirectory("templates/assets/images") + "/connect-simulacao-logo.png"
	assets["PNG_BV"] = utils.GetBaseDirectory("templates/assets/images") + "/bv.png"

	ig := html2image.New()
	require.NoError(t, ig.LoadDynamicTemplate(utils.GetBaseDirectory("templates")+"/connect-simulacao-fgts.html", assets, map[string]string{}))
	ig.SetDestination("teste.png")
	ig.SetScreenWidth(640)
	require.NoError(t, ig.GenerateImage())
}

func TestHTML2ImageLoadDynamicTemplateFiliacao(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skipping test on github")
	}
	assets := make(map[string]string)
	assets["PNG_LOGOMARCA"] = utils.GetBaseDirectory("templates/assets/images") + "/logo-horizontal.png"
	assets["PNG_ASSINATURA"] = utils.GetBaseDirectory("templates/assets/images") + "/thales-assinatura.png"
	data := make(map[string]string)
	data["NOME"] = "Joaquim de Teste"
	data["CPF"] = "123.456.789-00"
	data["FILIACAO"] = "13-05-1992"
	data["DATA"] = time.Now().Format("02/01/2006")

	ig := html2image.New()
	require.NoError(t, ig.LoadDynamicTemplate(utils.GetBaseDirectory("templates")+"/sindireceita-declaracao-filiacao.html", assets, data))
	ig.SetDestination("teste.png")
	ig.SetScreenWidth(1080)
	require.NoError(t, ig.GenerateImage())
}

func htmlSource() string {
	return `<html>
  <head>
    <meta charset="utf-8"/>
		<style>
/* latin-ext */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 200;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofIO-aBTMnFcQIG.woff2") + `) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}
/* latin */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 200;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofINeaBTMnFcQ.woff2") + `) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
/* latin-ext */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofIO-aBTMnFcQIG.woff2") + `) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}
/* latin */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofINeaBTMnFcQ.woff2") + `) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
/* latin-ext */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 600;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofIO-aBTMnFcQIG.woff2") + `) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}
/* latin */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 600;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofINeaBTMnFcQ.woff2") + `) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
/* latin-ext */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 900;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofIO-aBTMnFcQIG.woff2") + `) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}
/* latin */
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 900;
  font-display: swap;
  src: url(` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/fonts")+"/XRXV3I6Li01BKofINeaBTMnFcQ.woff2") + `) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

    </style>
    <style>
      body {
        background-image:url("https://verz.com.br/lnk/img/connect-simulacao-base.png");
        background-image:url("` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/images")+"/connect-simulacao-base.png") + `");
        background-repeat: no-repeat;
        background-size: 640px;
        width: 640px;
				margin: 0;
      }
      div.saudacao {
        font-family: nunito, sans-serif;
				font-weight: 900;
				font-style: normal;
        font-size: 32px;
        color: #f8b301;
        font-weight: 900;
        text-align: left;
        margin-top: 38.4px;
        margin-left: 320px;
        margin-right: 32px;
        line-height: 1em;
      }
      div.veja {
        font-family: nunito, sans-serif;
				font-weight: 900;
				font-style: normal;
        font-size: 32px;
        color: #f8b301;
        font-weight: 900;
        text-align: left;
        margin-top: 70.4px;
        margin-left: 64px;
        margin-right: 32px;
        line-height: 1em;
      }
      div.confira {
        font-family: nunito, sans-serif;
				font-weight: 600;
				font-style: normal;
        font-size: 17.28;
        color: #fff;
        font-weight: 600;
        font-style: bold;
        text-align: left;
        margin-top: 64px;
        margin-left: 320px;
        margin-right: 32px;
        line-height: 1em;
      }
      div.aproximado {
        font-family: nunito, sans-serif;
        font-size: 15.36px;
        color: #000;
        font-weight: 100;
        font-style: normal;
        text-align: left;
        margin-top: 12.8px;
        margin-left: 64px;
        margin-right: 32px;
        line-height: 1em;
      }
      div.titulo {
        font-weight: 400;
        font-size: 23.68px;
        color: #000;
        font-weight: 100;
        font-style: normal;
        text-align: left;
        margin-top: 51.2px;
        margin-left: 64px;
        margin-right: 32px;
        line-height: 1em;
      }
      span.titulo-negrito {
        font-family: nunito, sans-serif;
				font-weight: 900;
				font-style: normal;
        font-weight: 900;
      }
      div.label {
        font-family: nunito, sans-serif;
				font-weight: 600;
				font-style: normal;
        font-size: 12.8px;
        color: #000;
        font-weight: 100;
        font-style: normal;
        text-align: left;
        text-transform: uppercase;
        margin-top: 51.2px;
        margin-left: 64px;
        margin-right: 32px;
        line-height: 1em;
      }
      div.valor {
        font-style: normal;
        font-size: 32px;
        font-weight: 100;
        color: #000;
        font-style: normal;
        text-align: left;
        text-transform: uppercase;
        margin-top: 6.4px;
        margin-left: 64px;
        margin-right: 32px;
        line-height: 1em;
      }
      div.valor-negrito {
        font-family: nunito, sans-serif;
				font-weight: 900;
				font-style: normal;
        font-size: 32px;
        font-weight: 900;
        color: #000;
        font-style: normal;
        text-align: left;
        text-transform: uppercase;
        margin-top: 6.4px;
        margin-left: 64px;
        margin-right: 32px;
        line-height: 1em;
      }
      hr.separador {
        margin-top: 19.2px;
        margin-left: 64px;
        margin-right: 32px;
        width: 544px;
        border: 0.3px solid  #f8b301;
      }
      img.connect {
        margin-top: 132px;
        margin-left: 64px;
        margin-right: 32px;
        height: 64px;
        border: 0px;
        float: left;
      }
      img.banco {
        margin-top: 132px;
        margin-left: 64px;
        margin-right: 32px;
        height: 64px;
        border: 0px;
        float: right;
      }
    </style>
  </head>
  <body>
    <div class="saudacao">
      Oba! A simulação da sua proposta chegou!
    </div>
    <div class="confira">
      Confira abaixo mais informações e conte com a gente!
    </div>
    <div class="veja">
      João, veja sua simulação:
    </div>
    <div class="aproximado">
      *Valores aproximados simulados na data de hoje.
    </div>
    <div class="titulo">
      Antecipação <span class="titulo-negrito">Saque Aniversário FGTS</span>
    </div>
    <div class="label">
      Saldo
    </div>
    <div class="valor">
      R$ 5.000,00
    </div>
    <hr class="separador">
    <div class="label">
      Valor Liberado
    </div>
    <div class="valor-negrito">
      R$ 3.367,48
    </div>
    <hr class="separador">
    <div class="label">
      Parcelas Antecipadas
    </div>
    <div class="valor">
      10
    </div>
    <hr class="separador">
    <div class="label">
      Taxa de Juros
    </div>
    <div class="valor">
      1,98%
    </div>
    <hr class="separador">
    <img class="connect" src="` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/images")+"/connect-simulacao-logo.png") + `">
    <img class="connect" src="` + html2image.AssetToBase64(utils.GetBaseDirectory("templates/assets/images")+"/bv.png") + `">
  </body>
</html>`
}
