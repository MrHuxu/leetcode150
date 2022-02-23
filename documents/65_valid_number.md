## 题意

给定一个字符串 `s`, 判断它是不是一个合法的数字.

> "0" => true  
" 0.1 " => true  
"abc" => false  
"1 a" => false  
"2e10" => true  
" -90e3   " => true  
" 1e" => false  
"e3" => false  
" 6e-1" => true  
" 99e2.5 " => false  
"53.5e93" => true  
" --6 " => false  
"-+3" => false  
"95a54e53" => false

## 解答

根据科学记数法规则, 用 e 将整个字符串分成两个部分, 再分别判断两个部分是否为合法数字即可.

最好不要用 `strings.Trim` 和 `strings.Split` 方法, 而是直接在代码里使用 `str[startIdx:endIdx]` 的方式进行截取和判断, 因为 Trim 和 Split 都是会申请新的内存, 而直接截取字符串是不会的.