# 15 Rounds - Rook evades Bishop

Bishop stationary at `c3`, Rook starts at `h1`.  The Rook takes a random (board wrapping) walk (iterations) while attempting to evade the Bishop's reach.

## Development considerations

- Prompt asked for OOP, but GO seems to not have classes in the classical sense.  Instead a clever way of a Struct dedicated to a package and appropriately typed functions and interfaces.
- Interfaces, to achieve [Liskov Subtitution Principal](https://en.wikipedia.org/wiki/Liskov_substitution_principle): `type piece interface { ToString() string }` implement in the concrete 'class' `func (b bishop) ToString() string { return "I am a Bishop" }` then back at the 'piece' level define `func PrintName(p piece) { fmt.Println(p.ToString()) }`
- CanCapture() method, decided to put as a part of the Pieces interface, and implement at Bishop level

## Things I learned...

- Associate a function with a type: `func (e board) ExamplePrintln() {`