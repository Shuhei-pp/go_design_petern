package adapter

func Main() {
	b := PrintBanner{Banner{"Hello"}}
	b.PrintWeak()
	b.PrintStrong()
}

type Print interface {
	PrintWeak()
	PrintStrong()
}

type Banner struct {
	Str string
}

func (b Banner) ShowWithParen() {
	println("(", b.Str, ")")
}

func (b Banner) ShowWithAster() {
	println("*", b.Str, "*")
}

type PrintBanner struct {
	Banner Banner
}

func (pb PrintBanner) PrintWeak() {
	pb.Banner.ShowWithParen()
}

func (pb PrintBanner) PrintStrong() {
	pb.Banner.ShowWithAster()
}
