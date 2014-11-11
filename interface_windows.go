// +build windows

// interface_windows
package mso2pdf

type Exporter interface {
	Export(inFile, outDir string) (outFile string, err error)
}
