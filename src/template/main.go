package template

func Main() {
	da := DisplayArray{
		Displays: []AbstractDisplay{
			ErrorDisplay{Error{"404", "Not Found", "The requested resource could not be found."}},
			ErrorDisplay{Error{"500", "Internal Server Error", "The server encountered an unexpected condition."}},
			ErrorDisplay{Error{"503", "Service Unavailable", "The server is currently unavailable."}},
			CharDisplay{Char{"H"}},
		},
	}

	for da.Iterator().HasNext() {
		display(da.Iterator().Next())
		println()
	}
}

type DisplayArrayIterator struct {
	Arr *DisplayArray
}

type DisplayArray struct {
	Displays []AbstractDisplay
}

func (da *DisplayArray) Iterator() DisplayArrayIterator {
	return DisplayArrayIterator{da}
}

func (dai DisplayArrayIterator) HasNext() bool {
	return len(dai.Arr.Displays) > 0
}

func (dai DisplayArrayIterator) Next() AbstractDisplay {
	display := dai.Arr.Displays[0]
	dai.Arr.Displays = dai.Arr.Displays[1:]
	return display
}

type AbstractDisplay interface {
	open()
	print()
	close()
}

type Error struct {
	code    string
	title   string
	message string
}

type ErrorDisplay struct {
	Error
}

func (ed ErrorDisplay) open() {
	println("Error: " + ed.title + " (" + ed.code + ")")
	println("====================================")
}

func (ed ErrorDisplay) print() {
	println("Message: " + ed.message)
}

func (ed ErrorDisplay) close() {
	println("====================================")
}

func display(ad AbstractDisplay) {
	ad.open()
	ad.print()
	ad.close()
}

type Char struct {
	char string
}

type CharDisplay struct {
	Char
}

func (cd CharDisplay) open() {
	print("<<")
}

func (cd CharDisplay) print() {
	print(cd.char)
}

func (cd CharDisplay) close() {
	println(">>")
}
