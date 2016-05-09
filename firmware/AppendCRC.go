package firmware

/*
   This is a direct implementation of the alorithm written from the Rockbox forum member ScaV (thank you!).

   Ref: http://forums.rockbox.org/index.php/topic,43525.msg230595.html#msg230595 (forum post)
   Ref: http://pastebin.com/raw/TXxT7157 (algorithm implementation in C)
*/

import (
	"bytes"
	"encoding/binary"
	"os"
)

// Size of the block of data to read
const blockSize int = 0x80000

// getBlockCRC calculates CRC for any single block given
func getBlockCRC(dataSize int, dataArray []byte, prevCRC int32) int32 {
	// magic values, got from the official Fiio RE work
	values := [256]int{
		0, 4129, 8258, 12387, 16516, 20645, 24774, 28903, 33032, 37161, 41290, 45419, 49548, 53677, 57806, 61935, 4657, 528, 12915, 8786, 21173, 17044, 29431, 25302, 37689, 33560, 45947, 41818, 54205, 50076, 62463, 58334, 9314, 13379, 1056, 5121, 25830, 29895, 17572, 21637, 42346, 46411, 34088, 38153, 58862, 62927, 50604, 54669, 13907, 9842, 5649, 1584, 30423, 26358, 22165, 18100, 46939, 42874, 38681, 34616, 63455, 59390, 55197, 51132, 18628, 22757, 26758, 30887, 2112, 6241, 10242, 14371, 51660, 55789, 59790, 63919, 35144, 39273, 43274, 47403, 23285, 19156, 31415, 27286, 6769, 2640, 14899, 10770, 56317, 52188, 64447, 60318, 39801, 35672, 47931, 43802, 27814, 31879, 19684, 23749, 11298, 15363, 3168, 7233, 60846, 64911, 52716, 56781, 44330, 48395, 36200, 40265, 32407, 28342, 24277, 20212, 15891, 11826, 7761, 3696, 65439, 61374, 57309, 53244, 48923, 44858, 40793, 36728, 37256, 33193, 45514, 41451, 53516, 49453, 61774, 57711, 4224, 161, 12482, 8419, 20484, 16421, 28742, 24679, 33721, 37784, 41979, 46042, 49981, 54044, 58239, 62302, 689, 4752, 8947, 13010, 16949, 21012, 25207, 29270, 46570, 42443, 38312, 34185, 62830, 58703, 54572, 50445, 13538, 9411, 5280, 1153, 29798, 25671, 21540, 17413, 42971, 47098, 34713, 38840, 59231, 63358, 50973, 55100, 9939, 14066, 1681, 5808, 26199, 30326, 17941, 22068, 55628, 51565, 63758, 59695, 39368, 35305, 47498, 43435, 22596, 18533, 30726, 26663, 6336, 2273, 14466, 10403, 52093, 56156, 60223, 64286, 35833, 39896, 43963, 48026, 19061, 23124, 27191, 31254, 2801, 6864, 10931, 14994, 64814, 60687, 56684, 52557, 48554, 44427, 40424, 36297, 31782, 27655, 23652, 19525, 15522, 11395, 7392, 3265, 61215, 65342, 53085, 57212, 44955, 49082, 36825, 40952, 28183, 32310, 20053, 24180, 11923, 16050, 3793, 7920,
	}

	result := int32(prevCRC)
	for i := 0; i < dataSize; i++ {
		/*
		   This is a tricky part.

		   Basically, for any given data block passed to this function, iterate through and calculate "result".

		   dataArray[i] == current rune of the block being worked on
		   (result>>8)&0xff) == change endianess to big?
		   (result << 8) == change endianess back to little?

		   All the data is parsed as an "int32", "int" is too large and makes us work with garbage data.
		*/
		result = int32(values[rune(dataArray[i])^rune((result>>8)&0xff)]) ^ (result << 8)
	}
	return result
}

// AppendCRC calculates the CRC for any given file and appends it to the end
func AppendCRC(filename string) error {
	fwFile, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	defer fwFile.Close()

	if err != nil {
		return err
	}

	var fileposBegin int
	var fileposEnd int
	var dataSize int
	var crc int32

	fileInfo, err := fwFile.Stat()

	if err != nil {
		return err
	}

	fileposEnd = int(fileInfo.Size())

	for fileposBegin < fileposEnd {
		dataArray := make([]byte, blockSize)

		if (fileposEnd - fileposBegin) < blockSize {
			dataSize = fileposEnd - fileposBegin
		} else {
			dataSize = blockSize
		}

		fwFile.Read(dataArray)

		crc = getBlockCRC(dataSize, dataArray, crc)
		fileposBegin += dataSize
	}

	fwFile.Seek(int64(os.SEEK_END), 0)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, crc)

	_, err = fwFile.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
