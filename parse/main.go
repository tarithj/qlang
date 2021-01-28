package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type QPack struct {
	XMLName xml.Name `xml:"qPack"`
	QSets   []QSet   `xml:"qSet"`
	Head    string   `xml:"head"`
	Info    Info     `xml:"info"`
}
type Head struct {
	Text string `xml:"text,attr"`
}
type Info struct {
	Author string `xml:"text,attr"`
}
type QSet struct {
	XMLName xml.Name `xml:"qSet"`
	Text    string   `xml:"text,attr"`
	Q       []Q      `xml:"q"`
}
type Q struct {
	XMLName xml.Name `xml:"q"`
	Text    string   `xml:"text,attr"`
	Choice  []Choice `xml:"choice"`
}

type Choice struct {
	XMLName xml.Name `xml:"choice"`
	Correct string   `xml:"correct,attr"`
	Text    string   `xml:"text,attr"`
}

func (question *QSet) GetResponse(index int) map[string]Choice {
	fmt.Printf("%d) %s\n", index, question.Text)

	var choices = make(map[string]Choice)
	for i := 0; i < len(question.Q); i++ {
		displayQuestion := question.Q[i]
		fmt.Printf("	%d?) %s\n", i, displayQuestion.Text)
		for ii := 0; ii < len(displayQuestion.Choice); ii++ {
			displayChoice := displayQuestion.Choice[ii]
			fmt.Printf("		%d) %s\n", ii, displayChoice.Text)
		}
	a:
		fmt.Println("Enter your answers index number: ")
		var input int
		fmt.Scanln(&input)
		if input <= len(displayQuestion.Choice) && input > -1 {
			choices[displayQuestion.Text] = displayQuestion.Choice[input]
		} else {
			fmt.Println("Wrong input")
			goto a
		}
	}
	return choices
}
func Parse(data io.Reader) QPack {
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(data)

	// we initialize our Users array
	var p QPack
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	if err := xml.Unmarshal(byteValue, &p); err != nil {
		log.Fatalln(err)
	}
	return p
}
