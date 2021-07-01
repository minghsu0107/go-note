package chain

func ExampleChain() {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manager = c1

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
	// Output:
	// Project manager permit bob 400 fee request
	// Dep manager permit tom 1400 fee request
	// General manager permit ada 10000 fee request
	// Project manager don't permit floar 400 fee request

}

/*
職責鏈的最大優點是可以將所有節點解藕，無論誰是請求者誰是接收者，在鏈結中可以任意的拆解或是新增節點，
工程師也不必在意每個節點實際要做的內容，更可以手動指定起始節點，
假設我們今天客戶流程更改，我們就只需要更改鏈結內節點順序就好

但這樣的模式仍然有些缺點，首先是我們沒辦法保證請求一定會被處理，過程中的節點可能都會直接跳過這個請求，
為避免這狀況發生，我們可以在鏈尾增加一個處理這種狀況的節點。
另外一點也要注意的是，也許某些請求中大部分的節點都只是傳遞至下一節點的功能而已，
從效能方面考量我們要避免過長的職責鏈所帶來的效能耗損。
*/