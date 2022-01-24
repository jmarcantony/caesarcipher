package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jmarcantony/caesercipher"
)

var (
	sh  = flag.Int("s", 0, "shift value")
	enc = flag.Bool("e", false, "encode")
	dec = flag.Bool("d", false, "decode")
	f   = flag.String("f", "", "file to encode/decode")
	o   = flag.String("O", "", "output file to write encoded/decoded string to")
)

func write(s []byte) {
	f, err := os.OpenFile(*o, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(s)
	fmt.Printf("Wrote to %q succesfully", *o)
}

func main() {
	flag.Parse()
	s := flag.Arg(0)
	if *f != "" {
		data, err := ioutil.ReadFile(*f)
		if err != nil {
			log.Fatal(err)
		}
		s = string(data)
	}
	if *sh < 0 {
		log.Fatalln("Shift value cannot be negative")
	}
	if *enc {
		encoded := caesercipher.Encode(s, *sh)
		if *o != "" {
			write(encoded)
		} else {
			fmt.Printf("%s\n", encoded)
		}
	}
	if *dec {
		decoded := caesercipher.Decode(s, *sh)
		if *o != "" {
			write(decoded)
		} else {
			fmt.Printf("%s\n", decoded)
		}
	}
}
