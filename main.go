/** snowflake algorithm **/
package main

import (
	"fmt"
	"time"
)

var lastTimeStamp uint64
var sequence uint64

const (
	SEQUENCE_MASK    uint64 = 0xfff /** 4095 **/
	WORKERID_OFFSET  uint64 = 12
	TIMESTAMP_OFFSET uint64 = 22
)

func main() {
	for j := 0; j < 30; j++ {
		for i := 0; i < 1000; i++ {
			fmt.Println(nextId(uint64(j)))
		}
	}

}

/** 生成方法 **/
func nextId(workerid uint64) uint64 {
	/** ms **/
	var timeStamp uint64
	timeStamp = getTimestampUint64()
	if timeStamp < lastTimeStamp {
		panic("system clock wrong")
	}
	if timeStamp == lastTimeStamp {
		sequence = (sequence + 1) & SEQUENCE_MASK
		if sequence == 0 {
			/** 该毫秒内4095个已经完成，到下一毫秒 **/
			for timeStamp == lastTimeStamp {
				timeStamp = getTimestampUint64()
			}
		}
	} else {
		sequence = 0
	}
	lastTimeStamp = timeStamp
	return (timeStamp<<TIMESTAMP_OFFSET)&uint64(0x7fffffffffc00000) | (workerid<<WORKERID_OFFSET)&uint64(0x3ff000) | sequence
}

func getTimestampUint64() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}
