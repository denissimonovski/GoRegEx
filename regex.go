package main

import (
	"time"
	"regexp"
	"io/ioutil"
	"sort"
	"os"
	"fmt"
)

func get_date(filename string) string {
	reg := regexp.MustCompile(`.* Report \((\d{2}\.\d{2}\.)\-\d{2}\.\d{2}\.\)\.txt`)
	a := reg.FindAllStringSubmatch(filename, -1)
	if a != nil {
		return a[0][1]
	} else {
		return ""
	}
}

func date_from_name(filename string) time.Time {
	layout := "02.01."
	vreme, _ := time.Parse(layout, get_date(filename))
	return vreme
}

func main() {
	r := regexp.MustCompile("Nesto")
	fajlovi, _ := ioutil.ReadDir("D:\\GoEnvironment\\GoProjects\\Regex")
	for i := len(fajlovi) - 1; i >= 0; i-- {
		if get_date(fajlovi[i].Name()) == "" {
			fajlovi = append(fajlovi[:i], fajlovi[i+1:]...)
		}
	}
	sort.Slice(fajlovi, func(i, j int) bool {
		return date_from_name(fajlovi[i].Name()).Before(date_from_name(fajlovi[j].Name()))
	})
	for _, i := range fajlovi {
		fmt.Println(i.Name(), "==>", r.ReplaceAllString(i.Name(), "Nekoj"))
		os.Rename(i.Name(), r.ReplaceAllString(i.Name(), "Nekoj"))
	}
}
