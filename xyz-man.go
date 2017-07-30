package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	args := os.Args
	//args := [3]string{"1", "8.xyz", "-123.45"}
	if len(args) != 3 {
		fmt.Println("xyz-man " + " input-file.xyz correction-value")
		fmt.Println("        input-file.xyz      XYZ-file")
		fmt.Println("        correction-value    any flot number (e.g 1.23)")
		fmt.Println("       Output: input-file-(correction-value).xyz with the corrected depth values")
		os.Exit(0)
	}

	fin, err := ioutil.ReadFile(args[1])
	check(err)

	var i = strings.Index(args[1], ".")
	var fouts string
	if i == -1 {
		fouts = args[1] + "-" + args[2] + ".xyz"
	} else {
		fouts = args[1][0:i] + "-" + args[2] + ".xyz"
	}

	fout, err := os.Create(fouts)
	check(err)
	foutb := bufio.NewWriter(fout)

	correction, err := strconv.ParseFloat(args[2], 64)
	check(err)

	//fmt.Printf("%T, %v\n", string(fin), len(string(fin)))

	sa := strings.Split(string(fin), "\r")
	//fmt.Printf("%T, %v\n", sa, len(sa))

	for i := 0; i < len(sa); i++ {
		//fmt.Println(sa[i])
		t := strings.Replace(sa[i], " ", "\t", -1)
		ss := strings.Split(t, "\t")
		if len(ss) != 3 {
			fmt.Printf("#%v - no 3 elements in line: '%v'\n", i+1, sa[i])
		} else {
			ss[2] = strings.Replace(ss[2], "\r", "", -1)
			depth, err := strconv.ParseFloat(ss[2], 64)
			//			fmt.Printf(string(depth))
			if err != nil {
				fmt.Printf("#" + string(i+1) + " - last parameter invalid: '" + sa[i] + "'\n")
				//fmt.Println(err)
			} else {
				depthC := depth + correction
				var r = fmt.Sprintf("%s %s %.3f\n", ss[0], ss[1], depthC)
				// fmt.Print(r)
				foutb.WriteString(r)
				check(err)
			}
		}
	}
	foutb.Flush()
	defer fout.Close()

}
