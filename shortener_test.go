package main

import (
    "testing"
)

const iterations = 100

func TestEncodeDecode(t *testing.T) {
    for i:=0; i<iterations; i++ {
        if Decode(Encode(i)) != i {
            t.Errorf("Encode() Decode() Error")
        }
    }
}
