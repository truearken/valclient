# ValClient Go Module

A Go module that wraps the Valorant game API.

## Installation

To use the `valclient` module in your Go project, you can install it using `go get`:

```bash
go get github.com/truearken/valclient
```

## Usage

Here's a basic example of how to use the `valclient` module in your Go application:

```go
package main

import (
	"log"

	"github.com/truearken/valclient/valclient"
)

func main() {
	client, err := valclient.NewClient(valclient.REGION_EU)
	if err != nil {
		panic(err)
	}

	loadout, err := client.GetPlayerLoadout()
	if err != nil {
		panic(err)
	}

	log.Print(loadout)
}
```

## Features

- Fetch player details
- Fetch and set loadout
- Retrieve match history
- And more...

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

This project is inspired by [valclient.py](https://github.com/colinhartigan/valclient.py) by Colin Hartigan.
