// encodeData - encodes incoming data buffer into
// dataBlocks+parityBlocks returns a 2 dimensional byte array.
func encodeData(dataBuffer []byte, dataBlocks, parityBlocks int) ([][]byte, error) {
        rs, err := reedsolomon.New(dataBlocks, parityBlocks)
        if err != nil {
                return nil, traceError(err)
        }
        // Split the input buffer into data and parity blocks.
        var blocks [][]byte
        blocks, err = rs.Split(dataBuffer)
        if err != nil {
                return nil, traceError(err)
        }

        // Encode parity blocks using data blocks.
        err = rs.Encode(blocks)
        if err != nil {
                return nil, traceError(err)
        }

        // Return encoded blocks.
        return blocks, nil
}

