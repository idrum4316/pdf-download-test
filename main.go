package main

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jung-kurt/gofpdf"
	"time"
)

/*
 * Index Controller
 */
type IndexController struct {
	beego.Controller
}
func (this *IndexController) Get() {
	this.TplName = "index.html"
}

/*
 * Form/PDF Controller
 */
type FormController struct {
	beego.Controller
}
func (this *FormController) Get() {
	this.TplName = "form.html"
}
func (this *FormController) Post() {
	current := time.Now()
	time_string := current.Format("Mon Jan 2, 2006 15:04:05")

	width := this.GetString("width")
	height := this.GetString("height")
	resolution := fmt.Sprintf("Screen Resolution: %sx%s", width, height)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, fmt.Sprintf("PDF Generated on %s",time_string))
	pdf.Ln(10)
	pdf.Cell(0, 10, resolution)

	blob := new(bytes.Buffer)
	_ = pdf.Output(blob)

	this.Ctx.Output.Header("Content-Type", "application/pdf")
	_, _ = this.Ctx.ResponseWriter.Write(blob.Bytes())
}

/*
 * PDF Controller
 */
type PdfController struct {
	beego.Controller
}
func (this *PdfController) Get() {
	current := time.Now()
	time_string := current.Format("Mon Jan 2, 2006 15:04:05")

	width := this.GetString("width")
	height := this.GetString("height")
	resolution := fmt.Sprintf("Screen Resolution: %sx%s", width, height)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, fmt.Sprintf("PDF Generated on %s",time_string))
	pdf.Ln(10)
	pdf.Cell(0, 10, resolution)

	blob := new(bytes.Buffer)
	_ = pdf.Output(blob)

	this.Ctx.Output.Header("Content-Type", "application/pdf")
	_, _ = this.Ctx.ResponseWriter.Write(blob.Bytes())
}

/*
 * Main function
 */
func main() {
	beego.Router("/", &IndexController{})
	beego.Router("/pdf", &PdfController{})
	beego.Router("/get-pdf", &FormController{})
	beego.Run()
}