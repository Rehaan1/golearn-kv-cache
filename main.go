package main

import (
	"fmt"
	"kv-file-cache/kvcache"
)

func main() {

	var c kvcache.Cache

	c.Set("username", []byte("rehaan"))
	c.Set("role", []byte("admin"))

	// Get an existing key
	if value, ok := c.Get("username"); ok {
		fmt.Printf("username = %s\n", value)
	} else {
		fmt.Println("username not found")
	}

	// Get a missing key — no error, just false
	if _, ok := c.Get("email"); !ok {
		fmt.Println("email not found (as expected)")
	}

	// Overwrite an existing key
	c.Set("role", []byte("staff engineer"))
	value, _ := c.Get("role")
	fmt.Printf("role updated to = %s\n", value)

	// Set an empty value — should still be a hit, not treated as missing
	c.Set("note", []byte{})
	if value, ok := c.Get("note"); ok {
		fmt.Printf("note = %q (len=%d), present=%v\n", value, len(value), ok)
	}
}
