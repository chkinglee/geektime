// Package utils
// @Author      : lilinzhen
// @Time        : 2022/3/27 19:45:16
// @Description :
package utils

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader         = "geektime"
	ConstHeaderLength   = 8
	ConstSaveDataLength = 4

	ConstFixLength         = 100
	ConstFixLengthFullChar = "@"

	ConstDelimiter = '@'
)

func PacketWithHeader(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

func UnpackWithHeader(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i++ {
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			readerChannel <- data

			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

func PacketWithFixLength(message []byte) []byte {
	if len(message) < ConstFixLength {
		emptyLength := ConstFixLength - len(message)
		for emptyLength > 0 {
			emptyLength--
			message = append(message, []byte(ConstFixLengthFullChar)...)
		}
	}
	return message
}

func UnpackWithFixLength(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)
	for length >= ConstFixLength {
		length -= ConstFixLength
		message := buffer[:ConstFixLength]
		for i, ch := range message {
			if string(ch) == ConstFixLengthFullChar {
				readerChannel <- message[:i]
				break
			}
		}
		buffer = buffer[ConstFixLength:]
	}
	return buffer
}

func PacketWithDelimiter(message []byte) []byte {
	return append(message, ConstDelimiter)
}

// IntToBytes 整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
