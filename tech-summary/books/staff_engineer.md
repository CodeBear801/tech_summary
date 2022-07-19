# Staff Engineer

[website](https://staffeng.com/book) 

## Work on what matters 

充分運用你為數不多的工作時間 特別是當你的人生有了其他除了工作以外的事情需要處理


<img src="https://user-images.githubusercontent.com/16873751/147709477-93376ce2-464e-4b75-b7bc-dc06dd10fbe5.png" alt="staff_engineer_p1" width="600"/>  

+ find low hangling fruits
+ stop snacking(吃零食) -> 花很少時間可以有很少結果的工作
+ stop preening(爱慕虚荣) -> low impact & high visibility work
  + short term plan(preening works for short term) vs long term plan
+ stop chasing ghost(不要被过去的经验蒙蔽)
  + 右下角那塊high effort low impact那塊 不會說全公司都沒有人做 通常是交給剛進來的senior(XD) 如果那個人是你 你必須要有個底線 知道這些東西是你進來前沒人要做的東西 不要花太多時間
+ Existential issues(寻找存在的问题)
  + 你應該去找 你們組織現在正在經歷什麼樣的risk或是問題
  + 去找那些最嚴重的問題 如果這些問題沒解決 其他事就無關緊要了 的那些問題
+ Folster growth(促进增长)
  + Hiring 是很重要的一環 包含onboarding, mentoring, coaching 這些很容易被公司忽略的問題

### Edit

有很多的專案 只跟成功差距不遠 就只差一個小小的改變 比如說只要一個小修改就可以解鎖新機會 或是一個討論就可以達到共識

Staff engineer會去找這些被需要的”改變” 來改變組上專案的approach 要做到這些 你必須要善用你的

1.組織上的特權 (organizational privilege)  
2.你在公司內部建立的關係  
3.你從過去專案經驗學到的ability to see around the corner  

利用這些你的知識 去了解組上的專案 去改變組上的專案

### Finish things

你會發現很多有天份的工程師在他的生涯初期 不了解怎麼應對專案的unexpected surprise 進而去rescope專案

畢竟一個專案要被recognize還是必須要完整的finish 把你的專業度發揮在幫助他們完成最後一哩路 比如說把六個月的口號轉換成兩個禮拜的sprint 能給組上帶來的收穫非常非常大

### What only you can

最後一個分類 就是去做那些只有你能夠完成的事情

雖然有很多工作 是你做得比其他人都好的 但更重要的是那些如果你不做就沒有人會做的工作

這可能是兩種工作的交集

1.你非常擅長 做得特別好的工作  
2.你真正在乎的工作  

Summary: 不断的提高自己的软技能

## Write Engineering strategy

去領導你們的組織針對公司business objective的方法 從architecture, technology selection和organizational structure的角度

### When and Why

Strategy是一個工具 這個工具可以讓團隊前進的更快更有信心 這個工具可以讓團隊的所有人在意見分歧的時候快速釐清方向  
-> 如果你發現你們組上三番兩次的在同一件事情上沒有結論 這就是時候去寫一個strategy  

如果你發現未來太過模糊 讓大家不知道應該投資哪些值得做的事情 這就是時候去寫一個vision  

如果你們組上都沒有這些問題 那你應該先去忙別的事 以後再回來處理strategy和vision的問題  


### Write five design docs

[Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/) | [How to write AWESOME TECH specs from lyft](https://eng.lyft.com/awesome-tech-specs-86eea8e45bb9) | [Technical Decision-Making and Alignment in a Remote Culture](https://multithreaded.stitchfix.com/blog/2020/12/07/remote-decision-making/)


### 如何写文档

一個好的RFC應該
1.描述一個明確的問題(specific problem)  
2.survey各種可能的解法  
3.解釋不同解法的細節  

當你不知道某一個專案是否需要值得寫一個RFC時 某些法則可以拿來檢視

1.當這個專案會被很多其他未來的專案所使用  
2.當這個專案會對於你的使用者有著重大影響  
3.當這個專案需要一個月以上的engineer effort  

如剛剛所說 五個design doc就會是一個寫好strategy的元素 因為design doc有著strategy缺乏的東西 - **detailed specifics grounded in reality**

寫好design doc的好建議
- 1.Start from the problem
  - Problem statement寫得越清楚 solution就會越清楚 當你覺得解法不那麼顯而易見時 回頭看看你的問題是不是定義的不夠好
  - 當你真的卡在定義問題的時候 找五個從沒看過這個問題的工程師 問他們what is missing 你會發現旁觀者清
- 2.Keep the template simple
  - 只在風險度大的專案中堅持寫下所有細節
- 3.Gather and review together, write alone
  - 通常你不會一個人就了解所有的相關知識和細節 所以我建議當你在深入研究某個主題之前 先跟周圍的人收集feedback 特別是那些未來需要依賴你的output of design doc的人
  - Gather perspectives widely but write alone
  - 然而 收集完feedback之後 還是要記得自己把RFC寫完 因為每個人都認爲自己是個好writer 但事實上要所有人一起把所有想法組織起來是很困難的 最好就是由一個人寫 然後讓其他人review
- Prefer good over perfect
  - 建議大家 與其花很多時間寫到你認為完美再給大家看 不如先寫個ok的版本趕快先分享
  - 別忘了 當你在看別人的design的時候 不要一開始就預期別人寫的就應該跟你心中的最好解法一樣 特別是當你越資深 你就要越注意不要對別人寫的design太過toxic 要有耐心的給別人建議 讓別人更好
  - 寫一個好的design需要時間練習 如果你真的想進步的話 作者的建議是 當你實做完這個專案之後 回頭重新讀你的design 找出那些你當初忽略的地方 或是實作起來跟當初設計的時候不同的地方 然後找出當初忽略的原因

### Synthesize those five design docs into a strategy

- 當你們的組織中有五個design doc之後 坐下來好好的看一下這五個design doc 找尋這些文件中 有爭議的部分 衝突的部分
- 好的strategy會引導大家如何從tradeoff中選擇 並提供很好的解釋和context
- 不好的strategy就只是告訴大家應該做什麼不應該做什麼 但卻沒告訴大家為什麼
-  沒有context的話 你的strategy就會越來越難以理解 而當context開始轉換的時候 大家也不知道strategy應該相對應的如何轉換

如何生成strategy



什麼事情值得去寫stragety?
- “什麼時候我們需要寫design doc” 這個問題本身就可以寫一個strategy
- “在哪些use case的情況下我們要用哪種database” 這個問題也可以寫一個strategy
- “我們什麼時候該從monolith migrate to services” 這個問題也可以寫一個strategy


### Extrapolate five strategies into a vision

當你有越來越多的strategies 你就會發現很難釐清各個strategy之間的關聯互動或因果關係。也許某個strategy要我們少點自己寫的software 多用cloud上的solution， 但另一個strategy也許是說 database的複雜性應該盡可能的被降低  
那如果團隊遇到個選擇 發現一個複雜度高的database在cloud上 而複雜度低的不在cloud上 那該怎麼辦？  
拿出你們最近的五個strategy 把它們的tradeoff extrapolate到兩三年之後 你會越來越容易去解決那些彼此的衝突 進而寫成engineering vision  
一個好的vision會讓大家很好理解目前存在的每個strategy之間的關係 也讓未來的新strategy更加經的起推敲

幾個寫vision的建議

- 1.Write two to three years out
  - 小組 大組 甚至公司都改變得太快 如果想得太遠的話 會讓人失去目標而憂心重重 但如果想得太近也不行 如果一個strategy六個月後就會失效 那也沒啥用 想想你們六個月內是能寫出幾個strategy? 試著把目標定在2-3年之後
- 2.Ground in your business and your users
  - 一個好的vision應該不只跟business息息相關 跟使用者也應該息息相關
- 3.Be optimistic rather than audacious
  - vision應該要ambitious而不是audacious. 應該要考慮有限資源的情況下的best possible version
- 4.Stay concrete and specific
  - 越specific的vision越有用 generic的vision很容易被大家同意 但組織裡真正遇到問題的時候 很難幫忙得出有用結論
- 5.Keep it one to two pages long
  - 大家不喜歡讀冗長的文件 如果你光vision就寫了五六頁 那很容易大家都不會讀完
  - Vision文件本身小一些 提供一些reference讓那些有興趣了解更多的人去dive deep
  - 寫完之後呢 你也不要太爽的就急著分享給組織裡的所有人


怎麼判斷vision文档的好坏： 找一個兩年前的design doc 再找一個你的vision推出來之後才被寫出的新design doc 如果感覺起來有明顯的進步 那麼你的vision就很好
