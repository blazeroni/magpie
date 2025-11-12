# magpie

[![Go Reference](https://pkg.go.dev/badge/github.com/blazeroni/magpie/.svg)](https://pkg.go.dev/github.com/blazeroni/magpie/)
[![Go Report Card](https://goreportcard.com/badge/github.com/blazeroni/magpie)](https://goreportcard.com/report/github.com/blazeroni/magpie)
[![codecov](https://codecov.io/github/blazeroni/magpie/graph/badge.svg?token=ZMLR3TBDY5)](https://codecov.io/github/blazeroni/magpie)
[![blazingly fast](https://blazingly.fast/api/badge.svg?repo=blazeroni%2Fmagpie)](https://blazingly.fast)

**magpie** is a high-performance, expressive image blending and composition library in pure Go.

## Installation

```sh
go get github.com/blazeroni/magpie
```

## Quick Start

Here's a simple example of blending two images using the `Multiply` blend mode.

```go
package main

import (
	"image"
	"log"

	"github.com/blazeroni/magpie/pkg/blend"
	"github.com/blazeroni/magpie/pkg"
)

func main() {
	// 1. Load destination and source images.
	dst := loadImage("background.png")
	src := loadImage("foreground.png")

	// 2. Define the blend operation.
	op := blend.Multiply()

	// 3. Apply the operation.
	output, err := magpie.Draw(dst, dst.Bounds(), src, image.Point{}, op, magpie.ToNewImage())
	if err != nil {
		log.Fatal(err)
	}

	// 4. Save the output.
	saveImage("output.png", output)
}
```

## API Concepts

*   **`magpie.Draw`**: The primary entry point for all drawing operations.
*   **`op.Op`**: Defines the operation to be performed (e.g., `op.BlendOp`, `op.CompositeOp`).
*   **`magpie.Context`**: For advanced use cases, a `Context` can be created to control concurrency and other settings.

> [!NOTE]
> The project is still in early stages. Breaking API changes may occur leading up to a stable release.

## Future

* Improve README & project documentation
* Benchmarks and comparisons with other libraries
* Additional image processing operations
  * HSV blending
  * Point operations: (effects, filters, etc.)
  * Convolutions: (blur, sharpen, etc.)
* Additional optimizations

## Contributing

Bug reports and feedback are welcome!

## License

This project is licensed under the [MIT License](LICENSE).

---

### Trivia

"magpie" is derived from the first six letters of "<b><u>image</u> <u>p</u></b>rocessing".
