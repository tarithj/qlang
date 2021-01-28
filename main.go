package main

import (
	"flag"
	"fmt"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	"os"
	"qlang/parse"
	_ "qlang/statik"
)

func main() {
	init := flag.Bool("init", false, "use to init the current dir")
	location := flag.String("location", "", "the location of the xml")
	flag.Parse()

	if *init && *location == "" {
		initalize()
	} else {
		f, err := os.Open(*location)
		if err != nil {
			log.Fatalln(err)
		}

		Run(parse.Parse(f))
	}
}
func initalize() {
	sfs, err := fs.New()
	if err != nil {
		log.Fatalf("Unable to init internal fs %v\n", err)
	}
	baseXmlReader, err := sfs.Open("/base.xml")
	if err != nil {
		log.Fatalf("Unable to open base.xml in internal fs %v\n", err)
	}

	baseXml, err := ioutil.ReadAll(baseXmlReader)
	if err != nil {
		log.Fatalf("Unable to read base.xml in internal fs %v\n", err)
	}

	langDefReader, err := sfs.Open("/qlang.xsd")
	if err != nil {
		log.Fatalf("Unable to open qlang.xsd in internal fs %v\n", err)
	}

	langDef, err := ioutil.ReadAll(langDefReader)
	if err != nil {
		log.Fatalf("Unable to read qlang.xsd in internal fs %v\n", err)
	}

	cd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get wd %v\n", err)
	}
	f, err := os.Create(cd + "/index.xml")
	if err != nil {
		log.Fatalf("Unable to create index.xml%v\n", err)
	}
	_, err = f.Write(baseXml)
	if err != nil {
		log.Fatalf("Unable to write base.xml to index.xml%v\n", err)
	}

	err = os.Mkdir(cd+"/def", os.ModeDir)
	if err != nil {
		log.Fatalf("unable to create /def directory%v\n", err)
	}

	f, err = os.Create(cd + "/def/qlang.xsd")
	if err != nil {
		log.Fatalf("unable to create qlang.xsd in def directory %v\n", err)
	}

	_, err = f.Write(langDef)
	if err != nil {
		log.Fatalf("unable to write qlang.xsd in def directory %v\n", err)
	}

}

func Run(p parse.QPack) {
	var marks int64
	var report string

	if p.Info.Author != "" {
		fmt.Printf("questions made by %s\n\n", p.Info.Author)
	}

	for i := 0; i < len(p.QSets); i++ {
		res := p.QSets[i].GetResponse(i)

		var ii = 0
		for _, val := range res {
			if val.Correct == "true" {
				marks += 1
				val.Correct = "correct"
			} else {
				val.Correct = "incorrect"
			}
			report += fmt.Sprintf("%d) %d) your answer '%s' is %s\n", i, ii, val.Text, val.Correct)
			ii++
		}
	}
	fmt.Println(report)
	fmt.Printf("you got %d questions correct", marks)
}
