package pnpm

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"caddy/src/config"
	"caddy/src/tools"
)

func Uncompress(archivePath string, version string) error {
	archiveFile, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	uncompressedStream, err := gzip.NewReader(archiveFile)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)
	destination := filepath.Join(config.SystemPaths.Pnpm, version)

	if err := os.MkdirAll(destination, 0755); err != nil {
		return err
	}

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// preserve full path from archive
		target := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, header.FileInfo().Mode()); err != nil {
				return err
			}

		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}

			outfile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, header.FileInfo().Mode())
			if err != nil {
				return err
			}

			if _, err := io.Copy(outfile, tarReader); err != nil {
				outfile.Close()
				return err
			}

			outfile.Close()

		case tar.TypeSymlink:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}
			if err := os.Symlink(header.Linkname, target); err != nil && !os.IsExist(err) {
				return err
			}

		default:
			fmt.Printf("Skipping unsupported tar entry type %v for: %s\n", header.Typeflag, header.Name)
		}
	}

	binPath := tools.ToolBin(config.CaddyTool.Pnpm, version)
	fmt.Printf("Setting executable bit on: %s\n", binPath)
	os.Chmod(binPath, 0755)

	return nil
}
