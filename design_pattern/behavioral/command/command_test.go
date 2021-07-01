package command

func ExampleCommand() {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()
	box2.RunAll()
	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
	// system rebooting
	// system starting
}

/*
命令模式有幾個優點：

1. 它能較容易的設計一個命令序列。
2. 在需要的狀況下，可以較容易的將命令記入日誌。
3. 允許接收請求的一方決定是否要否決請求。
4. 可以容易的實現對請求的取消和重做。
5. 由於加進新的具體命令類別不影響其他類別，因此增加新的具體命令類別很容易。

最後、最大的優點是將請求的物件和執行的物件分開。
*/
