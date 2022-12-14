package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/Nik-U/pbc"
)

func main() {
	params := pbc.GenerateA(160, 512)
	pairing := params.NewPairing()

	M := big.NewInt(123456789)

	s := pairing.NewZr().Rand()
	P := pairing.NewG1().Rand()
	tem1 := pairing.NewGT().Rand()

	//Setup, system parameters generation
	Ppub := pairing.NewG1().MulZn(P, s)

	//Extract, key calculation
	Qid := pairing.NewG1().SetFromStringHash("cody", sha256.New())
	dID := pairing.NewG1().MulZn(Qid, s)

	//Encrypt encrypt M with ID
	r := pairing.NewZr().Rand()
	U := pairing.NewG1().MulZn(P, r)

	gID := pairing.NewGT().Pair(Qid, Ppub)
	gIDr := pairing.NewGT().MulZn(gID, r)

	z := gIDr.X()
	V := M.Xor(M, z)

	//Decrypt decrypt C = <U,V>
	tem1 = pairing.NewGT().Pair(dID, U)
	z = tem1.X()
	M = V.Xor(V, z)

	fmt.Println(M)
}
