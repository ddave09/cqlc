package generator

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func generator_tmpl_binding_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x5a,
		0x5d, 0x4f, 0xdb, 0xcc, 0x12, 0xbe, 0xcf, 0xaf, 0xd8, 0x22, 0x84, 0x1c,
		0x9a, 0x93, 0xf6, 0x3a, 0xe7, 0x70, 0x41, 0x43, 0x4a, 0x2d, 0xe5, 0x04,
		0xda, 0x98, 0xa2, 0x57, 0x55, 0x55, 0x6d, 0x9c, 0x0d, 0x58, 0x38, 0xeb,
		0xd4, 0xde, 0xd0, 0x22, 0x2b, 0xff, 0xfd, 0x9d, 0x59, 0xaf, 0xe3, 0xf5,
		0x7a, 0xed, 0x3a, 0x04, 0xde, 0x8f, 0x8a, 0xbd, 0x21, 0xb6, 0x67, 0xe7,
		0xe3, 0x99, 0x67, 0xf6, 0x93, 0x37, 0x6f, 0x88, 0xf7, 0xc1, 0x9d, 0x92,
		0xf7, 0xee, 0x78, 0x44, 0xae, 0x4f, 0xa7, 0xe4, 0xf4, 0xca, 0xbb, 0x38,
		0x1f, 0x4d, 0x46, 0x9f, 0x4e, 0xbd, 0xd1, 0x19, 0xf9, 0x0f, 0x39, 0x9d,
		0xfc, 0x41, 0x46, 0x67, 0xae, 0x37, 0x25, 0xde, 0x45, 0x26, 0x7a, 0xed,
		0x8e, 0xc7, 0xe4, 0xdd, 0x88, 0x8c, 0x2f, 0xa6, 0x1e, 0xb9, 0xfe, 0x30,
		0x9a, 0x10, 0xd7, 0x23, 0xf0, 0xfe, 0xd3, 0x68, 0xdb, 0xaf, 0x03, 0x6a,
		0x0b, 0x25, 0xa7, 0x1e, 0x49, 0x53, 0xd2, 0xbf, 0x8c, 0xa3, 0x7b, 0xc6,
		0x29, 0xf7, 0x59, 0xdf, 0x0b, 0x96, 0x2c, 0x11, 0x74, 0xb9, 0x22, 0x9b,
		0x0d, 0xb9, 0x9a, 0xba, 0x93, 0x73, 0xe2, 0x7f, 0x0f, 0x7d, 0xf2, 0x79,
		0xf4, 0x69, 0xea, 0x5e, 0x4c, 0x4c, 0xf1, 0xcf, 0x2c, 0x4e, 0x82, 0x88,
		0x83, 0x70, 0xa7, 0x03, 0x9f, 0x0e, 0x93, 0x87, 0xe5, 0x2c, 0x0a, 0x13,
		0x32, 0x38, 0x21, 0xfd, 0x8b, 0x95, 0x80, 0x4f, 0x49, 0x7f, 0xaa, 0xde,
		0xa1, 0xcc, 0x8a, 0xfa, 0x77, 0xf4, 0x86, 0x49, 0x35, 0xf9, 0xf7, 0x4b,
		0xf5, 0x0e, 0xbf, 0x07, 0xcb, 0x55, 0x14, 0x0b, 0xe2, 0x74, 0x08, 0xb4,
		0x34, 0x8d, 0x29, 0x87, 0x0f, 0x87, 0xdf, 0x7a, 0xe4, 0x70, 0x45, 0xc5,
		0xad, 0x54, 0xeb, 0x4a, 0x91, 0x04, 0xa4, 0x89, 0x6a, 0x07, 0x69, 0x2a,
		0x3f, 0x6f, 0x36, 0x07, 0xaa, 0x1f, 0xe3, 0x73, 0xf8, 0xde, 0xed, 0x74,
		0x7c, 0x30, 0x90, 0xab, 0x1b, 0x7e, 0x1c, 0x0f, 0xbf, 0xe5, 0x71, 0x9c,
		0x60, 0xaf, 0x9a, 0x50, 0x0e, 0xb0, 0x67, 0xc9, 0xb8, 0xbf, 0x90, 0xa6,
		0x87, 0x51, 0xb8, 0x5e, 0xf2, 0xf7, 0x74, 0x19, 0x84, 0x01, 0x43, 0x0f,
		0x94, 0x39, 0x72, 0x38, 0x15, 0xf1, 0xda, 0x17, 0xde, 0xc3, 0x8a, 0xa1,
		0x60, 0xc2, 0xe9, 0x1d, 0xf3, 0xa2, 0x21, 0x5d, 0xb2, 0x10, 0x3b, 0xf7,
		0x27, 0xf0, 0x8b, 0x14, 0xf2, 0xba, 0xe6, 0x28, 0xc4, 0x1e, 0x28, 0x94,
		0x69, 0xd7, 0x03, 0x43, 0xcd, 0xf0, 0xb6, 0x51, 0x79, 0x14, 0x6e, 0xb5,
		0xeb, 0xdd, 0x3e, 0xae, 0x69, 0x18, 0x2c, 0x02, 0x36, 0xaf, 0xf6, 0x5f,
		0xc5, 0x01, 0x17, 0x25, 0x97, 0x0d, 0x23, 0x9a, 0x2a, 0x81, 0xcf, 0x00,
		0xaf, 0x5d, 0xdd, 0x66, 0x93, 0xf9, 0x4c, 0x12, 0xf9, 0x2a, 0xdd, 0x76,
		0x53, 0x5e, 0x04, 0x8b, 0xcc, 0xc1, 0xe9, 0x7a, 0x25, 0x73, 0x36, 0x0c,
		0xd7, 0x89, 0x60, 0x60, 0xfe, 0x46, 0xb7, 0x81, 0x6d, 0xce, 0x12, 0x9f,
		0xcc, 0xa2, 0x28, 0x34, 0x55, 0x40, 0x26, 0x75, 0x59, 0x05, 0x21, 0xb6,
		0xc5, 0x9a, 0xfb, 0xc4, 0x99, 0x91, 0xe3, 0x16, 0xfe, 0x75, 0x49, 0xf6,
		0x03, 0x91, 0x72, 0xba, 0xe8, 0x2e, 0xfa, 0x50, 0xf6, 0x37, 0x66, 0x62,
		0x1d, 0x73, 0x49, 0xa6, 0x1c, 0xd4, 0x9c, 0x50, 0x7b, 0x59, 0xf6, 0x22,
		0xe7, 0x9e, 0x86, 0x6b, 0x46, 0x8e, 0xd3, 0x54, 0xfe, 0xc8, 0x20, 0x07,
		0x1b, 0x9b, 0x4d, 0x57, 0x16, 0x97, 0xca, 0xfc, 0xbb, 0x80, 0xcf, 0x6b,
		0xdd, 0xaa, 0xca, 0xa5, 0xd9, 0xd3, 0x80, 0xcc, 0x7a, 0xe4, 0x33, 0xea,
		0x1d, 0x10, 0xa9, 0xde, 0x0a, 0x96, 0x9e, 0x0c, 0x06, 0x45, 0x31, 0xa7,
		0xf1, 0x83, 0xcb, 0xe7, 0xec, 0xa7, 0x99, 0x88, 0xdd, 0x62, 0x1b, 0x7d,
		0x57, 0xb1, 0x35, 0x84, 0x06, 0xce, 0x62, 0x9d, 0x1b, 0x61, 0x61, 0xf3,
		0x33, 0x35, 0xc0, 0xc9, 0xa3, 0x5f, 0x1a, 0x4b, 0x37, 0x95, 0xee, 0x33,
		0x85, 0x17, 0xf4, 0x6f, 0x40, 0x27, 0x33, 0x52, 0x07, 0x91, 0x1d, 0x65,
		0xe5, 0x72, 0xaa, 0x74, 0x0d, 0x72, 0x53, 0x3d, 0x72, 0x19, 0xb3, 0x79,
		0xe0, 0x53, 0x01, 0x9a, 0xa4, 0xec, 0xe8, 0xfb, 0xf6, 0x4d, 0x59, 0x69,
		0xa9, 0x1a, 0x15, 0x89, 0x8d, 0x6c, 0x04, 0xc9, 0x30, 0x5a, 0x73, 0xa8,
		0x07, 0x05, 0xa7, 0x1c, 0x0b, 0xf6, 0x4a, 0xc7, 0x90, 0x72, 0x97, 0xfb,
		0x31, 0x5b, 0x32, 0x2e, 0x80, 0xe6, 0x58, 0x4f, 0x16, 0xd8, 0x55, 0xac,
		0xa0, 0x80, 0x35, 0xb8, 0x1c, 0x26, 0xcc, 0x74, 0xc6, 0x52, 0xd1, 0x97,
		0x34, 0x16, 0x12, 0x2b, 0x4b, 0x4d, 0xff, 0xd5, 0x74, 0xc2, 0xb6, 0x27,
		0xa5, 0xb0, 0x3d, 0x25, 0xad, 0x34, 0xb8, 0x9f, 0x8e, 0x5a, 0xe5, 0x5c,
		0x3d, 0x0e, 0xe9, 0x6d, 0xde, 0xde, 0x3d, 0x38, 0xa5, 0x41, 0xa8, 0x06,
		0x57, 0x15, 0xc4, 0xac, 0x85, 0x23, 0x1a, 0x4b, 0xdc, 0x64, 0x4c, 0x13,
		0x31, 0x8c, 0x60, 0xc6, 0xe6, 0x40, 0x49, 0x1b, 0x41, 0x76, 0x77, 0xdd,
		0xe5, 0x8a, 0x24, 0xfd, 0x7e, 0xff, 0xd1, 0x3c, 0xc1, 0xf6, 0x04, 0x5c,
		0xc1, 0xf6, 0xd4, 0x7c, 0xd1, 0xe0, 0xde, 0x99, 0x33, 0x2e, 0x6f, 0xe0,
		0x4c, 0x6d, 0xba, 0x8c, 0x39, 0xb6, 0xf1, 0x6d, 0xf3, 0x94, 0xbe, 0x27,
		0x2b, 0x95, 0xb2, 0xeb, 0x40, 0xdc, 0xd6, 0xcd, 0xd2, 0x06, 0x42, 0xb3,
		0xbe, 0x3e, 0xb1, 0x5b, 0xd8, 0xb9, 0xa7, 0x47, 0x67, 0xb0, 0x2e, 0xd9,
		0x16, 0x48, 0xe6, 0x9d, 0x14, 0xfe, 0x65, 0xa5, 0xb4, 0xa0, 0x14, 0xae,
		0x79, 0x06, 0x72, 0x18, 0xb6, 0xd5, 0xf7, 0x9e, 0x8e, 0xbb, 0x09, 0xba,
		0xce, 0x24, 0x4f, 0xea, 0xa7, 0x82, 0x12, 0x92, 0xe8, 0x4f, 0x8b, 0xfa,
		0xae, 0x29, 0x78, 0x1e, 0x89, 0x36, 0xeb, 0x8b, 0xc7, 0x45, 0xb3, 0xff,
		0xc4, 0x80, 0xed, 0x37, 0x2c, 0xf8, 0xc6, 0x49, 0x02, 0x5b, 0x53, 0xc1,
		0xbf, 0x0c, 0xdd, 0xff, 0xa0, 0x4c, 0x3e, 0xd9, 0xd0, 0xbd, 0x7b, 0x5a,
		0xce, 0xc5, 0xcb, 0xb2, 0x6b, 0x9b, 0x87, 0x73, 0xf1, 0x8c, 0xcb, 0xae,
		0x73, 0xf6, 0x82, 0x74, 0x81, 0x34, 0x7b, 0x46, 0xa4, 0xc7, 0x2f, 0x9c,
		0x2e, 0x90, 0x1e, 0x3f, 0x27, 0xa7, 0xc7, 0x2f, 0x9c, 0xd6, 0x90, 0x6e,
		0xcf, 0xe9, 0xea, 0xe8, 0x6d, 0xcc, 0xcc, 0xf9, 0x31, 0xa6, 0x7c, 0xc8,
		0x0f, 0xe2, 0x74, 0x44, 0xd4, 0xb1, 0x9b, 0x86, 0x6a, 0xfb, 0xa3, 0xc5,
		0x4c, 0xda, 0x7e, 0x96, 0x08, 0x9a, 0xab, 0xb9, 0xd4, 0x6c, 0x64, 0x6e,
		0x65, 0x11, 0xed, 0x78, 0xa2, 0x99, 0x51, 0x2b, 0xc9, 0xa8, 0xa5, 0xc7,
		0xd2, 0x6d, 0xf0, 0x46, 0xa6, 0xcb, 0xe9, 0x5a, 0x9c, 0xb2, 0x9f, 0x95,
		0x25, 0xfd, 0x7a, 0x5d, 0x9d, 0x72, 0x32, 0x7e, 0x89, 0xf1, 0x19, 0x5b,
		0xec, 0x0f, 0xb3, 0x88, 0xae, 0x56, 0x2b, 0x16, 0x1b, 0x08, 0x67, 0x6c,
		0x6c, 0x09, 0xb1, 0x04, 0x0e, 0xf9, 0x67, 0xf8, 0xe7, 0x04, 0xb0, 0x25,
		0x21, 0xc7, 0x37, 0x11, 0xd0, 0xaf, 0xef, 0xc2, 0xef, 0x2e, 0x71, 0xbe,
		0x7c, 0x35, 0x84, 0x7a, 0x84, 0xc5, 0x71, 0x04, 0x9f, 0x8a, 0x10, 0x68,
		0x1c, 0xd3, 0x07, 0xf4, 0x7b, 0x09, 0x38, 0xd9, 0x7a, 0xbc, 0x2d, 0x76,
		0x53, 0xd0, 0x19, 0x25, 0xff, 0x4f, 0x57, 0x36, 0xe3, 0x3d, 0xe9, 0x9b,
		0x23, 0xaa, 0x19, 0x75, 0x70, 0xd7, 0x61, 0xb1, 0x5d, 0xd8, 0x3f, 0x21,
		0x14, 0x80, 0xe1, 0x73, 0x47, 0x3e, 0xf6, 0x88, 0xe8, 0xda, 0xf2, 0x89,
		0x5b, 0xa4, 0x1e, 0xe1, 0x41, 0x71, 0x38, 0xbc, 0x29, 0xe4, 0x94, 0x8c,
		0x52, 0x00, 0xb6, 0x2a, 0xa8, 0xd5, 0xf8, 0xad, 0x83, 0xd6, 0x23, 0x3e,
		0x0d, 0xc3, 0x19, 0xf5, 0xef, 0xda, 0x05, 0xd3, 0xcd, 0xfe, 0x6a, 0x31,
		0x65, 0xc9, 0x94, 0x17, 0x1f, 0xa8, 0x3c, 0xa7, 0x82, 0xb6, 0x27, 0x8d,
		0xa3, 0x1f, 0x1a, 0xde, 0x01, 0x1e, 0xfc, 0x2d, 0xa8, 0xcf, 0x52, 0xc0,
		0x3a, 0x64, 0xdc, 0x51, 0x0a, 0xba, 0x5d, 0xed, 0x9c, 0xb9, 0x64, 0x41,
		0x32, 0x14, 0x35, 0x18, 0xae, 0xa5, 0xc6, 0x42, 0x1e, 0x7b, 0x05, 0x28,
		0xf7, 0xf6, 0xbf, 0xf0, 0xf7, 0x7f, 0x25, 0xe5, 0xf0, 0xe6, 0xf5, 0x6b,
		0xcb, 0x28, 0x9c, 0xfc, 0x08, 0x84, 0x7f, 0x9b, 0x07, 0xf1, 0x25, 0xf8,
		0x9a, 0xdd, 0x29, 0x54, 0x05, 0x77, 0x63, 0xfe, 0x16, 0x1c, 0x9a, 0x30,
		0xf3, 0x60, 0x7d, 0x80, 0x80, 0x80, 0x25, 0xe0, 0xc0, 0x91, 0x68, 0x53,
		0xb0, 0x85, 0x0b, 0x45, 0x71, 0xe8, 0x6d, 0xce, 0x16, 0x74, 0x1d, 0x8a,
		0x81, 0xd5, 0x83, 0x30, 0xba, 0xe9, 0xbf, 0xa7, 0x82, 0x86, 0xce, 0xc1,
		0x9a, 0xdf, 0x52, 0x3e, 0x0f, 0xd9, 0x5c, 0x85, 0x3b, 0x20, 0x07, 0x3d,
		0x33, 0x72, 0xdb, 0x49, 0x42, 0xfd, 0x13, 0xec, 0x99, 0x5e, 0xc9, 0xa4,
		0x4f, 0x7d, 0xca, 0x1d, 0x08, 0x0b, 0xb6, 0x37, 0x26, 0xdf, 0xb1, 0xcd,
		0x62, 0x46, 0xef, 0x0c, 0x3d, 0x06, 0xdf, 0xe9, 0x7c, 0xc2, 0x7e, 0x8a,
		0x5e, 0x5e, 0x72, 0x39, 0x2b, 0x1d, 0xa3, 0x30, 0xc0, 0x22, 0x4a, 0xbc,
		0x3a, 0xc1, 0xaa, 0xa8, 0x3f, 0xe2, 0xcd, 0xab, 0xa1, 0xde, 0xed, 0xdc,
		0x62, 0xbd, 0x0e, 0xbd, 0xec, 0xca, 0x3a, 0x34, 0xe7, 0x0d, 0xd9, 0xe2,
		0x62, 0x0c, 0x6c, 0xc0, 0x56, 0x52, 0x9d, 0x73, 0x93, 0x36, 0x13, 0x01,
		0x0e, 0xb8, 0x78, 0x7c, 0xa1, 0xfa, 0x78, 0x74, 0x16, 0x32, 0xfb, 0x01,
		0x86, 0xed, 0x1c, 0x3b, 0x1f, 0xd5, 0xe5, 0xf9, 0x75, 0x7b, 0x73, 0xf9,
		0x69, 0xd6, 0xd5, 0x2a, 0x61, 0x71, 0xcd, 0xd1, 0x79, 0xb3, 0xb9, 0x62,
		0x26, 0xef, 0xb4, 0xb0, 0x27, 0xa3, 0xaa, 0xbb, 0x8c, 0xd2, 0x2f, 0xa2,
		0x16, 0xa5, 0x7b, 0x28, 0x7d, 0x54, 0xab, 0x57, 0x8e, 0xb3, 0x84, 0x73,
		0x5f, 0x1d, 0xc1, 0xe4, 0x02, 0x45, 0x5a, 0xae, 0xde, 0x34, 0xf9, 0xea,
		0xd2, 0xf6, 0xcb, 0x57, 0xcb, 0xf2, 0xe9, 0xd1, 0xf3, 0xde, 0xa3, 0x6e,
		0x2f, 0xf3, 0xae, 0x4f, 0x74, 0x83, 0x29, 0xa3, 0xab, 0x5f, 0x12, 0x1e,
		0x35, 0x98, 0x2a, 0xd6, 0x9c, 0xc5, 0x92, 0xb1, 0x61, 0xa4, 0xda, 0xf4,
		0xac, 0x13, 0x78, 0xc1, 0x13, 0x2d, 0xb9, 0x95, 0x5c, 0xa4, 0xf2, 0x41,
		0xfa, 0x63, 0xa6, 0x14, 0xcd, 0x2b, 0x84, 0xe5, 0x12, 0x36, 0xa9, 0x2e,
		0x0c, 0x1a, 0x98, 0x16, 0x01, 0x15, 0x8e, 0x5f, 0xb8, 0xf0, 0x0c, 0x5c,
		0x38, 0xfa, 0xd7, 0x91, 0x21, 0xeb, 0x08, 0xbf, 0x03, 0x2e, 0x37, 0x34,
		0xb0, 0x3c, 0x29, 0x67, 0xb9, 0x3a, 0x0a, 0x95, 0x3e, 0x9b, 0x17, 0xfc,
		0xbb, 0xae, 0x05, 0xf6, 0x60, 0xc1, 0x33, 0x30, 0x01, 0x5b, 0xcb, 0x94,
		0x1b, 0x71, 0xdb, 0xf2, 0xa9, 0x67, 0xc0, 0x00, 0x5f, 0x66, 0x13, 0x50,
		0x07, 0xb4, 0xcd, 0x4a, 0xc4, 0xbc, 0x54, 0x30, 0xb7, 0xa5, 0xfd, 0xb7,
		0x43, 0xde, 0xbe, 0x2b, 0x6a, 0x5d, 0x84, 0xed, 0x33, 0xa2, 0x0e, 0xb5,
		0x93, 0xe2, 0xbf, 0x90, 0x72, 0xb1, 0x7b, 0x1a, 0x97, 0xdc, 0xc8, 0x27,
		0x5b, 0x52, 0x59, 0x63, 0x17, 0x19, 0x34, 0x26, 0xfc, 0x1d, 0xb7, 0xbf,
		0x7f, 0xff, 0x3f, 0xf4, 0x6c, 0x47, 0x88, 0x2a, 0xc7, 0xba, 0x16, 0xf7,
		0x72, 0xc4, 0xe5, 0xee, 0xbb, 0xb2, 0x61, 0xb5, 0xaf, 0x90, 0xda, 0x65,
		0x10, 0x3b, 0xd4, 0x6d, 0xc6, 0xd5, 0x8f, 0x3f, 0x03, 0x00, 0x00, 0xff,
		0xff, 0x3c, 0xe3, 0x1c, 0x5e, 0x02, 0x27, 0x00, 0x00,
		},
		"generator/tmpl/binding.tmpl",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"generator/tmpl/binding.tmpl": generator_tmpl_binding_tmpl,

}
