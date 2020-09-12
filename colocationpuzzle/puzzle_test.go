package colocationpuzzle

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/cachecashproject/go-cachecash/testutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ColocationPuzzleTestSuite struct {
	suite.Suite

	plaintextBlocks  [][]byte
	ciphertextBlocks [][]byte
	innerKeys        [][]byte
	innerIVs         [][]byte
	params           Parameters
}

const (
	BlockQty = 8
)

func TestColocationPuzzleTestSuite(t *testing.T) {
	suite.Run(t, new(ColocationPuzzleTestSuite))
}

// XXX: TODO: Add a test case that covers the single-block case; the implementation fails!

func (suite *ColocationPuzzleTestSuite) SetupSuite() {
	t := suite.T()

	suite.params = Parameters{
		Rounds:      2,
		StartOffset: 0, // XXX: Not used yet.
		StartRange:  0,
	}

	for i := 0; i < BlockQty; i++ {
		k := testutil.RandBytes(16)
		suite.innerKeys = append(suite.innerKeys, k)

		iv := testutil.RandBytes(aes.BlockSize)
		suite.innerIVs = append(suite.innerIVs, iv)

		// TODO: The blocks need not all be the same size.
		b := testutil.RandBytes(aes.BlockSize)
		suite.plaintextBlocks = append(suite.plaintextBlocks, b)

		// Set up our AES-CTR encryption gizmo.
		block, err := aes.NewCipher(k)
		if err != nil {
			t.Fatal(errors.Wrap(err, "failed to create block cipher"))
		}
		stream := cipher.NewCTR(block, iv)

		// Encrypt the plaintext block.
		cb := make([]byte, len(b))
		stream.XORKeyStream(cb, b)
		suite.ciphertextBlocks = append(suite.ciphertextBlocks, cb)
	}

	fmt.Printf("ciphertextblocks len=%v", len(suite.ciphertextBlocks))
}

func (suite *ColocationPuzzleTestSuite) TestGenerateAndSolve() {
	suite.testGenerateAndSolve(0, 4)
}

func (suite *ColocationPuzzleTestSuite) TestGenerateAndSolveWithOffset() {
	// This case differs from the above because block indices do not match the indices into the key/IV slices (they are
	// offset by 4).
	suite.testGenerateAndSolve(4, 8)
}

func (suite *ColocationPuzzleTestSuite) TestGenerateAndSolveSingleBlock() {
	suite.testGenerateAndSolve(0, 1)
}

func (suite *ColocationPuzzleTestSuite) testGenerateAndSolve(rangeBegin, rangeEnd int) {
	t := suite.T()

	puzzle, err := Generate(suite.params, suite.plaintextBlocks[rangeBegin:rangeEnd],
		suite.innerKeys[rangeBegin:rangeEnd], suite.innerIVs[rangeBegin:rangeEnd])
	if !assert.Nil(t, err, "failed to generate puzzle") {
		return
	}

	assert.Equal(t, sha512.Size384, len(puzzle.Secret), "unexpected secret length")
	assert.Equal(t, sha512.Size384, len(puzzle.Goal), "unxpected goal length")

	secret, offset, err := Solve(suite.params, suite.ciphertextBlocks[rangeBegin:rangeEnd], puzzle.Goal)
	if !assert.Nil(t, err, "failed to solve puzzle") {
		return
	}

	// It's only actually important that the secret be the same on both sides, but it would be very odd if multiple
	// offsets led to correct solutions to the puzzle.
	assert.Equal(t, puzzle.Offset, offset, "generator and solver do not agree on starting offset")
	assert.Equal(t, puzzle.Secret, secret, "generator and solver do not agree on secret")
}

func (suite *ColocationPuzzleTestSuite) TestRunPuzzle() {
	t := suite.T()

	chunks := [][]byte{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
		{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63},
	}
	rounds := uint32(2)
	offset := uint32(2)

	expectedGoal := []byte{0x11, 0x88, 0x53, 0xf9, 0x6d, 0xc6, 0x70, 0xe9, 0xd6, 0x6a, 0xab, 0xee, 0xf3, 0x4a, 0xed, 0x53, 0x5d, 0x2, 0xd2, 0xa9, 0x2b, 0xf0, 0xe0, 0x80, 0x9e, 0xc9, 0xb3, 0x12, 0xcd, 0xa0, 0x83, 0xfc, 0x5a, 0x5c, 0x94, 0x7c, 0xef, 0xba, 0xd7, 0x68, 0xe2, 0x3f, 0x64, 0xef, 0xd8, 0x8, 0x87, 0x20}
	expectedSecret := []byte{0x58, 0x72, 0x17, 0xdd, 0x1e, 0xfd, 0x61, 0x12, 0xb2, 0xb5, 0xb6, 0x41, 0xd2, 0x7a, 0xa5, 0xfd, 0x47, 0x2f, 0x27, 0xb6, 0x8f, 0x19, 0x4b, 0x8c, 0x2f, 0x9, 0x2, 0x9e, 0xdb, 0x63, 0xca, 0x5f, 0x2b, 0xf4, 0xd0, 0x91, 0x6b, 0xbc, 0x26, 0xa2, 0x92, 0x92, 0xe3, 0x11, 0xae, 0x5a, 0xb5, 0x18}

	getBlockFn := func(chunkIdx, offset uint32) ([]byte, error) {
		offset = offset % uint32(len(chunks[chunkIdx])/aes.BlockSize)
		return chunks[chunkIdx][offset*aes.BlockSize : (offset+1)*aes.BlockSize], nil
	}
	goal, secret, err := runPuzzle(rounds, uint32(len(chunks)), offset, getBlockFn, dCrypto)

	assert.Nil(t, err)
	assert.Equal(t, expectedGoal, goal, "goal mismatched")
	assert.Equal(t, expectedSecret, secret, "secret mismatched")
}

func (suite *ColocationPuzzleTestSuite) TestRunPuzzleMD5() {
	start := time.Now()

	t := suite.T()

	chunks := [][]byte{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
		{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63},
	}
	rounds := uint32(2)
	offset := uint32(2)

	expectedGoal := []byte{0x14, 0xd8, 0xb3, 0xc0, 0xfb, 0x18, 0x85, 0x8b, 0xb9, 0x19, 0x2d, 0x7e, 0xf3, 0x99, 0xd9, 0xc8, 0x14, 0xd8, 0xb3, 0xc0, 0xfb, 0x18, 0x85, 0x8b, 0xb9, 0x19, 0x2d, 0x7e, 0xf3, 0x99, 0xd9, 0xc8, 0x14, 0xd8, 0xb3, 0xc0, 0xfb, 0x18, 0x85, 0x8b, 0xb9, 0x19, 0x2d, 0x7e, 0xf3, 0x99, 0xd9, 0xc8}
	expectedSecret := []byte{0x91, 0xc5, 0xac, 0xb0, 0x47, 0xa8, 0x16, 0xdf, 0x1d, 0xfd, 0xbf, 0xdb, 0x2e, 0x6f, 0xb2, 0x63, 0x91, 0xc5, 0xac, 0xb0, 0x47, 0xa8, 0x16, 0xdf, 0x1d, 0xfd, 0xbf, 0xdb, 0x2e, 0x6f, 0xb2, 0x63, 0x91, 0xc5, 0xac, 0xb0, 0x47, 0xa8, 0x16, 0xdf, 0x1d, 0xfd, 0xbf, 0xdb, 0x2e, 0x6f, 0xb2, 0x63}

	getBlockFn := func(chunkIdx, offset uint32) ([]byte, error) {
		offset = offset % uint32(len(chunks[chunkIdx])/aes.BlockSize)
		return chunks[chunkIdx][offset*aes.BlockSize : (offset+1)*aes.BlockSize], nil
	}
	goal, secret, err := runPuzzle(rounds, uint32(len(chunks)), offset, getBlockFn, "md5")

	assert.Nil(t, err)
	assert.Equal(t, expectedGoal, goal, "goal mismatched")
	assert.Equal(t, expectedSecret, secret, "secret mismatched")
	elapsed := time.Since(start)
	log.Printf("MD5 Test took %s", elapsed)
}

func (suite *ColocationPuzzleTestSuite) TestRunPuzzleCRC32() {
	start := time.Now()

	t := suite.T()

	chunks := [][]byte{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
		{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63},
	}
	rounds := uint32(2)
	offset := uint32(2)

	expectedGoal := []byte{0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4, 0xd1, 0x91, 0xb, 0x4}
	expectedSecret := []byte{0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13, 0x3c, 0xfb, 0xec, 0x13}

	getBlockFn := func(chunkIdx, offset uint32) ([]byte, error) {
		offset = offset % uint32(len(chunks[chunkIdx])/aes.BlockSize)
		return chunks[chunkIdx][offset*aes.BlockSize : (offset+1)*aes.BlockSize], nil
	}
	goal, secret, err := runPuzzle(rounds, uint32(len(chunks)), offset, getBlockFn, "crc32")

	assert.Nil(t, err)
	assert.Equal(t, expectedGoal, goal, "goal mismatched")
	assert.Equal(t, expectedSecret, secret, "secret mismatched")
	elapsed := time.Since(start)
	log.Printf("CRC32 Test took %s", elapsed)
}

func (suite *ColocationPuzzleTestSuite) TestRunPuzzleSHA512() {
	start := time.Now()

	t := suite.T()

	chunks := [][]byte{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
		{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63},
	}
	rounds := uint32(2)
	offset := uint32(2)

	expectedGoal := []byte{0x11, 0x88, 0x53, 0xf9, 0x6d, 0xc6, 0x70, 0xe9, 0xd6, 0x6a, 0xab, 0xee, 0xf3, 0x4a, 0xed, 0x53, 0x5d, 0x2, 0xd2, 0xa9, 0x2b, 0xf0, 0xe0, 0x80, 0x9e, 0xc9, 0xb3, 0x12, 0xcd, 0xa0, 0x83, 0xfc, 0x5a, 0x5c, 0x94, 0x7c, 0xef, 0xba, 0xd7, 0x68, 0xe2, 0x3f, 0x64, 0xef, 0xd8, 0x8, 0x87, 0x20}
	expectedSecret := []byte{0x58, 0x72, 0x17, 0xdd, 0x1e, 0xfd, 0x61, 0x12, 0xb2, 0xb5, 0xb6, 0x41, 0xd2, 0x7a, 0xa5, 0xfd, 0x47, 0x2f, 0x27, 0xb6, 0x8f, 0x19, 0x4b, 0x8c, 0x2f, 0x9, 0x2, 0x9e, 0xdb, 0x63, 0xca, 0x5f, 0x2b, 0xf4, 0xd0, 0x91, 0x6b, 0xbc, 0x26, 0xa2, 0x92, 0x92, 0xe3, 0x11, 0xae, 0x5a, 0xb5, 0x18}

	getBlockFn := func(chunkIdx, offset uint32) ([]byte, error) {
		offset = offset % uint32(len(chunks[chunkIdx])/aes.BlockSize)
		return chunks[chunkIdx][offset*aes.BlockSize : (offset+1)*aes.BlockSize], nil
	}
	goal, secret, err := runPuzzle(rounds, uint32(len(chunks)), offset, getBlockFn, "sha384")

	assert.Nil(t, err)
	assert.Equal(t, expectedGoal, goal, "goal mismatched")
	assert.Equal(t, expectedSecret, secret, "secret mismatched")
	elapsed := time.Since(start)
	log.Printf("SHA512 Test took %s", elapsed)
}

func (suite *ColocationPuzzleTestSuite) TestSolutionVerify() {
	t := suite.T()

	chunks := [][]byte{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
		{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63},
	}
	expectedGoal := []byte{0x11, 0x88, 0x53, 0xf9, 0x6d, 0xc6, 0x70, 0xe9, 0xd6, 0x6a, 0xab, 0xee, 0xf3, 0x4a, 0xed, 0x53, 0x5d, 0x2, 0xd2, 0xa9, 0x2b, 0xf0, 0xe0, 0x80, 0x9e, 0xc9, 0xb3, 0x12, 0xcd, 0xa0, 0x83, 0xfc, 0x5a, 0x5c, 0x94, 0x7c, 0xef, 0xba, 0xd7, 0x68, 0xe2, 0x3f, 0x64, 0xef, 0xd8, 0x8, 0x87, 0x20}
	expectedSecret := []byte{0x58, 0x72, 0x17, 0xdd, 0x1e, 0xfd, 0x61, 0x12, 0xb2, 0xb5, 0xb6, 0x41, 0xd2, 0x7a, 0xa5, 0xfd, 0x47, 0x2f, 0x27, 0xb6, 0x8f, 0x19, 0x4b, 0x8c, 0x2f, 0x9, 0x2, 0x9e, 0xdb, 0x63, 0xca, 0x5f, 0x2b, 0xf4, 0xd0, 0x91, 0x6b, 0xbc, 0x26, 0xa2, 0x92, 0x92, 0xe3, 0x11, 0xae, 0x5a, 0xb5, 0x18}
	offset := uint32(2)

	secret, offset, err := VerifySolution(suite.params, chunks, expectedGoal, offset)
	assert.Equal(t, expectedSecret, secret)
	assert.Equal(t, uint32(2), offset)
	assert.Nil(t, err, "solution is invalid")
}

func (suite *ColocationPuzzleTestSuite) TestKeyAndIVFromSecret() {
	t := suite.T()

	puzzle, err := Generate(suite.params, suite.plaintextBlocks, suite.innerKeys, suite.innerIVs)
	if !assert.Nil(t, err, "failed to generate puzzle") {
		return
	}

	key := puzzle.Key()
	iv := puzzle.IV()

	assert.Equal(t, KeySize, len(key))
	assert.Equal(t, IVSize, len(iv))
}
