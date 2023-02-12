# Roll DX

D&D dice roller script from the Digital Ocean ['Build Your First Command Line Tool in Go'](https://www.digitalocean.com/community/tech-talks/build-your-first-command-line-tool-in-go) tutorial.

## Usage

The program can be run by entering `go run main.go` but using the flags listed below modifies the script's output.

### Flags

The script can be run with a number of flags:

* d - Set the number of faces for the dice. By default, this is set to d6 (a six sided dice)
* n - Set the number of times to roll the dice. By default this is 1
* s - Set to true to add up the values from each dice roll together
* adv - Set to true to get the higher value from the number of dice rolled
* dis - Set to true to get the lowest value from the number of dice rolled

**Example usage with flags**

`go run main.go -d d20 -n 3 -s=true -adv=true -dis=true`