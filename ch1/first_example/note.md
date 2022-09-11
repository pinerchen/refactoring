# Refactor Note

## 尚未看書我覺得可以改善的地方
1. 拆出 Tragedy/ Comedy PlayType 成物件，定義各自的 thisAmount, function 去處理超出預設 Audience 數量。
2. 計算 volumeCredits 也可以獨立出去，同上依 PlayType 分別處理。

## 書中教的方法

當需要加東西，但看到程式碼不知道從何處加進去，就表示結構不友善，不明確；就像是百貨公司缺少樓層規劃圖，我們需要讓自己和其他工程師能輕鬆地知道要從哪裡加新功能。

試想：first example 需要再加入 invoice printed in HTML

1. 加入結構, such as: set of functions, classes

# refactor 方向
- 處理 print invoice 方法要獨立出去
- playType 的處理要能輕易擴展
- 計算 amount, credit 要能動態輕易隨商業變動

# first example refactoring steps
1. identify points that separate different parts ov the overall behavior.

首先看到 switch-case 那段
光看程式碼無法快速知道做了什麼，用人腦記憶的資訊是種惡疾，需要做到只需快速看程式碼，就能知道在做什麼。

!需要注意 Extract Function 會影響到的變數 scope 

2. small changes each steps, and test it right away.
3. extracted method 內部變數也可以 refactor 一下
作者接受 Kent Beck 的建議，習慣依據變數類型在命名上加上不定冠詞，區分動態或特定的類型。
4. 盡量排除 function 內的 local variable，如果 refactor extract method，多呼叫幾次都比讓一次性變數待在 scope 內好。first example 的 thisAmount, format function 都是作者想辦法移除的 local variables
5. Iteration 內做了許多事 (計算 totalAmount, 計算 volumeCredits)，將他們分兩區塊，並且把對應的變數移到附近(Slide Statement)

## 用到了哪些方法
* Extract Method
* Replace Temp with Query
* Inline Variable
* Split Loop
* Slide Statement



