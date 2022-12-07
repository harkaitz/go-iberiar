package main
import (
	"github.com/harkaitz/go-iberiar"
	"github.com/pborman/getopt/v2"
	"log"
	"os"
	"fmt"
	"bufio"
)

const help string =
`Usage: iberiar OPTS... TXT...

Convert texts in basque to iberian script.

    -h           : Print this help.
    -s SILLABARY : Choose sillabary.
    -t TEMPLATE  : Choose SVG template to use.
    -l           : Print latin code.
    -S           : Print SVG.

Sillabaries:

    ne-nd : North East Non Dual.
    ne    : North East dual.

Templates:

    empty : Empty template.
`
const copyrightLine string =
`Bug reports, feature requests to gemini|https://harkadev.com/oss
Copyright (c) 2022 Harkaitz Agirre, harkaitz.aguirre@gmail.com`


func main() {

	hFlag := getopt.Bool('h')
	sFlag := getopt.String('s', "")
	tFlag := getopt.String('t', "")
	SFlag := getopt.Bool('S');
	lFlag := getopt.Bool('l');

	/* Parse command line arguments. */
	getopt.SetUsage(func() { fmt.Println(help + "\n\n" + copyrightLine) })
	getopt.Parse()
	if *hFlag {
		getopt.Usage()
		return
	}

	/* Get sillabary */
	runes, err := iberiar.GetSillabary(*sFlag)
	if err != nil {
		log.Fatal(err)
	}

	/* Read input. */
	scanner := bufio.NewScanner(os.Stdin)
	text    := ""
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	/* Translate. */
	iberian, latin, err := iberiar.ToIberian(text, runes)
	if err != nil {
		log.Fatal(err)
	}

	/* Print text. */
	if *SFlag {
		ins := iberiar.CreateInscription(iberian)
		if err != nil {
			log.Fatal(err)
		}
		err = iberiar.PrintInscription(ins, os.Stdout, *tFlag)
		if err != nil {
			log.Fatal(err)
		}
	} else if *lFlag {
		fmt.Println(latin)
	} else {
		fmt.Println(iberian)
	}
	
}

