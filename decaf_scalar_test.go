package ed448

import (
	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_NewScalar(c *C) {
	a := [fieldBytes]byte{
		0x25, 0x8a, 0x52, 0x63, 0xd9, 0xf0, 0xfa, 0xad,
		0x9d, 0x50, 0x40, 0x8a, 0xf0, 0x76, 0x66, 0xe3,
		0x3d, 0xc2, 0x86, 0x1b, 0x01, 0x54, 0x18, 0xb8,
		0x1b, 0x3b, 0x76, 0xcd, 0x55, 0x18, 0xa2, 0xfd,
		0xf1, 0xf2, 0x64, 0xee, 0xae, 0xae, 0xc5, 0xe7,
		0x68, 0xa4, 0x2e, 0xde, 0x76, 0x60, 0xe6, 0x4a,
		0x51, 0x12, 0xb1, 0x35, 0x3d, 0xac, 0x04, 0x08,
	}

	sc := NewDecafScalar(a)

	expected := &decafScalar{
		0x63528a25, 0xadfaf0d9,
		0x8a40509d, 0xe36676f0,
		0x1b86c23d, 0xb8185401,
		0xcd763b1b, 0xfda21855,
		0xee64f2f1, 0xe7c5aeae,
		0xde2ea468, 0x4ae66076,
		0x35b11251, 0x0804ac3d,
	}

	c.Assert(sc, DeepEquals, expected)
}

func (s *Ed448Suite) Test_ScalarAddition(c *C) {
	s1 := &decafScalar{
		0x529eec33, 0x721cf5b5,
		0xc8e9c2ab, 0x7a4cf635,
		0x44a725bf, 0xeec492d9,
		0x0cd77058, 0x00000002,
	}
	s2 := &decafScalar{0x00000001}
	expected := decafScalar{
		0x529eec34, 0x721cf5b5,
		0xc8e9c2ab, 0x7a4cf635,
		0x44a725bf, 0xeec492d9,
		0x0cd77058, 0x00000002,
	}
	out := decafScalar{}
	out.scalarAdd(s1, s2)
	c.Assert(out, DeepEquals, expected)
}

func (s *Ed448Suite) Test_ScalarHalve(c *C) {
	expected := decafScalar{6}
	s1 := &decafScalar{12}
	s2 := &decafScalar{4}
	out := decafScalar{}
	out.scalarHalve(s1, s2)
	c.Assert(out, DeepEquals, expected)
}

func (s *Ed448Suite) Test_littleScalarMul_Identity(c *C) {
	x := &decafScalar{
		0xd013f18b, 0xa03bc31f,
		0xa5586c00, 0x5269ccea,
		0x80becb3f, 0x38058556,
		0x736c3c5b, 0x07909887,
		0x87190ede, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	y := &decafScalar{0x00000001}

	expected := &decafScalar{
		0xf19fb32f, 0x62bc6ae6,
		0xed626086, 0x0e2d81d7,
		0x7a83d54b, 0x38e73799,
		0x485ad3d6, 0x45399c9e,
		0x824b12d9, 0x5ae842c9,
		0x5ca5b606, 0x3c0978b3,
		0x893b4262, 0x22c93812,
	}

	out := &decafScalar{}
	out.montgomeryMultiply(x, y)
	c.Assert(out, DeepEquals, expected)
	out.montgomeryMultiply(out, scalarR2)
	c.Assert(out, DeepEquals, x)
}

func (s *Ed448Suite) Test_littleScalarMul_Zero(c *C) {
	x := &decafScalar{
		0xd013f18b, 0xa03bc31f,
		0xa5586c00, 0x5269ccea,
		0x80becb3f, 0x38058556,
		0x736c3c5b, 0x07909887,
		0x87190ede, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	y := &decafScalar{}

	out := &decafScalar{}
	out.montgomeryMultiply(x, y)
	c.Assert(out, DeepEquals, y)
}

func (s *Ed448Suite) Test_littleScalarMul_fullMultiplication(c *C) {
	x := &decafScalar{
		0xffb823a3, 0xc96a3c35,
		0x7f8ed27d, 0x087b8fb9,
		0x1d9ac30a, 0x74d65764,
		0xc0be082e, 0xa8cb0ae8,
		0xa8fa552b, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	y := &decafScalar{
		0xd8bedc42, 0x686eb329,
		0xe416b899, 0x17aa6d9b,
		0x1e30b38b, 0x188c6b1a,
		0xd099595b, 0xbc343bcb,
		0x1adaa0e7, 0x24e8d499,
		0x8e59b308, 0x0a92de2d,
		0xcae1cb68, 0x16c5450a,
	}

	expected := decafScalar{
		0x14aec10b, 0x426d3399,
		0x3f79af9e, 0xb1f67159,
		0x6aa5e214, 0x33819c2b,
		0x19c30a89, 0x480bdc8b,
		0x7b3e1c0f, 0x5e01dfc8,
		0x9414037f, 0x345954ce,
		0x611e7191, 0x19381160,
	}

	out := decafScalar{}
	out.montgomeryMultiply(x, y)
	c.Assert(out, DeepEquals, expected)
}

func (s *Ed448Suite) Test_Add(c *C) {
	one := &decafScalar{0x1}
	two := &decafScalar{0x2}
	three := &decafScalar{0x3}

	result := &decafScalar{}
	result.Add(one, two)

	c.Assert(result, DeepEquals, three)
}

func (s *Ed448Suite) Test_Sub(c *C) {
	twelve := &decafScalar{0xc}
	thirteen := &decafScalar{0xd}
	one := &decafScalar{0x1}

	result := &decafScalar{}
	result.Sub(thirteen, twelve)

	c.Assert(result, DeepEquals, one)
}

func (s *Ed448Suite) Test_Mul(c *C) {
	x := &decafScalar{
		0xffb823a3, 0xc96a3c35,
		0x7f8ed27d, 0x087b8fb9,
		0x1d9ac30a, 0x74d65764,
		0xc0be082e, 0xa8cb0ae8,
		0xa8fa552b, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}

	y := &decafScalar{
		0xd8bedc42, 0x686eb329,
		0xe416b899, 0x17aa6d9b,
		0x1e30b38b, 0x188c6b1a,
		0xd099595b, 0xbc343bcb,
		0x1adaa0e7, 0x24e8d499,
		0x8e59b308, 0x0a92de2d,
		0xcae1cb68, 0x16c5450a,
	}

	expected := &decafScalar{
		0xa18d010a, 0x1f5b3197,
		0x994c9c2b, 0x6abd26f5,
		0x08a3a0e4, 0x36a14920,
		0x74e9335f, 0x07bcd931,
		0xf2d89c1e, 0xb9036ff6,
		0x203d424b, 0xfccd61b3,
		0x4ca389ed, 0x31e055c1,
	}
	x.Mul(x, y)
	c.Assert(x, DeepEquals, expected)
}

func (s *Ed448Suite) Test_Copy(c *C) {
	expected := &decafScalar{
		0xffb823a3, 0xc96a3c35,
		0x7f8ed27d, 0x087b8fb9,
		0x1d9ac30a, 0x74d65764,
		0xc0be082e, 0xa8cb0ae8,
		0xa8fa552b, 0x2aae8688,
		0x2c3dc273, 0x47cf8cac,
		0x3b089f07, 0x1e63e807,
	}
	x := expected.Copy()
	c.Assert(x, DeepEquals, expected)
}

func (s *Ed448Suite) Test_ScalarDecode(c *C) {
	x := &decafScalar{0x00}

	expected := &decafScalar{
		0x2a1c3d02, 0x12f970e8,
		0x41d97de7, 0x6a547b38,
		0xdaa8c88e, 0x9f299b75,
		0x01075c7b, 0x3b874ad9,
		0xe1c0b914, 0xc8bd0b68,
		0xc3f34776, 0x2f2d9082,
		0x4b75d258, 0x34a8bc39,
	}

	buf := []byte{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72, 0x36,
		0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d, 0xc1, 0x8b,
		0x1e, 0xff, 0x7e, 0x89, 0xbf, 0x76, 0x78, 0x63,
		0x65, 0x80, 0xd1, 0x7d, 0xd8, 0x4a, 0x87, 0x3b,
		0x14, 0xb9, 0xc0, 0xe1, 0x68, 0x0b, 0xbd, 0xc8,
		0x76, 0x47, 0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f,
		0x58, 0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	}

	ok := x.decode(buf)
	c.Assert(x, DeepEquals, expected)
	c.Assert(ok, Equals, word(0x00))

	x1 := &decafScalar{0x00}

	expected1 := &decafScalar{
		0x08ed77fa, 0x85c49151,
		0xa0dd2874, 0x7188bced,
		0x9a34c3bd, 0xab1aeece,
		0xea37a24c, 0x8dd2eab4,
		0x8610f125, 0xb3eb60c0,
		0x8aaa9ab0, 0xf19e004b,
		0x78fe2593, 0x3aa1dd0f,
	}

	buf1 := []byte{
		0xfa, 0x77, 0xed, 0x08, 0x51, 0x91, 0xc4, 0x85,
		0x74, 0x28, 0xdd, 0xa0, 0xed, 0xbc, 0x88, 0x71,
		0xbd, 0xc3, 0x34, 0x9a, 0xce, 0xee, 0x1a, 0xab,
		0x4c, 0xa2, 0x37, 0xea, 0xb4, 0xea, 0xd2, 0x8d,
		0x25, 0xf1, 0x10, 0x86, 0xc0, 0x60, 0xeb, 0xb3,
		0xb0, 0x9a, 0xaa, 0x8a, 0x4b, 0x00, 0x9e, 0xf1,
		0x93, 0x25, 0xfe, 0x78, 0x0f, 0xdd, 0xa1, 0x3a,
	}

	ok2 := x1.decode(buf1)
	c.Assert(x1, DeepEquals, expected1)
	c.Assert(ok2, Equals, word(0xffffffff))
}
