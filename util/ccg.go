package util

import (
	"ccg/src/pb"
	"reflect"

	"github.com/topfreegames/pitaya/v2/serialize"
)

func WrapWithCcgMsg(s serialize.Serializer, v interface{}) (interface{}, error) {
	bb, err := SerializeOrRaw(s, v)
	if err != nil {
		return nil, err
	}
	// logger.Log.Debugf("反射出来的名字：%s", reflect.TypeOf(ret).Elem().Name())
	v = &pb.TcgMsg{
		LogicType: reflect.TypeOf(v).Elem().Name(),
		LogicData: bb,
	}
	return v, nil
}
