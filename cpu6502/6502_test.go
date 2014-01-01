package cpu6502

import (
	"fmt"
	"testing"
)

type CpuTest struct {
	Name string
	Mem  []byte
	End  Cpu
}

var CpuTests = []CpuTest{
	CpuTest{
		Name: "load, set",
		Mem:  []byte{0xa9, 0x01, 0x8d, 0x00, 0x02, 0xa9, 0x05, 0x8d, 0x01, 0x02, 0xa9, 0x08, 0x8d, 0x02, 0x02},
		End: Cpu{
			A:  0x08,
			S:  0xff,
			PC: 0x0610,
			P:  0x30,
		},
	},
	CpuTest{
		Name: "load, transfer, increment, add",
		Mem:  []byte{0xa9, 0xc0, 0xaa, 0xe8, 0x69, 0xc4, 0x00},
		End: Cpu{
			A:  0x84,
			X:  0xc1,
			S:  0xff,
			PC: 0x0607,
			P:  0xb1,
		},
	},
	CpuTest{
		Name: "bne",
		Mem:  []byte{0xa2, 0x08, 0xca, 0x8e, 0x00, 0x02, 0xe0, 0x03, 0xd0, 0xf8, 0x8e, 0x01, 0x02, 0x00},
		End: Cpu{
			X:  0x03,
			S:  0xff,
			PC: 0x060e,
			P:  0x33,
		},
	},
	CpuTest{
		Name: "relative",
		Mem:  []byte{0xa9, 0x01, 0xc9, 0x02, 0xd0, 0x02, 0x85, 0x22, 0x00},
		End: Cpu{
			A:  0x01,
			S:  0xff,
			PC: 0x0609,
			P:  0xb0,
		},
	},
	CpuTest{
		Name: "indirect",
		Mem:  []byte{0xa9, 0x01, 0x85, 0xf0, 0xa9, 0xcc, 0x85, 0xf1, 0x6c, 0xf0, 0x00},
		End: Cpu{
			A:  0xcc,
			S:  0xff,
			PC: 0xcc02,
			P:  0xb0,
		},
	},
	CpuTest{
		Name: "indexed indirect",
		Mem:  []byte{0xa2, 0x01, 0xa9, 0x05, 0x85, 0x01, 0xa9, 0x06, 0x85, 0x02, 0xa0, 0x0a, 0x8c, 0x05, 0x06, 0xa1, 0x00},
		End: Cpu{
			A:  0x0a,
			X:  0x01,
			Y:  0x0a,
			S:  0xff,
			PC: 0x0612,
			P:  0x30,
		},
	},
}

func Test6502(t *testing.T) {
	for _, test := range CpuTests {
		c := New()
		copy(c.Mem[c.PC:], test.Mem)
		fmt.Println(test.Name)
		c.Run()
		if c.A != test.End.A ||
			c.X != test.End.X ||
			c.Y != test.End.Y ||
			c.S != test.End.S ||
			c.PC != test.End.PC ||
			c.P != test.End.P {
			t.Fatalf("bad cpu state %s, got:\n%sexpected:\n%s", test.Name, c, &test.End)
		}
	}
}
