package reflect

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_proxy(defineClass0, "defineClass0", "(Ljava/lang/ClassLoader;Ljava/lang/String;[BII)Ljava/lang/Class;")
}

func _proxy(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/reflect/Proxy", name, desc, method)
}

// private static native Class<?> defineClass0(ClassLoader loader, String name,
//                                             byte[] b, int off, int len);
// (Ljava/lang/ClassLoader;Ljava/lang/String;[BII)Ljava/lang/Class;
func defineClass0(frame *rtda.Frame) {
	stack := frame.OperandStack()
	if stack.IsEmpty() {
		_loadClass(frame)
	}

	// init class
	class := stack.TopRef(0).Extra().(*rtc.Class)
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
	}
}

func _loadClass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//loader := vars.GetRef(0)
	nameObj := vars.GetRef(1)
	byteArr := vars.GetRef(2)
	off := vars.GetInt(3)
	_len := vars.GetInt(4)

	name := rtda.GoString(nameObj)
	name = rtc.DotToSlash(name)
	data := byteArr.GoBytes()
	data = data[off : off+_len]

	// todo
	class := frame.ClassLoader().DefineClass(name, data)
	stack := frame.OperandStack()
	stack.PushRef(class.JClass())
}
