# cipg

Command line parameter generator.

## Usage

1. Create a option struct for your application.
2. Put some strings into the tag for introducing usage.
3. Instantiate the option and put the default value into it.
4. Call `cipg.Generate(...)` to generate command line parameters according to the option.

## Supported base types

```go
bool
int
int64
uint
uint64
string
float64
struct
```

## Example

```go
type Option struct {
	DurationOpt time.Duration `usage:"Option used for time.Duration variable."`
	BoolOpt     bool          `usage:"Option used for bool variable."`
	IntOpt      int           `usage:"Option used for int variable."`
	Int64Opt    int64         `usage:"Option used for int64 variable."`
	UintOpt     uint          `usage:"Option used for uint variable."`
	Uint64Opt   uint64        `usage:"Option used for uint64 variable."`
	StringOpt   string        `usage:"Option used for string variable."`
	Float64Opt  float64       `usage:"Option used for float64 variable."`
	StructOpt   StructOption  `usage:"Option used for struct variable."`
}

type StructOption struct {
	DurationOpt time.Duration `usage:"Option used for time.Duration variable in struct."`
	BoolOpt     bool          `usage:"Option used for bool variable in struct."`
	IntOpt      int           `usage:"Option used for int variable in struct."`
	Int64Opt    int64         `usage:"Option used for int64 variable in struct."`
	UintOpt     uint          `usage:"Option used for uint variable in struct."`
	Uint64Opt   uint64        `usage:"Option used for uint64 variable in struct."`
	StringOpt   string        `usage:"Option used for string variable in struct."`
	Float64Opt  float64       `usage:"Option used for float64 variable in struct."`
}

func DefaultOption() Option {
	return Option{
		DurationOpt: time.Hour + time.Minute + time.Second,
		BoolOpt:     true,
		IntOpt:      123,
		Int64Opt:    456,
		UintOpt:     789,
		Uint64Opt:   123,
		StringOpt:   "456",
		Float64Opt:  789.123,
		StructOpt: StructOption{
			DurationOpt: time.Hour + time.Minute + time.Second,
			BoolOpt:     true,
			IntOpt:      456,
			Int64Opt:    789,
			UintOpt:     123,
			Uint64Opt:   456,
			StringOpt:   "789",
			Float64Opt:  123.456,
		},
	}
}

func main() {
	opt := DefaultOption()
	Generate(&opt, fmt.Println)
}
```