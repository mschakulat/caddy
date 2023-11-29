package node

import (
	"archive/tar"
	"caddy/src/config"
	"caddy/src/tools"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

	var toplevelDirectory string
	destination := fmt.Sprintf("%s", config.SystemPaths.Node)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		target := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if toplevelDirectory == "" {
				toplevelDirectory = header.Name
			}

			if !strings.HasPrefix(header.Name, toplevelDirectory) {
				continue
			}

			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}

		case tar.TypeReg:
			if !strings.HasPrefix(header.Name, toplevelDirectory) {
				continue
			}

			dir := filepath.Dir(target)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}

			outfile, err := os.Create(target)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outfile, tarReader); err != nil {
				return err
			}

			outfile.Close()
		case tar.TypeSymlink:
			dir := filepath.Dir(target)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if err := os.Symlink(header.Linkname, target); err != nil {
				return err
			}
		default:
			fmt.Printf("Can't untar type %c for file %s\n", header.Typeflag, header.Name)
		}
	}

	if err := os.Rename(
		filepath.Join(destination, toplevelDirectory), filepath.Join(destination, version),
	); err != nil {
		return err
	}

	os.Chmod(tools.ToolBin(config.CaddyTool.Node, version), 0755)
	os.Chmod(tools.ToolBin(config.CaddyTool.Npm, version), 0755)
	os.Chmod(tools.ToolBin(config.CaddyTool.Npx, version), 0755)

	return nil
}
