package condition

import "lafitd/expression/value"



//第一层map的key对应了leftValue的type.
//第二层map的key是操作类型.
//第二层map的value表示的rightValue的合法类型

//对照如下元数据可导出:
// true equal 1 不合法  true equal 0 不合法  true equal false 合法  true equal true 合法 ...其余示例省略
// 1 equal 1 合法 1 equal 1.33 合法 1 equal "324" 不合法 ...其余示例忽略
//其余示例忽略
var standard = map[string]map[string][]string {
	value.BoolType: {
		Equal: {value.BoolType},
	},
	value.IntType: {
		Equal: {value.IntType, value.FloatType},
		GreaterThan: {value.IntType, value.FloatType},
		LesserThan: {value.IntType, value.FloatType},
	},
	value.StringType: {
		Equal: {value.StringType},
	},
	value.FloatType: {
		Equal: {value.FloatType, value.IntType},
		GreaterThan: {value.FloatType, value.IntType},
		LesserThan: {value.FloatType, value.IntType},
	},
	value.CollectionStringType: {
		Equal: {value.CollectionStringType},
		Contain: {value.StringType},
	},
	value.CollectionIntType: {
		Equal: {value.CollectionIntType, value.CollectionFloatType},
		Contain: {value.IntType, value.FloatType},
	},
	value.CollectionFloatType: {
		Equal: {value.CollectionIntType, value.CollectionFloatType},
		Contain: {value.IntType, value.FloatType},
	},
}

func GetStandard() map[string]map[string][]string  {
	return standard
}