package crawler

import (
	"log"
)

type Parser interface {
	ListParse() error
	DetailParse(url string) error

}

func DoCrawl(parse string, isAll bool, whichSite string) error {
	log.Print("start crawling...")

	return nil
}
