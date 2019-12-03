package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"

	"github.com/Mohammed-Aadil/ds-storage/config"
	"github.com/Mohammed-Aadil/ds-storage/pkg/errors"
)

//FileToGzip it will Gzip single file
func FileToGzip(attachment IAttachment, tarFileWriter *tar.Writer) error {
	file, err := os.Open(attachment.GetPath())
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	// preparing tar header
	header := new(tar.Header)
	header.Name = fileInfo.Name()
	header.Size = fileInfo.Size()
	header.Mode = int64(fileInfo.Mode())
	header.ModTime = fileInfo.ModTime()
	// writing files into tarball
	err = tarFileWriter.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tarFileWriter, file)
	if err != nil {
		return err
	}
	return nil
}

//FilesToGzip it will gzip multiple files
func FilesToGzip(attachments []IAttachment) (string, error) {
	if len(attachments) == 0 {
		return "", &errors.NonFieldValidationError{Msg: "No files found."}
	}
	// creating tarfile at destination path
	tarfile, err := os.Create(config.DocTempDirPath + "/" + attachments[0].GetName() + ".gz")
	if err != nil {
		return "", err
	}
	defer tarfile.Close()
	//using write and close interface to access only those methods
	var fileWriter io.WriteCloser = tarfile
	// creating new gzip writer
	fileWriter = gzip.NewWriter(tarfile)
	defer fileWriter.Close()
	// creating tar file writer
	tarFileWriter := tar.NewWriter(fileWriter)
	defer tarFileWriter.Close()
	for _, attachment := range attachments {
		if err := FileToGzip(attachment, tarFileWriter); err != nil {
			return "", err
		}
	}
	return tarfile.Name(), nil
}
