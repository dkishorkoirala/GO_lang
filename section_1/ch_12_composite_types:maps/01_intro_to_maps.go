/*Introduction to Maps


Maps in Go are key-value pairs, similar to dictionaries in other languages. They let you store and retrieve values using unique keys.

Create a map that associates string names with integer ages:

ages := map[string]int{
    "Alice": 25,
    "Bob": 30,
}
Access a value using its key:

aliceAge := ages["Alice"]
fmt.Println(aliceAge)
The output will be:

25
Add a new key-value pair:

ages["Charlie"] = 22*/