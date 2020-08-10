package main

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/jung-kurt/gofpdf"
	"time"
)

type Controller struct {
	beego.Controller
}

func (this *Controller) Get() {
	this.TplName = "form.html"
}

func (this *Controller) Post() {
	timeString := time.Now().Format("Mon Jan 2, 2006 15:04:05")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, timeString)

	blob := new(bytes.Buffer)
	_ = pdf.Output(blob)

	this.Ctx.Output.Header("Content-Type", "application/pdf")
	_, _ = this.Ctx.ResponseWriter.Write(blob.Bytes())
}

func main() {
	beego.Router("/", &Controller{})
	beego.Run()
}
