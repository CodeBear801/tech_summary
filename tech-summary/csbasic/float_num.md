# Float number

How computer represent float point  
Decimal(fragment) => binary => float32/float64

## Conversion

### Binary fraction to decimal

float_exmaple_binary2decimal

### Decimal fraction to binary

(A great example from [quora](https://www.quora.com/How-do-I-convert-the-decimal-fraction-to-binary-with-a-maximum-of-6-places-to-the-right-of-the-radix-point-example-33-90625))

float_example_decimalfraction2binary

## Storage the value of float

### float32

**Single precision**, which uses 32 bits and has the following layout:
- 1 bit for the sign of the number. 0 means positive and 1 means negative.
- 8 bits for the exponent.
- 23 bits for the mantissa.
