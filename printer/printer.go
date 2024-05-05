package printer

import (
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"sync"

	"github.com/alexbrainman/printer"
	"github.com/conejoninja/go-escpos"
)

type ESCPOSPrinter struct {
	printer *escpos.Printer
	once    sync.Once
}

var instance *ESCPOSPrinter

func GetESCPOSPrinter() *ESCPOSPrinter {
	if instance == nil {
		instance = &ESCPOSPrinter{}
	}
	return instance
}

func (p *ESCPOSPrinter) InitPrinter() error {
	var err error
	p.once.Do(func() {
		printerPath, err := printer.Default()
		if err != nil {
			fmt.Println("Default failed:")
		}

		p.printer, err = escpos.NewWindowsPrinterByPath(printerPath) // empty string will do a self discovery
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	return err
}

func (p *ESCPOSPrinter) connect() {
	printerPath, err := printer.Default()
	if err != nil {
		fmt.Println("Default failed:")
	}

	p.printer, err = escpos.NewWindowsPrinterByPath(printerPath) // empty string will do a self discovery
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *ESCPOSPrinter) Print(receipt *receipt_presenter.ReceiptResponse) {
	if p.printer == nil {
		fmt.Println("Printer not initialized.")
		return
	}
	p.connect()

	p.printer.Init()
	p.printer.Smooth(true)

	p.printer.Size(2, 2)
	p.printer.Align(escpos.AlignCenter)
	p.printer.Underline(true)
	p.printer.PrintLn("HOTEL TRIVOLI")
	p.printer.Underline(false)
	p.printer.Size(1, 1)
	p.printer.PrintLn("Monteria, Cordoba")

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Servicio:")
	p.printer.Size(1, 1)
	p.printer.PrintLn(receipt.Service.Name)
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(receipt.Service.Price, ".00 COP"))

	p.printer.Align(escpos.AlignLeft)
	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.PrintLn("Productos:")
	p.printer.PrintLn("")

	p.printer.Size(1, 1)

	p.printer.Align(escpos.AlignLeft)
	for _, product := range receipt.Products {
		p.printer.Align(escpos.AlignLeft)
		p.printer.Print(fmt.Sprintln("(", product.Quantity, ")", product.Name))
		p.printer.Align(escpos.AlignRight)
		p.printer.Print(fmt.Sprintln(product.Price, ".00 COP"))
	}
	p.printer.Font(escpos.FontA)

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total:")
	p.printer.Size(1, 1)
	p.printer.PrintLn(receipt.Service.Name)
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(receipt.TotalPrice, ".00 COP"))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Size(1, 1)
	p.printer.Font(escpos.FontB)
	p.printer.PrintLn(receipt.CreatedAt.Format("01/02/2006"))

	p.printer.Size(2, 2)
	p.printer.Align(escpos.AlignCenter)
	p.printer.Barcode(fmt.Sprintln(receipt.ID), escpos.BarcodeTypeCODE39)

	p.printer.Font(escpos.FontA)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignCenter)
	p.printer.PrintLn("Gracias por confiar en  nosotros!")

	p.printer.Feed(2)
	p.printer.Cut()
	p.printer.End()
	p.printer.Close()
}

type PrintReport struct {
}

func (p *ESCPOSPrinter) PrintReport(
	products []receipt_presenter.ProductResponse,
	totalProduct float32,
	services []entities.Service,
	totalService float32,
) {
	if p.printer == nil {
		fmt.Println("Printer not initialized.")
		return
	}
	p.connect()

	p.printer.Init()
	p.printer.Smooth(true)

	p.printer.Size(2, 2)
	p.printer.Align(escpos.AlignCenter)
	p.printer.Size(1, 1)
	p.printer.PrintLn("Monteria, Cordoba")

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Servicio:")
	p.printer.Size(1, 1)

	for _, service := range services {
		p.printer.PrintLn(service.Name)
		p.printer.Align(escpos.AlignRight)
		p.printer.Print(fmt.Sprintln(service.Price, ".00 COP"))
	}

	p.printer.Align(escpos.AlignLeft)
	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.PrintLn("Productos:")
	p.printer.PrintLn("")

	p.printer.Size(1, 1)

	p.printer.Align(escpos.AlignLeft)
	for _, product := range products {
		p.printer.Align(escpos.AlignLeft)
		p.printer.Print(fmt.Sprintln("(", product.Quantity, ")", product.Name))
		p.printer.Align(escpos.AlignRight)
		p.printer.Print(fmt.Sprintln(product.Price, ".00 COP"))
	}
	p.printer.Font(escpos.FontA)

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Subtotal:")
	p.printer.Size(1, 1)
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(totalProduct, ".00 COP"))
	p.printer.Print(fmt.Sprintln(totalService, ".00 COP"))

	p.printer.PrintLn("Total venta turno:")
	p.printer.Print(fmt.Sprintln(totalProduct+totalService, ".00 COP"))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Size(1, 1)
	p.printer.Font(escpos.FontB)
	// p.printer.PrintLn(receipt.CreatedAt.Format("01/02/2006"))

	p.printer.Size(2, 2)
	p.printer.Align(escpos.AlignCenter)
	// p.printer.Barcode(fmt.Sprintln(receipt.ID), escpos.BarcodeTypeCODE39)

	p.printer.Font(escpos.FontA)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignCenter)
	p.printer.PrintLn("Gracias por confiar en  nosotros!")

	p.printer.Feed(2)
	p.printer.Cut()
	p.printer.End()
	p.printer.Close()
}
