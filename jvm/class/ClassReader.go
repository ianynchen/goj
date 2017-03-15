package class

import "bufio"

func ReadClassHeader(reader *bufio.Reader) (*ClassHeader, error) {

	header := new(ClassHeader)

	content := make([]byte, 4)
	read, err := reader.Read(content)

	if err != nil {
		return header, err
	} else {
		header.Magic, err = ParseUint32(content)
	}

	header.Initialize()
	return header, err
}
