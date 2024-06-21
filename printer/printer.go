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

	p.printer.PrintLn("Wilson Vanegas Hernandez")
	p.printer.PrintLn("Nit: 10.777.579-3")
	p.printer.PrintLn("Regimen simplificado")
	p.printer.PrintLn(fmt.Sprint("Fecha: ", receipt.CreatedAt.Format("01/02/2006")))
	p.printer.PrintLn("Direccion: Monteria, Cordoba")
	p.printer.PrintLn("NIT/CC: 22222222 2")
	p.printer.PrintLn(fmt.Sprint("Vendedor: ", receipt.User.Person.Firstname, " ", receipt.User.Person.Lastname))
	p.printer.PrintLn(fmt.Sprint("Habitacion: ", receipt.Room.Number))
	p.printer.PrintLn(fmt.Sprint("Numero de recibo: ", receipt.ID))

	var totalProducts = 0.0
	for _, product := range receipt.Products {
		totalProducts += float64(product.Price) * float64(product.Quantity)
	}

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
	p.printer.Print(fmt.Sprintln("Tiempo total: ", FormatDuration(receipt.TotalTime*time.Second)))
	p.printer.Print(fmt.Sprintln("Precio por tiempo:", ac.FormatMoney(receipt.TotalPrice-receipt.Service.Price-float32(totalProducts)), " COP"))

	// fmt.Println(fmt.Sprintln("Precio por tiempo:", ac.FormatMoney(receipt.TotalPrice-(receipt.Service.Price+float32(totalProducts))), " COP"))

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
	// p.printer.Write(fmt.Sprintf("\x1B", "\x70", "\x00"))
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
	p.printer.PrintLn("Wilson Vanegas Hernandez")
	p.printer.PrintLn("Nit: 10.777.579-3")
	p.printer.PrintLn("Regimen simplificado")
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
	// p.printer.Write(fmt.Sprintf("\x1B", "\x70", "\x00"))
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
	totalPrice float32,
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
	p.printer.PrintLn("Wilson Vanegas Hernández")
	p.printer.PrintLn("Nit: 10.777.579-3")
	p.printer.PrintLn("Regimen simplificado")
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
		var thisTotalService = (service.Price * float32(service.Quantity))

		p.printer.Align(escpos.AlignLeft)
		p.printer.Print(fmt.Sprintln("(", service.Quantity, ")", service.Name))
		p.printer.Align(escpos.AlignRight)
		p.printer.Print(fmt.Sprintln(ac.FormatMoney(thisTotalService), " COP"))
		// fmt.Println(fmt.Sprintln("service total this: ", ac.FormatMoney(thisTotalService), " COP"))
	}
	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Precio despues de hora:")
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(totalPrice-totalService-totalProduct), " COP"))
	// fmt.Println(fmt.Sprintln("total service time: ", ac.FormatMoney(totalPrice-totalService-totalProduct), " COP"))

	// p.printer.Print(fmt.Sprintln("Precio por tiempo:", ac.FormatMoney(receipt.TotalPrice-receipt.Service.Price-float32(totalProducts)), " COP"))

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
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(totalPrice-totalProduct), " COP"))

	p.printer.Align(escpos.AlignLeft)
	p.printer.PrintLn("Total venta del turno:")
	p.printer.Align(escpos.AlignRight)
	p.printer.Print(fmt.Sprintln(ac.FormatMoney(totalPrice), " COP"))

	// fmt.Println(fmt.Sprintln("total de servicios + el tiempo", ac.FormatMoney(totalPrice-totalProduct)))

	// fmt.Println(fmt.Sprintln("total de servicios", ac.FormatMoney(totalPrice-totalProduct)))
	// fmt.Println(fmt.Sprintln("total de productos", ac.FormatMoney(totalProduct)))
	// fmt.Println(fmt.Sprintln("total de price", ac.FormatMoney(totalPrice)))

	p.printer.Size(2, 2)
	p.printer.PrintLn("------------------------")

	p.printer.Size(1, 1)

	p.printer.Feed(2)
	p.printer.Cut()
	p.printer.End()
	p.printer.Close()
}
func FormatDuration(d time.Duration) string {
	minutes := int(d.Minutes())

	// if minutes < 30 {
	// 	return "00:30"
	// }

	if minutes <= 60 {
		return "01:00"
	}

	hours := minutes / 60
	minutes = minutes % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
