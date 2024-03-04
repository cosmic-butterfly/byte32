package byte32

type Byte32 [32]byte

// Return the standard byte array representation of the byte32
func (b *Byte32) Bytes() []byte {
	return b[:]
}

// Return the standard string representation of the byte32
func (b *Byte32) String() string {
	return string(b[:])
}

// Set the byte32 from a byte array
func (b *Byte32) SetBytes(bytes []byte) {
	copy(b[:], bytes)
}

// Set the byte32 from a string
func (b *Byte32) SetString(s string) {
	copy(b[:], s)
}

// Check if is zero
func (b *Byte32) IsZero() bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}

// Equality check
func (b *Byte32) IsEqual(b2 *Byte32) bool {
	return *b == *b2
}

// Check if is empty
func (b *Byte32) IsEmpty() bool {
	return b.IsZero()
}

// Clone the byte32
func (b *Byte32) Clone() *Byte32 {
	b2 := new(Byte32)
	copy(b2[:], b[:])
	return b2
}

// Clear the byte32
func (b *Byte32) Clear() {
	for i := range b {
		b[i] = 0
	}
}

// Copy into the byte32
func (b *Byte32) Copy(b2 *Byte32) {
	copy(b[:], b2[:])
}

// Set the byte32
func (b *Byte32) Set(b2 *Byte32) {
	copy(b[:], b2[:])
}

// Add two byte arrays
func (b *Byte32) Add(b2 *Byte32) {
	for i := range b {
		b[i] += b2[i]
	}
}

// Subtract two byte arrays
func (b *Byte32) Sub(b2 *Byte32) {
	for i := range b {
		b[i] -= b2[i]
	}
}

// Multiply two byte arrays
func (b *Byte32) Mul(b2 *Byte32) {
	for i := range b {
		b[i] *= b2[i]
	}
}

// Divide two byte arrays
func (b *Byte32) Div(b2 *Byte32) {
	for i := range b {
		b[i] /= b2[i]
	}
}

// Modulus two byte arrays
func (b *Byte32) Mod(b2 *Byte32) {
	for i := range b {
		b[i] %= b2[i]
	}
}

// Bitwise AND two byte arrays
func (b *Byte32) And(b2 *Byte32) {
	for i := range b {
		b[i] &= b2[i]
	}
}

// Bitwise OR two byte arrays
func (b *Byte32) Or(b2 *Byte32) {
	for i := range b {
		b[i] |= b2[i]
	}
}

// Bitwise XOR two byte arrays
func (b *Byte32) Xor(b2 *Byte32) {
	for i := range b {
		b[i] ^= b2[i]
	}
}

// Bitwise NOT the byte array
func (b *Byte32) Not() {
	for i := range b {
		b[i] = ^b[i]
	}
}

// Bitwise left shift the byte array
func (b *Byte32) Lsh(b2 *Byte32) {
	for i := range b {
		b[i] <<= b2[i]
	}
}

// Bitwise right shift the byte array
func (b *Byte32) Rsh(b2 *Byte32) {
	for i := range b {
		b[i] >>= b2[i]
	}
}

// Add a scalar to the byte array
func (b *Byte32) AddScalar(s uint8) {
	for i := range b {
		b[i] += s
	}
}

// Subtract a scalar from the byte array
func (b *Byte32) SubScalar(s uint8) {
	for i := range b {
		b[i] -= s
	}
}

// Multiply a scalar with the byte array
func (b *Byte32) MulScalar(s uint8) {
	for i := range b {
		b[i] *= s
	}
}

// Divide a scalar with the byte array
func (b *Byte32) DivScalar(s uint8) {
	for i := range b {
		b[i] /= s
	}
}

// Modulus a scalar with the byte array
func (b *Byte32) ModScalar(s uint8) {
	for i := range b {
		b[i] %= s
	}
}

// Bitwise AND a scalar with the byte array
func (b *Byte32) AndScalar(s uint8) {
	for i := range b {
		b[i] &= s
	}
}

// Bitwise OR a scalar with the byte array
func (b *Byte32) OrScalar(s uint8) {
	for i := range b {
		b[i] |= s
	}
}

// Bitwise XOR a scalar with the byte array
func (b *Byte32) XorScalar(s uint8) {
	for i := range b {
		b[i] ^= s
	}
}

// Bitwise left shift a scalar with the byte array
func (b *Byte32) LshScalar(s uint8) {
	for i := range b {
		b[i] <<= s
	}
}

// Bitwise right shift a scalar with the byte array
func (b *Byte32) RshScalar(s uint8) {
	for i := range b {
		b[i] >>= s
	}
}

// Add a scalar to the byte array and store in another byte array
func (b *Byte32) AddScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] + s
	}
}

// Subtract a scalar from the byte array and store in another byte array
func (b *Byte32) SubScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] - s
	}
}

// Multiply a scalar with the byte array and store in another byte array
func (b *Byte32) MulScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] * s
	}
}

// Divide a scalar with the byte array and store in another byte array
func (b *Byte32) DivScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] / s
	}
}

// Modulus a scalar with the byte array and store in another byte array
func (b *Byte32) ModScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] % s
	}
}

// Bitwise AND a scalar with the byte array and store in another byte array
func (b *Byte32) AndScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] & s
	}
}

// Bitwise OR a scalar with the byte array and store in another byte array
func (b *Byte32) OrScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] | s
	}
}

// Bitwise XOR a scalar with the byte array and store in another byte array
func (b *Byte32) XorScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] ^ s
	}
}

// Bitwise left shift a scalar with the byte array and store in another byte array
func (b *Byte32) LshScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] << s
	}
}

// Bitwise right shift a scalar with the byte array and store in another byte array
func (b *Byte32) RshScalarTo(s uint8, b2 *Byte32) {
	for i := range b {
		b2[i] = b[i] >> s
	}
}

// Add two byte arrays and store in another byte array
func (b *Byte32) AddTo(b2 *Byte32) {
	for i := range b {
		b2[i] += b[i]
	}
}
