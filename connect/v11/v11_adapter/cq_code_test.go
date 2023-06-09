package v11_adapter

import (
	"log"
	"testing"
)

func TestCQCodeToSegment11(t *testing.T) {
	log.Println(CQCodeToSegment11("[CQ:at,qq=123456789]"))
}

func TestCQCodeToSegment12(t *testing.T) {
	log.Println(Segment11To12(CQCodeToSegment11("[CQ:at,qq=all]")))
}

func TestMessage11ToSegment12(t *testing.T) {
	t.Log(Message11ToSegment12("测试[CQ:at,qq=all][CQ:at,qq=123456789]123dccdsivd[CQ:face,id=12]"))
}
