package errors

//
//Well, we are handling the error twice by logging it first and then by returning it to the caller of this function.
//
//Probably one of your team colleagues will use this function and when an error returns, he will log the error again. Then an error nightmare occurs in the system log

// 1. log 被处理了两次
// 2. 其他人又又可能将这个函数再包一层，然后再 log
//func someFunc() (Result, error) {
//	result, err := repository.Find(id)
//	if err != nil {
//		log.Errof(err)
//		return Result{}, err
//	}
//	return result, nil
//}

// 原则与做法
// Provide a good error stack trace   提供调用栈
// Log the error (e.g. the web infrastructure layer) 在基础库层，进行日志打印
// Provide a contextual error information to the user when necessary. (e.g. The provided email is not in the right format)
//