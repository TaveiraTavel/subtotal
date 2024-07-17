package static

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

var banner = "                                                          \n" +
	"               ,--.  ,--------.        ,--.          ,--. \n" +
	" ,---. ,--.,--.|  |-.'--.  .--',---. ,-'  '-. ,--,--.|  | \n" +
	"(  .-' |  ||  || .-. '  |  |  | .-. |'-.  .-'' ,-.  ||  | \n" +
	".-'  `)'  ''  '| `-' |  |  |  ' '-' '  |  |  \\ '-'  ||  | \n" +
	"`----'  `----'  `---'   `--'   `---'   `--'   `--`--'`--' \n" +
	"                           v1.1                           \n"

func ShowBanner() {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println(banner)
	}
}
