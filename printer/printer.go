package printer

import (
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"sync"
	"time"

	"github.com/alexbrainman/printer"
	"github.com/conejoninja/go-escpos"
	"github.com/leekchan/accounting"
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
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	p.connect()

	p.printer.Init()
	p.printer.Smooth(true)

	p.printer.Size(1, 1)
	p.printer.PrintLn(fmt.Sprint("Fecha: ", receipt.CreatedAt.Format("01/02/2006")))
	p.printer.PrintLn("Direccion: Monteria, Cordoba")
	p.printer.PrintLn("NIT/CC: 22222222 2")
	p.printer.PrintLn(fmt.Sprint("Vendedor: ", receipt.User.Person.Firstname, " ", receipt.User.Person.Lastname))
	p.printer.PrintLn(fmt.Sprint("Habitacion: ", receipt.Room.Number))
	p.printer.PrintLn(fmt.Sprint("Numero de recibo: ", receipt.ID))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Servicio:")
	p.printer.Size(1, 1)
	p.printer.PrintLn(receipt.Service.Name)
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(receipt.Service.Price), " COP"))
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Servicio:")
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(receipt.TotalTime))

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
		p.printer.Print(fmt.Sprintln(ac.FormatMoney(product.Price*float32(product.Quantity)), " COP"))
	}
	p.printer.Font(escpos.FontA)

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total:")
	p.printer.Size(1, 1)
	p.printer.PrintLn(receipt.Service.Name)
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(receipt.TotalPrice), " COP"))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Size(1, 1)

	p.printer.Feed(2)
	p.printer.Cut()
	p.printer.End()
	p.printer.Close()
}

func (p *ESCPOSPrinter) PrintIndividual(receipt *receipt_presenter.IndividualReceiptResponse) {
	if p.printer == nil {
		fmt.Println("Printer not initialized.")
		return
	}
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	p.connect()

	p.printer.Init()
	p.printer.Smooth(true)

	p.printer.Size(1, 1)
	p.printer.PrintLn(fmt.Sprint("Fecha: ", receipt.CreatedAt.Format("01/02/2006")))
	p.printer.PrintLn("Direccion: Monteria, Cordoba")
	p.printer.PrintLn("NIT/CC: 22222222 2")
	p.printer.PrintLn(fmt.Sprint("Vendedor: ", receipt.User.Person.Firstname, " ", receipt.User.Person.Lastname))
	p.printer.PrintLn(fmt.Sprint("Numero de recibo: ", receipt.ID))

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
		p.printer.Print(fmt.Sprintln(ac.FormatMoney(product.Price*float32(product.Quantity)), " COP"))
	}
	p.printer.Font(escpos.FontA)

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total:")
	p.printer.Size(1, 1)
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(receipt.TotalPrice), " COP"))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Size(1, 1)

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
	services []receipt_presenter.ServiceResponse,
	totalService float32,
	user entities.User,
	createdAt time.Time,
) {
	if p.printer == nil {
		fmt.Println("Printer not initialized.")
		return
	}
	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	p.connect()

	p.printer.Init()
	p.printer.Smooth(true)

	p.printer.Size(1, 1)
	p.printer.PrintLn(fmt.Sprint("Fecha: ", createdAt.Format("01/02/2006")))
	p.printer.PrintLn("Direccion: Monteria, Cordoba")
	p.printer.PrintLn("NIT/CC: 22222222 2")
	p.printer.PrintLn(fmt.Sprint("Vendedor: ", user.Person.Firstname, " ", user.Person.Lastname))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Servicios:")
	p.printer.Size(1, 1)

	for _, service := range services {
		p.printer.Align(escpos.AlignLeft)
		p.printer.Print(fmt.Sprintln("(", service.Quantity, ")", service.Name))
		p.printer.Align(escpos.AlignRight)
		p.printer.Print(fmt.Sprintln(ac.FormatMoney(service.Price*float32(service.Quantity)), " COP"))
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

		p.printer.Print(fmt.Sprintln(ac.FormatMoney(product.Price*float32(product.Quantity)), " COP"))
	}

	p.printer.Font(escpos.FontA)

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Subtotal:")
	p.printer.Size(1, 1)

	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total de productos:")
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(totalProduct), " COP"))

	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total de servicios:")
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(totalService), " COP"))

	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total venta del turno:")
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(totalProduct+totalService), " COP"))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")

	p.printer.Size(1, 1)

	p.printer.Feed(2)
	p.printer.Cut()
	p.printer.End()
	p.printer.Close()
}
