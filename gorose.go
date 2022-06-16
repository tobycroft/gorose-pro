package gorose

// GOROSE_IMG ...
const GOROSE_IMG = `

 ██████╗  ██████╗ ██████╗  ██████╗ ███████╗███████╗    ██████╗ ██████╗  ██████╗
██╔════╝ ██╔═══██╗██╔══██╗██╔═══██╗██╔════╝██╔════╝    ██╔══██╗██╔══██╗██╔═══██╗
██║  ███╗██║   ██║██████╔╝██║   ██║███████╗█████╗█████╗██████╔╝██████╔╝██║   ██║
██║   ██║██║   ██║██╔══██╗██║   ██║╚════██║██╔══╝╚════╝██╔═══╝ ██╔══██╗██║   ██║
╚██████╔╝╚██████╔╝██║  ██║╚██████╔╝███████║███████╗    ██║     ██║  ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝    ╚═╝     ╚═╝  ╚═╝ ╚═════╝

`

const (
	// VERSION_TEXT ...
	VERSION_TEXT = "\ngolang orm of gorose's version : "
	// VERSION_NO ...
	VERSION_NO = "v1.3.0"
	// VERSION ...
	VERSION = VERSION_TEXT + VERSION_NO + GOROSE_IMG
)

// Open ...
func Open(conf ...interface{}) (engin *Engin, err error) {
	// 驱动engin
	engin, err = NewEngin(conf...)
	if err != nil {
		if engin.GetLogger().EnableErrorLog() {
			engin.GetLogger().Error(err.Error())
		}
		return
	}

	return
}
