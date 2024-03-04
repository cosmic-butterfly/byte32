# Byte32

This package provides a `Byte32` type in Go, which is a fixed-size array of 32 bytes. It includes methods for manipulating these arrays, including arithmetic operations, bitwise operations, and comparisons.

## Methods

Here are the main methods provided by the `Byte32` type:

- `Bytes() []byte`: Returns the byte array representation of the `Byte32`.
- `String() string`: Returns the string representation of the `Byte32`.
- `SetBytes(bytes []byte)`: Sets the `Byte32` from a byte array.
- `SetString(s string)`: Sets the `Byte32` from a string.
- `IsZero() bool`: Checks if the `Byte32` is zero.
- `IsEqual(b2 *Byte32) bool`: Checks if two `Byte32` are equal.
- `IsEmpty() bool`: Checks if the `Byte32` is empty.
- `Clone() *Byte32`: Returns a new `Byte32` that is a copy of the current one.
- `Clear()`: Sets all bytes in the `Byte32` to zero.
- `Copy(b2 *Byte32)`: Copies the bytes from another `Byte32` into the current one.
- `Set(b2 *Byte32)`: Sets the `Byte32` from another `Byte32`.

The package also includes methods for arithmetic and bitwise operations on `Byte32` instances, both with other `Byte32` instances and with scalar values. These operations include addition, subtraction, multiplication, division, modulus, bitwise AND, bitwise OR, bitwise XOR, bitwise NOT, bitwise left shift, and bitwise right shift.

Each operation has a corresponding method that performs the operation and stores the result in another `Byte32` instance.

## Usage

To use this package, import it in your Go code:

```go
go get github.com/cosmic-butterfly/byte32
```

Then, you can create a new `Byte32` and use its methods:

```go
b := new(byte32.Byte32)
b.SetString("Hello, world!")
fmt.Println(b.String()) // Outputs: "Hello, world!"
```

Please refer to the source code for more details on how to use these methods.
