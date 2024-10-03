package main

import "io"

func copySrcToDest(src io.Reader, dest io.Writer) error {
	bytes, err := io.ReadAll(src)

	if err != nil {
		return err
	}

	_, err = dest.Write(bytes)

	return nil
}
