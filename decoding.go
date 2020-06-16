package goqr

import (
	"math"
)

const maxPoly = 64

type galoisField struct {
	p   int
	log []uint8
	exp []uint8
}

var gf16Exp = []uint8{
	0x01, 0x02, 0x04, 0x08, 0x03, 0x06, 0x0c, 0x0b,
	0x05, 0x0a, 0x07, 0x0e, 0x0f, 0x0d, 0x09, 0x01,
}

var gf16Log = []uint8{
	0x00, 0x0f, 0x01, 0x04, 0x02, 0x08, 0x05, 0x0a,
	0x03, 0x0e, 0x09, 0x07, 0x06, 0x0d, 0x0b, 0x0c,
}

var gf16 = galoisField{
	p:   15,
	log: gf16Log,
	exp: gf16Exp,
}

var gf256Exp = [256]uint8{
	0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80,
	0x1d, 0x3a, 0x74, 0xe8, 0xcd, 0x87, 0x13, 0x26,
	0x4c, 0x98, 0x2d, 0x5a, 0xb4, 0x75, 0xea, 0xc9,
	0x8f, 0x03, 0x06, 0x0c, 0x18, 0x30, 0x60, 0xc0,
	0x9d, 0x27, 0x4e, 0x9c, 0x25, 0x4a, 0x94, 0x35,
	0x6a, 0xd4, 0xb5, 0x77, 0xee, 0xc1, 0x9f, 0x23,
	0x46, 0x8c, 0x05, 0x0a, 0x14, 0x28, 0x50, 0xa0,
	0x5d, 0xba, 0x69, 0xd2, 0xb9, 0x6f, 0xde, 0xa1,
	0x5f, 0xbe, 0x61, 0xc2, 0x99, 0x2f, 0x5e, 0xbc,
	0x65, 0xca, 0x89, 0x0f, 0x1e, 0x3c, 0x78, 0xf0,
	0xfd, 0xe7, 0xd3, 0xbb, 0x6b, 0xd6, 0xb1, 0x7f,
	0xfe, 0xe1, 0xdf, 0xa3, 0x5b, 0xb6, 0x71, 0xe2,
	0xd9, 0xaf, 0x43, 0x86, 0x11, 0x22, 0x44, 0x88,
	0x0d, 0x1a, 0x34, 0x68, 0xd0, 0xbd, 0x67, 0xce,
	0x81, 0x1f, 0x3e, 0x7c, 0xf8, 0xed, 0xc7, 0x93,
	0x3b, 0x76, 0xec, 0xc5, 0x97, 0x33, 0x66, 0xcc,
	0x85, 0x17, 0x2e, 0x5c, 0xb8, 0x6d, 0xda, 0xa9,
	0x4f, 0x9e, 0x21, 0x42, 0x84, 0x15, 0x2a, 0x54,
	0xa8, 0x4d, 0x9a, 0x29, 0x52, 0xa4, 0x55, 0xaa,
	0x49, 0x92, 0x39, 0x72, 0xe4, 0xd5, 0xb7, 0x73,
	0xe6, 0xd1, 0xbf, 0x63, 0xc6, 0x91, 0x3f, 0x7e,
	0xfc, 0xe5, 0xd7, 0xb3, 0x7b, 0xf6, 0xf1, 0xff,
	0xe3, 0xdb, 0xab, 0x4b, 0x96, 0x31, 0x62, 0xc4,
	0x95, 0x37, 0x6e, 0xdc, 0xa5, 0x57, 0xae, 0x41,
	0x82, 0x19, 0x32, 0x64, 0xc8, 0x8d, 0x07, 0x0e,
	0x1c, 0x38, 0x70, 0xe0, 0xdd, 0xa7, 0x53, 0xa6,
	0x51, 0xa2, 0x59, 0xb2, 0x79, 0xf2, 0xf9, 0xef,
	0xc3, 0x9b, 0x2b, 0x56, 0xac, 0x45, 0x8a, 0x09,
	0x12, 0x24, 0x48, 0x90, 0x3d, 0x7a, 0xf4, 0xf5,
	0xf7, 0xf3, 0xfb, 0xeb, 0xcb, 0x8b, 0x0b, 0x16,
	0x2c, 0x58, 0xb0, 0x7d, 0xfa, 0xe9, 0xcf, 0x83,
	0x1b, 0x36, 0x6c, 0xd8, 0xad, 0x47, 0x8e, 0x01,
}

var gf256Log = [256]uint8{
	0x00, 0xff, 0x01, 0x19, 0x02, 0x32, 0x1a, 0xc6,
	0x03, 0xdf, 0x33, 0xee, 0x1b, 0x68, 0xc7, 0x4b,
	0x04, 0x64, 0xe0, 0x0e, 0x34, 0x8d, 0xef, 0x81,
	0x1c, 0xc1, 0x69, 0xf8, 0xc8, 0x08, 0x4c, 0x71,
	0x05, 0x8a, 0x65, 0x2f, 0xe1, 0x24, 0x0f, 0x21,
	0x35, 0x93, 0x8e, 0xda, 0xf0, 0x12, 0x82, 0x45,
	0x1d, 0xb5, 0xc2, 0x7d, 0x6a, 0x27, 0xf9, 0xb9,
	0xc9, 0x9a, 0x09, 0x78, 0x4d, 0xe4, 0x72, 0xa6,
	0x06, 0xbf, 0x8b, 0x62, 0x66, 0xdd, 0x30, 0xfd,
	0xe2, 0x98, 0x25, 0xb3, 0x10, 0x91, 0x22, 0x88,
	0x36, 0xd0, 0x94, 0xce, 0x8f, 0x96, 0xdb, 0xbd,
	0xf1, 0xd2, 0x13, 0x5c, 0x83, 0x38, 0x46, 0x40,
	0x1e, 0x42, 0xb6, 0xa3, 0xc3, 0x48, 0x7e, 0x6e,
	0x6b, 0x3a, 0x28, 0x54, 0xfa, 0x85, 0xba, 0x3d,
	0xca, 0x5e, 0x9b, 0x9f, 0x0a, 0x15, 0x79, 0x2b,
	0x4e, 0xd4, 0xe5, 0xac, 0x73, 0xf3, 0xa7, 0x57,
	0x07, 0x70, 0xc0, 0xf7, 0x8c, 0x80, 0x63, 0x0d,
	0x67, 0x4a, 0xde, 0xed, 0x31, 0xc5, 0xfe, 0x18,
	0xe3, 0xa5, 0x99, 0x77, 0x26, 0xb8, 0xb4, 0x7c,
	0x11, 0x44, 0x92, 0xd9, 0x23, 0x20, 0x89, 0x2e,
	0x37, 0x3f, 0xd1, 0x5b, 0x95, 0xbc, 0xcf, 0xcd,
	0x90, 0x87, 0x97, 0xb2, 0xdc, 0xfc, 0xbe, 0x61,
	0xf2, 0x56, 0xd3, 0xab, 0x14, 0x2a, 0x5d, 0x9e,
	0x84, 0x3c, 0x39, 0x53, 0x47, 0x6d, 0x41, 0xa2,
	0x1f, 0x2d, 0x43, 0xd8, 0xb7, 0x7b, 0xa4, 0x76,
	0xc4, 0x17, 0x49, 0xec, 0x7f, 0x0c, 0x6f, 0xf6,
	0x6c, 0xa1, 0x3b, 0x52, 0x29, 0x9d, 0x55, 0xaa,
	0xfb, 0x60, 0x86, 0xb1, 0xbb, 0xcc, 0x3e, 0x5a,
	0xcb, 0x59, 0x5f, 0xb0, 0x9c, 0xa9, 0xa0, 0x51,
	0x0b, 0xf5, 0x16, 0xeb, 0x7a, 0x75, 0x2c, 0xd7,
	0x4f, 0xae, 0xd5, 0xe9, 0xe6, 0xe7, 0xad, 0xe8,
	0x74, 0xd6, 0xf4, 0xea, 0xa8, 0x50, 0x58, 0xaf,
}

var gf256 = galoisField{
	255,
	gf256Log[:],
	gf256Exp[:],
}

/************************************************************************
 * Polynomial operations
 */

func polyAdd(dst, src []uint8, c uint8, shift int, gf *galoisField) {

	logC := gf.log[c]
	if c == 0 {
		return
	}

	for i := 0; i < maxPoly; i++ {
		p := i + shift
		v := src[i]
		if p < 0 || p >= maxPoly {
			continue
		}
		if v == 0 {
			continue
		}
		pos := (int(gf.log[v]) + int(logC)) % gf.p

		dst[p] ^= gf.exp[pos]
	}
}

func polyEval(s []uint8, x uint8, gf *galoisField) uint8 {
	sum := uint8(0)
	logX := gf.log[x]

	if x == 0 {
		return s[0]
	}

	for i := 0; i < maxPoly; i++ {
		c := s[i]
		if c == 0 {
			continue
		}
		sum ^= gf.exp[(int(gf.log[c])+int(logX)*i)%gf.p]
	}
	return sum
}

/************************************************************************
 * Berlekamp-Massey algorithm for finding error locator polynomials.
 */

func berlekampMassey(s []uint8, num int, gf *galoisField, sigma []uint8) {

	C := make([]uint8, maxPoly)
	B := make([]uint8, maxPoly)

	L := 0
	m := 1
	b := uint8(1)

	B[0] = 1
	C[0] = 1

	for n := 0; n < num; n++ {
		d := s[n]

		for i := 1; i <= L; i++ {
			if !(C[i] > 0 && s[n-i] > 0) {
				continue
			}
			d ^= gf.exp[(int(gf.log[C[i]])+int(gf.log[s[n-i]]))%gf.p]
		}

		mult := gf.exp[(gf.p-int(gf.log[b])+int(gf.log[d]))%gf.p]

		if d == 0 {
			m++
		} else if L*2 <= n {
			T := make([]uint8, 0)
			T = append(T, C...)

			polyAdd(C, B, mult, m, gf)

			B = B[:0]
			B = append(B, T...)

			L = n + 1 - L
			b = d
			m = 1

		} else {
			polyAdd(C, B, mult, m, gf)
			m++
		}
	}
	sigma = sigma[:0]
	sigma = append(sigma, C...)
}

/************************************************************************
 * Code stream error correction
 *
 * Generator polynomial for GF(2^8) is x^8 + x^4 + x^3 + x^2 + 1
 */

func blockSyndromes(data []uint8, bs, npar int, s []uint8) int {
	nonzero := 0
	for i := 0; i < maxPoly; i++ {
		s[i] = 0
	}
	for i := 0; i < npar; i++ {
		for j := 0; j < bs; j++ {
			c := data[bs-j-1]
			if c == 0 {
				continue
			}
			s[i] ^= gf256Exp[((int)(gf256Log[c])+i*j)%255]
		}
		if s[i] != 0 {
			nonzero = 1
		}
	}
	return nonzero
}

func elocPoly(omega []uint8, s []uint8, sigma []uint8, npar int) {

	for i := 0; i < maxPoly; i++ {
		omega[i] = 0
	}

	for i := 0; i < npar; i++ {
		a := sigma[i]
		logA := gf256Log[a]

		if a == 0 {
			continue
		}

		for j := 0; j+1 < maxPoly; j++ {
			b := s[j+1]
			if i+j >= npar {
				break
			}

			if b == 0 {
				continue
			}
			omega[i+j] ^= gf256Exp[(int(logA)+int(gf256Log[b]))%255]
		}
	}
}

func correctBlock(data []uint8, ecc *qrRsParams) error {

	npar := ecc.bs - ecc.dw
	s := make([]uint8, maxPoly)

	// Compute syndrome vector
	if 0 == blockSyndromes(data, ecc.bs, npar, s) {
		return nil
	}

	sigma := make([]uint8, maxPoly)
	berlekampMassey(s, npar, &gf256, sigma)

	// Compute derivative of sigma
	sigmaDeriv := make([]uint8, maxPoly)
	for i := 0; i+1 < maxPoly; i += 2 {
		sigmaDeriv[i] = sigma[i+1]
	}

	// Compute error evaluator polynomial
	omega := make([]uint8, maxPoly)
	elocPoly(omega, s, sigma, npar-1)

	// Find error locations and magnitudes
	for i := 0; i < ecc.bs; i++ {
		xinv := gf256Exp[255-i]
		if 0 == polyEval(sigma, xinv, &gf256) {
			sdX := polyEval(sigmaDeriv, xinv, &gf256)
			omegaX := polyEval(omega, xinv, &gf256)
			error := gf256Exp[(255-int(gf256Log[sdX])+int(gf256Log[omegaX]))%255]
			data[ecc.bs-i-1] ^= error
		}
	}

	if blockSyndromes(data, ecc.bs, npar, s) != 0 {
		return ErrDataEcc
	}
	return nil
}

/************************************************************************
 * Format value error correction
 *
 * Generator polynomial for GF(2^4) is x^4 + x + 1
 */

const formatMaxError = 3
const formatSyndrome = formatMaxError * 2
const formatBits = 15

func formatSyndromes(u uint16, s []uint8) int {
	nonzero := 0
	for i := 0; i < maxPoly; i++ {
		s[i] = 0
	}
	for i := 0; i < formatSyndrome; i++ {
		s[i] = 0
		for j := 0; j < formatBits; j++ {
			if u&(uint16(1<<uint(j))) != 0 {
				s[i] ^= gf16Exp[((i+1)*j)%15]
			}
		}
		if s[i] != 0 {
			nonzero = 1
		}
	}
	return nonzero
}

func correctFormat(fRet *uint16) error {
	u := *fRet

	// Evaluate U (received codeword) at each of alpha_1 .. alpha_6
	// to get S_1 .. S_6 (but we index them from 0).

	s := make([]uint8, maxPoly)
	if 0 == formatSyndromes(u, s) {
		return nil
	}

	sigma := make([]uint8, maxPoly)
	berlekampMassey(s, formatSyndrome, &gf16, sigma)

	// Now, find the roots of the polynomial
	for i := 0; i < 15; i++ {
		if 0 == polyEval(sigma, gf16Exp[15-i], &gf16) {
			u ^= 1 << uint16(i)
		}
	}
	if 0 != formatSyndromes(u, s) {
		return ErrFormatEcc
	}

	*fRet = u
	return nil
}

/************************************************************************
 * Decoder algorithm
 */

type datastream struct {
	raw      [qrMaxPayload]uint8
	dataBits int
	ptr      int
	data     [qrMaxPayload]uint8
}

func maskBit(mask, i, j int) int {
	k := 0
	switch mask {
	case 0:
		k = (i + j) % 2
	case 1:
		k = i % 2
	case 2:
		k = j % 3
	case 3:
		k = (i + j) % 3
	case 4:
		k = ((i / 2) + (j / 3)) % 2
	case 5:
		k = (i*j)%2 + (i*j)%3
	case 6:
		k = ((i*j)%2 + (i*j)%3) % 2
	case 7:
		k = ((i*j)%3 + (i+j)%2) % 2
	default:
		return 0
	}
	if k != 0 {
		return 0
	}
	return 1
}

func reservedCell(version, i, j int) int {
	ver := &qrVersionDb[version]
	size := version*4 + 17
	ai := -1
	aj := -1
	a := 0

	// Finder + format: top left
	if i < 9 && j < 9 {
		return 1
	}

	// Finder + format: bottom left
	if i+8 >= size && j < 9 {
		return 1
	}

	// Finder + format: top right
	if i < 9 && j+8 >= size {
		return 1
	}
	// Exclude timing patterns
	if i == 6 || j == 6 {
		return 1
	}

	// Exclude Version info, if it exists. Version info sits adjacent to
	// the top-right and bottom-left finders in three rows, bounded by
	// the timing pattern.

	if version >= 7 {
		if i < 6 && j+11 >= size {
			return 1
		}

		if i+11 >= size && j < 6 {
			return 1
		}
	}
	// Exclude alignment patterns
	for a = 0; a < qrMaxAliment && ver.apat[a] != 0; a++ {
		p := ver.apat[a]

		if int(math.Abs(float64(p-i))) < 3 {
			ai = a
		}
		if int(math.Abs(float64(p-j))) < 3 {
			aj = a
		}
	}

	if ai >= 0 && aj >= 0 {
		a--
		if ai > 0 && ai < a {
			return 1
		}
		if aj > 0 && aj < a {
			return 1
		}
		if aj == a && ai == a {
			return 1
		}
	}
	return 0
}

func codestreamEcc(data *QRData, ds *datastream) error {
	ver := &qrVersionDb[data.Version]
	sbEcc := &ver.ecc[data.EccLevel]
	var lbEcc qrRsParams
	lbCount := (ver.dataBytes - sbEcc.bs*sbEcc.ns) / (sbEcc.bs + 1)

	bc := lbCount + sbEcc.ns
	eccOffset := sbEcc.dw*bc + lbCount

	dstOffset := 0
	lbEcc = *sbEcc

	lbEcc.dw++
	lbEcc.bs++

	for i := 0; i < bc; i++ {
		dst := ds.data[dstOffset:]
		ecc := sbEcc
		if i < sbEcc.ns {
			ecc = sbEcc
		} else {
			ecc = &lbEcc
		}
		numEc := ecc.bs - ecc.dw

		for j := 0; j < ecc.dw; j++ {
			dst[j] = ds.raw[j*bc+i]
		}

		for j := 0; j < numEc; j++ {
			dst[ecc.dw+j] = ds.raw[eccOffset+j*bc+i]
		}

		err := correctBlock(dst, ecc)
		if err != nil {
			return err
		}

		dstOffset += ecc.dw
	}

	ds.dataBits = dstOffset * 8
	return nil
}

func bitsRemaining(ds *datastream) int {
	return ds.dataBits - ds.ptr
}

func takeBits(ds *datastream, len int) int {
	ret := 0
	for len > 0 && (ds.ptr < ds.dataBits) {
		b := ds.data[ds.ptr>>3]
		bitpos := ds.ptr & 7
		ret <<= 1
		if ((b << uint(bitpos)) & 0x80) != 0 {
			ret |= 1
		}
		ds.ptr++
		len--
	}
	return ret
}

func numericTuple(data *QRData, ds *datastream, bits, digits int) int {
	if bitsRemaining(ds) < bits {
		return -1
	}
	tuple := takeBits(ds, bits)

	for i := digits - 1; i >= 0; i-- {
		data.Payload = append(data.Payload, uint8(tuple%10)+uint8('0'))

		tuple /= 10
	}
	return 0
}

func decodeNumeric(data *QRData, ds *datastream) error {
	bits := 14
	if data.Version < 10 {
		bits = 10
	} else if data.Version < 27 {
		bits = 12
	}

	count := takeBits(ds, bits)
	if len(data.Payload)+count+1 > qrMaxPayload {
		return ErrDataOverflow
	}
	for count >= 3 {
		if numericTuple(data, ds, 10, 3) < 0 {
			return ErrDataUnderflow
		}
		count -= 3
	}
	if count >= 2 {
		if numericTuple(data, ds, 7, 2) < 0 {
			return ErrDataUnderflow

		}
		count -= 2
	}
	if count > 0 {
		if numericTuple(data, ds, 4, 1) < 0 {
			return ErrDataUnderflow
		}
	}
	return nil
}

var alphaMap = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")

func alphaTuple(data *QRData, ds *datastream, bits, digits int) int {
	if bitsRemaining(ds) < bits {
		return -1
	}
	tuple := takeBits(ds, bits)
	data.Payload = append(data.Payload, make([]uint8, digits)...)
	payloadLen := len(data.Payload)
	for i := 0; i < digits; i++ {
		if payloadLen+digits-i-1 < len(data.Payload) {
			data.Payload[payloadLen+digits-i-1] = alphaMap[tuple%45]
			tuple /= 45
		}
	}
	return 0
}

func decodeAlpha(data *QRData, ds *datastream) error {
	bits := 13
	if data.Version < 10 {
		bits = 9
	} else if data.Version < 27 {
		bits = 11
	}
	count := takeBits(ds, bits)
	if len(data.Payload)+count+1 > qrMaxPayload {
		return ErrDataOverflow
	}
	for count >= 2 {
		if alphaTuple(data, ds, 11, 2) < 0 {
			return ErrDataUnderflow
		}
		count -= 2
	}
	if count > 0 {
		if alphaTuple(data, ds, 6, 1) < 0 {
			return ErrDataUnderflow
		}
	}
	return nil
}

func decodeByte(data *QRData, ds *datastream) error {
	bits := 16
	if data.Version < 10 {
		bits = 8
	}
	count := takeBits(ds, bits)
	if len(data.Payload)+count+1 > qrMaxPayload {
		return ErrDataOverflow
	}
	if bitsRemaining(ds) < count*8 {
		return ErrDataUnderflow
	}
	for i := 0; i < count; i++ {
		data.Payload = append(data.Payload, uint8(takeBits(ds, 8)))
	}
	return nil
}

func decodeKanji(data *QRData, ds *datastream) error {
	bits := 12

	if data.Version < 10 {
		bits = 8
	} else if data.Version < 27 {
		bits = 10
	}
	count := takeBits(ds, bits)

	if len(data.Payload)+count*2+1 > qrMaxPayload {
		return ErrDataOverflow
	}

	if bitsRemaining(ds) < count*13 {
		return ErrDataUnderflow
	}

	for i := 0; i < count; i++ {
		d := takeBits(ds, 13)
		msB := d / 0xc0
		lsB := d % 0xc0
		intermediate := uint16((msB << 8) | lsB)
		var sjw uint16

		if intermediate+0x8140 <= 0x9ffc {
			// bytes are in the range 0x8140 to 0x9FFC
			sjw = intermediate + 0x8140
		} else {
			// bytes are in the range 0xE040 to 0xEBBF
			sjw = intermediate + 0xc140
		}
		data.Payload = append(data.Payload, uint8(sjw>>8))
		data.Payload = append(data.Payload, uint8(sjw&0xff))
	}
	return nil
}

func decodeEci(data *QRData, ds *datastream) error {
	if bitsRemaining(ds) < 8 {
		return ErrDataOverflow
	}
	data.Eci = uint32(takeBits(ds, 8))
	if (data.Eci & 0xc0) == 0x80 {
		if bitsRemaining(ds) < 8 {
			return ErrDataUnderflow
		}
		data.Eci = (data.Eci << 8) | uint32(takeBits(ds, 8))
	} else if (data.Eci & 0xe0) == 0xc0 {
		if bitsRemaining(ds) < 16 {
			return ErrDataUnderflow
		}
		data.Eci = (data.Eci << 16) | uint32(takeBits(ds, 16))
	}
	return nil
}

func decodePayload(data *QRData, ds *datastream) error {

	for bitsRemaining(ds) >= 4 {
		var err error
		_type := takeBits(ds, 4)
		switch _type {
		case qrDataTypeNumeric:
			err = decodeNumeric(data, ds)
		case qrDataTypeAlpha:
			err = decodeAlpha(data, ds)
		case qrDataTypeByte:
			err = decodeByte(data, ds)
		case qrDataTypeKanji:
			err = decodeKanji(data, ds)
		case 7:
			err = decodeEci(data, ds)
		default:
			goto done
		}

		if err != nil {
			return err
		}
		if 0 == (_type&(_type-1)) && (_type > data.DataType) {
			data.DataType = _type
		}
	}

done:
	return nil
}
