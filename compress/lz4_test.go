package compress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLz4(t *testing.T) {
	data := `func CloneObject(a, b interface{}) []byte {
		buff := new(bytes.Buffer)
		enc := gob.NewEncoder(buff)
		dec := gob.NewDecoder(buff)
		err := enc.Encode(a)
		if err != nil {
			log.Panic("e1: ", err)
		}
		b1 := buff.Bytes()
		err = dec.Decode(b)
		if err != nil {
			log.Panic("e2: ", err)
		}
		return b1
	}`

	lz4 := GetLz4Compresser()

	compressRes, err := lz4.Compress([]byte(data))
	assert.Equal(t, err, nil)

	t.Log(len(compressRes), len(data))

	dataNew, err := lz4.UnCompress(compressRes)
	assert.Equal(t, err, nil)

	assert.Equal(t, []byte(data), dataNew)
}