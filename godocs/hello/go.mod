module hello

go 1.15

require (
	greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace greetings => ../greetings
