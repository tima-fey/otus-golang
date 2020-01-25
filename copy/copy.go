package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func copyChank(source io.Reader, destination io.Writer, chunkSize int) (bool, error) {
	buf := make([]byte, chunkSize)
	localOffset := 0
	isEndfile := false
	for localOffset < chunkSize {
		read, err := source.Read(buf[localOffset:])
		localOffset += read
		if err == io.EOF {
			isEndfile = true
			break
		}
		if err != nil {
			return false, err
		}
	}
	_, err := destination.Write(buf[:localOffset])
	if err != nil {
		fmt.Println("writing error")
	}
	return isEndfile, nil
}
func CustomCopyFile(source, destination *os.File, offset, limit int) error {
	stat, err := source.Stat()
	if err != nil {
		return err
	}
	toRead := stat.Size() - int64(offset)
	if limit > 0 {
		if int64(limit) < toRead {
			toRead = int64(limit)
		}
	}

	bar := pb.Full.Start64(toRead)
	barReader := bar.NewProxyReader(source)
	_, err = source.Seek(int64(offset), 0)
	if err != nil {
		return err
	}
	chankSize := 2
	for toRead > 0 {
		if int64(chankSize) > toRead {
			chankSize = int(toRead)
		}
		isEOF, err := copyChank(barReader, destination, chankSize)
		if err != nil {
			return err
		}
		if isEOF {
			break
		}
		toRead -= int64(chankSize)
	}
	bar.Finish()
	return nil
}
func CopyFile(source, destination *os.File, offset, limit int) error {
	stat, err := source.Stat()
	if err != nil {
		return err
	}
	toRead := stat.Size() - int64(offset)
	if limit > 0 {
		if int64(limit) < toRead {
			toRead = int64(limit)
		}
	}
	bar := pb.Full.Start64(toRead)
	barReader := bar.NewProxyReader(source)
	_, err = source.Seek(int64(offset), 0)
	if err != nil {
		return err
	}
	_, err = io.CopyN(destination, barReader, toRead)
	if err != nil {
		return err
	}
	bar.Finish()
	return nil
}

func Copy(offset, limit int, sourceName, destinationName string, isCustom bool) {
	var sourceD *os.File
	sourceD, err := os.Open(sourceName)
	if err != nil {
		fmt.Println("error")
	}
	defer sourceD.Close()

	var destinationD *os.File
	destinationD, err = os.Create(destinationName)
	if err != nil {
		fmt.Println("error")
	}
	defer destinationD.Close()
	if isCustom {
		err = CustomCopyFile(sourceD, destinationD, offset, limit)
	} else {
		err = CopyFile(sourceD, destinationD, offset, limit)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var offset int
	var sourceName string
	var destinationName string
	var limit int
	var isCustom bool
	flag.IntVar(&offset, "offset", 0, "offset in input file")
	flag.StringVar(&sourceName, "source", "", "osource file")
	flag.StringVar(&destinationName, "destination", "", "destination file")
	flag.IntVar(&limit, "limit", 0, "limit of copy")
	flag.BoolVar(&isCustom, "custom", false, "use custom copy func")
	flag.Parse()
	Copy(offset, limit, sourceName, destinationName, isCustom)
}
