package utils
/**
compare, apply to all var type
 */
import "reflect"

func Compare(a interface{}, b interface{}) int {
  // check type of the var is the same?
  aTyp := reflect.TypeOf(a).String()
  bTyp := reflect.TypeOf(b).String()

  if aTyp != bTyp {
    panic("cannot compare different type params")
  }

  switch a.(type) {
  case int:
    if a.(int) > b.(int) {
      return 1
    } else if a.(int) < b.(int) {
      return -1
    } else {
      return 0
    }

  case string:
    if a.(string) > b.(string) {
      return 1
    }  else if a.(string) < b.(string) {
      return -1
    } else {
      return 0
    }

  case float64:
    if a.(float64) > b.(float64) {
      return 1
    } else if a.(float64) < b.(float64) {
      return -1
    } else {
      return 0
    }

  default:
    panic("unsupport type params")
  }
}






