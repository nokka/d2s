package d2s

import "io"

// bitReader wraps an io.Reader and provides the ability to read bytes bit-by-bit from it.
type bitReader struct {
	r    io.ByteReader
	n    uint64
	bits uint
	err  error
}

func (br *bitReader) ReadByte() (byte, error) {
	return br.r.ReadByte()
}
func (br *bitReader) ReadBits(bits uint) (uint64, error) {
	var n uint64

	for bits > br.bits {
		b, err := br.r.ReadByte()
		if err == io.EOF {
			return n, io.ErrUnexpectedEOF
		}
		if err != nil {
			return n, err
		}
		b = reverseByte(b)
		br.n <<= 8
		br.n |= uint64(b)
		br.bits += 8
	}

	n = (br.n >> (br.bits - bits)) & ((1 << bits) - 1)
	n = reverseBits(n, bits)
	br.bits -= bits

	return n, nil
}

func reverseByte(b byte) byte {
	var d byte
	for i := 0; i < 8; i++ {
		d <<= 1
		d |= b & 1
		b >>= 1
	}
	return d
}

func reverseBits(b uint64, n uint) uint64 {
	var d uint64
	for i := 0; i < int(n); i++ {
		d <<= 1
		d |= b & 1
		b >>= 1
	}
	return d
}
