// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	err := license.SetLicenseKey(licenseKey, `Company Name`)
	if err != nil {
		panic(err)
	}
}

func main() {
	doc := document.New()
	defer doc.Close()
	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum.`

	// single spaced
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(lorem)
	run.AddText(lorem)
	run.AddBreak()

	// double spaced is twice the text height (24 points in this case as the text height is 12 points)
	para = doc.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(24*measurement.Point, wml.ST_LineSpacingRuleAuto)
	run = para.AddRun()
	run.AddText(lorem)
	run.AddText(lorem)
	run.AddBreak()

	if err := doc.Validate(); err != nil {
		log.Fatalf("error during validation: %s", err)
	}
	doc.SaveToFile("line-spacing.docx")
}
